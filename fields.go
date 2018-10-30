package nborm

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
)

//UpdateValue is used for bulk update
type UpdateValue struct {
	column string
	val    interface{}
	null   bool
}

//StringField represent char, varchar, text type in mysql
type StringField struct {
	super  Model
	val    string
	valid  bool
	null   bool
	column string
	pk     bool
	inc    bool
	uni    bool
}

//NewStringField create a StringField
func NewStringField(model Model, column string, pk bool, inc bool, uni bool) *StringField {
	return &StringField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *StringField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

//Super get the Model which is the field stored in
func (f *StringField) Super() Model {
	return f.super
}

//Column return the column name of the field
func (f *StringField) Column() string {
	return f.column
}

//IsPk return true if the column is a primary key
func (f *StringField) IsPk() bool {
	return f.pk
}

//IsInc return true if the column is a auto_increment column
func (f *StringField) IsInc() bool {
	return f.inc
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
func (f *StringField) IsUni() bool {
	return f.uni
}

//Set set a value for the field
func (f *StringField) Set(val string, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *StringField) Where() *Where {
	if f.null {
		return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
	} else {
		return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", f.val)
	}
}

//Eq generate a equal Where
func (f *StringField) Eq(val string) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", val)

}

//Neq generate a not equal Where
func (f *StringField) Neq(val string) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<>", val)
}

//Contains generate a contains Where
func (f *StringField) Contains(val string) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "LIKE", "%"+val+"%")
}

//Null generate a Where which represent this column is null
func (f *StringField) Null() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
}

//NotNull generate a Where which represent this column is not null
func (f *StringField) NotNull() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS NOT", "NULL")
}

//In generate a in Where
func (f *StringField) In(val []string) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IN", fmt.Sprintf("(%s)", strings.Join(val, ", ")))
}

//LessFunc generate a function for sort a ModelList
func (f *StringField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getByName(im, f.Column()).(*StringField), getByName(jm, f.Column()).(*StringField)
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
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

//IntField represent int, tinyint, midint, bigint types in mysql
type IntField struct {
	super  Model
	val    int64
	valid  bool
	null   bool
	column string
	pk     bool
	inc    bool
	uni    bool
}

//NewIntField return a IntField
func NewIntField(model Model, column string, pk bool, inc bool, uni bool) *IntField {
	return &IntField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *IntField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

//Super return the model which the field is stored in
func (f *IntField) Super() Model {
	return f.super
}

//Column return the column name of the field
func (f *IntField) Column() string {
	return f.column
}

//IsPk return true if this column is primary key
func (f *IntField) IsPk() bool {
	return f.pk
}

//IsInc return true if this column is auto_increment column
func (f *IntField) IsInc() bool {
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
func (f *IntField) IsUni() bool {
	return f.uni
}

//Set set a value for this field
func (f *IntField) Set(val int64, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *IntField) Where() *Where {
	if f.null {
		return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
	}
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", f.val)
}

//Eq generate a equal Where
func (f *IntField) Eq(val int64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", val)
}

//Neq generate a not equal Where
func (f *IntField) Neq(val int64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<>", val)
}

//Lt generate a less than Where
func (f *IntField) Lt(val int64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<", val)
}

//Gt generate a great than Where
func (f *IntField) Gt(val int64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, ">", val)
}

//Lte generate a less than equal Where
func (f *IntField) Lte(val int64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<=", val)
}

//Gte generate a great than equal Where
func (f *IntField) Gte(val int64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, ">=", val)
}

//Null generate null Where
func (f *IntField) Null() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *IntField) NotNull() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS NOT", "NULL")
}

//In generate a in Where
func (f *IntField) In(val []int) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IN", toListStr(val))
}

