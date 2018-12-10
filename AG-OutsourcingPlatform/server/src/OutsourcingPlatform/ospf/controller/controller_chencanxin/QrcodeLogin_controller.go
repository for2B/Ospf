package controller_chencanxin

import (
	"net/http"
	"OutsourcingPlatform/ospf/infrastructure"
	"OutsourcingPlatform/ospf/service/chencanxin/user"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/internal/httpapi"
)

type QrcodeLoginController struct {
	*infrastructure.Infrastructure
}

func (c *QrcodeLoginController) QrcodeLogin(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	u:=user.NewUser(c.Infrastructure)

	err:=u.QrcodeLogin(w,r)
	if err!=nil{
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "user oauth fail",err)
	}
	return nil,nil
}
