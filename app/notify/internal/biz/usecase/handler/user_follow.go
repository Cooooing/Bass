package handler

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/usecase"
)

type UserFollowHandler struct {
}

func NewUserFollowHandler() *UserFollowHandler {
	return &UserFollowHandler{}
}

func (h *UserFollowHandler) Build(_ context.Context, event *enums.Event) (*usecase.NotificationIntent, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetUserFollow()
	if payload == nil {
		return nil, nil
	}
	intent := &usecase.NotificationIntent{
		EventID:   event.EventId,
		EventType: commonenum.EventTypeUserFollow,
		Vars:      payload,
	}
	if payload.FollowedId != 0 {
		intent.Station = append(intent.Station, &usecase.StationInput{UserID: payload.FollowedId})
	}
	return intent, nil
}
