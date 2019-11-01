package mock_nborm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"gotest.tools/assert"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wangjun861205/nbcolor"
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
			defer func() {
				user.SetForDelete()
				if _, err := nborm.Delete(db, user); err != nil {
					t.Error(err)
				}
			}()
			code := randString(16)
			user.IntelUserCode.SetString(code)
			user.Phone.SetString(code)
			if err := nborm.InsertOne(db, user); err != nil {
				t.Error(err)
				return
			}
			u := model.NewUser()
			u.SelectFields(&u.IntelUserCode)
			u.Phone.AndWhere("=", code)
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, u.IntelUserCode.AnyValue(), code, "intelusercode is %s", u.IntelUserCode.AnyValue())
		},
	},
	{
		"delete",
		func(t *testing.T) {
			u := model.NewUser()
			defer bulkDelete(t, u)
			code := randString(16)
			u.IntelUserCode.SetString(code)
			u.Phone.SetString(code)
			if err := nborm.InsertOne(db, u); err != nil {
				t.Error(err)
				return
			}
			u.Phone.AndWhere("=", code)
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
			u.SetForDelete()
			if _, err := nborm.Delete(db, u); err != nil {
				t.Error(err)
				return
			}
		},
	},
	{
		"value update",
		func(t *testing.T) {
			u := model.NewUser()
			defer bulkDelete(t, u)
			u.IntelUserCode.SetString("test")
			u.Phone.SetString("13793148690")
			if err := nborm.InsertOne(db, u); err != nil {
				t.Error(err)
				return
			}
			u.Phone.AndWhere("=", "13793148690")
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
			defer bulkDelete(t, u)
			u.IntelUserCode.SetString("test")
			u.Phone.SetString("13793148690")
			if err := nborm.InsertOne(db, u); err != nil {
				t.Error(err)
				return
			}
			u.IntelUserCode.AndWhere("=", "test")
			u.Email.Update(nborm.NewExpr("CONCAT('hello', 'world')"))
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
			assert.Equal(t, u.Email.AnyValue(), "helloworld", "email is %s", u.Email.AnyValue())
		},
	},
	{
		"group where query",
		func(t *testing.T) {
			u := model.NewUser()
			defer bulkDelete(t, u)
			u.IntelUserCode.SetString("test")
			u.Phone.SetString("13793148690")
			u.Email.SetString("email")
			if err := nborm.InsertOne(db, u); err != nil {
				t.Error(err)
				return
			}
			u.IntelUserCode.ForSelect()
			u.AndModelWhereGroup(u.Phone.AndEqWhere("13793148690"), u.Email.AndEqWhere("email"))
			u.OrModelWhereGroup(u.IntelUserCode.AndEqWhere("aaa"), u.IntelUserCode.AndEqWhere("bbb"))
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, u.IntelUserCode.AnyValue(), "test", "intelusercode is %s", u.IntelUserCode.AnyValue())
		},
	},
	{
		"two table join query",
		func(t *testing.T) {
			u := model.NewUser()
			defer bulkDelete(t, u)
			u.IntelUserCode.SetString("test")
			u.Phone.SetString("13793148690")
			if err := nborm.InsertOne(db, u); err != nil {
				t.Error(err)
				return
			}
			b := model.NewStudentbasicinfo()
			defer func() {
				b.SetForDelete()
				if _, err := nborm.Delete(db, b); err != nil {
					fmt.Println(nbcolor.Red(err))
				}
			}()
			b.IntelUserCode.SetString("test")
			b.Class.SetString("class")
			if err := nborm.InsertOne(db, b); err != nil {
				t.Error(err)
				return
			}
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
			defer bulkDelete(t, user)
			user.IntelUserCode.SetString("test")
			user.Phone.SetString("123456789")
			if err := nborm.InsertOne(db, user); err != nil {
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
			user.Phone.ForSelect()
			user.IntelUserCode.AndWhereEq("13793148690")
			if err := nborm.Query(db, user); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, user.Phone.AnyValue(), "123456789")
		},
	},
	{
		"group by selected fields",
		func(t *testing.T) {
			user := model.NewUser()
			info := model.NewStudentbasicinfo()
			defer func() {
				bulkDelete(t, user)
				bulkDelete(t, info)
			}()
			for i := 0; i < 10; i++ {
				user.Id.SetNull()
				user.IntelUserCode.SetString(fmt.Sprintf("%d", i))
				if err := nborm.InsertOne(db, user); err != nil {
					t.Error(err)
					return
				}
				info.Id.SetNull()
				info.IntelUserCode.SetString(fmt.Sprintf("%d", i))
				if err := nborm.InsertOne(db, info); err != nil {
					t.Error(err)
					return
				}
			}
			user.SelectAll()
			count := user.IntAgg(nborm.NewExpr("COUNT(*)"), "count")
			user.BasicInfo.SetForJoin()
			user.GroupBySelectedFields()
			if err := nborm.Query(db, user); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, count.AnyValue(), 1)
		},
	},
	{
		"uall",
		func(t *testing.T) {
			user := model.NewUser()
			defer bulkDelete(t, user)
			user.IntelUserCode.SetString("abc")
			user.Phone.SetString("13793148690")
			if err := nborm.InsertOne(db, user); err != nil {
				t.Error(err)
				return
			}
			user.Email.SetString("444055828@qq.com")
			user.UAll()
			isInsert, err := nborm.InsertOrUpdateOne(db, user)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Assert(t, !isInsert, "is insert not update")
		},
	},
	{
		"eq query test",
		func(t *testing.T) {
			user := model.NewUser()
			defer func() {
				user.SetForDelete()
				if _, err := nborm.Delete(db, user); err != nil {
					t.Error(err)
				}
			}()
			user.IntelUserCode.SetString("1234567")
			user.Phone.SetString("13793148690")
			if err := nborm.InsertOne(db, user); err != nil {
				t.Error(err)
				return
			}
			u := model.NewUser()
			u.Phone.ForSelect()
			u.IntelUserCode.AndWhereEq("1234567")
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, u.Phone.AnyValue(), "13793148690", "phone is %s", u.Phone.AnyValue())
		},
	},
	{
		"triple table join",
		func(t *testing.T) {
			grade := model.NewGrade()
			defer bulkDelete(t, grade)
			gradecode := randString(16)
			grade.GradeCode.SetString(gradecode)
			grade.GradeName.SetString(gradecode)
			if err := nborm.InsertOne(db, grade); err != nil {
				t.Error(err)
				return
			}
			class := model.NewClass()
			defer bulkDelete(t, class)
			classcode := randString(16)
			class.ClassName.SetString(classcode)
			class.ClassCode.SetString(classcode)
			class.Grade.SetString(gradecode)
			if err := nborm.InsertOne(db, class); err != nil {
				t.Error(err)
				return
			}
			info := model.NewStudentbasicinfo()
			defer bulkDelete(t, info)
			intelusercode := randString(16)
			info.IntelUserCode.SetString(intelusercode)
			info.Class.SetString(classcode)
			if err := nborm.InsertOne(db, info); err != nil {
				t.Error(err)
				return
			}
			user := model.NewUser()
			defer bulkDelete(t, user)
			user.IntelUserCode.SetString(intelusercode)
			if err := nborm.InsertOne(db, user); err != nil {
				t.Error(err)
				return
			}
			u := model.NewUser()
			u.StudentClass.SetForJoin()
			u.StudentClass.ClassCode.ForSelect()
			u.IntelUserCode.AndWhereEq(intelusercode)
			if err := nborm.Query(db, u); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, u.StudentClass.ClassCode.AnyValue(), classcode)
		},
	},
	{
		"back query",
		func(t *testing.T) {
			g := model.NewGrade()
			defer bulkDelete(t, g)
			gradecode := randString(16)
			g.GradeName.SetString(gradecode)
			g.GradeCode.SetString(gradecode)
			if err := nborm.InsertOne(db, g); err != nil {
				t.Error(err)
				return
			}
			c := model.NewClass()
			defer bulkDelete(t, c)
			classcode := randString(16)
			c.Grade.SetString(gradecode)
			c.ClassName.SetString(classcode)
			c.ClassCode.SetString(classcode)
			if err := nborm.InsertOne(db, c); err != nil {
				t.Error(err)
				return
			}
			s := model.NewStudentbasicinfo()
			defer bulkDelete(t, s)
			for i := 0; i < 10; i++ {
				s.Id.SetNull()
				intelusercode := randString(16)
				s.IntelUserCode.SetString(intelusercode)
				s.Class.SetString(classcode)
				if err := nborm.InsertOne(db, s); err != nil {
					t.Error(err)
					return
				}
			}
			c.SelectFields(&c.Id)
			c.ClassCode.AndWhereEq(classcode)
			if err := nborm.Query(db, c); err != nil {
				t.Error(err)
				return
			}
			c.Students.Id.ForSelect()
			if err := nborm.BackQuery(db, c.Students); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, c.Students.Len(), 10)
		},
	},
	{
		"unmarshal",
		func(t *testing.T) {
			u := model.NewUser()
			b := []byte(`
			{
				"Name": "test",
				"InsertDatetime": "2019-10-30 15:04:05"
			}`)
			if err := json.Unmarshal(b, u); err != nil {
				t.Error(err)
				return
			}
			y, m, d := u.InsertDatetime.AnyValue().Date()
			assert.Assert(t, y == 2019 && m == 10 && d == 30)
		},
	},
	{
		"unmarshal list",
		func(t *testing.T) {
			us := model.NewUserList()
			b := []byte(`
			{
				"List": [{
					"Name": "test",
					"InsertDatetime": "2019-10-30 15:04:05"
				}],
				"Total": 1
			}`)
			if err := json.Unmarshal(b, us); err != nil {
				t.Error(err)
				return
			}
			y, m, d := us.List[0].InsertDatetime.AnyValue().Date()
			assert.Assert(t, y == 2019 && m == 10 && d == 30)
		},
	},
	{
		"unmarshal list meta",
		func(t *testing.T) {
			us := model.NewUserList()
			b := []byte(`
			{
				"Name": "test",
				"InsertDatetime": "2019-10-30 15:04:05"
			}`)
			if err := us.UnmarshalMeta(b); err != nil {
				t.Error(err)
				return
			}
			y, m, d := us.InsertDatetime.AnyValue().Date()
			assert.Assert(t, y == 2019 && m == 10 && d == 30)
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
	for i := 0; i < number; i++ {
		for _, fieldInfo := range m.FieldInfos() {
			if fieldInfo.ColName != "Id" {
				switch f := fieldInfo.Field.(type) {
				case *nborm.String:
					f.SetString(randString(8))
				case *nborm.Int:
					f.SetInt(rand.Intn(4))
				}
			} else {
				fieldInfo.Field.(*nborm.Int).SetInt(fieldInfo.Field.(*nborm.Int).AnyValue() + 1)
			}
		}
		if err := nborm.InsertOne(db, m); err != nil {
			return err
		}
	}
	return nil
}

func bulkDelete(t *testing.T, m nborm.Model) {
	m.SetForDelete()
	if _, err := nborm.Delete(db, m); err != nil {
		t.Error(err)
	}
}

func TestNBorm(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}
