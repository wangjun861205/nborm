package nborm

import (
	"database/sql"
	"fmt"
	"strings"
)

// DEBUG 调试标志位
var DEBUG = false

// SetDebug 设置调试标志位
func SetDebug(debug bool) {
	DEBUG = debug
}

func initModel(model Model) {
	if !model.checkStatus(inited) {
		if model.getConList() != nil {
			model.setParent(model.getConList().getParent())
			model.setIndex(model.getConList().getIndex())
		} else {
			model.genAlias()
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

func getFieldsForScan(classModel, instanceModel Model, models *[]Model, fields *[]interface{}) {
	instanceAllFields := getAllFields(instanceModel)
	for _, index := range classModel.getSelectedFieldIndexes() {
		*fields = append(*fields, instanceAllFields[index])
	}
	aggs := instanceModel.getAggs()
	for _, agg := range aggs {
		*fields = append(*fields, agg.getField())
	}
	*models = append(*models, instanceModel)
	for i, rel := range classModel.relations() {
		if rel.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
			if subClassModel, ok := rel.lastModel().(ModelList); ok {
				subInstanceModel := instanceModel.relations()[i].lastModel().(ModelList).NewModel()
				getFieldsForScan(subClassModel, subInstanceModel, models, fields)
			} else {
				if !instanceModel.checkStatus(relInited) {
					instanceModel.InitRel()
				}
				getFieldsForScan(rel.lastModel(), instanceModel.relations()[i].lastModel(), models, fields)
			}
		}
	}
}

func toInsert(field Field, cl *[]string, pl *[]string, vl *[]interface{}) {
	field.mustValid()
	*cl = append(*cl, field.colName())
	*pl = append(*pl, "?")
	*vl = append(*vl, field.value())
}

func joinQueryAndScan(exe Executor, model Model, stmt string, whereValues ...interface{}) error {
	rows, err := exe.Query(stmt, whereValues...)
	if err != nil {
		return err
	}
	for rows.Next() {
		models := make([]Model, 0, 8)
		fields := make([]interface{}, 0, 64)
		if l, ok := model.(ModelList); ok {
			m := l.NewModel()
			getFieldsForScan(l, m, &models, &fields)
		} else {
			getFieldsForScan(model, model, &models, &fields)
		}
		if err := rows.Scan(fields...); err != nil {
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
		l.SetTotal(l.Len())
		if limit, offset := l.getLimit(); limit > 0 {
			l.Slice(offset, offset+limit)
		}
	}
	return nil
}

func queryAndScan(exe Executor, model Model, stmt string, whereValues ...interface{}) error {
	if l, ok := model.(ModelList); ok {
		var needCommit bool
		if ex, ok := exe.(*sql.DB); ok {
			var err error
			if exe, err = ex.Begin(); err != nil {
				return err
			}
			needCommit = true
		}
		rows, err := exe.Query(stmt, whereValues...)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			models := make([]Model, 0, 4)
			fields := make([]interface{}, 0, 32)
			m := l.NewModel()
			getFieldsForScan(l, m, &models, &fields)
			if err := rows.Scan(fields...); err != nil {
				return err
			}
			for _, m := range models {
				if l := m.getConList(); l != nil {
					l.addModelStatus(synced)
				}
				m.addModelStatus(synced)
			}
			l.Collapse()
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
		if needCommit {
			return exe.(*sql.Tx).Commit()
		}
	} else {
		rows, err := exe.Query(stmt, whereValues...)
		if err != nil {
			return err
		}
		defer rows.Close()
		for rows.Next() {
			models := make([]Model, 0, 4)
			fields := make([]interface{}, 0, 32)
			getFieldsForScan(model, model, &models, &fields)
			if err := rows.Scan(fields...); err != nil {
				return err
			}
			for _, m := range models {
				if l := m.getConList(); l != nil {
					l.addModelStatus(synced)
				}
				m.addModelStatus(synced)
			}
			model.Collapse()
		}
		if err := rows.Err(); err != nil {
			return err
		}
	}
	return nil
}

func getUpdates(model Model, updates *exprList) {
	if model.checkStatus(forUpdate) {
		*updates = append(*updates, model.getUpdates()...)
	}
	for _, relInfo := range model.relations() {
		subModel := relInfo.lastModel()
		if subModel.checkStatus(forUpdate | containSubUpdate) {
			getUpdates(subModel, updates)
		}
	}
}

func genUpdateClause(model Model) (string, []interface{}) {
	updates := make(exprList, 0, 16)
	getUpdates(model, &updates)
	return updates.toClause(assignExpr)
}

func genSimpleUpdateClause(model Model) (string, []interface{}) {
	updates := make(exprList, 0, 16)
	getUpdates(model, &updates)
	return updates.toSimpleClause(assignExpr)
}

func getSelectedColumns(model Model, fields *[]interface{}) {
	allFields := getAllFields(model)
	for _, index := range model.getSelectedFieldIndexes() {
		*fields = append(*fields, allFields[index])
	}
	aggs := model.getAggs()
	for _, agg := range aggs {
		*fields = append(*fields, agg)
	}
	for _, relInfo := range model.relations() {
		if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
			getSelectedColumns(relInfo.lastModel(), fields)
		}
	}
}

func genSelectedClause(model Model) string {
	var builder strings.Builder
	builder.WriteString("SELECT ")
	if _, ok := model.(ModelList); ok {
		builder.WriteString("SQL_CALC_FOUND_ROWS ")
	}
	if model.checkStatus(distinct) {
		builder.WriteString("DISTINCT ")
	}
	fields := make([]interface{}, 0, 64)
	getSelectedColumns(model, &fields)
	for _, field := range fields {
		switch f := field.(type) {
		case Field:
			builder.WriteString(f.fullColName())
			builder.WriteString(", ")
		case aggregator:
			clause, _ := f.getExpr().toClause()
			builder.WriteString(clause)
			builder.WriteString(" AS ")
			builder.WriteString(f.getName())
			builder.WriteString(", ")
		default:
			panic(fmt.Errorf("invalid field type (%T)", field))
		}
	}
	return strings.TrimSuffix(builder.String(), ", ")
}

func getSelectFields(model Model) FieldList {
	selectFields := getFields(model, forSelect)
	if len(selectFields) == 0 {
		return getAllFields(model)
	}
	return selectFields
}

func getOrderBys(model Model, orderBys *[]refClauser) {
	*orderBys = append(*orderBys, model.getOrderBys()...)
	for _, relInfo := range model.relations() {
		if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
			getOrderBys(relInfo.lastModel(), orderBys)
		}
	}
}

func genOrderByClause(model Model) string {
	orderBys := make([]refClauser, 0, 8)
	getOrderBys(model, &orderBys)
	if len(orderBys) == 0 {
		return ""
	}
	var builder strings.Builder
	builder.WriteString("ORDER BY ")
	for _, orderBy := range orderBys {
		builder.WriteString(orderBy.toRefClause())
		builder.WriteString(", ")
	}
	return strings.TrimSuffix(builder.String(), ", ")
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
	for _, relInfo := range model.relations() {
		dstModel := relInfo.lastModel()
		switch {
		case dstModel.checkStatus(forJoin):
			*refs = append(*refs, relInfo.toClause(join))
			getTabRef(dstModel, refs)
		case dstModel.checkStatus(forLeftJoin):
			*refs = append(*refs, relInfo.toClause(leftJoin))
			getTabRef(dstModel, refs)
		case dstModel.checkStatus(forRightJoin):
			*refs = append(*refs, relInfo.toClause(rightJoin))
			getTabRef(dstModel, refs)
		}
	}
}

func genTabRefClause(model Model) string {
	refs := make([]string, 0, 4)
	refs = append(refs, model.fullTabName())
	getTabRef(model, &refs)
	return strings.Join(refs, " ")
}

func genBackTabRefClause(model Model) string {
	parent := model.getParent()
	if parent == nil {
		panic("no parent model for back query")
	}
	refs := make([]string, 0, 4)
	refs = append(refs, parent.fullTabName())
	var got bool
	for _, relInfo := range parent.relations() {
		if relInfo.lastModel() == model {
			refs = append(refs, relInfo.toClause(join))
			got = true
			break
		}
	}
	if !got {
		panic("cannot find relation")
	}
	getTabRef(model, &refs)
	return strings.Join(refs, " ")
}

func getWheres(model Model, wheres *exprList) {
	*wheres = append(*wheres, model.getWheres()...)
	for _, relInfo := range model.relations() {
		dstModel := relInfo.lastModel()
		if dstModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			getWheres(dstModel, wheres)
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

func genBackWhereClause(model Model) (string, []interface{}) {
	parent := model.getParent()
	if parent == nil {
		panic("no parent model for back query")
	}
	wheres := make(exprList, 0, 8)
	for _, k := range parent.PrimaryKey() {
		wheres = append(wheres, NewExpr("@ = ?", k, k.value()))
	}
	getWheres(model, &wheres)
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

func getGroupByFields(model Model, groupBys *[]refClauser) {
	for _, g := range model.getGroupBys() {
		*groupBys = append(*groupBys, g)
	}
	for _, relInfo := range model.relations() {
		if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
			getGroupByFields(relInfo.lastModel(), groupBys)
		}
	}
}

func genGroupByClause(model Model) string {
	groupBys := make([]refClauser, 0, 8)
	getGroupByFields(model, &groupBys)
	if len(groupBys) == 0 {
		return ""
	}
	var builder strings.Builder
	builder.WriteString("GROUP BY ")
	for _, g := range groupBys {
		builder.WriteString(g.toRefClause())
		builder.WriteString(", ")
	}
	return strings.TrimSuffix(builder.String(), ", ")
}

func getHavings(model Model, havings *exprList) {
	*havings = append(*havings, model.getHavings()...)
	for _, relInfo := range model.relations() {
		dstModel := relInfo.lastModel()
		if dstModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			getHavings(dstModel, havings)
		}
	}
}

func genHavingClause(model Model) (string, []interface{}) {
	havings := make(exprList, 0, 8)
	getHavings(model, &havings)
	if len(havings) == 0 {
		return "", nil
	}
	return havings.toClause(havingExpr)
}

func genInsertClause(model Model) (string, []interface{}) {
	validFields := getFields(model, valid)
	cl := make([]string, 0, len(validFields))
	vl := make([]interface{}, 0, len(validFields))
	ip := strings.TrimSuffix(strings.Repeat("?, ", len(validFields)), ", ")
	for _, f := range validFields {
		cl = append(cl, f.rawFullColName())
		vl = append(vl, f.value())
	}
	return fmt.Sprintf("(%s) VALUES (%s)", strings.Join(cl, ", "), ip), vl
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
		values = append(values, v.value())
	case []byte:
		values = append(values, fmt.Sprintf("X'%x'", v))
	case string, int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float32, float64:
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
			values = append(values, ev.value())
		}
	default:
		panic("unsupported value type to convert to placeholder")
	}
	return
}

func getDeleteModels(model Model, models *[]string) {
	if model.checkStatus(forDelete) {
		*models = append(*models, model.getAlias())
	}
	for _, relInfo := range model.relations() {
		dstModel := relInfo.lastModel()
		if dstModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			getDeleteModels(dstModel, models)
		}
	}
}

func genDeleteClause(model Model) string {
	deleteModels := make([]string, 0, 4)
	getDeleteModels(model, &deleteModels)
	return fmt.Sprintf("DELETE %s", strings.Join(deleteModels, ", "))
}
