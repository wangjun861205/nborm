package nborm

import (
	"encoding/json"
	"fmt"
	"io"
)

type modelStatus int

const (
	none             modelStatus = 0
	synced           modelStatus = 1
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
	selectForUpdate  modelStatus = 1 << 27
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

func (m *modelBaseInfo) toRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(m.fullTabName()))
	w.Write([]byte(" "))
}

func (m *modelBaseInfo) toSimpleRefClause(w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	w.Write([]byte(m.rawFullTabName()))
	w.Write([]byte(" "))
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
// func (m *modelBaseInfo) SelectDistinct() Model {
// 	m.addModelStatus(distinct)
// 	return m
// }

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
func (m *modelBaseInfo) SetForJoin() Model {
	m.addModelStatus(forJoin)
	return m
}

// SetForLeftJoin 左关联
func (m *modelBaseInfo) SetForLeftJoin() Model {
	m.addModelStatus(forLeftJoin)
	return m
}

// SetForRightJjoin 右关联
func (m *modelBaseInfo) SetForRightJoin() Model {
	m.addModelStatus(forRightJoin)
	return m
}

// SetForDelete 设置为删除
func (m *modelBaseInfo) SetForDelete() Model {
	m.addModelStatus(forDelete)
	return m
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
	selectors selectorList
	wheres    wherer
	updates   updateList
	havings   havinger
	orderBys  orderByList
	groupBys  groupByList
	aggs      aggList
	limit     [2]int
}

func (m *modelClause) SelectDistinct() Model {
	m.selectors.addStatus(selectorStatusDistinct)
	return m
}

func (m *modelClause) getInserts() insertList {
	l := make(insertList, 0, 8)
	for _, info := range m.FieldInfos() {
		info.Field.toInsert(&l)
	}
	return l
}

func (m *modelClause) getOrderBys() orderByList {
	return m.orderBys
}

func (m *modelClause) appendOrderBys(orderBy *orderBy) {
	m.orderBys = append(m.orderBys, orderBy)
}

func (m *modelClause) getWheres() wherer {
	return m.wheres
}

func (m *modelClause) appendWheres(w wherer) {
	lastWhere := m.wheres.lastNode()
	if lastWhere == nil {
		m.wheres = w
	} else {
		lastWhere.append(w)
	}
}

func (m *modelClause) getHavings() havinger {
	return m.havings
}

func (m *modelClause) appendHavings(h havinger) {
	lastHaving := m.havings.lastNode()
	if lastHaving == nil {
		m.havings = h
	} else {
		lastHaving.append(h)
	}
}

// SetLimit 设置Limit子句信息
func (m *modelClause) SetLimit(limit, offset int) {
	m.limit = [2]int{limit, offset}
}

func (m *modelClause) appendSelector(s selector) {
	m.selectors.list = append(m.selectors.list, s)
}

func (m *modelClause) getLimit() (limit, offset int) {
	return m.limit[0], m.limit[1]
}

func (m *modelClause) getAggs() aggList {
	return m.aggs
}

func (m *modelClause) getUpdates() updateList {
	return m.updates
}

func (m *modelClause) appendUpdate(update *update) Model {
	m.updates = append(m.updates, update)
	return m
}

func (m *modelClause) AscOrderBy(orderBy referencer) Model {
	m.appendOrderBys(newOrderBy(orderBy, asc))
	return m
}

func (m *modelClause) DescOrderBy(orderBy referencer) Model {
	m.appendOrderBys(newOrderBy(orderBy, desc))
	return m
}

func (m *modelClause) LookupAgg(name string) Field {
	for _, agg := range m.aggs {
		if agg.getField().colName() == name {
			return agg.getField()
		}
	}
	return nil
}

// AndExprWhere 添加表达式where(and关系)
func (m *modelClause) AndExprWhere(expr *Expr) Model {
	w := newWhere(expr, whereAnd)
	if m.wheres == nil {
		m.wheres = w
	} else {
		m.wheres.lastNode().append(w)
	}
	return m
}

// OrExprWhere 添加表达式where(or关系)
func (m *modelClause) OrExprWhere(expr *Expr) Model {
	w := newWhere(expr, whereOr)
	if m.wheres == nil {
		m.wheres = w
	} else {
		m.wheres.lastNode().append(w)
	}
	return m
}

// AndWhereGroup AndWhereGroup
func (m *modelClause) AndModelWhereGroup(wheres ...wherer) Model {
	group := groupWherers(whereAnd, wheres...)
	if m.wheres == nil {
		m.wheres = group
	} else {
		m.wheres.lastNode().append(group)
	}
	return m
}

// OrWhereGroup OrWhereGroup
func (m *modelClause) OrModelWhereGroup(wheres ...wherer) Model {
	group := groupWherers(whereOr, wheres...)
	if m.wheres == nil {
		m.wheres = group
	} else {
		m.wheres.lastNode().append(group)
	}
	return m
}

// AndHavingGroup AndHavingGroup
func (m *modelClause) AndHavingGroup(havings ...havinger) Model {
	group := groupHavings(whereAnd, havings...)
	if m.havings == nil {
		m.havings = group
	} else {
		m.havings.lastNode().append(group)
	}
	return m
}

// OrHavingGroup OrHavingGroup
func (m *modelClause) OrHavingGroup(havings ...havinger) Model {
	group := groupHavings(whereOr, havings...)
	if m.havings == nil {
		m.havings = group
	} else {
		m.havings.lastNode().append(group)
	}
	return m
}

// AndHaving 添加表达式having(and关系)
func (m *modelClause) AndHaving(expr *Expr) Model {
	h := newHaving(expr, whereAnd)
	if m.havings == nil {
		m.havings = h
	} else {
		m.havings.lastNode().append(h)
	}
	return m
}

// OrHaving 添加表达式having(or关系)
func (m *modelClause) OrHaving(expr *Expr) Model {
	h := newHaving(expr, whereOr)
	if m.havings == nil {
		m.havings = h
	} else {
		m.havings.lastNode().append(h)
	}
	return m
}

func (m *modelClause) appendAgg(agg aggregator) {
	m.aggs = append(m.aggs, agg)
}

// StrAgg 添加字符串结果的汇总
func (m *modelClause) StrAgg(expr *Expr, name string) *StrAgg {
	agg := newStrAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// IntAgg 添加整数结果的汇总
func (m *modelClause) IntAgg(expr *Expr, name string) *IntAgg {
	agg := newIntAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.appendSelector(agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// DateAgg 添加日期结果的汇总
func (m *modelClause) DateAgg(expr *Expr, name string) *DateAgg {
	agg := newDateAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.appendSelector(agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// DatetimeAgg 添加日期时间结果的汇总
func (m *modelClause) DatetimeAgg(expr *Expr, name string) *DatetimeAgg {
	agg := newDatetimeAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.appendSelector(agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// TimeAgg 添加时间结果的汇总
func (m *modelClause) TimeAgg(expr *Expr, name string) *TimeAgg {
	agg := newTimeAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.appendSelector(agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// DecAgg 添加浮点数结果的汇总
func (m *modelClause) DecAgg(expr *Expr, name string) *DecimalAgg {
	agg := newDecAgg(expr, name)
	m.aggs = append(m.aggs, agg)
	m.appendSelector(agg)
	m.addModelStatus(forModelAgg)
	return agg
}

// ExprUpdate 添加表达式更新
func (m *modelClause) ExprUpdate(field referencer, expr *Expr) Model {
	m.updates = append(m.updates, newUpdate(field, expr))
	m.addModelStatus(forUpdate)
	for parent := m.getParent(); parent != nil; parent = parent.getParent() {
		parent.addModelStatus(containSubUpdate)
	}
	return m
}

func (m *modelClause) getGroupBys() groupByList {
	return m.groupBys
}

func (m *modelClause) appendGroupBys(ref referencer) {
	m.groupBys = append(m.groupBys, &groupBy{ref})
}

func (m *modelClause) ModelGroupBy(groupBy referencer) Model {
	m.appendGroupBys(groupBy)
	return m
}

func (m *modelClause) setAggs(aggs aggList) {
	m.aggs = aggs
}

func (m *modelClause) CopyAggs(dst Model) {
	aggs := m.aggs.copy()
	dst.setAggs(aggs)
}

func (m *modelClause) getSelectors() *selectorList {
	return &m.selectors
}

// SelectAll 选择所有字段
func (m *modelClause) SelectAll() Model {
	for _, info := range m.Model.FieldInfos() {
		m.appendSelector(info.Field)
	}
	return m
}

// func (m *modelClause) groupBySelectedFields(refs *[]referencer) {
// 	fields := getSelectFields(m)
// 	for _, field := range fields {
// 		*refs = append(*refs, field)
// 	}
// 	for _, relInfo := range m.relations() {
// 		if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
// 			relInfo.lastModel().groupBySelectedFields(refs)
// 		}
// 	}
// }

func (m *modelClause) GroupBySelectedFields() Model {
	fields := getSelectFields(m)
	for _, field := range fields {
		m.appendGroupBys(field)
	}
	for _, relInfo := range m.relations() {
		if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
			relInfo.lastModel().GroupBySelectedFields()
		}
	}
	return m
}

// SelectFields 选择某些字段
func (m *modelClause) SelectFields(fields ...Field) Model {
	for _, f := range fields {
		m.appendSelector(f)
	}
	return m
}

func (m *modelClause) SelectExcept(fields ...Field) Model {
	var flag int
	for _, f := range fields {
		flag |= (1 << uint(f.getFieldIndex()))
	}
	for _, fieldInfo := range m.Model.FieldInfos() {
		if (1<<uint(fieldInfo.Index))&flag == 0 {
			m.appendSelector(fieldInfo.Field)
		}
	}
	return m
}

func (m *modelClause) SelectForUpdate() Model {
	m.addModelStatus(selectForUpdate)
	return m
}

// Meta Model的元信息
type Meta struct {
	modelBaseInfo
	modelClause
}

func (m Meta) String() string {
	return fmt.Sprintf("Aggs: %v", m.aggs)
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
