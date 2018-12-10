package middleware

import (
	"net/http"
	"OutsourcingPlatform/internal/httpapi"
	"OutsourcingPlatform/internal/httpparse"
	"OutsourcingPlatform/ospf/constant"
)

func DefaultDecode(next httpapi.APIHandler) httpapi.APIHandler {
	return func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		ctx, err := httpparse.Parse(r)
		if err != nil {
			return nil, httpapi.NewErr(constant.GLOBAL_SYS_ERR, "httpparse.Parse(r) func failed", err)
		}
		return next(w, r.WithContext(ctx))
	}
}
