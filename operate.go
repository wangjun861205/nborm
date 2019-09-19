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
	stmt, values := genInsertOrUpdateStmt(model)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", stmt, values)))
	}
	res, err := exe.Exec(stmt, values...)
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
			model.AutoIncField().Set(int(lastInsertID))
		}
	}
	return
}

// InsertOne 插入
func InsertOne(exe Executor, model Model) error {
	stmt, values := genInsertStmt(model)
	res, err := exe.Exec(stmt, values...)
	if err != nil {
		return err
	}
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", stmt, values)))
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

// Query 查询
func Query(exe Executor, m Model) error {
	stmt, values := genSelectStmt(m)
	if DEBUG {
		log.Println(nbcolor.Green(fmt.Sprintf("%s %v", stmt, values)))

	}
	return queryAndScan(exe, m, stmt, values...)
}

// CacheQuery 缓存查询
func CacheQuery(exe Executor, m Model, timeout time.Duration) error {
	stmt, values := genSelectStmt(m)
	if DEBUG {
		log.Println(nbcolor.Green(fmt.Sprintf("%s %v", stmt, values)))
	}
	var builder strings.Builder
	builder.WriteString(stmt)
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
	stmt, values := genBackQueryStmt(model)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", stmt, values)))
	}
	return queryAndScan(exe, model, stmt, values...)

}

// Update 更新
func Update(exe Executor, model Model) (sql.Result, error) {
	stmt, values := genUpdateStmt(model)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", stmt, values)))
	}
	return exe.Exec(stmt, values...)
}

// Delete 删除
func Delete(exe Executor, model Model) (sql.Result, error) {
	// whereClause, whereValues := genSimpleWhereClause(model)
	// stmt := fmt.Sprintf("DELETE FROM %s %s", model.rawFullTabName(), whereClause)
	stmt, values := genDeleteStmt(model)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", stmt, values)))
	}
	return exe.Exec(stmt, values...)
}

// Exists 查询是否存在
func Exists(exe Executor, model Model) (bool, error) {
	stmt, values := genSelectStmt(model)
	stmt = fmt.Sprintf("SELECT EXISTS(%s)", stmt)
	if DEBUG {
		fmt.Println(nbcolor.Green(fmt.Sprintf("%s %v", stmt, values)))
	}
	var exists int
	if err := exe.QueryRow(stmt, values...).Scan(&exists); err != nil {
		return false, err
	}
	return exists == 1, nil
}
