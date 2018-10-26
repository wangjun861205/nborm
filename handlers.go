package nborm

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func First(m Model) error {
	return queryAndScan(m, genSelect(m, nil, nil, nil, false))
}

func GetOne(m Model) (exists bool, err error) {
	err = queryAndScan(m, genSelect(m, genWhere(m), nil, nil, false))
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	exists = true
	return
}

func GetMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error {
		return queryAndScanContext(ctx, m, genSelect(m, genWhere(m), nil, nil, false))
	})
}

func JoinQueryOne(m Model, where *Where, relations ...relation) (exists bool, err error) {
	err = queryAndScan(m, genSelect(m, where, nil, nil, false, relations...))
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	exists = true
	return
}

func All(l ModelList, sorter *Sorter, pager *Pager) error {
	return queryAndScan(l, genSelect(l, nil, sorter, pager, false))
}

func AllWithFoundRows(l ModelList, sorter *Sorter, pager *Pager) (int, error) {
	return queryAndScanWithNum(l, genSelect(l, nil, sorter, pager, true))
}

func QueryOne(m Model, where *Where) error {
	err := queryAndScan(m, genSelect(m, where, nil, nil, false))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	return nil
}

func QueryMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error {
		err := queryAndScanContext(ctx, m, genSelect(m, genWhere(m), nil, nil, false))
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return err
		}
		return nil
	})
}

func Query(l ModelList, where *Where, sorter *Sorter, pager *Pager) error {
	return queryAndScan(l, genSelect(l, where, sorter, pager, false))
}

func QueryWithFoundRows(l ModelList, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	return queryAndScanWithNum(l, genSelect(l, where, sorter, pager, true))
}

func JoinQuery(l ModelList, where *Where, sorter *Sorter, pager *Pager, relations ...relation) error {
	return queryAndScan(l, genSelect(l, where, sorter, pager, false, relations...))
}

func JoinQueryWithFoundRows(l ModelList, where *Where, sorter *Sorter, pager *Pager, relations ...relation) (int, error) {
	return queryAndScanWithNum(l, genSelect(l, where, sorter, pager, true, relations...))
}

// func Query(tab interface{}, where *Where, sorter *Sorter, pager *Pager) error {
// 	if where == nil {
// 		switch obj := tab.(type) {
// 		case Model:
// 			where = genWhere(obj)
// 			stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", obj.Tab(), where.String())
// 			db := dbMap[obj.DB()]
// 			row := db.QueryRow(stmtStr)
// 			return scanRow(obj, row)
// 		case ModelList:
// 			wg := sync.WaitGroup{}
// 			doneChan := make(chan interface{})
// 			errChan := make(chan error)
// 			for i := 0; i < obj.Len(); i++ {
// 				wg.Add(1)
// 				go func(index int) {
// 					defer func() {
// 						recover()
// 						wg.Done()
// 					}()
// 					w := genWhere(obj.Index(index))
// 					stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", obj.Model().Tab(), w.String())
// 					db := dbMap[obj.Model().DB()]
// 					row := db.QueryRow(stmtStr)
// 					if err := scanRow(obj.Index(index), row); err != nil {
// 						errChan <- err
// 					}
// 				}(i)
// 			}
// 			go func() {
// 				wg.Wait()
// 				close(doneChan)
// 			}()
// 			select {
// 			case err := <-errChan:
// 				close(errChan)
// 				return err
// 			case <-doneChan:
// 				return nil
// 			}
// 		}
// 	}
// 	switch obj := tab.(type) {
// 	case Model:
// 		stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s %s %s", obj.Tab(), where.String(), sorter.toSQL(), pager.toSQL())
// 		db := dbMap[obj.DB()]
// 		row := db.QueryRow(stmtStr)
// 		return scanRow(obj, row)
// 	case ModelList:
// 		stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s %s %s", obj.Model().Tab(), where.String(), sorter.toSQL(), pager.toSQL())
// 		db := dbMap[obj.Model().DB()]
// 		rows, err := db.Query(stmtStr)
// 		if err != nil {
// 			return err
// 		}
// 		return scanRows(obj, rows)
// 	default:
// 		panic("nborm.Query() error: unsupported type")
// 	}
// }

