package nborm

import (
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
	parent := model.GetParent()
	if parent != nil && parent.getModelStatus()&forJoin == forJoin {
		if selectFields := getSelectFields(parent); len(selectFields) == 0 {
			if l, ok := parent.(ModelList); ok {
				newModel := l.NewModel()
				for _, f := range getAllFields(newModel) {
					addrs = append(addrs, f)
				}
				models = append(models, newModel)
			} else {
				for _, f := range getAllFields(parent) {
					addrs = append(addrs, f)
				}
				models = append(models, parent)
			}
		} else {
			if l, ok := parent.(ModelList); ok {
				newModel := l.NewModel()
				allFields := getAllFields(newModel)
			PARENT_OUTER:
				for _, sf := range selectFields {
					for _, af := range allFields {
						if sf.colName() == af.colName() {
							addrs = append(addrs, af)
							continue PARENT_OUTER
						}
					}
				}
				models = append(models, newModel)
			} else {
				for _, f := range selectFields {
					addrs = append(addrs, f)
				}
				models = append(models, parent)
			}
		}
	}
	if selectFields := getSelectFields(model); len(selectFields) == 0 {
		if l, ok := model.(ModelList); ok {
			newModel := l.NewModel()
			for _, f := range getAllFields(newModel) {
				addrs = append(addrs, f)
			}
			models = append(models, newModel)
		} else {
			for _, f := range getAllFields(model) {
				addrs = append(addrs, f)
			}
			models = append(models, model)
		}
	} else {
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
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.Object.(Model)
		if subModel.getModelStatus()&forJoin == forJoin {
			if selectFields := getSelectFields(subModel); len(selectFields) == 0 {
				if l, ok := subModel.(ModelList); ok {
					newModel := l.NewModel()
					for _, f := range getAllFields(newModel) {
						addrs = append(addrs, f)
					}
					models = append(models, newModel)
				} else {
					for _, f := range getAllFields(subModel) {
						addrs = append(addrs, f)
					}
					models = append(models, subModel)
				}
			} else {
				if l, ok := subModel.(ModelList); ok {
					newModel := l.NewModel()
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
					for _, f := range selectFields {
						addrs = append(addrs, f)
					}
					models = append(models, subModel)
				}
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

func queryAndScan(exe Executor, model Model, stmt string, whereValues ...interface{}) error {
	collFuncs := make([]func(), 0, 2)
	parent := model.GetParent()
	if parent != nil && parent.getModelStatus()&forJoin == forJoin {
		if l, ok := parent.(ModelList); ok {
			collFunc := func() { l.Collapse() }
			collFuncs = append(collFuncs, collFunc)
		}
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.Object.(Model)
		if subModel.getModelStatus()&forJoin == forJoin {
			if l, ok := model.(ModelList); ok {
				collFunc := func() { l.Collapse() }
				collFuncs = append(collFuncs, collFunc)
				break
			}
		}
	}
	switch {
	case parent != nil && parent.getModelStatus()&forJoin == forJoin:
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
			for i := len(collFuncs); i > 0; i-- {
				collFuncs[i-1]()
			}
			for _, m := range models {
				m.addModelStatus(synced)
			}
		}
		if err := rows.Err(); err != nil {
			return err
		}
		var rowCount int
		if err := exe.QueryRow(`SELECT FOUND_ROWS()`).Scan(&rowCount); err != nil {
			return err
		}
		if l, ok := parent.(ModelList); ok {
			l.SetTotal(rowCount)
		}
	default:
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
				for i := len(collFuncs); i > 0; i-- {
					collFuncs[i-1]()
				}
				for _, m := range models {
					m.addModelStatus(synced)
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
		} else {
			if len(collFuncs) > 0 {
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
					for i := len(collFuncs); i > 0; i-- {
						collFuncs[i-1]()
					}
					for _, m := range models {
						m.addModelStatus(synced)
					}
				}
				if err := rows.Err(); err != nil {
					return err
				}
			} else {
				fields, models := getFieldsForScan(model)
				if err := exe.QueryRow(stmt, whereValues...).Scan(fields...); err != nil {
					return err
				}
				for _, m := range models {
					m.addModelStatus(synced)
				}
			}
		}
	}
	return nil
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
	for _, relInfo := range model.Relations() {
		subModel := relInfo.Object.(Model)
		if subModel.getModelStatus()&forModelUpdate == forModelUpdate {
			for _, f := range getFields(subModel, forUpdate) {
				updates = append(updates, f.updateSet())
			}
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

func genTabRef(model Model) string {
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
			for _, relInfo := range model.Relations() {
				if relInfo.Object.(Model).getModelStatus()&forModelRef == forModelRef {
					builder.WriteString(relInfo.toAppendJoinClause())
				}
			}
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
		for _, relInfo := range model.Relations() {
			subModel := relInfo.Object.(Model)
			if subModel.getModelStatus()&forModelWhere == forModelWhere {
				for i, f := range relInfo.Fields[1:] {
					if i%2 == 0 {
						where = where.append(f.(Model).getWhere())
					}
				}
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
			for _, relInfo := range model.Relations() {
				subModel := relInfo.Object.(Model)
				if subModel.getModelStatus()&forModelWhere == forModelWhere {
					for i, f := range relInfo.Fields[1:] {
						if i%2 == 0 {
							where = where.append(f.(Model).getWhere())
						}
					}
				}
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
		subModel := relInfo.Object.(Model)
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
				builder.WriteString(fmt.Sprintf("%s, ", exp.expr.String()))
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
			builder.WriteString(fmt.Sprintf("%s, ", exp.expr.String()))
			fieldList = append(fieldList, exp.field.dup())
		}
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.Object.(Model)
		for _, f := range getFields(subModel, forSelect) {
			builder.WriteString(fmt.Sprintf("%s, ", f.fullColName()))
			fieldList = append(fieldList, f.dup())
		}
		if subModel.getModelStatus()&forModelAgg == forModelAgg {
			for _, exp := range subModel.getAggExps() {
				builder.WriteString(fmt.Sprintf("%s, ", exp.expr.String()))
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
		subModel := relInfo.Object.(Model)
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
	var having *where
	parent := model.GetParent()
	if parent != nil && parent.getModelStatus()&forModelHaving == forModelHaving {
		having = having.append(parent.getHaving())
	}
	if model.getModelStatus()&forModelHaving == forModelHaving {
		having = having.append(model.getHaving())
	}
	for _, relInfo := range model.Relations() {
		subModel := relInfo.Object.(Model)
		if subModel.getModelStatus()&forModelHaving == forModelHaving {
			having = having.append(subModel.getHaving())
		}
	}
	cl := make([]string, 0, 8)
	vl := make([]interface{}, 0, 8)
	having.toClause(&cl, &vl)
	if len(cl) == 0 {
		return "", nil
	}
	return fmt.Sprintf("HAVING %s", strings.TrimPrefix(strings.TrimPrefix(strings.Join(cl, " "), "AND "), "OR ")), vl
}
