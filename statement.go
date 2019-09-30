package nborm

import (
	"io"
)

func genInsertStmt(model Model, w io.Writer, vals *[]interface{}) {
	w.Write([]byte("INSERT INTO "))
	w.Write([]byte(model.rawFullTabName()))
	w.Write([]byte(" SET "))
	inserts := model.getInserts()
	isFirstGroup, isFirstNode := true, true
	inserts.toSimpleClause(assignExpr, w, vals, &isFirstGroup, &isFirstNode)
}

func genInsertOrUpdateStmt(model Model, w io.Writer, vals *[]interface{}) {
	w.Write([]byte("INSERT INTO "))
	w.Write([]byte(model.rawFullTabName()))
	w.Write([]byte(" SET "))
	isFirstGroup, isFirstNode := true, true
	model.getInserts().toSimpleClause(assignExpr, w, vals, &isFirstGroup, &isFirstNode)
	w.Write([]byte(" ON DUPLICATE KEY UPDATE "))
	isFirstGroup, isFirstNode = true, true
	model.getUpdates().toSimpleClause(w, vals, &isFirstGroup, &isFirstNode)
}

func genSelectStmt(model Model, w io.Writer, vals *[]interface{}) {
	isFirstGroup, isFirstNode := true, true
	genSelectedClause(model, w, vals, &isFirstGroup, &isFirstNode)
	isFirstGroup, isFirstNode = true, true
	genTabRefClause(model, nil, noJoin, w, vals)
	isFirstGroup, isFirstNode = true, true
	genWhereClause(model, w, vals, &isFirstGroup, &isFirstNode)
	isFirstGroup, isFirstNode = true, true
	genGroupByClause(model, w, vals, &isFirstGroup, &isFirstNode)
	isFirstGroup, isFirstNode = true, true
	genHavingClause(model, w, vals, &isFirstGroup, &isFirstNode)
	isFirstGroup, isFirstNode = true, true
	genOrderByClause(model, w, vals, &isFirstGroup, &isFirstNode)
	genLimitClause(model, w, vals)
	if model.checkStatus(selectForUpdate) {
		w.Write([]byte("FOR UPDATE "))
	}
}

func genBackQueryStmt(model Model, w io.Writer, vals *[]interface{}) {
	isFirstGroup, isFirstNode := true, true
	genSelectedClause(model, w, vals, &isFirstGroup, &isFirstNode)
	genBackTabRefClause(model, nil, noJoin, w, vals)
	genBackWhereClause(model, w, vals)
	isFirstGroup, isFirstNode = true, true
	genGroupByClause(model, w, vals, &isFirstGroup, &isFirstNode)
	isFirstGroup, isFirstNode = true, true
	genHavingClause(model, w, vals, &isFirstGroup, &isFirstNode)
	isFirstGroup, isFirstNode = true, true
	genOrderByClause(model, w, vals, &isFirstGroup, &isFirstNode)
	genLimitClause(model, w, vals)
	if model.checkStatus(selectForUpdate) {
		w.Write([]byte("FOR UPDATE "))
	}
}

func genUpdateStmt(model Model, w io.Writer, vals *[]interface{}) {
	isFirstGroup, isFirstNode := true, true
	genTabRefClause(model, nil, noJoin, w, vals)
	isFirstGroup, isFirstNode = true, true
	genUpdateClause(model, w, vals, &isFirstGroup, &isFirstNode)
	genWhereClause(model, w, vals, &isFirstGroup, &isFirstNode)
}

func genDeleteStmt(model Model, w io.Writer, vals *[]interface{}) {
	isFirstGroup, isFirstNode := true, true
	genDeleteClause(model, w, vals, true)
	isFirstGroup, isFirstNode = true, true
	genTabRefClause(model, nil, noJoin, w, vals)
	genWhereClause(model, w, vals, &isFirstGroup, &isFirstNode)
	isFirstGroup, isFirstNode = true, true
	genOrderByClause(model, w, vals, &isFirstGroup, &isFirstNode)
	genLimitClause(model, w, vals)
}

