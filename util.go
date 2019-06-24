package nborm

import (
	"database/sql"
	"fmt"
	"strings"
)

var DEBUG = false

func SetDebug(debug bool) {
	DEBUG = debug
}

func InitModel(model Model) {
	if model.getModelStatus()&inited == 0 {
		model.setModel(model)
		if model.getParent() == nil {
			model.setAlias("t0")
		}
		for _, fi := range model.FieldInfos() {
			fi.Field.setModel(model)
			fi.Field.setCol(fi.ColName)
			fi.Field.setField(fi.FieldName)
		}
		for _, f := range model.PrimaryKey() {
			f.setPrimaryKey()
		}
		if model.AutoIncField() != nil {
			model.AutoIncField().setAutoInc()
		}
		model.setModelStatus(inited)
	}
	for _, rel := range model.Relations() {
		relModel := rel.Object.(Model)
		relModel.setModel(relModel)
		relModel.setAlias(fmt.Sprintf("t%d", model.getIndex()))
		for _, fi := range relModel.FieldInfos() {
			fi.Field.setModel(relModel)
			fi.Field.setCol(fi.ColName)
			fi.Field.setField(fi.FieldName)
		}
		for _, f := range rel.Object.(Model).PrimaryKey() {
			f.setPrimaryKey()
		}
		if rel.Object.(Model).AutoIncField() != nil {
			rel.Object.(Model).AutoIncField().setAutoInc()
		}
		for i, rf := range rel.Fields[1 : len(rel.Fields)-1] {
			if i%2 == 0 {
				if rf.(Model).getModelStatus()&inited == 0 {
					rf.(Model).setAlias(fmt.Sprintf("t%d", model.getIndex()))
					rf.(Model).setParent(model)
				}
			}
		}
	}
}

func getAllFields(model Model) FieldList {
	fieldInfos := model.FieldInfos()
	fields := make(FieldList, len(fieldInfos))
	for i, fi := range fieldInfos {
		fields[i] = fi.Field
	}
	return fields
}

func getFields(model Model, fieldStatus fieldStatus) FieldList {
	fieldInfos := model.FieldInfos()
	fields := make(FieldList, 0, len(fieldInfos))
	for _, fi := range fieldInfos {
		fields = append(fields, fi.Field)
	}
	if fieldStatus == invalid {
		return fields
	}
	remains := make([]Field, 0, len(fields))
	for _, f := range fields {
		if f.getStatus()&fieldStatus > 0 {
			remains = append(remains, f)
		}
	}
	return remains
}

func getField(model Model, fieldName string) Field {
	fieldInfos := model.FieldInfos()
	for _, fieldInfo := range fieldInfos {
		if fieldInfo.FieldName == fieldName {
			return fieldInfo.Field
		}
	}
	panic(fmt.Sprintf("field not exists(%s.%s.%s)", model.DB(), model.Tab(), fieldName))
}

func getFieldsForScan(model Model, fields ...Field) []interface{} {
	allFields := getAllFields(model)
	addrs := make([]interface{}, 0, len(allFields))
	if len(fields) == 0 {
		for _, f := range allFields {
			addrs = append(addrs, f)
		}
		return addrs
	}
OUTER:
	for _, tf := range fields {
		for _, af := range allFields {
			if tf.colName() == af.colName() {
				addrs = append(addrs, af)
				continue OUTER
			}
		}
	}
	return addrs
}

func toInsert(field Field, cl *[]string, pl *[]string, vl *[]interface{}) {
	field.mustValid()
	*cl = append(*cl, field.colName())
	*pl = append(*pl, "?")
	*vl = append(*vl, field.Value())
}

func setRelWhere(model Model) {
	for _, info := range model.Relations() {
		if info.Fields[0].getStatus()&valid == valid {
			info.Object.(Model).AndExprWhere(NewExpr("@ = ?", info.Fields[len(info.Fields)-1]), info.Fields[0].Value())
			info.Object.(Model).setModelStatus(forModelWhere)
		}
	}
}

func scanRow(row *sql.Row, model Model, fields ...Field) error {
	scanFields := getFieldsForScan(model, fields...)
	addrs := make([]interface{}, 0, len(scanFields))
	for _, f := range scanFields {
		addrs = append(addrs, f)
	}
	if err := row.Scan(addrs...); err != nil {
		return err
	}
	if model.getModelStatus()&inited == 0 {
		InitModel(model)
	}
	setRelWhere(model)
	model.addModelStatus(synced)
	return nil
}

