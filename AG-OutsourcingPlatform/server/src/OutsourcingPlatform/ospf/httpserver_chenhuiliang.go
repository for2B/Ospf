package ospf

import (
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/middleware"
	"OutsourcingPlatform/ospf/controller/ProjectManage"
	"OutsourcingPlatform/utils"
	"net/http"

	"github.com/gorilla/mux"
	"OutsourcingPlatform/ospf/model/getProjectModel"
	"OutsourcingPlatform/ospf/model/projectOperateModel"
	"OutsourcingPlatform/ospf/model/projectRelatedPersonModel"
	"OutsourcingPlatform/ospf/controller/EditorFile"
	"OutsourcingPlatform/ospf/model/editorFileModel"
)

func (s *httpServer) initRouter_chenhuiliang() {
	s.RouteUploadFile(s.router)          //上传文件
	s.RouteDownFile(s.router)            //下载文件接口
	s.RouteGetProject(s.router) //获取项目管理
	s.RouteProjectOperate(s.router) //项目操作管理
	s.RouteProjectRelatedPerson(s.router) //项目相关人员管理


}
//-----------富文本处理---------------
//处理文件上传并保存
func (s *httpServer) RouteUploadFile(r *mux.Router) {
	path, _ := utils.GetProDir()
	c := &EditorFile.UpLoadFileController{
		UploadFileModel: EditorFileModel.GetUploadFileModel(s.ctx.ospf.infrastructure, path),
	}
	r.Handle("/api/upload_file", httpapi.Decorate(c.HandleUploadFile, s.AllowOrigin, middleware.Online, middleware.DefaultDecode, middleware.Log(s.ctx.ospf.infrastructure.Logger)))
}

//下载富文本文件接口
func (s *httpServer) RouteDownFile(r *mux.Router) {
	path, _ := utils.GetProDir()
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(path+"/project_item_src/"))))
}
//-----------富文本处理---------------


//----------获取项目管理------------------------------------------------start
func (s *httpServer) RouteGetProject( r *mux.Router){
	c := &ProjectManage.GetProjectController{
		GetProjectModel: getProjectModel.ProduceGetProjectModel(s.ctx.ospf.infrastructure),
	}
	//获取平台所有项目
	r.Handle("/api/get_project_list", httpapi.Decorate(c.HandleGetPlfLsProject, s.defaultMiddleWare))
	//获取我承包的项目列表
	r.Handle("/api/get_project_contract", httpapi.Decorate(c.HandleGetProjectContract, s.defaultMiddleWare))
	//获取我发布的项目列表
	r.Handle("/api/get_project_publish", httpapi.Decorate(c.HandleGetProjectPublish, s.defaultMiddleWare))
	//获取对应身份下各状态的项目总数
	r.Handle("/api/count_comparted_status", httpapi.Decorate(c.HandleCountCompartedStatus, s.defaultMiddleWare))
}
//----------获取项目管理------------------------------------------------end



//------------项目操作管理--------------------------------------------start
func (s *httpServer)RouteProjectOperate(r *mux.Router){
	c := &ProjectManage.ProjectOperateController{
		ProjectOperateModel:projectOperateModel.ProduceProjectOperateModel(s.ctx.ospf.infrastructure),
	}
	//发布项目
	r.Handle("/api/publish_project", httpapi.Decorate(c.HandlePublicProject, s.defaultMiddleWare))
	//修改项目
	r.Handle("/api/modification_project", httpapi.Decorate(c.HandleModificationProject, s.defaultMiddleWare))
	//更新项目状态
	r.Handle("/api/update_project_status", httpapi.Decorate(c.HandleUpdateProjectStatus, s.defaultMiddleWare))
	//参与项目竞标
	r.Handle("/api/join_bidding", httpapi.Decorate(c.HandleJoinBidding, s.defaultMiddleWare))
	//放弃项目竞标
	r.Handle("/api/give_up_bidding", httpapi.Decorate(c.HandleGiveUpBidding, s.defaultMiddleWare))
}
//------------项目操作管理--------------------------------------------end

//------项目相关人员管理----------------------------------------------start
func (s *httpServer)RouteProjectRelatedPerson(r *mux.Router){
	c := &ProjectManage.ProjectRelatedPersonController{
		ProjectRelatedPersonModel:projectRelatedPersonModel.ProduceProjectRelatedPersonModel(s.ctx.ospf.infrastructure),
	}
	//选择项目承包人接口
	r.Handle("/api/select_project_contractor", httpapi.Decorate(c.HandleSelectProjectContractor, s.defaultMiddleWare))
	//取消项目承包者
	r.Handle("/api/give_up_contract", httpapi.Decorate(c.HandleCancelContractor, s.defaultMiddleWare))
	//获取竞标列表
	r.Handle("/api/get_bidding_list", httpapi.Decorate(c.HandleGetBiddingList, s.defaultMiddleWare))
	//获取用户对某个项目的身份
	r.Handle("/api/get_project_identity", httpapi.Decorate(c.HandleGetProjectIdentity, s.defaultMiddleWare))
	//获取团队成员
	r.Handle("/api/get_project_teammate",httpapi.Decorate(c.HandleGetTeammate,s.defaultMiddleWare))
}
//------项目相关人员管理----------------------------------------------end
