package nborm

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

type Statement struct {
	stmtStr string
	args    []interface{}
}

func getColAndValList(fields ...Field) (colList []string, placeHolderList []string, valList []interface{}) {
	for _, field := range fields {
		colList = append(colList, field.columnName())
		placeHolderList = append(placeHolderList, "?")
		valList = append(valList, field.value())
	}
	return
}

func getColAndValAndUpdList(fields ...Field) (colList []string, valList []interface{}, updateList []string) {
	for _, field := range fields {
		if !field.isInc() {
			colList = append(colList, field.columnName())
			valList = append(valList, field.value())
			updateList = append(updateList, fmt.Sprintf("%s = ?", field.columnName()))
		}
	}
	return
}

func genUpdateColAndValList(updVals ...*UpdateValue) (colList []string, valList []interface{}) {
	for _, updVal := range updVals {
		col, val := updVal.toSQL()
		colList = append(colList, col)
		valList = append(valList, val)
	}
	return
}

func checkFields(fields ...Field) error {
	for _, field := range fields {
		if err := field.check(); err != nil {
			return err
		}
	}
	return nil
}

func genPlaceHolder(num int) string {
	return strings.Trim(strings.Repeat("?, ", num), ", ")
}

func genInsertStmt(addr uintptr, tabInfo *TableInfo) (*Statement, error) {
	colList, placeHolderList, valList, err := processInsertModel(addr, tabInfo)
	if err != nil {
		return nil, err
	}
	stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tabInfo.fullTabName(), strings.Join(colList, ", "), strings.Join(placeHolderList, ", "))
	return &Statement{stmtStr, valList}, nil
}

func genInsertOrUpdateStmt(addr uintptr, tabInfo *TableInfo) (*Statement, error) {
	colList, insertPlaceHolderList, updatePlaceHolderList, valList, err := processInsertOrUpdateModel(addr, tabInfo)
	if err != nil {
		return nil, err
	}
	stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", tabInfo.fullTabName(),
		strings.Join(colList, ", "), strings.Join(insertPlaceHolderList, ", "), tabInfo.Inc.colName(), tabInfo.Inc.colName(),
		strings.Join(updatePlaceHolderList, ", "))
	return &Statement{stmtStr, valList}, nil

}

func genSelectStmt(tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager, forUpdate, foundRows bool) *Statement {
	var forUpdateKeyWords, foundRowsKeyWords string
	if forUpdate {
		forUpdateKeyWords = "FOR UPDATE"
	}
	if foundRows {
		foundRowsKeyWords = "SQL_CALC_FOUND_ROWS"
	}
	whereClause, whereValues := where.toClause()
	stmtStr := fmt.Sprintf("SELECT %s * FROM %s %s %s %s %s", foundRowsKeyWords, tabInfo.fullTabName(), whereClause, sorter.toSQL(), pager.toSQL(),
		forUpdateKeyWords)
	return &Statement{stmtStr, whereValues}
}

func genDeleteStmt(tabInfo *TableInfo, where *Where) *Statement {
	whereClause, whereValues := where.toClause()
	stmtStr := fmt.Sprintf("DELETE FROM %s %s", tabInfo.fullTabName(), whereClause)
	return &Statement{stmtStr, whereValues}
}

func genUpdateStmt(tabInfo *TableInfo, where *Where, updVals ...*UpdateValue) (*Statement, error) {
	if len(updVals) == 0 {
		return nil, fmt.Errorf("nborm.genUpdateStmt() error: no update values (%s)", tabInfo.ModelName)
	}
	whereClause, whereValues := where.toClause()
	updColList, updValList := genUpdateColAndValList(updVals...)
	stmtStr := fmt.Sprintf("UPDATE %s SET %s %s", tabInfo.fullTabName(), strings.Join(updColList, ", "), whereClause)
	return &Statement{stmtStr, append(updValList, whereValues...)}, nil
}

func genRelationJoinSelectStmt(relation relation, where *Where, sorter *Sorter, pager *Pager, isForUpdate, withFoundRows bool) *Statement {
	var forUpdate, foundRows string
	if isForUpdate {
		forUpdate = "FOR UPDATE"
	}
	if withFoundRows {
		foundRows = "SQL_CALC_FOUND_ROWS"
	}
	whereClause, whereValues := where.toClause()
	joinClause, _ := genJoinClause(relation)
	stmtStr := fmt.Sprintf("SELECT %s %s.* FROM %s %s %s %s %s", foundRows, relation.getFullDstTab(), joinClause, whereClause, sorter.toSQL(),
		pager.toSQL(), forUpdate)
	return &Statement{stmtStr, whereValues}
}

