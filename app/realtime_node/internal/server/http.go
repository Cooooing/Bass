package server

import (
	commonclient "common/pkg/client"
	commonserver "common/pkg/server"
	servermiddleware "common/pkg/server/middleware"
	"fmt"
	"log/slog"

	"realtime_node/internal/config"

	"github.com/go-kratos/kratos/v3/transport/http"
)

func NewHTTPServer(config *config.Bootstrap, logger *slog.Logger, observer *commonclient.Observer, services []commonserver.Service) *http.Server {
	options := []http.ServerOption{
		http.Filter(servermiddleware.HTTPCORSFilter()),
		http.Middleware(servermiddleware.RequestLogContextMiddleware(), observer.ServerMiddleware()),
	}
	if config.GetHttp().GetNetwork() != "" {
		options = append(options, http.Network(config.GetHttp().GetNetwork()))
	}
	if config.GetHttp().GetHost() != "" && config.GetHttp().GetPort() != 0 {
		options = append(options, http.Address(fmt.Sprintf("%s:%d", config.GetHttp().GetHost(), config.GetHttp().GetPort())))
	}
	options = append(options, http.Timeout(0))
	server := http.NewServer(options...)
	for _, service := range services {
		service.RegisterHTTP(server)
	}
	return server
}
