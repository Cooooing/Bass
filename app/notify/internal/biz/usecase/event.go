package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	commonenum "common/pkg/enum"
	"common/proto/gen/common/enums"
	"notify/internal/biz/base"
	"notify/internal/biz/repo"
	"notify/internal/config"
	notifyenum "notify/internal/enum"

	"log/slog"

	"google.golang.org/protobuf/encoding/protojson"
)

type EventHandler interface {
	Build(ctx context.Context, event *enums.Event) (*NotificationContext, error)
}

type EventHandlers map[commonenum.EventType]EventHandler

type EventSubjects []commonenum.EventSubject

type EventUsecase struct {
	log            *slog.Logger
	conf           *config.Bootstrap
	tx             base.Tx
	inboxEventRepo repo.InboxEventRepo
	notifyUsecase  *NotifyUsecase
	eventHandlers  EventHandlers
}

func NewEventUsecase(logger *slog.Logger, conf *config.Bootstrap, tx base.Tx, inboxEventRepo repo.InboxEventRepo, notifyUsecase *NotifyUsecase, eventHandlers EventHandlers) *EventUsecase {
	return &EventUsecase{log: logger, conf: conf, tx: tx, inboxEventRepo: inboxEventRepo, notifyUsecase: notifyUsecase, eventHandlers: eventHandlers}
}

type EventHandleMessageReq struct {
	SubjectName string
	Payload     []byte
}

func (u *EventUsecase) HandleMessage(ctx context.Context, req *EventHandleMessageReq) error {
	if req == nil {
		return errors.New("event handle message request is required")
	}
	subjectName := req.SubjectName
	payload := req.Payload
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
	saveResp, err := u.inboxEventRepo.SaveProcessing(ctx, &repo.InboxEventSaveProcessingReq{EventID: event.EventId, EventType: event.Type, Subject: eventSubject, Payload: string(payload), Now: now})
	if err != nil {
		return err
	}
	inboxEvent := saveResp.Event
	claimed := saveResp.Claimed
	if inboxEvent.Status == commonenum.InboxEventStatusProcessed || inboxEvent.Status == commonenum.InboxEventStatusDead {
		return nil
	}
	if !claimed {
		claimedRetry, err := u.inboxEventRepo.ClaimRetry(ctx, &repo.InboxEventClaimRetryReq{EventID: event.EventId, Now: now, ProcessingTimeout: u.inboxProcessingTimeout(), MaxRetry: u.inboxMaxRetry()})
		if err != nil {
			return err
		}
		if !claimedRetry {
			return fmt.Errorf("event is processing: event_id=%s", event.EventId)
		}
	}
	rules, err := u.notifyUsecase.ListEnabledRules(ctx, &NotifyListEnabledRulesReq{
		EventType: eventType,
		Language:  notifyenum.LanguageZhCN,
	})
	if err != nil {
		return u.markFailed(ctx, &markFailedReq{EventID: event.EventId, Err: err})
	}
	if len(rules) == 0 {
		err = u.inboxEventRepo.MarkProcessed(ctx, &repo.InboxEventMarkProcessedReq{EventID: event.EventId, Now: now})
		return err
	}
	eventHandler, ok := u.eventHandlers[eventType]
	if !ok {
		err = u.inboxEventRepo.MarkProcessed(ctx, &repo.InboxEventMarkProcessedReq{EventID: event.EventId, Now: now})
		return err
	}
	notificationContext, err := eventHandler.Build(ctx, &event)
	if err != nil {
		return u.markFailed(ctx, &markFailedReq{EventID: event.EventId, Err: err})
	}
	if notificationContext == nil {
		err = u.inboxEventRepo.MarkProcessed(ctx, &repo.InboxEventMarkProcessedReq{EventID: event.EventId, Now: now})
		return err
	}
	notificationContext.EventID = event.EventId
	notificationContext.EventType = eventType
	notificationContext.Language = notifyenum.LanguageZhCN
	if err = u.notifyUsecase.Process(ctx, &NotifyProcessReq{NotificationContext: notificationContext, Rules: rules}); err != nil {
		return u.markFailed(ctx, &markFailedReq{EventID: event.EventId, Err: err})
	}
	err = u.inboxEventRepo.MarkProcessed(ctx, &repo.InboxEventMarkProcessedReq{EventID: event.EventId, Now: now})
	return err
}

type markFailedReq struct {
	EventID string
	Err     error
}

func (u *EventUsecase) markFailed(ctx context.Context, req *markFailedReq) error {
	eventID := req.EventID
	err := req.Err
	if markErr := u.inboxEventRepo.MarkFailed(ctx, &repo.InboxEventMarkFailedReq{EventID: eventID, LastError: err.Error(), MaxRetry: u.inboxMaxRetry()}); markErr != nil {
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
