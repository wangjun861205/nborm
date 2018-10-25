package nborm

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/go-sql-driver/mysql"
)

func All(l ModelList, sorter *Sorter, pager *Pager) error {
	stmtStr := fmt.Sprintf("SELECT * FROM %s %s %s", l.Tab(), sorter.toSQL(), pager.toSQL())
	rows, err := dbMap[l.DB()].Query(stmtStr)
	if err != nil {
		return err
	}
	return scanRows(l, rows)
}

func Query(tab table, where *Where, sorter *Sorter, pager *Pager) error {
	db := dbMap[tab.DB()]
	if where == nil {
		switch obj := tab.(type) {
		case Model:
			where = genWhere(obj)
			stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", tab.Tab(), where.String())
			row := db.QueryRow(stmtStr)
			return scanRow(obj, row)
		case ModelList:
			wg := sync.WaitGroup{}
			doneChan := make(chan interface{})
			errChan := make(chan error)
			for i := 0; i < obj.Len(); i++ {
				wg.Add(1)
				go func(index int) {
					defer func() {
						recover()
						wg.Done()
					}()
					w := genWhere(obj.Index(index))
					stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", tab.Tab(), w.String())
					row := db.QueryRow(stmtStr)
					if err := scanRow(obj.Index(index), row); err != nil {
						errChan <- err
					}
				}(i)
			}
			go func() {
				wg.Wait()
				close(doneChan)
			}()
			select {
			case err := <-errChan:
				close(errChan)
				return err
			case <-doneChan:
				return nil
			}
		}
	}
	stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s %s %s", tab.Tab(), where.String(), sorter.toSQL(), pager.toSQL())
	switch obj := tab.(type) {
	case Model:
		row := db.QueryRow(stmtStr)
		return scanRow(obj, row)
	case ModelList:
		rows, err := db.Query(stmtStr)
		if err != nil {
			return err
		}
		return scanRows(obj, rows)
	default:
		panic("nborm.Query() error: unsupported type")
	}
}

func QueryWithFoundRows(tab table, where *Where, sorter *Sorter, pager *Pager) (int, error) {
	db := dbMap[tab.DB()]
	if where == nil {
		switch obj := tab.(type) {
		case Model:
			where = genWhere(obj)
			stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", tab.Tab(), where.String())
			row := db.QueryRow(stmtStr)
			err := scanRow(obj, row)
			if err != nil {
				if err == sql.ErrNoRows {
					return 0, nil
				}
				return 0, err
			}
			return 1, nil
		case ModelList:
			wg := sync.WaitGroup{}
			doneChan := make(chan interface{})
			errChan := make(chan error)
			for i := 0; i < obj.Len(); i++ {
				wg.Add(1)
				go func(index int) {
					defer func() {
						recover()
						wg.Done()
					}()
					w := genWhere(obj.Index(index))
					stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", tab.Tab(), w.String())
					row := db.QueryRow(stmtStr)
					if err := scanRow(obj.Index(index), row); err != nil {
						errChan <- err
					}
				}(i)
			}
			go func() {
				wg.Wait()
				close(doneChan)
			}()
			select {
			case err := <-errChan:
				close(errChan)
				return 0, err
			case <-doneChan:
				return obj.Len(), nil
			}
		}
	}
	switch obj := tab.(type) {
	case Model:
		stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s %s %s", tab.Tab(), where.String(), sorter.toSQL(), pager.toSQL())
		row := db.QueryRow(stmtStr)
		err := scanRow(obj, row)
		if err != nil {
			if err == sql.ErrNoRows {
				return 0, nil
			}
			return 0, err
		}
		return 1, nil
	case ModelList:
		stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS * FROM %s WHERE %s %s %s", tab.Tab(), where.String(), sorter.toSQL(), pager.toSQL())
		rows, err := db.Query(stmtStr)
		if err != nil {
			return 0, err
		}
		err = scanRows(obj, rows)
		if err != nil {
			return 0, err
		}
		var num int
		numRow := db.QueryRow("SELECT FOUND_ROWS()")
		err = numRow.Scan(&num)
		if err != nil {
			return 0, err
		}
		return num, nil
	default:
		panic("nborm.Query() error: unsupported type")
	}
}

