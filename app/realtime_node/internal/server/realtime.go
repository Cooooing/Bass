package server

import (
	"context"

	"realtime_node/internal/biz/usecase"
)

type RealtimeServer struct {
	realtime *usecase.RealtimeUsecase
}

func NewRealtimeServer(realtime *usecase.RealtimeUsecase) *RealtimeServer {
	return &RealtimeServer{realtime: realtime}
}

func (s *RealtimeServer) Start(ctx context.Context) error {
	return s.realtime.Start(ctx)
}

func (s *RealtimeServer) Stop(ctx context.Context) error {
	return s.realtime.Stop(ctx)
}
