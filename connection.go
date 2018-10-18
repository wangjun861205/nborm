package nborm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var dbMap = map[string]*sql.DB{}
var dbLock sync.Mutex

func RegisterDB(username, password, address, db string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, address, db)
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	dbMap[db] = conn
	if db != "information_schema" {
		if _, err := os.Stat("dbinfo.json"); os.IsNotExist(err) {
			infoConn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/information_schema", username, password, address))
			if err != nil {
				panic(err)
			}
			dbMap["information_schema"] = infoConn
			GetDBInfo(db)
			MarshalDBInfo()
		} else {
			infoMap := make(map[string]map[string]*tableInfo)
			f, err := os.Open("dbinfo.json")
			if err != nil {
				panic(err)
			}
			b, err := ioutil.ReadAll(f)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(b, &infoMap)
			if err != nil {
				panic(err)
			}
			dbInfo = infoMap
		}
		initCacheByInfo()
	}
}

func CloseConns() error {
	dbLock.Lock()
	defer dbLock.Unlock()
	for _, db := range dbMap {
		err := db.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
