package ProjectManage

import (
	"net/http"
	"encoding/json"
	"OutsourcingPlatform/internal/httpparse"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/model/getProjectModel"
)

type GetProjectController struct {
	*getProjectModel.GetProjectModel
}

type ProjectSiftParameter struct {
	CycleRangeFirst int `json:"cycle_range_start"`
	CycleRangeLast  int `json:"cycle_range_end"`
	MoneyRangeFirst int `json:"money_range_start"`
	MoneyRangeLast  int `json:"money_range_end"`
	DateRangeFirst  int `json:"date_range_start"`
	DateRangeLast   int `json:"date_range_end"`
	CurrentPage     int `json:"current_page"`
	ShowQuantity    int `json:"show_nums"`
}

//获取平台所有项目列表
func (c *GetProjectController) HandleGetPlfLsProject(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	if r.Method == "POST" {

		body := r.Context()
		BodyData, ok := body.Value("Body").(string)
		if !ok {
			return nil, httpapi.Err{Code: constant.HTTP_ERR_READBODY, Msg: "get body param failed"}
		}

		var SiftParameter ProjectSiftParameter

		err := json.Unmarshal([]byte(BodyData), &SiftParameter)
		if err != nil {
			return nil, httpapi.NewErr(constant.HTTP_ERR_READBODY, "UnmJson failed", err)
		}

		session, err :=  httpparse.GetSession(r)
		if err!=nil {
			return nil,httpapi.NewErr(constant.SESSION_FETCH_ERR,"get session fail",err)
		}

		return c.GetPlatformProjectList(
			SiftParameter.CycleRangeFirst, SiftParameter.CycleRangeLast,
			SiftParameter.MoneyRangeFirst, SiftParameter.MoneyRangeLast,
			SiftParameter.DateRangeFirst, SiftParameter.DateRangeLast,
			SiftParameter.CurrentPage, SiftParameter.ShowQuantity,
			c.DB, c.Logger,session.Openid)
	}
	return nil, nil
}

//获取我承包的项目列表
func (c *GetProjectController)HandleGetProjectContract(w http.ResponseWriter,r *http.Request) (interface{},error){

	if r.Method == "GET"{

		value := r.URL.Query()
		ProjectStatus := value.Get("project_status")
		CurrentPag:=value.Get("current_page")
		ShowQuantity := value.Get("show_nums")

		session, err :=  httpparse.GetSession(r)
		if err!=nil {
			return nil,httpapi.NewErr(constant.SESSION_FETCH_ERR,"get session fail",err)
		}

		//session := constant.Session{Openid:"5577006791947779410"}
		return c.GetProjectContract(ProjectStatus,session.Openid,CurrentPag,ShowQuantity)

	}

	return nil,httpapi.NewErr(constant.REQUEST_METHOd_FAIL,"request method fail",nil)
}

//获取我发布的项目列表
func (c *GetProjectController) HandleGetProjectPublish(w http.ResponseWriter,r *http.Request) (interface{},error){
	if r.Method == "GET"{

		value := r.URL.Query()
		ProjectStatus := value.Get("project_status")
		CurrentPage:=value.Get("current_page")
		ShowQuantity := value.Get("show_nums")
		session, err :=  httpparse.GetSession(r)
		if err!=nil {
			return nil,httpapi.NewErr(constant.SESSION_FETCH_ERR,"get session fail",err)
		}

		//session := constant.Session{Openid:"6334824724549167320"}
		return c.GetProjectPublish(ProjectStatus,session.Openid,CurrentPage,ShowQuantity)
	}
	return nil,httpapi.NewErr(constant.REQUEST_METHOd_FAIL,"request method fail",nil)
}

//获取对应身份下各状态的项目总数
func (c *GetProjectController)HandleCountCompartedStatus(w http.ResponseWriter,r *http.Request)(interface{},error){
	if r.Method == "GET"{

		values := r.URL.Query()
		ListType := values.Get("list_type")
		session, err :=  httpparse.GetSession(r)
		if err!=nil {
			return nil,httpapi.NewErr(constant.SESSION_FETCH_ERR,"get session fail",err)
		}
		//session := constant.Session{Openid:"5577006791947779410"}

		return c.CountCompartedStatus(ListType,session.Openid)
	}
	return nil, nil
}
