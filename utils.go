package nborm

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"unsafe"
)

//UpdateValue is used for bulk update
type UpdateValue struct {
	column string
	val    interface{}
	null   bool
}

func (uv *UpdateValue) toSQL() (string, interface{}) {
	if uv.null {
		return fmt.Sprintf("%s = ?", uv.column), nil
	}
	return fmt.Sprintf("%s = ?", uv.column), uv.val
}

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

func filterValid(fields []Field) (validFields []Field) {
	for _, f := range fields {
		if f.IsValid() {
			validFields = append(validFields, f)
		}
	}
	return
}

func filterList(slice table, f func(uintptr) bool) {
	l := *(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice)) + uintptr(8)))
	i := 1
	for i < len(*l) {
		if f((*l)[i]) {
			*l = append((*l)[:i], (*l)[i+1:]...)
			continue
		}
		i++
	}
}

func iterList(slice table, f func(context.Context, uintptr) error) error {
	l := **(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice)) + uintptr(8)))
	doneChan := make(chan interface{})
	errChan := make(chan error, len(l))
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 1; i < len(l); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			err := f(ctx, l[index])
			if err != nil {
				errChan <- err
			}
		}(i)
	}
	func() {
		wg.Wait()
		close(doneChan)
	}()
	for {
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
}

func toListStr(val interface{}) string {
	return fmt.Sprintf("(%s)", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(val), "[]")), ", "))
}

var snakeCaseRe = regexp.MustCompile(`[A-Z]+[0-9a-z]*`)

func toSnakeCase(s string) string {
	newStr := snakeCaseRe.ReplaceAllStringFunc(s, func(v string) string {
		return "_" + strings.ToLower(v)
	})
	return strings.TrimLeft(strings.Replace(newStr, "___", "__", -1), "_")
}

func getTabAddr(tab table) uintptr {
	return *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&tab)) + uintptr(8)))
}

func setInc(addr uintptr, tabInfo *tableInfo, lastInsertId int64) {
	if tabInfo.inc != nil {
		inc := getIncWithTableInfo(addr, tabInfo)
		inc.setVal(lastInsertId, false)
	}
}

func genUpdVals(addr uintptr, tabInfo *tableInfo) []*UpdateValue {
	updVals := make([]*UpdateValue, 0, len(tabInfo.columns))
	for _, colInfo := range tabInfo.columns {
		field := getFieldByColumnInfo(addr, colInfo)
		if !colInfo.isInc && field.IsValid() {
			updVals = append(updVals, field.updateValue())
		}
	}
	return updVals
}

func getFoundRows(tx *sql.Tx) (int, error) {
	var num int
	row := tx.QueryRow("SELECT FOUND_ROWS()")
	if err := row.Scan(&num); err != nil {
		tx.Rollback()
		return -1, err
	}
	return num, nil
}

func getFoundRowsContext(ctx context.Context, tx *sql.Tx) (int, error) {
	var num int
	row := tx.QueryRowContext(ctx, "SELECT FOUND_ROWS()")
	if err := row.Scan(&num); err != nil {
		tx.Rollback()
		return -1, err
	}
	return num, nil
}

//NumRes return the number of query result set. Because the first element of model slice is an example model, so this function infact return
//ModelSlice[1:]
func NumRes(slice table) int {
	l := **(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice)) + uintptr(8)))
	return len(l) - 1
}

//ClsRes clear result set except the first element(example model)
func ClsRes(slice table) {
	l := *(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice)) + uintptr(8)))
	*l = (*l)[:1]
}

func wrap(s string) string {
	return fmt.Sprintf("`%s`", s)
}

func escap(s string) string {
	return strings.Trim(s, "`")
}
