package repo

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	"notify/internal/biz/model"
	"time"
)

type InboxEventRepo interface {
	SaveProcessing(ctx context.Context, req *InboxEventSave, now time.Time) (*model.InboxEvent, bool, error)
	ClaimRetry(ctx context.Context, eventID string, now time.Time, processingTimeout time.Duration) (bool, error)
	MarkProcessed(ctx context.Context, eventID string, now time.Time) error
	MarkFailed(ctx context.Context, eventID string, lastError string) error
}

type InboxEventSave struct {
	EventID   string
	EventType enums.EventType
	Subject   commonenum.EventSubject
	Payload   []byte
}
