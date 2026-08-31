package usecase

import (
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"log/slog"
	"strconv"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GameIdleEventUsecase 提供类型化事件发布入口。
type GameIdleEventUsecase struct {
	logger            *slog.Logger
	gameIdleEventRepo repo.GameIdleEventRepo
}

func NewGameIdleEventUsecase(logger *slog.Logger, gameIdleEventRepo repo.GameIdleEventRepo) *GameIdleEventUsecase {
	return &GameIdleEventUsecase{
		logger:            logger,
		gameIdleEventRepo: gameIdleEventRepo,
	}
}

func (u *GameIdleEventUsecase) PublishCloseSession(ctx context.Context, event *model.CharacterCloseSessionEvent) error {
	reason := commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_UNSPECIFIED
	switch event.Reason {
	case enum.CharacterCloseSessionReasonOccupied:
		reason = commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_OCCUPIED
	case enum.CharacterCloseSessionReasonTimeout:
		reason = commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_TIMEOUT
	}
	return u.publish(ctx, &commonenums.Event{
		Type: commonenums.EventType_EVENT_TYPE_GAME_IDLE_CLOSE_SESSION,
		Payload: &commonenums.Event_GameIdleCloseSession{
			GameIdleCloseSession: &commonenums.GameIdleCloseSessionPayload{
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
	return u.publish(ctx, &commonenums.Event{
		Type:      commonenums.EventType_EVENT_TYPE_GAME_IDLE_CHAT_MESSAGE,
		Timestamp: timestamppb.New(createdAt),
		Payload: &commonenums.Event_GameIdleChatMessage{
			GameIdleChatMessage: &commonenums.GameIdleChatMessagePayload{
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
	itemChanges := make([]*commonenums.GameIdleItemChange, 0, len(event.ItemChanges))
	for _, item := range event.ItemChanges {
		itemChanges = append(itemChanges, &commonenums.GameIdleItemChange{
			ItemId:        item.ItemID,
			QuantityDelta: item.QuantityDelta,
			QuantityAfter: item.QuantityAfter,
		})
	}
	abilityChanges := make([]*commonenums.GameIdleAbilityChange, 0, len(event.AbilityChanges))
	for _, ability := range event.AbilityChanges {
		abilityChanges = append(abilityChanges, &commonenums.GameIdleAbilityChange{
			AbilityId: ability.AbilityID,
			ExpDelta:  ability.ExpDelta,
			ExpAfter:  ability.ExpAfter,
		})
	}
	err := u.publish(ctx, &commonenums.Event{
		Type:      commonenums.EventType_EVENT_TYPE_GAME_IDLE_ACTION_COMPLETED,
		Timestamp: timestamppb.New(event.CompletedAt),
		Payload: &commonenums.Event_GameIdleActionCompleted{
			GameIdleActionCompleted: &commonenums.GameIdleActionCompletedPayload{
				CharacterId: event.CharacterID,
				Action: &commonenums.GameIdleActionCompletedAction{
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
		"character_id": strconv.FormatInt(event.CharacterID, 10),
	})
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle action completed event publish failed", constant.LogFieldErr, err, "character_id", event.CharacterID)
	}
}

func (u *GameIdleEventUsecase) PublishAbilityLeveledUp(ctx context.Context, event *model.AbilityLeveledUpEvent) {
	if event == nil {
		return
	}
	err := u.publish(ctx, &commonenums.Event{
		Type: commonenums.EventType_EVENT_TYPE_GAME_IDLE_ABILITY_LEVELED_UP,
		Payload: &commonenums.Event_GameIdleAbilityLeveledUp{
			GameIdleAbilityLeveledUp: &commonenums.GameIdleAbilityLeveledUpPayload{
				CharacterId:  event.CharacterID,
				AbilityId:    event.AbilityID,
				Level:        event.Level,
				Exp:          event.Exp,
				NextLevelExp: event.NextLevelExp,
			},
		},
	}, map[string]string{
		"character_id": strconv.FormatInt(event.CharacterID, 10),
	})
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle ability leveled up event publish failed", constant.LogFieldErr, err, "character_id", event.CharacterID)
	}
}

func (u *GameIdleEventUsecase) PublishActionQueueUpdated(ctx context.Context, queue *model.CharacterActionQueue) {
	items := make([]*commonenums.GameIdleActionQueueItem, 0, len(queue.Items))
	for _, item := range queue.Items {
		items = append(items, &commonenums.GameIdleActionQueueItem{
			ActionId:  item.ActionID,
			Times:     item.Times,
			CreatedAt: timestamppb.New(item.CreatedAt),
		})
	}
	updatedAt := time.Now()
	err := u.publish(ctx, &commonenums.Event{
		Type:      commonenums.EventType_EVENT_TYPE_GAME_IDLE_ACTION_QUEUE_UPDATED,
		Timestamp: timestamppb.New(updatedAt),
		Payload: &commonenums.Event_GameIdleActionQueueUpdated{
			GameIdleActionQueueUpdated: &commonenums.GameIdleActionQueueUpdatedPayload{
				CharacterId: queue.CharacterID,
				Items:       items,
				UpdatedAt:   timestamppb.New(updatedAt),
			},
		},
	}, map[string]string{
		"character_id": strconv.FormatInt(queue.CharacterID, 10),
	})
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle action queue updated event publish failed", constant.LogFieldErr, err, "character_id", queue.CharacterID)
	}
}

func (u *GameIdleEventUsecase) publish(ctx context.Context, event *commonenums.Event, header map[string]string) error {
	eventID := uuid.NewString()
	event.EventId = eventID
	event.Subject = commonenums.EventSubject_EVENT_SUBJECT_GAME_IDLE
	if event.Timestamp == nil {
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
		Subject: commonenum.EventSubjectGameIdle.String(),
		Data:    payload,
		Header:  header,
	})
}
