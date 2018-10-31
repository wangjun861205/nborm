package nborm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
)

//Pager is for pagerate
type Pager [2]int

//NewPager create Pager
func NewPager(numPerPage, pageNum int) *Pager {
	if numPerPage <= 0 || pageNum <= 0 {
		return nil
	}
	p := Pager([2]int{numPerPage, pageNum})
	return &p
}

func (p *Pager) toSQL() string {
	if p == nil {
		return ""
	}
	offset := (*p)[0] * ((*p)[1] - 1)
	return fmt.Sprintf("LIMIT %d, %d", offset, p[0])
}

//NextPage set to next page
func (p *Pager) NextPage() {
	(*p)[1]++
}

//PrevPage set to previous page
func (p *Pager) PrevPage() (ok bool) {
	if (*p)[1] != 0 {
		(*p)[1]--
		return true
	}
	return false
}

//Index index to one page
func (p *Pager) Index(i int) {
	(*p)[1] = i
}

//Order sort order
type Order struct {
	Field   Field
	Reverse bool
}

func NewOrder(field Field, reverse bool) *Order {
	return &Order{field, reverse}
}

//OrderBy create Sorter
func NewSorter(orders ...*Order) *Sorter {
	l := make([]string, len(orders))
	for i, order := range orders {
		var o string
		if order.Reverse {
			o = "DESC"
		} else {
			o = "ASC"
		}
		l[i] = fmt.Sprintf("%s.%s.%s %s", order.Field.superModel().DB(), order.Field.superModel().Tab(), order.Field.columnName(), o)
	}
	s := Sorter(fmt.Sprintf("ORDER BY %s", strings.Join(l, ", ")))
	return &s
}

//Sorter information for sql order by operation
type Sorter string

func (s *Sorter) toSQL() string {
	if s == nil {
		return ""
	}
	return string(*s)
}

