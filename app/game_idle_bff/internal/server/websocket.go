package server

import (
	"context"
	"game_idle_bff/internal/biz/usecase"
)

type WebSocketManagerServer struct {
	webSocketUsecase *usecase.WebSocketUsecase
}

func NewWebSocketManagerServer(webSocketUsecase *usecase.WebSocketUsecase) *WebSocketManagerServer {
	return &WebSocketManagerServer{
		webSocketUsecase: webSocketUsecase,
	}
}

func (s *WebSocketManagerServer) Start(ctx context.Context) error {
	return s.webSocketUsecase.Start(ctx)
}

func (s *WebSocketManagerServer) Stop(ctx context.Context) error {
	return s.webSocketUsecase.Stop(ctx)
}