//LessFunc return a function for sort ModelList
func (f *IntField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getByName(im, f.Column()).(*IntField), getByName(jm, f.Column()).(*IntField)
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
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

//SumFunc return a function for sum aggregate
func (f *IntField) SumFunc() func(ModelList) float64 {
	return func(l ModelList) float64 {
		var sum int
		for i := 0; i < l.Len(); i++ {
			m := l.Index(i)
			field := getByName(m, f.column)
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
				field := getByName(m, f.column)
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
			field := getByName(m, f.column)
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
	super  Model
	val    float64
	valid  bool
	null   bool
	column string
	pk     bool
	inc    bool
	uni    bool
}

//NewFloatField create new FloatField
func NewFloatField(model Model, column string, pk bool, inc bool, uni bool) *FloatField {
	return &FloatField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *FloatField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

//Super return Model which the field stored in
func (f *FloatField) Super() Model {
	return f.super
}

//Column return column name
func (f *FloatField) Column() string {
	return f.column
}

//IsPk return true if the column is primary key
func (f *FloatField) IsPk() bool {
	return f.pk
}

//IsInc return true if the column is auto_increment column
func (f *FloatField) IsInc() bool {
	return f.inc
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
func (f *FloatField) IsUni() bool {
	return f.uni
}

//Set set a value
func (f *FloatField) Set(val float64, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *FloatField) Where() *Where {
	if f.null {
		return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
	}
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", f.val)
}

//Eq generate a equal Where
func (f *FloatField) Eq(val float64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", val)
}

//Neq generate a not equal Where
func (f *FloatField) Neq(val float64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<>", val)
}

//Lt generate a less than Where
func (f *FloatField) Lt(val float64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<", val)
}

//Gt generate a great than Where
func (f *FloatField) Gt(val float64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, ">", val)
}

//Lte generate a less than equal Where
func (f *FloatField) Lte(val float64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<=", val)
}

//Gte generate a great than equal Where
func (f *FloatField) Gte(val float64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, ">=", val)
}

//Null generate a null Where
func (f *FloatField) Null() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
}

//NotNull generate a not null Where
func (f *FloatField) NotNull() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<>", "NULL")
}

//In generate in Where
func (f *FloatField) In(val []float64) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IN", toListStr(val))
}

//LessFunc return a func for sort ModelList
func (f *FloatField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getByName(im, f.Column()).(*FloatField), getByName(jm, f.Column()).(*FloatField)
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
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

//BoolField represent tinyint(1) type in mysql
type BoolField struct {
	super  Model
	val    bool
	valid  bool
	null   bool
	column string
	pk     bool
	inc    bool
	uni    bool
}

//NewBoolField return a new BoolField
func NewBoolField(model Model, column string, pk bool, inc bool, uni bool) *BoolField {
	return &BoolField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *BoolField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val
}

//Super return the Model which the field is stored in
func (f *BoolField) Super() Model {
	return f.super
}

//Column return column name
func (f *BoolField) Column() string {
	return f.column
}

//IsPk return true if the column is a primary key
func (f *BoolField) IsPk() bool {
	return f.pk
}

//IsInc return true if the column is a auto_increment column
func (f *BoolField) IsInc() bool {
	return f.inc
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
func (f *BoolField) Where() *Where {
	if f.null {
		return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
	}
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", f.val)
}

//Eq generate euqal Where
func (f *BoolField) Eq(val bool) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", val)
}

//Neq generate not equal Where
func (f *BoolField) Neq(val bool) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<>", val)
}

//Null generate null Where
func (f *BoolField) Null() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *BoolField) NotNull() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS NOT", "NULL")
}

//IsUni return true if column is a unique key
func (f *BoolField) IsUni() bool {
	return f.uni
}

//LessFunc return a func for sort ModelList
func (f *BoolField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getByName(im, f.Column()).(*BoolField), getByName(jm, f.Column()).(*BoolField)
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
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

//DateField represent date type in mysql
type DateField struct {
	super  Model
	val    time.Time
	valid  bool
	null   bool
	column string
	pk     bool
	inc    bool
	uni    bool
}

//NewDateField create a new DateField
func NewDateField(model Model, column string, pk bool, inc bool, uni bool) *DateField {
	return &DateField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *DateField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val.Format("2006-01-02")
}

//Super return Model which the Field is stored in
func (f *DateField) Super() Model {
	return f.super
}

//Column return column name
func (f *DateField) Column() string {
	return f.column
}

//IsPk return true if the column is a primary key
func (f *DateField) IsPk() bool {
	return f.pk
}

//IsInc return true if the column is a auto_increment column
func (f *DateField) IsInc() bool {
	return f.inc
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
func (f *DateField) IsUni() bool {
	return f.uni
}

//Set set a value
func (f *DateField) Set(val time.Time, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *DateField) Where() *Where {
	if f.null {
		return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
	}
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", f.val.Format("2006-01-02"))
}

//Eq generte euqal Where
func (f *DateField) Eq(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", val.Format("2006-01-02"))
}

//Neq generate not euqal Where
func (f *DateField) Neq(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<>", val.Format("2006-01-02"))
}

//Lt genereate less than Where
func (f *DateField) Lt(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<", val.Format("2006-01-02"))
}

//Gt generate great then Where
func (f *DateField) Gt(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, ">", val.Format("2006-01-02"))
}

//Lte generate less than equal Where
func (f *DateField) Lte(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<=", val.Format("2006-01-02"))
}

//Gte generate great than equal Where
func (f *DateField) Gte(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, ">=", val.Format("2006-01-02"))
}

