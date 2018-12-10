package controller_chencanxin

import (
	"net/http"
	"OutsourcingPlatform/utils"
	"time"
	"OutsourcingPlatform/internal/httpparse"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"net/url"
	"OutsourcingPlatform/ospf/model/model_chencanxin"
	"OutsourcingPlatform/ospf/infrastructure"
	"OutsourcingPlatform/ospf/service/chencanxin/qrcode"
	"strings"
)

type FecthInvitationCodeController struct {
	*model_chencanxin.FetchInvitationCodeModel
}

type Data struct {
	Qrcode_url string `json:"qrcode_url"`
}

/*
作用:返回该用户的邀请码图片url地址
输入:无
输出:返回二维码图片url地址
*/
func (c *FecthInvitationCodeController) Fetch(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	//check param
	regist_user_type := utils.QueryUrl("type", r)
	if regist_user_type != constant.REFEREE_USER_TYPE && regist_user_type != constant.NORMAL_USER_TYPE {
		return nil, httpapi.NewErr(constant.GLOBAL_PARAM_ERR, `the param can only be "NORMAL" or "REFEREE".`, nil)
	}
	var response Data = Data{Qrcode_url: ""}
	//fetch token from db firstly
	openid, err := httpparse.GetOpenid(r)
	token, err := c.GetQRcodeToken(openid, regist_user_type)
	if err != nil {
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "getqrcodetoken failed", err)
	}
	if token != "" {
		//if exist,check the token,and then if is valid,update the expire_time
		isValid, err := CheckQrcodeToken(token, c.Infrastructure, true)
		if err != nil {
			return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "qrcode.Check() failed ", err)
		}
		if isValid {
			//if valid,return image url
			sdomain := c.GetServerDomain()
			response.Qrcode_url, err = GenerateImageurlByGo_Qr_Code(sdomain, token)
			if err != nil {
				return response, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "set generate_imageurl() failed ", err)
			}
			return response, nil
		} else {
			//	delete img file
			err := DeleteQrcode(token)
			if err != nil {
				if !strings.Contains(err.Error(), "The system cannot find the file specified") {
					return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "delete qrcode fail", err)
				}

			}
		}
	}
	//or,generate new token
	response.Qrcode_url, err = c.generate_new_qrcode(r, openid)
	if err != nil {
		return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "generate_new_qrcode() failed", err)
	}
	return response, nil
}

/*
作用:生成新的邀请码token,邀请码图片url地址
输入:无
输出:返回二维码图片url地址
*/
func (c *FecthInvitationCodeController) generate_new_qrcode(r *http.Request, openid string) (string, error) {
	//generate token
	token := utils.GenerateRandomID()
	//new a qrcode
	var qrcode Qrcode = Qrcode{}
	var err error
	qrcode.Expired_time = time.Now().Add(time.Duration(c.GetOpts().QrcodeLifeTime) * time.Second)
	qrcode.Creator_openid, err = httpparse.GetOpenid(r)
	if err != nil {
		return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "getopenid() failed ", err)
	}
	regist_user_type := utils.QueryUrl("type", r)
	qrcode.Regist_user_type = regist_user_type
	//store token in cache and db
	err = c.CacheManager.Set(constant.QRCODE_BUCKEY_NAME, token, qrcode)
	if err != nil {
		return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "set qrcode token to cache () failed ", err)
	}
	err = c.InsertInvitationCodeToken(regist_user_type, openid, token)
	if err != nil {
		return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "InsertInvitationCodeToken() failed", err)
	}
	//return image url
	sdomain := c.GetServerDomain()
	imageurl, err := GenerateImageurlByGo_Qr_Code(sdomain, token)
	if err != nil {
		return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "set generate_imageurl() failed ", err)
	}
	return imageurl, nil
}

/*
作用:调用"liantu.com第三方api"生成二维码图片url地址
输入:token
输出:返回二维码图片url地址
*/
func GenerateImageurl(sdomain, token string) (string, error) {
	redirect_url := "http://" + sdomain + "/api/valid_invitation_code?token=" + token
	apiurl := "http://qr.liantu.com/api.php?w=200&text="
	responseUrl := url.QueryEscape(apiurl + redirect_url)
	return responseUrl, nil
}

/*
作用:调用类库生成二维码图片url地址
输入:token,后端服务器域名,
输出:返回二维码图片url地址
*/
func GenerateImageurlByGo_Qr_Code(sdomain, token string) (string, error) {
	redirect_url := "http://" + sdomain + "/api/valid_invitation_code?token=" + token
	qm := qrcode.NewQrcodeManager()
	path, err := qm.Create_Invitation_Qrcode(sdomain, redirect_url, token)
	if err != nil {
		return "", httpapi.NewErr(constant.GLOBAL_SYS_ERR, "qm.Create_Invitation_Qrcode() failed ", err)
	}
	responseUrl := url.QueryEscape(path)
	return responseUrl, nil
}

/*
作用:判断邀请码token的有效性,如果无效删除cache中此token和对应的文件
输入:token,是否更新有效期isUpdateExpire
输出:该token有效返回true,反之为false
*/
func CheckQrcodeToken(token string, i *infrastructure.Infrastructure, isUpdateExpire bool) (bool, error) {
	var _qrcode *Qrcode = &Qrcode{}
	err := i.CacheManager.Get(constant.QRCODE_BUCKEY_NAME, token, _qrcode)
	if err != nil {
		if err.Error() == "not found" {
			return false, nil
		}
		return false, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "CacheManager.Get() failed ", err)
	}
	//utils.ShowTime(_qrcode.Expired_time)
	//utils.ShowTime(time.Now())
	if _qrcode.Expired_time.Before(time.Now()) {
		err := i.CacheManager.Del(constant.QRCODE_BUCKEY_NAME, token)
		if err != nil {
			return false, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "CacheManager.Del() failed ", err)
		}
		DeleteQrcode(token)
		return false, nil
	}
	if isUpdateExpire {
		UpdateQrcodeExpired(_qrcode, token, i)
	}
	return true, nil
}

/*
作用:更新邀请码token的有效期
输入:token
输出:无
*/
func UpdateQrcodeExpired(qrcode *Qrcode, token string, i *infrastructure.Infrastructure) (error) {
	qrcode.Expired_time = time.Now().Add(time.Duration(i.GetOpts().QrcodeLifeTime) * time.Second)
	//store token in cache and db
	err := i.CacheManager.Set(constant.QRCODE_BUCKEY_NAME, token, qrcode)
	if err != nil {
		return httpapi.NewErr(constant.GLOBAL_SYS_ERR, "set qrcode token to cache () failed ", err)
	}
	return nil
}

/*
作用:删除指定token的邀请码图片
输入:token
输出:无
*/
func DeleteQrcode(token string) (error) {
	qm := qrcode.NewQrcodeManager()
	err := qm.Delete_Invitation_Qrcode(token)
	if err != nil {
		return err
	}
	return nil
}
