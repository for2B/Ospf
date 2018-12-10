package httpapi

import (
	"net/http"
)

type MiddlewareFunc func(APIHandler) APIHandler

type APIHandler func(w http.ResponseWriter, r *http.Request) (response interface{}, err error)

func Decorate(h APIHandler, ds ...MiddlewareFunc) http.HandlerFunc {
	decorated := h
	for i := len(ds) - 1; i >= 0; i-- {
		decorated = ds[i](decorated)
	}
	return func(w http.ResponseWriter, r *http.Request) {

		decorated(w, r)
	}
}

func CombinationMiddleware(ds ...MiddlewareFunc) MiddlewareFunc {
	return func(w APIHandler) APIHandler {
		for i := len(ds) - 1; i >= 0; i-- {
			w = ds[i](w)
		}
		return w
	}
}
