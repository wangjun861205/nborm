package nborm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func insert(addr uintptr, tabInfo *TableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.Columns))
	valList := make([]interface{}, 0, len(tabInfo.Columns))
	for _, info := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.IsInc && field.IsValid() {
			colList = append(colList, info.colName())
			valList = append(valList, field.value())
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", tabInfo.dbName(), tabInfo.tabName(), strings.Join(colList, ", "),
		strings.TrimRight(strings.Repeat("?,", len(colList)), ","))
	res, err := getConn(tabInfo.dbName()).Exec(stmt, valList...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertInTx(tx *sql.Tx, addr uintptr, tabInfo *TableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.Columns))
	valList := make([]interface{}, 0, len(tabInfo.Columns))
	for _, info := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.IsInc && field.IsValid() {
			colList = append(colList, info.colName())
			valList = append(valList, field.value())
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", tabInfo.dbName(), tabInfo.tabName(), strings.Join(colList, ", "),
		strings.TrimRight(strings.Repeat("?,", len(colList)), ","))
	res, err := tx.Exec(stmt, valList...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertContext(ctx context.Context, addr uintptr, tabInfo *TableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.Columns))
	valList := make([]interface{}, 0, len(tabInfo.Columns))
	for _, info := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.IsInc && field.IsValid() {
			colList = append(colList, info.colName())
			valList = append(valList, field.value())
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", tabInfo.dbName(), tabInfo.tabName(), strings.Join(colList, ", "),
		strings.TrimRight(strings.Repeat("?,", len(colList)), ","))
	res, err := getConn(tabInfo.dbName()).ExecContext(ctx, stmt, valList...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertContextInTx(tx *sql.Tx, ctx context.Context, addr uintptr, tabInfo *TableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.Columns))
	valList := make([]interface{}, 0, len(tabInfo.Columns))
	for _, info := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.IsInc && field.IsValid() {
			colList = append(colList, info.colName())
			valList = append(valList, field.value())
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", tabInfo.dbName(), tabInfo.tabName(), strings.Join(colList, ", "),
		strings.TrimRight(strings.Repeat("?,", len(colList)), ","))
	res, err := tx.ExecContext(ctx, stmt, valList...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdate(addr uintptr, tabInfo *TableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.Columns))
	valList := make([]interface{}, 0, len(tabInfo.Columns))
	updList := make([]string, 0, len(tabInfo.Columns))
	for _, info := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.IsInc && field.IsValid() {
			colList = append(colList, info.colName())
			valList = append(valList, field.value())
			updList = append(updList, fmt.Sprintf("%s = ?", info.colName()))
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", tabInfo.dbName(), tabInfo.tabName(),
		strings.Join(colList, ", "), strings.Trim(strings.Repeat("?, ", len(colList)), ", "), tabInfo.Inc.colName(), tabInfo.Inc.colName(),
		strings.Join(updList, ", "))
	res, err := getConn(tabInfo.dbName()).Exec(stmt, append(valList, valList...)...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdateInTx(tx *sql.Tx, addr uintptr, tabInfo *TableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.Columns))
	valList := make([]interface{}, 0, len(tabInfo.Columns))
	updList := make([]string, 0, len(tabInfo.Columns))
	for _, info := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.IsInc && field.IsValid() {
			colList = append(colList, info.colName())
			valList = append(valList, field.value())
			updList = append(updList, fmt.Sprintf("%s = ?", info.colName()))
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", tabInfo.dbName(), tabInfo.tabName(),
		strings.Join(colList, ", "), strings.Trim(strings.Repeat("?, ", len(colList)), ", "), tabInfo.Inc.colName(), tabInfo.Inc.colName(),
		strings.Join(updList, ", "))
	res, err := tx.Exec(stmt, append(valList, valList...)...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdateContext(ctx context.Context, addr uintptr, tabInfo *TableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.Columns))
	valList := make([]interface{}, 0, len(tabInfo.Columns))
	updList := make([]string, 0, len(tabInfo.Columns))
	for _, info := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.IsInc && field.IsValid() {
			colList = append(colList, info.colName())
			valList = append(valList, field.value())
			updList = append(updList, fmt.Sprintf("%s = ?", info.colName()))
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", tabInfo.dbName(), tabInfo.tabName(),
		strings.Join(colList, ", "), strings.Trim(strings.Repeat("?, ", len(colList)), ", "), tabInfo.Inc.colName(), tabInfo.Inc.colName(),
		strings.Join(updList, ", "))
	res, err := getConn(tabInfo.dbName()).ExecContext(ctx, stmt, append(valList, valList...)...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdateContextInTx(tx *sql.Tx, ctx context.Context, addr uintptr, tabInfo *TableInfo) (int64, error) {
	colList := make([]string, 0, len(tabInfo.Columns))
	valList := make([]interface{}, 0, len(tabInfo.Columns))
	updList := make([]string, 0, len(tabInfo.Columns))
	for _, info := range tabInfo.Columns {
		field := getFieldByColumnInfo(addr, info)
		if !info.IsInc && field.IsValid() {
			colList = append(colList, info.colName())
			valList = append(valList, field.value())
			updList = append(updList, fmt.Sprintf("%s = ?", info.colName()))
		}
	}
	stmt := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s", tabInfo.dbName(), tabInfo.tabName(),
		strings.Join(colList, ", "), strings.Trim(strings.Repeat("?, ", len(colList)), ", "), tabInfo.Inc.colName(), tabInfo.Inc.colName(),
		strings.Join(updList, ", "))
	res, err := tx.ExecContext(ctx, stmt, append(valList, valList...)...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func queryRow(tabInfo *TableInfo, where *Where) *sql.Row {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause)
	return getConn(tabInfo.dbName()).QueryRow(stmt, valueList...)
}

func queryRowInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where) *sql.Row {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s FOR UPDATE", tabInfo.dbName(), tabInfo.tabName(), whereClause)
	return tx.QueryRow(stmt, valueList...)
}

