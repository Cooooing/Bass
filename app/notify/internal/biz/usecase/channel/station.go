package channel

import (
	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"
	commonevent "common/proto/gen/common/enums/event"
	"context"
	"time"

	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	notifyenum "notify/internal/enum"

	"log/slog"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type StationUsecase struct {
	log                *slog.Logger
	stationMessageRepo repo.NotificationStationMessageRepo
	publisher          *nats.Client
}

func NewStationUsecase(
	log *slog.Logger,
	stationMessageRepo repo.NotificationStationMessageRepo,
	publisher *nats.Client,
) *StationUsecase {
	return &StationUsecase{
		log:                log,
		stationMessageRepo: stationMessageRepo,
		publisher:          publisher,
	}
}

type StationProcessReq struct {
	NotificationContext *model.NotificationContext
	Rule                *model.NotificationRule
	AccountsByUserID    map[int64]*model.UserAccount
}

func (u *StationUsecase) Process(ctx context.Context, req *StationProcessReq) (notifyenum.NotificationChannelStatus, error) {
	if req == nil || req.NotificationContext == nil || req.Rule == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	notificationContext := req.NotificationContext
	rule := req.Rule
	if rule.StationTemplate == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	seen := make(map[int64]struct{}, len(notificationContext.Recipients))
	written := false
	for _, recipient := range notificationContext.Recipients {
		if recipient == nil || recipient.UserID == 0 {
			continue
		}
		if _, ok := seen[recipient.UserID]; ok {
			continue
		}
		seen[recipient.UserID] = struct{}{}
		rendered, err := rule.StationTemplate.Render(ctx, notificationContext.TemplateData)
		if err != nil {
			continue
		}
		message, err := u.stationMessageRepo.Save(ctx, &model.NotificationStationMessage{
			EventID:          notificationContext.EventID,
			NotificationType: notificationContext.NotificationType,
			ReceiverID:       recipient.UserID,
			BusinessKey:      notificationContext.BusinessKey,
			Title:            rendered.Title,
			Content:          rendered.Content,
			Status:           notifyenum.NotificationChannelStatusSucceeded,
		})
		if err != nil {
			return notifyenum.NotificationChannelStatusInternalError, err
		}
		if message == nil {
			continue
		}
		written = true
		err = u.publishStationMessagesCreated(ctx, notificationContext.EventID, message.ID, []int64{recipient.UserID})
		if err != nil {
			u.log.Warn(
				"publish station message created event failed",
				"event_id", notificationContext.EventID,
				"notification_id", message.ID,
				"receiver_id", recipient.UserID,
				"error", err,
			)
		}
	}
	if !written {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	return notifyenum.NotificationChannelStatusSucceeded, nil
}

func (u *StationUsecase) publishStationMessagesCreated(ctx context.Context, eventID string, notificationID int64, receiverIDs []int64) error {
	if u.publisher == nil || eventID == "" || notificationID <= 0 || len(receiverIDs) == 0 {
		return nil
	}
	event := &commonevent.Event{
		EventId:   eventID,
		Subject:   commonevent.EventSubject_EVENT_SUBJECT_NOTIFY_STATION_MESSAGES_CREATED,
		Timestamp: timestamppb.New(time.Now()),
		Payload: &commonevent.Event_NotifyStationMessagesCreated{
			NotifyStationMessagesCreated: &commonevent.NotifyStationMessagesCreatedPayload{
				NotificationId: notificationID,
				ReceiverIds:    receiverIDs,
			},
		},
	}
	data, err := protojson.MarshalOptions{UseProtoNames: true}.Marshal(event)
	if err != nil {
		return err
	}
	return u.publisher.Publish(ctx, commonenum.EventSubjectNotifyStationMessagesCreated.String(), &nats.Message{Data: data})
}
