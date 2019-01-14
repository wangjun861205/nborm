package nborm

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

type ModelStatus struct {
	isInit bool
	isSync bool
}

func (s *ModelStatus) IsSync() bool {
	return s.isSync
}

func getStatus(addr uintptr, tabInfo *TableInfo) *ModelStatus {
	return (*ModelStatus)(unsafe.Pointer(addr + tabInfo.ModelStatus))
}

func setSync(addr uintptr, tabInfo *TableInfo) {
	status := getStatus(addr, tabInfo)
	status.isSync = true
}

func unsetSync(addr uintptr, tabInfo *TableInfo) {
	status := getStatus(addr, tabInfo)
	status.isSync = false
}

func getSync(addr uintptr, tabInfo *TableInfo) bool {
	status := getStatus(addr, tabInfo)
	return status.isSync
}

func setInit(addr uintptr, tabInfo *TableInfo) {
	status := getStatus(addr, tabInfo)
	status.isInit = true
}

func getInit(addr uintptr, tabInfo *TableInfo) bool {
	status := getStatus(addr, tabInfo)
	return status.isInit
}

//StringField represent char, varchar, text type in mysql
type StringField struct {
	db         string
	tab        string
	column     string
	nullable   bool
	pk         bool
	uni        bool
	defVal     interface{}
	offset     uintptr
	val        string
	valid      bool
	null       bool
	validators []Validator
	fieldName  string
	modelName  string
}

func (f *StringField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

func (f *StringField) dbName() string {
	return wrap(f.db)
}

func (f *StringField) tabName() string {
	return wrap(f.tab)
}

func (f *StringField) fullTabName() string {
	return fmt.Sprintf("%s.%s", f.dbName(), f.tabName())
}

func (f *StringField) columnName() string {
	return wrap(f.column)
}

func (f *StringField) getFieldName() string {
	return f.fieldName
}

func (f *StringField) fullColName() string {
	return fmt.Sprintf("%s.%s.%s", f.dbName(), f.tabName(), f.columnName())
}

func (f *StringField) getModelName() string {
	return f.modelName
}

//IsPk return true if the column is a primary key
func (f *StringField) isPk() bool {
	return f.pk
}

//IsInc return true if the column is a auto_increment column
func (f *StringField) isInc() bool {
	return false
}

//IsValid return true if the field has been set a value or been scaned
func (f *StringField) IsValid() bool {
	return f.valid
}

//IsNull return true if the value of this column is null
func (f *StringField) IsNull() bool {
	return f.null
}

func (f *StringField) isNullable() bool {
	return f.nullable
}

//IsUni return true if this column is a unique key
func (f *StringField) isUni() bool {
	return f.uni
}

//Set set a value for the field
func (f *StringField) Set(val string) {
	f.valid, f.null, f.val = true, false, val
}

func (f *StringField) SetNull() {
	f.valid, f.null, f.val = true, true, ""
}

func (f *StringField) setVal(val interface{}, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val.(string)
}

//SetByUpdateValue set by UpdateValue struct
func (f *StringField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(string)
}

//Get get the field value
func (f *StringField) Get() (val string, valid bool, null bool) {
	return f.val, f.valid, f.null
}

//Scan implement sql.Scanner interface
func (f *StringField) Scan(v interface{}) error {
	f.valid = true
	if v == nil {
		f.null = true
		return nil
	}
	f.null = false
	switch val := v.(type) {
	case []byte:
		f.val = string(val)
	case string:
		f.val = val
	default:
		return fmt.Errorf("not supported value type for StringField: %T", val)
	}
	return nil
}

//Invalidate set the field invalid
func (f *StringField) Invalidate() {
	f.valid = false
}

//MarshalJSON implement json.Marshaler interface
func (f *StringField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(fmt.Sprintf("%q", f.val)), nil
}

//BulkUpdateValue generate a UpdateValue by value which is passed in, it is for bulk update
func (f *StringField) UpdateValue(val string, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val, isNull}
}

func (f *StringField) updateValue() *UpdateValue {
	return &UpdateValue{f.column, f.val, f.null}
}

//Where generate a Where by self value
func (f *StringField) where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	} else {
		return newWhere(f.db, f.tab, f.column, "=", f.val)
	}
}

func (f *StringField) Where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	} else {
		return newWhere(f.db, f.tab, f.column, "=", f.val)
	}
}

//Eq generate a equal Where
func (f *StringField) Eq(val string) *Where {
	return newWhere(f.db, f.tab, f.column, "=", val)

}

//Neq generate a not equal Where
func (f *StringField) Neq(val string) *Where {
	return newWhere(f.db, f.tab, f.column, "<>", val)
}

//Contains generate a contains Where
func (f *StringField) Contains(val string) *Where {
	return newWhere(f.db, f.tab, f.column, "LIKE", "%"+val+"%")
}

//Null generate a Where which represent this column is null
func (f *StringField) Null() *Where {
	return newWhere(f.db, f.tab, f.column, "IS", "NULL")
}

//NotNull generate a Where which represent this column is not null
func (f *StringField) NotNull() *Where {
	return newWhere(f.db, f.tab, f.column, "IS NOT", "NULL")
}

//In generate a in Where
func (f *StringField) In(val []string) *Where {
	l := make([]string, len(val))
	for i, v := range val {
		l[i] = fmt.Sprintf("%q", v)
	}
	return newWhere(f.db, f.tab, f.column, "IN", fmt.Sprintf("(%s)", strings.Join(l, ", ")))
}

