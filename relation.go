package nborm

import (
	"io"
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

func (r *RelationInfo) toClause(joinType joinType, w io.Writer, vals *[]interface{}) {
	for rel := r; rel != nil; rel = rel.next {
		w.Write([]byte(string(joinType)))
		w.Write([]byte(" "))
		w.Write([]byte(rel.dstModel.fullTabName()))
		w.Write([]byte(" ON "))
		rel.on.toClause(w, vals)
	}
}

func (r *RelationInfo) toRevClause(srcModel Model, joinType joinType, w io.Writer, vals *[]interface{}) {
	last := r
	for last.next != nil {
		last = last.next
	}
	for last.prev != nil {
		w.Write([]byte(string(joinType)))
		w.Write([]byte(" "))
		w.Write([]byte(last.prev.dstModel.fullTabName()))
		w.Write([]byte(" ON "))
		last.on.toClause(w, vals)
		last = last.prev
	}
	w.Write([]byte(string(joinType)))
	w.Write([]byte(" "))
	w.Write([]byte(srcModel.fullTabName()))
	w.Write([]byte(" ON "))
	last.on.toClause(w, vals)
}

// RelationInfoList 关系列表
type RelationInfoList []*RelationInfo
