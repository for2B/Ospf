package oauth

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	fetchCodeFromWechatServer        = "https://open.weixin.qq." +
		"com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s" +
		"#wechat_redirect";
	fetchAccessTokenFromWechatServer = "";
)

func buildFetchCodeFromWechatServerUrl(appid , redirecturi,state string) string{
	return fmt.Sprintf(fetchCodeFromWechatServer,appid,redirecturi,state)

}

/*
作用:跳转到扫码登陆页面
*/
func (oauth *Oauth) RedirectQrcodeLoginPage(writer http.ResponseWriter, req *http.Request, redirectURI, scope, agentid,
	state string) error {
	location, err := oauth.GetRedirectURLQrcodeLogin(redirectURI, scope, agentid, state)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, 302)
	return nil
}

//GetRedirectURL 获取跳转的url地址
func (oauth *Oauth) GetRedirectURLQrcodeLogin(redirectURI, scope, agentid, state string) (string, error) {
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(fetchCodeFromWechatServer, oauth.AppID, urlStr, scope,state), nil
}

