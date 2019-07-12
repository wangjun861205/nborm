package nborm

import (
	"fmt"
	"strings"
)

type RelationInfo struct {
	Fields FieldList
	Object interface{}
	Name   string
}

type RelationInfoList []RelationInfo

func (r RelationInfo) toAppendJoinClause() string {
	if len(r.Fields)%2 != 0 {
		panic(fmt.Errorf("invalid RelationInfo.Fields length(%d)", len(r.Fields)))
	}
	var builder strings.Builder
	for i, field := range r.Fields[1:] {
		if i%2 == 0 {
			builder.WriteString(fmt.Sprintf(" JOIN %s ON %s = %s ", field.fullTabName(), r.Fields[i].fullColName(), field.fullColName()))
		}
	}
	return builder.String()
}

func (r RelationInfo) toRevAppendJoinClause() string {
	if len(r.Fields)%2 != 0 {
		panic(fmt.Errorf("invalid RelationInfo.Fields length(%d)", len(r.Fields)))
	}
	var builder strings.Builder
	for i := len(r.Fields) - 2; i >= 0; i-- {
		if i%2 == 0 {
			builder.WriteString(fmt.Sprintf(" JOIN %s ON %s = %s", r.Fields[i].fullTabName(), r.Fields[i+1].fullColName(), r.Fields[i].fullColName()))
		}
	}
	return builder.String()
}

func (r RelationInfo) getMidJoinWheres() whereList {
	if len(r.Fields)%2 != 0 {
		panic(fmt.Errorf("invalid RelationInfo.Fields length(%d)", len(r.Fields)))
	}
	wheres := make(whereList, 0, 8)
	for i, f := range r.Fields[1 : len(r.Fields)-1] {
		if i%2 == 0 && f.(Model).checkStatus(containJoinWhere) {
			wheres = append(wheres, f.(Model).getJoinWheres()...)
		}
	}
	return wheres
}
