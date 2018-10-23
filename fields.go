package nborm

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

type UpdateValue struct {
	column string
	sqlVal string
	val    interface{}
	null   bool
}

func (v *UpdateValue) String() string {
	return fmt.Sprintf("%s = %s", v.column, v.sqlVal)
}

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

func NewStringField(model Model, column string, pk bool, inc bool, uni bool) *StringField {
	return &StringField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *StringField) Super() Model {
	return f.super
}

func (f *StringField) Column() string {
	return f.column
}

func (f *StringField) IsPk() bool {
	return f.pk
}

func (f *StringField) IsInc() bool {
	return f.inc
}

func (f *StringField) IsValid() bool {
	return f.valid
}

func (f *StringField) IsNull() bool {
	return f.null
}

func (f *StringField) Set(val string, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
}

func (f *StringField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(string)
}

func (f *StringField) Get() (val string, valid bool, null bool) {
	return f.val, f.valid, f.null
}

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

func (f *StringField) SQLVal() string {
	return fmt.Sprintf("%q", f.val)
}

func (f *StringField) Invalidate() {
	f.valid = false
}

func (f *StringField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(fmt.Sprintf("%q", f.val)), nil
}

func (f *StringField) InsertValuePair() [2]string {
	return [2]string{f.column, f.SQLVal()}
}

func (f *StringField) UpdateValue() *UpdateValue {
	return &UpdateValue{f.column, f.SQLVal(), f.val, f.null}
}

func (f *StringField) BulkUpdateValue(val string, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, fmt.Sprintf("%q", val), val, isNull}
}

func (f *StringField) Where() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %s", f.super.DB(), f.super.Tab(), f.column, f.SQLVal())}
}

func (f *StringField) Eq(val string) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %q", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *StringField) Neq(val string) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s <> %q", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *StringField) Contains(val string) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s LIKE \"%%%s%%\"", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *StringField) Null() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *StringField) NotNull() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NOT NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (s *StringField) IsUni() bool {
	return s.uni
}

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

func (f *StringField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

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

func NewIntField(model Model, column string, pk bool, inc bool, uni bool) *IntField {
	return &IntField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *IntField) Super() Model {
	return f.super
}

func (f *IntField) Column() string {
	return f.column
}

func (f *IntField) IsPk() bool {
	return f.pk
}

func (f *IntField) IsInc() bool {
	return f.inc
}

func (f *IntField) IsValid() bool {
	return f.valid
}

func (f *IntField) IsNull() bool {
	return f.null
}

func (f *IntField) Set(val int64, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
}

func (f *IntField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(int64)
}

func (f *IntField) Get() (val int64, valid bool, null bool) {
	return f.val, f.valid, f.null
}

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

func (f *IntField) SQLVal() string {
	return fmt.Sprintf("%d", f.val)
}

func (f *IntField) Invalidate() {
	f.valid = false
}

func (f *IntField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(strconv.FormatInt(f.val, 10)), nil
}

func (f *IntField) InsertValuePair() [2]string {
	return [2]string{f.column, f.SQLVal()}
}

func (f *IntField) UpdateValue() *UpdateValue {
	return &UpdateValue{f.column, f.SQLVal(), f.val, f.null}
}

func (f *IntField) BulkUpdateValue(val int64, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, fmt.Sprintf("%d", val), val, isNull}
}

func (f *IntField) Where() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %s", f.super.DB(), f.super.Tab(), f.column, f.SQLVal())}
}

func (f *IntField) Eq(val int64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %d", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *IntField) Neq(val int64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s <> %d", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *IntField) Lt(val int64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s < %d", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *IntField) Gt(val int64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s > %d", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *IntField) Lte(val int64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s <= %d", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *IntField) Gte(val int64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s >= %d", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *IntField) Null() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *IntField) NotNull() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NOT NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *IntField) IsUni() bool {
	return f.uni
}

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

