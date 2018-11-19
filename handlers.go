package nborm

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"

	"github.com/go-sql-driver/mysql"
)

//First get the first record in database, no error return when no record in table, check Model synchronized status after query
func First(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where := genWhere(modAddr, tabInfo)
	row := queryRow(tabInfo, where)
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
	where := genWhere(modAddr, tabInfo)
	row := queryRow(tabInfo, where)
	return scanRow(modAddr, tabInfo, row)
}

//GetMul get multiple Models by Models's owned field value, if one of them not has conresponse record, it will return a sql.ErrNoRows error
func GetMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		where := genWhere(addr, tabInfo)
		row := queryRowContext(ctx, tabInfo, where)
		return scanRow(addr, tabInfo, row)
	})
}

//JoinQueryOne query one record by join tables, no error return when no conresponse record, check Model synchronized status after query
func JoinQueryOne(model table, where *Where, relations ...relation) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	row := joinQueryRow(tabInfo, where, relations...)
	return scanRow(modAddr, tabInfo, row)
}

//JoinQuery join query records by where
func JoinQuery(slice table, where *Where, sorter *Sorter, pager *Pager, relations ...relation) error {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, err := joinQueryRows(tabInfo, where, sorter, pager, relations...)
	if err != nil {
		return err
	}
	return scanRows(sliceAddr, tabInfo, rows)
}

//JoinQueryWithFoundRows join query records and get the number of found rows
func JoinQueryWithFoundRows(slice table, where *Where, sorter *Sorter, pager *Pager, relations ...relation) (int, error) {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, numRows, tx, err := joinQueryRowsAndFoundRows(tabInfo, where, sorter, pager, relations...)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return numRows, nil
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

//AllWithFoundRows get all records of one table and the number of records
func AllWithFoundRows(slice table, sorter *Sorter, pager *Pager) (int, error) {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, numRows, tx, err := queryRowsAndFoundRows(tabInfo, nil, sorter, pager)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		return -1, err
	}
	tx.Commit()
	return numRows, nil

}

//QueryOne query one record, no error will return when no conresponse record, check the synchronized status of Model after query
func QueryOne(model table, where *Where) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	row := queryRow(tabInfo, where)
	return scanRow(modAddr, tabInfo, row)
}

//QueryMul query multiple Models, no error will be returned if no conresponse record, check synchronized status of Models after query
func QueryMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		where := genWhere(addr, tabInfo)
		row := queryRowContext(ctx, tabInfo, where)
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

//QueryWithFoundRows query records by where and get the number of found rows
func QueryWithFoundRows(slice table, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	tabInfo := getTabInfo(slice)
	sliceAddr := getTabAddr(slice)
	rows, numRows, tx, err := queryRowsAndFoundRows(tabInfo, where, sorter, pager)
	if err != nil {
		return -1, err
	}
	if err := scanRows(sliceAddr, tabInfo, rows); err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return numRows, nil
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

//InsertOrUpdateOne insert or update one record
func InsertOrUpdateOne(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	lid, err := insertOrUpdate(modAddr, tabInfo)
	if err != nil {
		return err
	}
	setInc(modAddr, tabInfo, lid)
	return nil
}

//InsertMul insert multiple record
func InsertMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		lid, err := insertContext(ctx, addr, tabInfo)
		if err != nil {
			return err
		}
		inc := getIncWithTableInfo(addr, tabInfo)
		inc.setVal(lid, false)
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
		inc := getIncWithTableInfo(addr, tabInfo)
		inc.setVal(lid, false)
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
			where := genWhere(modAddr, tabInfo)
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
	inc.setVal(lid, false)
	return nil
}

//InsertOrGetMul insert multiple records or get them when they are already exist
func InsertOrGetMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		lid, err := insertContext(ctx, addr, tabInfo)
		if err != nil {
			if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
				where := genWhere(addr, tabInfo)
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

//UpdateOne update one record
func UpdateOne(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where := genWhere(modAddr, tabInfo)
	updVals := genUpdVals(modAddr, tabInfo)
	_, err := update(tabInfo, where, updVals...)
	return err
}

//UpdateMul update multiple records
func UpdateMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		updVals := genUpdVals(addr, tabInfo)
		where := genWhere(addr, tabInfo)
		_, err := update(tabInfo, where, updVals...)
		return err
	})
}

//BulkUpdate update records by where
func BulkUpdate(table table, where *Where, updVals ...*UpdateValue) error {
	tabInfo := getTabInfo(table)
	_, err := update(tabInfo, where, updVals...)
	return err
}

//DeleteOne delete one record
func DeleteOne(model table) error {
	tabInfo := getTabInfo(model)
	modAddr := getTabAddr(model)
	where := genWhere(modAddr, tabInfo)
	if _, err := delete(tabInfo, where); err != nil {
		return err
	}
	unsetSync(modAddr, tabInfo)
	return nil
}

//DeleteMul delete multiple records
func DeleteMul(slice table) error {
	tabInfo := getTabInfo(slice)
	return iterList(slice, func(ctx context.Context, addr uintptr) error {
		where := genWhere(addr, tabInfo)
		if _, err := delete(tabInfo, where); err != nil {
			return err
		}
		unsetSync(addr, tabInfo)
		return nil
	})
}

//BulkDelete delete by where
func BulkDelete(table table, where *Where) error {
	tabInfo := getTabInfo(table)
	_, err := delete(tabInfo, where)
	return err
}

//DeleteAll delete all records from one table
func DeleteAll(table table) error {
	tabInfo := getTabInfo(table)
	_, err := truncateTable(tabInfo)
	return err
}

//Count get the number of rows in one table
func Count(table table, where *Where) (int, error) {
	tabInfo := getTabInfo(table)
	return count(tabInfo, where)
}

//Sort sort ModelList
func Sort(slice table, reverse bool, funcs ...func(iaddr, jaddr uintptr) int) {
	o := &sorter{slice, funcs}
	if reverse {
		sort.Sort(sort.Reverse(o))
	} else {
		sort.Sort(o)
	}
}

//Distinct distinct Models in a ModelList by selected Fields
func Distinct(slice table, fields ...Field) {
	tabInfo := getTabInfoByName(slice.DB(), slice.Tab())
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
