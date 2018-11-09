package nborm

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/go-sql-driver/mysql"
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
		l[i] = fmt.Sprintf("%s.%s.%s %s", order.Field.dbName(), order.Field.tabName(), order.Field.columnName(), o)
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

// func scanRows(l ModelList, rows *sql.Rows) error {
// 	defer rows.Close()
// 	for rows.Next() {
// 		m := l.New()
// 		fields := m.Fields()
// 		scanList := make([]interface{}, len(fields))
// 		for i, f := range fields {
// 			scanList[i] = f
// 		}
// 		err := rows.Scan(scanList...)
// 		if err != nil {
// 			return err
// 		}
// 		m.SetSync(true)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return err
// 	}
// 	return nil
// }

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

// func scanRow(m Model, row *sql.Row) error {
// 	fields := m.Fields()
// 	scanList := make([]interface{}, len(fields))
// 	for i, f := range fields {
// 		scanList[i] = f
// 	}
// 	err := row.Scan(scanList...)
// 	if err != nil {
// 		return err
// 	}
// 	m.SetSync(true)
// 	return nil
// }

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

func queryAndScan(tab table, stmtStr string, valList []interface{}) error {
	db := dbMap[tab.DB()]
	switch obj := tab.(type) {
	case ModelList:
		rows, err := db.Query(stmtStr, valList...)
		if err != nil {
			return err
		}
		return scanRows(obj, rows)
	case Model:
		row := db.QueryRow(stmtStr, valList...)
		return scanRow(obj, row)
	default:
		return fmt.Errorf("nborm error: unsupported type (%T)", tab)
	}
}

func queryAndScanContext(ctx context.Context, tab table, stmtStr string, valList []interface{}) error {
	db := dbMap[tab.DB()]
	switch obj := tab.(type) {
	case ModelList:
		rows, err := db.QueryContext(ctx, stmtStr, valList...)
		if err != nil {
			return err
		}
		return scanRowsContext(ctx, obj, rows)
	case Model:
		row := db.QueryRowContext(ctx, stmtStr, valList...)
		return scanRowContext(ctx, obj, row)
	default:
		return fmt.Errorf("nborm error: unsupported type (%T)", tab)
	}
}

func queryAndScanWithNum(tab table, stmtStr string, valList []interface{}) (int, error) {
	if !strings.Contains(stmtStr, "SQL_CALC_FOUND_ROWS") {
		return -1, fmt.Errorf("the statement (%s) does not contains 'SQL_CALC_FOUND_ROWS'", stmtStr)
	}
	tx, err := dbMap[tab.DB()].Begin()
	if err != nil {
		return -1, err
	}
	switch obj := tab.(type) {
	case ModelList:
		rows, err := tx.Query(stmtStr, valList...)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
		if err := scanRows(obj, rows); err != nil {
			tx.Rollback()
			return -1, err
		}
		num, err := getFoundRows(tx)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
		tx.Commit()
		return num, nil
	case Model:
		row := tx.QueryRow(stmtStr, valList...)
		if err := scanRow(obj, row); err != nil {
			tx.Rollback()
			return -1, err
		}
		num, err := getFoundRows(tx)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
		tx.Commit()
		return num, nil
	default:
		return -1, fmt.Errorf("nborm error: unsupported type (%T)", tab)
	}
}

func queryAndScanWithNumContext(ctx context.Context, tab table, stmtStr string, valList []interface{}) (int, error) {
	if !strings.Contains(stmtStr, "SQL_CALC_FOUND_ROWS") {
		return -1, fmt.Errorf("the statement (%s) does not contains 'SQL_CALC_FOUND_ROWS'", stmtStr)
	}
	tx, err := dbMap[tab.DB()].Begin()
	if err != nil {
		return -1, err
	}
	switch obj := tab.(type) {
	case ModelList:
		rows, err := tx.QueryContext(ctx, stmtStr, valList...)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
		if err := scanRowsContext(ctx, obj, rows); err != nil {
			tx.Rollback()
			return -1, err
		}
		num, err := getFoundRowsContext(ctx, tx)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
		return num, nil
	case Model:
		row := tx.QueryRowContext(ctx, stmtStr, valList...)
		if err := scanRowContext(ctx, obj, row); err != nil {
			tx.Rollback()
			return -1, err
		}
		num, err := getFoundRowsContext(ctx, tx)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
		tx.Commit()
		return num, nil
	default:
		return -1, fmt.Errorf("nborm error: unsupported type (%T)", tab)
	}
}

