package usecase

import (
	"context"
	"log/slog"

	event "common/proto/gen/common/enums/event"

	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
)

type EventHandler interface {
	Handle(ctx context.Context, req *EventHandleReq) error
	Templates() []*model.NotificationTemplateDefinition
}

type EventHandleReq struct {
	Event            *event.Event
	NotificationType notifyenum.NotificationType
	Language         notifyenum.Language
}

type EventUsecase struct {
	log           *slog.Logger
	eventHandlers map[event.EventSubject][]EventHandler
}

func NewEventUsecase(
	logger *slog.Logger,
	eventHandlers map[event.EventSubject][]EventHandler,
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
	for _, handler := range u.eventHandlers[req.Event.GetSubject()] {
		if handler == nil {
			continue
		}
		if err := handler.Handle(ctx, req); err != nil {
			return err
		}
	}
	return nil
}
