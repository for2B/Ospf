package CreateQrcode

import (
	"fmt"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/ospf/wechat/util"
)

type Qrcode struct {
	*context.Context
}

type data struct {
	Expire_seconds int64 `json:"expire_seconds"`
	Action_name string `json:"action_name"`
	Action_info info `json:"action_info"`
}

type info struct {
	Scene scene `json:"scene"`
}

type scene struct {
	Scene_str string `json:"scene_str"`
}


func (q *Qrcode) fetchTicket () error{

	token,err:=q.GetAccessToken()
	if err!=nil{
		return  httpapi.NewErr(constant.GLOBAL_SYS_ERR,"get accesstoken failed in fetchTicket()",err)
	}
	url:=fmt.Sprint("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s",token)
	var d data=data{}
	//d.Expire_seconds=constant.QRCODE_EXPIRE
	d.Action_name="QR_STR_SCENE"
	//post json to wechat server
	util.PostJSON(url,d)
	return nil
}

func (q *Qrcode) fetchQrcode (){

}