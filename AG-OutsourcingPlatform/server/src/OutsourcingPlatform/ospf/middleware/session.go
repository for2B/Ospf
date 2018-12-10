package middleware

import (
	"context"
	"net/http"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/ospf/constant"
	"OutsourcingPlatform/ospf/infrastructure/session"
)

func Session(sf *session.SessionManager) httpapi.MiddlewareFunc {
	return func(next httpapi.APIHandler) httpapi.APIHandler {
		return func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
			sess, err := sf.FetchSession(w, r)
			if err != nil {
				return nil, httpapi.NewErr(constant.SESSION_FETCH_ERR,
					"FetchSession() faild in session middleware",
					err)
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, "session", sess)
			return next(w, r.WithContext(ctx))
		}
	}
}
