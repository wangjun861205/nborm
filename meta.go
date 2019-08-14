package nborm

import (
	"encoding/json"
	"fmt"
)

type modelStatus int

const (
	none             modelStatus = 0
	synced           modelStatus = 1
	distinct         modelStatus = 1 << 1
	forModelAgg      modelStatus = 1 << 2
	inited           modelStatus = 1 << 4
	relInited        modelStatus = 1 << 5
	forBackQuery     modelStatus = 1 << 6
	forUpdate        modelStatus = 1 << 7
	forModelOrder    modelStatus = 1 << 8
	forModelRef      modelStatus = 1 << 9
	forJoin          modelStatus = 1 << 10
	forLeftJoin      modelStatus = 1 << 11
	forRightJoin     modelStatus = 1 << 12
	containValue     modelStatus = 1 << 13
	forModelHaving   modelStatus = 1 << 15
	containSubUpdate modelStatus = 1 << 22
	containSubOrder  modelStatus = 1 << 25
	forDelete        modelStatus = 1 << 26
)

type modelBaseInfo struct {
	Model
	parent  Model
	status  modelStatus
	alias   string
	index   int
	rels    RelationInfoList
	conList ModelList
}

func (m *modelBaseInfo) rawFullTabName() string {
	if m.DB() == "*" {
		return fmt.Sprintf("`%s`", m.Tab())
	}
	return fmt.Sprintf("`%s`.`%s`", m.DB(), m.Tab())
}

func (m *modelBaseInfo) fullTabName() string {
	if m.alias != "" {
		return fmt.Sprintf("%s AS %s", m.rawFullTabName(), m.alias)
	}
	if m.DB() == "*" {
		return m.Tab()
	}
	return fmt.Sprintf("`%s`.`%s`", m.DB(), m.Tab())
}

func (m *modelBaseInfo) getAlias() string {
	return m.alias
}

func (m *modelBaseInfo) genAlias() {
	index := m.genIndex()
	m.index = index
	m.alias = fmt.Sprintf("t%d", index)
}

func (m *modelBaseInfo) getModelStatus() modelStatus {
	return m.status
}

func (m *modelBaseInfo) addModelStatus(status modelStatus) {
	m.status |= status
}

func (m *modelBaseInfo) setModelStatus(status modelStatus) {
	m.status = status
}

func (m *modelBaseInfo) removeModelStatus(status modelStatus) {
	m.status &^= status
}

func (m *modelBaseInfo) checkStatus(status modelStatus) bool {
	return m.status&status > 0
}

// SelectDistinct 设定去重标志位
func (m *modelBaseInfo) SelectDistinct() {
	m.addModelStatus(distinct)
}

// IsSynced 检查是否为synced
func (m *modelBaseInfo) IsSynced() bool {
	return m.status&synced == synced
}

// IsContainValue 是否包含Value(Scan或直接设置)
func (m *modelBaseInfo) IsContainValue() bool {
	return m.status&containValue == containValue
}

// IsRelInited 子关系是否已初始化
func (m *modelBaseInfo) IsRelInited() bool {
	return m.status&relInited == relInited
}

// AddRelInited 添加子关系初始化标志位
func (m *modelBaseInfo) AddRelInited() {
	m.addModelStatus(relInited)
}

// GetParent 获取当前Model的Father Model
func (m *modelBaseInfo) getParent() Model {
	return m.parent
}

// setParent 设置当前Model的Father Model
func (m *modelBaseInfo) setParent(parent Model) {
	m.parent = parent
}

func (m *modelBaseInfo) setIndex(index int) {
	m.index = index
}

func (m *modelBaseInfo) getIndex() int {
	return m.index
}

func (m *modelBaseInfo) genIndex() int {
	if m.parent != nil {
		return m.parent.genIndex()
	}
	m.index++
	return m.index
}

// SetForJoin 设置Join查询标志位(所有Father Model的containSubJoin标志位均会被置为1)
func (m *modelBaseInfo) SetForJoin() {
	m.addModelStatus(forJoin)
}

