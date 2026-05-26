package usecase

import (
	"common/api/gen/common/enums"
	commonenum "common/pkg/enum"
	"context"
	base "notify/internal/biz/base"
	"notify/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/proto"
)

// EventHandler 将一种跨服务事件转换为通知意图。
type EventHandler interface {
	Build(ctx context.Context, event *enums.Event) (*NotificationIntent, error)
}

type EventHandlers map[commonenum.EventType]EventHandler

// EventSubjects 是 notify 订阅的跨服务事件主题集合。
type EventSubjects []commonenum.EventSubject

type EventUsecase struct {
	log            *log.Helper
	tx             base.Tx
	inboxEventRepo repo.InboxEventRepo
	notifyUsecase  *NotifyUsecase
	eventHandlers  EventHandlers
}

func NewEventUsecase(
	logger log.Logger,
	tx base.Tx,
	inboxEventRepo repo.InboxEventRepo,
	notifyUsecase *NotifyUsecase,
	eventHandlers EventHandlers,
) *EventUsecase {
	return &EventUsecase{
		log:            log.NewHelper(logger),
		tx:             tx,
		inboxEventRepo: inboxEventRepo,
		notifyUsecase:  notifyUsecase,
		eventHandlers:  eventHandlers,
	}
}

func (u *EventUsecase) HandleMessage(ctx context.Context, subjectName string, payload []byte) error {
	var event enums.Event
	if err := proto.Unmarshal(payload, &event); err != nil {
		u.log.Errorf("unmarshal event failed: subject=%s err=%v", subjectName, err)
		return nil
	}
	if event.EventId == "" {
		u.log.Errorf("event id is empty: subject=%s type=%v", subjectName, event.Type)
		return nil
	}
	eventType, ok := commonenum.EventTypeMap.ToEnum(event.Type)
	if !ok {
		u.log.Errorf("unknown event type: subject=%s type=%v", subjectName, event.Type)
		return nil
	}
	subject := commonenum.EventSubject(subjectName)
	if _, ok := commonenum.EventSubjectMap.ToProto(subject); !ok {
		u.log.Errorf("unknown event subject: subject=%s type=%v", subjectName, event.Type)
		return nil
	}
	if expectedSubject, ok := commonenum.EventSubjectByEventType(event.Type); !ok || expectedSubject != subject {
		u.log.Errorf("event subject mismatch: subject=%s type=%v", subjectName, event.Type)
		return nil
	}

	inboxEvent, err := u.inboxEventRepo.SaveReceived(ctx, &repo.InboxEventSave{
		EventID:   event.EventId,
		EventType: event.Type,
		Subject:   subject,
		Payload:   payload,
	})
	if err != nil {
		return err
	}
	if inboxEvent.Status == commonenum.InboxEventStatusProcessed {
		return nil
	}

	claimed, err := u.inboxEventRepo.MarkProcessing(ctx, event.EventId)
	if err != nil {
		return err
	}
	if !claimed {
		return nil
	}

	eventHandler, ok := u.eventHandlers[eventType]
	if !ok {
		return u.inboxEventRepo.MarkProcessed(ctx, event.EventId)
	}
	intent, err := eventHandler.Build(ctx, &event)
	if err != nil {
		return u.markFailed(ctx, event.EventId, err)
	}
	prepared, err := u.notifyUsecase.prepare(ctx, intent)
	if err != nil {
		return u.markFailed(ctx, event.EventId, err)
	}

	err = u.tx(ctx, func(ctx context.Context) error {
		if prepared != nil && (prepared.station != nil || len(prepared.deliveries) > 0) {
			if err := u.notifyUsecase.savePrepared(ctx, prepared); err != nil {
				return err
			}
		}
		return u.inboxEventRepo.MarkProcessed(ctx, event.EventId)
	})
	if err != nil {
		return u.markFailed(ctx, event.EventId, err)
	}
	return nil
}

func (u *EventUsecase) markFailed(ctx context.Context, eventID string, err error) error {
	if markErr := u.inboxEventRepo.MarkFailed(ctx, eventID); markErr != nil {
		u.log.Errorf("mark inbox failed status failed: event_id=%s err=%v", eventID, markErr)
	}
	return err
}
