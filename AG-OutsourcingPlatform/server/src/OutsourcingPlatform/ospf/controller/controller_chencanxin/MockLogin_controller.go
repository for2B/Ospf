package controller_chencanxin

import (
	"net/http"
	"OutsourcingPlatform/ospf/service/chencanxin/user"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/ospf/infrastructure"
	"OutsourcingPlatform/internal/httpapi"
)

type MockLoginController struct {
	*infrastructure.Infrastructure
}

func (c *MockLoginController) MockLogin(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	u:=user.NewUser(c.Infrastructure)

	err:= u.MockOauth(w,r)
	if err!=nil{
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "user oauth fail",err)
	}
	return nil,nil
}