// func QueryWithFoundRows(tab interface{}, where *Where, sorter *Sorter, pager *Pager) (int, error) {
// 	if where == nil {
// 		switch obj := tab.(type) {
// 		case Model:
// 			where = genWhere(obj)
// 			stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", obj.Tab(), where.String())
// 			db := dbMap[obj.DB()]
// 			row := db.QueryRow(stmtStr)
// 			err := scanRow(obj, row)
// 			if err != nil {
// 				if err == sql.ErrNoRows {
// 					return 0, nil
// 				}
// 				return 0, err
// 			}
// 			return 1, nil
// 		case ModelList:
// 			wg := sync.WaitGroup{}
// 			doneChan := make(chan interface{})
// 			errChan := make(chan error)
// 			for i := 0; i < obj.Len(); i++ {
// 				wg.Add(1)
// 				go func(index int) {
// 					defer func() {
// 						recover()
// 						wg.Done()
// 					}()
// 					w := genWhere(obj.Index(index))
// 					stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", obj.Model().Tab(), w.String())
// 					db := dbMap[obj.Model().DB()]
// 					row := db.QueryRow(stmtStr)
// 					if err := scanRow(obj.Index(index), row); err != nil {
// 						errChan <- err
// 					}
// 				}(i)
// 			}
// 			go func() {
// 				wg.Wait()
// 				close(doneChan)
// 			}()
// 			select {
// 			case err := <-errChan:
// 				close(errChan)
// 				return 0, err
// 			case <-doneChan:
// 				return obj.Len(), nil
// 			}
// 		}
// 	}
// 	switch obj := tab.(type) {
// 	case Model:
// 		stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s %s %s", obj.Tab(), where.String(), sorter.toSQL(), pager.toSQL())
// 		db := dbMap[obj.DB()]
// 		row := db.QueryRow(stmtStr)
// 		err := scanRow(obj, row)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				return 0, nil
// 			}
// 			return 0, err
// 		}
// 		return 1, nil
// 	case ModelList:
// 		stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s WHERE %s %s %s", obj.Model().Tab(), where.String(), sorter.toSQL(), pager.toSQL())
// 		db := dbMap[obj.Model().DB()]
// 		rows, err := db.Query(stmtStr)
// 		if err != nil {
// 			return 0, err
// 		}
// 		err = scanRows(obj, rows)
// 		if err != nil {
// 			return 0, err
// 		}
// 		var num int
// 		numRow := db.QueryRow("SELECT FOUND_ROWS()")
// 		err = numRow.Scan(&num)
// 		if err != nil {
// 			return 0, err
// 		}
// 		return num, nil
// 	default:
// 		panic("nborm.Query() error: unsupported type")
// 	}
// }

// func JoinQuery(tab interface{}, where *Where, relations ...relation) error {
// 	joinList := make([]string, len(relations))
// 	for i, rel := range relations {
// 		joinList[i] = rel.joinClause()
// 	}
// 	if where == nil {
// 		switch obj := tab.(type) {
// 		case Model:
// 			where = genWhere(obj)
// 			stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", obj.DB(), obj.Tab(), obj.DB(), obj.Tab(), strings.Join(joinList, " "), where.String())
// 			db := dbMap[obj.DB()]
// 			row := db.QueryRow(stmtStr)
// 			return scanRow(obj, row)
// 		case ModelList:
// 			wg := sync.WaitGroup{}
// 			doneChan := make(chan interface{})
// 			errChan := make(chan error)
// 			for i := 0; i < obj.Len(); i++ {
// 				wg.Add(1)
// 				go func(index int) {
// 					defer func() {
// 						recover()
// 						wg.Done()
// 					}()
// 					m := obj.Index(index)
// 					w := genWhere(m)
// 					stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", obj.Model().DB(), obj.Model().Tab(), obj.Model().DB(), obj.Model().Tab(), strings.Join(joinList, " "), w.String())
// 					db := dbMap[obj.Model().DB()]
// 					row := db.QueryRow(stmtStr)
// 					err := scanRow(m, row)
// 					if err != nil {
// 						errChan <- err
// 					}
// 				}(i)
// 			}
// 			go func() {
// 				wg.Wait()
// 				close(doneChan)
// 			}()
// 			select {
// 			case err := <-errChan:
// 				close(errChan)
// 				return err
// 			}
// 			return nil
// 		default:
// 			panic("nborm.JoinQuery() error: unsupported type")
// 		}
// 	} else {
// 		switch obj := tab.(type) {
// 		case Model:
// 			stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", obj.DB(), obj.Tab(), obj.DB(), obj.Tab(), strings.Join(joinList, " "), where.String())
// 			db := dbMap[obj.DB()]
// 			row := db.QueryRow(stmtStr)
// 			return scanRow(obj, row)
// 		case ModelList:
// 			stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", obj.Model().DB(), obj.Model().Tab(), obj.Model().DB(), obj.Model().Tab(), strings.Join(joinList, " "), where.String())
// 			db := dbMap[obj.Model().DB()]
// 			rows, err := db.Query(stmtStr)
// 			if err != nil {
// 				return err
// 			}
// 			return scanRows(obj, rows)
// 		default:
// 			panic("nborm.JoinQuery() error: unsupported type")
// 		}
// 	}
// }

