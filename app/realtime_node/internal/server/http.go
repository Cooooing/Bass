package server

import (
	commonclient "common/pkg/client"
	commonserver "common/pkg/server"
	"fmt"
	"log/slog"
	"realtime_node/internal/config"

	"github.com/go-kratos/kratos/v3/transport/http"
)

func NewHTTPServer(config *config.Bootstrap, logger *slog.Logger, observer *commonclient.Observer, services []commonserver.Service) *http.Server {
	options := []http.ServerOption{
		http.Filter(commonserver.HTTPCORSFilter()),
		http.Middleware(commonserver.RequestLogContextMiddleware(), observer.ServerMiddleware()),
	}
	if config.Http.GetNetwork() != "" {
		options = append(options, http.Network(config.Http.GetNetwork()))
	}
	if config.Http.GetHost() != "" && config.Http.GetPort() != 0 {
		options = append(options, http.Address(fmt.Sprintf("%s:%d", config.Http.GetHost(), config.Http.GetPort())))
	}
	options = append(options, http.Timeout(0))
	server := http.NewServer(options...)
	for _, service := range services {
		service.RegisterHTTP(server)
	}
	return server
}
