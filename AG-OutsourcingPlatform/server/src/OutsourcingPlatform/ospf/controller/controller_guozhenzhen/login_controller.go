package controller_guozhenzhen

import (
	"OutsourcingPlatform/ospf/model/model_guozhenzhen"
	"net/http"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"encoding/json"
)

type Logincontroller struct {
	*model_guozhenzhen.UserModel
}


type Login struct {
	Phone    string
	Password string
}

type Back struct {
	Msg  string
}

//登录
//1.获取参数
//2.校验参数
//3.根据账号查询记录
//-如果没有记录，返回用户不存在
//4.匹配密码
//--如果匹配成功，返回成功
//--如果匹配失败，返回失败
func (c *Logincontroller) Login(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var back Back
	var err error
	if r.Method == "POST" {
		var loginUser Login
		var user model_guozhenzhen.UserModel
		login, ok := r.Context().Value("Body").(string)
		if ok == false {
			return nil, httpapi.NewErr(constant.GLOBAL_ASSERT_ERR, "assert is failed", nil)
		}
		err = json.Unmarshal([]byte(login), &loginUser)
		if err != nil {
			return nil, httpapi.NewErr(constant.UNMARSHAL_FAIL, "unmarshal is failed", nil)
		}
		if loginUser.Phone == "" {
			back.Msg = "phone is nil"
			return nil, httpapi.NewErr(constant.USERNAME_NIL, "username is nil", nil)
		} else {
			user, err = c.Getuser(loginUser.Phone)
			if err != nil {
				return nil, httpapi.NewErr(constant.GLOBAL_DB_ERR, "db error", err)
			}
			if user.Password == loginUser.Password && user.Password != "" {
				back.Msg = "login success"
				return back.Msg, nil
			} else if user.Password == "" {
				back.Msg = "not have the person"
				return nil, httpapi.NewErr(constant.USERNAME_NOT_EXIST, "user is not exist", nil)
			} else {
				back.Msg = "password wrong"
				return nil, httpapi.NewErr(constant.PASSWORD_WRONG, "password wrong", nil)
			}
		}
		back.Msg = "unknow error"
		return nil, httpapi.NewErr(constant.UNKNOW_ERROR,"unknow error",nil)
	}

	return nil,nil
}