func JoinQuery(tab table, where *Where, relations ...relation) error {
	db := dbMap[tab.DB()]
	joinList := make([]string, len(relations))
	for i, rel := range relations {
		switch r := rel.(type) {
		case *OneToOne, *ForeignKey, *ReverseForeignKey:
			joinList[i] = fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", r.dstDB(), r.dstTab(), tab.DB(), tab.Tab(), r.srcCol(),
				r.dstDB(), r.dstTab(), r.dstCol())
		case *ManyToMany:
			joinList[i] = fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s", r.midDB(), r.midTab(), tab.DB(), tab.Tab(),
				r.srcCol(), r.midDB(), r.midTab(), r.midLeftCol(), r.dstDB(), r.dstTab(), r.midDB(), r.midTab(), r.midRightCol(), r.dstDB(), r.dstTab(),
				r.dstCol())
		default:
			panic("nborm.JoinQuery() error: unsupported relation")
		}
	}
	if where == nil {
		switch obj := tab.(type) {
		case Model:
			where = genWhere(obj)
			stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", tab.DB(), tab.Tab(), tab.DB(), tab.Tab(), strings.Join(joinList, " "), where.String())
			row := db.QueryRow(stmtStr)
			return scanRow(obj, row)
		case ModelList:
			wg := sync.WaitGroup{}
			doneChan := make(chan interface{})
			errChan := make(chan error)
			for i := 0; i < obj.Len(); i++ {
				wg.Add(1)
				go func(index int) {
					defer func() {
						recover()
						wg.Done()
					}()
					m := obj.Index(index)
					w := genWhere(m)
					stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", tab.DB(), tab.Tab(), tab.DB(), tab.Tab(), strings.Join(joinList, " "), w.String())
					row := db.QueryRow(stmtStr)
					err := scanRow(m, row)
					if err != nil {
						errChan <- err
					}
				}(i)
			}
			go func() {
				wg.Wait()
				close(doneChan)
			}()
			select {
			case err := <-errChan:
				close(errChan)
				return err
			}
			return nil
		default:
			panic("nborm.JoinQuery() error: unsupported type")
		}
	} else {
		stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", tab.DB(), tab.Tab(), tab.DB(), tab.Tab(), strings.Join(joinList, " "), where.String())
		switch obj := tab.(type) {
		case Model:
			row := db.QueryRow(stmtStr)
			return scanRow(obj, row)
		case ModelList:
			rows, err := db.Query(stmtStr)
			if err != nil {
				return err
			}
			return scanRows(obj, rows)
		default:
			panic("nborm.JoinQuery() error: unsupported type")
		}
	}
}

func JoinQueryWithFoundRows(tab table, where *Where, sorter *Sorter, pager *Pager, relations ...relation) (int, error) {
	db := dbMap[tab.DB()]
	joinList := make([]string, len(relations))
	for i, rel := range relations {
		switch r := rel.(type) {
		case *OneToOne, *ForeignKey, *ReverseForeignKey:
			joinList[i] = fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", r.dstDB(), r.dstTab(), tab.DB(), tab.Tab(), r.srcCol(),
				r.dstDB(), r.dstTab(), r.dstCol())
		case *ManyToMany:
			joinList[i] = fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s", r.midDB(), r.midTab(), tab.DB(), tab.Tab(),
				r.srcCol(), r.midDB(), r.midTab(), r.midLeftCol(), r.dstDB(), r.dstTab(), r.midDB(), r.midTab(), r.midRightCol(), r.dstDB(), r.dstTab(),
				r.dstCol())
		default:
			panic("nborm.JoinQuery() error: unsupported relation")
		}
	}
	if where == nil {
		switch obj := tab.(type) {
		case Model:
			where = genWhere(obj)
			stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s %s %s", tab.DB(), tab.Tab(), tab.DB(), tab.Tab(), strings.Join(joinList, " "),
				where.String())
			row := db.QueryRow(stmtStr)
			err := scanRow(obj, row)
			if err != nil {
				return 0, err
			}
			return 1, nil
		case ModelList:
			wg := sync.WaitGroup{}
			doneChan := make(chan interface{})
			errChan := make(chan error)
			for i := 0; i < obj.Len(); i++ {
				wg.Add(1)
				go func(index int) {
					defer func() {
						recover()
						wg.Done()
					}()
					m := obj.Index(index)
					w := genWhere(m)
					stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", tab.DB(), tab.Tab(), tab.DB(), tab.Tab(), strings.Join(joinList, " "), w.String())
					row := db.QueryRow(stmtStr)
					err := scanRow(m, row)
					if err != nil {
						errChan <- err
					}
				}(i)
			}
			go func() {
				wg.Wait()
				close(doneChan)
			}()
			select {
			case err := <-errChan:
				close(errChan)
				return 0, err
			case <-doneChan:
				return obj.Len(), nil
			}
		default:
			panic("nborm.JoinQuery() error: unsupported type")
		}
	} else {
		switch obj := tab.(type) {
		case Model:
			stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", tab.DB(), tab.Tab(), tab.DB(), tab.Tab(), strings.Join(joinList, " "),
				where.String())
			row := db.QueryRow(stmtStr)
			if err := scanRow(obj, row); err != nil {
				return 0, err
			}
			return 1, nil
		case ModelList:
			stmtStr := fmt.Sprintf("SELECT SQL_CALC_FOUND_ROWS %s.%s.* FROM %s.%s %s WHERE %s %s %s", tab.DB(), tab.Tab(), tab.DB(), tab.Tab(),
				strings.Join(joinList, " "), where.String(), sorter.toSQL(), pager.toSQL())
			rows, err := db.Query(stmtStr)
			if err != nil {
				return 0, err
			}
			if err = scanRows(obj, rows); err != nil {
				return 0, err
			}
			var num int
			numRow := db.QueryRow("SELECT FOUND_ROWS()")
			if err := numRow.Scan(&num); err != nil {
				return 0, err
			}
			return num, nil
		default:
			panic("nborm.JoinQuery() error: unsupported type")
		}
	}
}

