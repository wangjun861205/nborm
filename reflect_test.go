package nborm

import (
	"fmt"
	"testing"
)

type Auth struct {
	Id            *IntField      `column:"id" auto_increment:"true"`
	Username      *StringField   `column:"username"`
	Password      *StringField   `column:"password"`
	Phone         *StringField   `column:"phone" primary_key:"true"`
	Status        *IntField      `column:"status"`
	SessionId     *StringField   `column:"session_id"`
	ExpireTime    *DatetimeField `column:"expire_time"`
	Email         *StringField   `column:"email" unique:"true"`
	CreateTime    *DatetimeField `column:"create_time"`
	LastLoginTime *DatetimeField `column:"last_login_time"`
	IsSuperUser   *BoolField     `column:"is_super_user"`
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
	InitSlice(&auths)
	fmt.Println(len(auths))
}