//LessFunc generate a function for sort a ModelList
func (f *StringField) LessFunc(reverse bool) func(iaddr, jaddr uintptr) int {
	if reverse {
		return func(iaddr, jaddr uintptr) int {
			iField, jField := (*StringField)(unsafe.Pointer(iaddr+f.offset)), (*StringField)(unsafe.Pointer(jaddr+f.offset))
			iVal, iValid, iNull := iField.Get()
			jVal, jValid, jNull := jField.Get()
			var iBit, jBit int
			if iValid {
				iBit |= 1 << 1
			}
			if !iNull {
				iBit |= 1
			}
			if jValid {
				jBit |= 1 << 1
			}
			if !jNull {
				jBit |= 1
			}
			if iBit != 3 || jBit != 3 {
				return -(iBit - jBit)
			}
			switch {
			case iVal < jVal:
				return 1
			case iVal > jVal:
				return -1
			default:
				return 0
			}
		}
	}
	return func(iaddr, jaddr uintptr) int {
		iField, jField := (*StringField)(unsafe.Pointer(iaddr+f.offset)), (*StringField)(unsafe.Pointer(jaddr+f.offset))
		iVal, iValid, iNull := iField.Get()
		jVal, jValid, jNull := jField.Get()
		var iBit, jBit int
		if iValid {
			iBit |= 1 << 1
		}
		if !iNull {
			iBit |= 1
		}
		if jValid {
			jBit |= 1 << 1
		}
		if !jNull {
			jBit |= 1
		}
		if iBit != 3 || jBit != 3 {
			return iBit - jBit
		}
		switch {
		case iVal < jVal:
			return -1
		case iVal > jVal:
			return 1
		default:
			return 0
		}
	}
}

//SortOrder return a string represent the sort order
func (f *StringField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.db, f.tab, f.column, o)
}

func (f *StringField) getDefVal() interface{} {
	return f.defVal
}

func (f *StringField) check() error {
	for _, validator := range f.validators {
		if err := validator(f); err != nil {
			return err
		}
	}
	return nil
}

//IntField represent int, tinyint, midint, bigint types in mysql
type IntField struct {
	db         string
	tab        string
	column     string
	nullable   bool
	pk         bool
	inc        bool
	uni        bool
	defVal     interface{}
	offset     uintptr
	val        int64
	valid      bool
	null       bool
	validators []Validator
	fieldName  string
	modelName  string
}

func (f *IntField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

func (f *IntField) dbName() string {
	return wrap(f.db)
}

func (f *IntField) tabName() string {
	return wrap(f.tab)
}

func (f *IntField) fullTabName() string {
	return fmt.Sprintf("%s.%s", f.dbName(), f.tabName())
}

func (f *IntField) columnName() string {
	return wrap(f.column)
}

func (f *IntField) getModelName() string {
	return f.modelName
}

func (f *IntField) getFieldName() string {
	return f.fieldName
}

func (f *IntField) fullColName() string {
	return fmt.Sprintf("%s.%s.%s", f.dbName(), f.tabName(), f.columnName())
}

//IsPk return true if this column is primary key
func (f *IntField) isPk() bool {
	return f.pk
}

//IsInc return true if this column is auto_increment column
func (f *IntField) isInc() bool {
	return f.inc
}

//IsValid return true if this field has been set a value or been scaned
func (f *IntField) IsValid() bool {
	return f.valid
}

//IsNull return true if this field value is null
func (f *IntField) IsNull() bool {
	return f.null
}

func (f *IntField) isNullable() bool {
	return f.nullable
}

//IsUni return true if the column is a unique key
func (f *IntField) isUni() bool {
	return f.uni
}

//Set set a value for this field
func (f *IntField) Set(val int64) {
	f.valid, f.null, f.val = true, false, val
}

func (f *IntField) SetNull() {
	f.valid, f.null, f.val = true, true, 0
}

func (f *IntField) setVal(val interface{}, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val.(int64)
}

//SetByUpdateValue set a value by UpdateValue struct
func (f *IntField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(int64)
}

//Get get value
func (f *IntField) Get() (val int64, valid bool, null bool) {
	return f.val, f.valid, f.null
}

//Scan implement sql.Scanner interface
func (f *IntField) Scan(v interface{}) error {
	f.valid = true
	if v == nil {
		f.null = true
		return nil
	}
	f.null = false
	switch val := v.(type) {
	case []byte:
		i64, err := strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return err
		}
		f.val = i64
	case int64:
		f.val = val
	default:
		return fmt.Errorf("not supported value type for IntField: %T", val)
	}
	return nil
}

//Invalidate invalidate the field
func (f *IntField) Invalidate() {
	f.valid = false
}

//MarshalJSON implement json.Marshaler interface
func (f *IntField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(strconv.FormatInt(f.val, 10)), nil
}

//BulkUpdateValue return a UpdateValue for bulk update
func (f *IntField) UpdateValue(val int64, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val, isNull}
}

func (f *IntField) updateValue() *UpdateValue {
	return &UpdateValue{f.column, f.val, f.null}
}

//Where generate a Where by self value
func (f *IntField) where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val)
}

func (f *IntField) Where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val)
}

//Eq generate a equal Where
func (f *IntField) Eq(val int64) *Where {
	return newWhere(f.db, f.tab, f.column, "=", val)
}

//Neq generate a not equal Where
func (f *IntField) Neq(val int64) *Where {
	return newWhere(f.db, f.tab, f.column, "<>", val)
}

//Lt generate a less than Where
func (f *IntField) Lt(val int64) *Where {
	return newWhere(f.db, f.tab, f.column, "<", val)
}

//Gt generate a great than Where
func (f *IntField) Gt(val int64) *Where {
	return newWhere(f.db, f.tab, f.column, ">", val)
}

//Lte generate a less than equal Where
func (f *IntField) Lte(val int64) *Where {
	return newWhere(f.db, f.tab, f.column, "<=", val)
}

//Gte generate a great than equal Where
func (f *IntField) Gte(val int64) *Where {
	return newWhere(f.db, f.tab, f.column, ">=", val)
}

