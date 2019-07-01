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
		model.setAlias()
		for _, fi := range model.FieldInfos() {
			fi.Field.setModel(model)
			fi.Field.setCol(fi.ColName)
			fi.Field.setField(fi.FieldName)
		}
		model.addModelStatus(inited)
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

// func scanRow(row *sql.Row, model Model, fields ...Field) error {
// 	scanFields := getFieldsForScan(model, fields...)
// 	addrs := make([]interface{}, 0, len(scanFields))
// 	for _, f := range scanFields {
// 		addrs = append(addrs, f)
// 	}
// 	if err := row.Scan(addrs...); err != nil {
// 		return err
// 	}
// 	if model.getModelStatus()&inited == 0 {
// 		InitModel(model)
// 	}
// 	// setRelWhere(model)
// 	model.addModelStatus(synced)
// 	return nil
// }

// func scanRows(rows *sql.Rows, model Model, fields ...Field) error {
// 	scanFields := getFieldsForScan(model, fields...)
// 	addrs := make([]interface{}, 0, len(scanFields))
// 	for _, f := range scanFields {
// 		addrs = append(addrs, f)
// 	}
// 	if err := rows.Scan(addrs...); err != nil {
// 		return err
// 	}
// 	if model.getModelStatus()&inited == 0 {
// 		InitModel(model)
// 	}
// 	// setRelWhere(model)
// 	model.addModelStatus(synced)
// 	return nil
// }

func scanRow(row *sql.Row, model Model) error {
	selectFields := getSelectFields(model)
	if l, ok := model.(ModelList); ok {
		newModel := l.NewModel()
		scanFields := getFieldsForScan(newModel, selectFields...)
		return row.Scan(scanFields...)
	}
	scanFields := getSelectFields(model)
	addrs := make([]interface{}, 0, len(scanFields))
	for _, f := range scanFields {
		addrs = append(addrs, f)
	}
	if err := row.Scan(addrs...); err != nil {
		return err
	}
	if model.getModelStatus()&relInited == 0 {
		model.InitRel()
	}
	model.addModelStatus(synced)
	return nil
}

func scanRows(rows *sql.Rows, l ModelList) error {
	defer rows.Close()
	selectFields := getSelectFields(l)
	for rows.Next() {
		newModel := l.NewModel()
		scanFields := getFieldsForScan(newModel, selectFields...)
		rows.Scan(scanFields...)
	}
	return rows.Err()
}

func joinScanRow(row *sql.Row, m Model) error {
	addrs := make([]interface{}, 0, 64)
	for _, f := range getSelectFields(m) {
		addrs = append(addrs, f)
	}
	for _, relInfo := range m.Relations() {
		if relInfo.Object.(Model).getModelStatus()&forJoin == forJoin {
			if l, ok := m.(ModelList); ok {
				selectFields := getSelectFields(l)
				nm := l.NewModel()
				for _, f := range getFieldsForScan(nm, selectFields...) {
					addrs = append(addrs, f)
				}
				nm.addModelStatus(synced)
			} else {
				for _, f := range getSelectFields(relInfo.Object.(Model)) {
					addrs = append(addrs, f)
				}
				relInfo.Object.(Model).addModelStatus(synced)
			}
		}
	}
	return row.Scan(addrs...)
}

func IsPrimaryKeyEqual(lm, rm Model) bool {
	lpk := lm.PrimaryKey()
	rpk := rm.PrimaryKey()
	for i := 0; i < len(lpk); i++ {
		if lpk[i].Value() != rpk[i].Value() {
			return false
		}
	}
	return true
}

func joinScanRows(rows *sql.Rows, m Model) error {
	defer rows.Close()
	classModels := make([]Model, 0, 8)
	classModels = append(classModels, m)
	for _, relInfo := range m.Relations() {
		relModel := relInfo.Object.(Model)
		if relModel.getModelStatus()&forJoin == forJoin {
			classModels = append(classModels, relModel)
		}
	}
	for rows.Next() {
		instModels := make([]Model, 0, len(classModels))
		for _, cm := range classModels {
			if cl, ok := cm.(ModelList); ok {
				instModels = append(instModels, cl.NewModel())
			} else {
				instModels = append(instModels, cm)
			}
		}
		addrs := make([]interface{}, 0, 64)
		for i, im := range instModels {
			im.addModelStatus(synced)
			selectFields := getSelectFields(classModels[i])
			addrs = append(addrs, getFieldsForScan(im, selectFields...)...)
		}
		if err := rows.Scan(addrs...); err != nil {
			return err
		}
		if l, ok := m.(ModelList); ok {
			l.Collapse()
		}
	}
	return rows.Err()
}

