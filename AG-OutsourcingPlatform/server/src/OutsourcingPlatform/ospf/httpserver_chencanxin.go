package ospf

import (
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/controller/controller_chencanxin"
	"OutsourcingPlatform/ospf/model/model_chencanxin"
	"net/http"

	"github.com/gorilla/mux"
	"OutsourcingPlatform/ospf/middleware"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/ospf/service/chencanxin/wechat"
)

func (s *httpServer) initRouter_chencanxin() {
	s.regWechatOauthController(s.router)
	s.regValidSmsVerificationCodeController(s.router)
	s.regGetSmsVerificationCodeController(s.router)
	s.regSelectTableDateController(s.router)
	s.regUpdateTableDateController(s.router)
	s.regWeChatConf(s.router)
	s.regMockOauthController(s.router)
	s.regValidQrCodeController(s.router)
	s.regFetchQrCodeController(s.router)
	s.regQrcodeLoginController(s.router)
	s.regClearQrcodeController(s.router)
	s.regResetPasswordController(s.router)
	s.regGenerateGolangApiCodeController(s.router)

}

func (s *httpServer) regWechatOauthController(r *mux.Router) {
	c := &controller_chencanxin.WcAuthController{
		model_chencanxin.GetWcAuthModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/wechat_login", httpapi.Decorate(c.WechatAuth,
		middleware.DefaultDecode,
		middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger),
		middleware.Online))
}

func (s *httpServer) regWeChatConf(r *mux.Router) {
	r.Handle("/api/wechat_config_domain", http.HandlerFunc(wechat.ProcRequest))
}

func (s *httpServer) regMockOauthController(r *mux.Router) {
	c := &controller_chencanxin.MockLoginController{
		s.ctx.ospf.infrastructure,
	}
	r.Handle("/api/mock_login", httpapi.Decorate(c.MockLogin,
		middleware.DefaultDecode,
		middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger),
		middleware.Online))
}

func (s *httpServer) regValidQrCodeController(r *mux.Router) {
	c := &controller_chencanxin.ValidQRCodeController{
		s.ctx.ospf.infrastructure,
	}
	r.Handle("/api/valid_invitation_code", httpapi.Decorate(c.Valid,
		s.AllowOrigin,
		middleware.DefaultDecode,
		middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger),
		middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
		middleware.Permission(constant.USER_NORMAL), middleware.Online))
}

func (s *httpServer) regFetchQrCodeController(r *mux.Router) {
	c := &controller_chencanxin.FecthInvitationCodeController{
		FetchInvitationCodeModel: model_chencanxin.GetFetchInviCodeModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/fetch_invitation_code", httpapi.Decorate(c.Fetch, s.AllowOrigin, middleware.DefaultDecode, middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger),
		middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
		middleware.Permission(constant.USER_NORMAL), middleware.Online))
}

func (s *httpServer) regQrcodeLoginController(r *mux.Router) {
	c := &controller_chencanxin.QrcodeLoginController{
		s.ctx.ospf.infrastructure,
	}
	r.Handle("/api/qrcode_login", httpapi.Decorate(c.QrcodeLogin,
		middleware.DefaultDecode,
		middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger),
		middleware.Online))
}

func (s *httpServer) regClearQrcodeController(r *mux.Router) {
	c := &controller_chencanxin.ClearQrcodeController{
		s.ctx.ospf.infrastructure,
	}
	r.Handle("/api/clear_qrcode", httpapi.Decorate(c.Clear,
		s.AllowOrigin, middleware.DefaultDecode, middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger),
		middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
		middleware.Permission(constant.USER_SUPER),
		middleware.Online))
}

func (s *httpServer) regResetPasswordController(r *mux.Router) {
	c := &controller_chencanxin.ResetPasswordController{
		ResetPasswordModel: model_chencanxin.GetResetPasswordModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/resetPassword", httpapi.Decorate(c.ToResetPassword,
		s.AllowOrigin,
		middleware.DefaultDecode,
		middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger),
		middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
		middleware.Permission(constant.USER_VISITOR),
		middleware.Online),
	)
}

func (s *httpServer) regGenerateGolangApiCodeController(r *mux.Router) {
	c := &controller_chencanxin.GenerateGolangApiCodeController{
		GenerateGolangApiCodeModel: model_chencanxin.GetGenerateGolangApiCodeModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/generateGolangApiCode", httpapi.Decorate(c.ToGenerateGolangApiCode,
		s.AllowOrigin,
		middleware.DefaultDecode,
		middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger),
		middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
		middleware.Permission(constant.USER_SUPER),
		middleware.Online),
	)
}



func (s *httpServer) regUpdateTableDateController(r *mux.Router) {
	c := &controller_chencanxin.UpdateTableDateController{
		UpdateTableDateModel: model_chencanxin.GetUpdateTableDateModel(s.ctx.ospf.infrastructure),
	}
	r.Handle("/api/updateTableDate", httpapi.Decorate(c.ToUpdateTableDate,
		s.AllowOrigin,
		middleware.DefaultDecode,
		middleware.Encode,
		middleware.Log(s.ctx.ospf.infrastructure.Logger),
		middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
		middleware.Permission(constant.USER_SUPER),
		middleware.Online),
	)
}

  
  func (s *httpServer) regSelectTableDateController(r *mux.Router) {
      c := &controller_chencanxin.SelectTableDateController{
        SelectTableDateModel: model_chencanxin.GetSelectTableDateModel(s.ctx.ospf.infrastructure),
      }
      r.Handle("/api/selectTableDate", httpapi.Decorate(c.ToSelectTableDate, 
        s.AllowOrigin,
        middleware.DefaultDecode, 
        middleware.Encode,
        middleware.Log(s.ctx.ospf.infrastructure.Logger),
        middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
        middleware.Permission(constant.USER_SUPER),
        middleware.Online),
        )
  }
    
  
  func (s *httpServer) regGetSmsVerificationCodeController(r *mux.Router) {
      c := &controller_chencanxin.GetSmsVerificationCodeController{
        GetSmsVerificationCodeModel: model_chencanxin.GetGetSmsVerificationCodeModel(s.ctx.ospf.infrastructure),
      }
      r.Handle("/api/getSmsVerificationCode", httpapi.Decorate(c.ToGetSmsVerificationCode, 
        s.AllowOrigin,
        middleware.DefaultDecode, 
        middleware.Encode,
        middleware.Log(s.ctx.ospf.infrastructure.Logger),
        middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
        middleware.Permission(constant.USER_SUPER),
        middleware.Online),
        )
  }
    
  
  func (s *httpServer) regValidSmsVerificationCodeController(r *mux.Router) {
      c := &controller_chencanxin.ValidSmsVerificationCodeController{
        ValidSmsVerificationCodeModel: model_chencanxin.GetValidSmsVerificationCodeModel(s.ctx.ospf.infrastructure),
      }
      r.Handle("/api/ValidSmsVerificationCode", httpapi.Decorate(c.ToValidSmsVerificationCode, 
        s.AllowOrigin,
        middleware.DefaultDecode, 
        middleware.Encode,
        middleware.Log(s.ctx.ospf.infrastructure.Logger),
        middleware.Session(s.ctx.ospf.infrastructure.SessionManager),
        middleware.Permission(constant.USER_VISITOR),
        middleware.Online),
        )
  }
    