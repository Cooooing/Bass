package server

import (
	"bff_game_idle/internal/biz/usecase"
	"context"
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
