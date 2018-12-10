package ProjectManage

import (
	"OutsourcingPlatform/ospf/model/projectOperateModel"
	"net/http"
	"encoding/json"
	"OutsourcingPlatform/internal/httpparse"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

type ProjectOperateController struct {
	*projectOperateModel.ProjectOperateModel
}

type ProjectBaseInfo struct {
	ProjectId int           `json:"project_id"`
	ProjectName string 		`json:"project_name"`
	ProjectCircle int		`json:"project_circle"`
	ProjectMoney float64	`json:"project_money"`
	ProjectDetail string	`json:"project_detail"`
}

//发布项目
func (c *ProjectOperateController) HandlePublicProject(w http.ResponseWriter, r *http.Request) (interface{},error){
	if r.Method == "POST"{

		BodyData ,err := httpparse.GetBody(r)
		if err!=nil  {
			return nil , httpapi.Err{Code:constant.HTTP_ERR_READBODY,Msg:"get body param failed"}
		}

		var projectBaseInfo ProjectBaseInfo
		err = json.Unmarshal([]byte(BodyData),&projectBaseInfo)
		if err!=nil{
			return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "UnmJson failed", err)
		}

		session, err :=  httpparse.GetSession(r)
		if err!=nil {
			return nil,httpapi.NewErr(constant.SESSION_FETCH_ERR,"get session fail",err)
		}

		//session := constant.Session{Openid:"11111"}
		err = c.PublishProject(projectBaseInfo.ProjectName,projectBaseInfo.ProjectDetail,projectBaseInfo.ProjectCircle,projectBaseInfo.ProjectMoney,session.Openid)
		if err!= nil {
			return nil ,err
		}
	}
	return nil,nil
}

//修改项目
func(c *ProjectOperateController) HandleModificationProject(w http.ResponseWriter,r *http.Request) (interface{},error){
	if r.Method == "POST"{

		BodyData ,err := httpparse.GetBody(r)
		if err!=nil  {
			return nil , httpapi.Err{Code:constant.HTTP_ERR_READBODY,Msg:"get body param failed"}
		}

		var projectBaseInfo ProjectBaseInfo
		err = json.Unmarshal([]byte(BodyData),&projectBaseInfo)
		if err!=nil{
			return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "UnmJson failed", err)
		}

		err = c.ModifyProject(projectBaseInfo.ProjectName,projectBaseInfo.ProjectDetail,projectBaseInfo.ProjectId,projectBaseInfo.ProjectCircle,projectBaseInfo.ProjectMoney)
		if err!= nil {
			return nil ,err
		}
	}
	return nil,nil
}

//更新项目状态
func (c *ProjectOperateController)HandleUpdateProjectStatus(w http.ResponseWriter,r *http.Request) (interface{},error){
	if  r.Method == "GET"{

		value := r.URL.Query()

		ProjectStatus := value.Get("project_status")
		ProjectId := value.Get("project_id")

		err:=c.UpdateProjectStatus(ProjectId,ProjectStatus)
		if err!=nil{
			return nil,err
		}

		return nil,nil
	}
	return nil,httpapi.NewErr(constant.REQUEST_METHOd_FAIL,"request method fail",nil)
}

//参与项目竞标
func (c *ProjectOperateController) HandleJoinBidding(w http.ResponseWriter, r *http.Request) (interface{},error){
	if r.Method == "GET"{

		value := r.URL.Query()
		ProjectId := value.Get("project_id")

		session, err :=  httpparse.GetSession(r)
		if err!=nil {
			return nil,httpapi.NewErr(constant.SESSION_FETCH_ERR,"get session fail",err)
		}
		//session := constant.Session{Openid:"6334824724549167320"}

		err = c.JoinBidding(ProjectId,session.Openid)
		if err!=nil{
			return nil,err
		}
		return nil,nil
	}
	return nil,httpapi.NewErr(constant.REQUEST_METHOd_FAIL,"request method fail",nil)
}

//放弃项目竞标
func (c *ProjectOperateController) HandleGiveUpBidding(w http.ResponseWriter, r *http.Request) (interface{},error){
	if r.Method == "GET"{

		value := r.URL.Query()
		ProjectId := value.Get("project_id")

		session, err :=  httpparse.GetSession(r)
		if err!=nil {
			return nil,httpapi.NewErr(constant.SESSION_FETCH_ERR,"get session fail",err)
		}
		//session := constant.Session{Openid:"11114"}

		err = c.GiveUpBidding(ProjectId,session.Openid)
		if err!=nil{
			return nil,err
		}
		return nil,nil
	}
	return nil,httpapi.NewErr(constant.REQUEST_METHOd_FAIL,"request method fail",nil)
}
