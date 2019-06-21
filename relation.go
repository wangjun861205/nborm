package nborm

import (
	"fmt"
	"strings"
)

type RelationInfo struct {
	Fields FieldList
	Object interface{}
}

type RelationInfoList []RelationInfo

func (r RelationInfo) toJoinClause() string {
	switch {
	case len(r.Fields)%2 != 0:
		panic(fmt.Errorf("invalid RelationInfo.Fields length(%d)", len(r.Fields)))
	case len(r.Fields) == 2:
		srcField := r.Fields[0]
		dstField := r.Fields[1]
		dstModel := dstField.(Model)
		dstModel.setAlias("t1")
		return fmt.Sprintf(`%s JOIN %s ON t0.%s = t1.%s`,
			srcField.fullTabName(),
			dstField.fullTabName(),
			srcField.colName(),
			dstField.colName())
	default:
		var builder strings.Builder
		for i, field := range r.Fields {
			switch {
			case i == 0:
				builder.WriteString(fmt.Sprintf("%s AS t%d ", field.rawFullTabName(), i))
			case i%2 == 1:
				field.(Model).setAlias(fmt.Sprintf("t%d", i/2+1))
				builder.WriteString(fmt.Sprintf("JOIN %s AS t%d ON t%d.%s = t%d.%s ",
					field.rawFullTabName(), i/2+1, i/2, r.Fields[i-1].colName(), i/2+1, field.colName()))
				for _, wf := range getFields(field.(Model), forWhere) {
					wl := wf.whereList()
					builder.WriteString(wl.toOnClause())
				}
			}
		}
		return builder.String()
	}
}

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
