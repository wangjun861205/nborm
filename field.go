package nborm

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type FieldInfo struct {
	ColName   string
	FieldName string
	Field     Field
}

type FieldInfoList []FieldInfo

type FieldList []Field

type modelStatus int

const (
	none            modelStatus = 0
	synced          modelStatus = 1
	distinct        modelStatus = 1 << 1
	forModelAgg     modelStatus = 1 << 2
	forModelWhere   modelStatus = 1 << 3
	inited          modelStatus = 1 << 4
	relInited       modelStatus = 1 << 5
	forModelUpdate  modelStatus = 1 << 6
	forModelOrder   modelStatus = 1 << 7
	forModelRef     modelStatus = 1 << 8
	forJoin         modelStatus = 1 << 9
	containValue    modelStatus = 1 << 10
	selectAll       modelStatus = 1 << 11
	forModelHaving  modelStatus = 1 << 12
	forReverseQuery modelStatus = 1 << 13
	containSubJoin  modelStatus = 1 << 14
)

type Meta struct {
	Model
	status  modelStatus
	midTabs []Model
	where   *where
	alias   string
	parent  Model
	index   int
	limit   [2]int
	aggExps []*aggExp
	having  *where
	updates updateList
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
		return fmt.Sprintf("`%s`", m.Tab())
	}
	return fmt.Sprintf("`%s`.`%s`", m.DB(), m.Tab())
}

func (m *Meta) fullTabName() string {
	if m.alias != "" {
		return fmt.Sprintf("%s AS %s", m.rawFullTabName(), m.alias)
	}
	if m.DB() == "*" {
		return m.Tab()
	}
	return fmt.Sprintf("`%s`.`%s`", m.DB(), m.Tab())
}

func (m *Meta) getAlias() string {
	return m.alias
}

func (m *Meta) setAlias() {
	m.alias = fmt.Sprintf("t%d", m.getIndex())
}

func (m *Meta) getWhere() *where {
	return m.where
}

func (m *Meta) getHaving() *where {
	return m.having
}

func (m *Meta) AndExprWhere(expr *Expr, val ...interface{}) Model {
	m.where = m.where.append(newWhere(and, expr, val...))
	m.addModelStatus(forModelWhere)
	m.addModelStatus(forModelRef)
	return m
}

func (m *Meta) OrExprWhere(expr *Expr, val ...interface{}) Model {
	m.where = m.where.append(newWhere(or, expr, val...))
	m.addModelStatus(forModelWhere)
	m.addModelStatus(forModelRef)
	return m
}

func (m *Meta) AndHaving(expr *Expr, val ...interface{}) Model {
	m.having = m.having.append(newWhere(and, expr, val...))
	m.addModelStatus(forModelHaving)
	return m
}