// func JoinQueryWithFoundRows(tab interface{}, where *Where, sorter *Sorter, pager *Pager, relations ...relation) (int, error) {
// 	joinList := make([]string, len(relations))
// 	for i, rel := range relations {
// 		joinList[i] = rel.joinClause()
// 	}
// 	if where == nil {
// 		switch obj := tab.(type) {
// 		case Model:
// 			where = genWhere(obj)
// 			stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s %s %s", obj.DB(), obj.Tab(), obj.DB(), obj.Tab(), strings.Join(joinList, " "),
// 				where.String())
// 			db := dbMap[obj.DB()]
// 			row := db.QueryRow(stmtStr)
// 			err := scanRow(obj, row)
// 			if err != nil {
// 				return 0, err
// 			}
// 			return 1, nil
// 		case ModelList:
// 			wg := sync.WaitGroup{}
// 			doneChan := make(chan interface{})
// 			errChan := make(chan error)
// 			for i := 0; i < obj.Len(); i++ {
// 				wg.Add(1)
// 				go func(index int) {
// 					defer func() {
// 						recover()
// 						wg.Done()
// 					}()
// 					m := obj.Index(index)
// 					w := genWhere(m)
// 					stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", obj.Model().DB(), obj.Model().Tab(), obj.Model().DB(),
// 						obj.Model().Tab(), strings.Join(joinList, " "), w.String())
// 					db := dbMap[obj.Model().DB()]
// 					row := db.QueryRow(stmtStr)
// 					err := scanRow(m, row)
// 					if err != nil {
// 						errChan <- err
// 					}
// 				}(i)
// 			}
// 			go func() {
// 				wg.Wait()
// 				close(doneChan)
// 			}()
// 			select {
// 			case err := <-errChan:
// 				close(errChan)
// 				return 0, err
// 			case <-doneChan:
// 				return obj.Len(), nil
// 			}
// 		default:
// 			panic("nborm.JoinQuery() error: unsupported type")
// 		}
// 	} else {
// 		switch obj := tab.(type) {
// 		case Model:
// 			stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", obj.DB(), obj.Tab(), obj.DB(), obj.Tab(), strings.Join(joinList, " "),
// 				where.String())
// 			db := dbMap[obj.DB()]
// 			row := db.QueryRow(stmtStr)
// 			if err := scanRow(obj, row); err != nil {
// 				return 0, err
// 			}
// 			return 1, nil
// 		case ModelList:
// 			stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s %s WHERE %s %s %s", obj.Model().DB(), obj.Model().Tab(),
// 				obj.Model().DB(), obj.Model().Tab(), strings.Join(joinList, " "), where.String(), sorter.toSQL(), pager.toSQL())
// 			db := dbMap[obj.Model().DB()]
// 			rows, err := db.Query(stmtStr)
// 			if err != nil {
// 				return 0, err
// 			}
// 			if err = scanRows(obj, rows); err != nil {
// 				return 0, err
// 			}
// 			var num int
// 			numRow := db.QueryRow("SELECT FOUND_ROWS()")
// 			if err := numRow.Scan(&num); err != nil {
// 				return 0, err
// 			}
// 			return num, nil
// 		default:
// 			panic("nborm.JoinQuery() error: unsupported type")
// 		}
// 	}
// }

