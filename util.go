package model

import (
	"database/sql"
	"fmt"
	"strings"
)

const DEBUG = true

func InitModel(model Model) {
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
}

func initRelation(model Model) {
	infos := model.Relations()
	for _, info := range infos {
		m := info.Object.(Model)
		InitModel(m)
		m.setRel("", fmt.Sprintf("%s.*", info.Fields[0].fullTabName()), info.toJoinClause(), info.Fields[0].genAndWhere("=", info.Fields[0].Value()))
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

func scanRow(row *sql.Row, model Model, fields ...Field) error {
	scanFields := getFieldsForScan(model, fields...)
	addrs := make([]interface{}, 0, len(scanFields))
	for _, f := range scanFields {
		addrs = append(addrs, f)
	}
	if err := row.Scan(addrs...); err != nil {
		return err
	}
	initRelation(model)
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
	initRelation(model)
	return nil
}

func genWhereClause(model Model) (string, []interface{}) {
	whereFields := getFields(model, forWhere)
	whereList := make(whereList, 0, len(whereFields)*2)
	for _, f := range whereFields {
		whereList = append(whereList, f.whereList()...)
	}
	return whereList.toClause()
}

func genUpdateSetClause(model Model) (string, []interface{}) {
	updateFields := getFields(model, forUpdate)
	updates := make(updateSetList, len(updateFields))
	for i, f := range updateFields {
		updates[i] = f.updateSet()
	}
	return updates.toClause()
}

func genWhereList(model Model) whereList {
	whereFields := getFields(model, forWhere)
	whereList := make(whereList, 0, len(whereFields)*2)
	for _, f := range whereFields {
		whereList = append(whereList, f.whereList()...)
	}
	return whereList
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
			builder.WriteString(fmt.Sprintf("%s.%s.*", model.DB(), model.Tab()))
		}
	} else {
		for _, f := range selectFields {
			builder.WriteString(fmt.Sprintf("%s, ", f.fullFieldName()))
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
