package nborm

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/wangjun861205/nbcolor"
)

func InsertOne(exe Executor, model Model) error {
	validFields := getFields(model, valid)
	cl := make([]string, 0, len(validFields))
	pl := make([]string, 0, len(validFields))
	vl := make([]interface{}, 0, len(validFields))
	for _, f := range validFields {
		toInsert(f, &cl, &pl, &vl)
	}
	stmt := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, model.rawFullTabName(), strings.Join(cl, ", "), strings.Join(pl, ", "))
	if DEBUG {
		fmt.Println(nbcolor.Green(stmt))
		fmt.Println(nbcolor.Green(vl))
	}
	res, err := exe.Exec(stmt, vl...)
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
	model.setModelStatus(synced)
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

func FullCount(exe Executor, model Model) (int, error) {
	whereClause, whereValues := genFullWhereClause(model)
	tabRef := genFullTabRef(model)
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

func QueryOne(exe Executor, model Model) error {
	// selectFields := getFields(model, forSelect)
	selectColumns := genSelectColumns(model)
	whereClause, whereValues := genWhereClause(model)
	tabRef := genTabRef(model)
	stmt := fmt.Sprintf("SELECT %s FROM %s %s", selectColumns, tabRef, whereClause)
	if DEBUG {
		log.Println(nbcolor.Green(stmt))
		log.Println(nbcolor.Green(whereValues))
	}
	err := scanRow(exe.QueryRow(stmt, whereValues...), model)
	if err != nil {
		return err
	}
	return nil
}

func FullQueryOne(exe Executor, model Model) error {
	// selectFields := getFields(model, forSelect)
	selectColumns := genSelectColumns(model)
	whereClause, whereValues := genFullWhereClause(model)
	tabRef := genFullTabRef(model)
	stmt := fmt.Sprintf("SELECT %s FROM %s %s", selectColumns, tabRef, whereClause)
	if DEBUG {
		log.Println(nbcolor.Green(stmt))
		log.Println(nbcolor.Green(whereValues))
	}
	err := scanRow(exe.QueryRow(stmt, whereValues...), model)
	if err != nil {
		return err
	}
	return nil
}

func Query(exe Executor, l ModelList) error {
	// selectFields := getFields(l, forSelect)
	selectColumns := genSelectColumns(l)
	whereClause, whereValues := genWhereClause(l)
	tabRef := genTabRef(l)
	orderClause := genOrderClause(l)
	limitClause := genLimitClause(l)
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s FROM %s %s %s %s", selectColumns, tabRef, whereClause, orderClause, limitClause)
	if DEBUG {
		log.Println(nbcolor.Green(stmt))
		log.Println(nbcolor.Green(whereValues))
	}
	rows, err := exe.Query(stmt, whereValues...)
	if err != nil {
		return err
	}
	// defer rows.Close()
	// for rows.Next() {
	// 	if err := scanRows(rows, l.NewModel(), selectFields...); err != nil {
	// 		return err
	// 	}
	// }
	// if err := rows.Err(); err != nil {
	// 	return err
	// }
	if err := scanRows(rows, l); err != nil {
		return err
	}
	stmt = `SELECT FOUND_ROWS()`
	var total int
	if err := exe.QueryRow(stmt).Scan(&total); err != nil {
		return err
	}
	l.SetTotal(total)
	l.setModelStatus(synced)
	return nil
}

func FullQuery(exe Executor, l ModelList) error {
	// selectFields := getFields(l, forSelect)
	selectColumns := genSelectColumns(l)
	whereClause, whereValues := genFullWhereClause(l)
	tabRef := genFullTabRef(l)
	orderClause := genFullOrderClause(l)
	limitClause := genLimitClause(l)
	stmt := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s FROM %s %s %s %s", selectColumns, tabRef, whereClause, orderClause, limitClause)
	if DEBUG {
		log.Println(nbcolor.Green(stmt))
		log.Println(nbcolor.Green(whereValues))
	}
	rows, err := exe.Query(stmt, whereValues...)
	if err != nil {
		return err
	}
	// defer rows.Close()
	// for rows.Next() {
	// 	if err := scanRows(rows, l.NewModel(), selectFields...); err != nil {
	// 		return err
	// 	}
	// }
	// if err := rows.Err(); err != nil {
	// 	return err
	// }
	if err := scanRows(rows, l); err != nil {
		return err
	}
	stmt = `SELECT FOUND_ROWS()`
	var total int
	if err := exe.QueryRow(stmt).Scan(&total); err != nil {
		return err
	}
	l.SetTotal(total)
	l.setModelStatus(synced)
	return nil
}

func JoinQuery(exe Executor, model Model) error {
	selectColumns := genJoinSelectColumns(model)
	tabRef := genJoinTabRef(model)
	whereClause, whereValues := genWhereClause(model)
	limitClause := genLimitClause(model)
	orderClause := genOrderClause(model)
	var stmt string
	if _, ok := model.(ModelList); ok {
		stmt = fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s FROM %s %s %s %s", selectColumns, tabRef, whereClause, orderClause, limitClause)
	} else {
		stmt = fmt.Sprintf("SELECT %s FROM %s %s %s %s", selectColumns, tabRef, whereClause, orderClause, limitClause)
	}
	if DEBUG {
		fmt.Println(nbcolor.Green(stmt))
		fmt.Println(nbcolor.Green(whereValues))
	}
	if l, ok := model.(ModelList); ok {
		rows, err := exe.Query(stmt, whereValues...)
		if err != nil {
			return err
		}
		if err := joinScanRows(rows, l); err != nil {
			return err
		}
		stmt = `SELECT FOUND_ROWS()`
		var total int
		if err := exe.QueryRow(stmt).Scan(&total); err != nil {
			return err
		}
		l.SetTotal(total)
		l.setModelStatus(synced)
		return nil
	}
	return joinScanRow(exe.QueryRow(stmt, whereValues...), model)

}

func Update(exe Executor, model Model) (sql.Result, error) {
	tabRef := genTabRef(model)
	updateClause, updateValues := genUpdateSetClause(model)
	whereClause, whereValues := genWhereClause(model)
	stmt := fmt.Sprintf(`UPDATE %s %s %s`, tabRef, updateClause, whereClause)
	if DEBUG {
		fmt.Println(nbcolor.Green(stmt))
		fmt.Println(nbcolor.Green(updateValues))
		fmt.Println(nbcolor.Green(whereValues))
	}
	return exe.Exec(stmt, append(updateValues, whereValues...)...)
}

func FullUpdate(exe Executor, model Model) (sql.Result, error) {
	tabRef := genFullTabRef(model)
	updateClause, updateValues := genFullUpdateSetClause(model)
	whereClause, whereValues := genFullWhereClause(model)
	stmt := fmt.Sprintf("UPDATE %s %s %s", tabRef, updateClause, whereClause)
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