//Null generate null Where
func (f *DateField) Null() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *DateField) NotNull() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS NOT", "NULL")
}

//In generate in Where
func (f *DateField) In(val []time.Time) *Where {
	l := make([]string, len(val))
	for i := range val {
		l[i] = val[i].Format("2006-01-02")
	}
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IN", fmt.Sprintf("(%s)", strings.Join(l, ", ")))

}

//LessFunc return a function for sort ModelList
func (f *DateField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getByName(im, f.Column()).(*DateField), getByName(jm, f.Column()).(*DateField)
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
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

//DatetimeField represent datetime, timestamp types in mysql
type DatetimeField struct {
	super  Model
	val    time.Time
	valid  bool
	null   bool
	column string
	pk     bool
	inc    bool
	uni    bool
}

//NewDatetimeField create a DatetimeField
func NewDatetimeField(model Model, column string, pk bool, inc bool, uni bool) *DatetimeField {
	return &DatetimeField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *DatetimeField) value() interface{} {
	if f.null {
		return "NULL"
	}
	return f.val.Format("2006-01-02 15:04:05")
}

//Super return Model where the field is stored in
func (f *DatetimeField) Super() Model {
	return f.super
}

//Column return column name
func (f *DatetimeField) Column() string {
	return f.column
}

//IsPk return true if column is a primary key
func (f *DatetimeField) IsPk() bool {
	return f.pk
}

//IsInc return true if column is a auto_increment column
func (f *DatetimeField) IsInc() bool {
	return f.inc
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
func (f *DatetimeField) IsUni() bool {
	return f.uni
}

//Set set value
func (f *DatetimeField) Set(val time.Time, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
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
func (f *DatetimeField) Where() *Where {
	if f.null {
		return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
	}
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", f.val.Format("2006-01-02 15:04:05"))
}

//Eq generate equal Where
func (f *DatetimeField) Eq(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", val.Format("2006-01-02 15:04:05"))
}

//Neq generate not euqal Where
func (f *DatetimeField) Neq(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<>", val.Format("2006-01-02 15:04:05"))
}

//Lt generate less than Where
func (f *DatetimeField) Lt(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<", val.Format("2006-01-02 15:04:05"))
}

//Gt generate great than Where
func (f *DatetimeField) Gt(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, ">", val.Format("2006-01-02 15:04:05"))
}

//Lte generate less then equal Where
func (f *DatetimeField) Lte(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<=", val.Format("2006-01-02 15:04:05"))
}

//Gte generate great than equal Where
func (f *DatetimeField) Gte(val time.Time) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, ">=", val.Format("2006-01-02 15:04:05"))
}

//Null generate null Where
func (f *DatetimeField) Null() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *DatetimeField) NotNull() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS NOT", "NULL")
}

//In generate in Where
func (f *DatetimeField) In(val []time.Time) *Where {
	l := make([]string, len(val))
	for i := range val {
		l[i] = val[i].Format("2006-01-02 15:04:05")
	}
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IN", fmt.Sprintf("(%s)", strings.Join(l, ", ")))
}

