package handler

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/usecase"
)

type ArticleWatchedHandler struct {
}

func NewArticleWatchedHandler() *ArticleWatchedHandler {
	return &ArticleWatchedHandler{}
}

func (h *ArticleWatchedHandler) Build(_ context.Context, event *enums.Event) (*usecase.NotificationIntent, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetArticleWatched()
	if payload == nil {
		return nil, nil
	}
	intent := &usecase.NotificationIntent{
		EventID:   event.EventId,
		EventType: commonenum.EventTypeContentArticleWatch,
		Vars:      payload,
	}
	if payload.AuthorId != 0 && payload.AuthorId != payload.SenderId {
		intent.Station = append(intent.Station, &usecase.StationInput{UserID: payload.AuthorId})
	}
	return intent, nil
}