//Null generate null Where
func (f *IntField) Null() *Where {
	return newWhere(f.db, f.tab, f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *IntField) NotNull() *Where {
	return newWhere(f.db, f.tab, f.column, "IS NOT", "NULL")
}

//In generate a in Where
func (f *IntField) In(val []int) *Where {
	return newWhere(f.db, f.tab, f.column, "IN", toListStr(val))
}

//LessFunc return a function for sort ModelList
func (f *IntField) LessFunc(reverse bool) func(iaddr, jaddr uintptr) int {
	if reverse {
		return func(iaddr, jaddr uintptr) int {
			iField, jField := *(**IntField)(unsafe.Pointer(iaddr + f.offset)), *(**IntField)(unsafe.Pointer(jaddr + f.offset))
			iVal, iValid, iNull := iField.Get()
			jVal, jValid, jNull := jField.Get()
			var iBit, jBit int
			if iValid {
				iBit |= 1 << 1
			}
			if !iNull {
				iBit |= 1
			}
			if jValid {
				jBit |= 1 << 1
			}
			if !jNull {
				jBit |= 1
			}
			if iBit != 3 || jBit != 3 {
				return -(iBit - jBit)
			}
			switch {
			case iVal < jVal:
				return 1
			case iVal > jVal:
				return -1
			default:
				return 0
			}
		}
	}
	return func(iaddr, jaddr uintptr) int {
		iField, jField := *(**IntField)(unsafe.Pointer(iaddr + f.offset)), *(**IntField)(unsafe.Pointer(jaddr + f.offset))
		iVal, iValid, iNull := iField.Get()
		jVal, jValid, jNull := jField.Get()
		var iBit, jBit int
		if iValid {
			iBit |= 1 << 1
		}
		if !iNull {
			iBit |= 1
		}
		if jValid {
			jBit |= 1 << 1
		}
		if !jNull {
			jBit |= 1
		}
		if iBit != 3 || jBit != 3 {
			return iBit - jBit
		}
		switch {
		case iVal < jVal:
			return -1
		case iVal > jVal:
			return 1
		default:
			return 0
		}
	}
}

//SortOrder return a string represent a sort order
func (f *IntField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.db, f.tab, f.column, o)
}

func (f *IntField) getDefVal() interface{} {
	return f.defVal
}

func (f *IntField) check() error {
	for _, validator := range f.validators {
		if err := validator(f); err != nil {
			return err
		}
	}
	return nil
}

//SumFunc return a function for sum aggregate
func (f *IntField) SumFunc() func(slice table) float64 {
	return func(slice table) float64 {
		l := **(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice)) + uintptr(8)))
		var sum int
		for i := 0; i < len(l); i++ {
			field := (*IntField)(unsafe.Pointer(l[i] + f.offset))
			val, valid, null := field.Get()
			if valid && !null {
				sum += int(val)
			}
		}
		return float64(sum)
	}
}

//AvgFunc return a funcion for average aggregate
func (f *IntField) AvgFunc(ignoreNull bool) func(slice table) float64 {
	if ignoreNull {
		return func(slice table) float64 {
			l := **(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice)) + uintptr(8)))
			var num, sum int
			for i := 0; i < len(l); i++ {
				field := (*IntField)(unsafe.Pointer(l[i] + f.offset))
				val, valid, null := field.Get()
				if valid && !null {
					num++
					sum += int(val)
				}
			}
			if num == 0 {
				return 0
			}
			return float64(sum) / float64(num)
		}
	}
	return func(slice table) float64 {
		l := **(**[]uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice)) + uintptr(8)))
		var num, sum int
		for i := 0; i < len(l); i++ {
			field := (*IntField)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice)) + uintptr(8)))
			val, valid, null := field.Get()
			num++
			if valid && !null {
				sum += int(val)
			}
		}
		if num == 0 {
			return 0
		}
		return float64(sum) / float64(num)
	}
}

//FloatField represent float, decimal types in mysql
type FloatField struct {
	db         string
	tab        string
	column     string
	nullable   bool
	pk         bool
	uni        bool
	defVal     interface{}
	offset     uintptr
	val        float64
	valid      bool
	null       bool
	validators []Validator
	fieldName  string
	modelName  string
}

//NewFloatField create new FloatField
func NewFloatField(db, tab, column string, nullable, pk, uni bool, defVal interface{}, offset uintptr) *FloatField {
	return &FloatField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal, offset: offset}
}

func (f *FloatField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

func (f *FloatField) dbName() string {
	return wrap(f.db)
}

func (f *FloatField) tabName() string {
	return wrap(f.tab)
}

func (f *FloatField) columnName() string {
	return wrap(f.column)
}

func (f *FloatField) getModelName() string {
	return f.modelName
}

func (f *FloatField) getFieldName() string {
	return f.fieldName
}

func (f *FloatField) fullTabName() string {
	return fmt.Sprintf("%s.%s", f.dbName(), f.tabName())
}

func (f *FloatField) fullColName() string {
	return fmt.Sprintf("%s.%s.%s", f.dbName(), f.tabName(), f.columnName())
}

//IsPk return true if the column is primary key
func (f *FloatField) isPk() bool {
	return f.pk
}

//IsInc return true if the column is auto_increment column
func (f *FloatField) isInc() bool {
	return false
}

//IsValid return true if the field has been set a value or been scanned
func (f *FloatField) IsValid() bool {
	return f.valid
}

//IsNull return true if the field value is null
func (f *FloatField) IsNull() bool {
	return f.null
}

func (f *FloatField) isNullable() bool {
	return f.nullable
}

//IsUni return true if the column is a unique key
func (f *FloatField) isUni() bool {
	return f.uni
}

//Set set a value
func (f *FloatField) Set(val float64) {
	f.valid, f.null, f.val = true, false, val
}

func (f *FloatField) SetNull() {
	f.valid, f.null, f.val = true, true, 0
}

func (f *FloatField) setVal(val interface{}, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val.(float64)
}

//SetByUpdateValue set a value by UpdateValue struct
func (f *FloatField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(float64)
}

//Get get value
func (f *FloatField) Get() (val float64, valid bool, null bool) {
	return f.val, f.valid, f.null
}

//Scan implement sql.Scanner interface
func (f *FloatField) Scan(v interface{}) error {
	f.valid = true
	if v == nil {
		f.null = true
		return nil
	}
	f.null = false
	switch val := v.(type) {
	case []byte:
		f64, err := strconv.ParseFloat(string(val), 64)
		if err != nil {
			return err
		}
		f.val = f64
	case float64:
		f.val = val
	default:
		return fmt.Errorf("not supported value type for IntField: %T", val)
	}
	return nil
}

//Invalidate invalidate field
func (f *FloatField) Invalidate() {
	f.valid = false
}

//MarshalJSON implement json.Marshaler interface
func (f *FloatField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(strconv.FormatFloat(f.val, 'f', -1, 64)), nil
}

//BulkUpdateValue generate a UpdateValue for bulk update
func (f *FloatField) UpdateValue(val float64, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val, isNull}
}

