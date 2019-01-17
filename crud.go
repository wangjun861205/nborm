package nborm

import (
	"context"
	"database/sql"
)

func insert(addr uintptr, tabInfo *TableInfo) (int64, error) {
	stmt, err := genInsertStmt(addr, tabInfo)
	if err != nil {
		return -1, err
	}
	res, err := getConn(tabInfo.dbName()).Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertInTx(tx *sql.Tx, addr uintptr, tabInfo *TableInfo) (int64, error) {
	stmt, err := genInsertStmt(addr, tabInfo)
	if err != nil {
		return -1, err
	}
	res, err := tx.Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertContext(ctx context.Context, addr uintptr, tabInfo *TableInfo) (int64, error) {
	stmt, err := genInsertStmt(addr, tabInfo)
	if err != nil {
		return -1, err
	}
	res, err := getConn(tabInfo.dbName()).ExecContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertContextInTx(tx *sql.Tx, ctx context.Context, addr uintptr, tabInfo *TableInfo) (int64, error) {
	stmt, err := genInsertStmt(addr, tabInfo)
	if err != nil {
		return -1, err
	}
	res, err := tx.ExecContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdate(addr uintptr, tabInfo *TableInfo) (int64, error) {
	stmt, err := genInsertOrUpdateStmt(addr, tabInfo)
	if err != nil {
		return -1, err
	}
	res, err := getConn(tabInfo.dbName()).Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdateInTx(tx *sql.Tx, addr uintptr, tabInfo *TableInfo) (int64, error) {
	stmt, err := genInsertOrUpdateStmt(addr, tabInfo)
	if err != nil {
		return -1, err
	}
	res, err := tx.Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdateContext(ctx context.Context, addr uintptr, tabInfo *TableInfo) (int64, error) {
	stmt, err := genInsertOrUpdateStmt(addr, tabInfo)
	if err != nil {
		return -1, err
	}
	res, err := getConn(tabInfo.dbName()).ExecContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func insertOrUpdateContextInTx(tx *sql.Tx, ctx context.Context, addr uintptr, tabInfo *TableInfo) (int64, error) {
	stmt, err := genInsertOrUpdateStmt(addr, tabInfo)
	if err != nil {
		return -1, err
	}
	res, err := tx.ExecContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func queryRow(tabInfo *TableInfo, where *Where) *sql.Row {
	stmt := genSelectStmt(tabInfo, where, nil, nil, false, false)
	return getConn(tabInfo.dbName()).QueryRow(stmt.stmtStr, stmt.args...)
}

func queryRowInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where) *sql.Row {
	stmt := genSelectStmt(tabInfo, where, nil, nil, true, false)
	return tx.QueryRow(stmt.stmtStr, stmt.args...)
}

func queryRowContext(ctx context.Context, tabInfo *TableInfo, where *Where) *sql.Row {
	stmt := genSelectStmt(tabInfo, where, nil, nil, false, false)
	return getConn(tabInfo.dbName()).QueryRowContext(ctx, stmt.stmtStr, stmt.args...)
}

func queryRowContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where) *sql.Row {
	stmt := genSelectStmt(tabInfo, where, nil, nil, true, false)
	return tx.QueryRowContext(ctx, stmt.stmtStr, stmt.args...)
}

func queryRows(tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	stmt := genSelectStmt(tabInfo, where, sorter, pager, false, false)
	return getConn(tabInfo.dbName()).Query(stmt.stmtStr, stmt.args...)
}

func queryRowsInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	stmt := genSelectStmt(tabInfo, where, sorter, pager, true, false)
	return tx.Query(stmt.stmtStr, stmt.args...)
}

func queryRowsAndFoundRows(tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	stmt := genSelectStmt(tabInfo, where, sorter, pager, false, true)
	tx, err := getConn(tabInfo.dbName()).Begin()
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.Query(stmt.stmtStr, stmt.args...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil
}

func queryRowsAndFoundRowsInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	stmt := genSelectStmt(tabInfo, where, sorter, pager, true, true)
	rows, err := tx.Query(stmt.stmtStr, stmt.args...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func queryRowsContext(ctx context.Context, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	stmt := genSelectStmt(tabInfo, where, sorter, pager, false, false)
	return getConn(tabInfo.dbName()).QueryContext(ctx, stmt.stmtStr, stmt.args...)
}

func queryRowsContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	stmt := genSelectStmt(tabInfo, where, sorter, pager, true, false)
	return tx.QueryContext(ctx, stmt.stmtStr, stmt.args...)
}

func queryRowsAndFoundRowsContext(ctx context.Context, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	stmt := genSelectStmt(tabInfo, where, sorter, pager, false, true)
	tx, err := getConn(tabInfo.dbName()).Begin()
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.QueryContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil

}

func queryRowsAndFoundRowsContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	stmt := genSelectStmt(tabInfo, where, sorter, pager, true, true)
	rows, err := tx.QueryContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func delete(tabInfo *TableInfo, where *Where) (int64, error) {
	stmt := genDeleteStmt(tabInfo, where)
	res, err := getConn(tabInfo.dbName()).Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func deleteInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where) (int64, error) {
	stmt := genDeleteStmt(tabInfo, where)
	res, err := tx.Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func deleteContext(ctx context.Context, tabInfo *TableInfo, where *Where) (int64, error) {
	stmt := genDeleteStmt(tabInfo, where)
	res, err := getConn(tabInfo.dbName()).ExecContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func deleteContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where) (int64, error) {
	stmt := genDeleteStmt(tabInfo, where)
	res, err := tx.ExecContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func update(tabInfo *TableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	stmt, err := genUpdateStmt(tabInfo, where, updVals...)
	if err != nil {
		return -1, err
	}
	res, err := getConn(tabInfo.dbName()).Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func updateInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	stmt, err := genUpdateStmt(tabInfo, where, updVals...)
	if err != nil {
		return -1, err
	}
	res, err := tx.Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func updateContext(ctx context.Context, tabInfo *TableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	stmt, err := genUpdateStmt(tabInfo, where, updVals...)
	if err != nil {
		return -1, err
	}
	res, err := getConn(tabInfo.dbName()).ExecContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func updateContextInTx(tx *sql.Tx, ctx context.Context, tabInfo *TableInfo, where *Where, updVals ...*UpdateValue) (int64, error) {
	stmt, err := genUpdateStmt(tabInfo, where, updVals...)
	if err != nil {
		return -1, err
	}
	res, err := tx.ExecContext(ctx, stmt.stmtStr, stmt.args...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func relationQueryRow(relation relation, where *Where) *sql.Row {
	where = relation.where().And(where)
	stmt := genRelationJoinSelectStmt(relation, where, nil, nil, false, false)
	return getConn(relation.getDstDB()).QueryRow(stmt.stmtStr, stmt.args...)
}

func relationQueryRowInTx(tx *sql.Tx, relation relation, where *Where) *sql.Row {
	where = relation.where().And(where)
	stmt := genRelationJoinSelectStmt(relation, where, nil, nil, true, false)
	return tx.QueryRow(stmt.stmtStr, stmt.args...)
}

func relationQueryRows(relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	where = relation.where().And(where)
	stmt := genRelationJoinSelectStmt(relation, where, sorter, pager, false, false)
	return getConn(relation.getDstDB()).Query(stmt.stmtStr, stmt.args...)
}

func relationQueryRowsInTx(tx *sql.Tx, relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, error) {
	where = relation.where().And(where)
	stmt := genRelationJoinSelectStmt(relation, where, sorter, pager, true, false)
	return tx.Query(stmt.stmtStr, stmt.args...)
}

func relationQueryRowsAndFoundRows(relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	where = relation.where().And(where)
	stmt := genRelationJoinSelectStmt(relation, where, sorter, pager, false, true)
	tx, err := getConn(relation.getDstDB()).Begin()
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.Query(stmt.stmtStr, stmt.args...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil
}

func relationQueryRowsAndFoundRowsInTx(tx *sql.Tx, relation relation, where *Where, sorter *Sorter, pager *Pager) (*sql.Rows, *sql.Tx, error) {
	where = relation.where().And(where)
	stmt := genRelationJoinSelectStmt(relation, where, sorter, pager, true, true)
	rows, err := tx.Query(stmt.stmtStr, stmt.args...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func relationCount(relation relation, where *Where) (int, error) {
	where = relation.where().And(where)
	stmt, _ := genJoinCountStmt(where, false, relation)
	row := getConn(relation.getDstDB()).QueryRow(stmt.stmtStr, stmt.args...)
	var num int
	if err := row.Scan(&num); err != nil {
		return 0, err
	}
	return num, nil
}

func relationCountInTx(tx *sql.Tx, relation relation, where *Where) (int, error) {
	where = relation.where().And(where)
	stmt, _ := genJoinCountStmt(where, true, relation)
	row := tx.QueryRow(stmt.stmtStr, stmt.args...)
	var num int
	if err := row.Scan(&num); err != nil {
		return 0, err
	}
	return num, nil
}

func joinQueryRow(tabInfo *TableInfo, where *Where, relations ...relation) (*sql.Row, error) {
	stmt, err := genUnionJoinSelectStmt([]*TableInfo{tabInfo}, where, nil, nil, false, false, false, relations...)
	if err != nil {
		return nil, err
	}
	return getConn(tabInfo.dbName()).QueryRow(stmt.stmtStr, stmt.args...), nil
}

func joinQueryRowInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, relations ...relation) (*sql.Row, error) {
	stmt, err := genUnionJoinSelectStmt([]*TableInfo{tabInfo}, where, nil, nil, true, false, false, relations...)
	if err != nil {
		return nil, err
	}
	return tx.QueryRow(stmt.stmtStr, stmt.args...), nil
}

func unionQueryRow(tabInfos []*TableInfo, where *Where, relations ...relation) (*sql.Row, error) {
	stmt, err := genUnionJoinSelectStmt(tabInfos, where, nil, nil, false, false, false, relations...)
	if err != nil {
		return nil, err
	}
	return getConn(tabInfos[0].dbName()).QueryRow(stmt.stmtStr, stmt.args...), nil
}

func unionQueryRowInTx(tx *sql.Tx, tabInfos []*TableInfo, where *Where, relations ...relation) (*sql.Row, error) {
	stmt, err := genUnionJoinSelectStmt(tabInfos, where, nil, nil, true, false, false, relations...)
	if err != nil {
		return nil, err
	}
	return tx.QueryRow(stmt.stmtStr, stmt.args...), nil
}

func joinQueryRows(tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, error) {
	stmt, err := genUnionJoinSelectStmt([]*TableInfo{tabInfo}, where, sorter, pager, false, false, distinct, relations...)
	if err != nil {
		return nil, err
	}
	return getConn(tabInfo.dbName()).Query(stmt.stmtStr, stmt.args...)
}

func joinQueryRowsInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, error) {
	stmt, err := genUnionJoinSelectStmt([]*TableInfo{tabInfo}, where, sorter, pager, true, false, distinct, relations...)
	if err != nil {
		return nil, err
	}
	return tx.Query(stmt.stmtStr, stmt.args...)
}

func unionQueryRows(tabInfos []*TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, error) {
	stmt, err := genUnionJoinSelectStmt(tabInfos, where, sorter, pager, false, false, distinct, relations...)
	if err != nil {
		return nil, err
	}
	return getConn(tabInfos[0].dbName()).Query(stmt.stmtStr, stmt.args...)
}

func unionQueryRowsInTx(tx *sql.Tx, tabInfos []*TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, error) {
	stmt, err := genUnionJoinSelectStmt(tabInfos, where, sorter, pager, true, false, distinct, relations...)
	if err != nil {
		return nil, err
	}
	return tx.Query(stmt.stmtStr, stmt.args...)
}

func joinQueryRowsAndFoundRows(tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, *sql.Tx, error) {
	stmt, err := genUnionJoinSelectStmt([]*TableInfo{tabInfo}, where, sorter, pager, false, true, distinct, relations...)
	if err != nil {
		return nil, nil, err
	}
	tx, err := getConn(tabInfo.dbName()).Begin()
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.Query(stmt.stmtStr, stmt.args...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil
}

func unionQueryRowsAndFoundRows(tabInfos []*TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, *sql.Tx, error) {
	stmt, err := genUnionJoinSelectStmt(tabInfos, where, sorter, pager, false, true, distinct, relations...)
	if err != nil {
		return nil, nil, err
	}
	tx, err := getConn(tabInfos[0].dbName()).Begin()
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.Query(stmt.stmtStr, stmt.args...)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}
	return rows, tx, nil
}

func joinQueryRowsAndFoundRowsInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, *sql.Tx, error) {
	stmt, err := genUnionJoinSelectStmt([]*TableInfo{tabInfo}, where, sorter, pager, true, true, distinct, relations...)
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.Query(stmt.stmtStr, stmt.args...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func unionQueryRowsAndFoundRowsInTx(tx *sql.Tx, tabInfos []*TableInfo, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (*sql.Rows, *sql.Tx, error) {
	stmt, err := genUnionJoinSelectStmt(tabInfos, where, sorter, pager, true, true, distinct, relations...)
	if err != nil {
		return nil, nil, err
	}
	rows, err := tx.Query(stmt.stmtStr, stmt.args...)
	if err != nil {
		return nil, nil, err
	}
	return rows, tx, nil
}

func deleteAll(tabInfo *TableInfo) (int64, error) {
	stmt := genDeleteStmt(tabInfo, nil)
	res, err := getConn(tabInfo.dbName()).Exec(stmt.stmtStr)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func deleteAllInTx(tx *sql.Tx, tabInfo *TableInfo) (int64, error) {
	stmt := genDeleteStmt(tabInfo, nil)
	res, err := tx.Exec(stmt.stmtStr)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func truncateTable(tabInfo *TableInfo) (int64, error) {
	stmt := genTruncateStmt(tabInfo)
	res, err := getConn(tabInfo.dbName()).Exec(stmt.stmtStr)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func truncateTableInTx(tx *sql.Tx, tabInfo *TableInfo) (int64, error) {
	stmt := genTruncateStmt(tabInfo)
	res, err := tx.Exec(stmt.stmtStr)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func count(tabInfo *TableInfo, where *Where) (int, error) {
	stmt := genCountStmt(tabInfo, where)
	row := getConn(tabInfo.rawDBName()).QueryRow(stmt.stmtStr, stmt.args...)
	var num int
	if err := row.Scan(&num); err != nil {
		return -1, err
	}
	return num, nil
}

func countInTx(tx *sql.Tx, tabInfo *TableInfo, where *Where) (int, error) {
	stmt := genCountStmt(tabInfo, where)
	row := tx.QueryRow(stmt.stmtStr, stmt.args...)
	var num int
	if err := row.Scan(&num); err != nil {
		return -1, err
	}
	return num, nil
}

func insertMiddleTable(relation complexRelation, dstAddr uintptr, dstTabInfo *TableInfo) error {
	stmt, err := genMiddleTableInsertStmt(relation, dstAddr, dstTabInfo)
	if err != nil {
		return err
	}
	_, err = getConn(relation.getRawMidDB()).Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return err
	}
	return nil
}

func insertMiddleTableInTx(tx *sql.Tx, relation complexRelation, dstAddr uintptr, dstTabInfo *TableInfo) error {
	stmt, err := genMiddleTableInsertStmt(relation, dstAddr, dstTabInfo)
	if err != nil {
		return err
	}
	_, err = tx.Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return err
	}
	return nil
}

func deleteMiddleTable(relation complexRelation, dstAddr uintptr, dstTabInfo *TableInfo) error {
	stmt, err := genMiddleTableDeleteStmt(relation, dstAddr, dstTabInfo)
	if err != nil {
		return err
	}
	_, err = getConn(relation.getRawMidDB()).Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return err
	}
	return nil
}

func deleteMiddleTableInTx(tx *sql.Tx, relation complexRelation, dstAddr uintptr, dstTabInfo *TableInfo) error {
	stmt, err := genMiddleTableDeleteStmt(relation, dstAddr, dstTabInfo)
	if err != nil {
		return err
	}
	_, err = tx.Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		return err
	}
	return nil
}

func bulkInsert(addr uintptr, tabInfo *TableInfo) (int64, error) {
	stmt, err := genBulkInsertStmt(addr, tabInfo)
	if err != nil {
		return 0, err
	}
	tx, err := getConn(tabInfo.rawDBName()).Begin()
	if err != nil {
		return -1, err
	}
	res, err := tx.Exec(stmt.stmtStr, stmt.args...)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	return lastInsertID, nil
}