func (f *IntField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

func (f *IntField) SumFunc() func(modelList) float64 {
	return func(l modelList) float64 {
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

func (f *IntField) AvgFunc(ignoreNull bool) func(modelList) float64 {
	if ignoreNull {
		return func(l modelList) float64 {
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
	return func(l modelList) float64 {
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

func NewFloatField(model Model, column string, pk bool, inc bool, uni bool) *FloatField {
	return &FloatField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *FloatField) Super() Model {
	return f.super
}

func (f *FloatField) Column() string {
	return f.column
}

func (f *FloatField) IsPk() bool {
	return f.pk
}

func (f *FloatField) IsInc() bool {
	return f.inc
}

func (f *FloatField) IsValid() bool {
	return f.valid
}

func (f *FloatField) IsNull() bool {
	return f.null
}

func (f *FloatField) Set(val float64, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
}

func (f *FloatField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(float64)
}

func (f *FloatField) Get() (val float64, valid bool, null bool) {
	return f.val, f.valid, f.null
}

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

func (f *FloatField) SQLVal() string {
	return fmt.Sprintf("%f", f.val)
}

func (f *FloatField) Invalidate() {
	f.valid = false
}

func (f *FloatField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(strconv.FormatFloat(f.val, 'f', -1, 64)), nil
}

func (f *FloatField) InsertValuePair() [2]string {
	return [2]string{f.column, f.SQLVal()}
}

func (f *FloatField) UpdateValue() *UpdateValue {
	return &UpdateValue{f.column, f.SQLVal(), f.val, f.null}
}

func (f *FloatField) BulkUpdateValue(val float64, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, fmt.Sprintf("%f", val), val, isNull}
}

func (f *FloatField) Where() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %s", f.super.DB(), f.super.Tab(), f.column, f.SQLVal())}
}

func (f *FloatField) Eq(val float64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %f", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *FloatField) Neq(val float64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s <> %f", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *FloatField) Lt(val float64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s < %f", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *FloatField) Gt(val float64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s > %f", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *FloatField) Lte(val float64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s <= %f", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *FloatField) Gte(val float64) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s >= %f", f.super.DB(), f.super.Tab(), f.column, val)}
}

func (f *FloatField) Null() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *FloatField) NotNull() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NOT NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *FloatField) IsUni() bool {
	return f.uni
}

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

func (f *FloatField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

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

func NewBoolField(model Model, column string, pk bool, inc bool, uni bool) *BoolField {
	return &BoolField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *BoolField) Super() Model {
	return f.super
}

func (f *BoolField) Column() string {
	return f.column
}

func (f *BoolField) IsPk() bool {
	return f.pk
}

func (f *BoolField) IsInc() bool {
	return f.inc
}

func (f *BoolField) IsValid() bool {
	return f.valid
}

func (f *BoolField) IsNull() bool {
	return f.null
}

func (f *BoolField) Set(val bool, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
}

func (f *BoolField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(bool)
}

func (f *BoolField) Get() (val bool, valid bool, null bool) {
	return f.val, f.valid, f.null
}

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
	default:
		return fmt.Errorf("not supported value type for BoolField: %T", val)
	}
	return nil
}

func (f *BoolField) SQLVal() string {
	return fmt.Sprintf("%t", f.val)
}

func (f *BoolField) Invalidate() {
	f.valid = false
}

func (f *BoolField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(strconv.FormatBool(f.val)), nil
}

func (f *BoolField) InsertValuePair() [2]string {
	return [2]string{f.column, f.SQLVal()}
}

func (f *BoolField) UpdateValue() *UpdateValue {
	return &UpdateValue{f.column, f.SQLVal(), f.val, f.null}
}

func (f *BoolField) BulkUpdateValue(val bool, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, fmt.Sprintf("%t", val), val, isNull}
}

func (f *BoolField) Where() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %s", f.super.DB(), f.super.Tab(), f.column, f.SQLVal())}
}

func (f *BoolField) Eq(val bool) *Where {
	if val {
		return &Where{fmt.Sprintf("%s.%s.%s = 1", f.super.DB(), f.super.Tab(), f.column)}
	} else {
		return &Where{fmt.Sprintf("%s.%s.%s = 0", f.super.DB(), f.super.Tab(), f.column)}
	}
}

func (f *BoolField) Neq(val bool) *Where {
	if val {
		return &Where{fmt.Sprintf("%s.%s.%s <> 1", f.super.DB(), f.super.Tab(), f.column)}
	} else {
		return &Where{fmt.Sprintf("%s.%s.%s <> 0", f.super.DB(), f.super.Tab(), f.column)}
	}
}

