package nborm

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"unicode"
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

// var snakeCaseRe = regexp.MustCompile(`[A-Z]+[0-9a-z]*`)

// func toSnakeCase(s string) string {
// 	newStr := snakeCaseRe.ReplaceAllStringFunc(s, func(v string) string {
// 		return "_" + strings.ToLower(v)
// 	})
// 	return strings.TrimLeft(strings.Replace(newStr, "___", "__", -1), "_")
// }

func toSnakeCase(s string) string {
	if s == "" {
		panic(errors.New("nborm.toSnakeCase() empty string"))
	}
	var builder strings.Builder
	buffer := bytes.NewBuffer(make([]byte, 0, len(s)))
	const (
		start int = iota
		lower
		upper
		other
	)
	var flag int
	reader := bufio.NewReader(strings.NewReader(s))
OUTER:
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				builder.Write(buffer.Bytes())
				break
			} else {
				panic(fmt.Errorf("nborm.toSnakeCase() error: %v", err))
			}
		}
		switch {
		case unicode.IsUpper(r):
			switch flag {
			case start:
				buffer.WriteRune(r)
				flag = upper
			case lower, other:
				builder.Write(buffer.Bytes())
				builder.WriteRune('_')
				buffer.Reset()
				buffer.WriteRune(r)
				flag = upper
			case upper:
				next, err := reader.Peek(1)
				if err != nil {
					if err == io.EOF {
						buffer.WriteRune(r)
						builder.Write(buffer.Bytes())
						buffer.Reset()
						break OUTER
					} else {
						panic(fmt.Errorf("nborm.toSnakeCase() error: %v", err))
					}
				}
				switch {
				case unicode.IsUpper(rune(next[0])):
					buffer.WriteRune(r)
				case unicode.IsLower(rune(next[0])):
					buffer.WriteByte('_')
					builder.Write(buffer.Bytes())
					buffer.Reset()
					buffer.WriteRune(r)
					flag = lower
				default:
					builder.Write(buffer.Bytes())
					buffer.Reset()
					buffer.WriteRune(r)
					flag = other

				}
			}
		case unicode.IsLower(r):
			switch flag {
			case start:
				buffer.WriteRune(r)
				flag = lower
			case lower:
				buffer.WriteRune(r)
			case upper, other:
				buffer.WriteRune(r)
				flag = lower
			}
		default:
			buffer.WriteRune(r)
			flag = other
		}
	}
	return strings.ToLower(builder.String())
}

func getTabAddr(tab table) uintptr {
	return *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&tab)) + uintptr(8)))
}

func setInc(addr uintptr, tabInfo *TableInfo, lastInsertId int64) {
	if tabInfo.Inc != nil {
		inc := getIncWithTableInfo(addr, tabInfo)
		inc.setVal(lastInsertId, false)
	}
}

