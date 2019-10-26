package mock_nborm

import (
	"database/sql"
	"testing"

	"gotest.tools/assert"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wangjun861205/nborm"
	"github.com/wangjun861205/nborm/mock/model"
)

var db *sql.DB

func init() {
	nborm.SetDebug(true)
	d, err := sql.Open("mysql", "wangjun:Wt20110523@tcp(localhost:3306)/qdxg")
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

	{
		name: "insert",
		f: func(t *testing.T) {
			user := model.NewUser()
			user.IntelUserCode.SetString("test")
			user.Phone.SetString("test")
			if err := nborm.InsertOne(db, user); err != nil {
				t.Error(err)
				return
			}
			u := model.NewUser()
			u.SelectFields(&u.IntelUserCode)
			u.Phone.AndWhere("=", "test")

			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, u.IntelUserCode.AnyValue(), "test", "intelusercode is %s", u.IntelUserCode.AnyValue())
		},
	},
	{
		"delete",
		func(t *testing.T) {
			u := model.NewUser()
			u.Phone.AndWhere("=", "test")
			count := u.IntAgg(nborm.NewExpr("COUNT(*)"), "count")
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Assert(t, count.AnyValue() > 0, "no rows for delete")
			u.SetForDelete()
			if _, err := nborm.Delete(db, u); err != nil {
				t.Error(err)
				return
			}
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Assert(t, count.AnyValue() == 0, "still exists rows which should be deleted")
		},
	},
	{
		"expr insert",
		func(t *testing.T) {
			u := model.NewUser()
			u.IntelUserCode.SetExpr(nborm.NewExpr("CONCAT('hello', 'world')"))
			u.Phone.SetString("test")
			if err := nborm.InsertOne(db, u); err != nil {
				t.Error(err)
				return
			}
			u.IntelUserCode.ForSelect()
			u.Phone.AndW()
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, u.IntelUserCode.AnyValue(), "helloworld", "intelusercode is %s", u.IntelUserCode.AnyValue())
		},
	},
	{
		"value update",
		func(t *testing.T) {
			u := model.NewUser()
			u.Phone.AndWhere("=", "test")
			u.Email.Update("test email")
			if res, err := nborm.Update(db, u); err != nil {
				t.Error(err)
				return
			} else if num, err := res.RowsAffected(); err != nil {
				t.Error(err)
				return
			} else if num == 0 {
				t.Error("no rows updated")
				return
			}
			u.Email.ForSelect()
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, u.Email.AnyValue(), "test email", "email is %s", u.Email.AnyValue())
		},
	},
	{
		"expr update",
		func(t *testing.T) {
			u := model.NewUser()
			u.Phone.AndWhere("=", "test")
			u.Email.Update(nborm.NewExpr("SUBSTR(@, 6, 5)", &u.Email))
			if res, err := nborm.Update(db, u); err != nil {
				t.Error(err)
				return
			} else if num, err := res.RowsAffected(); err != nil {
				t.Error(err)
				return
			} else if num == 0 {
				t.Error("no rows updated")
				return
			}
			u.Email.ForSelect()
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, u.Email.AnyValue(), "email", "email is %s", u.Email.AnyValue())
		},
	},
	{
		"group where query",
		func(t *testing.T) {
			u := model.NewUser()
			u.IntelUserCode.ForSelect()
			u.AndModelWhereGroup(u.Phone.AndEqWhere("test"), u.Email.AndEqWhere("email"))
			u.OrModelWhereGroup(u.IntelUserCode.AndEqWhere("aaa"), u.IntelUserCode.AndEqWhere("bbb"))
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, u.IntelUserCode.AnyValue(), "helloworld", "intelusercode is %s", u.IntelUserCode.AnyValue())
		},
	},
	{
		"two table join query",
		func(t *testing.T) {
			l := model.NewUserList()
			l.IntelUserCode.ForSelect()
			l.BasicInfo.SetForJoin()
			l.BasicInfo.IntelUserCode.ForSelect()
			l.SetLimit(10, 0)
			if err := nborm.Query(db, l); err != nil {
				t.Error(err)
				return
			}
			for _, u := range l.List {
				assert.Equal(t, u.IntelUserCode.AnyValue(), u.BasicInfo.IntelUserCode.AnyValue(),
					"user intelusercode is %s, basicinfo intelusercode is %s", u.IntelUserCode.AnyValue(), u.BasicInfo.IntelUserCode.AnyValue())
			}
		},
	},
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
	// {
	// 	name: "aggregate",
	// 	f: func(t *testing.T) {
	// 		ent := model.NewEmployEnterpriseList()
	// 		ent.IntAgg(nborm.NewExpr("SUM(IF(@=1 OR @=2, 1, 0))", &ent.Status, &ent.Status), "count")
	// 		if err := nborm.Query(db, ent); err != nil {
	// 			t.Error(err)
	// 			return
	// 		}
	// 		fmt.Println(ent)
	// 	},
	// },
}

func TestNBorm(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}
