package ProjectManage

import (
	"OutsourcingPlatform/ospf/model/projectRelatedPersonModel"
	"net/http"
	"OutsourcingPlatform/internal/httpparse"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/internal/httpapi"
)

type ProjectRelatedPersonController struct {
	*projectRelatedPersonModel.ProjectRelatedPersonModel
}

//取消项目承包者
func (c *ProjectRelatedPersonController)HandleCancelContractor(w http.ResponseWriter,r *http.Request)(interface{},error){
	if r.Method == "GET"{

		values := r.URL.Query()
		ProjectId := values.Get("project_id")

		err:=c.Cancelcontractor(ProjectId)
		if err!=nil{
			return nil ,err
		}
		return nil,nil
	}
	return nil, nil
}

//获取项目竞标列表
func(c *ProjectRelatedPersonController)HandleGetBiddingList(w http.ResponseWriter,r *http.Request)(interface{},error){
	if r.Method == "GET" {

		//get parameter
		value := r.URL.Query()
		ProjectId := value.Get("project_id")
		CurrentPage := value.Get("current_page")
		ShowQuantity:=value.Get("show_nums")
		SortOrder:=value.Get("sort_order")

		//get BiddingList
		BiddingList,err:=c.GetBiddingList(ProjectId,CurrentPage,ShowQuantity,SortOrder)
		if err!=nil {
			return nil,err
		}
		return BiddingList,err
	}
	return nil,nil
}

//获取用户相对于某个项目的身份
func (c *ProjectRelatedPersonController) HandleGetProjectIdentity(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method == "GET" {

		value := r.URL.Query()
		ProjectId := value.Get("project_id")
		session, err :=  httpparse.GetSession(r)
		if err!=nil {
			return nil,httpapi.NewErr(constant.SESSION_FETCH_ERR,"get session fail",err)
		}
		//session := constant.Session{Openid: "11113"}

		return c.GetProjectIdentity(ProjectId, session.Openid)
	}
	return nil, nil
}

//获取项目团队成员
func (c *ProjectRelatedPersonController)HandleGetTeammate(w http.ResponseWriter,r *http.Request) (interface{},error){
	if r.Method == "GET" {
		value := r.URL.Query()
		ProjectId := value.Get("project_id")

		session, err :=  httpparse.GetSession(r)
		if err!=nil {
			return nil,httpapi.NewErr(constant.SESSION_FETCH_ERR,"get session fail",err)
		}

		//session := constant.Session{Openid:"6334824724549167320"}
		return c.GetTeammate(ProjectId,session.Openid)
	}
	return nil,httpapi.NewErr(constant.REQUEST_METHOd_FAIL,"Request method error , it should be POST",nil)
}

//选择项目承包者
func (c *ProjectRelatedPersonController)HandleSelectProjectContractor(w http.ResponseWriter,r *http.Request)(interface{},error){
	if r.Method == "GET"{

		value := r.URL.Query()
		UserId := value.Get("user_id")
		ProjectId := value.Get("project_id")

		err := c.SelectProjectContractor(ProjectId,UserId)
		if err!=nil{
			return nil,err
		}
		return nil,nil
	}
	return nil,nil
}
