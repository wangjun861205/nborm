package nborm

import (
	"fmt"
	"strings"
)

type RelationInfo struct {
	Name       string
	SrcModel   Model
	SrcFields  []Field
	MidModels  []Model
	DstModel   Model
	DstFields  []Field
	JoinWheres exprList
}

func (r RelationInfo) toAppendJoinClause() string {
	var builder strings.Builder
	for _, mm := range r.MidModels {
		builder.WriteString(fmt.Sprintf("JOIN %s", mm.fullTabName()))
	}
	builder.WriteString(fmt.Sprintf("JOIN %s", r.DstModel.fullTabName()))
	return builder.String()
}

func (r RelationInfo) toAppendLeftJoinClause() string {
	var builder strings.Builder
	prevModel := r.SrcModel
	for _, mm := range r.MidModels {
		builder.WriteString(fmt.Sprintf("LEFT JOIN %s", mm.fullTabName()))
	}
	builder.WriteString(fmt.Sprintf("LEFT JOIN %s", r.DstModel.fullTabName()))
	return builder.String()
}

func (r RelationInfo) toAppendRightJoinClause() string {
	var builder strings.Builder
	for _, mm := range r.MidModels {
		builder.WriteString(fmt.Sprintf("RIGHT JOIN %s", mm.fullTabName()))
	}
	builder.WriteString(fmt.Sprintf("RIGHT JOIN %s", r.DstModel.fullTabName()))
	return builder.String()
}

func (r RelationInfo) toRevAppendJoinClause() string {
	var builder strings.Builder
	for _, mm := range r.MidModels {
		builder.WriteString(fmt.Sprintf("JOIN %s", mm.fullTabName()))
	}
	builder.WriteString(fmt.Sprintf("JOIN %s", r.SrcModel.fullTabName()))
	return builder.String()
}

func (r RelationInfo) getMidJoinWheres() exprList {
	return r.JoinWheres
}

// type RelationInfo struct {
// 	Fields FieldList
// 	Object interface{}
// 	Name   string
// }

type RelationInfoList []RelationInfo

// func (r RelationInfo) toAppendJoinClause() string {
// 	if len(r.Fields)%2 != 0 {
// 		panic(fmt.Errorf("invalid RelationInfo.Fields length(%d)", len(r.Fields)))
// 	}
// 	var builder strings.Builder
// 	for i, field := range r.Fields[1:] {
// 		if i%2 == 0 {
// 			builder.WriteString(fmt.Sprintf(" JOIN %s ON %s = %s ", field.fullTabName(), r.Fields[i].fullColName(), field.fullColName()))
// 		}
// 	}
// 	return builder.String()
// }

// func (r RelationInfo) toRevAppendJoinClause() string {
// 	if len(r.Fields)%2 != 0 {
// 		panic(fmt.Errorf("invalid RelationInfo.Fields length(%d)", len(r.Fields)))
// 	}
// 	var builder strings.Builder
// 	for i := len(r.Fields) - 2; i >= 0; i-- {
// 		if i%2 == 0 {
// 			builder.WriteString(fmt.Sprintf(" JOIN %s ON %s = %s", r.Fields[i].fullTabName(), r.Fields[i+1].fullColName(), r.Fields[i].fullColName()))
// 		}
// 	}
// 	return builder.String()
// }

// func (r RelationInfo) getMidJoinWheres() exprList {
// 	if len(r.Fields)%2 != 0 {
// 		panic(fmt.Errorf("invalid RelationInfo.Fields length(%d)", len(r.Fields)))
// 	}
// 	wheres := make(exprList, 0, 8)
// 	for i, f := range r.Fields[1 : len(r.Fields)-1] {
// 		if i%2 == 0 && f.(Model).checkStatus(containJoinWhere) {
// 			wheres = append(wheres, f.(Model).getJoinWheres()...)
// 		}
// 	}
// 	return wheres
// }
