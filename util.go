package nborm

import (
	"bytes"
	"encoding/json"
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

func getFieldsForScan(model Model) ([]interface{}, []Model) {
	addrs := make([]interface{}, 0, 32)
	models := make([]Model, 0, 16)
	selectFields := getSelectFields(model)
	if l, ok := model.(ModelList); ok {
		newModel := l.NewModel()
		allFields := getAllFields(newModel)
	OUTER:
		for _, sf := range selectFields {
			for _, af := range allFields {
				if sf.colName() == af.colName() {
					addrs = append(addrs, af)
					continue OUTER
				}
			}
		}
		models = append(models, newModel)
	} else {
		for _, f := range selectFields {
			addrs = append(addrs, f)
		}
		models = append(models, model)
	}
	for i, relInfo := range model.Relations() {
		subModel := relInfo.DstModel
		if subModel.getModelStatus()&forJoin == forJoin {
			selectFields := getSelectFields(subModel)
			if _, ok := subModel.(ModelList); ok {
				sl := models[0].Relations()[i].DstModel.(ModelList)
				newModel := sl.NewModel()
				allFields := getAllFields(newModel)
			SUB_OUTER:
				for _, sf := range selectFields {
					for _, af := range allFields {
						if sf.colName() == af.colName() {
							addrs = append(addrs, af)
							continue SUB_OUTER
						}
					}
				}
				models = append(models, newModel)
			} else {
				newModel := models[0].Relations()[i].DstModel
				allFields := getAllFields(newModel)
			SUB_OUTER2:
				for _, sf := range selectFields {
					for _, af := range allFields {
						if sf.colName() == af.colName() {
							addrs = append(addrs, af)
							continue SUB_OUTER2
						}
					}
				}
				models = append(models, newModel)
			}
		}
	}
	return addrs, models
}

func toInsert(field Field, cl *[]string, pl *[]string, vl *[]interface{}) {
	field.mustValid()
	*cl = append(*cl, field.colName())
	*pl = append(*pl, "?")
	*vl = append(*vl, field.Value())
}

func getJoinModels(classModel Model, instanceModel Model, models *[]Model, fields *[]Field) {
	if instanceModel == nil {
		if l, ok := classModel.(ModelList); ok {
			m, fs := newModelAndSelectFields(l, l)
			*models = append(*models, m)
			*fields = append(*fields, fs...)
			instanceModel = m
		} else {
			*models = append(*models, classModel)
			*fields = append(*fields, getSelectFields(classModel)...)
			instanceModel = classModel
		}
	}
	if classModel.getModelStatus()&containSubJoin == containSubJoin {
		for i, relInfo := range classModel.Relations() {
			subClassModel := relInfo.DstModel
			subInstanceModel := instanceModel.Relations()[i].DstModel
			if subClassModel.getModelStatus()&forJoin == forJoin {
				if l, ok := subClassModel.(ModelList); ok {
					newSubModel, selectFields := newModelAndSelectFields(l, subInstanceModel.(ModelList))
					*models = append(*models, newSubModel)
					*fields = append(*fields, selectFields...)
					getJoinModels(subClassModel, newSubModel, models, fields)
				} else {
					*models = append(*models, subInstanceModel)
					*fields = append(*fields, getJoinInstanceSelectFields(subClassModel, subInstanceModel)...)
					getJoinModels(subClassModel, subInstanceModel, models, fields)
				}
			}
		}
	}
}

func getJoinSelectFields(model Model, fields *[]Field) {
	*fields = append(*fields, getSelectFields(model)...)
	if model.getModelStatus()&containSubJoin == containSubJoin {
		for _, relInfo := range model.Relations() {
			subModel := relInfo.DstModel
			if subModel.getModelStatus()&forJoin == forJoin {
				getJoinSelectFields(subModel, fields)
			}
		}
	}
}

func genJoinSelectClause(model Model) string {
	fields := make([]Field, 0, 64)
	getJoinSelectFields(model, &fields)
	var builder strings.Builder
	// if _, ok := model.(ModelList); ok {
	// 	builder.WriteString("SELECT SQL_CALC_FOUND_ROWS ")
	// } else {
	builder.WriteString("SELECT ")
	// }
	if model.getModelStatus()&distinct == distinct {
		builder.WriteString("DISTINCT ")
	}
	for _, field := range fields {
		builder.WriteString(fmt.Sprintf("%s, ", field.fullColName()))
	}
	return strings.TrimSuffix(builder.String(), ", ")
}

func joinQueryAndScan(exe Executor, model Model, stmt string, whereValues ...interface{}) error {
	rows, err := exe.Query(stmt, whereValues...)
	if err != nil {
		return err
	}
	for rows.Next() {
		models := make([]Model, 0, 8)
		fields := make([]Field, 0, 64)
		getJoinModels(model, nil, &models, &fields)
		addrs := make([]interface{}, 0, len(fields))
		for _, f := range fields {
			addrs = append(addrs, f)
		}
		if err := rows.Scan(addrs...); err != nil {
			return err
		}
		for _, m := range models {
			m.addModelStatus(synced)
			if conList := m.getConList(); conList != nil {
				conList.addModelStatus(synced)
			}
		}
		model.Collapse()
	}
	if err := rows.Err(); err != nil {
		return err
	}
	if err := rows.Close(); err != nil {
		return err
	}
	if l, ok := model.(ModelList); ok {
		// var foundRows int
		// if err := exe.QueryRow("SELECT FOUND_ROWS()").Scan(&foundRows); err != nil {
		// 	return err
		// }
		// l.SetTotal(foundRows)
		l.SetTotal(l.Len())
	}
	return nil
}

func activeBackQuery(model Model) {
	for _, relInfo := range model.Relations() {
		relInfo.DstModel.addModelStatus(forBackQuery)
	}
}

func queryAndScan(exe Executor, model Model, stmt string, whereValues ...interface{}) error {
	if l, ok := model.(ModelList); ok {
		rows, err := exe.Query(stmt, whereValues...)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			fields, models := getFieldsForScan(model)
			if err := rows.Scan(fields...); err != nil {
				return err
			}
			for _, m := range models {
				m.addModelStatus(synced)
				activeBackQuery(m)
			}
		}
		if err := rows.Err(); err != nil {
			return err
		}
		var rowCount int
		if err := exe.QueryRow(`SELECT FOUND_ROWS()`).Scan(&rowCount); err != nil {
			return err
		}
		l.SetTotal(rowCount)
		if l.Len() > 0 {
			l.addModelStatus(synced)
		}
	} else {
		fields, models := getFieldsForScan(model)
		if err := exe.QueryRow(stmt, whereValues...).Scan(fields...); err != nil {
			return err
		}
		for _, m := range models {
			m.addModelStatus(synced)
			activeBackQuery(m)
		}
	}
	return nil
}

