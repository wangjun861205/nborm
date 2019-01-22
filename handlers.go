package nborm

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"unsafe"

	"github.com/go-sql-driver/mysql"
)

//First get the first record in database, no error return when no record in table, check Model synchronized status after query
func First(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where, _ := genWhere(modAddr, tabInfo)
	row := queryRow(tabInfo, where)
	err := scanRow(modAddr, tabInfo, row)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func FirstInTx(tx *sql.Tx, model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where, _ := genWhere(modAddr, tabInfo)
	row := queryRowInTx(tx, tabInfo, where)
	err := scanRow(modAddr, tabInfo, row)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

//GetOne get one record by Model owned field value, if no record, it will return a sql.ErrNoRows error
func GetOne(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where, typ := genWhere(modAddr, tabInfo)
	if typ == otherType {
		return fmt.Errorf("nborm.GetOne() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
	}
	row := queryRow(tabInfo, where)
	return scanRow(modAddr, tabInfo, row)
}

func GetOneInTx(tx *sql.Tx, model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where, typ := genWhere(modAddr, tabInfo)
	if typ == otherType {
		return fmt.Errorf("nborm.GetOne() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
	}
	row := queryRowInTx(tx, tabInfo, where)
	return scanRow(modAddr, tabInfo, row)
}

//GetMul get multiple Models by Models's owned field value, if one of them not has conresponse record, it will return a sql.ErrNoRows error
func GetMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		where, typ := genWhere(addr, tabInfo)
		if typ == otherType {
			return fmt.Errorf("nborm.GetMul() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
		}
		row := queryRowContext(ctx, tabInfo, where)
		return scanRow(addr, tabInfo, row)
	})
}

func GetMulInTx(tx *sql.Tx, slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		where, typ := genWhere(addr, tabInfo)
		if typ == otherType {
			return fmt.Errorf("nborm.GetMul() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
		}
		row := queryRowContextInTx(tx, ctx, tabInfo, where)
		return scanRow(addr, tabInfo, row)
	})
}

//JoinQueryOne query one record by join tables, no error return when no conresponse record, check Model synchronized status after query
func JoinQueryOne(model table, where *Where, relations ...relation) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	row, err := joinQueryRow(tabInfo, where, relations...)
	if err != nil {
		return err
	}
	return scanRow(modAddr, tabInfo, row)
}

func UnionQueryOne(models Union, where *Where, relations ...relation) error {
	tabInfos := models.tabInfos()
	modAddrs := models.addrs()
	row, err := unionQueryRow(tabInfos, where, relations...)
	if err != nil {
		return err
	}
	return unionScanRow(modAddrs, tabInfos, row)
}

func JoinQueryOneInTx(tx *sql.Tx, model table, where *Where, relations ...relation) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	row, err := joinQueryRowInTx(tx, tabInfo, where, relations...)
	if err != nil {
		return err
	}
	return scanRow(modAddr, tabInfo, row)
}

func UnionQueryOneInTx(tx *sql.Tx, models Union, where *Where, relations ...relation) error {
	tabInfos := models.tabInfos()
	modAddrs := models.addrs()
	row, err := unionQueryRowInTx(tx, tabInfos, where, relations...)
	if err != nil {
		return err
	}
	return unionScanRow(modAddrs, tabInfos, row)
}

