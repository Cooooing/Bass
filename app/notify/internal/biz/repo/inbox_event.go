package repo

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/model"
)

type InboxEventRepo interface {
	SaveReceived(ctx context.Context, req *InboxEventSave) (*model.InboxEvent, error)
	MarkProcessing(ctx context.Context, eventID string) (bool, error)
	MarkProcessed(ctx context.Context, eventID string) error
	MarkFailed(ctx context.Context, eventID string) error
}

type InboxEventSave struct {
	EventID   string
	EventType enums.EventType
	Subject   commonenum.EventSubject
	Payload   []byte
}
