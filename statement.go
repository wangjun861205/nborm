package nborm

import "fmt"

func genInsertStmt(model Model) (string, []interface{}) {
	inserts := model.getInserts()
	insertClause, insertValues := inserts.toSimpleClause(assignExpr)
	return fmt.Sprintf("INSERT INTO %s SET %s", model.rawFullTabName(), insertClause), insertValues
}

func genInsertOrUpdateStmt(model Model) (string, []interface{}) {
	insertClause, insertValues := model.getInserts().toSimpleClause(assignExpr)
	updateClause, updateValues := model.getUpdates().toSimpleClause(assignExpr)
	return fmt.Sprintf("INSERT INTO %s SET %s ON DUPLICATE KEY UPDATE %s", model.rawFullTabName(), insertClause, updateClause), append(insertValues, updateValues...)
}

func genSelectStmt(model Model) (string, []interface{}) {
	selectedClause := genSelectedClause(model)
	tabRefClause := genTabRefClause(model)
	whereClause, whereValues := genWhereClause(model)
	groupByClause := genGroupByClause(model)
	havingClause, havingValues := genHavingClause(model)
	orderByClause := genOrderByClause(model)
	limitClause := genLimitClause(model)
	return fmt.Sprintf("%s FROM %s %s %s %s %s %s", selectedClause, tabRefClause, whereClause, groupByClause, havingClause, orderByClause, limitClause),
		append(whereValues, havingValues...)

}

func genBackQueryStmt(model Model) (string, []interface{}) {
	selectedClause := genSelectedClause(model)
	tabRefClause := genBackTabRefClause(model)
	whereClause, whereValues := genBackWhereClause(model)
	groupByClause := genGroupByClause(model)
	havingClause, havingValues := genHavingClause(model)
	orderByClause := genOrderByClause(model)
	limitClause := genLimitClause(model)
	return fmt.Sprintf("%s FROM %s %s %s %s %s %s", selectedClause, tabRefClause, whereClause, groupByClause, havingClause, orderByClause, limitClause),
		append(whereValues, havingValues...)
}

func genUpdateStmt(model Model) (string, []interface{}) {
	tabRefClause := genTabRefClause(model)
	updateClause, updateValues := genUpdateClause(model)
	whereClause, whereValues := genWhereClause(model)
	return fmt.Sprintf(`UPDATE %s SET %s %s`, tabRefClause, updateClause, whereClause), append(updateValues, whereValues...)
}