func IsPrimaryKeyEqual(lm, rm Model) bool {
	lpk := lm.PrimaryKey()
	rpk := rm.PrimaryKey()
	for i := 0; i < len(lpk); i++ {
		if !lpk[i].IsValid() || !rpk[i].IsValid() || lpk[i].Value() != rpk[i].Value() {
			return false
		}
	}
	return true
}

func getUpdates(model Model, updates *exprList) {
	if model.checkStatus(forUpdate) {
		*updates = append(*updates, model.getUpdateList()...)
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.DstModel
		if subModel.checkStatus(forUpdate | containSubUpdate) {
			getUpdates(subModel, updates)
		}
	}
}

func genUpdateClause(model Model) (string, []interface{}) {
	updates := make(exprList, 0, 16)
	getUpdates(model, &updates)
	return updates.toClause(updateExpr)
}

func genSimpleUpdateClause(model Model) (string, []interface{}) {
	updates := make(exprList, 0, 16)
	getUpdates(model, &updates)
	return updates.toSimpleClause(updateExpr)
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
		relModel := relInfo.DstModel
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
		if relInfo.DstModel.getModelStatus()&forModelOrder == forModelOrder {
			for _, f := range getFields(relInfo.DstModel, forAscOrder|forDscOrder) {
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

func getTabRef(model Model, refs *[]string) {
	*refs = append(*refs, model.fullTabName())
	if model.checkStatus(forBackQuery) {
		for _, relInfo := range model.GetParent().Relations() {
			if relInfo.DstModel == model {
				*refs = append(*refs, relInfo.toRevAppendJoinClause())
			}
		}
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.DstModel
		if subModel.checkStatus(containWhere | containSubWhere) {
			getTabRef(subModel, refs)
		}
	}
}

func genTabRef(model Model) string {
	refs := make([]string, 0, 8)
	getTabRef(model, &refs)
	return strings.Join(refs, " ")
}

func getJoinTabRef(model Model, refs *[]string) {
	if parent := model.GetParent(); parent != nil {
		for _, refInfo := range parent.Relations() {
			m := refInfo.DstModel
			if m.getAlias() == model.getAlias() {
				*refs = append(*refs, refInfo.toAppendJoinClause())
			}
		}
	} else {
		*refs = append(*refs, model.fullTabName())
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.DstModel
		if subModel.checkStatus(forJoin | containSubJoin | containSubWhere | containWhere) {
			getJoinTabRef(subModel, refs)
		}
	}
}

func genJoinTabRef(model Model) string {
	refs := make([]string, 0, 8)
	getJoinTabRef(model, &refs)
	return strings.Join(refs, " ")
}

func getWheres(model Model, wheres *exprList) {
	if model.checkStatus(forBackQuery) {
		for _, relInfo := range model.GetParent().Relations() {
			if relInfo.DstModel == model {
				*wheres = append(*wheres, relInfo.JoinWheres...)
				for _, pk := range relInfo.SrcModel.PrimaryKey() {
					*wheres = append(*wheres, NewExpr(" AND @ = ?", pk, pk.Value()))
				}
			}
		}
	}
	if model.checkStatus(containWhere) {
		*wheres = append(*wheres, model.getWheres()...)
	}
	if model.checkStatus(containSubWhere) {
		for _, relInfo := range model.Relations() {
			subModel := relInfo.DstModel
			getWheres(subModel, wheres)
		}
	}
}

func genWhereClause(model Model) (string, []interface{}) {
	wheres := make(exprList, 0, 8)
	getWheres(model, &wheres)
	if len(wheres) == 0 {
		return "", nil
	}
	return wheres.toClause(whereExpr)
}

func getJoinWheres(model Model, wheres *exprList) {
	if parent := model.GetParent(); parent != nil {
		for _, relInfo := range model.GetParent().Relations() {
			if relInfo.DstModel == model {
				*wheres = append(*wheres, relInfo.JoinWheres...)
			}
		}
	}
	if model.checkStatus(forJoin | containSubJoin) {
		if model.checkStatus(containJoinWhere) {
			*wheres = append(*wheres, model.getJoinWheres()...)
		}
	}
	if model.checkStatus(containWhere) {
		*wheres = append(*wheres, model.getWheres()...)
	}
	if model.checkStatus(containSubWhere | containSubJoin) {
		for _, relInfo := range model.Relations() {
			getJoinWheres(relInfo.DstModel, wheres)
		}
	}
}

func genJoinWhereClause(model Model) (string, []interface{}) {
	wheres := make(exprList, 0, 8)
	getJoinWheres(model, &wheres)
	if len(wheres) == 0 {
		return "", nil
	}
	return wheres.toClause(whereExpr)
}

func genSimpleWhereClause(model Model) (string, []interface{}) {
	wheres := model.getWheres()
	if len(wheres) == 0 {
		return "", nil
	}
	return wheres.toSimpleClause(whereExpr)
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
	var foundRows string
	var dist string
	parent := model.GetParent()
	if parent != nil && parent.getModelStatus()&forJoin == forJoin {
		if _, ok := parent.(ModelList); ok {
			foundRows = "SQL_CALC_FOUND_ROWS"
		}
		if parent.getModelStatus()&distinct == distinct {
			dist = "DISTINCT"
		}
		if parent.getModelStatus()&selectAll == selectAll {
			builder.WriteString(fmt.Sprintf("%s.*, ", parent.fullTabName()))
		} else {
			for _, f := range getSelectFields(parent) {
				builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
			}
		}
	}
	if _, ok := model.(ModelList); ok {
		foundRows = "SQL_CALC_FOUND_ROWS"
	}
	if model.getModelStatus()&distinct == distinct {
		dist = "DISTINCT"
	}
	if model.getModelStatus()&selectAll == selectAll {
		builder.WriteString(fmt.Sprintf("%s.*, ", model.fullTabName()))
	} else {
		for _, f := range getSelectFields(model) {
			builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
		}
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.DstModel
		if subModel.getModelStatus()&forJoin == forJoin {
			if subModel.getModelStatus()&selectAll == selectAll {
				builder.WriteString(fmt.Sprintf("%s.*, ", subModel.fullTabName()))
			} else {
				for _, f := range getSelectFields(subModel) {
					builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
				}
			}
		}
	}
	return fmt.Sprintf("SELECT %s %s %s", foundRows, dist, strings.TrimSuffix(builder.String(), ", "))
}

func genAggSelectClause(model Model) (string, FieldList) {
	var builder strings.Builder
	var foundRows string
	var dist string
	fieldList := make(FieldList, 0, 16)
	parent := model.GetParent()
	if parent != nil {
		if _, ok := parent.(ModelList); ok {
			foundRows = "SQL_CALC_FOUND_ROWS"
		}
		if parent.getModelStatus()&distinct == distinct {
			dist = "DISTINCT"
		}
		for _, f := range getFields(parent, forSelect) {
			builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
			fieldList = append(fieldList, f.dup())
		}
		if parent.getModelStatus()&forModelAgg == forModelAgg {
			for _, exp := range parent.getAggExps() {
				c, _ := exp.expr.toClause()
				builder.WriteString(fmt.Sprintf("%s, ", c))
				fieldList = append(fieldList, exp.field.dup())
			}
		}
	}
	if _, ok := model.(ModelList); ok {
		foundRows = "SQL_CALC_FOUND_ROWS"
	}
	if model.getModelStatus()&distinct == distinct {
		dist = "DISTINCT"
	}
	for _, f := range getFields(model, forSelect) {
		builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
		fieldList = append(fieldList, f.dup())
	}
	if model.getModelStatus()&forModelAgg == forModelAgg {
		for _, exp := range model.getAggExps() {
			c, _ := exp.expr.toClause()
			builder.WriteString(fmt.Sprintf("%s, ", c))
			fieldList = append(fieldList, exp.field.dup())
		}
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.DstModel
		for _, f := range getFields(subModel, forSelect) {
			builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
			fieldList = append(fieldList, f.dup())
		}
		if subModel.getModelStatus()&forModelAgg == forModelAgg {
			for _, exp := range subModel.getAggExps() {
				c, _ := exp.expr.toClause()
				builder.WriteString(fmt.Sprintf("%s, ", c))
				fieldList = append(fieldList, exp.field.dup())
			}
		}
	}
	return fmt.Sprintf("SELECT %s %s %s", foundRows, dist, strings.TrimSuffix(builder.String(), ", ")), fieldList
}

func genGroupByClause(model Model) string {
	l := make([]string, 0, 16)
	parent := model.GetParent()
	if parent != nil {
		for _, f := range getFields(parent, forGroup) {
			l = append(l, fmt.Sprintf("%s", f.fullColName()))
		}
	}
	for _, f := range getFields(model, forGroup) {
		l = append(l, fmt.Sprintf("%s", f.fullColName()))
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.DstModel
		for _, f := range getFields(subModel, forGroup) {
			l = append(l, fmt.Sprintf("%s", f.fullColName()))
		}
	}
	if len(l) == 0 {
		return ""
	}
	return fmt.Sprintf("GROUP BY %s", strings.Join(l, ", "))

}

func genHavingClause(model Model) (string, []interface{}) {
	havings := make(exprList, 0, 8)
	parent := model.GetParent()
	if parent != nil && parent.getModelStatus()&forModelHaving == forModelHaving {
		havings = append(havings, parent.getHavings()...)
	}
	if model.getModelStatus()&forModelHaving == forModelHaving {
		havings = append(havings, model.getHavings()...)
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.DstModel
		if subModel.getModelStatus()&forModelHaving == forModelHaving {
			havings = append(havings, subModel.getHavings()...)
		}
	}
	if len(havings) == 0 {
		return "", nil
	}
	return havings.toClause(whereExpr)
}

func genInsertClause(model Model) (string, []interface{}) {
	validFields := getFields(model, valid)
	cl := make([]string, 0, len(validFields))
	vl := make([]interface{}, 0, len(validFields))
	ip := strings.TrimSuffix(strings.Repeat("?, ", len(validFields)), ", ")
	for _, f := range validFields {
		cl = append(cl, f.rawFullColName())
		vl = append(vl, f.Value())
	}
	return fmt.Sprintf("(%s) VALUES (%s)", strings.Join(cl, ", "), ip), vl
}

func newModelAndSelectFields(classList ModelList, instanceList ModelList) (Model, []Field) {
	newModel := instanceList.NewModel()
	classSelectFields := getSelectFields(classList)
	instanceSelectFields := make([]Field, 0, len(classSelectFields))
OUTER:
	for _, lsf := range classSelectFields {
		for _, mf := range getAllFields(newModel) {
			if lsf.colName() == mf.colName() {
				instanceSelectFields = append(instanceSelectFields, mf)
				continue OUTER
			}
		}
	}
	return newModel, instanceSelectFields
}

func getJoinInstanceSelectFields(classModel, instanceModel Model) (selectFields []Field) {
	classSelectFields := getSelectFields(classModel)
	instanceFields := getAllFields(instanceModel)
OUTER:
	for _, cf := range classSelectFields {
		for _, inf := range instanceFields {
			if cf.colName() == inf.colName() {
				selectFields = append(selectFields, inf)
				continue OUTER
			}
		}
	}
	return
}

func searchRelation(parent Model, child Model) RelationInfo {
	for _, relInfo := range parent.Relations() {
		subModel := relInfo.DstModel
		if subModel.getAlias() == child.getAlias() {
			return relInfo
		}
	}
	panic(fmt.Errorf("cannot find relation (parent: %s, child: %s)", parent.fullTabName(), child.fullTabName()))
}

func getSelectedFieldsForCount(model Model, fieldList *FieldList) {
	if model.GetParent() == nil || model.checkStatus(forJoin) {
		selectedFields := getFields(model, forSelect)
		*fieldList = append(*fieldList, selectedFields...)
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.DstModel
		if subModel.checkStatus(forJoin | containSubJoin) {
			getSelectedFieldsForCount(subModel, fieldList)
		}
	}
}

func genCountClause(model Model) string {
	selectedFields := make(FieldList, 0, 8)
	getSelectedFieldsForCount(model, &selectedFields)
	if len(selectedFields) == 0 {
		return "SELECT COUNT(*)"
	}
	var builder strings.Builder
	for _, f := range selectedFields {
		builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
	}
	return fmt.Sprintf("SELECT COUNT(DISTINCT CONCAT(%s))", strings.TrimSuffix(builder.String(), ", "))
}

func marshalModel(model Model, bs *[]byte) {
	if l, ok := model.(ModelList); ok {
		*bs = append(*bs, []byte(`{ "List": [`)...)
		for _, m := range l.GetList() {
			marshalModel(m, bs)
			*bs = append(*bs, []byte(`, `)...)
		}
		*bs = bytes.TrimSuffix(*bs, []byte(", "))
		*bs = append(*bs, []byte(fmt.Sprintf(`], "Total": %d }`, l.GetTotal()))...)
	} else {
		if model.checkStatus(synced | containValue) {
			*bs = append(*bs, []byte(`{`)...)
			for _, f := range getFields(model, valid) {
				*bs = append(*bs, []byte(fmt.Sprintf(`"%s": `, f.fieldName()))...)
				b, _ := json.Marshal(f)
				*bs = append(*bs, b...)
				*bs = append(*bs, []byte(`, `)...)
			}
			for _, relInfo := range model.Relations() {
				subModel := relInfo.DstModel
				if subModel.checkStatus(synced | containValue) {
					*bs = append(*bs, []byte(fmt.Sprintf(`"%s":`, relInfo.Name))...)
					marshalModel(subModel, bs)
				} else {
					*bs = append(*bs, []byte(fmt.Sprintf(`"%s":`, relInfo.Name))...)
					if _, ok := subModel.(ModelList); ok {
						*bs = append(*bs, []byte("[], ")...)
					} else {
						*bs = append(*bs, []byte("null, ")...)
					}
				}
			}
			*bs = bytes.TrimSuffix(*bs, []byte(", "))
			*bs = append(*bs, []byte("}")...)
		} else {
			*bs = append(*bs, []byte("null, ")...)
		}
	}
}

func MarshalModel(model Model) []byte {
	bs := make([]byte, 0, 1024)
	marshalModel(model, &bs)
	return bs
}

func UnmarshalModel(bs []byte, model Model) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(bs, &m); err != nil {
		return err
	}
	if l, ok := model.(ModelList); ok {
		list := m["List"].([]interface{})
		for _, obj := range list {
			mm := obj.(map[string]interface{})
			newModel := l.NewModel()
			bytes := []byte(strings.Replace(fmt.Sprintf("%#v", mm), "map[string]interface {}", "", -1))
			if err := UnmarshalModel(bytes, newModel); err != nil {
				return err
			}
		}
		total := int(m["Total"].(float64))
		l.SetTotal(total)
		l.addModelStatus(containValue)
	} else {
		fields := getAllFields(model)
		for _, f := range fields {
			if obj, ok := m[f.fieldName()]; ok {
				bytes := []byte(fmt.Sprintf("%#v", obj))
				if err := json.Unmarshal(bytes, f); err != nil {
					return err
				}
				f.addStatus(valid)
			}
		}
		for _, relInfo := range model.Relations() {
			if obj, ok := m[relInfo.Name]; ok {
				if l, ok := relInfo.DstModel.(ModelList); ok {
					bytes := []byte(fmt.Sprintf("%#v", obj))
					if err := UnmarshalModel(bytes, l); err != nil {
						return err
					}
				} else {
					bytes := []byte(strings.Replace(fmt.Sprintf("%#v", obj), "map[string]interface {}", "", -1))
					if err := UnmarshalModel(bytes, relInfo.DstModel); err != nil {
						return err
					}
				}
			}
		}
		model.addModelStatus(containValue)
	}
	return nil
}

func valToPlaceholder(val interface{}) string {
	switch v := val.(type) {
	case string, []byte, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64, Field:
		return "?"
	case []string:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []int:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []uint:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []int8:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []int16:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []uint16:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []int32:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []uint32:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []int64:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []uint64:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []float32:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []float64:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	case []Field:
		return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(v)), ", "))
	default:
		panic("unsupported value type to convert to placeholder")
	}
}

func expandArg(val interface{}) (values []interface{}) {
	switch v := val.(type) {
	case Field:
		values = append(values, v.Value())
	case string, []byte, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
		values = append(values, v)
	case []string:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []int:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []uint:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []int8:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []int16:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []uint16:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []int32:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []uint32:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []int64:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []uint64:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []float32:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []float64:
		for _, ev := range v {
			values = append(values, ev)
		}
	case []Field:
		for _, ev := range v {
			values = append(values, ev.Value())
		}
	default:
		panic("unsupported value type to convert to placeholder")
	}
	return
}