func scanRows(rows *sql.Rows, model Model, fields ...Field) error {
	scanFields := getFieldsForScan(model, fields...)
	addrs := make([]interface{}, 0, len(scanFields))
	for _, f := range scanFields {
		addrs = append(addrs, f)
	}
	if err := rows.Scan(addrs...); err != nil {
		return err
	}
	if model.getModelStatus()&inited == 0 {
		InitModel(model)
	}
	setRelWhere(model)
	model.addModelStatus(synced)
	return nil
}

func genUpdateSetClause(model Model) (string, []interface{}) {
	updateFields := getFields(model, forUpdate)
	updates := make(updateSetList, len(updateFields))
	for i, f := range updateFields {
		updates[i] = f.updateSet()
	}
	return updates.toClause()
}

func getSelectColumns(model Model) string {
	var builder strings.Builder
	if model.getModelStatus()&distinct == distinct {
		builder.WriteString("DISTINCT ")
	}
	selectFields := getFields(model, forSelect)
	if len(selectFields) == 0 {
		if model.getAlias() != "" {
			builder.WriteString(fmt.Sprintf("%s.*", model.getAlias()))
		} else {
			builder.WriteString(fmt.Sprintf("%s.*", model.rawFullTabName()))
		}
	} else {
		for _, f := range selectFields {
			switch {
			case f.getStatus()&forSum == forSum:
				builder.WriteString(fmt.Sprintf("IFNULL(SUM(%s), 0), ", f.fullColName()))
			default:
				builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
			}
		}
	}
	return strings.Trim(builder.String(), " ,")
}

func getSelectFields(model Model) FieldList {
	selectFields := getFields(model, forSelect)
	if len(selectFields) == 0 {
		return getAllFields(model)
	}
	return selectFields
}

func getTabRef(model Model) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s ", model.fullTabName()))
	for _, rel := range model.Relations() {
		if rel.Object.(Model).getWhere() != nil {
			builder.WriteString(rel.toAppendJoinClause())
		}
	}
	return builder.String()
}

func getOrderClause(model Model) string {
	orderFields := getFields(model, forAscOrder|forDscOrder)
	if len(orderFields) == 0 {
		return ""
	}
	colList := make([]string, 0, len(orderFields))
	for _, f := range orderFields {
		if f.getStatus()|forAscOrder == forAscOrder {
			colList = append(colList, f.fullColName())
		} else {
			colList = append(colList, fmt.Sprintf("%s DESC", f.fullColName()))
		}
	}
	return fmt.Sprintf(" ORDER BY %s ", strings.Join(colList, ", "))
}

func genPlaceHolder(val []interface{}) string {
	if len(val) == 0 {
		return ""
	}
	switch v := val[0].(type) {
	case []int:
		return fmt.Sprintf("(%s)", strings.Trim(strings.Repeat("?, ", len(v)), ", "))
	case []float32:
		return fmt.Sprintf("(%s)", strings.Trim(strings.Repeat("?, ", len(v)), ", "))
	case []float64:
		return fmt.Sprintf("(%s)", strings.Trim(strings.Repeat("?, ", len(v)), ", "))
	case []string:
		return fmt.Sprintf("(%s)", strings.Trim(strings.Repeat("?, ", len(v)), ", "))
	default:
		return "?"
	}
}

func genWhereClause(model Model) (string, []interface{}) {
	cl := make([]string, 0, 8)
	vl := make([]interface{}, 0, 8)
	where := model.getWhere()
	for _, rel := range model.Relations() {
		if rel.Object.(Model).getWhere() != nil {
			if where == nil {
				where = rel.Object.(Model).getWhere()
			} else {
				where.append(rel.Object.(Model).getWhere())
			}
		}
	}
	where.toClause(&cl, &vl)
	if len(cl) == 0 {
		return "", nil
	}
	return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(strings.Join(cl, " "), "AND "), "OR ")), vl
}

func genSimpleWhereClause(model Model) (string, []interface{}) {
	cl := make([]string, 0, 8)
	vl := make([]interface{}, 0, 8)
	model.getWhere().toClause(&cl, &vl)
	if len(cl) == 0 {
		return "", nil
	}
	return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(strings.Join(cl, " "), "AND "), "OR ")), vl
}