//LessFunc return a function for sorting ModelList
func (f *DatetimeField) LessFunc() func(Model, Model) int {
	return func(im, jm Model) int {
		iField, jField := getByName(im, f.Column()).(*DatetimeField), getByName(jm, f.Column()).(*DatetimeField)
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
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

//BinaryField represent blob type in mysql
type BinaryField struct {
	super  Model
	column string
	val    []byte
	valid  bool
	null   bool
	pk     bool
	inc    bool
	uni    bool
}

//NewBinaryField create a BinaryField
func NewBinaryField(model Model, column string, pk bool, inc bool, uni bool) *BinaryField {
	return &BinaryField{super: model, column: column, pk: pk, inc: inc, uni: uni}
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

//Column return column name
func (f *BinaryField) Column() string {
	return f.column
}

//IsPk return true if column is a primary key
func (f *BinaryField) IsPk() bool {
	return f.pk
}

//IsInc return true if column is a auto_increment column
func (f *BinaryField) IsInc() bool {
	return f.inc
}

//IsUni return true if column is a unique key
func (f *BinaryField) IsUni() bool {
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
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", val)
}

//Where generate Where by self value
func (f *BinaryField) Where() *Where {
	if f.null {
		return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
	}
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "=", fmt.Sprintf("X'%x'", f.val))
}

//Neq generate not equal Where
func (f *BinaryField) Neq(val []byte) *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "<>", val)
}

//Null generate null Where
func (f *BinaryField) Null() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS", "NULL")
}

//NotNull generate not null Where
func (f *BinaryField) NotNull() *Where {
	return newWhere(f.super.DB(), f.super.Tab(), f.column, "IS NOT", "NULL")
}

//OneToOne represent a one point on relation
type OneToOne struct {
	srcField   Field
	dstColName string
	new        func() Model
	cache      Model
	result     Model
}

//NewOneToOne create a OneToOne
func NewOneToOne(srcField Field, dstColName string, f func() Model) *OneToOne {
	return &OneToOne{srcField: srcField, dstColName: dstColName, new: f}
}

func (oto *OneToOne) srcDB() string {
	return oto.srcField.Super().DB()
}

func (oto *OneToOne) srcTab() string {
	return oto.srcField.Super().Tab()
}

func (oto *OneToOne) srcCol() string {
	return oto.srcField.Column()
}

func (oto *OneToOne) dstDB() string {
	if oto.cache == nil {
		oto.cache = oto.new()
	}
	return oto.cache.DB()
}

func (oto *OneToOne) dstTab() string {
	if oto.cache == nil {
		oto.cache = oto.new()
	}
	return oto.cache.Tab()
}

func (oto *OneToOne) dstCol() string {
	return oto.dstColName
}

//Query query related table by OneToOne relation
func (oto *OneToOne) Query() error {
	m := oto.new()
	if oto.cache == nil {
		oto.cache = m
	}
	db := dbMap[m.DB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = ?", m.DB(), m.Tab(), oto.dstColName)
	row := db.QueryRow(stmtStr, oto.srcField.value())
	if err := scanRow(m, row); err != nil {
		return err
	}
	oto.result = m
	return nil
}

//Result get query result, if no rows return nil
func (oto *OneToOne) Result() Model {
	return oto.result
}

//Model return Model which the OneToOne is stored in
func (oto *OneToOne) Model() Model {
	if oto.cache == nil {
		oto.cache = oto.new()
	}
	return oto.cache
}

//MarshalJSON implement json.Marshaler interface
func (oto *OneToOne) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(oto.result, "\t", "\t")
}

func (oto *OneToOne) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", oto.dstDB(), oto.dstTab(), oto.srcDB(), oto.srcTab(), oto.srcCol(), oto.dstDB(),
		oto.dstTab(), oto.dstCol())
}

//ForeignKey represent a one point many relation
type ForeignKey struct {
	srcField   Field
	dstColName string
	new        func() Model
	cache      Model
	result     Model
}

//NewForeignKey create a ForeignKey
func NewForeignKey(srcField Field, dstColName string, f func() Model) *ForeignKey {
	return &ForeignKey{srcField: srcField, dstColName: dstColName, new: f}
}

func (fk *ForeignKey) srcDB() string {
	return fk.srcField.Super().DB()
}

func (fk *ForeignKey) srcTab() string {
	return fk.srcField.Super().Tab()
}

func (fk *ForeignKey) srcCol() string {
	return fk.srcField.Column()
}

func (fk *ForeignKey) dstDB() string {
	if fk.cache == nil {
		fk.cache = fk.new()
	}
	return fk.cache.DB()
}

func (fk *ForeignKey) dstTab() string {
	if fk.cache == nil {
		fk.cache = fk.new()
	}
	return fk.cache.Tab()
}

func (fk *ForeignKey) dstCol() string {
	return fk.dstColName
}

