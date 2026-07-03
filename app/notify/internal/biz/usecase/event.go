package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	commonenum "common/pkg/enum"
	"common/proto/gen/common/enums"
	base "notify/internal/biz/base"
	"notify/internal/biz/repo"
	"notify/internal/conf"
	notifyenum "notify/internal/enum"

	"google.golang.org/protobuf/encoding/protojson"
	"log/slog"
)

type EventHandler interface {
	Build(ctx context.Context, event *enums.Event) (*NotificationContext, error)
}

type EventHandlers map[commonenum.EventType]EventHandler

type EventSubjects []commonenum.EventSubject

type EventUsecase struct {
	log            *slog.Logger
	conf           *conf.Bootstrap
	tx             base.Tx
	inboxEventRepo repo.InboxEventRepo
	notifyUsecase  *NotifyUsecase
	eventHandlers  EventHandlers
}

func NewEventUsecase(logger *slog.Logger, conf *conf.Bootstrap, tx base.Tx, inboxEventRepo repo.InboxEventRepo, notifyUsecase *NotifyUsecase, eventHandlers EventHandlers) *EventUsecase {
	return &EventUsecase{log: logger, conf: conf, tx: tx, inboxEventRepo: inboxEventRepo, notifyUsecase: notifyUsecase, eventHandlers: eventHandlers}
}

func (u *EventUsecase) HandleMessage(ctx context.Context, subjectName string, payload []byte) error {
	var event enums.Event
	if err := protojson.Unmarshal(payload, &event); err != nil {
		u.log.Error(fmt.Sprintf("unmarshal event failed: subject=%s err=%v", subjectName, err))
		return err
	}
	if event.EventId == "" {
		u.log.Error(fmt.Sprintf("event id is empty, inbox status cannot be saved: subject=%s type=%v", subjectName, event.Type))
		return errors.New("event id is required")
	}

	eventType, ok := commonenum.EventTypeMap.ToEnum(event.Type)
	if !ok {
		err := fmt.Errorf("unknown event type: %s", event.Type.String())
		u.log.Error(fmt.Sprintf("event type invalid: event_id=%s subject=%s type=%v err=%v", event.EventId, subjectName, event.Type, err))
		return err
	}
	eventSubject := commonenum.EventSubject(subjectName)
	if _, ok := commonenum.EventSubjectMap.ToProto(eventSubject); !ok {
		err := fmt.Errorf("unknown event subject: %s", subjectName)
		u.log.Error(fmt.Sprintf("event subject invalid: event_id=%s subject=%s type=%v err=%v", event.EventId, subjectName, event.Type, err))
		return err
	}
	if expectedSubject, ok := commonenum.EventSubjectByEventType(event.Type); !ok || expectedSubject != eventSubject {
		err := fmt.Errorf("event subject mismatch: subject=%s type=%s", subjectName, event.Type.String())
		u.log.Error(fmt.Sprintf("event subject mismatch: event_id=%s subject=%s type=%v err=%v", event.EventId, subjectName, event.Type, err))
		return err
	}

	now := time.Now()
	inboxEvent, claimed, err := u.inboxEventRepo.SaveProcessing(ctx, &repo.InboxEventSave{EventID: event.EventId, EventType: event.Type, Subject: eventSubject, Payload: string(payload)}, now)
	if err != nil {
		return err
	}
	if inboxEvent.Status == commonenum.InboxEventStatusProcessed || inboxEvent.Status == commonenum.InboxEventStatusDead {
		return nil
	}
	if !claimed {
		claimed, err = u.inboxEventRepo.ClaimRetry(ctx, event.EventId, now, u.inboxProcessingTimeout(), u.inboxMaxRetry())
		if err != nil {
			return err
		}
		if !claimed {
			return fmt.Errorf("event is processing: event_id=%s", event.EventId)
		}
	}
	rules, err := u.notifyUsecase.ListEnabledRules(ctx, eventType, notifyenum.LanguageZhCN)
	if err != nil {
		return u.markFailed(ctx, event.EventId, err)
	}
	if len(rules) == 0 {
		return u.inboxEventRepo.MarkProcessed(ctx, event.EventId, now)
	}
	eventHandler, ok := u.eventHandlers[eventType]
	if !ok {
		return u.inboxEventRepo.MarkProcessed(ctx, event.EventId, now)
	}
	notificationContext, err := eventHandler.Build(ctx, &event)
	if err != nil {
		return u.markFailed(ctx, event.EventId, err)
	}
	if notificationContext == nil {
		return u.inboxEventRepo.MarkProcessed(ctx, event.EventId, now)
	}
	notificationContext.EventID = event.EventId
	notificationContext.EventType = eventType
	notificationContext.Language = notifyenum.LanguageZhCN
	if err = u.notifyUsecase.Process(ctx, notificationContext, rules); err != nil {
		return u.markFailed(ctx, event.EventId, err)
	}
	return u.inboxEventRepo.MarkProcessed(ctx, event.EventId, now)
}

func (u *EventUsecase) markFailed(ctx context.Context, eventID string, err error) error {
	if markErr := u.inboxEventRepo.MarkFailed(ctx, eventID, err.Error(), u.inboxMaxRetry()); markErr != nil {
		u.log.Error(fmt.Sprintf("mark inbox failed status failed: event_id=%s err=%v", eventID, markErr))
	}
	return err
}

func (u *EventUsecase) inboxMaxRetry() int32 {
	if u.conf != nil && u.conf.GetEvent() != nil && u.conf.GetEvent().GetInbox() != nil && u.conf.GetEvent().GetInbox().GetMaxRetry() > 0 {
		return u.conf.GetEvent().GetInbox().GetMaxRetry()
	}
	return 10
}

func (u *EventUsecase) inboxProcessingTimeout() time.Duration {
	if u.conf != nil && u.conf.GetEvent() != nil && u.conf.GetEvent().GetInbox() != nil && u.conf.GetEvent().GetInbox().GetProcessingTimeout() != nil && u.conf.GetEvent().GetInbox().GetProcessingTimeout().AsDuration() > 0 {
		return u.conf.GetEvent().GetInbox().GetProcessingTimeout().AsDuration()
	}
	return 10 * time.Minute
}