func scanRows(l ModelList, rows *sql.Rows) error {
	defer rows.Close()
	for rows.Next() {
		m := l.New()
		fields := m.Fields()
		scanList := make([]interface{}, len(fields))
		for i, f := range fields {
			scanList[i] = f
		}
		err := rows.Scan(scanList...)
		if err != nil {
			return err
		}
		m.SetSync(true)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func scanRowsContext(ctx context.Context, l ModelList, rows *sql.Rows) error {
	defer rows.Close()
	for rows.Next() {
		select {
		case <-ctx.Done():
			return nil
		default:
			m := l.New()
			fields := m.Fields()
			scanList := make([]interface{}, len(fields))
			for i, f := range fields {
				scanList[i] = f
			}
			err := rows.Scan(scanList...)
			if err != nil {
				return err
			}
			m.SetSync(true)
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func scanRow(m Model, row *sql.Row) error {
	fields := m.Fields()
	scanList := make([]interface{}, len(fields))
	for i, f := range fields {
		scanList[i] = f
	}
	err := row.Scan(scanList...)
	if err != nil {
		return err
	}
	m.SetSync(true)
	return nil
}

func scanRowContext(ctx context.Context, m Model, row *sql.Row) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		fields := m.Fields()
		scanList := make([]interface{}, len(fields))
		for i, f := range fields {
			scanList[i] = f
		}
		err := row.Scan(scanList...)
		if err != nil {
			return err
		}
		m.SetSync(true)
		return nil
	}
}

func queryAndScan(tab interface{}, stmtStr string, valList []interface{}) error {
	switch obj := tab.(type) {
	case Model:
		db := dbMap[obj.DB()]
		row := db.QueryRow(stmtStr, valList...)
		return scanRow(obj, row)
	case ModelList:
		db := dbMap[obj.Model().DB()]
		rows, err := db.Query(stmtStr, valList...)
		if err != nil {
			return err
		}
		return scanRows(obj, rows)
	default:
		return fmt.Errorf("nborm error: unsupported type (%T)", tab)
	}
}

func queryAndScanContext(ctx context.Context, tab interface{}, stmtStr string, valList []interface{}) error {
	switch obj := tab.(type) {
	case Model:
		db := dbMap[obj.DB()]
		row := db.QueryRowContext(ctx, stmtStr, valList...)
		return scanRowContext(ctx, obj, row)
	case ModelList:
		db := dbMap[obj.Model().DB()]
		rows, err := db.QueryContext(ctx, stmtStr, valList...)
		if err != nil {
			return err
		}
		return scanRowsContext(ctx, obj, rows)
	default:
		return fmt.Errorf("nborm error: unsupported type (%T)", tab)
	}
}

func queryAndScanWithNum(tab interface{}, stmtStr string, valList []interface{}) (int, error) {
	if !strings.Contains(stmtStr, "SQL_CALC_FOUND_ROWS") {
		return -1, fmt.Errorf("the statement (%s) does not contains 'SQL_CALC_FOUND_ROWS'", stmtStr)
	}
	switch obj := tab.(type) {
	case Model:
		db := dbMap[obj.DB()]
		row := db.QueryRow(stmtStr, valList...)
		if err := scanRow(obj, row); err != nil {
			return -1, err
		}
		return getFoundRows(db)
	case ModelList:
		db := dbMap[obj.Model().DB()]
		rows, err := db.Query(stmtStr, valList...)
		if err != nil {
			return -1, err
		}
		if err := scanRows(obj, rows); err != nil {
			return -1, err
		}
		return getFoundRows(db)
	default:
		return -1, fmt.Errorf("nborm error: unsupported type (%T)", tab)
	}
}

func queryAndScanWithNumContext(ctx context.Context, tab interface{}, stmtStr string, valList []interface{}) (int, error) {
	if !strings.Contains(stmtStr, "SQL_CALC_FOUND_ROWS") {
		return -1, fmt.Errorf("the statement (%s) does not contains 'SQL_CALC_FOUND_ROWS'", stmtStr)
	}
	switch obj := tab.(type) {
	case Model:
		db := dbMap[obj.DB()]
		row := db.QueryRowContext(ctx, stmtStr, valList...)
		if err := scanRowContext(ctx, obj, row); err != nil {
			return -1, err
		}
		return getFoundRowsContext(ctx, db)
	case ModelList:
		db := dbMap[obj.Model().DB()]
		rows, err := db.QueryContext(ctx, stmtStr, valList...)
		if err != nil {
			return -1, err
		}
		if err := scanRowsContext(ctx, obj, rows); err != nil {
			return -1, err
		}
		return getFoundRowsContext(ctx, db)
	default:
		return -1, fmt.Errorf("nborm error: unsupported type (%T)", tab)
	}
}

func getPk(m Model) (pk Field, others []Field) {
	info := getTableCache(m)
	fields := m.Fields()
	pk = fields[info.pk]
	others = append(fields[:info.pk], fields[info.pk+1:]...)
	return
}

func getInc(m Model) (inc Field, others []Field) {
	info := getTableCache(m)
	fields := m.Fields()
	inc = fields[info.inc]
	others = append(fields[:info.inc], fields[info.inc+1:]...)
	return
}

func getUni(m Model) (unis []Field, others []Field) {
	fields := m.Fields()
	for _, f := range fields {
		if f.IsPk() || f.IsInc() || f.IsUni() {
			unis = append(unis, f)
		} else {
			others = append(others, f)
		}
	}
	return
}

func filterValid(fields []Field) (validFields []Field) {
	for _, f := range fields {
		if f.IsValid() {
			validFields = append(validFields, f)
		}
	}
	return
}

func getByName(m Model, name string) Field {
	info := getTableCache(m)
	fields := m.Fields()
	return fields[info.fieldMap[name]]
}

type sortObj struct {
	ModelList
	funcs []func(Model, Model) int
}

//Less implement sort.Sorter interface
func (o *sortObj) Less(i, j int) bool {
	iMod, jMod := o.ModelList.Index(i), o.ModelList.Index(j)
	for _, f := range o.funcs {
		v := f(iMod, jMod)
		switch {
		case v < 0:
			return true
		case v > 0:
			return false
		default:
			continue
		}
	}
	return false
}

func filterList(l ModelList, f func(Model) bool) {
	var i int
	for i < l.Len() {
		if f(l.Index(i)) {
			l.Delete(i)
			continue
		}
		i++
	}
}

// func updateModel(m Model, values ...*UpdateValue) {
// 	for _, val := range values {
// 		getByName(m, val.column).SetByUpdateValue(val)
// 	}
// }

func genWhere(m Model) *Where {
	fields := m.Fields()
	info := getTableCache(m)
	switch {
	case info.inc != -1 && fields[info.inc].IsValid():
		return fields[info.inc].Where()
	case info.pk != -1 && fields[info.pk].IsValid():
		return fields[info.pk].Where()
	case len(info.unis) > 0:
		for _, index := range info.unis {
			if fields[index].IsValid() {
				return fields[index].Where()
			}
		}
		fallthrough
	default:
		var w *Where
		for _, f := range fields {
			if f.IsValid() {
				if w == nil {
					w = f.Where()
				} else {
					w.And(f.Where())
				}
			}
		}
		if w == nil {
			panic("nborm.genWhere() error: empty where")
		}
		return w
	}
}

func getFoundRows(db *sql.DB) (int, error) {
	var num int
	row := db.QueryRow("SELECT FOUND_ROWS()")
	err := row.Scan(&num)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func getFoundRowsContext(ctx context.Context, db *sql.DB) (int, error) {
	var num int
	row := db.QueryRowContext(ctx, "SELECT FOUND_ROWS()")
	err := row.Scan(&num)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func genSelect(tab interface{}, where *Where, sorter *Sorter, pager *Pager, withFoundRows bool, relations ...relation) (string, []interface{}) {
	var dbName, tabName string
	switch obj := tab.(type) {
	case Model:
		dbName, tabName = obj.DB(), obj.Tab()
	case ModelList:
		dbName, tabName = obj.Model().DB(), obj.Model().Tab()
	default:
		panic(fmt.Errorf("nborm error: unsupported type (%T)", tab))
	}
	if where == nil {
		if withFoundRows {
			if len(relations) == 0 {
				return fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s %s %s", dbName, tabName, sorter.toSQL(), pager.toSQL()), nil
			} else {
				joinList := make([]string, len(relations))
				for i, rel := range relations {
					joinList[i] = rel.joinClause()
				}
				return fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s %s %s %s", dbName, tabName, dbName, tabName,
					strings.Join(joinList, " "), sorter.toSQL(), pager.toSQL()), nil
			}
		} else {
			if len(relations) == 0 {
				return fmt.Sprintf("SELECT * FROM %s.%s %s %s", dbName, tabName, sorter.toSQL(), pager.toSQL()), nil
			} else {
				joinList := make([]string, len(relations))
				for i, rel := range relations {
					joinList[i] = rel.joinClause()
				}
				return fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s %s", dbName, tabName, dbName, tabName, strings.Join(joinList, " "),
					sorter.toSQL(), pager.toSQL()), nil
			}
		}
	} else {
		colStr, valList := where.toClause()
		if withFoundRows {
			if len(relations) == 0 {
				return fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s %s %s %s", dbName, tabName, colStr, sorter.toSQL(),
					pager.toSQL()), valList
			} else {
				joinList := make([]string, len(relations))
				for i, rel := range relations {
					joinList[i] = rel.joinClause()
				}
				return fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s %s %s %s %s", dbName, tabName, dbName, tabName,
					strings.Join(joinList, " "), colStr, sorter.toSQL(), pager.toSQL()), valList
			}
		} else {
			if len(relations) == 0 {
				return fmt.Sprintf("SELECT * FROM %s.%s %s %s %s", dbName, tabName, colStr, sorter.toSQL(), pager.toSQL()), valList
			} else {
				joinList := make([]string, len(relations))
				for i, rel := range relations {
					joinList[i] = rel.joinClause()
				}
				return fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s %s %s", dbName, tabName, dbName, tabName, strings.Join(joinList, " "),
					colStr, sorter.toSQL(), pager.toSQL()), valList
			}
		}
	}
}

