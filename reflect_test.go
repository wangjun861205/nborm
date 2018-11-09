package nborm

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

type Auth struct {
	Id            *IntField `auto_increment:"true"`
	Username      *StringField
	Password      *StringField
	Phone         *StringField `primary_key:"true"`
	Status        *IntField
	SessionId     *StringField
	ExpireTime    *DatetimeField
	Email         *StringField `unique:"true"`
	CreateTime    *DatetimeField
	LastLoginTime *DatetimeField
	IsSuperUser   *BoolField
	_isSync       bool
}

func (m *Auth) DB() string {
	return "bk_dalian"
}

func (m *Auth) Tab() string {
	return "auth"
}

func (m *Auth) Fields() []Field {
	return nil
}

func (m *Auth) SetSync(val bool) {
	m._isSync = val
}

func (m *Auth) GetSync() bool {
	return m._isSync
}

type AuthList []*Auth

func (l *AuthList) DB() string {
	return "bk_dalian"
}

func (l *AuthList) Tab() string {
	return "auth"
}

func TestReflect(t *testing.T) {
	db, err := sql.Open("mysql", "wangjun:Wt20110523@tcp(127.0.0.1:12345)/bk_dalian")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM auth")
	if err != nil {
		log.Fatal(err)
	}
	auths := make(AuthList, 0, 8)
	err = scanRows(&auths, rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(auths[0].Username.Get())
}
