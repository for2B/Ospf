package user

import (
	"OutsourcingPlatform/internal/httpapi"
	"net/http"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/utils"
	"fmt"
	"strconv"
)

/*
用户扫码登陆的方法
*/
func (u *User) QrcodeLogin(w http.ResponseWriter, r *http.Request) error {
	if u == nil {
		return ErrOfNilUser
	}
	wc := u.Wechat
	auth := wc.GetOauth(w, r)

	// fetch redirectURL
	redirectURL := auth.Query("redirectURL")

	//fecth user's session
	session, err := u.SessionManager.FetchSessionForAuth(w, r)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "FetchSessionForAuth() failed", nil)
	}

	//if user had logined (note:accesstoken is not stored.)
	if session.Level != constant.USER_VISITOR {
		utils.ToOpenURL(w, "http://"+u.GetOpts().ClientDomain+redirectURL)
	}

	if redirectURL != "" {
		session.RedirectURL = redirectURL
		err := u.Infrastructure.SessionManager.UpdateSessionAndCookie(w, r, session)
		if err != nil {
			return httpapi.NewErr(constant.SESSION_SET_ERR, "Set redirectURL Session failed", err)
		}
	}

	//fetch code:
	auth.Context.Code = auth.Query("code")
	if auth.Context.Code == "" {
		state:=utils.GenerateRandomID()
		auth.RedirectQrcodeLoginPage(w, r, "http://"+u.GetOpts().ClientDomain+constant.CONST_QRCODE_LOGIN_URL, "snsapi_login",
			u.GetOpts().AppID, state)
		return nil
	}

	// fetch AccessToken,openid
	at, err := auth.GetUserAccessTokenFromWechatServer()
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "auth.GetUserAccessToken execuse failed", err)
	}
	fmt.Println("openid=", at.Openid)

	// 获取用户详细信息
	rt, err := auth.GetUserDetailInfo(at.AccessToken, at.Openid)
	if rt.CommonError.ErrCode != 0 {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "wechat auth: fecth user info failed:"+rt.
			ErrMsg+strconv.Itoa(int(
			rt.ErrCode)), nil)
	}
	fmt.Println("user struct=", rt)
	fmt.Println("user Nickname=", rt.Nickname)
	fmt.Println("user Headimgurl=", rt.Headimgurl)

	//fecth user detail more from db
	//session.level=

	// update session
	session.Openid = at.Openid
	err = u.Infrastructure.SessionManager.UpdateSessionAndCookie(w, r, session)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "Set redirectURL Session failed", err)
	}

	// redirect to redirectURL
	fmt.Println("redirectURL=", session.RedirectURL)
	utils.ToOpenURL(w, "http://"+u.GetOpts().ClientDomain+session.RedirectURL)

	return nil
}