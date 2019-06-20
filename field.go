package nborm

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type FieldInfo struct {
	ColName   string
	FieldName string
	Field     Field
}

type FieldInfoList []FieldInfo

type FieldList []Field

func (l FieldList) toWhereClause() (string, []interface{}) {
	whereList := make(whereList, 0, len(l)*2)
	for _, f := range l {
		whereList = append(whereList, f.whereList()...)
	}
	return whereList.toClause()
}

type modelStatus int

const (
	none          modelStatus = 0
	synced        modelStatus = 1
	distinct      modelStatus = 1 << 1
	forAgg        modelStatus = 1 << 2
	forModelWhere modelStatus = 1 << 3
)

type Meta struct {
	Model
	status   modelStatus
	alias    string
	relJoin  string
	relWhere *where
	midTabs  []Model
}

func (m *Meta) GetMidTabs() []Model {
	return m.midTabs
}

func (m *Meta) AppendMidTab(model Model) {
	m.midTabs = append(m.midTabs, model)
}

func (m *Meta) setModel(model Model) {
	m.Model = model
}

func (m *Meta) rawFullTabName() string {
	if m.DB() == "*" {
		return m.Tab()
	}
	return fmt.Sprintf("%s.%s", m.DB(), m.Tab())
}

func (m *Meta) fullTabName() string {
	if m.alias != "" {
		return fmt.Sprintf("%s AS %s", m.rawFullTabName(), m.alias)
	}
	if m.DB() == "*" {
		return m.Tab()
	}
	return fmt.Sprintf("%s.%s", m.DB(), m.Tab())
}

func (m *Meta) setRel(relJoin string, where *where) {
	m.relJoin, m.relWhere = relJoin, where
}

func (m *Meta) getAlias() string {
	return m.alias
}

func (m *Meta) setAlias(alias string) {
	m.alias = alias
}

func (m *Meta) getRelJoin() string {
	return m.relJoin
}

func (m *Meta) setRelJoin(join string) {
	m.relJoin = join
}

func (m *Meta) getRelWhere() *where {
	return m.relWhere
}

func (m *Meta) setRelWhere(where *where) {
	m.relWhere = where
}

func (m *Meta) getModelStatus() modelStatus {
	return m.status
}

func (m *Meta) addModelStatus(status modelStatus) {
	m.status |= status
}

func (m *Meta) setModelStatus(status modelStatus) {
	m.status = status
}

func (m *Meta) removeModelStatus(status modelStatus) {
	m.status &^= status
}

func (m *Meta) SelectDistinct() {
	m.addModelStatus(distinct)
}

func (m *Meta) IsSynced() bool {
	return m.status&synced == synced
}

type fieldStatus int

const (
	invalid    fieldStatus = 0
	valid      fieldStatus = 1
	notNull    fieldStatus = 1 << 1
	primaryKey fieldStatus = 1 << 2
	autoInc    fieldStatus = 1 << 3
	forWhere   fieldStatus = 1 << 4
	forUpdate  fieldStatus = 1 << 5
	forSelect  fieldStatus = 1 << 6
	forSum     fieldStatus = 1 << 7
)

type baseField struct {
	Model
	col    string
	field  string
	status fieldStatus
}

func (f *baseField) setModel(model Model) {
	f.Model = model
}

func (f *baseField) dbName() string {
	return f.DB()
}

func (f *baseField) tabName() string {
	return f.Tab()
}

func (f *baseField) colName() string {
	return f.col
}

func (f *baseField) setCol(col string) {
	f.col = col
}

func (f *baseField) fieldName() string {
	return f.field
}

func (f *baseField) setField(field string) {
	f.field = field
}

func (f *baseField) getStatus() fieldStatus {
	return f.status
}

func (f *baseField) setStatus(status fieldStatus) {
	f.status = status
}

func (f *baseField) addStatus(status fieldStatus) {
	f.status |= status
}

func (f *baseField) removeStatus(status fieldStatus) {
	f.status &^= status
}

func (f *baseField) IsValid() bool {
	return f.status&valid == valid
}

func (f *baseField) setValid() {
	f.addStatus(valid)
}