func InsertOne(m Model) error {
	return insertAndGetInc(m, false)
}

func InsertOrUpdateOne(m Model) error {
	return insertAndGetInc(m, true)
}

func InsertMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error { return insertAndGetIncContext(ctx, m, false) })
}

func InsertOrUpdateMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error { return insertAndGetIncContext(ctx, m, true) })
}

func InsertOrGetOne(m Model) error {
	err := insertAndGetInc(m, false)
	if err != nil {
		if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
			return queryAndScan(m, genSelect(m, genWhere(m), nil, nil, false))
		}
		return err
	}
	return nil
}

func InsertOrGetMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error {
		err := insertAndGetIncContext(ctx, m, false)
		if err != nil {
			if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
				return queryAndScanContext(ctx, m, genSelect(m, genWhere(m), nil, nil, false))
			}
			return err
		}
		return nil
	})
}

// func Insert(tab interface{}, valuePairs ...[2]string) error {
// 	if len(valuePairs) != 0 {
// 		var db *sql.DB
// 		var stmtStr string
// 		colList := make([]string, len(valuePairs))
// 		valList := make([]string, len(valuePairs))
// 		for i, p := range valuePairs {
// 			colList[i] = p[0]
// 			valList[i] = p[1]
// 		}
// 		switch obj := tab.(type) {
// 		case Model:
// 			db = dbMap[obj.DB()]
// 			stmtStr = fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", obj.DB(), obj.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
// 		case ModelList:
// 			stmtStr = fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", obj.Model().DB(), obj.Model().Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
// 			db = dbMap[obj.Model().DB()]
// 		default:
// 			panic("nborm.Insert(): unsupported type")
// 		}
// 		_, err := db.Exec(stmtStr)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		switch obj := tab.(type) {
// 		case ModelList:
// 			var wg sync.WaitGroup
// 			doneChan := make(chan interface{})
// 			errChan := make(chan error)
// 			for i := 0; i < obj.Len(); i++ {
// 				wg.Add(1)
// 				go func(index int) {
// 					defer func() {
// 						recover()
// 						wg.Done()
// 					}()
// 					m := obj.Index(index)
// 					inc, others := getInc(m)
// 					others = filterValid(others)
// 					colList := make([]string, len(others))
// 					valList := make([]string, len(others))
// 					for i, f := range others {
// 						p := f.InsertValuePair()
// 						colList[i] = p[0]
// 						valList[i] = p[1]
// 					}
// 					stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", obj.Model().DB(), obj.Model().Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
// 					db := dbMap[obj.Model().DB()]
// 					res, err := db.Exec(stmtStr)
// 					if err != nil {
// 						errChan <- err
// 					}
// 					lastInsertId, err := res.LastInsertId()
// 					if err != nil {
// 						errChan <- err
// 						return
// 					}
// 					inc.(*IntField).Set(lastInsertId, false)
// 				}(i)
// 			}
// 			go func() {
// 				wg.Wait()
// 				close(doneChan)
// 			}()
// 			select {
// 			case err := <-errChan:
// 				close(errChan)
// 				return err
// 			case <-doneChan:
// 				return nil
// 			}
// 		case Model:
// 			inc, others := getInc(obj)
// 			others = filterValid(others)
// 			colList := make([]string, len(others))
// 			valList := make([]string, len(others))
// 			for i, f := range others {
// 				p := f.InsertValuePair()
// 				colList[i] = p[0]
// 				valList[i] = p[1]
// 			}
// 			stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", obj.DB(), obj.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
// 			db := dbMap[obj.DB()]
// 			res, err := db.Exec(stmtStr)
// 			if err != nil {
// 				return err
// 			}
// 			lastInsertId, err := res.LastInsertId()
// 			if err != nil {
// 				return err
// 			}
// 			inc.(*IntField).Set(lastInsertId, false)
// 		default:
// 			panic("nborm.Insert() error: unsupported type")
// 		}
// 	}
// 	return nil
// }