//JoinQuery join query records by where
func JoinQuery(slice table, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) error {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, err := joinQueryRows(tabInfo, where, sorter, pager, distinct, relations...)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

func UnionQuery(slices Union, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) error {
	tabInfos := slices.tabInfos()
	sliceAddrs := slices.addrs()
	rows, err := unionQueryRows(tabInfos, where, sorter, pager, distinct, relations...)
	if err != nil {
		return err
	}
	return unionScanRows(sliceAddrs, tabInfos, rows)
}

func JoinQueryInTx(tx *sql.Tx, slice table, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) error {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, err := joinQueryRowsInTx(tx, tabInfo, where, sorter, pager, distinct, relations...)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

func UnionQueryInTx(tx *sql.Tx, slices Union, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) error {
	tabInfos := slices.tabInfos()
	sliceAddrs := slices.addrs()
	rows, err := unionQueryRowsInTx(tx, tabInfos, where, sorter, pager, distinct, relations...)
	if err != nil {
		return err
	}
	return unionScanRows(sliceAddrs, tabInfos, rows)
}

//JoinQueryWithFoundRows join query records and get the number of found rows
func JoinQueryWithFoundRows(slice table, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (int, error) {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := joinQueryRowsAndFoundRows(tabInfo, where, sorter, pager, distinct, relations...)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

func UnionQueryWithFoundRows(slices Union, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (int, error) {
	tabInfos := slices.tabInfos()
	sliceAddrs := slices.addrs()
	rows, tx, err := unionQueryRowsAndFoundRows(tabInfos, where, sorter, pager, distinct, relations...)
	if err != nil {
		return -1, err
	}
	if err := unionScanRows(sliceAddrs, tabInfos, rows); err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

func JoinQueryWithFoundRowsInTx(tx *sql.Tx, slice table, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (int, error) {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := joinQueryRowsAndFoundRowsInTx(tx, tabInfo, where, sorter, pager, distinct, relations...)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		return -1, err
	}
	return num, nil
}

func UnionQueryWithFoundRowsInTx(tx *sql.Tx, slices Union, where *Where, sorter *Sorter, pager *Pager, distinct bool, relations ...relation) (int, error) {
	tabInfos := slices.tabInfos()
	sliceAddrs := slices.addrs()
	rows, tx, err := unionQueryRowsAndFoundRowsInTx(tx, tabInfos, where, sorter, pager, distinct, relations...)
	if err != nil {
		return -1, err
	}
	if err := unionScanRows(sliceAddrs, tabInfos, rows); err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

//All get all records of one table
func All(slice table, sorter *Sorter, pager *Pager) error {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, err := queryRows(tabInfo, nil, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

func AllInTx(tx *sql.Tx, slice table, sorter *Sorter, pager *Pager) error {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, err := queryRowsInTx(tx, tabInfo, nil, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

//AllWithFoundRows get all records of one table and the number of records
func AllWithFoundRows(slice table, sorter *Sorter, pager *Pager) (int, error) {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := queryRowsAndFoundRows(tabInfo, nil, sorter, pager)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

func AllWithFoundRowsInTx(tx *sql.Tx, slice table, sorter *Sorter, pager *Pager) (int, error) {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := queryRowsAndFoundRowsInTx(tx, tabInfo, nil, sorter, pager)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		return -1, err
	}
	return num, nil
}

//QueryOne query one record, no error will return when no conresponse record, check the synchronized status of Model after query
func QueryOne(model table, where *Where) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	row := queryRow(tabInfo, where)
	return scanRow(modAddr, tabInfo, row)
}

func QueryOneInTx(tx *sql.Tx, model table, where *Where) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	row := queryRowInTx(tx, tabInfo, where)
	return scanRow(modAddr, tabInfo, row)
}

//QueryMul query multiple Models, no error will be returned if no conresponse record, check synchronized status of Models after query
func QueryMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		where, _ := genWhere(addr, tabInfo)
		row := queryRowContext(ctx, tabInfo, where)
		err := scanRow(addr, tabInfo, row)
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	})
}

func QueryMulInTx(tx *sql.Tx, slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		where, _ := genWhere(addr, tabInfo)
		row := queryRowContextInTx(tx, ctx, tabInfo, where)
		err := scanRow(addr, tabInfo, row)
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	})
}

//Query query records by where
func Query(slice table, where *Where, sorter *Sorter, pager *Pager) error {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, err := queryRows(tabInfo, where, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

func QueryInTx(tx *sql.Tx, slice table, where *Where, sorter *Sorter, pager *Pager) error {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, err := queryRowsInTx(tx, tabInfo, where, sorter, pager)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

//QueryWithFoundRows query records by where and get the number of found rows
func QueryWithFoundRows(slice table, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := queryRowsAndFoundRows(tabInfo, where, sorter, pager)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		tx.Rollback()
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return num, nil
}

func QueryWithFoundRowsInTx(tx *sql.Tx, slice table, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, tx, err := queryRowsAndFoundRowsInTx(tx, tabInfo, where, sorter, pager)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		return -1, err
	}
	num, err := getFoundRows(tx)
	if err != nil {
		return -1, err
	}
	return num, nil
}

//InsertOne insert one record
func InsertOne(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	lid, err := insert(modAddr, tabInfo)
	if err != nil {
		return err
	}
	setInc(modAddr, tabInfo, lid)
	return nil
}

func InsertOneInTx(tx *sql.Tx, model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	lid, err := insertInTx(tx, modAddr, tabInfo)
	if err != nil {
		return err
	}
	setInc(modAddr, tabInfo, lid)
	return nil
}

//InsertOrUpdateOne insert or update one record
func InsertOrUpdateOne(model table) (bool, error) {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	lid, isCreated, err := insertOrUpdate(modAddr, tabInfo)
	if err != nil {
		return false, err
	}
	setInc(modAddr, tabInfo, lid)
	setSync(modAddr, tabInfo)
	return isCreated, nil
}

func InsertOrUpdateOneInTx(tx *sql.Tx, model table) (bool, error) {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	lid, isCreated, err := insertOrUpdateInTx(tx, modAddr, tabInfo)
	if err != nil {
		return false, err
	}
	setInc(modAddr, tabInfo, lid)
	setSync(modAddr, tabInfo)
	return isCreated, nil
}

//InsertMul insert multiple record
func InsertMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		lid, err := insertContext(ctx, addr, tabInfo)
		if err != nil {
			return err
		}
		setInc(addr, tabInfo, lid)
		setSync(addr, tabInfo)
		return nil
	})
}

func InsertMulInTx(tx *sql.Tx, slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		lid, err := insertContextInTx(tx, ctx, addr, tabInfo)
		if err != nil {
			return err
		}
		setInc(addr, tabInfo, lid)
		setSync(addr, tabInfo)
		return nil
	})
}

//InsertOrUpdateMul insert or update multiple record
func InsertOrUpdateMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		lid, err := insertOrUpdateContext(ctx, addr, tabInfo)
		if err != nil {
			return err
		}
		setInc(addr, tabInfo, lid)
		setSync(addr, tabInfo)
		return nil
	})
}