//Query query related table by this relation
func (fk *ForeignKey) Query() error {
	m := fk.new()
	db := dbMap[m.DB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = ?", m.DB(), m.Tab(), fk.dstColName)
	row := db.QueryRow(stmtStr, fk.srcField.value())
	if err := scanRow(m, row); err != nil {
		return err
	}
	fk.result = m
	return nil
}

//Result get query result, if no rows return nil
func (fk *ForeignKey) Result() Model {
	return fk.result
}

//MarshalJSON implement json.Marshaler interface
func (fk *ForeignKey) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(fk.Result(), "\t", "\t")
}

//Model return the example Model
func (fk *ForeignKey) Model() Model {
	if fk.cache == nil {
		fk.cache = fk.new()
	}
	return fk.cache
}

func (fk *ForeignKey) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", fk.dstDB(), fk.dstTab(), fk.srcDB(), fk.srcTab(), fk.srcCol(), fk.dstDB(),
		fk.dstTab(), fk.dstCol())
}

//ReverseForeignKey represent many point one relation
type ReverseForeignKey struct {
	srcField   Field
	dstColName string
	new        func() ModelList
	cache      ModelList
	result     ModelList
}

//NewReverseForeignKey create ReverseForeignKey
func NewReverseForeignKey(srcField Field, dstColName string, f func() ModelList) *ReverseForeignKey {
	return &ReverseForeignKey{srcField: srcField, dstColName: dstColName, new: f}
}

func (rfk *ReverseForeignKey) srcDB() string {
	return rfk.srcField.Super().DB()
}

func (rfk *ReverseForeignKey) srcTab() string {
	return rfk.srcField.Super().DB()
}

func (rfk *ReverseForeignKey) srcCol() string {
	return rfk.srcField.Column()
}

func (rfk *ReverseForeignKey) dstDB() string {
	if rfk.cache == nil {
		rfk.cache = rfk.new()
	}
	return rfk.cache.Model().DB()
}

func (rfk *ReverseForeignKey) dstTab() string {
	if rfk.cache == nil {
		rfk.cache = rfk.new()
	}
	return rfk.cache.Model().Tab()
}

func (rfk *ReverseForeignKey) dstCol() string {
	return rfk.dstColName
}

//All query all records in related table by this relation
func (rfk *ReverseForeignKey) All() error {
	l := rfk.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = ?", l.Model().DB(), l.Model().Tab(), rfk.dstColName)
	rows, err := db.Query(stmtStr, rfk.srcField.value())
	if err != nil {
		return err
	}
	if err = scanRows(l, rows); err != nil {
		return err
	}
	rfk.result = l
	return nil
}

//AllWithFoundRows query all records in related table by this relation and the number of found rows
func (rfk *ReverseForeignKey) AllWithFoundRows(sorter *Sorter, pager *Pager) (int, error) {
	l := rfk.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s WHERE %s = ? %s %s", l.Model().DB(), l.Model().Tab(),
		rfk.dstColName, sorter.toSQL(), pager.toSQL())
	rows, err := db.Query(stmtStr, rfk.srcField.value())
	if err != nil {
		return -1, err
	}
	if err = scanRows(l, rows); err != nil {
		return -1, err
	}
	rfk.result = l
	return getFoundRows(db)
}

//Query query related table by this relation
func (rfk *ReverseForeignKey) Query(where *Where) error {
	l := rfk.new()
	db := dbMap[l.Model().DB()]
	whereStr, whereList := where.toAndClause()
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = ? %s", l.Model().DB(), l.Model().Tab(), rfk.dstColName, whereStr)
	rows, err := db.Query(stmtStr, append([]interface{}{rfk.srcField.value()}, whereList...)...)
	if err != nil {
		return err
	}
	if err := scanRows(l, rows); err != nil {
		return err
	}
	rfk.result = l
	return nil
}

//QueryWithFoundRows query related table by this realtion and number of found rows
func (rfk *ReverseForeignKey) QueryWithFoundRows(where *Where, sorter *Sorter, pager *Pager) (int, error) {
	l := rfk.new()
	db := dbMap[l.Model().DB()]
	whereStr, whereList := where.toAndClause()
	stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s WHERE %s = ? %s %s %s", l.Model().DB(), l.Model().Tab(),
		rfk.dstColName, whereStr, sorter.toSQL(), pager.toSQL())
	rows, err := db.Query(stmtStr, append([]interface{}{rfk.srcField.value()}, whereList...))
	if err != nil {
		return -1, err
	}
	if err := scanRows(l, rows); err != nil {
		return -1, err
	}
	rfk.result = l
	return getFoundRows(db)
}

