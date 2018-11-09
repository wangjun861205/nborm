package nborm

import (
	"fmt"
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
	Test          *ForeignKey `source_column:"id" destination_column:"test.test.test"`
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
	auth := &Auth{}
	InitModel(auth)
	fmt.Println(auth.Test.dstCol)
}
