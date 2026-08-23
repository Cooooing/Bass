package server

import (
	"context"
	"errors"
	monolithmodule "monolith/internal/module"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
	ProvideServers,
)

func ProvideServers(grpcServer *grpc.Server, httpServer *http.Server, registry *monolithmodule.Registry) []transport.Server {
	servers := []transport.Server{
		grpcServer,
		httpServer,
	}
	internalServers := registry.Servers()
	if len(internalServers) > 0 {
		servers = append(servers, newInternalServer(internalServers))
	}
	return servers
}

type internalServer struct {
	servers []transport.Server
	started []transport.Server
}

func newInternalServer(servers []transport.Server) *internalServer {
	return &internalServer{servers: servers}
}

func (s *internalServer) Start(ctx context.Context) error {
	for _, server := range s.servers {
		if server == nil {
			continue
		}
		if err := server.Start(ctx); err != nil {
			return errors.Join(err, s.stopStarted(ctx))
		}
		s.started = append(s.started, server)
	}
	return nil
}

func (s *internalServer) Stop(ctx context.Context) error {
	return s.stopStarted(ctx)
}

func (s *internalServer) stopStarted(ctx context.Context) error {
	var err error
	for i := len(s.started) - 1; i >= 0; i-- {
		if stopErr := s.started[i].Stop(ctx); stopErr != nil {
			err = errors.Join(err, stopErr)
		}
	}
	s.started = nil
	return err
}