func (f *FloatField) updateValue() *UpdateValue {
	return &UpdateValue{f.column, f.val, f.null}
}

//Where generate a Where by self value
func (f *FloatField) where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val)
}

func (f *FloatField) Where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val)
}

//Eq generate a equal Where
func (f *FloatField) Eq(val float64) *Where {
	return newWhere(f.db, f.tab, f.column, "=", val)
}

//Neq generate a not equal Where
func (f *FloatField) Neq(val float64) *Where {
	return newWhere(f.db, f.tab, f.column, "<>", val)
}

//Lt generate a less than Where
func (f *FloatField) Lt(val float64) *Where {
	return newWhere(f.db, f.tab, f.column, "<", val)
}

//Gt generate a great than Where
func (f *FloatField) Gt(val float64) *Where {
	return newWhere(f.db, f.tab, f.column, ">", val)
}

//Lte generate a less than equal Where
func (f *FloatField) Lte(val float64) *Where {
	return newWhere(f.db, f.tab, f.column, "<=", val)
}

//Gte generate a great than equal Where
func (f *FloatField) Gte(val float64) *Where {
	return newWhere(f.db, f.tab, f.column, ">=", val)
}

//Null generate a null Where
func (f *FloatField) Null() *Where {
	return newWhere(f.db, f.tab, f.column, "IS", "NULL")
}

//NotNull generate a not null Where
func (f *FloatField) NotNull() *Where {
	return newWhere(f.db, f.tab, f.column, "<>", "NULL")
}

//In generate in Where
func (f *FloatField) In(val []float64) *Where {
	return newWhere(f.db, f.tab, f.column, "IN", toListStr(val))
}

//LessFunc return a func for sort ModelList
func (f *FloatField) LessFunc(reverse bool) func(iaddr, jaddr uintptr) int {
	if reverse {
		return func(iaddr, jaddr uintptr) int {
			iField, jField := *(**FloatField)(unsafe.Pointer(iaddr + f.offset)), *(**FloatField)(unsafe.Pointer(jaddr + f.offset))
			iVal, iValid, iNull := iField.Get()
			jVal, jValid, jNull := jField.Get()
			var iBit, jBit int
			if iValid {
				iBit |= 1 << 1
			}
			if !iNull {
				iBit |= 1
			}
			if jValid {
				jBit |= 1 << 1
			}
			if !jNull {
				jBit |= 1
			}
			if iBit != 3 || jBit != 3 {
				return -(iBit - jBit)
			}
			switch {
			case iVal < jVal:
				return 1
			case iVal > jVal:
				return -1
			default:
				return 0
			}
		}
	}
	return func(iaddr, jaddr uintptr) int {
		iField, jField := *(**FloatField)(unsafe.Pointer(iaddr + f.offset)), *(**FloatField)(unsafe.Pointer(jaddr + f.offset))
		iVal, iValid, iNull := iField.Get()
		jVal, jValid, jNull := jField.Get()
		var iBit, jBit int
		if iValid {
			iBit |= 1 << 1
		}
		if !iNull {
			iBit |= 1
		}
		if jValid {
			jBit |= 1 << 1
		}
		if !jNull {
			jBit |= 1
		}
		if iBit != 3 || jBit != 3 {
			return iBit - jBit
		}
		switch {
		case iVal < jVal:
			return -1
		case iVal > jVal:
			return 1
		default:
			return 0
		}
	}
}

//SortOrder return a string represent sort order
func (f *FloatField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.db, f.tab, f.column, o)
}

func (f *FloatField) getDefVal() interface{} {
	return f.defVal
}

func (f *FloatField) check() error {
	for _, validator := range f.validators {
		if err := validator(f); err != nil {
			return err
		}
	}
	return nil
}

//BoolField represent tinyint(1) type in mysql
type BoolField struct {
	db         string
	tab        string
	column     string
	nullable   bool
	pk         bool
	uni        bool
	defVal     interface{}
	offset     uintptr
	val        bool
	valid      bool
	null       bool
	validators []Validator
	fieldName  string
	modelName  string
}

//NewBoolField return a new BoolField
func NewBoolField(db, tab, column string, nullable, pk, uni bool, defVal interface{}, offset uintptr) *BoolField {
	return &BoolField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal, offset: offset}
}

func (f *BoolField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

func (f *BoolField) dbName() string {
	return wrap(f.db)
}

func (f *BoolField) tabName() string {
	return wrap(f.tab)
}

func (f *BoolField) columnName() string {
	return wrap(f.column)
}

func (f *BoolField) getModelName() string {
	return f.modelName
}

func (f *BoolField) getFieldName() string {
	return f.fieldName
}

func (f *BoolField) fullTabName() string {
	return fmt.Sprintf("%s.%s", f.dbName(), f.tabName())
}

func (f *BoolField) fullColName() string {
	return fmt.Sprintf("%s.%s.%s", f.dbName(), f.tabName(), f.columnName())
}

//IsPk return true if the column is a primary key
func (f *BoolField) isPk() bool {
	return f.pk
}

//IsInc return true if the column is a auto_increment column
func (f *BoolField) isInc() bool {
	return false
}

//IsValid return true when the field has been set a value or been scanned
func (f *BoolField) IsValid() bool {
	return f.valid
}

//IsNull return true if value is null
func (f *BoolField) IsNull() bool {
	return f.null
}

func (f *BoolField) isNullable() bool {
	return f.nullable
}

//Set set a value
func (f *BoolField) Set(val bool) {
	f.valid, f.null, f.val = true, false, val
}

func (f *BoolField) SetNull() {
	f.valid, f.null, f.val = true, true, false
}

func (f *BoolField) setVal(val interface{}, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val.(bool)
}

//SetByUpdateValue set value by a UpdateValue struct
func (f *BoolField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(bool)
}

//Get get value
func (f *BoolField) Get() (val bool, valid bool, null bool) {
	return f.val, f.valid, f.null
}

//Scan implement sql.Scanner interface
func (f *BoolField) Scan(v interface{}) error {
	f.valid = true
	if v == nil {
		f.null = true
		return nil
	}
	f.null = false
	switch val := v.(type) {
	case []byte:
		b, err := strconv.ParseBool(string(val))
		if err != nil {
			return err
		}
		f.val = b
	case bool:
		f.val = val
	case int64:
		if val == 1 {
			f.val = true
		}
	default:
		return fmt.Errorf("not supported value type for BoolField: %T", val)
	}
	return nil
}

//Invalidate invalidate field
func (f *BoolField) Invalidate() {
	f.valid = false
}

//MarshalJSON implement json.Marshaler interface
func (f *BoolField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(strconv.FormatBool(f.val)), nil
}

//BulkUpdateValue return a UpdateValue struct for bulk update
func (f *BoolField) UpdateValue(val bool, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val, isNull}
}

