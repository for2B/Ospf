package oauth

import (
	"fmt"

	"encoding/json"

	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/ospf/service/chencanxin/wechat/util"
)

const (
	//微信授权的accesstoken
	accessTokenURL = "https://api.weixin.qq." +
		"com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"

	refreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"

	checkAccessTokenURL = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
)

// GetUserAccessTokenFromWechatServer
func (oauth *Oauth) GetUserAccessTokenFromWechatServer() (result ResAccessToken, err error) {
	urlStr := fmt.Sprintf(accessTokenURL, oauth.AppID, oauth.AppSecret, oauth.Code)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return result, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "GetUserAccessTokenFromWechatServer() util.HTTPGet failed",
			err)
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return result, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "GetUserAccessTokenFromWechatServer() error json."+
			"Unmarshal failed", err)
	}
	if result.ErrCode != 0 {
		return result, httpapi.NewErr(constant.GLOBAL_SYS_ERR, fmt.Sprintf("GetUserAccessTokenFromWechatServer("+
			") error : errcode=%v , errmsg=%v", result.ErrCode,
			result.ErrMsg), nil)
	}
	return
}
