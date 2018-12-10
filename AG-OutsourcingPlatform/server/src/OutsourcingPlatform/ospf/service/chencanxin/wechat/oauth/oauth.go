package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"OutsourcingPlatform/ospf/service/chencanxin/wechat/util"
	"OutsourcingPlatform/ospf/service/chencanxin/wechat/context"
)

const (
	redirectOauthURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&agentid=%s&state=%s#wechat_redirect"

	getUserDetailInfoURL = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
)

//Oauth 保存用户授权信息
type Oauth struct {
	*context.Context
}

//NewOauth 实例化授权信息
func NewOauth(context *context.Context, w http.ResponseWriter, r *http.Request) *Oauth {
	auth := new(Oauth)
	auth.Context = context
	auth.Context.Request = r
	auth.Context.Writer = w
	return auth
}

//GetRedirectURL 获取跳转的url地址
func (oauth *Oauth) GetRedirectURL(redirectURI, scope, agentid, state string) (string, error) {
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(redirectOauthURL, oauth.AppID, urlStr, scope, agentid, state), nil
}

//Redirect 跳转到网页授权
func (oauth *Oauth) Redirect(writer http.ResponseWriter, req *http.Request, redirectURI, scope, agentid, state string) error {
	location, err := oauth.GetRedirectURL(redirectURI, scope, agentid, state)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, 302)
	return nil
}

// ResAccessToken 获取用户授权access_token的返回结果
type ResAccessToken struct {
	util.CommonError
	ReAccessToken string `json:"refresh_token"`
	Openid        string `json:"openid"`
	AccessToken   string `json:"access_token"`
	ExpiresIn     int64  `json:"expires_in"`
}

// ResAccessToken 获取用户授权access_token的返回结果

type UserInfo struct {
	util.CommonError
	OpenId     string `json:"openid"`
	Nickname   string `json:"nickname"`
	Sex        string `json:"sex"`
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"`
}

// GetUserDetailInfo   https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
func (oauth *Oauth) GetUserDetailInfo(acceseeToken, openid string) (result UserInfo, err error) {
	urlStr := fmt.Sprintf(getUserDetailInfoURL, acceseeToken, openid)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserDetailInfo error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
