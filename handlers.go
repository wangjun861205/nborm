package nborm

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func All(l modelList) error {
	stmtStr := fmt.Sprintf("SELECT * FROM %s", l.Tab())
	rows, err := dbMap[l.DB()].Query(stmtStr)
	if err != nil {
		return err
	}
	return scanRows(l, rows)
}

func Query(tab table, where *Where) error {
	db := dbMap[tab.DB()]
	if where == nil {
		switch obj := tab.(type) {
		case Model:
			where = genWhere(obj)
			stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", tab.Tab(), where.String())
			row := db.QueryRow(stmtStr)
			return scanRow(obj, row)
		case modelList:
			for i := 0; i < obj.Len(); i++ {
				where = genWhere(obj.Index(i))
				stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", tab.Tab(), where.String())
				row := db.QueryRow(stmtStr)
				err := scanRow(obj.Index(i), row)
				if err != nil {
					return err
				}
			}
			return nil
		}
	}
	stmtStr := fmt.Sprintf("SELECT * FROM %s WHERE %s", tab.Tab(), where.String())
	switch obj := tab.(type) {
	case Model:
		row := db.QueryRow(stmtStr)
		return scanRow(obj, row)
	case modelList:
		rows, err := db.Query(stmtStr)
		if err != nil {
			return err
		}
		return scanRows(obj, rows)
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
			joinList[i] = fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s", r.DstDB(), r.DstTab(), tab.DB(), tab.Tab(), r.SrcCol(),
				r.DstDB(), r.DstTab(), r.DstCol())
		case *ManyToMany:
			joinList[i] = fmt.Sprintf("JOIN %s.%s ON %s.%s.%s = %s.%s.%s JOIN %s.%s ON %s.%s.%s = %s.%s.%s", r.MidDB(), r.MidTab(), tab.DB(), tab.Tab(),
				r.SrcCol(), r.MidDB(), r.MidTab(), r.MidLeftCol(), r.DstDB(), r.DstTab(), r.MidDB(), r.MidTab(), r.MidRightCol(), r.DstDB(), r.DstTab(),
				r.DstCol())
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
		case modelList:
			for i := 0; i < obj.Len(); i++ {
				m := obj.Index(i)
				where = genWhere(m)
				stmtStr := fmt.Sprintf("SELECT %s.%s.* FROM %s.%s %s WHERE %s", tab.DB(), tab.Tab(), tab.DB(), tab.Tab(), strings.Join(joinList, " "), where.String())
				row := db.QueryRow(stmtStr)
				err := scanRow(m, row)
				if err != nil {
					return err
				}
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
		case modelList:
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
		case modelList:
			for i := 0; i < obj.Len(); i++ {
				m := obj.Index(i)
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
					return err
				}
				lastInsertId, err := res.LastInsertId()
				if err != nil {
					return err
				}
				inc.(*IntField).Set(lastInsertId, false)
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
		case modelList:
			info := getTableCache(obj.Index(0))
			if len(values) > 0 {
				valStrs := make([]string, len(values))
				for i, val := range values {
					valStrs[i] = val.String()
				}
				for i := 0; i < obj.Len(); i++ {
					stmtStr := fmt.Sprintf("UPDATE %s SET %s WHERE %s", obj.Tab(), strings.Join(valStrs, ", "), obj.Index(i).Fields()[info.inc].Where())
					_, err := db.Exec(stmtStr)
					if err != nil {
						return err
					}
					updateModel(obj.Index(i), values...)
				}
			} else {
				for i := 0; i < obj.Len(); i++ {
					fields := obj.Index(i).Fields()
					setValues := make([]string, 0, len(fields))
					for _, f := range fields {
						if !f.IsInc() && f.IsValid() {
							setValues = append(setValues, f.UpdateValue().String())
						}
					}
					stmtStr := fmt.Sprintf("UPDATE %s SET %s WHERE %s", obj.Tab(), strings.Join(setValues, ", "), fields[info.inc].Where())
					_, err := db.Exec(stmtStr)
					if err != nil {
						return err
					}
				}
			}
			return nil
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
		case modelList:
			for i := 0; i < obj.Len(); i++ {
				inc, fs := getInc(obj.Index(i))
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
				return Query(obj, nil)
			} else {
				return err
			}
		}
		lastInsertId, err := res.LastInsertId()
		if err != nil {
			return err
		}
		inc.(*IntField).Set(lastInsertId, false)
	case modelList:
		for i := 0; i < obj.Len(); i++ {
			inc, fs := getInc(obj.Index(i))
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
					err := Query(obj.Index(i), nil)
					if err != nil {
						return err
					}
					continue
				} else {
					return err
				}
			}
			lastInsertId, err := res.LastInsertId()
			if err != nil {
				return err
			}
			inc.(*IntField).Set(lastInsertId, false)
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
		case modelList:
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

func Sort(l modelList, reverse bool, fields ...Field) {
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

func Distinct(l modelList, fields ...Field) {
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
