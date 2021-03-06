package nborm

import (
	"database/sql"
	"fmt"
	"io"
	"reflect"
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
				if !instanceModel.checkStatus(relInited) {
					instanceModel.InitRel()
				}
				// 尝试修复即使没有要Scan的字段仍然在List中创建Model的问题
				if len(subClassModel.getSelectors().list) > 0 {
					subInstanceModel := instanceModel.relations()[i].lastModel().(ModelList).NewModel()
					getFieldsForScan(subClassModel, subInstanceModel, models, selectors)
				}
			} else {
				if !instanceModel.checkStatus(relInited) {
					instanceModel.InitRel()
				}
				// 尝试修复即使没有要Scan的字段仍然将Model放入要Scan的Model列表中的问题
				if len(rel.lastModel().getSelectors().list) > 0 {
					getFieldsForScan(rel.lastModel(), instanceModel.relations()[i].lastModel(), models, selectors)
				}
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
		if subModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			genUpdateClause(subModel, w, vals, isFirstGroup, isFirstNode)
		}
	}
}

func genBulkUpdateClause(model Model, w io.Writer, vals *[][]interface{}, isFirstGroup, isFirstNode *bool) {
	for _, fieldInfo := range model.FieldInfos() {
		if fieldInfo.Field.checkFieldStatus(forBulkUpdate) {
			if *isFirstNode {
				*isFirstNode = false
				w.Write([]byte("SET "))
				fieldInfo.Field.toRefClause(w, nil, isFirstGroup, isFirstNode)
				w.Write([]byte("= ? "))
			} else {
				w.Write([]byte(", "))
				fieldInfo.Field.toRefClause(w, nil, isFirstGroup, isFirstNode)
				w.Write([]byte("= ? "))
			}
		}
	}
	for _, relInfo := range model.relations() {
		subModel := relInfo.lastModel()
		if subModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			genBulkUpdateClause(subModel, w, vals, isFirstGroup, isFirstNode)
		}
	}
}

func getBulkUpdateValues(model Model, vals *[][]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstGroup {
		*isFirstGroup = false
		*isFirstNode = true
		*vals = append(*vals, make([]interface{}, 0, 8))
	}
	if l, ok := model.(ModelList); ok {
		for _, m := range l.GetList() {
			getBulkUpdateValues(m, vals, isFirstGroup, isFirstNode)
		}
	} else {
		if l := model.getConList(); l != nil {
			if *isFirstNode {
				for i, fieldInfo := range l.FieldInfos() {
					if fieldInfo.Field.checkFieldStatus(forBulkUpdate) {
						(*vals)[len(*vals)-1] = append((*vals)[len(*vals)-1], model.FieldInfos()[i].Field.getUpdateValue())
					}
				}
			} else {
				var count int
				for parent := l.getParent(); parent != nil; parent = parent.getParent() {
					for _, fieldInfo := range parent.FieldInfos() {
						if fieldInfo.Field.checkFieldStatus(forBulkUpdate) {
							count++
						}
					}
				}
				prevArgs := make([]interface{}, count)
				copy(prevArgs, (*vals)[len(*vals)-1][:count])
				*vals = append(*vals, prevArgs)
				for i, fieldInfo := range l.FieldInfos() {
					if fieldInfo.Field.checkFieldStatus(forBulkUpdate) {
						(*vals)[len(*vals)-1] = append((*vals)[len(*vals)-1], model.FieldInfos()[i].Field.getUpdateValue())
					}
				}
			}
			for _, relInfo := range model.relations() {
				if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
					getBulkUpdateValues(relInfo.lastModel(), vals, isFirstGroup, isFirstNode)
				}
			}
		} else {
			for _, fieldInfo := range model.FieldInfos() {
				if fieldInfo.Field.checkFieldStatus(forBulkUpdate) {
					(*vals)[len(*vals)-1] = append((*vals)[len(*vals)-1], fieldInfo.Field.getUpdateValue())
				}
			}
			for _, relInfo := range model.relations() {
				if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
					getBulkUpdateValues(relInfo.lastModel(), vals, isFirstGroup, isFirstNode)
				}
			}
		}
		*isFirstGroup = true
		*isFirstNode = false
	}
}

