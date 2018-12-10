package ospf

import (
	"github.com/gorilla/mux"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/controller/controller_panjiawei"
	"OutsourcingPlatform/ospf/model/model_panjiawei"
)

func (s *httpServer) initRouter_panjiawei() {
	s.teamMembersController(s.router)
	s.userManagement(s.router)
}

func (s *httpServer) teamMembersController(r *mux.Router) {
	c := &controller_panjiawei.GetTeamMemberController{
		model_panjiawei.New_GetMemberModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/getmember", httpapi.Decorate(c.GetMember,s.AllowOrigin, s.defaultMiddleWare))
}

func (s *httpServer) userManagement(r *mux.Router) {
	c := &controller_panjiawei.UserManagerController{
		model_panjiawei.NewUserManagerModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/alluser", httpapi.Decorate(c.GetUsers, s.defaultMiddleWare,s.AllowOrigin))
	r.Handle("/api/getuser", httpapi.Decorate(c.GetOneUser,s.AllowOrigin, s.defaultMiddleWare))
	r.Handle("/api/alter_permission", httpapi.Decorate(c.ChangeUserPermissions,s.AllowOrigin, s.defaultMiddleWare))
	r.Handle("/api/del_users",httpapi.Decorate(c.DeleteUsers,s.AllowOrigin, s.defaultMiddleWare))
}