func UpdateOne(m Model) error {
	_, fs := getInc(m)
	fs = filterValid(fs)
	setValues := make([]string, 0, len(fs))
	for _, f := range fs {
		setValues = append(setValues, f.UpdateValue().String())
	}
	stmtStr := fmt.Sprintf("UPDATE %s.%s SET %s WHERE %s", m.DB(), m.Tab(), strings.Join(setValues, ", "), genWhere(m).String())
	db := dbMap[m.DB()]
	_, err := db.Exec(stmtStr)
	return err
}

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
			stmtStr := fmt.Sprintf("UPDATE %s.%s SET %s WHERE %s", m.DB(), m.Tab(), strings.Join(setValues, ", "), genWhere(m).String())
			db := dbMap[m.DB()]
			_, err := db.ExecContext(ctx, stmtStr)
			return err
		}
	})
}

func BulkUpdate(m Model, where *Where, values ...*UpdateValue) error {
	db := dbMap[m.DB()]
	setList := make([]string, len(values))
	for i, val := range values {
		setList[i] = val.String()
	}
	stmtStr := fmt.Sprintf("UPDATE %s.%s SET %s WHERE %s", m.DB(), m.Tab(), strings.Join(setList, ", "), where.String())
	_, err := db.Exec(stmtStr)
	return err
}

// func Update(tab interface{}, where *Where, values ...*UpdateValue) error {
// 	if where != nil {
// 		var db *sql.DB
// 		var stmtStr string
// 		valStrs := make([]string, len(values))
// 		for i, val := range values {
// 			valStrs[i] = val.String()
// 		}
// 		switch obj := tab.(type) {
// 		case Model:
// 			db = dbMap[obj.DB()]
// 			stmtStr = fmt.Sprintf("UPDATE %s.%s SET %s WHERE %s", obj.DB(), obj.Tab(), strings.Join(valStrs, ", "), where.String())
// 		case ModelList:
// 			db = dbMap[obj.Model().DB()]
// 			stmtStr = fmt.Sprintf("UPDATE %s.%s SET %s WHERE %s", obj.Model().DB(), obj.Model().Tab(), strings.Join(valStrs, ", "), where.String())
// 		default:
// 			panic("nborm.Update() error: unsupported type")
// 		}
// 		_, err := db.Exec(stmtStr)
// 		return err
// 	} else {
// 		switch obj := tab.(type) {
// 		case Model:
// 			_, fs := getInc(obj)
// 			fs = filterValid(fs)
// 			setValues := make([]string, 0, len(fs))
// 			for _, f := range fs {
// 				setValues = append(setValues, f.UpdateValue().String())
// 			}
// 			stmtStr := fmt.Sprintf("UPDATE %s.%s SET %s WHERE %s", obj.DB(), obj.Tab(), strings.Join(setValues, ", "), genWhere(obj).String())
// 			db := dbMap[obj.DB()]
// 			_, err := db.Exec(stmtStr)
// 			return err
// 		case ModelList:
// 			var wg sync.WaitGroup
// 			doneChan := make(chan interface{})
// 			errChan := make(chan error)
// 			if len(values) > 0 {
// 				valStrs := make([]string, len(values))
// 				for i, val := range values {
// 					valStrs[i] = val.String()
// 				}
// 				for i := 0; i < obj.Len(); i++ {
// 					wg.Add(1)
// 					go func(index int) {
// 						defer func() {
// 							recover()
// 							wg.Done()
// 						}()
// 						stmtStr := fmt.Sprintf("UPDATE %s.%s SET %s WHERE %s", obj.Model().DB(), obj.Model().Tab(), strings.Join(valStrs, ", "), genWhere(obj.Index(index)))
// 						db := dbMap[obj.Model().DB()]
// 						_, err := db.Exec(stmtStr)
// 						if err != nil {
// 							errChan <- err
// 							return
// 						}
// 						updateModel(obj.Index(index), values...)
// 					}(i)
// 				}
// 			} else {
// 				for i := 0; i < obj.Len(); i++ {
// 					wg.Add(1)
// 					go func(index int) {
// 						defer func() {
// 							recover()
// 							wg.Done()
// 						}()
// 						_, others := getInc(obj.Index(index))
// 						others = filterValid(others)
// 						setValues := make([]string, len(others))
// 						for j, f := range others {
// 							setValues[j] = f.UpdateValue().String()
// 						}
// 						stmtStr := fmt.Sprintf("UPDATE %s.%s SET %s WHERE %s", obj.Model().DB(), obj.Model().Tab(), strings.Join(setValues, ", "), genWhere(obj.Index(index)))
// 						db := dbMap[obj.Model().DB()]
// 						_, err := db.Exec(stmtStr)
// 						if err != nil {
// 							errChan <- err
// 						}
// 					}(i)
// 				}
// 			}
// 			go func() {
// 				wg.Wait()
// 				close(doneChan)
// 			}()
// 			select {
// 			case err := <-errChan:
// 				close(errChan)
// 				return err
// 			case <-doneChan:
// 				return nil
// 			}
// 		default:
// 			panic("nborm.Update() error: unsupported type")
// 		}
// 	}
// }

