package nborm

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/wangjun861205/nbcolor"
)

// Field Field
type Field interface {
	ClauseField
	ValueField
	dup() Field
	Init(Model Model, colName, fieldName, formName string, uriName string, index int)
	// binding.Binding
	binding.BindingBody
	binding.StructValidator
}

// FieldInfo Field的基本信息
type FieldInfo struct {
	ColName   string
	FieldName string
	Field     Field
	Index     int
}

// FieldInfoList FieldInfo的列表
type FieldInfoList []FieldInfo

// FieldList Field的列表
type FieldList []Field

// String 字符串Field
type String struct {
	clauseField
	stringValueField
}

// Init 初始化方法
func (f *String) Init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.stringValueField
	}
	f.stringValueField.init(model, colName, fieldName, formName, uriName, index)
}

// String 实现Stringer接口
func (f String) String() string {
	return f.stringValueField.String()
}

func (f *String) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *String) dup() Field {
	nf := *f
	return &nf
}

func (f *String) Name() string {
	return f.formName()
}

func (f *String) Bind(req *http.Request, value interface{}) error {
	val, ok := value.(string)
	if !ok {
		// return fmt.Errorf("invalid binding value type(%T)", value)
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid value type for Bind() (value type: %T)", value), nil)
	}
	f.SetString(val)
	return nil
}

func (f *String) BindBody(b []byte, value interface{}) error {
	fmt.Println(nbcolor.Red("==========================="))
	return nil
}

func (f *String) ValidateStruct(interface{}) error {
	return nil
}

func (f *String) Engine() interface{} {
	return nil
}

//=============================================================================================================

// Int 整数型字段
type Int struct {
	clauseField
	intValueField
}

// Init 初始化方法
func (f *Int) Init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.intValueField
	}
	f.intValueField.init(model, colName, fieldName, formName, uriName, index)
}

func (f Int) String() string {
	return f.intValueField.String()
}

func (f *Int) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Int) dup() Field {
	nf := *f
	return &nf
}

func (f *Int) Name() string {
	return f.formName()
}

func (f *Int) Bind(req *http.Request, value interface{}) error {
	switch val := value.(type) {
	case string:
		i64, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		f.SetInt(int(i64))
	case int:
		f.SetInt(val)
	case float32:
		f.SetInt(int(val))
	case float64:
		f.SetInt(int(val))
	default:
		// return fmt.Errorf("invalid binding value type (%T)", value)
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid value type for Bind() (value type: %T)", value), nil)
	}
	return nil
}

func (f *Int) BindBody(b []byte, value interface{}) error {
	fmt.Println(nbcolor.Red("==========================="))
	return nil
}

func (f *Int) ValidateStruct(interface{}) error {
	return nil
}

func (f *Int) Engine() interface{} {
	return nil
}

//=======================================================================================================

// Date 日期型字段
type Date struct {
	clauseField
	dateValueField
}

// Init 初始化方法
func (f *Date) Init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.dateValueField
	}
	f.dateValueField.init(model, colName, fieldName, formName, uriName, index)
}

func (f Date) String() string {
	return f.dateValueField.String()
}

func (f *Date) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Date) dup() Field {
	nf := *f
	return &nf
}

func (f *Date) Name() string {
	return f.formName()
}

func (f *Date) Bind(req *http.Request, value interface{}) error {
	switch val := value.(type) {
	case string:
		if t, err := time.ParseInLocation("2006-01-02", val, time.Local); err == nil {
			f.SetDate(t)
			return nil
		}
		it, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		f.SetDate(time.Unix(it, 0))
	case int:
		f.SetDate(time.Unix(int64(val), 0))
	default:
		// return fmt.Errorf("invalid binding value type for dateValueField (%T)", value)
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid value type for Bind() (value type: %T)", value), nil)
	}
	return nil
}

func (f *Date) BindBody(b []byte, value interface{}) error {
	fmt.Println(nbcolor.Red("==========================="))
	return nil
}

func (f *Date) ValidateStruct(interface{}) error {
	return nil
}

func (f *Date) Engine() interface{} {
	return nil
}

//=========================================================================================

// Datetime 日期时间型字段
type Datetime struct {
	clauseField
	datetimeValueField
}

// Init 初始化方法
func (f *Datetime) Init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.datetimeValueField
	}
	f.datetimeValueField.init(model, colName, fieldName, formName, uriName, index)
}

func (f Datetime) String() string {
	return f.datetimeValueField.String()
}

func (f *Datetime) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Datetime) dup() Field {
	nf := *f
	return &nf
}

func (f *Datetime) Name() string {
	return f.formName()
}