// SetForLeftJoin 左关联
func (m *modelBaseInfo) SetForLeftJoin() {
	m.addModelStatus(forLeftJoin)
}

// SetForRightJjoin 右关联
func (m *modelBaseInfo) SetForRightJoin() {
	m.addModelStatus(forRightJoin)
}

// SetForDelete 设置为删除
func (m *modelBaseInfo) SetForDelete() {
	m.addModelStatus(forDelete)
}

// SetConList 设置当前Model的Container List
func (m *modelBaseInfo) setConList(l ModelList) {
	m.conList = l
}

func (m *modelBaseInfo) getConList() ModelList {
	return m.conList
}

// Relations 获取当前Model的子关系
func (m *modelBaseInfo) relations() RelationInfoList {
	return m.rels
}

func (m *modelBaseInfo) AppendRelation(rel *RelationInfo) {
	m.rels = append(m.rels, rel)
}

func (m *modelBaseInfo) dup() modelBaseInfo {
	d := *m
	d.rels = nil
	return d
}

type modelClause struct {
	Model
	selectedFieldIndexes []int
	inserts              exprList
	wheres               exprList
	updates              exprList
	havings              exprList
	orderBys             []refClauser
	groupBys             []refClauser
	aggs                 aggList
	limit                [2]int
}

func (m *modelClause) getInserts() exprList {
	return m.inserts
}

func (m *modelClause) appendInserts(value *Expr) {
	m.inserts = append(m.inserts, value)
}

func (m *modelClause) getOrderBys() []refClauser {
	return m.orderBys
}

func (m *modelClause) appendOrderBys(orderBy refClauser) {
	m.orderBys = append(m.orderBys, orderBy)
}

func (m *modelClause) getWheres() exprList {
	return m.wheres
}

func (m *modelClause) appendWheres(expr *Expr) {
	m.wheres = append(m.wheres, expr)
}

func (m *modelClause) getHavings() exprList {
	return m.havings
}

func (m *modelClause) appendHavings(having *Expr) {
	m.havings = append(m.havings, having)
}

// SetLimit 设置Limit子句信息
func (m *modelClause) SetLimit(limit, offset int) {
	m.limit = [2]int{limit, offset}
}

func (m *modelClause) getLimit() (limit, offset int) {
	return m.limit[0], m.limit[1]
}

func (m *modelClause) getAggs() aggList {
	return m.aggs
}

func (m *modelClause) getUpdates() exprList {
	return m.updates
}

func (m *modelClause) appendWhere(exprs ...*Expr) {
	m.wheres = append(m.wheres, exprs...)
}

func (m *modelClause) AscOrderBy(orderBy refClauser) {
	m.appendOrderBys(newOrderBy(orderBy, asc))
}

func (m *modelClause) DescOrderBy(orderBy refClauser) {
	m.appendOrderBys(newOrderBy(orderBy, desc))
}

func (m *modelClause) LookupAgg(name string) Field {
	for _, agg := range m.aggs {
		if agg.field.colName() == name {
			return agg.field
		}
	}
	return nil
}

// AndExprWhere 添加表达式where(and关系)
func (m *modelClause) AndExprWhere(expr *Expr) Model {
	expr.exp = fmt.Sprintf("AND %s", expr.exp)
	m.wheres = append(m.wheres, expr)
	return m
}

// OrExprWhere 添加表达式where(or关系)
func (m *modelClause) OrExprWhere(expr *Expr) Model {
	expr.exp = fmt.Sprintf("OR %s", expr.exp)
	m.wheres = append(m.wheres, expr)
	return m
}

// AndWhereGroup AndWhereGroup
func (m *modelClause) AndWhereGroup(wheres ...*condition) Model {
	m.appendWheres(conditionList(wheres).group(and).toExpr())
	return m
}

// OrWhereGroup OrWhereGroup
func (m *modelClause) OrWhereGroup(wheres ...*condition) Model {
	m.appendWheres(conditionList(wheres).group(or).toExpr())
	return m
}