func getBulkWhereValues(model Model, vals *[][]interface{}, isFirstGroup, isFirstNode *bool) {
	if *isFirstGroup {
		*isFirstGroup = false
		*isFirstNode = true
		*vals = append(*vals, make([]interface{}, 0, 8))
	}
	if l, ok := model.(ModelList); ok {
		for _, m := range l.GetList() {
			getBulkWhereValues(m, vals, isFirstGroup, isFirstNode)
		}
	} else {
		if l := model.getConList(); l != nil {
			if *isFirstNode {
				for i, fieldInfo := range l.FieldInfos() {
					if fieldInfo.Field.checkFieldStatus(forBulkWhere) {
						(*vals)[len(*vals)-1] = append((*vals)[len(*vals)-1], model.FieldInfos()[i].Field.getBulkWhereValues()...)
					}
				}
			} else {
				var count int
				for parent := l.getParent(); parent != nil; parent = parent.getParent() {
					for _, fieldInfo := range parent.FieldInfos() {
						if fieldInfo.Field.checkFieldStatus(forBulkUpdate) {
							count++
						}
					}
				}
				prevArgs := make([]interface{}, count)
				copy(prevArgs, (*vals)[len(*vals)-1][:count])
				*vals = append(*vals, prevArgs)
				for i, fieldInfo := range l.FieldInfos() {
					if fieldInfo.Field.checkFieldStatus(forBulkUpdate) {
						(*vals)[len(*vals)-1] = append((*vals)[len(*vals)-1], model.FieldInfos()[i].Field.getBulkWhereValues()...)
					}
				}
			}
			for _, relInfo := range model.relations() {
				if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
					getBulkUpdateValues(relInfo.lastModel(), vals, isFirstGroup, isFirstNode)
				}
			}
		} else {
			for _, fieldInfo := range model.FieldInfos() {
				if fieldInfo.Field.checkFieldStatus(forBulkUpdate) {
					(*vals)[len(*vals)-1] = append((*vals)[len(*vals)-1], fieldInfo.Field.getBulkWhereValues()...)
				}
			}
			for _, relInfo := range model.relations() {
				if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
					getBulkUpdateValues(relInfo.lastModel(), vals, isFirstGroup, isFirstNode)
				}
			}
		}
		*isFirstGroup = true
		*isFirstNode = false
	}
}

func genSimpleUpdateClause(model Model, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	model.getUpdates().toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	for _, relInfo := range model.relations() {
		subModel := relInfo.lastModel()
		if subModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
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
		// panic("unsupported value type to convert to placeholder")
		refVal := reflect.ValueOf(val)
		switch refVal.Kind() {
		case reflect.Int, reflect.Uint, reflect.Int8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			return "?"
		case reflect.Array:
			return strings.TrimSuffix(strings.Repeat("?, ", refVal.Len()), ", ")
		case reflect.Slice:
			return fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", refVal.Len()), ", "))
		default:
			panic("unsupported value type to convert to placeholder")
		}
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
		// panic("unsupported value type to convert to placeholder")
		refVal := reflect.ValueOf(val)
		switch refVal.Kind() {
		case reflect.Int, reflect.Uint, reflect.Int8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32, reflect.Int64, reflect.Uint64,
			reflect.Float32, reflect.Float64:
			values = append(values, refVal.Interface())
		case reflect.Array:
			for i := 0; i < refVal.Len(); i++ {
				values = append(values, refVal.Index(i).Interface())
			}
		case reflect.Slice:
			for i := 0; i < refVal.Len(); i++ {
				values = append(values, refVal.Index(i).Interface())
			}
		default:
			panic("unsupported value type to convert to placeholder")
		}
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
