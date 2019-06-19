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
	stmt := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, model.fullTabName(), strings.Join(cl, ", "), strings.Join(pl, ", "))
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
	return nil
}

func Count(exe Executor, model Model) (int, error) {
	fields := getFields(model, forWhere)
	whereList := make(whereList, 0, len(fields)*2)
	for _, f := range fields {
		whereList = append(whereList, f.whereList()...)
	}
	clause, values := whereList.toClause()
	stmt := fmt.Sprintf("SELECT COUNT(*) FROM %s %s", model.fullTabName(), clause)
	var count int
	if err := exe.QueryRow(stmt, values...).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func QueryOne(exe Executor, model Model) error {
	var whereClause string
	var whereValues []interface{}
	selectColumns := getSelectColumns(model)
	selectFields := getFields(model, forSelect)
	var stmt string
	if model.getRelJoin() != "" {
		whereList := make(whereList, 0, 8)
		whereList = append(whereList, model.getRelWhere())
		whereList = append(whereList, genWhereList(model)...)
		whereClause, whereValues = whereList.toClause()
		stmt = fmt.Sprintf("SELECT %s FROM %s %s", selectColumns, model.getRelJoin(), whereClause)
	} else {
		whereClause, whereValues = genWhereClause(model)
		stmt = fmt.Sprintf("SELECT %s FROM %s %s", selectColumns, model.fullTabName(), whereClause)
	}
	if DEBUG {
		log.Println(nbcolor.Green(stmt))
		log.Println(nbcolor.Green(whereValues))
	}
	err := scanRow(exe.QueryRow(stmt, whereValues...), model, selectFields...)
	if err != nil {
		return err
	}
	initRelation(model)
	return nil
}

func Query(exe Executor, l ModelList, limit, offset int) error {
	var whereClause string
	var whereValues []interface{}
	var stmt string
	selectFields := getFields(l, forSelect)
	selectColumns := getSelectColumns(l)
	if l.getRelJoin() != "" {
		whereList := make(whereList, 0, 8)
		whereList = append(whereList, l.getRelWhere())
		whereList = append(whereList, genWhereList(l)...)
		whereClause, whereValues = whereList.toClause()
		stmt = fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s FROM %s %s", selectColumns, l.getRelJoin(), whereClause)
	} else {
		whereClause, whereValues = genWhereClause(l)
		stmt = fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s FROM %s %s", selectColumns, l.fullTabName(), whereClause)
	}
	if limit > 0 && offset >= 0 {
		stmt = fmt.Sprintf("%s LIMIT %d, %d", stmt, offset, limit)
	}
	if DEBUG {
		log.Println(nbcolor.Green(stmt))
		log.Println(nbcolor.Green(whereValues))
	}
	rows, err := exe.Query(stmt, whereValues...)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		if err := scanRows(rows, l.NewModel(), selectFields...); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	stmt = `SELECT FOUND_ROWS()`
	var total int
	if err := exe.QueryRow(stmt).Scan(&total); err != nil {
		return err
	}
	l.SetTotal(total)
	initRelation(l)
	l.setModelStatus(synced)
	return nil
}

func Update(exe Executor, model Model) (sql.Result, error) {
	updateClause, updateValues := genUpdateSetClause(model)
	whereClause, whereValues := genWhereClause(model)
	stmt := fmt.Sprintf(`UPDATE %%s %s %s`, model.fullTabName(), updateClause, whereClause)
	return exe.Exec(stmt, append(updateValues, whereValues...))
}