// func InsertOrUpdate(tab interface{}, valuePairs ...[2]string) error {
// 	if len(valuePairs) > 0 {
// 		colList := make([]string, len(valuePairs))
// 		valList := make([]string, len(valuePairs))
// 		updateList := make([]string, len(valuePairs))
// 		for i, p := range valuePairs {
// 			colList[i] = p[0]
// 			valList[i] = p[1]
// 			updateList[i] = p[0] + "=" + p[1]
// 		}
// 		var db *sql.DB
// 		var stmtStr string
// 		switch obj := tab.(type) {
// 		case Model:
// 			db = dbMap[obj.DB()]
// 			stmtStr = fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s", obj.DB(), obj.Tab(), strings.Join(colList, ", "),
// 				strings.Join(valList, ", "), strings.Join(updateList, ", "))
// 		case ModelList:
// 			db = dbMap[obj.Model().DB()]
// 			stmtStr = fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s", obj.Model().DB(), obj.Model().Tab(), strings.Join(colList, ", "),
// 				strings.Join(valList, ", "), strings.Join(updateList, ", "))
// 		}
// 		_, err := db.Exec(stmtStr)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		switch obj := tab.(type) {
// 		case Model:
// 			inc, fs := getInc(obj)
// 			validFields := filterValid(fs)
// 			colList := make([]string, len(validFields))
// 			valList := make([]string, len(validFields))
// 			updateList := make([]string, len(validFields))
// 			for i, f := range validFields {
// 				valuePair := f.InsertValuePair()
// 				colList[i], valList[i], updateList[i] = valuePair[0], valuePair[1], valuePair[0]+"="+valuePair[1]
// 			}
// 			stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s",
// 				obj.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "), inc.Column(), inc.Column(), strings.Join(updateList, ", "))
// 			db := dbMap[obj.DB()]
// 			res, err := db.Exec(stmtStr)
// 			if err != nil {
// 				return err
// 			}
// 			lastInsertId, err := res.LastInsertId()
// 			if err != nil {
// 				return err
// 			}
// 			inc.(*IntField).Set(lastInsertId, false)
// 		case ModelList:
// 			var wg sync.WaitGroup
// 			doneChan := make(chan interface{})
// 			errChan := make(chan error)
// 			for i := 0; i < obj.Len(); i++ {
// 				wg.Add(1)
// 				go func(index int) {
// 					defer func() {
// 						recover()
// 						wg.Done()
// 					}()
// 					m := obj.Index(index)
// 					inc, others := getInc(m)
// 					others = filterValid(others)
// 					colList := make([]string, len(others))
// 					valList := make([]string, len(others))
// 					updateList := make([]string, len(others))
// 					for j, f := range others {
// 						valuePair := f.InsertValuePair()
// 						colList[j], valList[j], updateList[j] = valuePair[0], valuePair[1], valuePair[0]+"="+valuePair[1]
// 					}
// 					stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s",
// 						obj.Model().Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "), inc.Column(), inc.Column(), strings.Join(updateList, ", "))
// 					db := dbMap[obj.Model().DB()]
// 					res, err := db.Exec(stmtStr)
// 					if err != nil {
// 						errChan <- err
// 						return
// 					}
// 					lastInsertId, err := res.LastInsertId()
// 					if err != nil {
// 						errChan <- err
// 						return
// 					}
// 					inc.(*IntField).Set(lastInsertId, false)
// 				}(i)
// 			}
// 			go func() {
// 				wg.Wait()
// 				close(doneChan)
// 			}()
// 			select {
// 			case err := <-errChan:
// 				close(errChan)
// 				return err
// 			case <-doneChan:
// 				return nil
// 			}
// 		default:
// 			panic("nborm.InsertOrUpdate() error: unsupported type")
// 		}
// 	}
// 	return nil
// }

