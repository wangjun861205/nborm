package nborm

import (
	"fmt"
	"strings"
)

type RelationInfo struct {
	Name      string
	SrcModel  Model
	SrcFields []Field
	MidModels []Model
	DstModel  Model
	DstFields []Field
}

func (r RelationInfo) toAppendJoinClause() string {
	var builder strings.Builder
	for _, mm := range r.MidModels {
		on, _ := mm.getOnCond().toClause()
		builder.WriteString(fmt.Sprintf("JOIN %s ON %s", mm.fullTabName(), on))
	}
	on, _ := r.DstModel.getOnCond().toClause()
	builder.WriteString(fmt.Sprintf("JOIN %s ON %s", r.DstModel.fullTabName(), on))
	return builder.String()
}

func (r RelationInfo) toAppendLeftJoinClause() string {
	var builder strings.Builder
	for _, mm := range r.MidModels {
		on, _ := mm.getOnCond().toClause()
		builder.WriteString(fmt.Sprintf("LEFT JOIN %s ON %s", mm.fullTabName(), on))
	}
	on, _ := r.DstModel.getOnCond().toClause()
	builder.WriteString(fmt.Sprintf("LEFT JOIN %s ON %s", r.DstModel.fullTabName(), on))
	return builder.String()
}

func (r RelationInfo) toAppendRightJoinClause() string {
	var builder strings.Builder
	for _, mm := range r.MidModels {
		on, _ := mm.getOnCond().toClause()
		builder.WriteString(fmt.Sprintf("RIGHT JOIN %s ON %s", mm.fullTabName(), on))
	}
	on, _ := r.DstModel.getOnCond().toClause()
	builder.WriteString(fmt.Sprintf("RIGHT JOIN %s ON %s", r.DstModel.fullTabName(), on))
	return builder.String()
}

func (r RelationInfo) toRevAppendJoinClause() string {
	var builder strings.Builder
	ons := make(exprList, 0, 8)
	models := make([]Model, 0, 8)
	ons = append(ons, r.DstModel.getOnCond())
	for i := len(r.MidModels); i > 0; i-- {
		ons = append(ons, r.MidModels[i-1].getOnCond())
		models = append(models, r.MidModels[i-1])
	}
	models = append(models, r.SrcModel)
	for i, m := range models {
		on, _ := ons[i].toClause()
		builder.WriteString(fmt.Sprintf("JOIN %s ON %s", m.fullTabName(), on))
	}
	return builder.String()
}

type RelationInfoList []RelationInfo