func InsertOrUpdateMulInTx(tx *sql.Tx, slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		lid, err := insertOrUpdateContextInTx(tx, ctx, addr, tabInfo)
		if err != nil {
			return err
		}
		setInc(addr, tabInfo, lid)
		setSync(addr, tabInfo)
		return nil
	})
}

//InsertOrGetOne insert one record or get it when it is already exists
func InsertOrGetOne(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	lid, err := insert(modAddr, tabInfo)
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
			where, typ := genWhere(modAddr, tabInfo)
			if typ == otherType {
				return fmt.Errorf("nborm.InsertOrGetOne() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
			}
			row := queryRow(tabInfo, where)
			err := scanRow(modAddr, tabInfo, row)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	inc := getIncWithTableInfo(modAddr, tabInfo)
	inc.setVal(lid)
	return nil
}

func InsertOrGetOneInTx(tx *sql.Tx, model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	lid, err := insert(modAddr, tabInfo)
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
			where, typ := genWhere(modAddr, tabInfo)
			if typ == otherType {
				return fmt.Errorf("nborm.InsertOrGetOne() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
			}
			row := queryRowInTx(tx, tabInfo, where)
			err := scanRow(modAddr, tabInfo, row)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	inc := getIncWithTableInfo(modAddr, tabInfo)
	inc.setVal(lid)
	return nil
}

//InsertOrGetMul insert multiple records or get them when they are already exist
func InsertOrGetMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		lid, err := insertContext(ctx, addr, tabInfo)
		if err != nil {
			if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
				where, typ := genWhere(addr, tabInfo)
				if typ == otherType {
					return fmt.Errorf("nborm.InsertOrGetMul() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
				}
				row := queryRowContext(ctx, tabInfo, where)
				err := scanRow(addr, tabInfo, row)
				if err != nil {
					return err
				}
				return nil
			}
			return err
		}
		setInc(addr, tabInfo, lid)
		return nil
	})
}

func InsertOrGetMulInTx(tx *sql.Tx, slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		lid, err := insertContext(ctx, addr, tabInfo)
		if err != nil {
			if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
				where, typ := genWhere(addr, tabInfo)
				if typ == otherType {
					return fmt.Errorf("nborm.InsertOrGetMul() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
				}
				row := queryRowContextInTx(tx, ctx, tabInfo, where)
				err := scanRow(addr, tabInfo, row)
				if err != nil {
					return err
				}
				return nil
			}
			return err
		}
		setInc(addr, tabInfo, lid)
		return nil
	})
}

func BulkInsert(slice table) error {
	addr := getTabAddr(slice)
	tabInfo := getTabInfo(slice)
	lastInsertID, err := bulkInsert(addr, tabInfo)
	if err != nil {
		return err
	}
	l := *(*[]uintptr)(unsafe.Pointer(addr))
	for i, modAddr := range l[1:] {
		setSync(modAddr, tabInfo)
		setInc(modAddr, tabInfo, lastInsertID+int64(i))
	}
	return nil
}

//UpdateOne update one record
func UpdateOne(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where, typ := genWhere(modAddr, tabInfo)
	if typ == otherType {
		return fmt.Errorf("nborm.UpdateOne() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
	}
	updVals := genUpdVals(modAddr, tabInfo)
	_, err := update(tabInfo, where, updVals...)
	if err != nil {
		return err
	}
	setSync(modAddr, tabInfo)
	return nil
}

func UpdateOneInTx(tx *sql.Tx, model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where, typ := genWhere(modAddr, tabInfo)
	if typ == otherType {
		return fmt.Errorf("nborm.UpdateOne() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
	}
	updVals := genUpdVals(modAddr, tabInfo)
	_, err := updateInTx(tx, tabInfo, where, updVals...)
	if err != nil {
		return err
	}
	setSync(modAddr, tabInfo)
	return nil
}

//UpdateMul update multiple records
func UpdateMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		updVals := genUpdVals(addr, tabInfo)
		where, typ := genWhere(addr, tabInfo)
		if typ == otherType {
			return fmt.Errorf("nborm.UpdateMul() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
		}
		_, err := update(tabInfo, where, updVals...)
		if err != nil {
			return err
		}
		setSync(addr, tabInfo)
		return nil
	})
}

