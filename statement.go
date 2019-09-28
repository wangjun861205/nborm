package nborm

import (
	"io"
)

func genInsertStmt(model Model, w io.Writer, vals *[]interface{}) {
	w.Write([]byte("INSERT INTO "))
	w.Write([]byte(model.rawFullTabName()))
	w.Write([]byte(" SET "))
	inserts := model.getInserts()
	inserts.toSimpleClause(assignExpr, w, vals, true, true)
}

func genInsertOrUpdateStmt(model Model, w io.Writer, vals *[]interface{}) {
	w.Write([]byte("INSERT INTO "))
	w.Write([]byte(model.rawFullTabName()))
	w.Write([]byte(" SET "))
	model.getInserts().toSimpleClause(assignExpr, w, vals, true, true)
	w.Write([]byte(" ON DUPLICATE KEY UPDATE "))
	model.getUpdates().toSimpleClause(assignExpr, w, vals, true, true)
}

func genSelectStmt(model Model, w io.Writer, vals *[]interface{}) {
	genSelectedClause(model, w, vals, true)
	genTabRefClause(model, w, vals, true)
	genWhereClause(model, w, vals, true, true)
	genGroupByClause(model, w, vals, true)
	genHavingClause(model, w, vals, true)
	genOrderByClause(model, w, vals, true)
	genLimitClause(model, w, vals)
	if model.checkStatus(selectForUpdate) {
		w.Write([]byte("FOR UPDATE "))
	}
}

func genBackQueryStmt(model Model, w io.Writer, vals *[]interface{}) {
	genSelectedClause(model, w, vals, true)
	genBackTabRefClause(model, w, vals)
	genBackWhereClause(model, w, vals)
	genGroupByClause(model, w, vals, true)
	genHavingClause(model, w, vals, true)
	genOrderByClause(model, w, vals, true)
	genLimitClause(model, w, vals)
	if model.checkStatus(selectForUpdate) {
		w.Write([]byte("FOR UPDATE "))
	}
}

func genUpdateStmt(model Model, w io.Writer, vals *[]interface{}) {
	genTabRefClause(model, w, vals, true)
	genUpdateClause(model, w, vals, true)
	genWhereClause(model, w, vals, true, true)
}

func genDeleteStmt(model Model, w io.Writer, vals *[]interface{}) {
	genDeleteClause(model, w, vals, true)
	genTabRefClause(model, w, vals, true)
	genWhereClause(model, w, vals, true, true)
	genOrderByClause(model, w, vals, true)
	genLimitClause(model, w, vals)
}

func genSelectedClause(model Model, w io.Writer, vals *[]interface{}, isFirst bool) {
	fieldInfos := model.FieldInfos()
	fieldLength := len(model.getSelectedFieldIndexes())
	for i, idx := range model.getSelectedFieldIndexes() {
		if isFirst {
			isFirst = false
			w.Write([]byte("SELECT "))
			if _, ok := model.(ModelList); ok {
				w.Write([]byte("SQL_CALC_FOUND_ROWS "))
			}
			if model.checkStatus(distinct) {
				w.Write([]byte("DISTINCT "))
			}
		} else if i == 0 {
			w.Write([]byte(", "))
		}
		fieldInfos[idx].Field.toRefClause(w, vals, false, false)
		if i != (fieldLength - 1) {
			w.Write([]byte(", "))
		}
	}
	aggs := model.getAggs()
	aggLength := len(aggs)
	for i, agg := range model.getAggs() {
		if isFirst {
			isFirst = false
			w.Write([]byte("SELECT "))
			if _, ok := model.(ModelList); ok {
				w.Write([]byte("SQL_CALC_FOUND_ROWS "))
			}
			if model.checkStatus(distinct) {
				w.Write([]byte("DISTINCT "))
			}
		} else if i == 0 {
			w.Write([]byte(", "))
		}
		agg.toClause(w, vals, false, false)
		if i != (aggLength - 1) {
			w.Write([]byte(", "))
		}
	}
	for _, relInfo := range model.relations() {
		if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
			genSelectedClause(relInfo.lastModel(), w, vals, isFirst)
		}
	}
}

func genWhereClause(model Model, w io.Writer, vals *[]interface{}, isFirstGroup bool, isFirstNode bool) {
	model.getWheres().toClause(w, vals, isFirstGroup, isFirstNode)
	for _, relInfo := range model.relations() {
		dstModel := relInfo.lastModel()
		if dstModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			genWhereClause(dstModel, w, vals, isFirstGroup, isFirstNode)
		}
	}
}

// func genSimpleWhereClause(model Model, w io.Writer, vals *[]interface{}) {
// 	wheres := model.getWheres()
// 	if len(wheres) == 0 {
// 		return
// 	}
// 	for _, where := range model.getWheres() {
// 		w.Write([]byte("WHERE "))
// 		where.toClause(w, vals)
// 	}
// }

func genBackWhereClause(model Model, w io.Writer, vals *[]interface{}) {
	parent := model.getParent()
	if parent == nil {
		panic("no parent model for back query")
	}
	for _, k := range parent.PrimaryKey() {
		w.Write([]byte("WHERE "))
		NewExpr("@ = ?", k, k.value()).toClause(w, vals, false, false)

	}
	genWhereClause(model, w, vals, false, false)
}
