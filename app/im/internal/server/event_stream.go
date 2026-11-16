package server

import (
	"context"

	"im/internal/biz/usecase"
)

type EventStreamServer struct {
	globalChat *usecase.GlobalChatUsecase
}

func NewEventStreamServer(globalChat *usecase.GlobalChatUsecase) *EventStreamServer {
	return &EventStreamServer{globalChat: globalChat}
}

func (s *EventStreamServer) Start(ctx context.Context) error {
	return s.globalChat.EnsureEventStream(ctx)
}

func (*EventStreamServer) Stop(context.Context) error {
	return nil
}