func (f *baseField) unsetValid() {
	f.removeStatus(valid)
}

func (f *baseField) IsNull() bool {
	return !(f.status&notNull == notNull)
}

func (f *baseField) SetNull() {
	f.addStatus(valid)
	f.removeStatus(notNull)
}

func (f *baseField) unsetNull() {
	f.addStatus(notNull)
}

func (f *baseField) isForWhere() bool {
	return f.status&forWhere == forWhere
}

func (f *baseField) setForWhere() {
	f.addStatus(forWhere)
}

func (f *baseField) unsetForWhere() {
	f.removeStatus(forWhere)
}

func (f *baseField) isForUpdate() bool {
	return f.status&forUpdate == forUpdate
}

func (f *baseField) setForUpdate() {
	f.addStatus(forUpdate)
}

func (f *baseField) unsetForUpdate() {
	f.removeStatus(forUpdate)
}

func (f *baseField) setPrimaryKey() {
	f.addStatus(primaryKey)
}

func (f *baseField) unsetPrimaryKey() {
	f.removeStatus(primaryKey)
}

func (f *baseField) isPrimaryKey() bool {
	return f.status&primaryKey == primaryKey
}

func (f *baseField) setAutoInc() {
	f.addStatus(autoInc)
}

func (f *baseField) unsetAutoInc() {
	f.removeStatus(autoInc)
}

func (f *baseField) isAutoInc() bool {
	return f.status&autoInc == autoInc
}

func (f *baseField) mustValid() {
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid field (%s.%s.%s(%s))", f.DB(), f.Tab(), f.col, f.field))
	}
}

func (f *baseField) fullColName() string {
	if f.Model.getAlias() != "" {
		return fmt.Sprintf("%s.%s", f.Model.getAlias(), f.col)
	}
	return fmt.Sprintf("%s.%s", f.rawFullTabName(), f.col)
}

func (f *baseField) ForSelect() {
	f.addStatus(forSelect)
}

func (f *baseField) ForSum() {
	f.addStatus(forSelect)
	f.addStatus(forSum)
	f.addModelStatus(forAgg)
}

func (f *baseField) Distinct() {
	f.Model.addModelStatus(distinct)
	f.addStatus(forSelect)
}

func (f *baseField) String() string {
	return fmt.Sprintf("%s.%s.%s", f.DB(), f.Tab(), f.col)
}

type clauseField struct {
	where  whereList
	update *updateSet
}

func (f *clauseField) andWhere(field Field, op string, value interface{}) {
	f.where = append(f.where, newWhere(and, field, op, value))
}

func (f *clauseField) orWhere(field Field, op string, value interface{}) {
	f.where = append(f.where, newWhere(or, field, op, value))
}

func (f *clauseField) setUpdate(field Field, value interface{}) {
	f.update = newUpdateSet(field, value)
}

func (f *clauseField) whereList() whereList {
	return f.where
}

func (f *clauseField) updateSet() *updateSet {
	return f.update
}

type String struct {
	baseField
	clauseField
	value string
}

func (f *String) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case string:
		f.value = val
	case []byte:
		f.value = string(val)
	default:
		return fmt.Errorf("invalid type for scan String(%T)", v)
	}
	return nil
}

func (f String) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	return json.Marshal(f.value)
}

func (f *String) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = ""
		return nil
	}
	f.unsetNull()
	return json.Unmarshal(b, &f.value)
}

func (f *String) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value
}

func (f *String) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value
}

func (f *String) SetString(v string) {
	f.setValid()
	f.unsetNull()
	f.value = v
}

func (f *String) Set(v interface{}) {
	f.setValid()
	f.unsetNull()
	f.value = v.(string)
}

func (f *String) String() string {
	return f.value
}

func (f *String) AndW() Field {
	f.mustValid()
	f.andWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *String) OrW() Field {
	f.mustValid()
	f.orWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *String) AndWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *String) OrWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *String) SetU() {
	f.setUpdate(f, f.Value())
}

func (f *String) SetUpdate(value interface{}) {
	f.setUpdate(f, value)
}

func (f *String) genAndWhere(op string, value interface{}) *where {
	return newWhere(and, f, op, value)
}

