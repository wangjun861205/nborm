package nborm

import (
	"fmt"
	"testing"
)

func TestSorter(t *testing.T) {
	auths := NewAuthList()
	auth1, auth2, auth3 := NewAuth(), NewAuth(), NewAuth()
	auth1.Username.Set("wangjun", false)
	auth2.Username.Set("tongyao", false)
	auth3.Username.Set("buxiong", false)
	*auths = append(*auths, auth1, auth2, auth3)
	Sort(auths, false, (*auth)[0].Username)
	fmt.Println((*auths)[1].Username.Get())
	fmt.Println((*auths)[2].Username.Get())
	fmt.Println((*auths)[3].Username.Get())

}
