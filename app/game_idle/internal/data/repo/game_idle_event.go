package repo

import (
	"common/pkg/client"
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"strconv"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ bizrepo.GameIdleEventRepo = (*GameIdleEventRepo)(nil)

// GameIdleEventRepo 发布挂机游戏内部事件。
type GameIdleEventRepo struct {
	natsClient *client.NatsClient
}

func NewGameIdleEventRepo(natsClient *client.NatsClient) bizrepo.GameIdleEventRepo {
	return &GameIdleEventRepo{
		natsClient: natsClient,
	}
}

func (r *GameIdleEventRepo) Publish(ctx context.Context, event *model.GameIdleEvent) error {
	eventID := uuid.NewString()
	message := &client.Message{
		Subject: commonenum.EventSubjectGameIdle.String(),
		Header: map[string]string{
			"event_id": eventID,
		},
	}
	envelope := &commonenums.Event{
		EventId:   eventID,
		Subject:   commonenums.EventSubject_EVENT_SUBJECT_GAME_IDLE,
		Timestamp: timestamppb.New(time.Now()),
	}
	if event.ChatMessage != nil {
		createdAt := time.Now()
		if event.ChatMessage.CreatedAt != nil {
			createdAt = *event.ChatMessage.CreatedAt
		}
		receiverCharacterID := int64(0)
		if event.ChatMessage.ReceiverCharacterID != nil {
			receiverCharacterID = *event.ChatMessage.ReceiverCharacterID
		}
		envelope.Type = commonenums.EventType_EVENT_TYPE_GAME_IDLE_CHAT_MESSAGE
		envelope.Timestamp = timestamppb.New(createdAt)
		envelope.Payload = &commonenums.Event_GameIdleChatMessage{
			GameIdleChatMessage: &commonenums.GameIdleChatMessagePayload{
				MessageId:           event.ChatMessage.ID,
				ChannelType:         event.ChatMessage.ChannelType.String(),
				ChannelId:           event.ChatMessage.ChannelID,
				SenderCharacterId:   event.ChatMessage.SenderCharacterID,
				ReceiverCharacterId: receiverCharacterID,
				Content:             event.ChatMessage.Content,
				SenderName:          event.ChatMessage.SenderName,
				CreatedAt:           timestamppb.New(createdAt),
			},
		}
		message.Header["message_id"] = strconv.FormatInt(event.ChatMessage.ID, 10)
	} else if event.CloseSession != nil {
		reason := commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_UNSPECIFIED
		switch event.CloseSession.Reason {
		case enum.CharacterCloseSessionReasonOccupied:
			reason = commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_OCCUPIED
		case enum.CharacterCloseSessionReasonTimeout:
			reason = commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_TIMEOUT
		}
		envelope.Type = commonenums.EventType_EVENT_TYPE_GAME_IDLE_CLOSE_SESSION
		envelope.Payload = &commonenums.Event_GameIdleCloseSession{
			GameIdleCloseSession: &commonenums.GameIdleCloseSessionPayload{
				SessionId:       event.CloseSession.SessionID,
				Reason:          reason,
				Message:         event.CloseSession.Message,
				ShouldReconnect: event.CloseSession.ShouldReconnect,
			},
		}
		message.Header["session_id"] = event.CloseSession.SessionID
	} else if event.ActionCompleted != nil {
		itemChanges := make([]*commonenums.GameIdleItemChange, 0, len(event.ActionCompleted.ItemChanges))
		for _, item := range event.ActionCompleted.ItemChanges {
			itemChanges = append(itemChanges, &commonenums.GameIdleItemChange{
				ItemId:        item.ItemID,
				QuantityDelta: item.QuantityDelta,
				QuantityAfter: item.QuantityAfter,
			})
		}
		abilityChanges := make([]*commonenums.GameIdleAbilityChange, 0, len(event.ActionCompleted.AbilityChanges))
		for _, ability := range event.ActionCompleted.AbilityChanges {
			abilityChanges = append(abilityChanges, &commonenums.GameIdleAbilityChange{
				AbilityId: ability.AbilityID,
				ExpDelta:  ability.ExpDelta,
				ExpAfter:  ability.ExpAfter,
			})
		}
		envelope.Type = commonenums.EventType_EVENT_TYPE_GAME_IDLE_ACTION_COMPLETED
		envelope.Timestamp = timestamppb.New(event.ActionCompleted.CompletedAt)
		envelope.Payload = &commonenums.Event_GameIdleActionCompleted{
			GameIdleActionCompleted: &commonenums.GameIdleActionCompletedPayload{
				CharacterId: event.ActionCompleted.CharacterID,
				Action: &commonenums.GameIdleActionCompletedAction{
					ActionId:       event.ActionCompleted.ActionID,
					TimesFinished:  event.ActionCompleted.TimesFinished,
					TimesRemaining: event.ActionCompleted.TimesRemaining,
					StartedAt:      timestamppb.New(event.ActionCompleted.StartedAt),
					CompletedAt:    timestamppb.New(event.ActionCompleted.CompletedAt),
				},
				ItemChanges:    itemChanges,
				AbilityChanges: abilityChanges,
			},
		}
		message.Header["character_id"] = strconv.FormatInt(event.ActionCompleted.CharacterID, 10)
	} else if event.AbilityLeveledUp != nil {
		envelope.Type = commonenums.EventType_EVENT_TYPE_GAME_IDLE_ABILITY_LEVELED_UP
		envelope.Payload = &commonenums.Event_GameIdleAbilityLeveledUp{
			GameIdleAbilityLeveledUp: &commonenums.GameIdleAbilityLeveledUpPayload{
				CharacterId:  event.AbilityLeveledUp.CharacterID,
				AbilityId:    event.AbilityLeveledUp.AbilityID,
				Level:        event.AbilityLeveledUp.Level,
				Exp:          event.AbilityLeveledUp.Exp,
				NextLevelExp: event.AbilityLeveledUp.NextLevelExp,
			},
		}
		message.Header["character_id"] = strconv.FormatInt(event.AbilityLeveledUp.CharacterID, 10)
	} else {
		return model.ErrChatMessageInvalid
	}
	payload, err := proto.Marshal(envelope)
	if err != nil {
		return err
	}
	message.Data = payload
	return r.natsClient.Publish(ctx, commonenum.EventSubjectGameIdle.String(), message)
}
