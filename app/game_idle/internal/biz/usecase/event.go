package usecase

import (
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	commonevent "common/proto/gen/common/enums/event/v1"
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GameIdleEventUsecase 提供类型化事件发布入口。
type GameIdleEventUsecase struct {
	logger            *slog.Logger
	gameIdleEventRepo repo.GameIdleEventRepo
	characterIDField  string
}

func NewGameIdleEventUsecase(logger *slog.Logger, gameIdleEventRepo repo.GameIdleEventRepo) *GameIdleEventUsecase {
	return &GameIdleEventUsecase{
		logger:            logger,
		gameIdleEventRepo: gameIdleEventRepo,
		characterIDField:  "character_id",
	}
}

func (u *GameIdleEventUsecase) PublishCloseSession(ctx context.Context, event *model.CharacterCloseSessionEvent) error {
	reason := commonevent.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_UNSPECIFIED
	switch event.Reason {
	case enum.CharacterCloseSessionReasonOccupied:
		reason = commonevent.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_OCCUPIED
	case enum.CharacterCloseSessionReasonTimeout:
		reason = commonevent.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_TIMEOUT
	}
	return u.publish(ctx, &commonevent.Event{
		Payload: &commonevent.Event_GameIdleCloseSession{
			GameIdleCloseSession: &commonevent.GameIdleCloseSessionPayload{
				SessionId:       event.SessionID,
				Reason:          reason,
				Message:         event.Message,
				ShouldReconnect: event.ShouldReconnect,
			},
		},
	}, map[string]string{
		"session_id": event.SessionID,
	})
}

func (u *GameIdleEventUsecase) PublishChatMessage(ctx context.Context, message *model.ChatMessage) error {
	createdAt := time.Now()
	if message.CreatedAt != nil {
		createdAt = *message.CreatedAt
	}
	receiverCharacterID := int64(0)
	if message.ReceiverCharacterID != nil {
		receiverCharacterID = *message.ReceiverCharacterID
	}
	return u.publish(ctx, &commonevent.Event{
		Timestamp: timestamppb.New(createdAt),
		Payload: &commonevent.Event_GameIdleChatMessage{
			GameIdleChatMessage: &commonevent.GameIdleChatMessagePayload{
				MessageId:           message.ID,
				ChannelType:         message.ChannelType.String(),
				ChannelId:           message.ChannelID,
				SenderCharacterId:   message.SenderCharacterID,
				ReceiverCharacterId: receiverCharacterID,
				Content:             message.Content,
				SenderName:          message.SenderName,
				CreatedAt:           timestamppb.New(createdAt),
			},
		},
	}, map[string]string{
		"message_id": strconv.FormatInt(message.ID, 10),
	})
}

func (u *GameIdleEventUsecase) PublishActionCompleted(ctx context.Context, event *model.ActionCompletedEvent) {
	itemChanges := make([]*commonevent.GameIdleItemChange, 0, len(event.ItemChanges))
	for _, item := range event.ItemChanges {
		itemChanges = append(itemChanges, &commonevent.GameIdleItemChange{
			ItemId:        item.ItemID,
			QuantityDelta: item.QuantityDelta,
			QuantityAfter: item.QuantityAfter,
		})
	}
	abilityChanges := make([]*commonevent.GameIdleAbilityChange, 0, len(event.AbilityChanges))
	for _, ability := range event.AbilityChanges {
		abilityChanges = append(abilityChanges, &commonevent.GameIdleAbilityChange{
			AbilityId: ability.AbilityID,
			ExpDelta:  ability.ExpDelta,
			ExpAfter:  ability.ExpAfter,
		})
	}
	err := u.publish(ctx, &commonevent.Event{
		Timestamp: timestamppb.New(event.CompletedAt),
		Payload: &commonevent.Event_GameIdleActionCompleted{
			GameIdleActionCompleted: &commonevent.GameIdleActionCompletedPayload{
				CharacterId: event.CharacterID,
				Action: &commonevent.GameIdleActionCompletedAction{
					ActionId:       event.ActionID,
					TimesFinished:  event.TimesFinished,
					TimesRemaining: event.TimesRemaining,
					StartedAt:      timestamppb.New(event.StartedAt),
					CompletedAt:    timestamppb.New(event.CompletedAt),
				},
				ItemChanges:    itemChanges,
				AbilityChanges: abilityChanges,
			},
		},
	}, map[string]string{
		u.characterIDField: strconv.FormatInt(event.CharacterID, 10),
	})
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle action completed event publish failed", constant.LogFieldErr, err, u.characterIDField, event.CharacterID)
	}
}

