package repo

import (
	commonenums "common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
)

// OutboxEventSave 定义一条待持久化的 outbox 事件。
type OutboxEventSave struct {
	EventType commonenum.EventType
	Subject   commonenum.EventSubject
	Event     *commonenums.Event
}

type OutboxEventRepo interface {
	Save(ctx context.Context, req *OutboxEventSave) error
}
