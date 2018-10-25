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

func (f *DatetimeField) Eq(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s = %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02 15:04:05"))}
}

func (f *DatetimeField) Neq(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s <> %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02 15:04:05"))}
}

func (f *DatetimeField) Lt(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s < %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02 15:04:05"))}
}

func (f *DatetimeField) Gt(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s > %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02 15:04:05"))}
}

func (f *DatetimeField) Lte(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s <= %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02 15:04:05"))}
}

func (f *DatetimeField) Gte(val time.Time) *Where {
	return &Where{fmt.Sprintf("%s.%s.%s > %q", f.super.DB(), f.super.Tab(), f.column, val.Format("2006-01-02 15:04:05"))}
}

func (f *DatetimeField) Null() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NULL", f.super.DB(), f.super.Tab(), f.column)}
}

func (f *DatetimeField) NotNull() *Where {
	return &Where{fmt.Sprintf("%s.%s.%s IS NOT NULL", f.super.DB(), f.super.Tab(), f.column)}
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
	dstColName string
	new        func() Model
	cache      Model
	result     Model
}

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

func (oto *OneToOne) Query() error {
	m := oto.new()
	if oto.cache == nil {
		oto.cache = m
	}
	db := dbMap[m.DB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = %s", m.DB(), m.Tab(), oto.dstColName, oto.srcField.SQLVal())
	row := db.QueryRow(stmtStr)
	if err := scanRow(m, row); err != nil {
		return err
	}
	oto.result = m
	return nil
}

func (oto *OneToOne) Model() Model {
	if oto.cache == nil {
		oto.cache = oto.new()
	}
	return oto.cache
}

func (oto *OneToOne) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(oto.result, "\t", "\t")
}

func (oto *OneToOne) joinClause() string {
	return fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", oto.dstDB(), oto.dstTab(), oto.srcDB(), oto.srcTab(), oto.srcCol(), oto.dstDB(),
		oto.dstTab(), oto.dstCol())
}

type ForeignKey struct {
	srcField   Field
	dstColName string
	new        func() Model
	cache      Model
	result     Model
}

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

func (fk *ForeignKey) Query() error {
	m := fk.new()
	db := dbMap[m.DB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = %s", m.DB(), m.Tab(), fk.dstColName, fk.srcField.SQLVal())
	row := db.QueryRow(stmtStr)
	if err := scanRow(m, row); err != nil {
		return err
	}
	fk.result = m
	return nil
}

func (fk *ForeignKey) Result() Model {
	return fk.result
}

func (fk *ForeignKey) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(fk.Result(), "\t", "\t")
}

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

type ReverseForeignKey struct {
	srcField   Field
	dstColName string
	new        func() ModelList
	cache      ModelList
	result     ModelList
}

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

func (rfk *ReverseForeignKey) All() error {
	l := rfk.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = %s", l.Model().DB(), l.Model().Tab(), rfk.dstColName, rfk.srcField.SQLVal())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return err
	}
	if err = scanRows(l, rows); err != nil {
		return err
	}
	rfk.result = l
	return nil
}

func (rfk *ReverseForeignKey) AllWithFoundRows(sorter *Sorter, pager *Pager) (int, error) {
	l := rfk.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s WHERE %s = %s %s %s", l.Model().DB(), l.Model().Tab(),
		rfk.dstColName, rfk.srcField.SQLVal(), sorter.toSQL(), pager.toSQL())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return -1, err
	}
	if err = scanRows(l, rows); err != nil {
		return -1, err
	}
	rfk.result = l
	return getFoundRows(db)
}

func (rfk *ReverseForeignKey) Query(where *Where) error {
	l := rfk.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT * FROM %s.%s WHERE %s = %s AND %s", l.Model().DB(), l.Model().Tab(), rfk.dstColName, rfk.srcField.SQLVal(), where.String())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return err
	}
	if err := scanRows(l, rows); err != nil {
		return err
	}
	rfk.result = l
	return nil
}

func (rfk *ReverseForeignKey) QueryWithFoundRows(where *Where, sorter *Sorter, pager *Pager) (int, error) {
	l := rfk.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s WHERE %s = %s AND %s %s %s", l.Model().DB(), l.Model().Tab(),
		rfk.dstColName, rfk.srcField.SQLVal(), where.String(), sorter.toSQL(), pager.toSQL())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return -1, err
	}
	if err := scanRows(l, rows); err != nil {
		return -1, err
	}
	rfk.result = l
	return getFoundRows(db)
}

func (rfk *ReverseForeignKey) Result() ModelList {
	return rfk.result
}

func (rfk *ReverseForeignKey) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(rfk.Result(), "\t", "\t")
}

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

