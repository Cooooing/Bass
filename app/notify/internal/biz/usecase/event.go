package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	event "common/proto/gen/common/enums/event/v1"

	"notify/internal/biz/model"
	notifyenum "notify/internal/enum"
)

var errEventHandlerNotFound = errors.New("event handler not found")

type EventHandler interface {
	Handle(ctx context.Context, req *EventHandleReq) error
	Templates() []*model.NotificationTemplateDefinition
}

type EventHandleReq struct {
	Event            *event.Event
	NotificationType notifyenum.NotificationType
	Language         notifyenum.Language
}

type EventHandlerBinding struct {
	Handler          EventHandler
	NotificationType notifyenum.NotificationType
}

type EventUsecase struct {
	log           *slog.Logger
	eventHandlers map[event.EventSubject][]EventHandlerBinding
}

func NewEventUsecase(
	logger *slog.Logger,
	eventHandlers map[event.EventSubject][]EventHandlerBinding,
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
	bindings := u.eventHandlers[req.Event.GetSubject()]
	if len(bindings) == 0 {
		return fmt.Errorf("%w: %s", errEventHandlerNotFound, req.Event.GetSubject())
	}
	handled := false
	for _, binding := range bindings {
		if binding.Handler == nil {
			continue
		}
		handled = true
		handlerReq := *req
		handlerReq.NotificationType = binding.NotificationType
		if err := binding.Handler.Handle(ctx, &handlerReq); err != nil {
			return err
		}
	}
	if !handled {
		return fmt.Errorf("%w: %s", errEventHandlerNotFound, req.Event.GetSubject())
	}
	return nil
}
