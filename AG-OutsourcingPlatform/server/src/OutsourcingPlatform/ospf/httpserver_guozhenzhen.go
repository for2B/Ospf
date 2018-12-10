package ospf

import (
	"github.com/gorilla/mux"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/middleware"
	"OutsourcingPlatform/ospf/controller/controller_guozhenzhen"
	"OutsourcingPlatform/ospf/model/model_guozhenzhen"
	"OutsourcingPlatform/ospf/constant"
)

func (s *httpServer) initRouter_guozhenzhen() {
       s.Login(s.router)
       s.SetBaseInfo(s.router)

}

//登录
func (s *httpServer) Login(r *mux.Router){
	c := &controller_guozhenzhen.Logincontroller{
		UserModel:model_guozhenzhen.GetUserModel(s.ctx.ospf.infrastructure),
	}
	r.Handle(constant.CONST_LOGIN_URL,httpapi.Decorate(c.Login,s.defaultMiddleWare))

}

//填写或修改基础信息
func (s *httpServer) SetBaseInfo(r *mux.Router){
	c := &controller_guozhenzhen.SetBaseInfocontroller{
		SetBaseInfoModel:model_guozhenzhen.SetBaseInformationModel(s.ctx.ospf.infrastructure),
	}
	r.Handle(constant.CONST_SET_BASE_INFORMATION,httpapi.Decorate(c.SetBaseInfo,s.AllowOrigin,middleware.DefaultDecode,
		middleware.Online,middleware.Log(s.ctx.ospf.infrastructure.Logger)))
}

//判断用户状态
func (s *httpServer) GetUserState(r *mux.Router){

}