type ManyToMany struct {
	srcField      Field
	midLeftField  Field
	midRightField Field
	dstColName    string
	new           func() ModelList
	cache         ModelList
	result        ModelList
}

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

func (mtm *ManyToMany) All() error {
	l := mtm.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = %s",
		l.Model().DB(), l.Model().Tab(), mtm.srcDB(), mtm.srcTab(), mtm.midDB(), mtm.midTab(), mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.midDB(),
		mtm.midTab(), mtm.midLeftCol(), l.Model().DB(), l.Model().Tab(), mtm.midDB(), mtm.midTab(), mtm.midRightCol(), l.Model().DB(), l.Model().Tab(),
		mtm.dstColName, mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.srcField.SQLVal())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return err
	}
	if err = scanRows(l, rows); err != nil {
		return err
	}
	mtm.result = l
	return nil
}

func (mtm *ManyToMany) AllWithFoundRows(sorter *Sorter, pager *Pager) (int, error) {
	l := mtm.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = %s %s %s",
		l.Model().DB(), l.Model().Tab(), mtm.srcDB(), mtm.srcTab(), mtm.midDB(), mtm.midTab(), mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.midDB(),
		mtm.midTab(), mtm.midLeftCol(), l.Model().DB(), l.Model().Tab(), mtm.midDB(), mtm.midTab(), mtm.midRightCol(), l.Model().DB(), l.Model().Tab(),
		mtm.dstColName, mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.srcField.SQLVal(), sorter.toSQL(), pager.toSQL())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return -1, err
	}
	if err = scanRows(l, rows); err != nil {
		return -1, err
	}
	mtm.result = l
	return getFoundRows(db)
}

func (mtm *ManyToMany) Query(where *Where) error {
	l := mtm.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = %s AND %s",
		l.Model().DB(), l.Model().Tab(), mtm.srcDB(), mtm.srcTab(), mtm.midDB(), mtm.midTab(), mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.midDB(),
		mtm.midTab(), mtm.midLeftCol(), l.Model().DB(), l.Model().Tab(), mtm.midDB(), mtm.midTab(), mtm.midRightCol(), l.Model().DB(), l.Model().Tab(),
		mtm.dstColName, mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.srcField.SQLVal(), where.String())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return err
	}
	if err = scanRows(mtm.result, rows); err != nil {
		return err
	}
	mtm.result = l
	return nil
}

func (mtm *ManyToMany) QueryWithFoundRows(where *Where, sorter *Sorter, pager *Pager) (int, error) {
	l := mtm.new()
	db := dbMap[l.Model().DB()]
	stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s WHERE %s.%s.%s = %s AND %s %s %s",
		l.Model().DB(), l.Model().Tab(), mtm.srcDB(), mtm.srcTab(), mtm.midDB(), mtm.midTab(), mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.midDB(),
		mtm.midTab(), mtm.midLeftCol(), l.Model().DB(), l.Model().Tab(), mtm.midDB(), mtm.midTab(), mtm.midRightCol(), l.Model().DB(), l.Model().Tab(),
		mtm.dstColName, mtm.srcDB(), mtm.srcTab(), mtm.srcCol(), mtm.srcField.SQLVal(), where.String(), sorter.toSQL(), pager.toSQL())
	rows, err := db.Query(stmtStr)
	if err != nil {
		return -1, err
	}
	if err = scanRows(mtm.result, rows); err != nil {
		return -1, err
	}
	mtm.result = l
	return getFoundRows(db)
}

func (mtm *ManyToMany) Add(m Model) error {
	dstField := getByName(m, mtm.dstColName)
	db := dbMap[mtm.midDB()]
	stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s, %s) VALUES (%s, %s)", mtm.midDB(), mtm.midTab(), mtm.midLeftCol(), mtm.midRightCol(),
		mtm.srcField.SQLVal(), dstField.SQLVal())
	_, err := db.Exec(stmtStr)
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
			return nil
		}
		return err
	}
	mtm.result = nil
	return nil
}

func (mtm *ManyToMany) Remove(m Model) error {
	l := mtm.new()
	db := dbMap[mtm.midDB()]
	dstField := getByName(m, mtm.dstColName)
	stmtStr := fmt.Sprintf("DELETE FROM %s.%s WHERE %s = %s AND %s = %s", l.Model().DB(), l.Model().Tab(), mtm.midLeftCol(), mtm.srcField.SQLVal(),
		mtm.midRightCol(), dstField.SQLVal())
	_, err := db.Exec(stmtStr)
	if err != nil {
		return err
	}
	mtm.result = nil
	return nil
}

func (mtm *ManyToMany) Result() ModelList {
	return mtm.result
}

func (mtm *ManyToMany) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(mtm.Result(), "\t", "\t")
}

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
