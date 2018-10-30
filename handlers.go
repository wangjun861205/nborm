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
func First(m Model) error {
	stmt, valList := genSelect(m, nil, nil, nil, false)
	err := queryAndScan(m, stmt, valList)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

//GetOne get one record by Model owned field value, if no record, it will return a sql.ErrNoRows error
func GetOne(m Model) error {
	stmt, valList := genSelect(m, genWhere(m), nil, nil, false)
	return queryAndScan(m, stmt, valList)
}

//GetMul get multiple Models by Models's owned field value, if one of them not has conresponse record, it will return a sql.ErrNoRows error
func GetMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error {
		stmt, valList := genSelect(m, genWhere(m), nil, nil, false)
		return queryAndScanContext(ctx, m, stmt, valList)
	})
}

//JoinQueryOne query one record by join tables, no error return when no conresponse record, check Model synchronized status after query
func JoinQueryOne(m Model, where *Where, relations ...relation) error {
	stmt, valList := genSelect(m, where, nil, nil, false, relations...)
	err := queryAndScan(m, stmt, valList)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

//All get all records of one table
func All(l ModelList, sorter *Sorter, pager *Pager) error {
	stmt, valList := genSelect(l, nil, sorter, pager, false)
	return queryAndScan(l, stmt, valList)
}

//AllWithFoundRows get all records of one table and the number of records
func AllWithFoundRows(l ModelList, sorter *Sorter, pager *Pager) (int, error) {
	stmt, valList := genSelect(l, nil, sorter, pager, true)
	return queryAndScanWithNum(l, stmt, valList)
}

//QueryOne query one record, no error will return when no conresponse record, check the synchronized status of Model after query
func QueryOne(m Model, where *Where) error {
	stmt, valList := genSelect(m, where, nil, nil, false)
	err := queryAndScan(m, stmt, valList)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

//QueryMul query multiple Models, no error will be returned if no conresponse record, check synchronized status of Models after query
func QueryMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error {
		stmt, valList := genSelect(m, genWhere(m), nil, nil, false)
		err := queryAndScanContext(ctx, m, stmt, valList)
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	})
}

//Query query records by where
func Query(l ModelList, where *Where, sorter *Sorter, pager *Pager) error {
	stmt, valList := genSelect(l, where, sorter, pager, false)
	return queryAndScan(l, stmt, valList)
}

//QueryWithFoundRows query records by where and get the number of found rows
func QueryWithFoundRows(l ModelList, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	stmt, valList := genSelect(l, where, sorter, pager, true)
	return queryAndScanWithNum(l, stmt, valList)
}

//JoinQuery join query records by where
func JoinQuery(l ModelList, where *Where, sorter *Sorter, pager *Pager, relations ...relation) error {
	stmt, valList := genSelect(l, where, sorter, pager, false, relations...)
	return queryAndScan(l, stmt, valList)
}

//JoinQueryWithFoundRows join query records and get the number of found rows
func JoinQueryWithFoundRows(l ModelList, where *Where, sorter *Sorter, pager *Pager, relations ...relation) (int, error) {
	stmt, valList := genSelect(l, where, sorter, pager, true, relations...)
	return queryAndScanWithNum(l, stmt, valList)
}

//InsertOne insert one record
func InsertOne(m Model) error {
	return insertAndGetInc(m, false)
}

//InsertOrUpdateOne insert or update one record
func InsertOrUpdateOne(m Model) error {
	return insertAndGetInc(m, true)
}

//InsertMul insert multiple record
func InsertMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error { return insertAndGetIncContext(ctx, m, false) })
}

//InsertOrUpdateMul insert or update multiple record
func InsertOrUpdateMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error { return insertAndGetIncContext(ctx, m, true) })
}

//InsertOrGetOne insert one record or get it when it is already exists
func InsertOrGetOne(m Model) error {
	err := insertAndGetInc(m, false)
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
			stmt, valList := genSelect(m, genWhere(m), nil, nil, false)
			return queryAndScan(m, stmt, valList)
		}
		return err
	}
	return nil
}

//InsertOrGetMul insert multiple records or get them when they are already exist
func InsertOrGetMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error {
		err := insertAndGetIncContext(ctx, m, false)
		if err != nil {
			if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
				stmt, valList := genSelect(m, genWhere(m), nil, nil, false)
				return queryAndScanContext(ctx, m, stmt, valList)
			}
			return err
		}
		return nil
	})
}