// func InsertOrGet(tab interface{}) error {
// 	switch obj := tab.(type) {
// 	case Model:
// 		inc, fs := getInc(obj)
// 		validFields := filterValid(fs)
// 		colList := make([]string, len(validFields))
// 		valList := make([]string, len(validFields))
// 		for i, f := range validFields {
// 			valuePair := f.InsertValuePair()
// 			colList[i], valList[i] = valuePair[0], valuePair[1]
// 		}
// 		stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", obj.DB(), obj.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
// 		db := dbMap[obj.DB()]
// 		res, err := db.Exec(stmtStr)
// 		if err != nil {
// 			if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
// 				return Query(obj, nil, nil, nil)
// 			} else {
// 				return err
// 			}
// 		}
// 		lastInsertId, err := res.LastInsertId()
// 		if err != nil {
// 			return err
// 		}
// 		inc.(*IntField).Set(lastInsertId, false)
// 	case ModelList:
// 		var wg sync.WaitGroup
// 		doneChan := make(chan interface{})
// 		errChan := make(chan error)
// 		for i := 0; i < obj.Len(); i++ {
// 			wg.Add(1)
// 			go func(index int) {
// 				defer func() {
// 					recover()
// 					wg.Done()
// 				}()
// 				m := obj.Index(index)
// 				inc, others := getInc(m)
// 				others = filterValid(others)
// 				colList := make([]string, len(others))
// 				valList := make([]string, len(others))
// 				for j, f := range others {
// 					valuePair := f.InsertValuePair()
// 					colList[j], valList[j] = valuePair[0], valuePair[1]
// 				}
// 				stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", obj.Model().DB(), obj.Model().Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
// 				db := dbMap[obj.Model().DB()]
// 				res, err := db.Exec(stmtStr)
// 				if err != nil {
// 					if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
// 						if err := Query(m, nil, nil, nil); err != nil {
// 							errChan <- err
// 						}
// 						return
// 					} else {
// 						errChan <- err
// 						return
// 					}
// 				}
// 				lastInsertId, err := res.LastInsertId()
// 				if err != nil {
// 					errChan <- err
// 					return
// 				}
// 				inc.(*IntField).Set(lastInsertId, false)
// 			}(i)
// 			func() {
// 				wg.Wait()
// 				close(doneChan)
// 			}()
// 			select {
// 			case err := <-errChan:
// 				close(errChan)
// 				return err
// 			case <-doneChan:
// 				return nil
// 			}
// 		}
// 	default:
// 		panic("nborm.InsertOrUpdate() error: unsupported type")
// 	}
// 	return nil
// }

func DeleteOne(m Model) error {
	db := dbMap[m.DB()]
	stmtStr := fmt.Sprintf("DELETE FROM %s.%s WHERE %s", m.DB(), m.Tab(), genWhere(m).String())
	_, err := db.Exec(stmtStr)
	return err
}

func DeleteMul(l ModelList) error {
	return iterList(l, func(ctx context.Context, m Model) error {
		db := dbMap[m.DB()]
		stmtStr := fmt.Sprintf("DELETE FROM %s.%s WHERE %s", m.DB(), m.Tab(), genWhere(m).String())
		_, err := db.ExecContext(ctx, stmtStr)
		return err
	})
}

