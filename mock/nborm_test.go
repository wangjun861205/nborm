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
	// {
	// 	name: "insert or update",
	// 	f: func(t *testing.T) {
	// 		acct := model.NewEmployAccount()
	// 		acct.ID.SetExpr(nborm.NewExpr("UUID()"))
	// 		acct.Phone.SetString("13793148692")
	// 		acct.Phone.Update("13793148693")
	// 		acct.CreateTime.Update(time.Now())
	// 		acct.Password.SetExpr(nborm.NewExpr("MD5(?)", "123456"))
	// 		if isInsert, err := nborm.InsertOrUpdateOne(db, acct); err != nil {
	// 			t.Error(err)
	// 		} else {
	// 			fmt.Println(isInsert)
	// 		}
	// 	},
	// },
	// {
	// 	name: "update",
	// 	f: func(t *testing.T) {
	// 		acct := model.NewEmployAccount()
	// 		acct.Phone.AndWhere("=", "13793148690")
	// 		acct.Password.Update(nborm.NewExpr("MD5(?)", "123456"))
	// 		acct.CreateTime.Update(time.Now())
	// 		if _, err := nborm.Update(db, acct); err != nil {
	// 			t.Error(err)
	// 			return
	// 		}
	// 	},
	// },
	// {
	// 	name: "query",
	// 	f: func(t *testing.T) {
	// 		acct := model.NewEmployAccount()
	// 		acct.Phone.AndWhere("=", "15665793333")
	// 		acct.SelectAll()
	// 		if err := nborm.Query(db, acct); err != nil {
	// 			t.Error(err)
	// 			return
	// 		}
	// 		fmt.Println(acct)
	// 	},
	// },
	// {
	// 	name: "join query",
	// 	f: func(t *testing.T) {
	// 		acct := model.NewEmployAccount()
	// 		acct.SelectAll()
	// 		acct.Phone.AndWhere("=", "13793148690")
	// 		acct.Enterprise.SetForLeftJoin()
	// 		acct.Enterprise.SelectAll()
	// 		if err := nborm.Query(db, acct); err != nil {
	// 			t.Error(err)
	// 			return
	// 		}
	// 		fmt.Println(acct)
	// 	},
	// },
	// {
	// 	name: "delete",
	// 	f: func(t *testing.T) {
	// 		acct := model.NewEmployAccount()
	// 		acct.Phone.AndWhere("=", "13793148691")
	// 		acct.SetForDelete()
	// 		acct.Enterprise.SetForJoin()
	// 		acct.Enterprise.SetForDelete()
	// 		if _, err := nborm.Delete(db, acct); err != nil {
	// 			t.Error(err)
	// 			return
	// 		}
	// 	},
	// },
	{
		name: "aggregate",
		f: func(t *testing.T) {
			ent := model.NewEmployEnterpriseList()
			ent.IntAgg(nborm.NewExpr("SUM(IF(@=1 OR @=2, 1, 0))", &ent.Status, &ent.Status), "count")
			if err := nborm.Query(db, ent); err != nil {
				t.Error(err)
				return
			}
			fmt.Println(ent)
		},
	},
}

func TestNBorm(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}