func (f *BoolField) updateValue() *UpdateValue {
	return &UpdateValue{f.column, f.val, f.null}
}

//Where generate a Where by self value
func (f *BoolField) where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val)
}

func (f *BoolField) Where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val)
}

//Eq generate euqal Where
func (f *BoolField) Eq(val bool) *Where {
	return newWhere(f.db, f.tab, f.column, "=", val)
}

//Neq generate not equal Where
func (f *BoolField) Neq(val bool) *Where {
	return newWhere(f.db, f.tab, f.column, "<>", val)
}

//Null generate null Where
func (f *BoolField) Null() *Where {
	return newWhere(f.db, f.tab, f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *BoolField) NotNull() *Where {
	return newWhere(f.db, f.tab, f.column, "IS NOT", "NULL")
}

//IsUni return true if column is a unique key
func (f *BoolField) isUni() bool {
	return f.uni
}

//LessFunc return a func for sort ModelList
func (f *BoolField) LessFunc(reverse bool) func(iaddr, jaddr uintptr) int {
	if reverse {
		return func(iaddr, jaddr uintptr) int {
			iField, jField := *(**BoolField)(unsafe.Pointer(iaddr + f.offset)), *(**BoolField)(unsafe.Pointer(jaddr + f.offset))
			iVal, iValid, iNull := iField.Get()
			jVal, jValid, jNull := jField.Get()
			var iBit, jBit int
			if iValid {
				iBit |= 1 << 2
			}
			if !iNull {
				iBit |= 1 << 1
			}
			if iVal {
				iBit |= 1
			}
			if jValid {
				jBit |= 1 << 2
			}
			if !jNull {
				jBit |= 1 << 1
			}
			if jVal {
				jBit |= 1
			}
			return -(iBit - jBit)
		}
	}
	return func(iaddr, jaddr uintptr) int {
		iField, jField := *(**BoolField)(unsafe.Pointer(iaddr + f.offset)), *(**BoolField)(unsafe.Pointer(jaddr + f.offset))
		iVal, iValid, iNull := iField.Get()
		jVal, jValid, jNull := jField.Get()
		var iBit, jBit int
		if iValid {
			iBit |= 1 << 2
		}
		if !iNull {
			iBit |= 1 << 1
		}
		if iVal {
			iBit |= 1
		}
		if jValid {
			jBit |= 1 << 2
		}
		if !jNull {
			jBit |= 1 << 1
		}
		if jVal {
			jBit |= 1
		}
		return iBit - jBit
	}
}

//SortOrder return a string represent sort order
func (f *BoolField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.db, f.tab, f.column, o)
}

func (f *BoolField) getDefVal() interface{} {
	return f.defVal
}

func (f *BoolField) check() error {
	for _, validator := range f.validators {
		if err := validator(f); err != nil {
			return err
		}
	}
	return nil
}

//DateField represent date type in mysql
type DateField struct {
	db         string
	tab        string
	column     string
	nullable   bool
	pk         bool
	uni        bool
	defVal     interface{}
	offset     uintptr
	val        time.Time
	valid      bool
	null       bool
	validators []Validator
	fieldName  string
	modelName  string
}

//NewDateField create a new DateField
func NewDateField(db, tab, column string, nullable, pk, uni bool, defVal interface{}, offset uintptr) *DateField {
	return &DateField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal, offset: offset}
}

func (f *DateField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val.Format("2006-01-02")
}

func (f *DateField) dbName() string {
	return wrap(f.db)
}

func (f *DateField) tabName() string {
	return wrap(f.tab)
}

func (f *DateField) columnName() string {
	return wrap(f.column)
}

func (f *DateField) getModelName() string {
	return f.modelName
}

func (f *DateField) getFieldName() string {
	return f.fieldName
}

func (f *DateField) fullTabName() string {
	return fmt.Sprintf("%s.%s", f.dbName(), f.tabName())
}

func (f *DateField) fullColName() string {
	return fmt.Sprintf("%s.%s.%s", f.dbName(), f.tabName(), f.columnName())
}

//IsPk return true if the column is a primary key
func (f *DateField) isPk() bool {
	return f.pk
}

//IsInc return true if the column is a auto_increment column
func (f *DateField) isInc() bool {
	return false
}

//IsValid return true if the field has been set a value or been scanned
func (f *DateField) IsValid() bool {
	return f.valid
}

//IsNull return true if value is null
func (f *DateField) IsNull() bool {
	return f.null
}

func (f *DateField) isNullable() bool {
	return f.nullable
}

//IsUni return true if the column is a unique key
func (f *DateField) isUni() bool {
	return f.uni
}

//Set set a value
func (f *DateField) Set(val time.Time) {
	f.valid, f.null, f.val = true, false, val
}

func (f *DateField) SetNull() {
	f.valid, f.null, f.val = true, true, time.Time{}
}

func (f *DateField) setVal(val interface{}, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val.(time.Time)
}

//SetByUpdateValue set value by UpdateValue struct
func (f *DateField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(time.Time)
}

