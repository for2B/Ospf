package httpapi

import (
	"fmt"
	"go.uber.org/zap"
	"net"
	"net/http"
	"strings"
)

func Serve(listener net.Listener, handler http.Handler, logger *zap.Logger) {
	server := &http.Server{
		Handler: handler,
	}
	err := server.Serve(listener)
	if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
		logger.Error(fmt.Sprintf("http.Serve() - %s", err))
	}
	logger.Info(fmt.Sprintf("%s: closing", "http"),
		zap.String("port", listener.Addr().String()))
}
