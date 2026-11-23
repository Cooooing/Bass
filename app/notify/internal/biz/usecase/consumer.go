package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	commonenum "common/pkg/enum"
	event "common/proto/gen/common/enums/event/v1"

	"notify/internal/biz/repo"
	"notify/internal/config"
	notifyenum "notify/internal/enum"
)

type ConsumerUsecase struct {
	log            *slog.Logger
	conf           *config.Bootstrap
	inboxEventRepo repo.InboxEventRepo
	eventUsecase   *EventUsecase
}

func NewConsumerUsecase(
	logger *slog.Logger,
	conf *config.Bootstrap,
	inboxEventRepo repo.InboxEventRepo,
	eventUsecase *EventUsecase,
) *ConsumerUsecase {
	return &ConsumerUsecase{
		log:            logger,
		conf:           conf,
		inboxEventRepo: inboxEventRepo,
		eventUsecase:   eventUsecase,
	}
}

type ConsumerHandleReq struct {
	Event   *event.Event
	Subject commonenum.EventSubject
	Payload string
}

func (u *ConsumerUsecase) Handle(ctx context.Context, req *ConsumerHandleReq) error {
	if req == nil || req.Event == nil || req.Event.GetEventId() == "" {
		return nil
	}
	notificationType, ok := u.notificationType(req.Event.GetSubject())
	if !ok {
		return fmt.Errorf("unsupported event subject: %s", req.Event.GetSubject())
	}
	maxRetry := int32(3)
	if u.conf != nil && u.conf.GetEvent() != nil && u.conf.GetEvent().GetInbox() != nil && u.conf.GetEvent().GetInbox().GetMaxRetry() > 0 {
		maxRetry = u.conf.GetEvent().GetInbox().GetMaxRetry()
	}
	processingTimeout := time.Minute
	if u.conf != nil && u.conf.GetEvent() != nil && u.conf.GetEvent().GetInbox() != nil && u.conf.GetEvent().GetInbox().GetProcessingTimeout() != nil && u.conf.GetEvent().GetInbox().GetProcessingTimeout().AsDuration() > 0 {
		processingTimeout = u.conf.GetEvent().GetInbox().GetProcessingTimeout().AsDuration()
	}
	now := time.Now()
	saveResp, err := u.inboxEventRepo.SaveProcessing(ctx, &repo.InboxEventSaveProcessingReq{
		EventID: req.Event.GetEventId(),
		Subject: req.Subject,
		Payload: req.Payload,
		Now:     now,
	})
	if err != nil {
		return err
	}
	if saveResp.Event != nil && (saveResp.Event.Status == commonenum.InboxEventStatusProcessed || saveResp.Event.Status == commonenum.InboxEventStatusDead) {
		return nil
	}
	if !saveResp.Claimed {
		claimedRetry, err := u.inboxEventRepo.ClaimRetry(ctx, &repo.InboxEventClaimRetryReq{
			EventID:           req.Event.GetEventId(),
			Now:               now,
			ProcessingTimeout: processingTimeout,
			MaxRetry:          maxRetry,
		})
		if err != nil {
			return err
		}
		if !claimedRetry {
			return fmt.Errorf("event is processing: event_id=%s", req.Event.GetEventId())
		}
	}
	if err := u.eventUsecase.Dispatch(ctx, &EventHandleReq{
		Event:            req.Event,
		NotificationType: notificationType,
		Language:         notifyenum.LanguageZhCN,
	}); err != nil {
		if markErr := u.inboxEventRepo.MarkFailed(ctx, &repo.InboxEventMarkFailedReq{
			EventID:   req.Event.GetEventId(),
			LastError: err.Error(),
			MaxRetry:  maxRetry,
		}); markErr != nil && u.log != nil {
			u.log.Error(fmt.Sprintf("mark inbox failed status failed: event_id=%s err=%v", req.Event.GetEventId(), markErr))
		}
		return err
	}
	return u.inboxEventRepo.MarkProcessed(ctx, &repo.InboxEventMarkProcessedReq{
		EventID: req.Event.GetEventId(),
		Now:     now,
	})
}

func (u *ConsumerUsecase) notificationType(subject event.EventSubject) (notifyenum.NotificationType, bool) {
	switch subject {
	case event.EventSubject_EVENT_SUBJECT_USER_FOLLOW:
		return notifyenum.NotificationTypeUserFollow, true
	case event.EventSubject_EVENT_SUBJECT_USER_REGISTER:
		return notifyenum.NotificationTypeUserRegister, true
	case event.EventSubject_EVENT_SUBJECT_USER_EMAIL_VERIFICATION_CODE:
		return notifyenum.NotificationTypeUserEmailVerificationCode, true
	case event.EventSubject_EVENT_SUBJECT_USER_PHONE_VERIFICATION_CODE:
		return notifyenum.NotificationTypeUserPhoneVerificationCode, true
	case event.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED:
		return notifyenum.NotificationTypeContentArticlePublish, true
	case event.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED:
		return notifyenum.NotificationTypeContentArticleLike, true
	case event.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED:
		return notifyenum.NotificationTypeContentArticleThank, true
	case event.EventSubject_EVENT_SUBJECT_ARTICLE_REWARDED:
		return notifyenum.NotificationTypeContentArticleReward, true
	case event.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED:
		return notifyenum.NotificationTypeContentArticleCollect, true
	case event.EventSubject_EVENT_SUBJECT_ARTICLE_WATCHED:
		return notifyenum.NotificationTypeContentArticleWatch, true
	case event.EventSubject_EVENT_SUBJECT_ARTICLE_ACCEPTED_ANSWER:
		return notifyenum.NotificationTypeContentArticleAcceptAnswer, true
	case event.EventSubject_EVENT_SUBJECT_ARTICLE_LOTTERY_WON:
		return notifyenum.NotificationTypeContentArticleLotteryWon, true
	case event.EventSubject_EVENT_SUBJECT_COMMENT_PUBLISHED:
		return notifyenum.NotificationTypeContentCommentPublish, true
	case event.EventSubject_EVENT_SUBJECT_COMMENT_LIKED:
		return notifyenum.NotificationTypeContentCommentLike, true
	case event.EventSubject_EVENT_SUBJECT_COMMENT_THANKED:
		return notifyenum.NotificationTypeContentCommentThank, true
	default:
		return "", false
	}
}
