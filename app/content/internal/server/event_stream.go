package server

import (
	"context"

	"content/internal/biz/usecase"
)

type EventStreamServer struct {
	outbox *usecase.OutboxUsecase
}

func NewEventStreamServer(outbox *usecase.OutboxUsecase) *EventStreamServer {
	return &EventStreamServer{outbox: outbox}
}

func (s *EventStreamServer) Start(ctx context.Context) error {
	return s.outbox.EnsureEventStream(ctx)
}

func (*EventStreamServer) Stop(context.Context) error {
	return nil
}