func (f *String) genOrWhere(op string, value interface{}) *where {
	return newWhere(or, f, op, value)
}

//=============================================================================================================

type Int struct {
	baseField
	clauseField
	value int
}

func (f *Int) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case int64:
		f.value = int(val)
	case []byte:
		i, err := strconv.ParseInt(string(val), 10, 64)
		if err != nil {
			return err
		}
		f.value = int(i)
	default:
		return fmt.Errorf("invalid type for scan Int(%T)", v)
	}
	return nil
}

func (f Int) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	return json.Marshal(f.value)
}

func (f *Int) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = 0
		return nil
	}
	f.unsetNull()
	return json.Unmarshal(b, &f.value)
}

func (f *Int) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value
}

func (f *Int) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value
}

func (f *Int) SetInt(v int) {
	f.setValid()
	f.unsetNull()
	f.value = v
}

func (f *Int) Set(v interface{}) {
	f.setValid()
	f.unsetNull()
	f.value = v.(int)
}

func (f *Int) Int() int {
	return f.value
}

func (f *Int) AndW() Field {
	f.mustValid()
	f.andWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Int) OrW() Field {
	f.mustValid()
	f.orWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Int) AndWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Int) OrWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Int) SetU() {
	f.setUpdate(f, f.Value())
}

func (f *Int) SetUpdate(value interface{}) {
	f.setUpdate(f, value)
}

func (f *Int) genAndWhere(op string, value interface{}) *where {
	return newWhere(and, f, op, value)
}

func (f *Int) genOrWhere(op string, value interface{}) *where {
	return newWhere(or, f, op, value)
}

//=======================================================================================================

type Date struct {
	baseField
	clauseField
	value time.Time
}

func (f *Date) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05", string(val))
		if err != nil {
			return err
		}
		f.value = t
	case string:
		t, err := time.Parse("2006-01-02 15:04:05", val)
		if err != nil {
			return err
		}
		f.value = t
	case time.Time:
		f.value = val
	default:
		return fmt.Errorf("invalid type for scan Date(%T)", v)
	}
	return nil
}

func (f *Date) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	s := f.value.Format("2006-01-02")
	return json.Marshal(s)
}

func (f *Date) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = time.Time{}
		return nil
	}
	f.unsetNull()
	t, err := time.Parse("2006-01-02", string(b))
	if err != nil {
		return err
	}
	f.value = t
	return nil
}

func (f *Date) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value.Format("2006-01-02")
}

func (f *Date) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value.Format("2006-01-02")
}

func (f *Date) SetDate(v time.Time) {
	f.setValid()
	f.unsetNull()
	f.value = v
}

func (f *Date) Set(v interface{}) {
	switch val := v.(type) {
	case string:
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			panic(err)
		}
		f.value = t
	case time.Time:
		f.value = val
	default:
		panic(fmt.Errorf("invalid value for Date.Set(): want string or time.Time got (%T)", v))
	}
	f.setValid()
	f.unsetNull()
}

func (f *Date) Date() time.Time {
	return f.value
}

func (f *Date) AndW() Field {
	f.mustValid()
	f.andWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Date) OrW() Field {
	f.mustValid()
	f.orWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Date) AndWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Date) OrWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Date) SetU() {
	f.setUpdate(f, f.Value())
}

func (f *Date) SetUpdate(value interface{}) {
	f.setUpdate(f, value)
}

func (f *Date) genAndWhere(op string, value interface{}) *where {
	return newWhere(and, f, op, value)
}

func (f *Date) genOrWhere(op string, value interface{}) *where {
	return newWhere(or, f, op, value)
}

//=========================================================================================

type Datetime struct {
	baseField
	clauseField
	value time.Time
}

func (f *Datetime) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05", string(val))
		if err != nil {
			return err
		}
		f.value = t
	case string:
		t, err := time.Parse("2006-01-02 15:04:05", val)
		if err != nil {
			return err
		}
		f.value = t
	case time.Time:
		f.value = val
	default:
		return fmt.Errorf("invalid type for scan Date(%T)", v)
	}
	return nil
}

