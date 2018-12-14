package ospf

import (
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/ospf/middleware"
	"OutsourcingPlatform/ospf/view"
	"net/http"

	"OutsourcingPlatform/ospf/controller/controller_chencanxin"
	"OutsourcingPlatform/ospf/model/model_chencanxin"

	"github.com/gorilla/mux"
)

type context struct {
	ospf *OSPF
}

// 拦截非前端路由
var frontEndRoutes = []string{
	"/admin/mainshow",
	"/admin/adminmanage",
}

type httpServer struct {
	ctx               *context
	router            *mux.Router
	defaultMiddleWare httpapi.MiddlewareFunc
}

func newHTTPServer(ctx *context) *httpServer {

	router := mux.NewRouter()
	s := &httpServer{
		ctx:    ctx,
		router: router,
		defaultMiddleWare: httpapi.CombinationMiddleware(ctx.ospf.infrastructure.AllowOrigin,
			middleware.Online,
			middleware.DefaultDecode,
			middleware.Encode,
			middleware.Log(ctx.ospf.infrastructure.Logger),
			middleware.Session(ctx.ospf.infrastructure.SessionManager),
			middleware.Permission(constant.USER_VISITOR)),
	}
	s.initRouter()
	initFrontEndMux(router)

	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		view.GzipFileServe(w, r, r.URL.Path, ctx.ospf.infrastructure.GetOpts().FrontDir)
	})))
	router.PathPrefix("/assets/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		view.GzipFileServe(w, r, r.URL.Path, ctx.ospf.infrastructure.GetOpts().FrontDir)
	})
	//router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	view.GzipServeFile(w, r, ctx.ospf.infrastructure.GetOpts().FrontDir+"index.html")
	//})
	router.PathPrefix("/ospf/").Handler(http.StripPrefix("/ospf/", http.FileServer(http.Dir(ctx.ospf.infrastructure.GetOpts().FrontDir))))

	return s
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.Context()
	s.router.ServeHTTP(w, req)
}

func (s *httpServer) AllowOrigin(next httpapi.APIHandler) httpapi.APIHandler {
	return func(w http.ResponseWriter, r *http.Request) (response interface{}, err error) {
		w.Header().Set("Access-Control-Allow-Origin",s.ctx.ospf.infrastructure.GetOpts().AllowOrigin)
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		w.Header().Set("content-type", "text/html;application/json;charset=utf-8")
		return next(w, r)
	}
}

func initFrontEndMux(r *mux.Router) {
	for _, item := range frontEndRoutes {
		r.HandleFunc(item, view.LoadTemplate)
	}
}

func (s *httpServer) initRouter() {
	s.regLogController(s.router)
	s.initRouter_panjiawei()
	s.initRouter_chencanxin()
	s.initRouter_chenhuiliang()
	s.initRouter_guochuting()
	s.initRouter_guozhenzhen()
}

func (s *httpServer) regLogController(r *mux.Router) {
	c := &controller_chencanxin.LogController{
		LogModel: model_chencanxin.GetLogModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/test", httpapi.Decorate(c.GetLogs, s.AllowOrigin, middleware.Online,middleware.DefaultDecode,
		middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger)))
}
