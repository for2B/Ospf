package controller_panjiawei

import (
	"OutsourcingPlatform/ospf/model/model_panjiawei"
	"net/http"
	"strconv"
	"OutsourcingPlatform/internal/httpapi"
)

type GetTeamMemberController struct{
	*model_panjiawei.GetMemberModel
}

type UrlParams struct {
	project_id string
}
func (m *GetTeamMemberController) GetMember(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method != "GET"{
		return "",httpapi.NewErr(505,"【func GetMember】 method is not 'GET'" ,nil)
	}
	user_id ,err:= strconv.ParseInt(r.URL.Query().Get("uid"), 10, 64)
	if err!= nil {
		return "",httpapi.NewErr(502,"user_id URL参数错误",err)
	}
	response,err := m.GetMemberInDB(user_id)
	return response,err
}