//UpdateOne update one record
func UpdateOne(m Model) error {
	_, fs := getInc(m)
	fs = filterValid(fs)
	setValues := make([]string, 0, len(fs))
	for _, f := range fs {
		setValues = append(setValues, f.UpdateValue().String())
	}
	colStr, valList := genWhere(m).toSQL()
	stmtStr := fmt.Sprintf("UPDATE %s.%s SET %s %s", m.DB(), m.Tab(), strings.Join(setValues, ", "), colStr)
	db := dbMap[m.DB()]
	_, err := db.Exec(stmtStr, valList)
	return err
}

//UpdateMul update multiple records
func UpdateMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, fs := getInc(m)
			fs = filterValid(fs)
			setValues := make([]string, 0, len(fs))
			for _, f := range fs {
				setValues = append(setValues, f.UpdateValue().String())
			}
			colStr, valList := genWhere(m).toSQL()
			stmtStr := fmt.Sprintf("UPDATE %s.%s SET %s %s", m.DB(), m.Tab(), strings.Join(setValues, ", "), colStr)
			db := dbMap[m.DB()]
			_, err := db.ExecContext(ctx, stmtStr, valList)
			return err
		}
	})
}

//BulkUpdate update records by where
func BulkUpdate(m Model, where *Where, values ...*UpdateValue) error {
	db := dbMap[m.DB()]
	setList := make([]string, len(values))
	for i, val := range values {
		setList[i] = val.String()
	}
	colStr, valList := where.toClause()
	stmtStr := fmt.Sprintf("UPDATE %s.%s SET %s %s", m.DB(), m.Tab(), strings.Join(setList, ", "), colStr)
	_, err := db.Exec(stmtStr, valList)
	return err
}

//DeleteOne delete one record
func DeleteOne(m Model) error {
	db := dbMap[m.DB()]
	colStr, valList := genWhere(m).toSQL()
	stmtStr := fmt.Sprintf("DELETE FROM %s.%s %s", m.DB(), m.Tab(), colStr)
	_, err := db.Exec(stmtStr, valList)
	if err != nil {
		return err
	}
	m.SetSync(false)
	return nil
}

//DeleteMul delete multiple records
func DeleteMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error {
		db := dbMap[m.DB()]
		colStr, valList := genWhere(m).toSQL()
		stmtStr := fmt.Sprintf("DELETE FROM %s.%s %s", m.DB(), m.Tab(), colStr)
		_, err := db.ExecContext(ctx, stmtStr, valList)
		if err != nil {
			return err
		}
		m.SetSync(false)
		return nil
	})
}

//BulkDelete delete by where
func BulkDelete(m Model, where *Where) error {
	db := dbMap[m.DB()]
	colStr, valList := where.toClause()
	stmtStr := fmt.Sprintf("DELETE FROM %s.%s %s", m.DB(), m.Tab(), colStr)
	_, err := db.Exec(stmtStr, valList)
	return err
}

//DeleteAll delete all records from one table
func DeleteAll(m Model) error {
	db := dbMap[m.DB()]
	stmtStr := fmt.Sprintf("TRUNCATE TABLE %s.%s", m.DB(), m.Tab())
	_, err := db.Exec(stmtStr)
	return err
}

//Count get the number of rows in one table
func Count(m Model, where *Where) (int, error) {
	db := dbMap[m.DB()]
	var stmtStr string
	colStr, valList := where.toClause()
	stmtStr = fmt.Sprintf("SELECT COUNT(*) FROM %s.%s %s", m.DB(), m.Tab(), colStr)
	var num int
	row := db.QueryRow(stmtStr, valList)
	err := row.Scan(&num)
	if err != nil {
		return -1, err
	}
	return num, nil

}

//Sort sort ModelList
func Sort(l ModelList, reverse bool, fields ...Field) {
	funcs := make([]func(Model, Model) int, len(fields))
	for i, f := range fields {
		funcs[i] = f.LessFunc()
	}
	o := &sortObj{l, funcs}
	if reverse {
		sort.Sort(sort.Reverse(o))
	} else {
		sort.Sort(o)
	}
}

//Distinct distinct Models in a ModelList by selected Fields
func Distinct(l ModelList, fields ...Field) {
	distMap := make(map[string]bool)
	f := func(m Model) bool {
		builder := strings.Builder{}
		for _, field := range fields {
			builder.WriteString(getByName(m, field.Column()).SQLVal())
		}
		id := builder.String()
		if distMap[id] {
			return true
		}
		distMap[id] = true
		return false
	}
	filterList(l, f)
}
