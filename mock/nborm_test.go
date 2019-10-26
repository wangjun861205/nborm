package mock_nborm

import (
	"database/sql"
	"math/rand"
	"strings"
	"testing"
	"time"

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
	{
		"insert or update",
		func(t *testing.T) {
			user := model.NewUser()
			user.SetForDelete()
			if _, err := nborm.Delete(db, user); err != nil {
				t.Error(err)
				return
			}
			basicInfo := model.NewStudentbasicinfo()
			basicInfo.SetForDelete()
			if _, err := nborm.Delete(db, basicInfo); err != nil {
				t.Error(err)
				return
			}
			if err := bulkInsert(user, 1); err != nil {
				t.Error(err)
				return
			}
			user.Phone.Update("13793148690")
			isInsert, err := nborm.InsertOrUpdateOne(db, user)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Assert(t, !isInsert, "still insert")
		},
	},
}

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randString(charNum int) string {
	var builder strings.Builder
	for i := 0; i < charNum; i++ {
		builder.WriteRune(runes[rand.Intn(len(runes))])
	}
	return builder.String()
}

func bulkInsert(m nborm.Model, number int) error {
	randStr := randString(8)
	for i := 0; i < number; i++ {
		for _, fieldInfo := range m.FieldInfos() {
			if fieldInfo.ColName != "Id" {
				switch f := fieldInfo.Field.(type) {
				case *nborm.String:
					f.SetString(randStr)
				case *nborm.Int:
					f.SetInt(rand.Intn(4))
				}
			}
		}
		if err := nborm.InsertOne(db, m); err != nil {
			return err
		}
	}
	return nil
}

func TestNBorm(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}