func (f *Datetime) Bind(req *http.Request, value interface{}) error {
	switch val := value.(type) {
	case string:
		if t, err := time.ParseInLocation("2006-01-02 15:04:05", val, time.Local); err == nil {
			f.SetDatetime(t)
			return nil
		}
		ti, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		f.SetDatetime(time.Unix(ti, 0))
	case int:
		f.SetDatetime(time.Unix(int64(val), 0))
	default:
		// return fmt.Errorf("invalid binding type for datetimeValueField (%T)", value)
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid value type for Bind() (value type: %T)", value), nil)
	}
	return nil
}

func (f *Datetime) BindBody(b []byte, value interface{}) error {
	fmt.Println(nbcolor.Red("==========================="))
	return nil
}

func (f *Datetime) ValidateStruct(interface{}) error {
	return nil
}

func (f *Datetime) Engine() interface{} {
	return nil
}

//=============================================================================================================

// Decimal 浮点数字段
type Decimal struct {
	clauseField
	decimalValueField
}

// Init 初始化方法
func (f *Decimal) Init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.decimalValueField
	}
	f.decimalValueField.init(model, colName, fieldName, formName, uriName, index)
}

func (f Decimal) String() string {
	return f.decimalValueField.String()
}

func (f *Decimal) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Decimal) dup() Field {
	nf := *f
	return &nf
}

func (f *Decimal) Name() string {
	return f.formName()
}

func (f *Decimal) Bind(req *http.Request, value interface{}) error {
	switch val := value.(type) {
	case string:
		v, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return err
		}
		f.SetDecimal(v)
	case float32:
		f.SetDecimal(float64(val))
	case float64:
		f.SetDecimal(val)
	default:
		// return fmt.Errorf("invalid binding value type for decimalValueField (%T)", value)
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid value type for Bind() (value type: %T)", value), nil)
	}
	return nil
}

func (f *Decimal) BindBody(b []byte, value interface{}) error {
	fmt.Println(nbcolor.Red("==========================="))
	return nil
}

func (f *Decimal) ValidateStruct(interface{}) error {
	return nil
}

func (f *Decimal) Engine() interface{} {
	return nil
}

// ===================================================================================================

// Time 日期型字段
type Time struct {
	clauseField
	timeValueField
}

// Init 初始化方法
func (f *Time) Init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.timeValueField
	}
	f.timeValueField.init(model, colName, fieldName, formName, uriName, index)
}

func (f Time) String() string {
	return f.timeValueField.String()
}

func (f *Time) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Time) dup() Field {
	nf := *f
	return &nf
}

func (f *Time) Name() string {
	return f.formName()
}

func (f *Time) Bind(req *http.Request, value interface{}) error {
	switch val := value.(type) {
	case string:
		if t, err := time.ParseInLocation("15:04:05", val, time.Local); err == nil {
			f.SetTime(t)
			return nil
		}
		ti, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		f.SetTime(time.Unix(ti, 0))
	case int:
		f.SetTime(time.Unix(int64(val), 0))
	default:
		// return fmt.Errorf("invalid binding value type for timeValueField (%T)", value)
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid value type for Bind() (value type: %T)", value), nil)
	}
	return nil
}

func (f *Time) BindBody(b []byte, value interface{}) error {
	fmt.Println(nbcolor.Red("==========================="))
	return nil
}

func (f *Time) ValidateStruct(interface{}) error {
	return nil
}

func (f *Time) Engine() interface{} {
	return nil
}

// Bytes 二进制字段
type Bytes struct {
	clauseField
	byteValueField
}

// Init 初始化方法
func (f *Bytes) Init(model Model, colName, fieldName, formName, uriName string, index int) {
	f.clauseField.valueField = func() ValueField {
		return &f.byteValueField
	}
	f.byteValueField.init(model, colName, fieldName, formName, uriName, index)
}

func (f Bytes) String() string {
	return f.byteValueField.String()
}

func (f *Bytes) sqlLiteral() string {
	return f.clauseField.sqlLiteral()
}

func (f *Bytes) dup() Field {
	nf := *f
	return &nf
}

func (f *Bytes) Name() string {
	return f.formName()
}

func (f *Bytes) Bind(req *http.Request, value interface{}) error {
	switch val := value.(type) {
	case string:
		v, err := hex.DecodeString(val)
		if err != nil {
			return err
		}
		f.SetBytes(v)
	default:
		// return fmt.Errorf("invalid binding value type for bytValueField (%T)", value)
		return newErr(ErrCodeInvalidValueType, fmt.Sprintf("invalid value type for Bind() (value type: %T)", value), nil)
	}
	return nil
}

func (f *Bytes) BindBody(b []byte, value interface{}) error {
	fmt.Println(nbcolor.Red("==========================="))
	return nil
}

func (f *Bytes) ValidateStruct(interface{}) error {
	return nil
}

func (f *Bytes) Engine() interface{} {
	return nil
}