//Get get value
func (f *DateField) Get() (val time.Time, valid bool, null bool) {
	return f.val, f.valid, f.null
}

//Scan implement sql.Scanner interface
func (f *DateField) Scan(v interface{}) error {
	f.valid = true
	if v == nil {
		f.null = true
		return nil
	}
	f.null = false
	switch val := v.(type) {
	case []byte:
		t, err := time.Parse("2006-01-02", string(val))
		if err != nil {
			return err
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		return fmt.Errorf("not supported value type for DateField: %T", val)
	}
	return nil
}

//Invalidate invalidate field
func (f *DateField) Invalidate() {
	f.valid = false
}

//MarshalJSON implement json.Marshaler interface
func (f *DateField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(fmt.Sprintf("%q", f.val.Format("2006-01-02"))), nil
}

//BulkUpdateValue generate a UpdateValue for bulk update
func (f *DateField) UpdateValue(val time.Time, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val.Format("2006-01-02"), isNull}
}

func (f *DateField) updateValue() *UpdateValue {
	return &UpdateValue{f.column, f.val, f.null}
}

//Where generate where by self value
func (f *DateField) where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val.Format("2006-01-02"))
}

func (f *DateField) Where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val.Format("2006-01-02"))
}

//Eq generte euqal Where
func (f *DateField) Eq(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, "=", val.Format("2006-01-02"))
}

//Neq generate not euqal Where
func (f *DateField) Neq(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, "<>", val.Format("2006-01-02"))
}

//Lt genereate less than Where
func (f *DateField) Lt(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, "<", val.Format("2006-01-02"))
}

//Gt generate great then Where
func (f *DateField) Gt(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, ">", val.Format("2006-01-02"))
}

//Lte generate less than equal Where
func (f *DateField) Lte(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, "<=", val.Format("2006-01-02"))
}

//Gte generate great than equal Where
func (f *DateField) Gte(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, ">=", val.Format("2006-01-02"))
}

//Null generate null Where
func (f *DateField) Null() *Where {
	return newWhere(f.db, f.tab, f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *DateField) NotNull() *Where {
	return newWhere(f.db, f.tab, f.column, "IS NOT", "NULL")
}

//In generate in Where
func (f *DateField) In(val []time.Time) *Where {
	l := make([]string, len(val))
	for i := range val {
		l[i] = val[i].Format("2006-01-02")
	}
	return newWhere(f.db, f.tab, f.column, "IN", fmt.Sprintf("(%s)", strings.Join(l, ", ")))

}

//LessFunc return a function for sort ModelList
func (f *DateField) LessFunc(reverse bool) func(iaddr, jaddr uintptr) int {
	if reverse {
		return func(iaddr, jaddr uintptr) int {
			iField, jField := *(**DateField)(unsafe.Pointer(iaddr + f.offset)), *(**DateField)(unsafe.Pointer(jaddr + f.offset))
			iVal, iValid, iNull := iField.Get()
			jVal, jValid, jNull := jField.Get()
			var iBit, jBit int
			if iValid {
				iBit |= 1 << 1
			}
			if !iNull {
				iBit |= 1
			}
			if jValid {
				jBit |= 1 << 1
			}
			if !jNull {
				jBit |= 1
			}
			if iBit != 3 || jBit != 3 {
				return -(iBit - jBit)
			}
			switch {
			case iVal.Before(jVal):
				return 1
			case iVal.After(jVal):
				return -1
			default:
				return 0
			}
		}
	}
	return func(iaddr, jaddr uintptr) int {
		iField, jField := *(**DateField)(unsafe.Pointer(iaddr + f.offset)), *(**DateField)(unsafe.Pointer(jaddr + f.offset))
		iVal, iValid, iNull := iField.Get()
		jVal, jValid, jNull := jField.Get()
		var iBit, jBit int
		if iValid {
			iBit |= 1 << 1
		}
		if !iNull {
			iBit |= 1
		}
		if jValid {
			jBit |= 1 << 1
		}
		if !jNull {
			jBit |= 1
		}
		if iBit != 3 || jBit != 3 {
			return iBit - jBit
		}
		switch {
		case iVal.Before(jVal):
			return -1
		case iVal.After(jVal):
			return 1
		default:
			return 0
		}
	}
}

//SortOrder return a string represent a sort order
func (f *DateField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.db, f.tab, f.column, o)
}

func (f *DateField) getDefVal() interface{} {
	return f.defVal
}

func (f *DateField) check() error {
	for _, validator := range f.validators {
		if err := validator(f); err != nil {
			return err
		}
	}
	return nil
}

func (f *DateField) ISOFormat() string {
	return f.val.Format("2006-01-02")
}

func (f *DateField) Format(format string) string {
	return f.val.Format(format)
}

//DatetimeField represent datetime, timestamp types in mysql
type DatetimeField struct {
	db         string
	tab        string
	column     string
	nullable   bool
	pk         bool
	uni        bool
	defVal     interface{}
	offset     uintptr
	val        time.Time
	valid      bool
	null       bool
	validators []Validator
	fieldName  string
	modelName  string
}

//NewDatetimeField create a DatetimeField
func NewDatetimeField(db, tab, column string, nullable, pk, uni bool, defVal interface{}, offset uintptr) *DatetimeField {
	return &DatetimeField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal, offset: offset}
}

func (f *DatetimeField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val.Format("2006-01-02 15:04:05")
}

func (f *DatetimeField) dbName() string {
	return wrap(f.db)
}

func (f *DatetimeField) tabName() string {
	return wrap(f.tab)
}

func (f *DatetimeField) columnName() string {
	return wrap(f.column)
}

func (f *DatetimeField) getModelName() string {
	return f.modelName
}

func (f *DatetimeField) getFieldName() string {
	return f.fieldName
}

func (f *DatetimeField) fullTabName() string {
	return fmt.Sprintf("%s.%s", f.dbName(), f.tabName())
}

func (f *DatetimeField) fullColName() string {
	return fmt.Sprintf("%s.%s.%s", f.dbName(), f.tabName(), f.columnName())
}

