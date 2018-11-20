package nborm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func insert(addr uintptr, tabInfo *tableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.columns))
	valList := make([]interface{}, 0, len(tabInfo.columns))
	for _, info := range tabInfo.columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.isInc && field.IsValid() {
			colList = append(colList, info.colName)
			valList = append(valList, field.value)
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", tabInfo.db, tabInfo.tab, strings.Join(colList, ", "),
		strings.TrimRight(strings.Repeat("?,", len(colList)), ","))
	res, err := getConn(tabInfo.db).Exec(stmt, valList...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertContext(ctx context.Context, addr uintptr, tabInfo *tableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.columns))
	valList := make([]interface{}, 0, len(tabInfo.columns))
	for _, info := range tabInfo.columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.isInc && field.IsValid() {
			colList = append(colList, info.colName)
			valList = append(valList, field.value)
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", tabInfo.db, tabInfo.tab, strings.Join(colList, ", "),
		strings.TrimRight(strings.Repeat("?,", len(colList)), ","))
	res, err := getConn(tabInfo.db).ExecContext(ctx, stmt, valList...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdate(addr uintptr, tabInfo *tableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.columns))
	valList := make([]interface{}, 0, len(tabInfo.columns))
	updList := make([]string, 0, len(tabInfo.columns))
	for _, info := range tabInfo.columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.isInc && field.IsValid() {
			colList = append(colList, info.colName)
			valList = append(valList, field.value)
			updList = append(updList, fmt.Sprintf("%s = ?", info.colName))
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", tabInfo.db, tabInfo.tab,
		strings.Join(colList, ", "), strings.Trim(strings.Repeat("?, ", len(colList)), ", "), tabInfo.inc.colName, tabInfo.inc.colName,
		strings.Join(updList, ", "))
	res, err := getConn(tabInfo.db).Exec(stmt, append(valList, valList...))
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdateContext(ctx context.Context, addr uintptr, tabInfo *tableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.columns))
	valList := make([]interface{}, 0, len(tabInfo.columns))
	updList := make([]string, 0, len(tabInfo.columns))
	for _, info := range tabInfo.columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.isInc && field.IsValid() {
			colList = append(colList, info.colName)
			valList = append(valList, field.value)
			updList = append(updList, fmt.Sprintf("%s = ?", info.colName))
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", tabInfo.db, tabInfo.tab,
		strings.Join(colList, ", "), strings.Trim(strings.Repeat("?, ", len(colList)), ", "), tabInfo.inc.colName, tabInfo.inc.colName,
		strings.Join(updList, ", "))
	res, err := getConn(tabInfo.db).ExecContext(ctx, stmt, append(valList, valList...))
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func queryRow(tabInfo *tableInfo, where *Where) *sql.Row {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s", tabInfo.db, tabInfo.tab, whereClause)
	return getConn(tabInfo.db).QueryRow(stmt, valueList...)
}

func queryRowContext(ctx context.Context, tabInfo *tableInfo, where *Where) *sql.Row {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s", tabInfo.db, tabInfo.tab, whereClause)
	return getConn(tabInfo.db).QueryRowContext(ctx, stmt, valueList...)
}

func queryRows(tabInfo *tableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s %s %s", tabInfo.db, tabInfo.tab, whereClause, sorter.toSQL(), pager.toSQL())
	return getConn(tabInfo.db).Query(stmt, valueList...)
}

func queryRowsAndFoundRows(tabInfo *tableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, int, *sql.Tx, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s %s %s %s", tabInfo.db, tabInfo.tab, whereClause, sorter.toSQL(), pager.toSQL())
	tx, err := getConn(tabInfo.db).Begin()
	if err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	rows, err := tx.Query(stmt, valueList...)
	if err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	row := tx.QueryRow("SELECT FOUND_ROWS()")
	var numRows int
	err = row.Scan(&numRows)
	if err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	return rows, numRows, tx, nil
}

func queryRowsContext(ctx context.Context, tabInfo *tableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s %s %s", tabInfo.db, tabInfo.tab, whereClause, sorter.toSQL(), pager.toSQL())
	return getConn(tabInfo.db).QueryContext(ctx, stmt, valueList...)
}

func queryRowsAndFoundRowsContext(ctx context.Context, tabInfo *tableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, int,
	*sql.Tx, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s %s %s %s", tabInfo.db, tabInfo.tab, whereClause, sorter.toSQL(), pager.toSQL())
	tx, err := getConn(tabInfo.db).Begin()
	if err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	rows, err := tx.QueryContext(ctx, stmt, valueList...)
	if err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	row := tx.QueryRowContext(ctx, "SELECT FOUND_ROWS()")
	var numRows int
	err = row.Scan(&numRows)
	if err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	return rows, numRows, tx, nil

}

