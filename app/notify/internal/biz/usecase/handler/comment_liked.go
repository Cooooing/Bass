package handler

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/usecase"
)

type CommentLikedHandler struct {
}

func NewCommentLikedHandler() *CommentLikedHandler {
	return &CommentLikedHandler{}
}

func (h *CommentLikedHandler) Build(_ context.Context, event *enums.Event) (*usecase.NotificationIntent, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetCommentLiked()
	if payload == nil {
		return nil, nil
	}
	intent := &usecase.NotificationIntent{
		EventID:   event.EventId,
		EventType: commonenum.EventTypeContentCommentLike,
		Vars:      payload,
	}
	if payload.CommentAuthorId != 0 && payload.CommentAuthorId != payload.SenderId {
		intent.Station = append(intent.Station, &usecase.StationInput{UserID: payload.CommentAuthorId})
	}
	return intent, nil
}