func UpdateMulInTx(tx *sql.Tx, slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		updVals := genUpdVals(addr, tabInfo)
		where, typ := genWhere(addr, tabInfo)
		if typ == otherType {
			return fmt.Errorf("nborm.UpdateMul() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
		}
		_, err := updateInTx(tx, tabInfo, where, updVals...)
		if err != nil {
			return err
		}
		setSync(addr, tabInfo)
		return nil
	})
}

//BulkUpdate update records by where
func BulkUpdate(model table, where *Where, updVals ...*UpdateValue) error {
	tabInfo := getTabInfo(model)
	_, err := update(tabInfo, where, updVals...)
	return err
}

func BulkUpdateInTx(tx *sql.Tx, model table, where *Where, updVals ...*UpdateValue) error {
	tabInfo := getTabInfo(model)
	_, err := updateInTx(tx, tabInfo, where, updVals...)
	return err
}

//DeleteOne delete one record
func DeleteOne(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where, typ := genWhere(modAddr, tabInfo)
	if typ == otherType {
		return fmt.Errorf("nborm.DeleteOne() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
	}
	if _, err := delete(tabInfo, where); err != nil {
		return err
	}
	unsetSync(modAddr, tabInfo)
	return nil
}

func DeleteOneInTx(tx *sql.Tx, model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where, typ := genWhere(modAddr, tabInfo)
	if typ == otherType {
		return fmt.Errorf("nborm.DeleteOne() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
	}
	if _, err := deleteInTx(tx, tabInfo, where); err != nil {
		return err
	}
	unsetSync(modAddr, tabInfo)
	return nil
}

//DeleteMul delete multiple records
func DeleteMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		where, typ := genWhere(addr, tabInfo)
		if typ == otherType {
			return fmt.Errorf("nborm.DeleteMul() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
		}
		if _, err := delete(tabInfo, where); err != nil {
			return err
		}
		unsetSync(addr, tabInfo)
		return nil
	})
}

func DeleteMulInTx(tx *sql.Tx, slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		where, typ := genWhere(addr, tabInfo)
		if typ == otherType {
			return fmt.Errorf("nborm.DeleteMul() error: no valid unique field value for locating record (%s.%s)", tabInfo.dbName(), tabInfo.tabName())
		}
		if _, err := deleteInTx(tx, tabInfo, where); err != nil {
			return err
		}
		unsetSync(addr, tabInfo)
		return nil
	})
}

//BulkDelete delete by where
func BulkDelete(model table, where *Where) error {
	tabInfo := getTabInfo(model)
	_, err := delete(tabInfo, where)
	return err
}

func BulkDeleteInTx(tx *sql.Tx, model table, where *Where) error {
	tabInfo := getTabInfo(model)
	_, err := deleteInTx(tx, tabInfo, where)
	return err
}

//DeleteAll delete all records from one table
func DeleteAll(table table) error {
	tabInfo := getTabInfo(table)
	_, err := deleteAll(tabInfo)
	return err
}

func DeleteAllInTx(tx *sql.Tx, table table) error {
	tabInfo := getTabInfo(table)
	_, err := deleteAllInTx(tx, tabInfo)
	return err
}

func TruncateTable(table table) error {
	tabInfo := getTabInfo(table)
	_, err := truncateTable(tabInfo)
	return err
}

func TruncateTableInTx(tx *sql.Tx, table table) error {
	tabInfo := getTabInfo(table)
	_, err := truncateTableInTx(tx, tabInfo)
	return err
}

//Count get the number of rows in one table
func Count(model table, where *Where) (int, error) {
	tabInfo := getTabInfo(model)
	return count(tabInfo, where)
}

func CountInTx(tx *sql.Tx, model table, where *Where) (int, error) {
	tabInfo := getTabInfo(model)
	return countInTx(tx, tabInfo, where)
}

//Sort sort ModelList
func Sort(slice table, funcs ...func(iaddr, jaddr uintptr) int) {
	o := &sorter{slice, funcs}
	sort.Sort(o)
}

//Distinct distinct Models in a ModelList by selected Fields
func Distinct(slice table, fields ...Field) {
	tabInfo := getTabInfo(slice)
	distMap := make(map[string]bool)
	f := func(addr uintptr) bool {
		builder := strings.Builder{}
		for _, field := range fields {
			builder.WriteString(fmt.Sprintf("%v", getFieldByName(addr, field.columnName(), tabInfo).value()))
		}
		id := builder.String()
		if distMap[id] {
			return true
		}
		distMap[id] = true
		return false
	}
	filterList(slice, f)
}