// AndHavingGroup AndHavingGroup
func (m *modelClause) AndHavingGroup(havings ...*condition) Model {
	m.appendHavings(conditionList(havings).group(and).toExpr())
	return m
}

// OrHavingGroup OrHavingGroup
func (m *modelClause) OrHavingGroup(havings ...*condition) Model {
	m.appendHavings(conditionList(havings).group(or).toExpr())
	return m
}

// AndHaving 添加表达式having(and关系)
func (m *modelClause) AndHaving(expr *Expr) Model {
	expr.exp = fmt.Sprintf("AND %s", expr.exp)
	m.havings = append(m.havings, expr)
	return m
}

// OrHaving 添加表达式having(or关系)
func (m *modelClause) OrHaving(expr *Expr, val ...interface{}) Model {
	expr.exp = fmt.Sprintf("OR %s", expr.exp)
	m.havings = append(m.havings, expr)
	return m
}

// StrAgg 添加字符串结果的汇总
func (m *modelClause) StrAgg(expr *Expr, name string) *agg {
	agg := newStrAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// IntAgg 添加整数结果的汇总
func (m *modelClause) IntAgg(expr *Expr, name string) *agg {
	agg := newIntAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// DateAgg 添加日期结果的汇总
func (m *modelClause) DateAgg(expr *Expr, name string) *agg {
	agg := newDateAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// DatetimeAgg 添加日期时间结果的汇总
func (m *modelClause) DatetimeAgg(expr *Expr, name string) *agg {
	agg := newDatetimeAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// DecAgg 添加浮点数结果的汇总
func (m *modelClause) DecAgg(expr *Expr, name string) *agg {
	agg := newDecAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// ExprUpdate 添加表达式更新
func (m *modelClause) ExprUpdate(expr *Expr) {
	m.updates = append(m.updates, expr)
	m.addModelStatus(forUpdate)
	for parent := m.getParent(); parent != nil; parent = parent.getParent() {
		parent.addModelStatus(containSubUpdate)
	}
}

func (m *modelClause) getGroupBys() []refClauser {
	return m.groupBys
}

func (m *modelClause) appendGroupBys(groupBy refClauser) {
	m.groupBys = append(m.groupBys, groupBy)
}

func (m *modelClause) GroupBy(groupBy refClauser) {
	m.appendGroupBys(groupBy)
}

func (m *modelClause) setAggs(aggs aggList) {
	m.aggs = aggs
}

func (m *modelClause) CopyAggs(dst Model) {
	aggs := m.aggs.copy()
	dst.setAggs(aggs)
}

func (m *modelClause) appendSelectedFieldIndexes(index int) {
	m.selectedFieldIndexes = append(m.selectedFieldIndexes, index)
}

func (m *modelClause) getSelectedFieldIndexes() []int {
	return m.selectedFieldIndexes
}

func (m *modelClause) SelectAll() {
	for _, fieldInfos := range m.Model.FieldInfos() {
		m.selectedFieldIndexes = append(m.selectedFieldIndexes, fieldInfos.Index)
	}
}

func (m *modelClause) SelectFields(fields ...Field) {
	for _, f := range fields {
		m.selectedFieldIndexes = append(m.selectedFieldIndexes, f.getFieldIndex())
	}
}

func (m *modelClause) SelectExcept(fields ...Field) {
	var flag int
	for _, f := range fields {
		flag |= (1 << uint(f.getFieldIndex()))
	}
	for _, fieldInfo := range m.Model.FieldInfos() {
		if (1<<uint(fieldInfo.Index))&flag == 0 {
			m.selectedFieldIndexes = append(m.selectedFieldIndexes, fieldInfo.Index)
		}
	}
}

// Meta Model的元信息
type Meta struct {
	modelBaseInfo
	modelClause
}

// Init 初始化Meta
func (m *Meta) Init(model, parent Model, conList ModelList) {
	m.modelBaseInfo.Model = model
	m.modelBaseInfo.parent = parent
	m.modelBaseInfo.conList = conList
	m.modelClause.Model = model
	initModel(model)
}

// MarshalJSON MarshalJSON
func (m Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.modelClause.aggs)
}
