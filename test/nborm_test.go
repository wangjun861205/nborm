package test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/wangjun861205/nborm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wangjun861205/nborm/test/model"
)

var db *sql.DB

func TestDelete(t *testing.T) {
	e := model.NewEnterprise()
	e.ID.AndWhere("=", 1)
	e.Account.ID.AndWhere("=", 1)
	if res, err := nborm.Delete(db, e); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(res.RowsAffected())
	}

}

func TestInsert(t *testing.T) {
	e := model.NewEnterprise()
	e.Contact.SetString("contact")
	e.Introduction.SetString("introduction")
	e.NatureID.SetInt(1)
	e.SectorID.SetInt(2)
	e.ScopeID.SetInt(3)
	e.Name.SetString("name")
	e.RegisterCityID.SetInt(4)
	e.RegisterAddress.SetString("address")
	e.UniformCode.SetString("uniform code 1")
	e.Website.SetString("website")
	e.ZipCode.SetString("zipcode")
	e.EmployeeFromThis.SetInt(100)
	if err := nborm.InsertOne(db, e); err != nil {
		log.Fatal(err)
	}
}

func TestQuery(t *testing.T) {
	e := model.NewEnterprise()
	e.Account.EnterpriseID.AndWhere("=", 1)
	if err := nborm.QueryOne(db, e); err != nil {
		t.Fatal(err)
	}
	fmt.Println("success")
}

func init() {
	conn, err := sql.Open("mysql", "wangjun:Wt20110523@tcp(127.0.0.1:3306)/qdxg")
	if err != nil {
		panic(err)
	}
	db = conn
	nborm.SetDebug(true)
}
