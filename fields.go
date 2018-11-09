package nborm

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//UpdateValue is used for bulk update
type UpdateValue struct {
	column string
	val    interface{}
	null   bool
}

//StringField represent char, varchar, text type in mysql
type StringField struct {
	db       string
	tab      string
	column   string
	nullable bool
	pk       bool
	uni      bool
	defVal   interface{}
	val      string
	valid    bool
	null     bool
}

//NewStringField create a StringField
func NewStringField(db, tab, column string, nullable, pk, uni bool, defVal interface{}) *StringField {
	return &StringField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal}
}

func (f *StringField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

func (f *StringField) dbName() string {
	return f.db
}

func (f *StringField) tabName() string {
	return f.tab
}

func (f *StringField) columnName() string {
	return f.column
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

//IsUni return true if this column is a unique key
func (f *StringField) isUni() bool {
	return f.uni
}

//Set set a value for the field
func (f *StringField) Set(val string, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *StringField) BulkUpdateValue(val string, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val, isNull}
}

//Where generate a Where by self value
func (f *StringField) where() *Where {
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
	return newWhere(f.db, f.tab, f.column, "IN", fmt.Sprintf("(%s)", strings.Join(val, ", ")))
}

//LessFunc generate a function for sort a ModelList
func (f *StringField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getFieldByName(im, f.columnName()).(*StringField), getFieldByName(jm, f.columnName()).(*StringField)
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

//IntField represent int, tinyint, midint, bigint types in mysql
type IntField struct {
	db       string
	tab      string
	column   string
	nullable bool
	pk       bool
	inc      bool
	uni      bool
	defVal   interface{}
	val      int64
	valid    bool
	null     bool
}

//NewIntField return a IntField
func NewIntField(db, tab, column string, nullable, pk, inc, uni bool, defVal interface{}) *IntField {
	return &IntField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, inc: inc, uni: uni, defVal: defVal}
}

func (f *IntField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

func (f *IntField) dbName() string {
	return f.db
}

func (f *IntField) tabName() string {
	return f.tab
}

func (f *IntField) columnName() string {
	return f.column
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

//IsUni return true if the column is a unique key
func (f *IntField) isUni() bool {
	return f.uni
}

//Set set a value for this field
func (f *IntField) Set(val int64, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *IntField) BulkUpdateValue(val int64, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val, isNull}
}

//Where generate a Where by self value
func (f *IntField) where() *Where {
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
func (f *IntField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getFieldByName(im, f.columnName()).(*IntField), getFieldByName(jm, f.columnName()).(*IntField)
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

//SumFunc return a function for sum aggregate
func (f *IntField) SumFunc() func(ModelList) float64 {
	return func(l ModelList) float64 {
		var sum int
		for i := 0; i < l.Len(); i++ {
			m := l.Index(i)
			field := getFieldByName(m, f.column)
			val, valid, null := field.(*IntField).Get()
			if valid && !null {
				sum += int(val)
			}
		}
		return float64(sum)
	}
}

//AvgFunc return a funcion for average aggregate
func (f *IntField) AvgFunc(ignoreNull bool) func(ModelList) float64 {
	if ignoreNull {
		return func(l ModelList) float64 {
			var num, sum int
			for i := 0; i < l.Len(); i++ {
				m := l.Index(i)
				field := getFieldByName(m, f.column)
				val, valid, null := field.(*IntField).Get()
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
	return func(l ModelList) float64 {
		var num, sum int
		for i := 0; i < l.Len(); i++ {
			m := l.Index(i)
			field := getFieldByName(m, f.column)
			val, valid, null := field.(*IntField).Get()
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
	db       string
	tab      string
	column   string
	nullable bool
	pk       bool
	uni      bool
	defVal   interface{}
	val      float64
	valid    bool
	null     bool
}

//NewFloatField create new FloatField
func NewFloatField(db, tab, column string, nullable, pk, uni bool, defVal interface{}) *FloatField {
	return &FloatField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal}
}

func (f *FloatField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

func (f *FloatField) dbName() string {
	return f.db
}

func (f *FloatField) tabName() string {
	return f.tab
}

func (f *FloatField) columnName() string {
	return f.column
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

//IsUni return true if the column is a unique key
func (f *FloatField) isUni() bool {
	return f.uni
}

//Set set a value
func (f *FloatField) Set(val float64, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *FloatField) BulkUpdateValue(val float64, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val, isNull}
}

//Where generate a Where by self value
func (f *FloatField) where() *Where {
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
func (f *FloatField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getFieldByName(im, f.columnName()).(*FloatField), getFieldByName(jm, f.columnName()).(*FloatField)
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

//BoolField represent tinyint(1) type in mysql
type BoolField struct {
	db       string
	tab      string
	column   string
	nullable bool
	pk       bool
	uni      bool
	defVal   interface{}
	val      bool
	valid    bool
	null     bool
}

//NewBoolField return a new BoolField
func NewBoolField(db, tab, column string, nullable, pk, uni bool, defVal interface{}) *BoolField {
	return &BoolField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal}
}

func (f *BoolField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

func (f *BoolField) dbName() string {
	return f.db
}

func (f *BoolField) tabName() string {
	return f.tab
}

func (f *BoolField) columnName() string {
	return f.column
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

//Set set a value
func (f *BoolField) Set(val bool, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *BoolField) BulkUpdateValue(val bool, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val, isNull}
}

//Where generate a Where by self value
func (f *BoolField) where() *Where {
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
func (f *BoolField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getFieldByName(im, f.columnName()).(*BoolField), getFieldByName(jm, f.columnName()).(*BoolField)
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

//DateField represent date type in mysql
type DateField struct {
	db       string
	tab      string
	column   string
	nullable bool
	pk       bool
	uni      bool
	defVal   interface{}
	val      time.Time
	valid    bool
	null     bool
}

//NewDateField create a new DateField
func NewDateField(db, tab, column string, nullable, pk, uni bool, defVal interface{}) *DateField {
	return &DateField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal}
}

func (f *DateField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val.Format("2006-01-02")
}

func (f *DateField) dbName() string {
	return f.db
}

func (f *DateField) tabName() string {
	return f.tab
}

func (f *DateField) columnName() string {
	return f.column
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

//IsUni return true if the column is a unique key
func (f *DateField) isUni() bool {
	return f.uni
}

//Set set a value
func (f *DateField) Set(val time.Time, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *DateField) BulkUpdateValue(val time.Time, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val.Format("2006-01-02"), isNull}
}

//Where generate where by self value
func (f *DateField) where() *Where {
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
func (f *DateField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getFieldByName(im, f.columnName()).(*DateField), getFieldByName(jm, f.columnName()).(*DateField)
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

//DatetimeField represent datetime, timestamp types in mysql
type DatetimeField struct {
	db       string
	tab      string
	column   string
	nullable bool
	pk       bool
	uni      bool
	defVal   interface{}
	val      time.Time
	valid    bool
	null     bool
}

//NewDatetimeField create a DatetimeField
func NewDatetimeField(db, tab, column string, nullable, pk, uni bool, defVal interface{}) *DatetimeField {
	return &DatetimeField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal}
}

func (f *DatetimeField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val.Format("2006-01-02 15:04:05")
}

func (f *DatetimeField) dbName() string {
	return f.db
}

func (f *DatetimeField) tabName() string {
	return f.tab
}

func (f *DatetimeField) columnName() string {
	return f.column
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

//IsUni return true if column is a unique key
func (f *DatetimeField) isUni() bool {
	return f.uni
}

//Set set value
func (f *DatetimeField) Set(val time.Time, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *DatetimeField) BulkUpdateValue(val time.Time, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val.Format("2006-01-02 15:04:05"), isNull}
}

//Where generate Where by self value
func (f *DatetimeField) where() *Where {
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
func (f *DatetimeField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getFieldByName(im, f.columnName()).(*DatetimeField), getFieldByName(jm, f.columnName()).(*DatetimeField)
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

//BinaryField represent blob type in mysql
type BinaryField struct {
	db       string
	tab      string
	column   string
	nullable bool
	pk       bool
	uni      bool
	defVal   interface{}
	val      []byte
	valid    bool
	null     bool
}

//NewBinaryField create a BinaryField
func NewBinaryField(db, tab, column string, nullable, pk, uni bool, defVal interface{}) *BinaryField {
	return &BinaryField{db: db, tab: tab, column: column, nullable: nullable, pk: pk, uni: uni, defVal: defVal}
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

//Scan implement sql.Scanner interface
func (f *BinaryField) Scan(val interface{}) error {
	f.valid = true
	if val == nil {
		f.null = true
		return nil
	}
	f.null = false
	f.val = val.([]byte)
	return nil
}

func (f *BinaryField) dbName() string {
	return f.db
}

func (f *BinaryField) tabName() string {
	return f.tab
}

func (f *BinaryField) columnName() string {
	return f.column
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
func (f *BinaryField) BulkUpdateValue(val []byte, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, fmt.Sprintf("X'%x'", val), isNull}
}

func (f *BinaryField) toSQL() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

//LessFunc generate function for sorting ModelList (this will return a fake function)
func (f *BinaryField) LessFunc() func(Model, Model) int {
	return func(iMod, jMod Model) int {
		return 0
	}
}

//Set set value
func (f *BinaryField) Set(val []byte, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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

//OneToOne represent a one point on relation
type OneToOne struct {
	srcDB   string
	srcTab  string
	srcCol  string
	dstDB   string
	dstTab  string
	dstCol  string
	srcValF func() interface{}
}

//NewOneToOne create a OneToOne
func NewOneToOne(srcDB, srcTab, srcCol, dstDB, dstTab, dstCol string, srcValF func() interface{}) *OneToOne {
	return &OneToOne{srcDB, srcTab, srcCol, dstDB, dstTab, dstCol, srcValF}
}

//Query query related table by OneToOne relation
func (oto *OneToOne) Query(m Model) error {
	if m.DB() != oto.dstDB || m.Tab() != oto.dstTab {
		return fmt.Errorf("nborm.OneToOne.Query() error: required %s.%s supported %s.%s", oto.dstDB, oto.dstTab, m.DB(), m.Tab())
	}
	return relationQuery(m, oto, nil, nil, nil)
}

func (oto *OneToOne) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", oto.dstDB, oto.dstTab, oto.srcDB, oto.srcTab, oto.srcCol, oto.dstDB,
		oto.dstTab, oto.dstCol)
}

func (oto *OneToOne) where() *Where {
	return newWhere(oto.srcDB, oto.srcTab, oto.srcCol, "=", oto.srcValF())
}

func (oto *OneToOne) getSrcDB() string {
	return oto.srcDB
}

func (oto *OneToOne) getSrcTab() string {
	return oto.srcTab
}

//ForeignKey represent a one point many relation
type ForeignKey struct {
	srcDB   string
	srcTab  string
	srcCol  string
	dstDB   string
	dstTab  string
	dstCol  string
	srcValF func() interface{}
}

//NewForeignKey create a ForeignKey
func NewForeignKey(srcDB, srcTab, srcCol, dstDB, dstTab, dstCol string, srcValF func() interface{}) *ForeignKey {
	return &ForeignKey{srcDB, srcTab, srcCol, dstDB, dstTab, dstCol, srcValF}
}

//Query query related table by this relation
func (fk *ForeignKey) Query(m Model) error {
	if m.DB() != fk.dstDB || m.Tab() != fk.dstTab {
		return fmt.Errorf("nborm.ForeignKey.Query() error: required %s.%s supported %s.%s", fk.dstDB, fk.dstTab, m.DB(), m.Tab())
	}
	return relationQuery(m, fk, nil, nil, nil)
}

func (fk *ForeignKey) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", fk.dstDB, fk.dstTab, fk.srcDB, fk.srcTab, fk.srcCol, fk.dstDB,
		fk.dstTab, fk.dstCol)
}

func (fk *ForeignKey) where() *Where {
	return newWhere(fk.srcDB, fk.srcTab, fk.srcCol, "=", fk.srcValF())
}

func (fk *ForeignKey) getSrcDB() string {
	return fk.srcDB
}

func (fk *ForeignKey) getSrcTab() string {
	return fk.srcTab
}

//ReverseForeignKey represent many point one relation
type ReverseForeignKey struct {
	srcDB   string
	srcTab  string
	srcCol  string
	dstDB   string
	dstTab  string
	dstCol  string
	srcValF func() interface{}
}

//NewReverseForeignKey create ReverseForeignKey
func NewReverseForeignKey(srcDB, srcTab, srcCol, dstDB, dstTab, dstCol string, srcValF func() interface{}) *ReverseForeignKey {
	return &ReverseForeignKey{srcDB, srcTab, srcCol, dstDB, dstTab, dstCol, srcValF}
}

//All query all records in related table by this relation
func (rfk *ReverseForeignKey) All(l ModelList, sorter *Sorter, pager *Pager) error {
	if l.Model().DB() != rfk.dstDB || l.Model().Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.All() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab, l.Model().DB(), l.Model().Tab())
	}
	return relationQuery(l, rfk, nil, sorter, pager)
}

//AllWithFoundRows query all records in related table by this relation and the number of found rows
func (rfk *ReverseForeignKey) AllWithFoundRows(l ModelList, sorter *Sorter, pager *Pager) (int, error) {
	if l.Model().DB() != rfk.dstDB || l.Model().Tab() != rfk.dstTab {
		return -1, fmt.Errorf("nborm.ReverseForeignKey.AllWithFoundRows() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab,
			l.Model().DB(), l.Model().Tab())
	}
	return relationQueryWithFoundRows(l, rfk, nil, sorter, pager)
}

//Query query related table by this relation
func (rfk *ReverseForeignKey) Query(l ModelList, where *Where, sorter *Sorter, pager *Pager) error {
	if l.Model().DB() != rfk.dstDB || l.Model().Tab() != rfk.dstTab {
		return fmt.Errorf("nborm.ReverseForeignKey.Query() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab,
			l.Model().DB(), l.Model().Tab())
	}
	return relationQuery(l, rfk, where, sorter, pager)
}

//QueryWithFoundRows query related table by this realtion and number of found rows
func (rfk *ReverseForeignKey) QueryWithFoundRows(l ModelList, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	if l.Model().DB() != rfk.dstDB || l.Model().Tab() != rfk.dstTab {
		return -1, fmt.Errorf("nborm.ReverseForeignKey.QueryWithFoundRows() error: required %s.%s supported %s.%s", rfk.dstDB, rfk.dstTab,
			l.Model().DB(), l.Model().Tab())
	}
	return relationQueryWithFoundRows(l, rfk, where, sorter, pager)
}

func (rfk *ReverseForeignKey) Add(m Model) error {
	if rfk.dstDB != m.DB() || rfk.dstTab != m.Tab() {
		return fmt.Errorf("nborm.ReverseForeignKey.Add() error: type inconsitent, required %s.%s supported %s.%s", rfk.dstDB,
			rfk.dstTab, m.DB(), m.Tab())
	}
	return relationAddOne(rfk, m)
}

func (rfk *ReverseForeignKey) Remove(m Model) error {
	if rfk.dstDB != m.DB() || rfk.dstTab != m.Tab() {
		return fmt.Errorf("nborm.ReverseForeignKey.Remove() error: type inconsitent, required %s.%s supported %s.%s", rfk.dstDB,
			rfk.dstTab, m.DB(), m.Tab())
	}
	return relationRemoveOne(rfk, m)
}

func (rfk *ReverseForeignKey) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", rfk.dstDB, rfk.dstTab, rfk.srcDB, rfk.srcTab, rfk.srcCol, rfk.dstDB,
		rfk.dstTab, rfk.dstCol)
}

func (rfk *ReverseForeignKey) where() *Where {
	return newWhere(rfk.srcDB, rfk.srcTab, rfk.srcCol, "=", rfk.srcValF())
}

func (rfk *ReverseForeignKey) getSrcDB() string {
	return rfk.srcDB
}

func (rfk *ReverseForeignKey) getSrcTab() string {
	return rfk.srcTab
}

//ManyToMany represent many point many relation
type ManyToMany struct {
	srcDB       string
	srcTab      string
	srcCol      string
	midDB       string
	midTab      string
	midLeftCol  string
	midRightCol string
	dstDB       string
	dstTab      string
	dstCol      string
	srcValF     func() interface{}
}

//NewManyToMany create ManyToMany
func NewManyToMany(srcDB, srcTab, srcCol, midDB, midTab, midLeftCol, midRightCol, dstDB, dstTab, dstCol string, srcValF func() interface{}) *ManyToMany {
	return &ManyToMany{srcDB, srcTab, srcCol, midDB, midTab, midLeftCol, midRightCol, dstDB, dstTab, dstCol, srcValF}
}

//All query all records in related table by this relation
func (mtm *ManyToMany) All(l ModelList, sorter *Sorter, pager *Pager) error {
	if l.Model().DB() != mtm.dstDB || l.Model().Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.All() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, l.Model().DB(), l.Model().Tab())
	}
	return relationQuery(l, mtm, nil, sorter, pager)
}

//AllWithFoundRows query all records in related table and number of found rows by this relation
func (mtm *ManyToMany) AllWithFoundRows(l ModelList, sorter *Sorter, pager *Pager) (int, error) {
	if l.Model().DB() != mtm.dstDB || l.Model().Tab() != mtm.dstTab {
		return -1, fmt.Errorf("nborm.ManyToMany.AllWithFoundRows() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, l.Model().DB(), l.Model().Tab())
	}
	return relationQueryWithFoundRows(l, mtm, nil, sorter, pager)
}

//Query query records in related table by this relation
func (mtm *ManyToMany) Query(l ModelList, where *Where, sorter *Sorter, pager *Pager) error {
	if l.Model().DB() != mtm.dstDB || l.Model().Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.Query() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, l.Model().DB(), l.Model().Tab())
	}
	return relationQuery(l, mtm, where, sorter, pager)
}

//QueryWithFoundRows query records in related table and number of found rows by this relation
func (mtm *ManyToMany) QueryWithFoundRows(l ModelList, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	if l.Model().DB() != mtm.dstDB || l.Model().Tab() != mtm.dstTab {
		return -1, fmt.Errorf("nborm.ManyToMany.QueryWithFoundRows() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, l.Model().DB(), l.Model().Tab())
	}
	return relationQueryWithFoundRows(l, mtm, where, sorter, pager)
}

//Add add a relation record to middle table
func (mtm *ManyToMany) Add(m Model) error {
	if m.DB() != mtm.dstDB || m.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.Add() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, m.DB(), m.Tab())
	}
	return relationAddOne(mtm, m)
}

//Remove remove a record from middle table
func (mtm *ManyToMany) Remove(m Model) error {
	if m.DB() != mtm.dstDB || m.Tab() != mtm.dstTab {
		return fmt.Errorf("nborm.ManyToMany.Remove() error: require %s.%s supported %s.%s", mtm.dstDB, mtm.dstTab, m.DB(), m.Tab())
	}
	return relationRemoveOne(mtm, m)
}

func (mtm *ManyToMany) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s", mtm.midDB, mtm.midTab, mtm.srcDB, mtm.srcTab,
		mtm.srcCol, mtm.midDB, mtm.midTab, mtm.midLeftCol, mtm.dstDB, mtm.dstTab, mtm.midDB, mtm.midTab, mtm.midRightCol, mtm.dstDB,
		mtm.dstTab, mtm.dstCol)
}

func (mtm *ManyToMany) where() *Where {
	return newWhere(mtm.srcDB, mtm.srcTab, mtm.srcCol, "=", mtm.srcValF())
}

func (mtm *ManyToMany) getSrcDB() string {
	return mtm.srcDB
}

func (mtm *ManyToMany) getSrcTab() string {
	return mtm.srcTab
}