//Result return result of query, if no rows return nil
func (rfk *ReverseForeignKey) Result() ModelList {
	return rfk.result
}

//MarshalJSON implement json.Marshaler interface
func (rfk *ReverseForeignKey) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(rfk.Result(), "\t", "\t")
}

//Model return example Model
func (rfk *ReverseForeignKey) Model() Model {
	if rfk.cache == nil {
		rfk.cache = rfk.new()
	}
	return rfk.cache.Model()
}

func (rfk *ReverseForeignKey) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", rfk.dstDB(), rfk.dstTab(), rfk.srcDB(), rfk.srcTab(), rfk.srcCol(), rfk.dstDB(),
		rfk.dstTab(), rfk.dstCol())
}

//ManyToMany represent many point many relation
type ManyToMany struct {
	srcField      Field
	midLeftField  Field
	midRightField Field
	dstColName    string
	new           func() ModelList
	cache         ModelList
	result        ModelList
}

//NewManyToMany create ManyToMany
func NewManyToMany(srcField, midLeftField, midRightField Field, dstColName string, f func() ModelList) *ManyToMany {
	if midLeftField.Super() != midRightField.Super() {
		panic("nborm.NewManyToMany() error: require the same middle tab")
	}
	return &ManyToMany{srcField: srcField, midLeftField: midLeftField, midRightField: midRightField, dstColName: dstColName, new: f}
}

func (mtm *ManyToMany) srcDB() string {
	return mtm.srcField.Super().DB()
}

func (mtm *ManyToMany) srcTab() string {
	return mtm.srcField.Super().Tab()
}

func (mtm *ManyToMany) srcCol() string {
	return mtm.srcField.Column()
}

func (mtm *ManyToMany) midDB() string {
	return mtm.midLeftField.Super().DB()
}

func (mtm *ManyToMany) midTab() string {
	return mtm.midLeftField.Super().Tab()
}

func (mtm *ManyToMany) midLeftCol() string {
	return mtm.midLeftField.Column()
}

func (mtm *ManyToMany) midRightCol() string {
	return mtm.midRightField.Column()
}

func (mtm *ManyToMany) dstDB() string {
	if mtm.cache == nil {
		mtm.cache = mtm.new()
	}
	return mtm.cache.Model().DB()
}

func (mtm *ManyToMany) dstTab() string {
	if mtm.cache == nil {
		mtm.cache = mtm.new()
	}
	return mtm.cache.Model().Tab()
}

func (mtm *ManyToMany) dstCol() string {
	return mtm.dstColName
}

//All query all records in related table by this relation
func (mtm *ManyToMany) All() error {
	l := mtm.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = ?",
		l.Model().DB(), l.Model().Tab(), mtm.srcDB(), mtm.srcTab(), mtm.midDB(), mtm.midTab(), mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.midDB(),
		mtm.midTab(), mtm.midLeftCol(), l.Model().DB(), l.Model().Tab(), mtm.midDB(), mtm.midTab(), mtm.midRightCol(), l.Model().DB(), l.Model().Tab(),
		mtm.dstColName, mtm.srcDB(), mtm.srcTab(), mtm.srcCol())
	rows, err := db.Query(stmtStr, mtm.srcField.value())
	if err != nil {
		return err
	}
	if err = scanRows(l, rows); err != nil {
		return err
	}
	mtm.result = l
	return nil
}

//AllWithFoundRows query all records in related table and number of found rows by this relation
func (mtm *ManyToMany) AllWithFoundRows(sorter *Sorter, pager *Pager) (int, error) {
	l := mtm.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = ? %s %s",
		l.Model().DB(), l.Model().Tab(), mtm.srcDB(), mtm.srcTab(), mtm.midDB(), mtm.midTab(), mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.midDB(),
		mtm.midTab(), mtm.midLeftCol(), l.Model().DB(), l.Model().Tab(), mtm.midDB(), mtm.midTab(), mtm.midRightCol(), l.Model().DB(), l.Model().Tab(),
		mtm.dstColName, mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), sorter.toSQL(), pager.toSQL())
	rows, err := db.Query(stmtStr, mtm.srcField.value())
	if err != nil {
		return -1, err
	}
	if err = scanRows(l, rows); err != nil {
		return -1, err
	}
	mtm.result = l
	return getFoundRows(db)
}