func (u *GameIdleEventUsecase) PublishAbilityLeveledUp(ctx context.Context, event *model.AbilityLeveledUpEvent) {
	if event == nil {
		return
	}
	err := u.publish(ctx, &commonevent.Event{
		Payload: &commonevent.Event_GameIdleAbilityLeveledUp{
			GameIdleAbilityLeveledUp: &commonevent.GameIdleAbilityLeveledUpPayload{
				CharacterId:  event.CharacterID,
				AbilityId:    event.AbilityID,
				Level:        event.Level,
				Exp:          event.Exp,
				NextLevelExp: event.NextLevelExp,
			},
		},
	}, map[string]string{
		u.characterIDField: strconv.FormatInt(event.CharacterID, 10),
	})
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle ability leveled up event publish failed", constant.LogFieldErr, err, u.characterIDField, event.CharacterID)
	}
}

func (u *GameIdleEventUsecase) PublishActionQueueUpdated(ctx context.Context, queue *model.CharacterActionQueue) {
	items := make([]*commonevent.GameIdleActionQueueItem, 0, len(queue.Items))
	for _, item := range queue.Items {
		items = append(items, &commonevent.GameIdleActionQueueItem{
			ActionId:  item.ActionID,
			Times:     item.Times,
			CreatedAt: timestamppb.New(item.CreatedAt),
		})
	}
	updatedAt := time.Now()
	err := u.publish(ctx, &commonevent.Event{
		Timestamp: timestamppb.New(updatedAt),
		Payload: &commonevent.Event_GameIdleActionQueueUpdated{
			GameIdleActionQueueUpdated: &commonevent.GameIdleActionQueueUpdatedPayload{
				CharacterId: queue.CharacterID,
				Items:       items,
				UpdatedAt:   timestamppb.New(updatedAt),
			},
		},
	}, map[string]string{
		u.characterIDField: strconv.FormatInt(queue.CharacterID, 10),
	})
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle action queue updated event publish failed", constant.LogFieldErr, err, u.characterIDField, queue.CharacterID)
	}
}

func (u *GameIdleEventUsecase) publish(ctx context.Context, event *commonevent.Event, header map[string]string) error {
	eventID := uuid.NewString()
	event.EventId = eventID
	subject, err := u.eventSubject(event)
	if err != nil {
		return err
	}
	event.Subject = commonenum.EventSubjectMap.MustToProto(subject)
	if event.GetTimestamp() == nil {
		event.Timestamp = timestamppb.Now()
	}
	if header == nil {
		header = map[string]string{}
	}
	header["event_id"] = eventID
	payload, err := proto.Marshal(event)
	if err != nil {
		return err
	}
	return u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEventMessage{
		Subject: subject.String(),
		Data:    payload,
		Header:  header,
	})
}

func (u *GameIdleEventUsecase) eventSubject(event *commonevent.Event) (commonenum.EventSubject, error) {
	switch event.GetPayload().(type) {
	case *commonevent.Event_GameIdleChatMessage:
		return commonenum.EventSubjectGameIdleChatMessage, nil
	case *commonevent.Event_GameIdleCloseSession:
		return commonenum.EventSubjectGameIdleCloseSession, nil
	case *commonevent.Event_GameIdleActionCompleted:
		return commonenum.EventSubjectGameIdleActionCompleted, nil
	case *commonevent.Event_GameIdleAbilityLeveledUp:
		return commonenum.EventSubjectGameIdleAbilityLeveledUp, nil
	case *commonevent.Event_GameIdleActionQueueUpdated:
		return commonenum.EventSubjectGameIdleActionQueueUpdated, nil
	default:
		return "", fmt.Errorf("unsupported game idle event payload")
	}
}
