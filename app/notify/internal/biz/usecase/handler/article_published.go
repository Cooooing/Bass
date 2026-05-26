package handler

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/usecase"
)

type ArticlePublishedHandler struct {
	userClient usecase.UserClient
}

func NewArticlePublishedHandler(userClient usecase.UserClient) *ArticlePublishedHandler {
	return &ArticlePublishedHandler{userClient: userClient}
}

func (h *ArticlePublishedHandler) Build(ctx context.Context, event *enums.Event) (*usecase.NotificationIntent, error) {
	if event == nil || event.EventId == "" {
		return nil, nil
	}
	payload := event.GetArticlePublished()
	if payload == nil {
		return nil, nil
	}
	receiverIDs := []int64(nil)
	if h.userClient != nil {
		var err error
		receiverIDs, err = h.userClient.ListFollowerIDs(ctx, payload.SenderId)
		if err != nil {
			return nil, err
		}
	}
	intent := &usecase.NotificationIntent{
		EventID:   event.EventId,
		EventType: commonenum.EventTypeContentArticlePublish,
		Vars:      payload,
	}
	for _, receiverID := range receiverIDs {
		if receiverID != 0 {
			intent.Station = append(intent.Station, &usecase.StationInput{UserID: receiverID})
		}
	}
	return intent, nil
}