func (f *BoolField) Null() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *BoolField) NotNull() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NOT NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *BoolField) IsUni() bool {
	return f.uni
}

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

func (f *BoolField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

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

func NewDateField(model Model, column string, pk bool, inc bool, uni bool) *DateField {
	return &DateField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *DateField) Super() Model {
	return f.super
}

func (f *DateField) Column() string {
	return f.column
}

func (f *DateField) IsPk() bool {
	return f.pk
}

func (f *DateField) IsInc() bool {
	return f.inc
}

func (f *DateField) IsValid() bool {
	return f.valid
}

func (f *DateField) IsNull() bool {
	return f.null
}

func (f *DateField) Set(val time.Time, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
}

func (f *DateField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(time.Time)
}

func (f *DateField) Get() (val time.Time, valid bool, null bool) {
	return f.val, f.valid, f.null
}

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

func (f *DateField) SQLVal() string {
	return fmt.Sprintf("%q", f.val.Format("2006-01-02"))
}

func (f *DateField) Invalidate() {
	f.valid = false
}

func (f *DateField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(fmt.Sprintf("%q", f.val.Format("2006-01-02"))), nil
}

func (f *DateField) InsertValuePair() [2]string {
	return [2]string{f.column, f.SQLVal()}
}

func (f *DateField) UpdateValue() *UpdateValue {
	return &UpdateValue{f.column, f.SQLVal(), f.val, f.null}
}

func (f *DateField) BulkUpdateValue(val time.Time, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val.Format("2006-01-02"), val, isNull}
}

func (f *DateField) Where() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %s", f.super.DB(), f.super.Tab(), f.column, f.SQLVal())}
}

func (f *DateField) Eq(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02"))}
}

func (f *DateField) Neq(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s <> %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02"))}
}

func (f *DateField) Lt(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s < %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02"))}
}

func (f *DateField) Gt(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s > %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02"))}
}

func (f *DateField) Lte(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s <= %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02"))}
}

func (f *DateField) Gte(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s >= %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02"))}
}

func (f *DateField) Null() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *DateField) NotNull() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NOT NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *DateField) IsUni() bool {
	return f.uni
}

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

func (f *DateField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

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

func NewDatetimeField(model Model, column string, pk bool, inc bool, uni bool) *DatetimeField {
	return &DatetimeField{super: model, column: column, pk: pk, inc: inc, uni: uni}
}

func (f *DatetimeField) Super() Model {
	return f.super
}

func (f *DatetimeField) Column() string {
	return f.column
}

func (f *DatetimeField) IsPk() bool {
	return f.pk
}

func (f *DatetimeField) IsInc() bool {
	return f.inc
}

func (f *DatetimeField) IsValid() bool {
	return f.valid
}

func (f *DatetimeField) IsNull() bool {
	return f.null
}

func (f *DatetimeField) Set(val time.Time, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
}

func (f *DatetimeField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.(time.Time)
}

func (f *DatetimeField) Get() (val time.Time, valid bool, null bool) {
	return f.val, f.valid, f.null
}

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

func (f *DatetimeField) SQLVal() string {
	return fmt.Sprintf("%q", f.val.Format("2006-01-02 15:04:05"))
}

func (f *DatetimeField) Invalidate() {
	f.valid = false
}

func (f *DatetimeField) MarshalJSON() ([]byte, error) {
	if !f.valid {
		return []byte("\"invalid\""), nil
	}
	if f.null {
		return []byte("\"NULL\""), nil
	}
	return []byte(fmt.Sprintf("%q", f.val.Format("2006-01-02 15:04:05"))), nil
}

func (f *DatetimeField) InsertValuePair() [2]string {
	return [2]string{f.column, f.SQLVal()}
}

func (f *DatetimeField) UpdateValue() *UpdateValue {
	return &UpdateValue{f.column, f.SQLVal(), f.val, f.null}
}

func (f *DatetimeField) BulkUpdateValue(val time.Time, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, val.Format("2006-01-02 15:04:05"), val, isNull}
}

func (f *DatetimeField) Where() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %s", f.super.DB(), f.super.Tab(), f.column, f.SQLVal())}
}

func (c *DatetimeField) Eq(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s = %q", c.column, val.Format("2006-01-02 15:04:05"))}
}

