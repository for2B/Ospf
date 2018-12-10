package controller_guozhenzhen

import (
	"net/http"
	"OutsourcingPlatform/ospf/model/model_guozhenzhen"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"encoding/json"

	"log"
)

type SetBaseInfocontroller struct {
	*model_guozhenzhen.SetBaseInfoModel
}

type Info struct {
	Username string       `json:"Username"`
	Sex      string       `json:"Sex"`
	Password string       `json:"Password"`
	Wechat_account string `json:"Wechat_account"`
	Phone string          `json:"Phone"`
	Identity string       `json:"Identity"`
	Credit string         `json:"Credit"`
	Location string       `json:"Location"`
}




func (c *SetBaseInfocontroller) SetBaseInfo(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var back Back
	var err error
	if r.Method == "POST" {
		//var setbaseinfo model_guozhenzhen.SetBaseInfoModel
		var setbase model_guozhenzhen.SetBaseInfoModel
		body,ok:=r.Context().Value("Body").(string)
		if ok == false {
			return nil,httpapi.NewErr(constant.GLOBAL_ASSERT_ERR,"assert is failed",nil)
		}
		log.Print(body)
		err = json.Unmarshal([]byte(body), &setbase)
		log.Println(setbase)
		if err != nil {
			return nil, httpapi.NewErr(constant.UNMARSHAL_FAIL,"unmarshal is failed",err)
		}
		_,err=c.SetBaseInfor(setbase)
		if err != nil {
			return nil, httpapi.NewErr(constant.UNMARSHAL_FAIL,"",err)
		}
		back.Msg = "unknow error"
		bac, erro := json.Marshal(back)
		if erro != nil {
			return nil, httpapi.NewErr(constant.MARSHAL_FAIL,"marshal is failed",nil)
		}
		w.Write(bac)
		//return nil, httpapi.NewErr(constant.UNKNOW_ERROR,"unknow error",nil)
	}
	return nil,nil

}
