package mock_nborm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/wangjun861205/nborm"
	"github.com/wangjun861205/nborm/mock/model"

	"github.com/golang/mock/gomock"
)

func TestAccountQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockExecutor(ctrl)
	m.EXPECT().QueryRow(gomock.Any()).Return(nil).AnyTimes()
	m.EXPECT().Query(gomock.Any()).Return(nil, errors.New("mock error")).AnyTimes()
	acct := model.NewEmployAccount()
	acct.ID.ForSelect()
	if err := nborm.Query(m, acct); err != nil {
		fmt.Println(err)
	}
	accts := model.NewEmployAccountList()
	accts.ID.ForSelect()
	if err := nborm.Query(m, accts); err != nil {
		fmt.Println(err)
	}
}
