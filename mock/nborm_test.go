package mock_nborm

import (
	"database/sql"
	"testing"
	"time"

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

type test struct {
	name string
	f    func(t *testing.T)
}

var tests = []test{
	// {
	// 	name: "insert",
	// 	f: func(t *testing.T) {
	// 		acct := model.NewEmployAccount()
	// 		acct.ID.SetExpr(nborm.NewExpr("UUID()"))
	// 		acct.Phone.SetString("13793148691")
	// 		acct.Password.SetExpr(nborm.NewExpr("MD5(?)", "123456"))
	// 		if err := nborm.InsertOne(db, acct); err != nil {
	// 			t.Error(err)
	// 		}
	// 	},
	// },
	{
		name: "update",
		f: func(t *testing.T) {
			acct := model.NewEmployAccount()
			acct.Phone.AndWhere("=", "13793148690")
			acct.Password.Update(nborm.NewExpr("MD5(?)", "123456"))
			acct.CreateTime.Update(time.Now())
			if _, err := nborm.Update(db, acct); err != nil {
				t.Error(err)
				return
			}
		},
	},
}

func TestNBorm(t *testing.T) {
	for _, tt := range tests {
		if ok := t.Run(tt.name, tt.f); ok {
			t.Logf("%s test success", tt.name)
		} else {
			t.Logf("%s test field", tt.name)
		}
	}
}
