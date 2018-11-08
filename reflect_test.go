package nborm

import (
	"fmt"
	"testing"
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

type AuthList []*Auth

func (l *AuthList) DB() string {
	return "bk_dalian"
}

func (l *AuthList) Tab() string {
	return "auth"
}

func TestReflect(t *testing.T) {
	auths := make(AuthList, 0, 8)
	tabInfo := getTabInfo(&auths)
	fmt.Println(tabInfo.columnMap["session_id"])
}