func BulkDelete(m Model, where *Where) error {
	db := dbMap[m.DB()]
	stmtStr := fmt.Sprintf("DELETE FROM %s.%s WHERE %s", m.DB(), m.Tab(), where.String())
	_, err := db.Exec(stmtStr)
	return err
}

func DeleteAll(m Model) error {
	db := dbMap[m.DB()]
	stmtStr := fmt.Sprintf("DELETE FROM %s.%s")
	_, err := db.Exec(stmtStr)
	return err
}

func Count(m Model, where *Where) (int, error) {
	db := dbMap[m.DB()]
	var stmtStr string
	if where == nil {
		stmtStr = fmt.Sprintf("SELECT COUNT(*) FROM %s.%s")
	} else {
		stmtStr = fmt.Sprintf("SELECT COUNT(*) FROM %s.%s WHERE %s", where.String())
	}
	var num int
	row := db.QueryRow(stmtStr)
	err := row.Scan(&num)
	if err != nil {
		return -1, err
	}
	return num, nil

}

// func Delete(tab interface{}, where *Where) error {
// 	if where != nil {
// 		var db *sql.DB
// 		var stmtStr string
// 		switch obj := tab.(type) {
// 		case Model:
// 			db = dbMap[obj.DB()]
// 			stmtStr = fmt.Sprintf("DELETE FROM %s.%s WHERE %s", obj.DB(), obj.Tab(), where.String())
// 		case ModelList:
// 			db = dbMap[obj.Model().DB()]
// 			stmtStr = fmt.Sprintf("DELETE FROM %s.%s WHERE %s", obj.Model().DB(), obj.Model().Tab(), where.String())
// 		}
// 		_, err := db.Exec(stmtStr)
// 		if err != nil {
// 			return err
// 		}
// 	} else {
// 		switch obj := tab.(type) {
// 		case Model:
// 			pk, _ := getPk(obj)
// 			stmtStr := fmt.Sprintf("DELETE FROM %s WHERE %s", obj.Tab(), pk.Where().String())
// 			db := dbMap[obj.DB()]
// 			_, err := db.Exec(stmtStr)
// 			if err != nil {
// 				return err
// 			}
// 			invalidateModel(obj)
// 		case ModelList:
// 			for i := 0; i < obj.Len(); i++ {
// 				pk, _ := getPk(obj.Index(i))
// 				stmtStr := fmt.Sprintf("DELETE FROM %s WHERE %s", obj.Model().Tab(), pk.Where().String())
// 				db := dbMap[obj.Model().DB()]
// 				_, err := db.Exec(stmtStr)
// 				if err != nil {
// 					return err
// 				}
// 				invalidateModel(obj.Index(i))
// 			}
// 		default:
// 			panic("nborm.Delete() error: unsupported type")
// 		}
// 	}
// 	return nil
// }

// func DeleteAll(tab interface{}) error {
// 	var db *sql.DB
// 	var err error
// 	switch obj := tab.(type) {
// 	case Model:
// 		db = dbMap[obj.DB()]
// 		_, err = db.Exec(fmt.Sprintf("DELETE FROM %s.%s", obj.DB(), obj.Tab()))
// 	case ModelList:
// 		db = dbMap[obj.Model().DB()]
// 		_, err = db.Exec(fmt.Sprintf("DELETE FROM %s.%s", obj.Model().DB(), obj.Model().Tab()))
// 	}
// 	return err
// }

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

// func Count(tab interface{}, where *Where) (int, error) {
// 	var m Model
// 	switch obj := tab.(type) {
// 	case Model:
// 		m = obj
// 	case ModelList:
// 		m = obj.Model()
// 	}
// 	var stmtStr string
// 	if where != nil {
// 		stmtStr = fmt.Sprintf("SELECT COUNT(*) FROM %s.%s WHERE %s", m.DB(), m.Tab(), where.String())
// 	} else {
// 		stmtStr = fmt.Sprintf("SELECT COUNT(*) FROM %s.%s", m.DB(), m.Tab())
// 	}
// 	db := dbMap[m.DB()]
// 	row := db.QueryRow(stmtStr)
// 	var num int
// 	err := row.Scan(&num)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return num, nil
// }