func filterValid(fields []Field) (validFields []Field) {
	for _, f := range fields {
		if f.IsValid() {
			validFields = append(validFields, f)
		}
	}
	return
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
	info := getTabInfo(m)
	inc := getInc(m)
	if inc != nil && inc.IsValid() {
		return inc.where()
	}
	pks := getPks(m)
	var where *Where
	for _, pk := range pks {
		if !pk.IsValid() {
			where = nil
			break
		}
		where = where.And(pk.where())
	}
	if where != nil {
		return where
	}
	unis := getUnis(m)
	if len(unis) > 0 {
		for _, uni := range unis {
			if uni.IsValid() {
				return uni.where()
			}
		}
	}
	for _, col := range info.columns {
		field := getFieldByName(m, col.colName, info)
		if field.IsValid() {
			where = where.And(field.where())
		}
	}
	return where
}

func getFoundRows(tx *sql.Tx) (int, error) {
	var num int
	row := tx.QueryRow("SELECT FOUND_ROWS()")
	err := row.Scan(&num)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func getFoundRowsContext(ctx context.Context, tx *sql.Tx) (int, error) {
	var num int
	row := tx.QueryRowContext(ctx, "SELECT FOUND_ROWS()")
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
	inc, others := getIncAndOthers(m)
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
		inc, others := getIncAndOthers(m)
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

func relationQuery(tab table, relation relation, where *Where, sorter *Sorter, pager *Pager) error {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s %s %s", tab.DB(), tab.Tab(), relation.getSrcDB(), relation.getSrcTab(), relation.joinClause(),
		whereClause, sorter.toSQL(), pager.toSQL())
	return queryAndScan(tab, stmt, whereValues)
}

func relationQueryWithFoundRows(l ModelList, relation relation, where *Where, sorter *Sorter, pager *Pager) (rowsNum int, err error) {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s %s %s %s %s", l.Model().DB(), l.Model().Tab(), relation.getSrcDB(),
		relation.getSrcTab(), relation.joinClause(), whereClause, sorter.toSQL(), pager.toSQL())
	return queryAndScanWithNum(l, stmt, whereValues)
}

func relationAddOne(relation relation, m Model) error {
	tabInfo := getTabInfo(m)
	switch rel := relation.(type) {
	case *ReverseForeignKey:
		getFieldByName(m, rel.dstCol, tabInfo).setVal(rel.srcValF(), false)
		return insertAndGetInc(m, false)
	case *ManyToMany:
		err := insertAndGetInc(m, false)
		if err != nil {
			if e, ok := err.(*mysql.MySQLError); !ok || e.Number != 1062 {
				return err
			}
		}
		stmt := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", rel.midDB, rel.midTab, rel.midLeftCol, rel.midRightCol)
		_, err = dbMap[rel.midDB].Exec(stmt, rel.srcValF(), getFieldByName(m, rel.dstCol, tabInfo).value())
		if err != nil {
			if e, ok := err.(*mysql.MySQLError); !ok || e.Number != 1062 {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("nborm.relationAddOne() error: unsupported relation (%T)", relation)
	}
}

func relationRemoveOne(relation relation, m Model) error {
	tabInfo := getTabInfo(m)
	switch rel := relation.(type) {
	case *ReverseForeignKey:
		return DeleteOne(m)
	case *ManyToMany:
		stmt := fmt.Sprintf("DELETE FROM %s.%s WHERE %s = ? AND %s = ?", rel.midDB, rel.midTab, rel.midLeftCol, rel.midRightCol)
		_, err := dbMap[rel.midDB].Exec(stmt, rel.srcValF(), getFieldByName(m, rel.dstCol, tabInfo).value())
		return err
	default:
		return fmt.Errorf("nborm.relationRemoveOne() error: unsupported relation (%T)", relation)
	}
}

var snakeCaseRe = regexp.MustCompile(`[A-Z]+[0-9a-z]*`)

func toSnakeCase(s string) string {
	newStr := snakeCaseRe.ReplaceAllStringFunc(s, func(v string) string {
		return "_" + strings.ToLower(v)
	})
	return strings.TrimLeft(strings.Replace(newStr, "___", "__", -1), "_")
}