func (m *Meta) OrHaving(expr *Expr, val ...interface{}) Model {
	m.having = m.having.append(newWhere(and, expr, val...))
	m.addModelStatus(forModelHaving)
	return m
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

func (m *Meta) IsContainValue() bool {
	return m.status&containValue == containValue
}

func (m *Meta) IsRelInited() bool {
	return m.status&relInited == relInited
}

func (m *Meta) AddRelInited() {
	m.addModelStatus(relInited)
}

func (m *Meta) GetParent() Model {
	return m.parent
}

func (m *Meta) SetParent(parent Model) {
	m.parent = parent
}

func (m *Meta) getIndex() int {
	if m.parent != nil {
		return m.parent.getIndex()
	}
	m.index++
	return m.index
}

func (m *Meta) SetLimit(limit, offset int) {
	m.limit = [2]int{limit, offset}
}

func (m *Meta) getLimit() (limit, offset int) {
	return m.limit[0], m.limit[1]
}

func (m *Meta) SetForJoin() {
	m.GetParent().addModelStatus(containSubJoin)
	m.addModelStatus(forJoin)
	m.addModelStatus(forModelRef)
	m.addModelStatus(forReverseQuery)
}

func (m *Meta) SelectAll() {
	m.addModelStatus(selectAll)
}

func (m *Meta) StrAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newStrAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

func (m *Meta) IntAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newIntAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

func (m *Meta) DateAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newDateAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

func (m *Meta) DatetimeAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newDatetimeAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

func (m *Meta) DecAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newDecAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

func (m *Meta) getAggExps() []*aggExp {
	return m.aggExps
}

func (m *Meta) ForReverseQuery() {
	m.addModelStatus(forReverseQuery)
}

func (m *Meta) appendUpdate(u *update) {
	m.updates = append(m.updates, u)
}

func (m *Meta) getUpdateList() updateList {
	return m.updates
}

func (m *Meta) NewUpdate(field Field, expr *Expr, values ...interface{}) {
	expr.exp = fmt.Sprintf("@ = %s", expr.exp)
	expr.fields = append([]Field{field}, expr.fields...)
	m.appendUpdate(newUpdate(expr, values...))
}

type fieldStatus int

const (
	invalid     fieldStatus = 0
	valid       fieldStatus = 1
	notNull     fieldStatus = 1 << 1
	primaryKey  fieldStatus = 1 << 2
	autoInc     fieldStatus = 1 << 3
	forWhere    fieldStatus = 1 << 4
	forUpdate   fieldStatus = 1 << 5
	forSelect   fieldStatus = 1 << 6
	forSum      fieldStatus = 1 << 7
	forAscOrder fieldStatus = 1 << 8
	forDscOrder fieldStatus = 1 << 9
	forAgg      fieldStatus = 1 << 10
	forGroup    fieldStatus = 1 << 11
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

func (f *baseField) mustValid() {
	if !f.IsValid() {
		panic(fmt.Sprintf("invalid field (%s.%s.%s(%s))", f.DB(), f.Tab(), f.col, f.field))
	}
}

func (f *baseField) fullColName() string {
	if f.Model == nil {
		return f.col
	}
	if f.Model.getAlias() != "" {
		return fmt.Sprintf("%s.`%s`", f.Model.getAlias(), f.col)
	}
	return fmt.Sprintf("%s.`%s`", f.rawFullTabName(), f.col)
}

func (f *baseField) rawFullColName() string {
	return fmt.Sprintf("%s.`%s`", f.rawFullTabName(), f.col)
}

func (f *baseField) ForSelect() {
	f.addStatus(forSelect)
}

func (f *baseField) ForSum() {
	f.addStatus(forSelect)
	f.addStatus(forSum)
	f.addModelStatus(forModelAgg)
}

func (f *baseField) AscOrder() {
	f.removeStatus(forDscOrder)
	f.addStatus(forAscOrder)
	f.addModelStatus(forModelOrder)
	f.addModelStatus(forModelRef)
}

func (f *baseField) DscOrder() {
	f.removeStatus(forAscOrder)
	f.addStatus(forDscOrder)
	f.addModelStatus(forModelOrder)
	f.addModelStatus(forModelRef)
}

func (f *baseField) Distinct() {
	f.Model.addModelStatus(distinct)
	f.addStatus(forSelect)
}

func (f *baseField) String() string {
	return fmt.Sprintf("%s.%s.%s", f.DB(), f.Tab(), f.col)
}

func (f *baseField) GroupBy() {
	f.addStatus(forGroup | forSelect)
}

type String struct {
	baseField
	// clauseField
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
	f.addModelStatus(containValue)
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

func (f *String) SetString(v string) *String {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v
	return f
}

func (f *String) Set(v interface{}) Field {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v.(string)
	return f
}

func (f *String) String() string {
	return f.value
}

func (f *String) AndW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *String) OrW() Field {
	f.OrExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *String) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]string))), ", ")), f), value)
	case "IS NULL", "IS NOT NULL":
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *String) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]string))), ", ")), f), value)
	case "IS NULL", "IS NOT NULL":
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *String) SetU() Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), f.Value()))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *String) SetUpdate(value interface{}) Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), value))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *String) dup() Field {
	nf := *f
	return &nf
}

//=============================================================================================================

type Int struct {
	baseField
	// clauseField
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
	f.addModelStatus(containValue)
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

func (f *Int) SetInt(v int) *Int {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v
	return f
}

func (f *Int) Set(v interface{}) Field {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v.(int)
	return f
}

func (f *Int) Int() int {
	return f.value
}

func (f *Int) AndW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Int) OrW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Int) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]int))), ", ")), f), value)
	case "IS NULL", "IS NOT NULL":
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Int) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]int))), ", ")), f), value)
	case "IS NULL", "IS NOT NULL":
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Int) SetU() Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), f.Value()))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Int) SetUpdate(value interface{}) Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), value))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Int) dup() Field {
	nf := *f
	return &nf
}

//=======================================================================================================

