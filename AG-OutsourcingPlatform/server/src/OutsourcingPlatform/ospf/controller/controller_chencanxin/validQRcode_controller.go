package controller_chencanxin

import (
	"net/http"
	"OutsourcingPlatform/ospf/infrastructure"
	"OutsourcingPlatform/utils"
	"time"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/utils/htmltool"
)

type ValidQRCodeController struct {
	*infrastructure.Infrastructure
}
type Qrcode struct {
	Creator_openid string
	Expired_time   time.Time
	Regist_user_type string
}





/*
作用:验证用户当前识别的二维码的有效性
输入:二维码token
输出:若有效跳转至完善信息页面;否则返回失败信息
*/
func (c *ValidQRCodeController) Valid(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	//check the token param
	token := utils.QueryUrl("token", r)
	isValid,err:=CheckQrcodeToken(token,c.Infrastructure,false)
	if err!=nil{
		return nil,httpapi.NewErr(constant.GLOBAL_SYS_ERR,"CheckQrcodeToken() failed",err)
	}
	if !isValid{
		htmltool.ResponHtml(w, "抱歉,该邀请码已失效", `
				<h1>
				抱歉,该邀请码已失效.
				</h1>
								`)
		return nil,httpapi.NewErr(constant.QRCODE_TOKEN_INVALID,"qrcode token is invalid",nil)
	}
	//if token is valid,redirect to regist page
	cdomain:=c.GetOpts().ClientDomain
	utils.ToOpenURL(w, "http://"+cdomain+"/user/regist")
	return nil,nil
}
