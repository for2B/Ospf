package ospf

import (
	"github.com/gorilla/mux"
	"OutsourcingPlatform/internal/httpapi"

	"OutsourcingPlatform/ospf/controller/controller_guochuting"
	"OutsourcingPlatform/ospf/model/model_guochuting"
	// "OutsourcingPlatform/ospf/constant"
)

func (s *httpServer) initRouter_guochuting() {
	s.GetEvaluationController(s.router)
	s.PostEvaluateController(s.router)
}

func (s *httpServer) GetEvaluationController(r *mux.Router) {
	c := &controller_guochuting.GetEvaluationController{
		GetEvaluationModel:model_guochuting.EvaluationModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/get_evaluation_to_other", httpapi.Decorate(c.GetEvaluationController, s.defaultMiddleWare))

	//r.Handle("/api/get_evaluation_to_other", httpapi.Decorate(c.GetEvaluationController,middleware.Online,middleware.DefaultDecode,middleware.Log(s.ctx.ospf.infrastructure.Logger)))
}

func (s *httpServer) PostEvaluateController(r *mux.Router) {
	c := &controller_guochuting.PostEvaluateController{
		PostEvaluateModel:model_guochuting.EvaluateModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/write_evaluate", httpapi.Decorate(c.PostEvaluateController, s.defaultMiddleWare))
}