func genSelectedClause(model Model, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	selectors := model.getSelectors()
	if _, ok := model.(ModelList); ok {
		selectors.addStatus(selectorStatusFoundRows)
	}
	selectors.toClause(w, vals, isFirstGroup, isFirstNode)
	for _, relInfo := range model.relations() {
		if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
			genSelectedClause(relInfo.lastModel(), w, vals, isFirstGroup, isFirstNode)
		}
	}
}

func genWhereClause(model Model, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	wheres := model.getWheres()
	if wheres != nil {
		model.getWheres().toClause(w, vals, isFirstGroup, isFirstNode)
	}
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
		NewExpr("@ = ?", k, k.value()).toClause(w, vals, nil, nil)

	}
	isFirstGroup, isFirstNode := false, false
	genWhereClause(model, w, vals, &isFirstGroup, &isFirstNode)
}

func genOrderByClause(model Model, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	orderBys := model.getOrderBys()
	orderBys.toClause(w, vals, isFirstGroup, isFirstNode)
	for _, relInfo := range model.relations() {
		if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
			genOrderByClause(relInfo.lastModel(), w, vals, isFirstGroup, isFirstNode)
		}
	}
}

func genGroupByClause(model Model, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	groupBys := model.getGroupBys()
	groupBys.toClause(w, vals, isFirstGroup, isFirstNode)
	for _, relInfo := range model.relations() {
		if model.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			genGroupByClause(relInfo.lastModel(), w, vals, isFirstGroup, isFirstNode)
		}
	}
}

func genHavingClause(model Model, w io.Writer, vals *[]interface{}, isFirstGroup, isFirstNode *bool) {
	havings := model.getHavings()
	if havings != nil {
		havings.toClause(w, vals, isFirstGroup, isFirstNode)
	}
	for _, relInfo := range model.relations() {
		dstModel := relInfo.lastModel()
		if dstModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			genHavingClause(dstModel, w, vals, isFirstGroup, isFirstNode)
		}
	}
}

func genTabRefClause(model Model, relInfo *RelationInfo, joinType joinType, w io.Writer, vals *[]interface{}) {
	if relInfo == nil {
		w.Write([]byte("FROM "))
		w.Write([]byte(model.fullTabName()))
		w.Write([]byte(" "))
	} else {
		relInfo.toClause(joinType, w, vals)
	}
	for _, relInfo := range model.relations() {
		dstModel := relInfo.lastModel()
		if dstModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			switch {
			case dstModel.checkStatus(forJoin):
				genTabRefClause(dstModel, relInfo, join, w, vals)
			case dstModel.checkStatus(forLeftJoin):
				genTabRefClause(dstModel, relInfo, leftJoin, w, vals)
			case dstModel.checkStatus(forRightJoin):
				genTabRefClause(dstModel, relInfo, rightJoin, w, vals)
			}
		}
	}
}

func genBackTabRefClause(model Model, relInfo *RelationInfo, joinType joinType, w io.Writer, vals *[]interface{}) {
	if relInfo == nil {
		w.Write([]byte("FROM "))
		parent := model.getParent()
		if parent == nil {
			panic("no parent model for back query")
		}
		isFirstGroup, isFirstNode := true, true
		parent.toRefClause(w, vals, &isFirstGroup, &isFirstNode)
		var got bool
		for _, relInfo := range parent.relations() {
			if relInfo.lastModel() == model {
				relInfo.toClause(join, w, vals)
				got = true
				break
			}
		}
		if !got {
			panic("cannot find relation")
		}
	} else {
		relInfo.toClause(joinType, w, vals)
	}
	for _, relInfo := range model.relations() {
		dstModel := relInfo.lastModel()
		if dstModel.checkStatus(forJoin | forLeftJoin | forRightJoin) {
			switch {
			case dstModel.checkStatus(forJoin):
				genBackTabRefClause(dstModel, relInfo, join, w, vals)
			case dstModel.checkStatus(forLeftJoin):
				genBackTabRefClause(dstModel, relInfo, leftJoin, w, vals)
			case dstModel.checkStatus(forRightJoin):
				genBackTabRefClause(dstModel, relInfo, rightJoin, w, vals)
			}
		}
	}
}