func Insert(tab table, valuePairs ...[2]string) error {
	db := dbMap[tab.DB()]
	if len(valuePairs) != 0 {
		colList := make([]string, len(valuePairs))
		valList := make([]string, len(valuePairs))
		for i, p := range valuePairs {
			colList[i] = p[0]
			valList[i] = p[1]
		}
		stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", tab.DB(), tab.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
		_, err := db.Exec(stmtStr)
		if err != nil {
			return err
		}
	} else {
		switch obj := tab.(type) {
		case ModelList:
			var wg sync.WaitGroup
			doneChan := make(chan interface{})
			errChan := make(chan error)
			for i := 0; i < obj.Len(); i++ {
				wg.Add(1)
				go func(index int) {
					defer func() {
						recover()
						wg.Done()
					}()
					m := obj.Index(index)
					inc, others := getInc(m)
					others = filterValid(others)
					colList := make([]string, len(others))
					valList := make([]string, len(others))
					for i, f := range others {
						p := f.InsertValuePair()
						colList[i] = p[0]
						valList[i] = p[1]
					}
					stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", tab.DB(), tab.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
					res, err := db.Exec(stmtStr)
					if err != nil {
						errChan <- err
					}
					lastInsertId, err := res.LastInsertId()
					if err != nil {
						errChan <- err
						return
					}
					inc.(*IntField).Set(lastInsertId, false)
				}(i)
			}
			go func() {
				wg.Wait()
				close(doneChan)
			}()
			select {
			case err := <-errChan:
				close(errChan)
				return err
			case <-doneChan:
				return nil
			}
		case Model:
			inc, others := getInc(obj)
			others = filterValid(others)
			colList := make([]string, len(others))
			valList := make([]string, len(others))
			for i, f := range others {
				p := f.InsertValuePair()
				colList[i] = p[0]
				valList[i] = p[1]
			}
			stmtStr := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES (%s)", tab.DB(), tab.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
			res, err := db.Exec(stmtStr)
			if err != nil {
				return err
			}
			lastInsertId, err := res.LastInsertId()
			if err != nil {
				return err
			}
			inc.(*IntField).Set(lastInsertId, false)
		default:
			panic("nborm.Insert() error: unsupported type")
		}
	}
	return nil
}

