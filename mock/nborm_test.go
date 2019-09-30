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
	// ents.Account.SetForJoin().SelectAll()
	ents.Name.ForSelect()
	ents.SID.ForSelect()
	ents.EmployFromThis.AndWhere(">", 1)
	if err := nborm.Query(db, ents); err != nil {
		t.Fatal(err)
	}
	for _, ent := range ents.List {
		// fmt.Println(ent)
		ent.Account.SelectAll()
		if err := nborm.BackQuery(db, ent.Account); err != nil {
			t.Fatal(err)
		}
	}
	for _, ent := range ents.List {
		fmt.Println(ent.Account)
	}
}
