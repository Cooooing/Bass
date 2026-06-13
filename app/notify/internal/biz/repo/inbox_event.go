package repo

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common"
	"common/proto/gen/common/enums"
	"context"
	"notify/internal/biz/model"
	"time"
)

type InboxEventRepo interface {
	SaveProcessing(ctx context.Context, req *InboxEventSave, now time.Time) (*model.InboxEvent, bool, error)
	Get(ctx context.Context, req *InboxEventQuery) (*model.InboxEvent, error)
	List(ctx context.Context, req *InboxEventQuery) ([]*model.InboxEvent, error)
	Map(ctx context.Context, req *InboxEventQuery) (map[int64]*model.InboxEvent, error)
	Count(ctx context.Context, req *InboxEventQuery) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *InboxEventQuery) ([]*model.InboxEvent, *common.PageReply, error)
	ClaimRetry(ctx context.Context, eventID string, now time.Time, processingTimeout time.Duration, maxRetry int32) (bool, error)
	MarkProcessed(ctx context.Context, eventID string, now time.Time) error
	MarkFailed(ctx context.Context, eventID string, lastError string, maxRetry int32) error
}

type InboxEventSave struct {
	EventID   string
	EventType enums.EventType
	Subject   commonenum.EventSubject
	Payload   string
}

type InboxEventQuery struct {
	ID        *int64
	IDs       []int64
	EventID   *string
	EventIDs  []string
	EventType *commonenum.EventType
	Subject   *commonenum.EventSubject
	Status    *commonenum.InboxEventStatus
}
