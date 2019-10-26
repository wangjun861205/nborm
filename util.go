package nborm

import (
	"database/sql"
	"fmt"
	"io"
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
			model.genAlias()
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

func getFieldsForScan(classModel, instanceModel Model, models *[]Model, selectors *[]interface{}) {
	for _, sel := range classModel.getSelectors().list {
		sel.toScan(instanceModel, selectors)
	}
	*models = append(*models, instanceModel)
	for i, rel := range classModel.relations() {
		if rel.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
			if subClassModel, ok := rel.lastModel().(ModelList); ok {
				subInstanceModel := instanceModel.relations()[i].lastModel().(ModelList).NewModel()
				getFieldsForScan(subClassModel, subInstanceModel, models, selectors)
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
			selectors := make([]interface{}, 0, 32)
			m := l.NewModel()
			getFieldsForScan(l, m, &models, &selectors)
			if err := rows.Scan(selectors...); err != nil {
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
			selectors := make([]interface{}, 0, 32)
			getFieldsForScan(model, model, &models, &selectors)
			if err := rows.Scan(selectors...); err != nil {
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

func genUpdateClause(model Model, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	model.getUpdates().toClause(w, vals, isFirstGroup, isFirstNode)
	for _, relInfo := range model.relations() {
		subModel := relInfo.lastModel()
		if subModel.checkStatus(forUpdate | containSubUpdate) {
			genUpdateClause(subModel, w, vals, isFirstGroup, isFirstNode)
		}
	}
}

func genSimpleUpdateClause(model Model, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	model.getUpdates().toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	for _, relInfo := range model.relations() {
		subModel := relInfo.lastModel()
		if subModel.checkStatus(forUpdate | containSubUpdate) {
			genSimpleUpdateClause(subModel, w, vals, isFirstGroup, isFirstNode)
		}
	}
}

func getSelectFields(model Model) FieldList {
	selectFields := getFields(model, forSelect)
	if len(selectFields) == 0 {
		return getAllFields(model)
	}
	return selectFields
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

func genLimitClause(model Model, w io.Writer, vals *[]interface{}) {
	limit, offset := model.getLimit()
	if limit == 0 {
		return
	}
	w.Write([]byte(fmt.Sprintf("LIMIT %d, %d ", offset, limit)))
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

func genDeleteClause(model Model, w io.Writer, vals *[]interface{}, isFirstNode *bool) {
	if model.checkStatus(forDelete) {
		if *isFirstNode {
			*isFirstNode = false
			w.Write([]byte("DELETE "))
		} else {
			w.Write([]byte(", "))
		}
		w.Write([]byte(model.getAlias()))
		w.Write([]byte(" "))
	}
	for _, relInfo := range model.relations() {
		dstModel := relInfo.lastModel()
		if dstModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			genDeleteClause(dstModel, w, vals, isFirstNode)
		}
	}
}