func Update(tab table, where *Where, values ...*UpdateValue) error {
	db := dbMap[tab.DB()]
	if where != nil {
		valStrs := make([]string, len(values))
		for i, val := range values {
			valStrs[i] = val.String()
		}
		stmtStr := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tab.Tab(), strings.Join(valStrs, ", "), where.String())
		_, err := db.Exec(stmtStr)
		return err
	} else {
		switch obj := tab.(type) {
		case Model:
			_, fs := getInc(obj)
			fs = filterValid(fs)
			setValues := make([]string, 0, len(fs))
			for _, f := range fs {
				setValues = append(setValues, f.UpdateValue().String())
			}
			stmtStr := fmt.Sprintf("UPDATE %s SET %s WHERE %s", obj.Tab(), strings.Join(setValues, ", "), genWhere(obj).String())
			_, err := db.Exec(stmtStr)
			return err
		case ModelList:
			var wg sync.WaitGroup
			doneChan := make(chan interface{})
			errChan := make(chan error)
			if len(values) > 0 {
				valStrs := make([]string, len(values))
				for i, val := range values {
					valStrs[i] = val.String()
				}
				for i := 0; i < obj.Len(); i++ {
					wg.Add(1)
					go func(index int) {
						defer func() {
							recover()
							wg.Done()
						}()
						stmtStr := fmt.Sprintf("UPDATE %s SET %s WHERE %s", obj.Tab(), strings.Join(valStrs, ", "), genWhere(obj.Index(index)))
						_, err := db.Exec(stmtStr)
						if err != nil {
							errChan <- err
							return
						}
						updateModel(obj.Index(index), values...)
					}(i)
				}
			} else {
				for i := 0; i < obj.Len(); i++ {
					wg.Add(1)
					go func(index int) {
						defer func() {
							recover()
							wg.Done()
						}()
						_, others := getInc(obj.Index(index))
						others = filterValid(others)
						setValues := make([]string, len(others))
						for j, f := range others {
							setValues[j] = f.UpdateValue().String()
						}
						stmtStr := fmt.Sprintf("UPDATE %s SET %s WHERE %s", obj.Tab(), strings.Join(setValues, ", "), genWhere(obj.Index(index)))
						_, err := db.Exec(stmtStr)
						if err != nil {
							errChan <- err
						}
					}(i)
				}
			}
			go func() {
				wg.Wait()
				close(doneChan)
			}()
			select {
			case err := <-errChan:
				close(errChan)
				return err
			case <-doneChan:
				return nil
			}
		default:
			panic("nborm.Update() error: unsupported type")
		}
	}
}

func InsertOrUpdate(tab table, valuePairs ...[2]string) error {
	db := dbMap[tab.DB()]
	if len(valuePairs) > 0 {
		colList := make([]string, len(valuePairs))
		valList := make([]string, len(valuePairs))
		updateList := make([]string, len(valuePairs))
		for i, p := range valuePairs {
			colList[i] = p[0]
			valList[i] = p[1]
			updateList[i] = p[0] + "=" + p[1]
		}
		stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s", tab.Tab(), strings.Join(colList, ", "),
			strings.Join(valList, ", "), strings.Join(updateList, ", "))
		_, err := db.Exec(stmtStr)
		if err != nil {
			return err
		}
	} else {
		switch obj := tab.(type) {
		case Model:
			inc, fs := getInc(obj)
			validFields := filterValid(fs)
			colList := make([]string, len(validFields))
			valList := make([]string, len(validFields))
			updateList := make([]string, len(validFields))
			for i, f := range validFields {
				valuePair := f.InsertValuePair()
				colList[i], valList[i], updateList[i] = valuePair[0], valuePair[1], valuePair[0]+"="+valuePair[1]
			}
			stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s",
				obj.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "), inc.Column(), inc.Column(), strings.Join(updateList, ", "))
			res, err := db.Exec(stmtStr)
			if err != nil {
				return err
			}
			lastInsertId, err := res.LastInsertId()
			if err != nil {
				return err
			}
			inc.(*IntField).Set(lastInsertId, false)
		case ModelList:
			var wg sync.WaitGroup
			doneChan := make(chan interface{})
			errChan := make(chan error)
			for i := 0; i < obj.Len(); i++ {
				wg.Add(1)
				go func(index int) {
					defer func() {
						recover()
						wg.Done()
					}()
					m := obj.Index(index)
					inc, others := getInc(m)
					others = filterValid(others)
					colList := make([]string, len(others))
					valList := make([]string, len(others))
					updateList := make([]string, len(others))
					for j, f := range others {
						valuePair := f.InsertValuePair()
						colList[j], valList[j], updateList[j] = valuePair[0], valuePair[1], valuePair[0]+"="+valuePair[1]
					}
					stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) ON DUPLICATE KEY UPDATE %s = LAST_INSERT_ID(%s), %s",
						obj.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "), inc.Column(), inc.Column(), strings.Join(updateList, ", "))
					res, err := db.Exec(stmtStr)
					if err != nil {
						errChan <- err
						return
					}
					lastInsertId, err := res.LastInsertId()
					if err != nil {
						errChan <- err
						return
					}
					inc.(*IntField).Set(lastInsertId, false)
				}(i)
			}
			go func() {
				wg.Wait()
				close(doneChan)
			}()
			select {
			case err := <-errChan:
				close(errChan)
				return err
			case <-doneChan:
				return nil
			}
		default:
			panic("nborm.InsertOrUpdate() error: unsupported type")
		}
	}
	return nil
}

