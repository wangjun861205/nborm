package nborm

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var connMap = map[string]*sql.DB{}
var connLock sync.RWMutex

func getConn(db string) *sql.DB {
	db = escap(db)
	connLock.RLock()
	conn, ok := connMap[db]
	if !ok {
		connLock.RUnlock()
		var err error
		conn, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, db))
		if err != nil {
			panic(err)
		}
		connLock.Lock()
		connMap[db] = conn
		connLock.Unlock()
		return conn
	}
	connLock.RUnlock()
	return conn
}

func Begin() (*sql.Tx, error) {
	db := getConn("information_schema")
	return db.Begin()
}

//CloseConns close all connections which is stored in connMap
func CloseConns() error {
	connLock.Lock()
	defer connLock.Unlock()
	for _, conn := range connMap {
		err := conn.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
