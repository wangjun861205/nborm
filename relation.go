package nborm

import (
	"strings"
)

type joinType string

const (
	join      joinType = "JOIN"
	leftJoin  joinType = "LEFT JOIN"
	rightJoin joinType = "RIGHT JOIN"
)

// RelationInfo 关系信息
type RelationInfo struct {
	name     string
	dstModel Model
	on       *Expr
	next     *RelationInfo
	prev     *RelationInfo
}

// NewRelationInfo 新建关系
func NewRelationInfo(name string, dst Model, on *Expr) *RelationInfo {
	return &RelationInfo{
		name,
		dst,
		on,
		nil,
		nil,
	}
}

func (r *RelationInfo) lastModel() Model {
	last := r
	for last.next != nil {
		last = last.next
	}
	return last.dstModel
}

// Append Append
func (r *RelationInfo) Append(name string, dst Model, on *Expr) *RelationInfo {
	nr := &RelationInfo{
		name,
		dst,
		on,
		nil,
		r,
	}
	if r == nil {
		return nr
	}
	last := r
	for last.next != nil {
		last = last.next
	}
	last.next = nr
	nr.prev = last
	return r
}

func (r *RelationInfo) toClause(joinType joinType) string {
	var builder strings.Builder
	for rel := r; rel != nil; rel = rel.next {
		builder.WriteString(string(joinType))
		builder.WriteString(" ")
		builder.WriteString(rel.dstModel.fullTabName())
		builder.WriteString(" ON ")
		onClause, _ := rel.on.toClause()
		builder.WriteString(onClause)
		builder.WriteString(" ")
	}
	return builder.String()
}

func (r *RelationInfo) toRevClause(srcModel Model, joinType joinType) string {
	last := r
	for last.next != nil {
		last = last.next
	}
	var builder strings.Builder
	for last.prev != nil {
		builder.WriteString(string(joinType))
		builder.WriteString(" ")
		builder.WriteString(last.prev.dstModel.fullTabName())
		builder.WriteString(" ON ")
		onClause, _ := last.on.toClause()
		builder.WriteString(onClause)
		builder.WriteString(" ")
		last = last.prev
	}
	builder.WriteString(string(joinType))
	builder.WriteString(" ")
	builder.WriteString(srcModel.fullTabName())
	builder.WriteString(" ON ")
	onClause, _ := last.on.toClause()
	builder.WriteString(onClause)
	return builder.String()
}

// RelationInfoList 关系列表
type RelationInfoList []*RelationInfo