func genUnionJoinSelectStmt(tabInfos []*TableInfo, where *Where, sorter *Sorter, pager *Pager, isForUpdate, withFoundRows, distinct bool,
	relations ...relation) (*Statement, error) {
	if len(relations) == 0 {
		return nil, errors.New("nborm.genMultiJoinSelectStmt() error: no relation")
	}
	tabNames := make([]string, len(tabInfos))
	for i, tabInfo := range tabInfos {
		tabNames[i] = tabInfo.fullTabName() + ".*"
	}
	var forUpdate, foundRows, dist string
	if isForUpdate {
		forUpdate = "FOR UPDATE"
	}
	if withFoundRows {
		foundRows = "SQL_CALC_FOUND_ROWS"
	}
	if distinct {
		dist = "DISTINCT"
	}
	whereClause, whereValues := where.toClause()
	joinClause, err := genJoinClause(relations...)
	if err != nil {
		return nil, err
	}
	stmtStr := fmt.Sprintf("SELECT %s %s %s FROM %s %s %s %s %s", foundRows, dist, strings.Join(tabNames, ", "), joinClause, whereClause,
		sorter.toSQL(), pager.toSQL(), forUpdate)
	return &Statement{stmtStr, whereValues}, nil
}

func genCountStmt(tabInfo *TableInfo, where *Where) *Statement {
	whereClause, whereValues := where.toClause()
	stmtStr := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", tabInfo.fullTabName(), whereClause)
	return &Statement{stmtStr, whereValues}
}

func genJoinCountStmt(where *Where, isForUpdate bool, relations ...relation) (*Statement, error) {
	if len(relations) == 0 {
		return nil, errors.New("nborm.genJoinCountStmt() error: no realtion")
	}
	var forUpdateKeyWords string
	if isForUpdate {
		forUpdateKeyWords = "FOR UPDATE"
	}
	whereClause, whereValues := where.toClause()
	joinClause, err := genJoinClause(relations...)
	if err != nil {
		return nil, err
	}
	stmtStr := fmt.Sprintf("SELECT COUNT(*) FROM %s %s %s", joinClause, whereClause, forUpdateKeyWords)
	return &Statement{stmtStr, whereValues}, nil
}

func genTruncateStmt(tabInfo *TableInfo) *Statement {
	stmtStr := fmt.Sprintf("TRUNCATE TABLE %s", tabInfo.fullTabName())
	return &Statement{stmtStr: stmtStr}
}

func genMiddleTableInsertStmt(relation complexRelation, dstAddr uintptr, dstTabInfo *TableInfo) (*Statement, error) {
	leftVal := relation.getSrcVal()
	rightVal := getFieldByName(dstAddr, relation.getRawDstCol(), dstTabInfo).value()
	stmtStr := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES (?, ?)", relation.getFullMidTab(), relation.getMidLeftCol(), relation.getMidRightCol())
	return &Statement{stmtStr, []interface{}{leftVal, rightVal}}, nil
}

func genMiddleTableDeleteStmt(relation complexRelation, dstAddr uintptr, dstTabInfo *TableInfo) (*Statement, error) {
	leftVal := relation.getSrcVal()
	rightVal := getFieldByName(dstAddr, relation.getRawDstCol(), dstTabInfo).value()
	stmtStr := fmt.Sprintf("DELETE FROM %s WHERE %s = ? AND %s = ?", relation.getFullMidTab(), relation.getMidLeftCol(), relation.getMidRightCol())
	return &Statement{stmtStr, []interface{}{leftVal, rightVal}}, nil
}

func genBulkInsertStmt(addr uintptr, tabInfo *TableInfo) (*Statement, error) {
	l := *(*[]uintptr)(unsafe.Pointer(addr))
	placeHolderStr := strings.Trim(strings.Repeat("?,", len(tabInfo.Columns)), ",")
	placeHolderList := make([]string, len(l)-1)
	for i := 0; i < len(l)-1; i++ {
		placeHolderList[i] = fmt.Sprintf("(%s)", placeHolderStr)
	}
	stmtStr := fmt.Sprintf("INSERT INTO %s VALUES %s", tabInfo.fullTabName(), strings.Join(placeHolderList, ", "))
	argList := make([]interface{}, len(tabInfo.Columns)*(len(l)-1))
	for i, modAddr := range l[1:] {
		fields := getAllFieldsWithTableInfo(modAddr, tabInfo)
		for j, field := range fields {
			err := field.check()
			if err != nil {
				return nil, err
			}
			if field.IsValid() {
				if field.IsNull() {
					argList[i*len(tabInfo.Columns)+j] = nil
					continue
				}
				argList[i*len(tabInfo.Columns)+j] = field.value()
			} else {
				argList[i*len(tabInfo.Columns)+j] = field.getDefVal()
			}
		}
	}
	return &Statement{stmtStr, argList}, nil
}
