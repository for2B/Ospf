package controller_chencanxin


import (
	"net/http"
	"OutsourcingPlatform/ospf/service/chencanxin/user"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/ospf/model/model_chencanxin"
)

type WcAuthController struct {
	*model_chencanxin.WcAuthModel
}

func (c *WcAuthController) WechatAuth(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	u:=user.NewUser(c.Infrastructure)
	err:=u.Oauth(w,r)
	if err!=nil{
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "user oauth fail",err)
	}
	return nil,nil
}
