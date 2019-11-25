package nborm

import (
	"io"
	"strings"
)

func genFlags() (*bool, *bool, func()) {
	isFirstGroup, isFirstNode := true, true
	return &isFirstGroup, &isFirstNode, func() {
		isFirstGroup, isFirstNode = true, true
	}
}

func genInsertStmt(model Model, w io.Writer, vals *[]interface{}) {
	w.Write([]byte("INSERT INTO "))
	w.Write([]byte(model.rawFullTabName()))
	w.Write([]byte(" "))
	inserts := model.getInserts()
	isFirstGroup, isFirstNode, _ := genFlags()
	inserts.toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func genBulkInsertStmt(models []Model, w io.Writer, vals *[][]interface{}) {
	fakeVals := make([]interface{}, 0, 8)
	w.Write([]byte("INSERT INTO "))
	w.Write([]byte(models[0].rawFullTabName()))
	inserts := models[0].getInserts()
	isFirstGroup, isFirstNode, reset := genFlags()
	inserts.toSimpleRefClause(w, &fakeVals, isFirstGroup, isFirstNode)
	w.Write([]byte(") "))
	reset()
	inserts.toValueListClause(w, &fakeVals, isFirstGroup, isFirstNode)
	w.Write([]byte(")"))
	for _, m := range models {
		var fakeBuilder strings.Builder
		fg, fn, _ := genFlags()
		vs := make([]interface{}, 0, 8)
		inserts := m.getInserts()
		for _, inst := range inserts {
			inst.value.toClause(&fakeBuilder, &vs, fg, fn)
		}
		*vals = append(*vals, vs)
	}
}

func genLoadDataInfileStmt(models []Model, w io.Writer, fw io.Writer, filename string) {
	w.Write([]byte("LOAD DATA INFILE '"))
	w.Write([]byte(filename))
	w.Write([]byte("' INTO TABLE "))
	w.Write([]byte(models[0].rawFullTabName()))
	w.Write([]byte(` FIELDS TERMINATED BY ',' ENCLOSED BY '"' LINES TERMINATED BY '\n' IGNORE 1 ROWS `))
	inserts := models[0].getInserts()
	vals := make([]interface{}, 0, 8)
	fg, fn, reset := genFlags()
	inserts.toSimpleRefClause(w, &vals, fg, fn)
	w.Write([]byte(") "))
	reset()
	for _, inst := range inserts {
		inst.toLoadDataExpr(w, &vals, fg, fn)
	}
	reset()
	for _, inst := range inserts {
		inst.toCSVColName(fw, &vals, fg, fn)
	}
	fw.Write([]byte("\n"))
	for _, m := range models {
		insts := m.getInserts()
		reset()
		for _, inst := range insts {
			inst.toCSVVal(fw, &vals, fg, fn)
		}
		fw.Write([]byte("\n"))
	}
}

func genInsertOrUpdateStmt(model Model, w io.Writer, vals *[]interface{}) {
	w.Write([]byte("INSERT INTO "))
	w.Write([]byte(model.rawFullTabName()))
	w.Write([]byte(" "))
	isFirstGroup, isFirstNode, reset := genFlags()
	model.getInserts().toSimpleClause(w, vals, isFirstGroup, isFirstNode)
	w.Write([]byte(" ON DUPLICATE KEY UPDATE "))
	reset()
	*isFirstGroup = false
	model.getUpdates().toSimpleClause(w, vals, isFirstGroup, isFirstNode)
}

func genSelectStmt(model Model, w io.Writer, vals *[]interface{}) {
	isFirstGroup, isFirstNode, reset := genFlags()
	genSelectedClause(model, w, vals, isFirstGroup, isFirstNode)
	reset()
	genTabRefClause(model, nil, noJoin, w, vals, isFirstNode)
	reset()
	genWhereClause(model, w, vals, isFirstGroup, isFirstNode)
	reset()
	genGroupByClause(model, w, vals, isFirstGroup, isFirstNode)
	reset()
	genHavingClause(model, w, vals, isFirstGroup, isFirstNode)
	reset()
	genOrderByClause(model, w, vals, isFirstGroup, isFirstNode)
	genLimitClause(model, w, vals)
	if model.checkStatus(selectForUpdate) {
		w.Write([]byte("FOR UPDATE "))
	}
}

func genBackQueryStmt(model Model, w io.Writer, vals *[]interface{}) {
	isFirstGroup, isFirstNode, reset := genFlags()
	genSelectedClause(model, w, vals, isFirstGroup, isFirstNode)
	genBackTabRefClause(model, nil, noJoin, w, vals)
	genBackWhereClause(model, w, vals)
	reset()
	genGroupByClause(model, w, vals, isFirstGroup, isFirstNode)
	reset()
	genHavingClause(model, w, vals, isFirstGroup, isFirstNode)
	reset()
	genOrderByClause(model, w, vals, isFirstGroup, isFirstNode)
	genLimitClause(model, w, vals)
	if model.checkStatus(selectForUpdate) {
		w.Write([]byte("FOR UPDATE "))
	}
}

func genUpdateStmt(model Model, w io.Writer, vals *[]interface{}) {
	w.Write([]byte("UPDATE "))
	isFirstGroup, isFirstNode, reset := genFlags()
	*isFirstNode = false
	genTabRefClause(model, nil, noJoin, w, vals, isFirstNode)
	reset()
	genUpdateClause(model, w, vals, isFirstGroup, isFirstNode)
	reset()
	genWhereClause(model, w, vals, isFirstGroup, isFirstNode)
}

func genDeleteStmt(model Model, w io.Writer, vals *[]interface{}) {
	isFirstGroup, isFirstNode, reset := genFlags()
	genDeleteClause(model, w, vals, isFirstNode)
	reset()
	genTabRefClause(model, nil, noJoin, w, vals, isFirstNode)
	reset()
	genWhereClause(model, w, vals, isFirstGroup, isFirstNode)
	reset()
	genOrderByClause(model, w, vals, isFirstGroup, isFirstNode)
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
		if relInfo.lastModel().checkStatus(forJoin | forLeftJoin | forRightJoin) {
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

func genTabRefClause(model Model, relInfo *RelationInfo, joinType joinType, w io.Writer, vals *[]interface{}, isFirstNode *bool) {
	if *isFirstNode {
		*isFirstNode = false
		w.Write([]byte("FROM "))
	}
	if relInfo == nil {
		// w.Write([]byte("FROM "))
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
				genTabRefClause(dstModel, relInfo, join, w, vals, isFirstNode)
			case dstModel.checkStatus(forLeftJoin):
				genTabRefClause(dstModel, relInfo, leftJoin, w, vals, isFirstNode)
			case dstModel.checkStatus(forRightJoin):
				genTabRefClause(dstModel, relInfo, rightJoin, w, vals, isFirstNode)
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