func genFullUpdateSetClause(model Model) (string, []interface{}) {
	updates := make(updateSetList, 0, 16)
	parent := model.GetParent()
	if parent != nil && parent.getModelStatus()&forModelUpdate == forModelUpdate {
		for _, f := range getFields(model, forUpdate) {
			updates = append(updates, f.updateSet())
		}
	}
	if model.getModelStatus()&forModelUpdate == forModelUpdate {
		for _, f := range getFields(model, forUpdate) {
			updates = append(updates, f.updateSet())
		}
	}
	for _, relInfo := range model.Relations() {
		if relInfo.Object.(Model).getModelStatus()&forModelUpdate == forModelUpdate {
			for _, f := range getFields(relInfo.Object.(Model), forUpdate) {
				updates = append(updates, f.updateSet())
			}
		}
	}
	return updates.toClause()
}

func genUpdateSetClause(model Model) (string, []interface{}) {
	updates := make(updateSetList, 0, 16)
	parent := model.GetParent()
	if parent != nil && parent.getModelStatus()&forModelUpdate == forModelUpdate {
		for _, f := range getFields(model, forUpdate) {
			updates = append(updates, f.updateSet())
		}
	}
	if model.getModelStatus()&forModelUpdate == forModelUpdate {
		for _, f := range getFields(model, forUpdate) {
			updates = append(updates, f.updateSet())
		}
	}
	return updates.toClause()
}

func genSimpleSelectColumns(model Model) string {
	var builder strings.Builder
	if model.getModelStatus()&distinct == distinct {
		builder.WriteString("DISTINCT ")
	}
	selectFields := getFields(model, forSelect)
	if len(selectFields) == 0 {
		builder.WriteString(fmt.Sprintf("%s.*", model.rawFullTabName()))
	} else {
		for _, f := range selectFields {
			switch {
			case f.getStatus()&forSum == forSum:
				builder.WriteString(fmt.Sprintf("IFNULL(SUM(%s), 0), ", f.rawFullColName()))
			default:
				builder.WriteString(fmt.Sprintf("%s, ", f.rawFullColName()))
			}
		}
	}
	return strings.Trim(builder.String(), " ,")
}

