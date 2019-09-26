package mock_nborm

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wangjun861205/nborm"
	"github.com/wangjun861205/nborm/mock/model"
)

var db *sql.DB

func init() {
	nborm.SetDebug(true)
	d, err := sql.Open("mysql", "root:Ydkj@0517@tcp(192.168.175.2:3306)/employment")
	if err != nil {
		panic(err)
	}
	db = d
}

func TestNBorm(t *testing.T) {
	ents := model.NewEmployEnterpriseList()
	ents.ID.ForSelect()
	ents.EmployFromThis.AndWhere(">", 10)
	if err := nborm.Query(db, ents); err != nil {
		t.Fatal(err)
	}
	fmt.Println(ents)
}
