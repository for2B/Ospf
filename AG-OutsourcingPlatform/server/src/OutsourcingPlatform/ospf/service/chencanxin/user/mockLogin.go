package user

import (
	"net/http"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/utils"
	"fmt"
)

/*
用户模拟登陆的方法
*/
func (u *User) MockOauth(w http.ResponseWriter, r *http.Request) error {
	// 接受参数：模拟登录者的用户类型usertype，redirectURL,设置session

	//fecth user's session
	ses, err := u.SessionManager.FetchSessionForAuth(w, r)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "FetchSessionForAuth() failed in mockoauth()", nil)
	}
	//redirectURL := utils.QueryUrl("redirectURL", r)
	usertype := utils.QueryUrl("userType", r)
	switch usertype {
	case "0":
		//普通用户
		ses.Level = constant.USER_NORMAL
		ses.Name = "普通用户"
		ses.Openid="mockopenid0"

	case "1":
		//证明人
		ses.Level = constant.USER_REFEREE
		ses.Name = "证明人"
		ses.Openid="mockopenid1"

	case "2":
		//管理员
		ses.Level = constant.USER_ADMIN
		ses.Name = "管理员"
		ses.Openid="mockopenid2"

	case "3":
		//上帝
		ses.Level = constant.USER_SUPER
		ses.Name = "上帝"
		ses.Openid="mockopenid3"
	}

	ses.RedirectURL = utils.QueryUrl("redirectURL", r)
	fmt.Println("session=",ses.RedirectURL)
	err = u.Infrastructure.SessionManager.UpdateSession(ses)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "UpdateSessionAndCookie Session failed in mockoauth", err)
	}

	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "Set virtual Session failed", nil)
	}

	// 跳转至redirectURL
	utils.ToOpenURL(w, "http://"+u.GetOpts().ClientDomain+ses.RedirectURL)
	return nil
}