func genSelectColumns(model Model) string {
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

func genJoinSelectColumns(model Model) string {
	var builder strings.Builder
	if model.getModelStatus()&distinct == distinct {
		builder.WriteString("DISTINCT ")
	}
	selectFields := getFields(model, forSelect)
	if len(selectFields) == 0 {
		builder.WriteString(fmt.Sprintf("%s.*, ", model.getAlias()))
	} else {
		for _, f := range selectFields {
			builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
		}
	}
	for _, relInfo := range model.Relations() {
		relModel := relInfo.Object.(Model)
		if relModel.getModelStatus()&forJoin == forJoin {
			subSelectFields := getFields(relModel, forSelect)
			if len(subSelectFields) == 0 {
				builder.WriteString(fmt.Sprintf("%s.*, ", relModel.getAlias()))
			} else {
				for _, f := range subSelectFields {
					builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
				}
			}
		}
	}
	return strings.Trim(builder.String(), ", ")
}

func getSelectFields(model Model) FieldList {
	selectFields := getFields(model, forSelect)
	if len(selectFields) == 0 {
		return getAllFields(model)
	}
	return selectFields
}

func genFullOrderClause(model Model) string {
	colList := make([]string, 0, 8)
	parent := model.GetParent()
	if parent != nil && parent.getModelStatus()&forModelOrder == forModelOrder {
		for _, f := range getFields(parent, forAscOrder|forDscOrder) {
			if f.getStatus()&forAscOrder == forAscOrder {
				colList = append(colList, f.fullColName())
			} else {
				colList = append(colList, fmt.Sprintf("%s DESC"))
			}
		}
	}
	if model.getModelStatus()&forModelOrder == forModelOrder {
		for _, f := range getFields(model, forAscOrder|forDscOrder) {
			if f.getStatus()&forAscOrder == forAscOrder {
				colList = append(colList, f.fullColName())
			} else {
				colList = append(colList, fmt.Sprintf("%s DESC"))
			}
		}
	}
	for _, relInfo := range model.Relations() {
		if relInfo.Object.(Model).getModelStatus()&forModelOrder == forModelOrder {
			for _, f := range getFields(relInfo.Object.(Model), forAscOrder|forDscOrder) {
				if f.getStatus()&forAscOrder == forAscOrder {
					colList = append(colList, f.fullColName())
				} else {
					colList = append(colList, fmt.Sprintf("%s DESC"))
				}
			}
		}
	}
	if len(colList) == 0 {
		return ""
	}
	return fmt.Sprintf("ORDER BY %s", strings.Join(colList, ", "))
}

func genOrderClause(model Model) string {
	colList := make([]string, 0, 8)
	parent := model.GetParent()
	if parent != nil && parent.getModelStatus()&forModelOrder == forModelOrder {
		for _, f := range getFields(parent, forAscOrder|forDscOrder) {
			if f.getStatus()&forAscOrder == forAscOrder {
				colList = append(colList, f.fullColName())
			} else {
				colList = append(colList, fmt.Sprintf("%s DESC", f.fullColName()))
			}
		}
	}
	if model.getModelStatus()&forModelOrder == forModelOrder {
		for _, f := range getFields(model, forAscOrder|forDscOrder) {
			if f.getStatus()&forAscOrder == forAscOrder {
				colList = append(colList, f.fullColName())
			} else {
				colList = append(colList, fmt.Sprintf("%s DESC", f.fullColName()))
			}
		}
	}
	if len(colList) == 0 {
		return ""
	}
	return fmt.Sprintf("ORDER BY %s", strings.Join(colList, ", "))
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

func genFullTabRef(model Model) string {
	var builder strings.Builder
	if model.GetParent() == nil {
		builder.WriteString(model.fullTabName())
		for _, relInfo := range model.Relations() {
			if relInfo.Object.(Model).getModelStatus()&forModelRef == forModelRef {
				builder.WriteString(relInfo.toAppendJoinClause())
			}
		}
	} else {
		parent := model.GetParent()
		switch {
		case (parent.getModelStatus()&synced)|(parent.getModelStatus()&forModelRef) > 0:
			for _, relInfo := range parent.Relations() {
				if relInfo.Object.(Model) == model {
					builder.WriteString(parent.fullTabName())
					builder.WriteString(relInfo.toAppendJoinClause())
					break
				}
			}
		default:
			builder.WriteString(model.fullTabName())
		}
		for _, relInfo := range model.Relations() {
			if relInfo.Object.(Model).getModelStatus()&forModelRef == forModelRef {
				builder.WriteString(relInfo.toAppendJoinClause())
			}
		}
	}
	return builder.String()
}

func genTabRef(model Model) string {
	var builder strings.Builder
	if model.GetParent() == nil {
		builder.WriteString(model.fullTabName())
	} else {
		parent := model.GetParent()
		switch {
		case (parent.getModelStatus()&synced)|(parent.getModelStatus()&forModelRef) > 0:
			for _, relInfo := range parent.Relations() {
				if relInfo.Object.(Model) == model {
					builder.WriteString(parent.fullTabName())
					builder.WriteString(relInfo.toAppendJoinClause())
					break
				}
			}
		default:
			builder.WriteString(model.fullTabName())
		}
	}
	return builder.String()
}

func genJoinTabRef(model Model) string {
	var builder strings.Builder
	builder.WriteString(model.fullTabName())
	for _, relInfo := range model.Relations() {
		relModel := relInfo.Object.(Model)
		if relModel.getModelStatus()&forJoin == forJoin {
			builder.WriteString(relInfo.toAppendJoinClause())
		}
	}
	return builder.String()
}

func genFullWhereClause(model Model) (string, []interface{}) {
	var where *where
	if model.GetParent() == nil {
		where = where.append(model.getWhere())
		for _, relInfo := range model.Relations() {
			if relInfo.Object.(Model).getModelStatus()&forModelWhere == forModelWhere {
				where = where.append(relInfo.Object.(Model).getWhere())
			}
		}
	} else {
		parent := model.GetParent()
		switch {
		case parent.getModelStatus()&synced == synced:
			for _, relInfo := range parent.Relations() {
				if relInfo.Object.(Model) == model {
					where = where.append(newWhere(and, NewExpr("@ = ?", relInfo.Fields[0]), relInfo.Fields[0].Value()))
					for i, f := range relInfo.Fields[1:] {
						if i%2 == 0 {
							where = where.append(f.(Model).getWhere())
						}
					}
					break
				}
			}
		case parent.getModelStatus()&forModelWhere == forModelWhere:
			for _, relInfo := range parent.Relations() {
				if relInfo.Object.(Model) == model {
					where = where.append(parent.getWhere())
					for i, f := range relInfo.Fields[1:] {
						if i%2 == 0 {
							where = where.append(f.(Model).getWhere())
						}
					}
					break
				}
			}
		default:
			where = where.append(model.getWhere())
		}
		for _, relInfo := range model.Relations() {
			if relInfo.Object.(Model).getModelStatus()&forModelWhere == forModelWhere {
				where = where.append(relInfo.Object.(Model).getWhere())
			}
		}
	}
	cl := make([]string, 0, 8)
	vl := make([]interface{}, 0, 8)
	where.toClause(&cl, &vl)
	if len(cl) == 0 {
		return "", nil
	}
	return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(strings.Join(cl, " "), "AND "), "OR ")), vl
}

