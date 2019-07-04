package nborm

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/wangjun861205/nbcolor"
)

func InsertOrUpdateOne(exe Executor, model Model) (isInsert bool, err error) {
	insertClause, insertValues := genInsertClause(model)
	updateClause, updateValues := genSimpleUpdateClause(model)
	stmt := fmt.Sprintf("INSERT INTO %s %s ON DUPLICATE KEY UPDATE %s", model.rawFullTabName(), insertClause, updateClause)
	if DEBUG {
		fmt.Println(nbcolor.Green(stmt))
		fmt.Println(nbcolor.Green(insertValues))
		fmt.Println(nbcolor.Green(updateValues))
	}
	res, err := exe.Exec(stmt, append(insertValues, updateValues...)...)
	if err != nil {
		return
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affectedRows == 1 {
		isInsert = true
	}
	if affectedRows == 1 || affectedRows == 2 {
		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return false, err
		}
		model.AutoIncField().Set(int(lastInsertID))
	}
	return
}

func InsertOne(exe Executor, model Model) error {
	insertClause, insertValues := genInsertClause(model)
	stmt := fmt.Sprintf(`INSERT INTO %s %s`, model.rawFullTabName(), insertClause)
	if DEBUG {
		fmt.Println(nbcolor.Green(stmt))
		fmt.Println(nbcolor.Green(insertValues))
	}
	res, err := exe.Exec(stmt, insertValues...)
	if err != nil {
		return err
	}
	if model.AutoIncField() != nil {
		lid, err := res.LastInsertId()
		if err != nil {
			return err
		}
		model.AutoIncField().Set(int(lid))
	}
	model.addModelStatus(synced)
	return nil
}

func Count(exe Executor, model Model) (int, error) {
	whereClause, whereValues := genWhereClause(model)
	tabRef := genTabRef(model)
	stmt := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", tabRef, whereClause)
	if DEBUG {
		fmt.Println(nbcolor.Green(stmt))
		fmt.Println(nbcolor.Green(whereValues))
	}
	var count int
	if err := exe.QueryRow(stmt, whereValues...).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func Query(exe Executor, m Model) error {
	selectClause := genSelectClause(m)
	whereClause, whereValues := genWhereClause(m)
	tabRef := genTabRef(m)
	orderClause := genOrderClause(m)
	limitClause := genLimitClause(m)
	stmt := fmt.Sprintf("%s FROM %s %s %s %s", selectClause, tabRef, whereClause, orderClause, limitClause)
	if DEBUG {
		log.Println(nbcolor.Green(stmt))
		log.Println(nbcolor.Green(whereValues))
	}
	return queryAndScan(exe, m, stmt, whereValues...)
}

func Update(exe Executor, model Model) (sql.Result, error) {
	tabRef := genTabRef(model)
	updateClause, updateValues := genUpdateClause(model)
	whereClause, whereValues := genWhereClause(model)
	stmt := fmt.Sprintf(`UPDATE %s SET %s %s`, tabRef, updateClause, whereClause)
	if DEBUG {
		fmt.Println(nbcolor.Green(stmt))
		fmt.Println(nbcolor.Green(updateValues))
		fmt.Println(nbcolor.Green(whereValues))
	}
	return exe.Exec(stmt, append(updateValues, whereValues...)...)
}

func Delete(exe Executor, model Model) (sql.Result, error) {
	whereClause, whereValues := genSimpleWhereClause(model)
	stmt := fmt.Sprintf("DELETE FROM %s %s", model.rawFullTabName(), whereClause)
	if DEBUG {
		fmt.Println(nbcolor.Green(stmt))
		fmt.Println(nbcolor.Green(whereValues))
	}
	return exe.Exec(stmt, whereValues...)
}

func Agg(exe Executor, model Model) (AggResult, error) {
	selectClause, tempFields := genAggSelectClause(model)
	tabRef := genTabRef(model)
	whereClause, whereValues := genWhereClause(model)
	groupByClause := genGroupByClause(model)
	havingClause, havingValues := genHavingClause(model)
	stmt := fmt.Sprintf("%s FROM %s %s %s %s", selectClause, tabRef, whereClause, groupByClause, havingClause)
	if DEBUG {
		fmt.Println(nbcolor.Green(stmt))
		fmt.Println(nbcolor.Green(whereValues))
	}
	result := make(AggResult, 0, 64)
	rows, err := exe.Query(stmt, append(whereValues, havingValues...)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		nl := make(FieldList, 0, len(tempFields))
		al := make([]interface{}, 0, len(tempFields))
		for i := 0; i < len(tempFields); i++ {
			nf := tempFields[i].dup()
			nl = append(nl, nf)
			al = append(al, nf)
		}
		if err := rows.Scan(al...); err != nil {
			return nil, err
		}
		result = append(result, nl)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