func (c *DatetimeField) Neq(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s <> %q", c.column, val.Format("2006-01-02 15:04:05"))}
}

func (c *DatetimeField) Lt(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s < %q", c.column, val.Format("2006-01-02 15:04:05"))}
}

func (c *DatetimeField) Gt(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s > %q", c.column, val.Format("2006-01-02 15:04:05"))}
}

func (c *DatetimeField) Lte(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s <= %q", c.column, val.Format("2006-01-02 15:04:05"))}
}

func (c *DatetimeField) Gte(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s > %q", c.column, val.Format("2006-01-02 15:04:05"))}
}

func (c *DatetimeField) Null() *Where {
	return &Where{fmt.Sprintf("%s IS NULL", c.column)}
}

func (f *DatetimeField) NotNull() *Where {
	return &Where{fmt.Sprintf("%s IS NOT NULL", f.column)}
}

func (f *DatetimeField) IsUni() bool {
	return f.uni
}

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

func (f *DatetimeField) SortOrder(reverse bool) string {
	var o string
	if reverse {
		o = "DESC"
	} else {
		o = "ASC"
	}
	return fmt.Sprintf("%s.%s.%s %s", f.super.DB(), f.super.Tab(), f.column, o)
}

type BinaryField struct {
	column string
	val    []byte
	valid  bool
	null   bool
	pk     bool
	inc    bool
	uni    bool
}

func NewBinaryField(column string, pk bool, inc bool, uni bool) *BinaryField {
	return &BinaryField{column: column, pk: pk, inc: inc, uni: uni}
}

func (f *BinaryField) Get() (val []byte, isValid, isNull bool) {
	return f.val, f.valid, f.null
}

func (f *BinaryField) SQLVal() string {
	if f.null {
		return "NULL"
	}
	return fmt.Sprintf("X'%x'", f.val)
}

func (f *BinaryField) IsValid() bool {
	return f.valid
}

func (f *BinaryField) IsNull() bool {
	return f.null
}

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

func (f *BinaryField) Column() string {
	return f.column
}

func (f *BinaryField) IsPk() bool {
	return f.pk
}

func (f *BinaryField) IsInc() bool {
	return f.inc
}

func (f *BinaryField) IsUni() bool {
	return f.uni
}

func (f *BinaryField) InsertValuePair() [2]string {
	return [2]string{f.column, f.SQLVal()}
}

func (f *BinaryField) UpdateValue() *UpdateValue {
	return &UpdateValue{f.column, f.SQLVal(), f.val, f.null}
}

func (f *BinaryField) BulkUpdateValue(val []byte, isNull bool) *UpdateValue {
	return &UpdateValue{f.column, fmt.Sprintf("X'%x'", val), val, isNull}
}

func (f *BinaryField) Where() *Where {
	return &Where{fmt.Sprintf("%s = %s", f.column, f.SQLVal())}
}

func (f *BinaryField) LessFunc() func(Model, Model) int {
	return func(iMod, jMod Model) int {
		return 0
	}
}

func (f *BinaryField) Set(val []byte, isNull bool) {
	f.valid, f.null, f.val = true, isNull, val
}

func (f *BinaryField) SetByUpdateValue(val *UpdateValue) {
	f.valid, f.null, f.val = true, val.null, val.val.([]byte)
}

func (f *BinaryField) Invalidate() {
	f.valid = false
}

func (f *BinaryField) Eq(val []byte) *Where {
	return &Where{fmt.Sprintf("%s = %s", f.column, fmt.Sprintf("X'%x'", val))}
}

func (f *BinaryField) Neq(val []byte) *Where {
	return &Where{fmt.Sprintf("%s <> %s", f.column, fmt.Sprintf("X'%x'", val))}
}

func (f *BinaryField) Null() *Where {
	return &Where{fmt.Sprintf("%s IS NULL", f.column)}
}

func (f *BinaryField) NotNull() *Where {
	return &Where{fmt.Sprintf("%s IS NOT NULL", f.column)}
}

type OneToOne struct {
	srcField   Field
	dstField   Field
	result     Model
	isResultOk bool
}

func NewOneToOne(srcField, dstField Field, model Model) *OneToOne {
	return &OneToOne{srcField, dstField, model, false}
}