//Query query records in related table by this relation
func (mtm *ManyToMany) Query(where *Where) error {
	l := mtm.new()
	db := dbMap[l.Model().DB()]
	whereStr, whereList := where.toAndClause()
	stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = ? %s",
		l.Model().DB(), l.Model().Tab(), mtm.srcDB(), mtm.srcTab(), mtm.midDB(), mtm.midTab(), mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.midDB(),
		mtm.midTab(), mtm.midLeftCol(), l.Model().DB(), l.Model().Tab(), mtm.midDB(), mtm.midTab(), mtm.midRightCol(), l.Model().DB(), l.Model().Tab(),
		mtm.dstColName, mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), whereStr)
	rows, err := db.Query(stmtStr, append([]interface{}{mtm.srcField.value()}, whereList...))
	if err != nil {
		return err
	}
	if err = scanRows(mtm.result, rows); err != nil {
		return err
	}
	mtm.result = l
	return nil
}

//QueryWithFoundRows query records in related table and number of found rows by this relation
func (mtm *ManyToMany) QueryWithFoundRows(where *Where, sorter *Sorter, pager *Pager) (int, error) {
	l := mtm.new()
	db := dbMap[l.Model().DB()]
	whereStr, whereList := where.toAndClause()
	stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = ? %s %s %s",
		l.Model().DB(), l.Model().Tab(), mtm.srcDB(), mtm.srcTab(), mtm.midDB(), mtm.midTab(), mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.midDB(),
		mtm.midTab(), mtm.midLeftCol(), l.Model().DB(), l.Model().Tab(), mtm.midDB(), mtm.midTab(), mtm.midRightCol(), l.Model().DB(), l.Model().Tab(),
		mtm.dstColName, mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), whereStr, sorter.toSQL(), pager.toSQL())
	rows, err := db.Query(stmtStr, append([]interface{}{mtm.srcField.value()}, whereList...)...)
	if err != nil {
		return -1, err
	}
	if err = scanRows(mtm.result, rows); err != nil {
		return -1, err
	}
	mtm.result = l
	return getFoundRows(db)
}

//Add add a relation record to middle table
func (mtm *ManyToMany) Add(m Model) error {
	dstField := getByName(m, mtm.dstColName)
	db := dbMap[mtm.midDB()]
	stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (?, ?)", mtm.midDB(), mtm.midTab(), mtm.midLeftCol(), mtm.midRightCol())
	_, err := db.Exec(stmtStr, mtm.srcField.value(), dstField.value())
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
			return nil
		}
		return err
	}
	mtm.result = nil
	return nil
}

//Remove remove a record from middle table
func (mtm *ManyToMany) Remove(m Model) error {
	l := mtm.new()
	db := dbMap[mtm.midDB()]
	dstField := getByName(m, mtm.dstColName)
	stmtStr := fmt.Sprintf("DELETE FROM %s.%s WHERE %s = ? AND %s = ?", l.Model().DB(), l.Model().Tab(), mtm.midLeftCol(), mtm.midRightCol())
	_, err := db.Exec(stmtStr, mtm.srcField.value(), dstField.value())
	if err != nil {
		return err
	}
	mtm.result = nil
	return nil
}

//Result return query result
func (mtm *ManyToMany) Result() ModelList {
	return mtm.result
}

//MarshalJSON implement json.Marshaler interface
func (mtm *ManyToMany) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(mtm.Result(), "\t", "\t")
}

//Model return example Model
func (mtm *ManyToMany) Model() Model {
	if mtm.cache == nil {
		mtm.cache = mtm.new()
	}
	return mtm.cache.Model()
}

func (mtm *ManyToMany) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s", mtm.midDB(), mtm.midTab(), mtm.srcDB(), mtm.srcTab(),
		mtm.srcCol(), mtm.midDB(), mtm.midTab(), mtm.midLeftCol(), mtm.dstDB(), mtm.dstTab(), mtm.midDB(), mtm.midTab(), mtm.midRightCol(), mtm.dstDB(),
		mtm.dstTab(), mtm.dstCol())
}
