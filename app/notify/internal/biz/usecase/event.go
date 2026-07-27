package usecase

import (
	"context"
	"log/slog"

	commonenum "common/pkg/enum"
	"common/proto/gen/common/enums"
	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
)

type EventHandler interface {
	Handle(ctx context.Context, req *EventHandleReq) error
	Templates() []*model.NotificationTemplateDefinition
}

type EventHandleReq struct {
	Event     *enums.Event
	EventType commonenum.EventType
	Language  notifyenum.Language
}

type EventUsecase struct {
	log           *slog.Logger
	eventHandlers map[commonenum.EventType]EventHandler
}

func NewEventUsecase(
	logger *slog.Logger,
	eventHandlers map[commonenum.EventType]EventHandler,
) *EventUsecase {
	return &EventUsecase{
		log:           logger,
		eventHandlers: eventHandlers,
	}
}

func (u *EventUsecase) Dispatch(ctx context.Context, req *EventHandleReq) error {
	if req == nil || req.Event == nil || req.Event.GetEventId() == "" {
		return nil
	}
	eventHandler := u.eventHandlers[req.EventType]
	if eventHandler == nil {
		return nil
	}
	return eventHandler.Handle(ctx, req)
}