func (oto *OneToOne) DstDB() string {
	return oto.dstField.Super().DB()
}

func (oto *OneToOne) DstTab() string {
	return oto.dstField.Super().Tab()
}

func (oto *OneToOne) DstCol() string {
	return oto.dstField.Column()
}

func (oto *OneToOne) SrcCol() string {
	return oto.srcField.Column()
}

func (oto *OneToOne) Query() error {
	db := dbMap[oto.DstDB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = %s", oto.DstDB(), oto.DstTab(), oto.DstCol(), oto.srcField.SQLVal())
	row := db.QueryRow(stmtStr)
	err := scanRow(oto.result, row)
	if err != nil {
		return err
	}
	oto.isResultOk = true
	return nil
}

func (oto *OneToOne) Result() Model {
	if oto.isResultOk {
		return oto.result
	}
	return nil
}

func (oto *OneToOne) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(oto.Result(), "\t", "\t")
}

type ForeignKey struct {
	srcField   Field
	dstField   Field
	result     Model
	isResultOk bool
}

func NewForeignKey(srcField, dstField Field, model Model) *ForeignKey {
	return &ForeignKey{srcField, dstField, model, false}
}

func (fk *ForeignKey) DstDB() string {
	return fk.dstField.Super().DB()
}

func (fk *ForeignKey) DstTab() string {
	return fk.dstField.Super().Tab()
}

func (fk *ForeignKey) DstCol() string {
	return fk.dstField.Column()
}

func (fk *ForeignKey) SrcCol() string {
	return fk.srcField.Column()
}

func (fk *ForeignKey) Query() error {
	db := dbMap[fk.DstDB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = %s", fk.DstDB(), fk.DstTab(), fk.DstCol(), fk.srcField.SQLVal())
	row := db.QueryRow(stmtStr)
	err := scanRow(fk.result, row)
	if err != nil {
		return err
	}
	fk.isResultOk = true
	return nil
}

func (fk *ForeignKey) Result() Model {
	if fk.isResultOk {
		return fk.result
	}
	return nil
}

func (fk *ForeignKey) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(fk.Result(), "\t", "\t")
}

type ReverseForeignKey struct {
	srcField   Field
	dstField   Field
	result     modelList
	isResultOk bool
}

func NewReverseForeignKey(srcField, dstField Field, modelList modelList) *ReverseForeignKey {
	return &ReverseForeignKey{srcField, dstField, modelList, false}
}

func (rfk *ReverseForeignKey) DstDB() string {
	return rfk.dstField.Super().DB()
}

func (rfk *ReverseForeignKey) DstTab() string {
	return rfk.dstField.Super().DB()
}

func (rfk *ReverseForeignKey) DstCol() string {
	return rfk.dstField.Column()
}

func (rfk *ReverseForeignKey) SrcCol() string {
	return rfk.srcField.Column()
}

func (rfk *ReverseForeignKey) All() error {
	db := dbMap[rfk.DstDB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = %s", rfk.DstDB(), rfk.DstTab(), rfk.DstCol(), rfk.srcField.SQLVal())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return err
	}
	err = scanRows(rfk.result, rows)
	if err != nil {
		return err
	}
	rfk.isResultOk = true
	return nil
}

func (rfk *ReverseForeignKey) Query(where *Where) error {
	db := dbMap[rfk.DstDB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = %s AND %s", rfk.DstDB(), rfk.DstTab(), rfk.DstCol(), rfk.srcField.SQLVal(), where.String())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return err
	}
	return scanRows(rfk.result, rows)
}

func (rfk *ReverseForeignKey) Result() modelList {
	if rfk.isResultOk {
		return rfk.result
	}
	return rfk.result
}

func (rfk *ReverseForeignKey) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(rfk.Result(), "\t", "\t")
}

type ManyToMany struct {
	srcField      Field
	midLeftField  Field
	midRightField Field
	dstField      Field
	result        modelList
	isResultOk    bool
}

func NewManyToMany(srcField, midLeftField, midRightField, dstField Field, modelList modelList) *ManyToMany {
	if midLeftField.Super() != midRightField.Super() {
		panic("nborm.NewManyToMany() error: require the same middle tab")
	}
	return &ManyToMany{srcField, midLeftField, midRightField, dstField, modelList, false}
}

