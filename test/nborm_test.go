package test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/wangjun861205/nbcolor"
	"github.com/wangjun861205/nborm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wangjun861205/nborm/test/model"
)

var db *sql.DB

type test struct {
	name string
	f    func() error
}

var testList = []test{
	{
		name: "deleteTest",
		f: func() error {
			e := model.NewEnterpriseJob()
			if _, err := nborm.Delete(db, e); err != nil {
				return err
			}
			return nil
		},
	},
	{
		name: "insertTest",
		f: func() error {
			for i := 0; i < 10; i++ {
				j := model.NewEnterpriseJob()
				j.Address.SetString(fmt.Sprintf("address%d", i))
				j.CityID.SetInt(i)
				j.Comment.SetString(fmt.Sprintf("comment%d", i))
				j.DegreeID.SetInt(i)
				j.Description.SetString(fmt.Sprintf("Description%d", i))
				j.EnterpriseID.SetInt(i)
				j.ExpiredAt.SetDatetime(time.Now())
				j.Gender.SetString("男")
				j.LanguageSkillID.SetInt(i)
				j.MajorCode.SetString(fmt.Sprintf("major_code%d", i))
				j.Name.SetString(fmt.Sprintf("name%d", i))
				j.SalaryRangeID.SetInt(i)
				j.Status.SetString("待审核")
				j.Vacancies.SetInt(i)
				j.Welfare.SetString(fmt.Sprintf("welfare%d", 1))
				j.TypeID.SetInt(i)
				if err := nborm.InsertOne(db, j); err != nil {
					return err
				}
			}
			return nil
		},
	},
	{
		name: "queryTest",
		f: func() error {
			j := model.NewEnterpriseJobList()
			j.Address.AndWhere("IN", []string{"address1", "address2"})
			j.AndExprWhere(nborm.NewExpr("(@ = 'test1' OR @ = 'test2')", &j.Comment, &j.Comment))
			if err := nborm.Query(db, j, -1, -1); err != nil {
				return err
			}
			return nil
		},
	},
	{
		name: "backQuery",
		f: func() error {
			j := model.NewEnterpriseJobList()
			j.EnterpriseID.AndWhere("=", 1)
			j.StudentResumes.IntelUserCode.AndWhere("=", "xxx")
			if err := nborm.Query(db, j, -1, -1); err != nil {
				return err
			}
			return nil
		},
	},
	{
		name: "spanQuary",
		f: func() error {
			j := model.NewEnterpriseJobList()
			j.StudentResumes.InitRel()
			j.StudentResumes.StudentSkill.ID.AndWhere("=", 1)
			if err := nborm.Query(db, j.StudentResumes, -1, -1); err != nil {
				return err
			}
			return nil
		},
	},
	{
		name: "isNullOpQuery",
		f: func() error {
			j := model.NewEnterpriseJobList()
			j.EnterpriseID.AndWhere("IS NOT NULL", nil)
			if err := nborm.Query(db, j, -1, -1); err != nil {
				return err
			}
			return nil
		},
	},
}

func TestNBorm(t *testing.T) {
	for _, test := range testList {
		if err := test.f(); err != nil {
			t.Fatal(nbcolor.Red(fmt.Sprintf("%s failed(%v)", test.name, err)))
		}
		fmt.Println(nbcolor.LightCyan(fmt.Sprintf("%s passed", test.name)))
	}
}

func init() {
	conn, err := sql.Open("mysql", "wangjun:Wt20110523@tcp(127.0.0.1:3306)/qdxg")
	if err != nil {
		panic(err)
	}
	db = conn
	nborm.SetDebug(true)
}
