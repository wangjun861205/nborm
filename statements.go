package nborm

import (
	"errors"
	"fmt"
	"strings"
)

type Statement struct {
	stmtStr string
	args    []interface{}
}

func getColAndValList(fields ...Field) (colList []string, valList []interface{}) {
	for _, field := range fields {
		if !field.isInc() {
			colList = append(colList, field.columnName())
			valList = append(valList, field.value())
		}
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
	validFields := getValidFieldsWithTableInfo(addr, tabInfo)
	if len(validFields) == 0 {
		return nil, fmt.Errorf("nborm.genInsertStmt() error: no valid field (%s)", tabInfo.ModelName)
	}
	if err := checkFields(validFields...); err != nil {
		return nil, err
	}
	colList, valList := getColAndValList(validFields...)
	stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tabInfo.fullTabName(), strings.Join(colList, ", "), genPlaceHolder(len(validFields)))
	return &Statement{stmtStr, valList}, nil
}

func genInsertOrUpdateStmt(addr uintptr, tabInfo *TableInfo) (*Statement, error) {
	validFields := getValidFieldsWithTableInfo(addr, tabInfo)
	if len(validFields) == 0 {
		return nil, fmt.Errorf("nborm.genInsertOrUpdateStmt() error: no valid field (%s)", tabInfo.ModelName)
	}
	if err := checkFields(validFields...); err != nil {
		return nil, err
	}
	colList, valList, updList := getColAndValAndUpdList(validFields...)
	stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", tabInfo.fullTabName(),
		strings.Join(colList, ", "), genPlaceHolder(len(validFields)), tabInfo.Inc.colName(), tabInfo.Inc.colName(), strings.Join(updList, ", "))
	return &Statement{stmtStr, append(valList, valList...)}, nil

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