func (mtm *ManyToMany) DstDB() string {
	return mtm.dstField.Super().DB()
}

func (mtm *ManyToMany) DstTab() string {
	return mtm.dstField.Super().DB()
}

func (mtm *ManyToMany) DstCol() string {
	return mtm.dstField.Column()
}

func (mtm *ManyToMany) SrcDB() string {
	return mtm.srcField.Super().DB()
}

func (mtm *ManyToMany) SrcTab() string {
	return mtm.srcField.Super().Tab()
}

func (mtm *ManyToMany) SrcCol() string {
	return mtm.srcField.Column()
}

func (mtm *ManyToMany) MidDB() string {
	return mtm.midLeftField.Super().DB()
}

func (mtm *ManyToMany) MidTab() string {
	return mtm.midLeftField.Super().Tab()
}

func (mtm *ManyToMany) MidLeftCol() string {
	return mtm.midLeftField.Column()
}

func (mtm *ManyToMany) MidRightCol() string {
	return mtm.midRightField.Column()
}

func (mtm *ManyToMany) All() error {
	db := dbMap[mtm.DstDB()]
	stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = %s",
		mtm.DstDB(), mtm.DstTab(), mtm.SrcDB(), mtm.SrcTab(), mtm.MidDB(), mtm.MidTab(), mtm.SrcDB(), mtm.SrcTab(), mtm.SrcCol(), mtm.MidDB(),
		mtm.MidTab(), mtm.MidLeftCol(), mtm.DstDB(), mtm.DstTab(), mtm.MidDB(), mtm.MidTab(), mtm.MidRightCol(), mtm.DstDB(), mtm.DstTab(),
		mtm.DstCol(), mtm.SrcDB(), mtm.SrcTab(), mtm.SrcCol(), mtm.srcField.SQLVal())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return err
	}
	err = scanRows(mtm.result, rows)
	if err != nil {
		return err
	}
	mtm.isResultOk = true
	return nil
}

func (mtm *ManyToMany) Query(where *Where) error {
	db := dbMap[mtm.DstDB()]
	stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = %s AND %s",
		mtm.DstDB(), mtm.DstTab(), mtm.SrcDB(), mtm.SrcTab(), mtm.MidDB(), mtm.MidTab(), mtm.SrcDB(), mtm.SrcTab(), mtm.SrcCol(), mtm.MidDB(),
		mtm.MidTab(), mtm.MidLeftCol(), mtm.DstDB(), mtm.DstTab(), mtm.MidDB(), mtm.MidTab(), mtm.MidRightCol(), mtm.DstDB(), mtm.DstTab(),
		mtm.DstCol(), mtm.SrcDB(), mtm.SrcTab(), mtm.SrcCol(), mtm.srcField.SQLVal(), where.String())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return err
	}
	err = scanRows(mtm.result, rows)
	if err != nil {
		return err
	}
	mtm.isResultOk = true
	return nil
}

func (mtm *ManyToMany) Add(m Model) error {
	dstField := getByName(m, mtm.DstCol())
	db := dbMap[mtm.MidDB()]
	stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (%s, %s)", mtm.MidDB(), mtm.MidTab(), mtm.MidLeftCol(), mtm.MidRightCol(),
		mtm.srcField.SQLVal(), dstField.SQLVal())
	_, err := db.Exec(stmtStr)
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
			return nil
		}
		return err
	}
	mtm.isResultOk = false
	return nil
}

func (mtm *ManyToMany) Remove(m Model) error {
	db := dbMap[mtm.MidDB()]
	dstField := getByName(m, mtm.DstCol())
	stmtStr := fmt.Sprintf("DELETE FROM %s.%s WHERE %s = %s AND %s = %s", mtm.DstDB(), mtm.DstTab(), mtm.MidLeftCol(), mtm.srcField.SQLVal(),
		mtm.MidRightCol(), dstField.SQLVal())
	_, err := db.Exec(stmtStr)
	if err != nil {
		return err
	}
	mtm.isResultOk = false
	return nil
}

func (mtm *ManyToMany) Result() modelList {
	if mtm.isResultOk {
		return mtm.result
	}
	return nil
}

func (mtm *ManyToMany) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(mtm.Result(), "\t", "\t")
}