func insertAndGetInc(m Model, update bool) error {
	db := dbMap[m.DB()]
	inc, others := getInc(m)
	others = filterValid(others)
	colList := make([]string, len(others))
	valList := make([]interface{}, len(others))
	updateList := make([]string, len(others))
	for i, f := range others {
		colList[i] = f.columnName()
		valList[i] = f.value()
		updateList[i] = f.columnName() + " = ?"
	}
	var res sql.Result
	var err error
	if update {
		stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", m.DB(), m.Tab(),
			strings.Join(colList, ", "), strings.Join(strings.Fields(strings.Repeat("? ", len(colList))), ", "), inc.columnName(), inc.columnName(),
			strings.Join(updateList, ", "))
		res, err = db.Exec(stmtStr, append(valList, valList...)...)
	} else {
		stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", m.DB(), m.Tab(), strings.Join(colList, ", "),
			strings.Join(strings.Fields(strings.Repeat("? ", len(colList))), ", "))
		res, err = db.Exec(stmtStr, valList...)
	}
	if err != nil {
		return err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	inc.(*IntField).Set(lastInsertId, false)
	return nil
}

func insertAndGetIncContext(ctx context.Context, m Model, update bool) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		db := dbMap[m.DB()]
		inc, others := getInc(m)
		others = filterValid(others)
		colList := make([]string, len(others))
		valList := make([]interface{}, len(others))
		updateList := make([]string, len(others))
		for i, f := range others {
			colList[i] = f.columnName()
			valList[i] = f.value()
			updateList[i] = f.columnName() + " = ?"
		}
		var res sql.Result
		var err error
		if update {
			stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", m.DB(), m.Tab(),
				strings.Join(colList, ", "), strings.Join(strings.Fields(strings.Repeat("? ", len(colList))), ", "), inc.columnName(), inc.columnName(),
				strings.Join(updateList, ", "))
			res, err = db.ExecContext(ctx, stmtStr, append(valList, valList...)...)
		} else {
			stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", m.DB(), m.Tab(), strings.Join(colList, ", "),
				strings.Join(strings.Fields(strings.Repeat("? ", len(colList))), ", "))
			res, err = db.ExecContext(ctx, stmtStr, valList...)
		}
		if err != nil {
			return err
		}
		lastInsertId, err := res.LastInsertId()
		if err != nil {
			return err
		}
		inc.(*IntField).Set(lastInsertId, false)
		return nil
	}
}

func iterList(l ModelList, f func(context.Context, Model) error) error {
	doneChan := make(chan interface{})
	errChan := make(chan error, l.Len())
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < l.Len(); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			err := f(ctx, l.Index(index))
			if err != nil {
				errChan <- err
			}
		}(i)
	}
	func() {
		wg.Wait()
		close(doneChan)
	}()
	select {
	case err := <-errChan:
		cancel()
		wg.Wait()
		close(errChan)
		return err
	case <-doneChan:
		close(errChan)
		return nil
	}
}

func toListStr(val interface{}) string {
	return strings.Join(strings.Fields(strings.Trim(fmt.Sprint(val), "[]")), ", ")
}
