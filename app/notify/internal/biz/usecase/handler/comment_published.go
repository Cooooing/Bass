package handler

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/usecase"
)

type CommentPublishedHandler struct {
}

func NewCommentPublishedHandler() *CommentPublishedHandler {
	return &CommentPublishedHandler{}
}

func (h *CommentPublishedHandler) Build(_ context.Context, event *enums.Event) (*usecase.NotificationIntent, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetCommentPublished()
	if payload == nil {
		return nil, nil
	}
	intent := &usecase.NotificationIntent{
		EventID:   event.EventId,
		EventType: commonenum.EventTypeContentCommentPublish,
		Vars:      payload,
	}
	if payload.AuthorId != 0 && payload.AuthorId != payload.SenderId {
		intent.Station = append(intent.Station, &usecase.StationInput{UserID: payload.AuthorId})
	}
	return intent, nil
}