type Date struct {
	baseField
	// clauseField
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
		t, err := time.Parse("2006-01-02", string(val))
		if err != nil {
			return err
		}
		f.value = t
	case string:
		t, err := time.Parse("2006-01-02", val)
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
	t, err := time.Parse("2006-01-02", strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	f.value = t
	f.addModelStatus(containValue)
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

func (f *Date) SetDate(v time.Time) *Date {
	f.setValid()
	f.unsetNull()
	f.value = v
	f.addModelStatus(containValue)
	return f
}

func (f *Date) Set(v interface{}) Field {
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
	f.addModelStatus(containValue)
	return f
}

func (f *Date) Date() time.Time {
	return f.value
}

func (f *Date) AndW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Date) OrW() Field {
	f.OrExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Date) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		vals := make([]string, 0, len(value.([]time.Time)))
		for _, t := range value.([]time.Time) {
			vals = append(vals, t.Format("2006-01-02"))
		}
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]time.Time))), ", ")), f), vals)
	case "IS NULL", "IS NOT NULL":
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Date) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		vals := make([]string, 0, len(value.([]time.Time)))
		for _, t := range value.([]time.Time) {
			vals = append(vals, t.Format("2006-01-02"))
		}
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]time.Time))), ", ")), f), vals)
	case "IS NULL", "IS NOT NULL":
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Date) SetU() Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), f.Value()))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Date) SetUpdate(value interface{}) Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), value))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Date) dup() Field {
	nf := *f
	return &nf
}

//=========================================================================================

type Datetime struct {
	baseField
	// clauseField
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
	t, err := time.Parse("2006-01-02 15:04:05", strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}
	f.value = t
	f.addModelStatus(containValue)
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

func (f *Datetime) SetDatetime(v time.Time) *Datetime {
	f.setValid()
	f.unsetNull()
	f.value = v
	f.addModelStatus(containValue)
	return f
}

func (f *Datetime) Set(v interface{}) Field {
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
	f.addModelStatus(containValue)
	return f
}

func (f *Datetime) Datetime() time.Time {
	return f.value
}

func (f *Datetime) AndW() Field {
	f.mustValid()
	f.AndExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Datetime) OrW() Field {
	f.mustValid()
	f.OrExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Datetime) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		vals := make([]string, 0, len(value.([]time.Time)))
		for _, t := range value.([]time.Time) {
			vals = append(vals, t.Format("2006-01-02 15:04:05"))
		}
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]time.Time))), ", ")), f), vals)
	case "IS NULL", "IS NOT NULL":
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Datetime) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		vals := make([]string, 0, len(value.([]time.Time)))
		for _, t := range value.([]time.Time) {
			vals = append(vals, t.Format("2006-01-02 15:04:05"))
		}
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]time.Time))), ", ")), f), vals)
	case "IS NULL", "IS NOT NULL":
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Datetime) SetU() Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), f.Value()))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Datetime) SetUpdate(value interface{}) Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), value))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Datetime) dup() Field {
	nf := *f
	return &nf
}

//=============================================================================================================

type Decimal struct {
	baseField
	// clauseField
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
	f.addModelStatus(containValue)
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

func (f *Decimal) SetDecimal(v float64) *Decimal {
	f.setValid()
	f.unsetNull()
	f.addModelStatus(containValue)
	f.value = v
	return f
}

func (f *Decimal) Set(v interface{}) Field {
	f.setValid()
	f.unsetNull()
	f.value = v.(float64)
	f.addModelStatus(containValue)
	return f
}

func (f *Decimal) Decimal() float64 {
	return f.value
}

func (f *Decimal) AndW() Field {
	f.AndExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Decimal) OrW() Field {
	f.OrExprWhere(NewExpr("@ = ?", f), f.Value())
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Decimal) AndWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]float64))), ", ")), f), value)
	case "IS NULL", "IS NOT NULL":
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.AndExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Decimal) OrWhere(op string, value interface{}) Field {
	checkOp(op)
	switch op {
	case "IN", "NOT IN":
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s (%s)", op, strings.Trim(strings.Repeat("?, ", len(value.([]float64))), ", ")), f), value)
	case "IS NULL", "IS NOT NULL":
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s", op), f))
	default:
		f.OrExprWhere(NewExpr(fmt.Sprintf("@ %s ?", op), f), value)
	}
	f.addModelStatus(forModelWhere)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Decimal) SetU() Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), f.Value()))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Decimal) SetUpdate(value interface{}) Field {
	f.appendUpdate(newUpdate(NewExpr("@ = ?", f), value))
	f.addModelStatus(forModelUpdate)
	f.addModelStatus(forModelRef)
	return f
}

func (f *Decimal) dup() Field {
	nf := *f
	return &nf
}