//IsPk return true if column is a primary key
func (f *DatetimeField) isPk() bool {
	return f.pk
}

//IsInc return true if column is a auto_increment column
func (f *DatetimeField) isInc() bool {
	return false
}

//IsValid return true if field has been set a value or has been scanned
func (f *DatetimeField) IsValid() bool {
	return f.valid
}

//IsNull return true if field value is null
func (f *DatetimeField) IsNull() bool {
	return f.null
}

func (f *DatetimeField) isNullable() bool {
	return f.nullable
}

//IsUni return true if column is a unique key
func (f *DatetimeField) isUni() bool {
	return f.uni
}

//Set set value
func (f *DatetimeField) Set(val time.Time) {
	f.valid, f.null, f.val = true, false, val
}

func (f *DatetimeField) SetNull() {
	f.valid, f.null, f.val = true, true, time.Time{}
}

func (f *DatetimeField) setVal(val interface{}, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val.(time.Time)
}

//SetByUpdateValue set value by UpdateValue struct
func (f *DatetimeField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(time.Time)
}

//Get get value
func (f *DatetimeField) Get() (val time.Time, valid bool, null bool) {
	return f.val, f.valid, f.null
}

//Scan implement sql.Scanner interface
func (f *DatetimeField) Scan(v interface{}) error {
	f.valid = true
	if v == nil {
		f.null = true
		return nil
	}
	f.null = false
	switch val := v.(type) {
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05", string(val))
		if err != nil {
			return err
		}
		f.val = t
	case time.Time:
		f.val = val
	default:
		return fmt.Errorf("not supported value type for DatetimeField: %T", val)
	}
	return nil
}

//Invalidate invalidate field
func (f *DatetimeField) Invalidate() {
	f.valid = false
}

//MarshalJSON implement json.Marshaler interface
func (f *DatetimeField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(fmt.Sprintf("%q", f.val.Format("2006-01-02 15:04:05"))), nil
}

//BulkUpdateValue return a UpdateValue for bulk update
func (f *DatetimeField) UpdateValue(val time.Time, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val.Format("2006-01-02 15:04:05"), isNull}
}

func (f *DatetimeField) updateValue() *UpdateValue {
	return &UpdateValue{f.column, f.val, f.null}
}

//Where generate Where by self value
func (f *DatetimeField) where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val.Format("2006-01-02 15:04:05"))
}

func (f *DatetimeField) Where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", f.val.Format("2006-01-02 15:04:05"))
}

//Eq generate equal Where
func (f *DatetimeField) Eq(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, "=", val.Format("2006-01-02 15:04:05"))
}

//Neq generate not euqal Where
func (f *DatetimeField) Neq(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, "<>", val.Format("2006-01-02 15:04:05"))
}

//Lt generate less than Where
func (f *DatetimeField) Lt(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, "<", val.Format("2006-01-02 15:04:05"))
}

//Gt generate great than Where
func (f *DatetimeField) Gt(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, ">", val.Format("2006-01-02 15:04:05"))
}

//Lte generate less then equal Where
func (f *DatetimeField) Lte(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, "<=", val.Format("2006-01-02 15:04:05"))
}

//Gte generate great than equal Where
func (f *DatetimeField) Gte(val time.Time) *Where {
	return newWhere(f.db, f.tab, f.column, ">=", val.Format("2006-01-02 15:04:05"))
}

//Null generate null Where
func (f *DatetimeField) Null() *Where {
	return newWhere(f.db, f.tab, f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *DatetimeField) NotNull() *Where {
	return newWhere(f.db, f.tab, f.column, "IS NOT", "NULL")
}

//In generate in Where
func (f *DatetimeField) In(val []time.Time) *Where {
	l := make([]string, len(val))
	for i := range val {
		l[i] = val[i].Format("2006-01-02 15:04:05")
	}
	return newWhere(f.db, f.tab, f.column, "IN", fmt.Sprintf("(%s)", strings.Join(l, ", ")))
}

//LessFunc return a function for sorting ModelList
func (f *DatetimeField) LessFunc(reverse bool) func(iaddr, jaddr uintptr) int {
	if reverse {
		return func(iaddr, jaddr uintptr) int {
			iField, jField := *(**DatetimeField)(unsafe.Pointer(iaddr + f.offset)), *(**DatetimeField)(unsafe.Pointer(jaddr + f.offset))
			iVal, iValid, iNull := iField.Get()
			jVal, jValid, jNull := jField.Get()
			var iBit, jBit int
			if iValid {
				iBit |= 1 << 1
			}
			if !iNull {
				iBit |= 1
			}
			if jValid {
				jBit |= 1 << 1
			}
			if !jNull {
				jBit |= 1
			}
			if iBit != 3 || jBit != 3 {
				return -(iBit - jBit)
			}
			switch {
			case iVal.Before(jVal):
				return 1
			case iVal.After(jVal):
				return -1
			default:
				return 0
			}
		}
	}
	return func(iaddr, jaddr uintptr) int {
		iField, jField := *(**DatetimeField)(unsafe.Pointer(iaddr + f.offset)), *(**DatetimeField)(unsafe.Pointer(jaddr + f.offset))
		iVal, iValid, iNull := iField.Get()
		jVal, jValid, jNull := jField.Get()
		var iBit, jBit int
		if iValid {
			iBit |= 1 << 1
		}
		if !iNull {
			iBit |= 1
		}
		if jValid {
			jBit |= 1 << 1
		}
		if !jNull {
			jBit |= 1
		}
		if iBit != 3 || jBit != 3 {
			return iBit - jBit
		}
		switch {
		case iVal.Before(jVal):
			return -1
		case iVal.After(jVal):
			return 1
		default:
			return 0
		}
	}
}

//SortOrder return a string represent sort order
func (f *DatetimeField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.db, f.tab, f.column, o)
}

func (f *DatetimeField) getDefVal() interface{} {
	return f.defVal
}

func (f *DatetimeField) check() error {
	for _, validator := range f.validators {
		if err := validator(f); err != nil {
			return err
		}
	}
	return nil
}

func (f *DatetimeField) Format(format string) string {
	return f.val.Format(format)
}