func genUpdVals(addr uintptr, tabInfo *TableInfo) []*UpdateValue {
	updVals := make([]*UpdateValue, 0, len(tabInfo.Columns))
	for _, colInfo := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, colInfo)
		if !colInfo.IsInc && field.IsValid() {
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

func genMiddleTableName(srcModelName, srcFieldName, dstModelName, dstFieldName string) string {
	srcFactor := srcModelName + srcFieldName
	dstFactor := dstModelName + dstFieldName
	var length int
	if len(srcFactor) >= len(dstFactor) {
		length = len(srcFactor)

	} else {
		length = len(dstFactor)
	}
	factorA := make([]byte, length)
	factorB := make([]byte, length)
	result := make([]byte, length)
	copy(factorA, []byte(srcFactor))
	copy(factorB, []byte(dstFactor))
	for i := range result {
		result[i] = factorA[i] & factorB[i]
	}
	return fmt.Sprintf("%x", result)
}

type JSONFieldNameChoice int

const (
	FieldName JSONFieldNameChoice = iota
	ColumnName
)

func JsonifyModels(models Union, fieldNameChoice JSONFieldNameChoice, fieldNames ...string) (map[string]interface{}, error) {
	if len(models) == 0 {
		return nil, errors.New("nborm.JsonifyModels() error: no models")
	}
	if len(fieldNames) == 0 {
		return nil, errors.New("nborm.Jsonify() error: no field name")
	}
	m := make(map[string]interface{})
	tabInfos := make([]*TableInfo, len(models))
	tabAddrs := make([]uintptr, len(models))
	for i, model := range models {
		tabInfos[i] = getTabInfo(model)
		tabAddrs[i] = getTabAddr(model)
	}
	for _, fieldName := range fieldNames {
		l := strings.Split(fieldName, ".")
		if len(l) != 2 {
			return nil, fmt.Errorf("nborm.JsonifyModels() error: invalid field name (%s) want <modelName>.<fieldName>", fieldName)
		}
		var tabInfo *TableInfo
		var tabAddr uintptr
		for i, ti := range tabInfos {
			if ti.ModelName == l[0] {
				tabInfo = ti
				tabAddr = tabAddrs[i]
				break
			}
		}
		if tabInfo == nil {
			return nil, fmt.Errorf("nborm.JsonifyModels() error: model not exists (%s)", l[0])
		}
		field, err := getFieldByFieldName(tabAddr, l[1], tabInfo)
		if err != nil {
			return nil, err
		}
		switch fieldNameChoice {
		case FieldName:
			switch f := field.(type) {
			case *DateField:
				m[l[1]] = f.ISOFormat()
			case *DatetimeField:
				m[l[1]] = f.ISOFormat()
			default:
				m[l[1]] = f.value()
			}
		case ColumnName:
			switch f := field.(type) {
			case *DateField:
				m[f.column] = f.ISOFormat()
			case *DatetimeField:
				m[f.column] = f.ISOFormat()
			default:
				m[escap(f.columnName())] = f.value()
			}
		}
	}
	return m, nil
}

func JsonifySlices(slices Union, fieldNameChoice JSONFieldNameChoice, fieldNames ...string) ([]map[string]interface{}, error) {
	if len(slices) == 0 {
		return nil, errors.New("nborm.JsonifySlices() error: no slice")
	}
	if len(fieldNames) == 0 {
		return nil, errors.New("nborm.JsonifySlices() error: no field name")
	}
	tabInfos := make([]*TableInfo, len(slices))
	for i, slice := range slices {
		tabInfos[i] = getTabInfo(slice)
	}
	sliceLength := len(*(*[]uintptr)(unsafe.Pointer(getTabAddr(slices[0]))))
	l := make([][]uintptr, sliceLength-1)
	for _, slice := range slices {
		ptrSlice := *(*[]uintptr)(unsafe.Pointer(getTabAddr(slice)))
		for i, ptr := range ptrSlice[1:] {
			l[i] = append(l[i], ptr)
		}
	}
	ml := make([]map[string]interface{}, sliceLength-1)
	colNameList := make([][][2]string, len(slices))
	var complete bool
	for mi, addrs := range l {
		m := make(map[string]interface{})
		if !complete {
		OUTER:
			for _, fieldName := range fieldNames {
				l := strings.Split(fieldName, ".")
				if len(l) != 2 {
					return nil, fmt.Errorf("nborm.JsonifyModels() error: invalid field name (%s) want <modelName>.<fieldName>", fieldName)
				}
				for i, ti := range tabInfos {
					if ti.ModelName == l[0] {
						tabInfo := ti
						tabAddr := addrs[i]
						field, err := getFieldByFieldName(tabAddr, l[1], tabInfo)
						if err != nil {
							return nil, err
						}
						colNameList[i] = append(colNameList[i], [2]string{escap(field.columnName()), l[1]})
						switch fieldNameChoice {
						case FieldName:
							switch f := field.(type) {
							case *DateField:
								m[l[1]] = f.ISOFormat()
							case *DatetimeField:
								m[l[1]] = f.ISOFormat()
							default:
								m[l[1]] = f.value()
							}
						case ColumnName:
							switch f := field.(type) {
							case *DateField:
								m[f.column] = f.ISOFormat()
							case *DatetimeField:
								m[f.column] = f.ISOFormat()
							default:
								m[escap(f.columnName())] = f.value()
							}
						}
						continue OUTER
					}
				}
				return nil, fmt.Errorf("nborm.JsonifySlices() error: model not exists (%s)", l[0])
			}
			ml[mi] = m
			complete = true
		} else {
			for i, colNames := range colNameList {
				tabInfo := tabInfos[i]
				addr := addrs[i]
				for _, colName := range colNames {
					field := getFieldByName(addr, colName[0], tabInfo)
					switch fieldNameChoice {
					case FieldName:
						switch f := field.(type) {
						case *DateField:
							m[colName[1]] = f.ISOFormat()
						case *DatetimeField:
							m[colName[1]] = f.ISOFormat()
						default:
							m[colName[1]] = f.value()
						}
					case ColumnName:
						switch f := field.(type) {
						case *DateField:
							m[f.column] = f.ISOFormat()
						case *DatetimeField:
							m[f.column] = f.ISOFormat()
						default:
							m[escap(f.columnName())] = f.value()
						}
					}
				}
			}
			ml[mi] = m
		}
	}
	return ml, nil
}
