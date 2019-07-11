package test

import (
	"database/sql"
	"encoding/json"
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
	// {
	// 	name: "query",
	// 	f: func() error {
	// 		j := model.NewEnterpriseJobList()
	// 		j.Address.AndWhere("IN", []string{"address1", "address2"})
	// 		j.AndExprWhere(nborm.NewExpr("(@ = 'test1' OR @ = 'test2')", &j.Comment, &j.Comment))
	// 		j.EnterpriseID.DscOrder()
	// 		j.SetLimit(10, 0)
	// 		if err := nborm.Query(db, j); err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "backQuery",
	// 	f: func() error {
	// 		j := model.NewEnterpriseJobList()
	// 		j.EnterpriseID.AndWhere("=", 1)
	// 		j.StudentResumes.IntelUserCode.AndWhere("=", "xxx")
	// 		j.StudentResumes.ForReverseQuery()
	// 		if err := nborm.Query(db, j); err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "spanQuary",
	// 	f: func() error {
	// 		j := model.NewEnterpriseJobList()
	// 		j.StudentResumes.InitRel()
	// 		j.StudentResumes.StudentSkill.ForReverseQuery()
	// 		j.StudentResumes.StudentSkill.ID.AndWhere("=", 1)
	// 		if err := nborm.Query(db, j); err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "isNullOpQuery",
	// 	f: func() error {
	// 		j := model.NewEnterpriseJobList()
	// 		j.EnterpriseID.AndWhere("IS NOT NULL", nil)
	// 		if err := nborm.Query(db, j); err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "addedCondQuery",
	// 	f: func() error {
	// 		a := model.NewEnterpriseAccountList()
	// 		a.Enterprise.InitRel()
	// 		a.Email.AndWhere("=", "parent_query")
	// 		a.Enterprise.ForReverseQuery()
	// 		a.Enterprise.UniformCode.AndWhere("=", "local_query")
	// 		a.Enterprise.Sector.ForReverseQuery()
	// 		a.Enterprise.Sector.Name.AndWhere("=", "children_query")
	// 		if err := nborm.Query(db, a); err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "update",
	// 	f: func() error {
	// 		e := model.NewEnterpriseAccountList()
	// 		e.Email.SetUpdate("test")
	// 		e.Enterprise.ForReverseQuery()
	// 		e.Enterprise.ID.AndWhere("=", 1)
	// 		e.Enterprise.Contact.SetUpdate("yyy")
	// 		if _, err := nborm.Update(db, e); err != nil {
	// 			return err
	// 		}
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "joinQuery",
	// 	f: func() error {
	// 		aForDel := model.NewEnterpriseAccount()
	// 		eForDel := model.NewEnterprise()
	// 		if _, err := nborm.Delete(db, aForDel); err != nil {
	// 			return err
	// 		}
	// 		if _, err := nborm.Delete(db, eForDel); err != nil {
	// 			return err
	// 		}
	// 		a := model.NewEnterpriseAccount()
	// 		a.Email.SetString("email")
	// 		a.Password.SetString("password")
	// 		a.Phone.SetString("phone")
	// 		if err := nborm.InsertOne(db, a); err != nil {
	// 			return err
	// 		}
	// 		for i := 0; i < 10; i++ {
	// 			e := model.NewEnterprise()
	// 			e.AccountID.SetInt(a.ID.Int())
	// 			e.Contact.SetString(fmt.Sprintf("contact%d", i))
	// 			e.EmployeeFromThis.SetInt(i)
	// 			e.Introduction.SetString(fmt.Sprintf("introduction%d", i))
	// 			e.Name.SetString(fmt.Sprintf("name%d", i))
	// 			e.NatureID.SetInt(i)
	// 			e.RegisterAddress.SetString(fmt.Sprintf("register_address%d", i))
	// 			e.RegisterCityID.SetInt(i)
	// 			e.ScopeID.SetInt(i)
	// 			e.SectorID.SetInt(i)
	// 			e.UniformCode.SetString(fmt.Sprintf("uniform_code%d", i))
	// 			e.Website.SetString(fmt.Sprintf("website%d", i))
	// 			e.ZipCode.SetString(fmt.Sprintf("zipcode%d", i))
	// 			if err := nborm.InsertOne(db, e); err != nil {
	// 				return err
	// 			}
	// 		}
	// 		qa := model.NewEnterpriseAccountList()
	// 		qa.Enterprise.SetForJoin()
	// 		qa.Enterprise.InitRel()
	// 		qa.Enterprise.Sector.SetForJoin()
	// 		qa.Enterprise.Sector.Name.ForSelect()
	// 		qa.Enterprise.Sector.Name.AndWhere("=", "name 3")
	// 		// qa.Email.AndWhere("=", "email")
	// 		// qa.Enterprise.Contact.AndWhere(">", "contact5")
	// 		if err := nborm.JoinQuery(db, qa); err != nil {
	// 			return err
	// 		}
	// 		for _, e := range qa.Enterprise.List {
	// 			fmt.Println(nbcolor.Yellow(e.Contact.String()))
	// 		}
	// 		b, err := json.Marshal(qa)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(nbcolor.Yellow(string(b)))
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "setValueForJSON",
	// 	f: func() error {
	// 		e := model.NewEnterprise()
	// 		e.Contact.SetString("hello")
	// 		b, err := json.Marshal(e)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		ne := model.NewEnterprise()
	// 		nb, err := json.Marshal(ne)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(nbcolor.Yellow(string(b)))
	// 		fmt.Println(nbcolor.Yellow(string(nb)))
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "agg",
	// 	f: func() error {
	// 		j := model.NewEnterpriseJob()
	// 		j.Enterprise.AccountID.AndWhere(">", 1)
	// 		j.IntAgg(nborm.NewExpr("SUM(@)", &j.Vacancies), "Total")
	// 		j.StrAgg(nborm.NewExpr("GROUP_CONCAT(@)", &j.Comment), "Comment")
	// 		j.AndHaving(nborm.NewExpr("Total>5"))
	// 		j.CityID.GroupBy()
	// 		j.Address.GroupBy()
	// 		res, err := nborm.Agg(db, j)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(nbcolor.Yellow(res.ToSimpleList()))
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "unsignedTinyint",
	// 	f: func() error {
	// 		b := make([]byte, 0, 8)
	// 		// var i int64
	// 		if err := db.QueryRow("SELECT t FROM qdxg.test").Scan(&b); err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(nbcolor.Red(string(b)))
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "insertOrUpdate",
	// 	f: func() error {
	// 		flag := model.NewJobFlag()
	// 		flag.ID.SetInt(2)
	// 		flag.Name.SetString("name").SetUpdate("name 3")
	// 		flag.Type.SetString("type").SetUpdate("type 3")
	// 		flag.Status.SetString("正常")
	// 		flag.Comment.SetString("comment").SetUpdate("comment 3")
	// 		flag.Order.SetInt(1).SetUpdate(3)
	// 		isInsert, err := nborm.InsertOrUpdateOne(db, flag)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(nbcolor.Red(isInsert))
	// 		fmt.Println(nbcolor.Red(flag.ID.Int()))
	// 		return nil
	// 	},
	// },
	// {
	// 	name: "unmarshal",
	// 	f: func() error {
	// 		e := model.NewEnterprise()
	// 		if err := nborm.Query(db, e); err != nil {
	// 			return err
	// 		}
	// 		jb, err := json.Marshal(e)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		ne := model.NewEnterprise()
	// 		if err := json.Unmarshal(jb, ne); err != nil {
	// 			return err
	// 		}
	// 		nj, err := json.Marshal(ne)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(nbcolor.Yellow(string(nj)))
	// 		return nil
	// 	},
	// },
	{
		name: "newMarshalModel",
		f: func() error {
			e := model.NewEnterprise()
			e.Contact.ForSelect()
			if err := nborm.Query(db, e); err != nil {
				return err
			}
			bs, err := json.Marshal(e)
			if err != nil {
				return err
			}
			fmt.Println(nbcolor.Brown(string(bs)))
			return nil
		},
	},
	{
		name: "newMarshalModelList",
		f: func() error {
			es := model.NewEnterpriseList()
			es.Contact.ForSelect()
			es.Account.SetForJoin()
			es.Account.Email.ForSelect()
			if err := nborm.JoinQuery(db, es); err != nil {
				return err
			}
			bs, err := json.Marshal(es)
			if err != nil {
				return err
			}
			fmt.Println(nbcolor.Brown(string(bs)))
			nes := model.NewEnterpriseList()
			if err := json.Unmarshal(bs, nes); err != nil {
				return err
			}
			fmt.Println(nbcolor.Brown(nes.List[0].Contact.String()))
			fmt.Println(nbcolor.Brown(nes.List[0].Account.Phone.String()))
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