func (f *Datetime) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	s := f.value.Format("2006-01-02 15:04:05")
	return json.Marshal(s)
}

func (f *Datetime) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = time.Time{}
		return nil
	}
	f.unsetNull()
	t, err := time.Parse("2006-01-02 15:04:05", string(b))
	if err != nil {
		return err
	}
	f.value = t
	return nil
}

func (f *Datetime) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value.Format("2006-01-02 15:04:05")
}

func (f *Datetime) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value.Format("2006-01-02 15:04:05")
}

func (f *Datetime) SetDatetime(v time.Time) {
	f.setValid()
	f.unsetNull()
	f.value = v
}

func (f *Datetime) Set(v interface{}) {
	switch val := v.(type) {
	case string:
		t, err := time.Parse("2006-01-02", val)
		if err != nil {
			panic(err)
		}
		f.value = t
	case time.Time:
		f.value = val
	default:
		panic(fmt.Errorf("invalid value for Date.Set(): want string or time.Time got (%T)", v))
	}
	f.setValid()
	f.unsetNull()
}

func (f *Datetime) Datetime() time.Time {
	return f.value
}

func (f *Datetime) AndW() Field {
	f.mustValid()
	f.andWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Datetime) OrW() Field {
	f.mustValid()
	f.orWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Datetime) AndWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Datetime) OrWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Datetime) SetU() {
	f.setUpdate(f, f.Value())
}

func (f *Datetime) SetUpdate(value interface{}) {
	f.setUpdate(f, value)
}

func (f *Datetime) genAndWhere(op string, value interface{}) *where {
	return newWhere(and, f, op, value)
}

func (f *Datetime) genOrWhere(op string, value interface{}) *where {
	return newWhere(or, f, op, value)
}

//=============================================================================================================

type Decimal struct {
	baseField
	clauseField
	value   float64
	exprVal *Expr
}

func (f *Decimal) Scan(v interface{}) error {
	f.setValid()
	if v == nil {
		f.SetNull()
		return nil
	}
	f.unsetNull()
	switch val := v.(type) {
	case float64:
		f.value = val
	case float32:
		f.value = float64(val)
	default:
		return fmt.Errorf("invalid type for scan Decimal(%T)", v)
	}
	return nil
}

func (f *Decimal) MarshalJSON() ([]byte, error) {
	if !f.IsValid() || f.IsNull() {
		return json.Marshal(nil)
	}
	return json.Marshal(f.value)
}

func (f *Decimal) UnmarshalJSON(b []byte) error {
	f.addStatus(valid)
	if string(b) == "null" {
		f.SetNull()
		f.value = 0.0
		return nil
	}
	f.unsetNull()
	return json.Unmarshal(b, &f.value)

}

func (f *Decimal) Value() interface{} {
	f.mustValid()
	if f.IsNull() {
		return nil
	}
	return f.value
}

func (f *Decimal) JSONValue() interface{} {
	if !f.IsValid() || f.IsNull() {
		return nil
	}
	return f.value
}

func (f *Decimal) SetDecimal(v float64) {
	f.setValid()
	f.unsetNull()
	f.value = v
}

func (f *Decimal) Set(v interface{}) {
	f.setValid()
	f.unsetNull()
	f.value = v.(float64)
}

func (f *Decimal) Decimal() float64 {
	return f.value
}

func (f *Decimal) AndW() Field {
	f.mustValid()
	f.andWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Decimal) OrW() Field {
	f.mustValid()
	f.orWhere(f, "=", f.Value())
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Decimal) AndWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Decimal) OrWhere(op string, value interface{}) Field {
	f.andWhere(f, op, value)
	f.addStatus(forWhere)
	f.addModelStatus(forModelWhere)
	return f
}

func (f *Decimal) SetU() {
	f.setUpdate(f, f.Value())
}

func (f *Decimal) SetUpdate(value interface{}) {
	f.setUpdate(f, value)
}

func (f *Decimal) genAndWhere(op string, value interface{}) *where {
	return newWhere(and, f, op, value)
}

func (f *Decimal) genOrWhere(op string, value interface{}) *where {
	return newWhere(or, f, op, value)
}
