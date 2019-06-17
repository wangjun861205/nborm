package model

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
		srcModel := srcField.(Model)
		dstModel := dstField.(Model)
		srcModel.setAlias("t0")
		dstModel.setAlias("t1")
		return fmt.Sprintf(`%s AS t1 JOIN %s AS t2 ON t1.%s = t2.%s`,
			srcField.fullTabName(), dstField.fullTabName(), srcField.colName(), dstField.colName())
	default:
		var builder strings.Builder
		for i, field := range r.Fields {
			switch {
			case i == 0:
				builder.WriteString(fmt.Sprintf("%s AS t%d ", field.fullTabName(), i))
				field.(Model).setAlias("t0")
			case i%2 == 1:
				builder.WriteString(fmt.Sprintf("JOIN %s AS t%d ON t%d.%s = t%d.%s ",
					field.fullTabName(), i/2+1, i/2, r.Fields[i-1].colName(), i/2+1, field.colName()))
				field.(Model).setAlias(fmt.Sprintf("t%d", i/2+1))
			}
		}
		return builder.String()
	}
}
