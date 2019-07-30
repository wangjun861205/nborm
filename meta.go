package nborm

import "fmt"

type modelStatus int

const (
	none                modelStatus = 0
	synced              modelStatus = 1
	distinct            modelStatus = 1 << 1
	forModelAgg         modelStatus = 1 << 2
	inited              modelStatus = 1 << 4
	relInited           modelStatus = 1 << 5
	forBackQuery        modelStatus = 1 << 6
	forUpdate           modelStatus = 1 << 7
	forModelOrder       modelStatus = 1 << 8
	forModelRef         modelStatus = 1 << 9
	forJoin             modelStatus = 1 << 10
	containValue        modelStatus = 1 << 11
	selectAll           modelStatus = 1 << 12
	forModelHaving      modelStatus = 1 << 13
	forReverseQuery     modelStatus = 1 << 14
	containSubJoin      modelStatus = 1 << 15
	containSubWhere     modelStatus = 1 << 16
	containJoinWhere    modelStatus = 1 << 17
	containSubJoinWhere modelStatus = 1 << 18
	containWhere        modelStatus = 1 << 19
	containSubUpdate    modelStatus = 1 << 20
	forLeftJoin         modelStatus = 1 << 21
	forRightJoin        modelStatus = 1 << 22
	containSubLeftJoin  modelStatus = 1 << 23
	containSubRightJoin modelStatus = 1 << 24
	containSubOrder     modelStatus = 1 << 25
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

func (m *modelBaseInfo) setAlias() {
	m.alias = fmt.Sprintf("t%d", m.getIndex())
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

func (m *modelBaseInfo) getIndex() int {
	if m.parent != nil {
		return m.parent.getIndex()
	}
	m.index++
	return m.index
}

// SetForJoin 设置Join查询标志位(所有Father Model的containSubJoin标志位均会被置为1)
func (m *modelBaseInfo) SetForJoin() {
	m.getParent().addModelStatus(containSubJoin)
	m.addModelStatus(forJoin)
}

// SetForLeftJoin 左关联
func (m *modelBaseInfo) SetForLeftJoin() {
	m.getParent().addModelStatus(containSubLeftJoin)
	m.addModelStatus(forLeftJoin)
}

// SetForRightJjoin 右关联
func (m *modelBaseInfo) SetForRightJoin() {
	m.getParent().addModelStatus(containSubLeftJoin)
	m.addModelStatus(forRightJoin)
}

// SelectAll 显式设置查询所有字段，拼合Select语句时，该Model的字段将以alias.*的方式出现
func (m *modelBaseInfo) SelectAll() {
	m.addModelStatus(selectAll)
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

type modelClause struct {
	Model
	wheres  exprList
	havings exprList
	updates exprList
	aggExps []*aggExp
	limit   [2]int
}

func (m *modelClause) getWheres() exprList {
	return m.wheres
}

func (m *modelClause) getHavings() exprList {
	return m.havings
}

// SetLimit 设置Limit子句信息
func (m *modelClause) SetLimit(limit, offset int) {
	m.limit = [2]int{limit, offset}
}

func (m *modelClause) getLimit() (limit, offset int) {
	return m.limit[0], m.limit[1]
}

func (m *modelClause) getAggExps() []*aggExp {
	return m.aggExps
}

func (m *modelClause) getUpdateList() exprList {
	return m.updates
}

func (m *modelClause) appendWhere(exprs ...*Expr) {
	m.wheres = append(m.wheres, exprs...)
}

// AndExprWhere 添加表达式where(and关系)
func (m *modelClause) AndExprWhere(expr *Expr) Model {
	expr.exp = fmt.Sprintf("AND %s", expr.exp)
	m.wheres = append(m.wheres, expr)
	m.addModelStatus(containWhere)
	m.addModelStatus(forModelRef)
	for parent := m.getParent(); parent != nil; parent = parent.getParent() {
		parent.addModelStatus(containSubWhere)
	}
	return m
}

// OrExprWhere 添加表达式where(or关系)
func (m *modelClause) OrExprWhere(expr *Expr) Model {
	expr.exp = fmt.Sprintf("OR %s", expr.exp)
	m.wheres = append(m.wheres, expr)
	m.addModelStatus(containWhere)
	m.addModelStatus(forModelRef)
	for parent := m.getParent(); parent != nil; parent = parent.getParent() {
		parent.addModelStatus(containSubWhere)
	}
	return m
}

// AndHaving 添加表达式having(and关系)
func (m *modelClause) AndHaving(expr *Expr) Model {
	expr.exp = fmt.Sprintf("AND %s", expr.exp)
	m.havings = append(m.havings, expr)
	m.addModelStatus(forModelHaving)
	return m
}

// OrHaving 添加表达式having(or关系)
func (m *modelClause) OrHaving(expr *Expr, val ...interface{}) Model {
	expr.exp = fmt.Sprintf("OR %s", expr.exp)
	m.havings = append(m.havings, expr)
	m.addModelStatus(forModelHaving)
	return m
}

// StrAgg 添加字符串结果的汇总
func (m *modelClause) StrAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newStrAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

// IntAgg 添加整数结果的汇总
func (m *modelClause) IntAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newIntAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

// DateAgg 添加日期结果的汇总
func (m *modelClause) DateAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newDateAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

// DatetimeAgg 添加日期时间结果的汇总
func (m *modelClause) DatetimeAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newDatetimeAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

// DecAgg 添加浮点数结果的汇总
func (m *modelClause) DecAgg(expr *Expr, name string) {
	expr.exp = fmt.Sprintf("%s AS %s", expr.exp, name)
	m.aggExps = append(m.aggExps, newDecAgg(expr, name))
	m.addModelStatus(forModelAgg)
}

// ExprUpdate 添加表达式更新
func (m *modelClause) ExprUpdate(expr *Expr) {
	m.updates = append(m.updates, expr)
	m.addModelStatus(forUpdate)
	for parent := m.getParent(); parent != nil; parent = parent.getParent() {
		parent.addModelStatus(containSubUpdate)
	}
}

// Meta Model的元信息
type Meta struct {
	modelBaseInfo
	modelClause
}

func (m *Meta) setModel(model Model) {
	m.modelBaseInfo.Model = model
	m.modelClause.Model = model
}
