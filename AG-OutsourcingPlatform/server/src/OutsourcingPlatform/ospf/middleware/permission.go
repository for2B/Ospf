package middleware

import (
	"net/http"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
)

func Permission(LevelNeeded string) httpapi.MiddlewareFunc {
	return func(next httpapi.APIHandler) httpapi.APIHandler {
		return func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
			session, ok := r.Context().Value("session").(*constant.Session)
			if !ok {
				return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "session assert failed in permission middleware",
					nil)
			}
			userLevel := session.Level
			if session.Openid=="" {
				userLevel = constant.USER_VISITOR
			}
			//judge the level
			switch userLevel {
			case constant.USER_SUPER:
				return next(w, r)
			case constant.USER_ADMIN:
				if LevelNeeded != constant.USER_SUPER {
					return next(w, r)
				}
			case constant.USER_NORMAL:
				if LevelNeeded != constant.USER_SUPER&&LevelNeeded!=constant.USER_SUPER {
					return next(w, r)
				}

			case constant.USER_VISITOR:
				if LevelNeeded == constant.USER_VISITOR {
					return next(w, r)
				}

			}
			return nil, httpapi.Err{constant.GLOBAL_NO_PERMISSION, "No permission!"}


		}
	}
}