func genWhereClause(model Model) (string, []interface{}) {
	var where *where
	if model.GetParent() == nil {
		where = where.append(model.getWhere())
	} else {
		parent := model.GetParent()
		switch {
		case parent.getModelStatus()&synced == synced:
			for _, relInfo := range parent.Relations() {
				if relInfo.Object.(Model) == model {
					where = where.append(newWhere(and, NewExpr("@ = ?", relInfo.Fields[0]), relInfo.Fields[0].Value()))
					for i, f := range relInfo.Fields[1:] {
						if i%2 == 0 {
							where = where.append(f.(Model).getWhere())
						}
					}
					break
				}
			}
		case parent.getModelStatus()&forModelWhere == forModelWhere:
			for _, relInfo := range parent.Relations() {
				if relInfo.Object.(Model) == model {
					where = where.append(parent.getWhere())
					for i, f := range relInfo.Fields[1:] {
						if i%2 == 0 {
							where = where.append(f.(Model).getWhere())
						}
					}
					break
				}
			}
		default:
			where = where.append(model.getWhere())
		}
	}
	cl := make([]string, 0, 8)
	vl := make([]interface{}, 0, 8)
	where.toClause(&cl, &vl)
	if len(cl) == 0 {
		return "", nil
	}
	return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(strings.Join(cl, " "), "AND "), "OR ")), vl
}

func genJoinWhereClause(model Model) (string, []interface{}) {
	where := model.getWhere()
	for _, relInfo := range model.Relations() {
		relModel := relInfo.Object.(Model)
		if relModel.getModelStatus()&forModelWhere == forModelWhere {
			where = where.append(relModel.getWhere())
		}
	}
	cl := make([]string, 0, 8)
	vl := make([]interface{}, 0, 8)
	where.toClause(&cl, &vl)
	if len(cl) == 0 {
		return "", nil
	}
	return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(strings.Join(cl, " "), "AND "), "OR ")), vl
}

func genSimpleWhereClause(model Model) (string, []interface{}) {
	cl := make([]string, 0, 8)
	vl := make([]interface{}, 0, 8)
	model.getWhere().toSimpleClause(&cl, &vl)
	if len(cl) == 0 {
		return "", nil
	}
	return fmt.Sprintf("WHERE %s", strings.TrimPrefix(strings.TrimPrefix(strings.Join(cl, " "), "AND "), "OR ")), vl
}

func genLimitClause(model Model) string {
	limit, offset := model.getLimit()
	if limit == 0 {
		return ""
	}
	return fmt.Sprintf("LIMIT %d, %d", offset, limit)
}

func genSelectClause(model Model) string {
	var builder strings.Builder
	if _, ok := model.(ModelList); ok {
		builder.WriteString("SELECT SQL_CALC_FOUND_ROWS ")
		if model.getModelStatus()&distinct == distinct {
			builder.WriteString("DISTINCT ")
		}
	} else {
		builder.WriteString("SELECT ")
	}
	if model.getModelStatus()&selectAll == selectAll {
		builder.WriteString(fmt.Sprintf("%s.* ", model.fullTabName()))
	} else {
		if selectFields := getSelectFields(model); len(selectFields) == 0 {
			for _, f := range getAllFields(model) {
				builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
			}
		} else {
			for _, f := range selectFields {
				builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
			}
		}
	}
	return strings.TrimSuffix(builder.String(), ", ")
}

func genJoinSelectClause(model Model) string {
	var dist string
	var foundRows string
	var builder strings.Builder
	if _, ok := model.(ModelList); ok {
		foundRows = "SQL_CALC_FOUND_ROWS"
		if model.getModelStatus()&distinct == distinct {
			dist = "DISTINCT"
		}
	}
	if model.getModelStatus()&selectAll == selectAll {
		builder.WriteString(fmt.Sprintf("%s.* ", model.fullTabName()))
	} else {
		if selectFields := getSelectFields(model); len(selectFields) == 0 {
			for _, f := range getAllFields(model) {
				builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
			}
		} else {
			for _, f := range selectFields {
				builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
			}
		}
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.Object.(Model)
		if subModel.getModelStatus()&forJoin == forJoin {
			if subModel.getModelStatus()&distinct == distinct {
				dist = "DISTINCT"
			}
			if subModel.getModelStatus()&selectAll == selectAll {
				builder.WriteString(fmt.Sprintf("%s.*, ", subModel.fullTabName()))
			} else {
				if selectFields := getSelectFields(subModel); len(selectFields) == 0 {
					for _, f := range getAllFields(model) {
						builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
					}
				} else {
					for _, f := range selectFields {
						builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
					}
				}
			}
		}
	}
	return fmt.Sprintf("SELECT %s %s %s", foundRows, dist, strings.Trim(builder.String(), ", "))
}