func (f *DatetimeField) ISOFormat() string {
	return f.val.Format("2006-01-02 15:04:05")
}

//BinaryField represent blob type in mysql
type BinaryField struct {
	db         string
	tab        string
	column     string
	nullable   bool
	pk         bool
	uni        bool
	defVal     interface{}
	offset     uintptr
	val        []byte
	valid      bool
	null       bool
	validators []Validator
	fieldName  string
	modelName  string
}

//NewBinaryField create a BinaryField
func NewBinaryField(db, tab, column string, nullable, pk, uni bool, defVal interface{}, offset uintptr) *BinaryField {
	return &BinaryField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal, offset: offset}
}

func (f *BinaryField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return fmt.Sprintf("X'%x'", f.val)
}

//Get get value
func (f *BinaryField) Get() (val []byte, isValid, isNull bool) {
	return f.val, f.valid, f.null
}

//IsValid return true if field has been set a value or has been scanned
func (f *BinaryField) IsValid() bool {
	return f.valid
}

//IsNull return true if value is null
func (f *BinaryField) IsNull() bool {
	return f.null
}

func (f *BinaryField) isNullable() bool {
	return f.nullable
}

//Scan implement sql.Scanner interface
func (f *BinaryField) Scan(val interface{}) error {
	f.valid = true
	if val == nil {
		f.null = true
		return nil
	}
	f.null = false
	b := val.([]byte)
	s := strings.Trim(string(b), "X'")
	b = []byte(s)
	decodedLen := hex.DecodedLen(len(b))
	l := make([]byte, decodedLen)
	n, err := hex.Decode(l, b)
	if err != nil {
		return err
	}
	f.val = l[:n]
	return nil
}

func (f *BinaryField) dbName() string {
	return wrap(f.db)
}

func (f *BinaryField) tabName() string {
	return wrap(f.tab)
}

func (f *BinaryField) columnName() string {
	return wrap(f.column)
}

func (f *BinaryField) getModelName() string {
	return f.modelName
}

func (f *BinaryField) getFieldName() string {
	return f.fieldName
}

func (f *BinaryField) fullTabName() string {
	return fmt.Sprintf("%s.%s", f.dbName(), f.tabName())
}

func (f *BinaryField) fullColName() string {
	return fmt.Sprintf("%s.%s.%s", f.dbName(), f.tabName(), f.columnName())
}

//IsPk return true if column is a primary key
func (f *BinaryField) isPk() bool {
	return f.pk
}

//IsInc return true if column is a auto_increment column
func (f *BinaryField) isInc() bool {
	return false
}

//IsUni return true if column is a unique key
func (f *BinaryField) isUni() bool {
	return f.uni
}

//BulkUpdateValue generate UpdateValue for bulk update
func (f *BinaryField) UpdateValue(val []byte, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, fmt.Sprintf("X'%x'", val), isNull}
}

func (f *BinaryField) updateValue() *UpdateValue {
	return &UpdateValue{f.column, f.val, f.null}
}

func (f *BinaryField) toSQL() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

//LessFunc generate function for sorting ModelList (this will return a fake function)
func (f *BinaryField) LessFunc(reverse bool) func(iaddr, jaddr uintptr) int {
	if reverse {
		return func(iaddr, jaddr uintptr) int {
			iField, jField := *(**BinaryField)(unsafe.Pointer(iaddr + f.offset)), *(**BinaryField)(unsafe.Pointer(jaddr + f.offset))
			var iBit, jBit int
			if iField.valid {
				iBit |= 1 << 1
			}
			if !iField.null {
				iBit |= 1
			}
			if jField.valid {
				jBit |= 1 << 1
			}
			if !iField.null {
				jBit |= 1
			}
			return -(iBit - jBit)
		}
	}
	return func(iaddr, jaddr uintptr) int {
		iField, jField := *(**BinaryField)(unsafe.Pointer(iaddr + f.offset)), *(**BinaryField)(unsafe.Pointer(jaddr + f.offset))
		var iBit, jBit int
		if iField.valid {
			iBit |= 1 << 1
		}
		if !iField.null {
			iBit |= 1
		}
		if jField.valid {
			jBit |= 1 << 1
		}
		if !iField.null {
			jBit |= 1
		}
		return iBit - jBit
	}
}

//Set set value
func (f *BinaryField) Set(val []byte) {
	f.valid, f.null, f.val = true, false, val
}

func (f *BinaryField) SetNull() {
	f.valid, f.null, f.val = true, true, nil
}

func (f *BinaryField) setVal(val interface{}, null bool) {
	f.valid, f.null, f.val = true, null, val.([]byte)
}

//SetByUpdateValue set value by UpdateValue struct
func (f *BinaryField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.([]byte)
}

//Invalidate invalidate field
func (f *BinaryField) Invalidate() {
	f.valid = false
}

//Eq generate equal Where
func (f *BinaryField) Eq(val []byte) *Where {
	return newWhere(f.db, f.tab, f.column, "=", val)
}

//Where generate Where by self value
func (f *BinaryField) where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", fmt.Sprintf("X'%x'", f.val))
}

func (f *BinaryField) Where() *Where {
	if f.null {
		return newWhere(f.db, f.tab, f.column, "IS", "NULL")
	}
	return newWhere(f.db, f.tab, f.column, "=", fmt.Sprintf("X'%x'", f.val))
}

//Neq generate not equal Where
func (f *BinaryField) Neq(val []byte) *Where {
	return newWhere(f.db, f.tab, f.column, "<>", val)
}

//Null generate null Where
func (f *BinaryField) Null() *Where {
	return newWhere(f.db, f.tab, f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *BinaryField) NotNull() *Where {
	return newWhere(f.db, f.tab, f.column, "IS NOT", "NULL")
}

func (f *BinaryField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.db, f.tab, f.column, o)
}

func (f *BinaryField) getDefVal() interface{} {
	return f.defVal
}

func (f *BinaryField) check() error {
	for _, validator := range f.validators {
		if err := validator(f); err != nil {
			return err
		}
	}
	return nil
}