func queryRowContext(ctx context.Context, tabInfo *TableInfo, where *Where) *sql.Row {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause)
	return getConn(tabInfo.dbName()).QueryRowContext(ctx, stmt, valueList...)
}

func queryRowContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where) *sql.Row {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s FOR UPDATE", tabInfo.dbName(), tabInfo.tabName(), whereClause)
	return tx.QueryRowContext(ctx, stmt, valueList...)
}

func queryRows(tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s %s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause, sorter.toSQL(), pager.toSQL())
	return getConn(tabInfo.dbName()).Query(stmt, valueList...)
}

func queryRowsInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s %s %s FOR UPDATE", tabInfo.dbName(), tabInfo.tabName(), whereClause, sorter.toSQL(), pager.toSQL())
	return tx.Query(stmt, valueList...)
}

func queryRowsAndFoundRows(tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s %s %s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause, sorter.toSQL(), pager.toSQL())
	tx, err := getConn(tabInfo.dbName()).Begin()
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	rows, err := tx.Query(stmt, valueList...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil
}

func queryRowsAndFoundRowsInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s %s %s %s FOR UPDATE", tabInfo.dbName(), tabInfo.tabName(), whereClause,
		sorter.toSQL(), pager.toSQL())
	rows, err := tx.Query(stmt, valueList...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func queryRowsContext(ctx context.Context, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s %s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause, sorter.toSQL(), pager.toSQL())
	return getConn(tabInfo.dbName()).QueryContext(ctx, stmt, valueList...)
}

func queryRowsContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT * FROM %s.%s %s %s %s FOR UPDATE", tabInfo.dbName(), tabInfo.tabName(), whereClause, sorter.toSQL(), pager.toSQL())
	return tx.QueryContext(ctx, stmt, valueList...)
}

func queryRowsAndFoundRowsContext(ctx context.Context, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s %s %s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause, sorter.toSQL(), pager.toSQL())
	tx, err := getConn(tabInfo.dbName()).Begin()
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	rows, err := tx.QueryContext(ctx, stmt, valueList...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil

}

func queryRowsAndFoundRowsContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s.%s %s %s %s FOR UPDATE", tabInfo.dbName(), tabInfo.tabName(), whereClause, sorter.toSQL(), pager.toSQL())
	rows, err := tx.QueryContext(ctx, stmt, valueList...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func delete(tabInfo *TableInfo, where *Where) (int64, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("DELETE FROM %s.%s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause)
	res, err := getConn(tabInfo.dbName()).Exec(stmt, valueList...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func deleteInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where) (int64, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("DELETE FROM %s.%s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause)
	res, err := tx.Exec(stmt, valueList...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func deleteContext(ctx context.Context, tabInfo *TableInfo, where *Where) (int64, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("DELETE FROM %s.%s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause)
	res, err := getConn(tabInfo.dbName()).ExecContext(ctx, stmt, valueList...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func deleteContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where) (int64, error) {
	whereClause, valueList := where.toClause()
	stmt := fmt.Sprintf("DELETE FROM %s.%s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause)
	res, err := tx.ExecContext(ctx, stmt, valueList...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func update(tabInfo *TableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	whereClause, valueList := where.toClause()
	updColList := make([]string, len(updVals))
	updValList := make([]interface{}, len(updVals))
	for i, v := range updVals {
		updColList[i], updValList[i] = v.toSQL()
	}
	stmt := fmt.Sprintf("UPDATE %s.%s SET %s %s", tabInfo.dbName(), tabInfo.tabName(), strings.Join(updColList, ", "), whereClause)
	res, err := getConn(tabInfo.dbName()).Exec(stmt, append(updValList, valueList...)...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func updateInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	whereClause, valueList := where.toClause()
	updColList := make([]string, len(updVals))
	updValList := make([]interface{}, len(updVals))
	for i, v := range updVals {
		updColList[i], updValList[i] = v.toSQL()
	}
	stmt := fmt.Sprintf("UPDATE %s.%s SET %s %s", tabInfo.dbName(), tabInfo.tabName(), strings.Join(updColList, ", "), whereClause)
	res, err := tx.Exec(stmt, append(updValList, valueList...)...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func updateContext(ctx context.Context, tabInfo *TableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	whereClause, valueList := where.toClause()
	updColList := make([]string, len(updVals))
	updValList := make([]interface{}, len(updVals))
	for i, v := range updVals {
		updColList[i], updValList[i] = v.toSQL()
	}
	stmt := fmt.Sprintf("UPDATE %s.%s SET %s %s", tabInfo.dbName(), tabInfo.tabName(), strings.Join(updColList, ", "), whereClause)
	res, err := getConn(tabInfo.dbName()).ExecContext(ctx, stmt, append(updValList, valueList...)...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func updateContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	whereClause, valueList := where.toClause()
	updColList := make([]string, len(updVals))
	updValList := make([]interface{}, len(updVals))
	for i, v := range updVals {
		updColList[i], updValList[i] = v.toSQL()
	}
	stmt := fmt.Sprintf("UPDATE %s.%s SET %s %s", tabInfo.dbName(), tabInfo.tabName(), strings.Join(updColList, ", "), whereClause)
	res, err := tx.ExecContext(ctx, stmt, append(updValList, valueList...)...)
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

func relationQueryRowInTx(tx *sql.Tx, relation relation, where *Where) *sql.Row {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s FOR UPDATE", relation.getDstDB(), relation.getDstTab(), relation.getSrcDB(), relation.getSrcTab(),
		relation.joinClause(), whereClause)
	return tx.QueryRow(stmt, whereValues...)
}

func relationQueryRows(relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s %s %s", relation.getDstDB(), relation.getDstTab(), relation.getSrcDB(), relation.getSrcTab(),
		relation.joinClause(), whereClause, sorter.toSQL(), pager.toSQL())
	return getConn(relation.getDstDB()).Query(stmt, whereValues...)
}

func relationQueryRowsInTx(tx *sql.Tx, relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s %s %s FOR UPDATE", relation.getDstDB(), relation.getDstTab(), relation.getSrcDB(), relation.getSrcTab(),
		relation.joinClause(), whereClause, sorter.toSQL(), pager.toSQL())
	return tx.Query(stmt, whereValues...)
}

func relationQueryRowsAndFoundRows(relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s %s %s %s %s", relation.getDstDB(), relation.getDstTab(), relation.getSrcDB(), relation.getSrcTab(),
		relation.joinClause(), whereClause, sorter.toSQL(), pager.toSQL())

	tx, err := getConn(relation.getDstDB()).Begin()
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.Query(stmt, whereValues...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil
}

func relationQueryRowsAndFoundRowsInTx(tx *sql.Tx, relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s %s %s %s %s FOR UPDATE", relation.getDstDB(), relation.getDstTab(), relation.getSrcDB(), relation.getSrcTab(),
		relation.joinClause(), whereClause, sorter.toSQL(), pager.toSQL())

	rows, err := tx.Query(stmt, whereValues...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func relationCount(relation relation, where *Where) (int, error) {
	where = relation.where().And(where)
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT COUNT(*) FROM %s.%s %s %s", relation.getSrcDB(), relation.getSrcTab(), relation.joinClause(), whereClause)
	row := getConn(relation.getDstDB()).QueryRow(stmt, whereValues...)
	var num int
	if err := row.Scan(&num); err != nil {
		return 0, err
	}
	return num, nil
}

func relationCountInTx(tx *sql.Tx, relation relation, where *Where) (int, error) {
	where = relation.where().And(where)
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT COUNT(*) FROM %s.%s %s %s FOR UPDATE", relation.getSrcDB(), relation.getSrcTab(), relation.joinClause(), whereClause)
	row := tx.QueryRow(stmt, whereValues...)
	var num int
	if err := row.Scan(&num); err != nil {
		return 0, err
	}
	return num, nil
}

func joinQueryRow(tabInfo *TableInfo, where *Where, relations ...relation) *sql.Row {
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s", tabInfo.dbName(), tabInfo.tabName(), relations[0].getSrcDB(), relations[0].getSrcTab(),
		strings.Join(joinClauses, " "), whereClause)
	return getConn(tabInfo.dbName()).QueryRow(stmt, whereValues...)
}

func joinQueryRowInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, relations ...relation) *sql.Row {
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s %s FOR UPDATE", tabInfo.dbName(), tabInfo.tabName(), relations[0].getSrcDB(), relations[0].getSrcTab(),
		strings.Join(joinClauses, " "), whereClause)
	return tx.QueryRow(stmt, whereValues...)
}

func unionQueryRow(tabInfos []*TableInfo, where *Where, relations ...relation) *sql.Row {
	tabClause := make([]string, len(tabInfos))
	for i, tabInfo := range tabInfos {
		tabClause[i] = fmt.Sprintf("%s.%s.*", tabInfo.dbName(), tabInfo.tabName())
	}
	joinClause := make([]string, len(relations))
	for i, rel := range relations {
		joinClause[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s FROM %s.%s %s %s", strings.Join(tabClause, ", "), relations[0].getSrcDB(),
		relations[0].getSrcTab(), strings.Join(joinClause, " "), whereClause)
	return getConn(tabInfos[0].dbName()).QueryRow(stmt, whereValues...)
}

func unionQueryRowInTx(tx *sql.Tx, tabInfos []*TableInfo, where *Where, relations ...relation) *sql.Row {
	tabClause := make([]string, len(tabInfos))
	for i, tabInfo := range tabInfos {
		tabClause[i] = fmt.Sprintf("%s.%s.*", tabInfo.dbName(), tabInfo.tabName())
	}
	joinClause := make([]string, len(relations))
	for i, rel := range relations {
		joinClause[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s FROM %s.%s %s %s FOR UPDATE", strings.Join(tabClause, ", "), relations[0].getSrcDB(),
		relations[0].getSrcTab(), strings.Join(joinClause, " "), whereClause)
	return tx.QueryRow(stmt, whereValues...)
}

func joinQueryRows(tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, error) {
	distMap := map[bool]string{
		true:  "DISTINCT",
		false: "",
	}
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s %s.%s.* FROM %s.%s %s %s %s %s", distMap[distinct], tabInfo.dbName(), tabInfo.tabName(), relations[0].getSrcDB(),
		relations[0].getSrcTab(), strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	return getConn(tabInfo.dbName()).Query(stmt, whereValues...)
}

func joinQueryRowsInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, error) {
	distMap := map[bool]string{
		true:  "DISTINCT",
		false: "",
	}
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s %s.%s.* FROM %s.%s %s %s %s %s FOR UPDATE", distMap[distinct], tabInfo.dbName(), tabInfo.tabName(), relations[0].getSrcDB(),
		relations[0].getSrcTab(), strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	return tx.Query(stmt, whereValues...)
}

func unionQueryRows(tabInfos []*TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, error) {
	distMap := map[bool]string{
		true:  "DISTINCT",
		false: "",
	}
	tabClause := make([]string, len(tabInfos))
	for i, tabInfo := range tabInfos {
		tabClause[i] = fmt.Sprintf("%s.%s.*", tabInfo.dbName(), tabInfo.tabName())
	}
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s %s FROM %s.%s %s %s %s %s", distMap[distinct], strings.Join(tabClause, ", "), relations[0].getSrcDB(),
		relations[0].getSrcTab(), strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	return getConn(tabInfos[0].dbName()).Query(stmt, whereValues...)
}

func unionQueryRowsInTx(tx *sql.Tx, tabInfos []*TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, error) {
	distMap := map[bool]string{
		true:  "DISTINCT",
		false: "",
	}
	tabClause := make([]string, len(tabInfos))
	for i, tabInfo := range tabInfos {
		tabClause[i] = fmt.Sprintf("%s.%s.*", tabInfo.dbName(), tabInfo.tabName())
	}
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT %s %s FROM %s.%s %s %s %s %s FOR UPDATE", distMap[distinct], strings.Join(tabClause, ", "), relations[0].getSrcDB(),
		relations[0].getSrcTab(), strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	return tx.Query(stmt, whereValues...)
}

func joinQueryRowsAndFoundRows(tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, *sql.Tx, error) {
	distMap := map[bool]string{
		true:  "DISTINCT",
		false: "",
	}
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s %s.%s.* FROM %s.%s %s %s %s %s", distMap[distinct], tabInfo.dbName(), tabInfo.tabName(),
		relations[0].getSrcDB(), relations[0].getSrcTab(), strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	tx, err := getConn(tabInfo.dbName()).Begin()
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.Query(stmt, whereValues...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil
}

func unionQueryRowsAndFoundRows(tabInfos []*TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, *sql.Tx, error) {
	distMap := map[bool]string{
		true:  "DISTINCT",
		false: "",
	}
	tabClause := make([]string, len(tabInfos))
	for i, tabInfo := range tabInfos {
		tabClause[i] = fmt.Sprintf("%s.%s.*", tabInfo.dbName(), tabInfo.tabName())
	}
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s %s FROM %s.%s %s %s %s %s", distMap[distinct], strings.Join(tabClause, ", "),
		relations[0].getSrcDB(), relations[0].getSrcTab(), strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	tx, err := getConn(tabInfos[0].dbName()).Begin()
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.Query(stmt, whereValues...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil
}

func joinQueryRowsAndFoundRowsInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, *sql.Tx, error) {
	distMap := map[bool]string{
		true:  "DISTINCT",
		false: "",
	}
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s %s.%s.* FROM %s.%s %s %s %s %s FOR UPDATE", distMap[distinct], tabInfo.dbName(), tabInfo.tabName(),
		relations[0].getSrcDB(), relations[0].getSrcTab(), strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	rows, err := tx.Query(stmt, whereValues...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func unionQueryRowsAndFoundRowsInTx(tx *sql.Tx, tabInfos []*TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, *sql.Tx, error) {
	distMap := map[bool]string{
		true:  "DISTINCT",
		false: "",
	}
	tabClause := make([]string, len(tabInfos))
	for i, tabInfo := range tabInfos {
		tabClause[i] = fmt.Sprintf("%s.%s.*", tabInfo.dbName(), tabInfo.tabName())
	}
	joinClauses := make([]string, len(relations))
	for i, rel := range relations {
		joinClauses[i] = rel.joinClause()
	}
	whereClause, whereValues := where.toClause()
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s %s FROM %s.%s %s %s %s %s FOR UPDATE", distMap[distinct], strings.Join(tabClause, ", "),
		relations[0].getSrcDB(), relations[0].getSrcTab(), strings.Join(joinClauses, " "), whereClause, sorter.toSQL(), pager.toSQL())
	rows, err := tx.Query(stmt, whereValues...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func deleteAll(tabInfo *TableInfo) (int64, error) {
	res, err := getConn(tabInfo.dbName()).Exec(fmt.Sprintf("DELETE FROM %s.%s", tabInfo.dbName(), tabInfo.tabName()))
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func deleteAllInTx(tx *sql.Tx, tabInfo *TableInfo) (int64, error) {
	res, err := tx.Exec(fmt.Sprintf("DELETE FROM %s.%s", tabInfo.dbName(), tabInfo.tabName()))
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func truncateTable(tabInfo *TableInfo) (int64, error) {
	res, err := getConn(tabInfo.dbName()).Exec(fmt.Sprintf("TRUNCATE TABLE %s.%s", tabInfo.dbName(), tabInfo.tabName()))
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func truncateTableInTx(tx *sql.Tx, tabInfo *TableInfo) (int64, error) {
	res, err := tx.Exec(fmt.Sprintf("TRUNCATE TABLE %s.%s", tabInfo.dbName(), tabInfo.tabName()))
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func count(tabInfo *TableInfo, where *Where) (int, error) {
	whereClause, whereValues := where.toClause()
	row := getConn(tabInfo.dbName()).QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s.%s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause), whereValues...)
	var num int
	if err := row.Scan(&num); err != nil {
		return -1, err
	}
	return num, nil
}

func countInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where) (int, error) {
	whereClause, whereValues := where.toClause()
	row := tx.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s.%s %s", tabInfo.dbName(), tabInfo.tabName(), whereClause), whereValues...)
	var num int
	if err := row.Scan(&num); err != nil {
		return -1, err
	}
	return num, nil
}