func InsertOrGet(tab table) error {
	db := dbMap[tab.DB()]
	switch obj := tab.(type) {
	case Model:
		inc, fs := getInc(obj)
		validFields := filterValid(fs)
		colList := make([]string, len(validFields))
		valList := make([]string, len(validFields))
		for i, f := range validFields {
			valuePair := f.InsertValuePair()
			colList[i], valList[i] = valuePair[0], valuePair[1]
		}
		stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tab.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
		res, err := db.Exec(stmtStr)
		if err != nil {
			if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
				return Query(obj, nil, nil, nil)
			} else {
				return err
			}
		}
		lastInsertId, err := res.LastInsertId()
		if err != nil {
			return err
		}
		inc.(*IntField).Set(lastInsertId, false)
	case ModelList:
		var wg sync.WaitGroup
		doneChan := make(chan interface{})
		errChan := make(chan error)
		for i := 0; i < obj.Len(); i++ {
			wg.Add(1)
			go func(index int) {
				defer func() {
					recover()
					wg.Done()
				}()
				m := obj.Index(index)
				inc, others := getInc(m)
				others = filterValid(others)
				colList := make([]string, len(others))
				valList := make([]string, len(others))
				for j, f := range others {
					valuePair := f.InsertValuePair()
					colList[j], valList[j] = valuePair[0], valuePair[1]
				}
				stmtStr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tab.Tab(), strings.Join(colList, ", "), strings.Join(valList, ", "))
				res, err := db.Exec(stmtStr)
				if err != nil {
					if e, ok := err.(*mysql.MySQLError); ok && e.Number == 1062 {
						if err := Query(m, nil, nil, nil); err != nil {
							errChan <- err
						}
						return
					} else {
						errChan <- err
						return
					}
				}
				lastInsertId, err := res.LastInsertId()
				if err != nil {
					errChan <- err
					return
				}
				inc.(*IntField).Set(lastInsertId, false)
			}(i)
			func() {
				wg.Wait()
				close(doneChan)
			}()
			select {
			case err := <-errChan:
				close(errChan)
				return err
			case <-doneChan:
				return nil
			}
		}
	default:
		panic("nborm.InsertOrUpdate() error: unsupported type")
	}
	return nil
}

func Delete(tab table, where *Where) error {
	db := dbMap[tab.DB()]
	if where != nil {
		stmtStr := fmt.Sprintf("DELETE FROM %s WHERE %s", tab.Tab(), where.String())
		_, err := db.Exec(stmtStr)
		if err != nil {
			return err
		}
	} else {
		switch obj := tab.(type) {
		case Model:
			pk, _ := getPk(obj)
			stmtStr := fmt.Sprintf("DELETE FROM %s WHERE %s", obj.Tab(), pk.Where().String())
			_, err := db.Exec(stmtStr)
			if err != nil {
				return err
			}
			invalidateModel(obj)
		case ModelList:
			for i := 0; i < obj.Len(); i++ {
				pk, _ := getPk(obj.Index(i))
				stmtStr := fmt.Sprintf("DELETE FROM %s WHERE %s", obj.Tab(), pk.Where().String())
				_, err := db.Exec(stmtStr)
				if err != nil {
					return err
				}
				invalidateModel(obj.Index(i))
			}
		default:
			panic("nborm.Delete() error: unsupported type")
		}
	}
	return nil
}

func DeleteAll(tab table) error {
	db := dbMap[tab.DB()]
	_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", tab.Tab()))
	return err
}

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

func Count(tab table, where *Where) (int, error) {
	var stmtStr string
	if where != nil {
		stmtStr = fmt.Sprintf("SELECT COUNT(*) FROM %s.%s WHERE %s", tab.DB(), tab.Tab(), where.String())
	} else {
		stmtStr = fmt.Sprintf("SELECT COUNT(*) FROM %s.%s", tab.DB(), tab.Tab())
	}
	db := dbMap[tab.DB()]
	row := db.QueryRow(stmtStr)
	var num int
	err := row.Scan(&num)
	if err != nil {
		return 0, err
	}
	return num, nil
}