func delete(tabInfo *tableInfo, where *Where) (int64, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("DELETE FROM %s.%s %s", tabInfo.db, tabInfo.tab, whereClause)
	res, err := getConn(tabInfo.db).Exec(stmt, valueList...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func deleteContext(ctx context.Context, tabInfo *tableInfo, where *Where) (int64, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("DELETE FROM %s.%s %s", tabInfo.db, tabInfo.tab, whereClause)
	res, err := getConn(tabInfo.db).ExecContext(ctx, stmt, valueList...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func update(tabInfo *tableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	whereClause, valueList := where.toClause()
	updColList := make([]string, len(updVals))
	updValList := make([]interface{}, len(updVals))
	for i, v := range updVals {
		updColList[i], updValList[i] = v.toSQL()
	}
	stmt := fmt.Sprintf("UPDATE %s.%s SET %s %s", tabInfo.db, tabInfo.tab, strings.Join(updColList, ", "), whereClause)
	res, err := getConn(tabInfo.db).Exec(stmt, append(updValList, valueList...)...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func updateContext(ctx context.Context, tabInfo *tableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	whereClause, valueList := where.toClause()
	updColList := make([]string, len(updVals))
	updValList := make([]interface{}, len(updVals))
	for i, v := range updVals {
		updColList[i], updValList[i] = v.toSQL()
	}
	stmt := fmt.Sprintf("UPDATE %s.%s SET %s %s", tabInfo.db, tabInfo.tab, strings.Join(updColList, ", "), whereClause)
	res, err := getConn(tabInfo.db).ExecContext(ctx, stmt, append(updValList, valueList...)...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func relationQueryRow(relation relation, where *Where) *sql.Row {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s", relation.getDstDB(), relation.getDstTab(), relation.getSrcDB(), relation.getSrcTab(),
		relation.joinClause(), whereClause)
	return getConn(relation.getDstDB()).QueryRow(stmt, whereValues...)
}

func relationQueryRows(relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s %s %s", relation.getDstDB(), relation.getDstTab(), relation.getSrcDB(), relation.getSrcTab(),
		relation.joinClause(), whereClause, sorter.toSQL(), pager.toSQL())
	return getConn(relation.getDstDB()).Query(stmt, whereValues...)
}

func relationQueryRowsAndFoundRows(relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, int, *sql.Tx, error) {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s %s %s %s %s", relation.getDstDB(), relation.getDstTab(), relation.getSrcDB(), relation.getSrcTab(),
		relation.joinClause(), whereClause, sorter.toSQL(), pager.toSQL())

	tx, err := getConn(relation.getDstDB()).Begin()
	if err != nil {
		return nil, -1, nil, err
	}
	rows, err := tx.Query(stmt, whereValues...)
	if err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	var numRows int
	err = tx.QueryRow("SELECT FOUND_ROWS()").Scan(&numRows)
	if err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	return rows, numRows, tx, nil
}

func joinQueryRow(tabInfo *tableInfo, where *Where, relations ...relation) *sql.Row {
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s", tabInfo.db, tabInfo.tab, relations[0].getSrcDB(), relations[0].getSrcTab(),
		strings.Join(joinClauses, " "), whereClause)
	return getConn(tabInfo.db).QueryRow(stmt, whereValues...)
}

func joinQueryRows(tabInfo *tableInfo, where *Where, sorter *Sorter, pager *Pager, relations ...relation) (*sql.Rows, error) {
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s %s %s", tabInfo.db, tabInfo.tab, relations[0].getSrcDB(), relations[0].getSrcTab(),
		strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	return getConn(tabInfo.db).Query(stmt, whereValues...)
}

func joinQueryRowsAndFoundRows(tabInfo *tableInfo, where *Where, sorter *Sorter, pager *Pager, relations ...relation) (*sql.Rows, int, *sql.Tx, error) {
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s %s %s %s %s", tabInfo.db, tabInfo.tab, relations[0].getSrcDB(),
		relations[0].getSrcTab(), strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	tx, err := getConn(tabInfo.db).Begin()
	if err != nil {
		return nil, -1, nil, err
	}
	rows, err := tx.Query(stmt, whereValues...)
	if err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	var numRows int
	if err := tx.QueryRow("SELECT FOUND_ROWS()").Scan(&numRows); err != nil {
		tx.Rollback()
		return nil, -1, nil, err
	}
	return rows, numRows, tx, nil
}

func truncateTable(tabInfo *tableInfo) (int64, error) {
	res, err := getConn(tabInfo.db).Exec(fmt.Sprintf("TRUNCATE TABLE %s.%s", tabInfo.db, tabInfo.tab))
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func count(tabInfo *tableInfo, where *Where) (int, error) {
	whereClause, whereValues := where.toClause()
	row := getConn(tabInfo.db).QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s.%s %s", tabInfo.db, tabInfo.tab, whereClause), whereValues...)
	var num int
	if err := row.Scan(&num); err != nil {
		return -1, err
	}
	return num, nil
}
