package nborm

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/wangjun861205/nbcolor"
)

// InsertOrUpdateOne 插入或更新
func InsertOrUpdateOne(exe Executor, model Model) (isInsert bool, err error) {
	var builder strings.Builder
	values := make([]interface{}, 0, 16)
	genInsertOrUpdateStmt(model, &builder, &values)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", builder.String(), values)))
	}
	res, err := exe.Exec(builder.String(), values...)
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
		if model.AutoIncField() != nil {
			model.AutoIncField().set(int(lastInsertID))
		}
	}
	return
}

// InsertOne 插入
func InsertOne(exe Executor, model Model) error {
	var builder strings.Builder
	values := make([]interface{}, 0, 16)
	genInsertStmt(model, &builder, &values)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", builder.String(), values)))
	}
	res, err := exe.Exec(builder.String(), values...)
	if err != nil {
		return err
	}
	if model.AutoIncField() != nil {
		lid, err := res.LastInsertId()
		if err != nil {
			return err
		}
		model.AutoIncField().set(int(lid))
	}
	model.addModelStatus(synced)
	return nil
}

// Query 查询
func Query(exe Executor, m Model) error {
	var builder strings.Builder
	values := make([]interface{}, 0, 16)
	genSelectStmt(m, &builder, &values)
	if DEBUG {
		log.Println(nbcolor.Green(fmt.Sprintf("%s %v", builder.String(), values)))

	}
	return queryAndScan(exe, m, builder.String(), values...)
}

// CacheQuery 缓存查询
func CacheQuery(exe Executor, m Model, timeout time.Duration) error {
	var builder strings.Builder
	values := make([]interface{}, 0, 16)
	genSelectStmt(m, &builder, &values)
	if DEBUG {
		log.Println(nbcolor.Green(fmt.Sprintf("%s %v", builder.String(), values)))
	}
	stmt := builder.String()
	for _, val := range values {
		builder.WriteString(fmt.Sprintf("%v", val))
	}
	hashValue := fmt.Sprintf("%x", md5.Sum([]byte(stmt)))
	if l, ok := m.(ModelList); ok {
		if l.GetListCache(hashValue, timeout) {
			if DEBUG {
				log.Println(nbcolor.Yellow("using cache"))
			}
			return nil
		}
	} else {
		if m.GetCache(hashValue, timeout) {
			if DEBUG {
				log.Println(nbcolor.Yellow("using cache"))
			}
			return nil
		}
	}
	if err := queryAndScan(exe, m, stmt, values...); err != nil {
		return err
	}
	if l, ok := m.(ModelList); ok {
		l.SetListCache(hashValue)
	} else {
		m.SetCache(hashValue)
	}
	return nil
}

// BackQuery 反向关联查询
func BackQuery(exe Executor, model Model) error {
	var builder strings.Builder
	values := make([]interface{}, 0, 16)
	genBackQueryStmt(model, &builder, &values)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", builder.String(), values)))
	}
	return queryAndScan(exe, model, builder.String(), values...)

}

// Update 更新
func Update(exe Executor, model Model) (sql.Result, error) {
	var builder strings.Builder
	values := make([]interface{}, 0, 16)
	genUpdateStmt(model, &builder, &values)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", builder.String(), values)))
	}
	return exe.Exec(builder.String(), values...)
}

// Delete 删除
func Delete(exe Executor, model Model) (sql.Result, error) {
	var builder strings.Builder
	values := make([]interface{}, 0, 16)
	genDeleteStmt(model, &builder, &values)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", builder.String(), values)))
	}
	return exe.Exec(builder.String(), values...)
}
