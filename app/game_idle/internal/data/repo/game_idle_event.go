package repo

import (
	"common/pkg/apperror"
	"common/pkg/client"
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
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

	switch {
	case event.ChatMessage != nil:
		r.encodeChatMessage(envelope, message, event.ChatMessage)
	case event.CloseSession != nil:
		r.encodeCloseSession(envelope, message, event.CloseSession)
	case event.ActionCompleted != nil:
		r.encodeActionCompleted(envelope, message, event.ActionCompleted)
	case event.AbilityLeveledUp != nil:
		r.encodeAbilityLeveledUp(envelope, message, event.AbilityLeveledUp)
	case event.ActionQueueUpdated != nil:
		r.encodeActionQueueUpdated(envelope, message, event.ActionQueueUpdated)
	default:
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHAT_MESSAGE_INVALID)
	}

	payload, err := proto.Marshal(envelope)
	if err != nil {
		return err
	}
	message.Data = payload
	return r.natsClient.Publish(ctx, commonenum.EventSubjectGameIdle.String(), message)
}

func (r *GameIdleEventRepo) encodeChatMessage(
	envelope *commonenums.Event,
	message *client.Message,
	event *model.ChatMessage,
) {
	createdAt := time.Now()
	if event.CreatedAt != nil {
		createdAt = *event.CreatedAt
	}
	receiverCharacterID := int64(0)
	if event.ReceiverCharacterID != nil {
		receiverCharacterID = *event.ReceiverCharacterID
	}
	envelope.Type = commonenums.EventType_EVENT_TYPE_GAME_IDLE_CHAT_MESSAGE
	envelope.Timestamp = timestamppb.New(createdAt)
	envelope.Payload = &commonenums.Event_GameIdleChatMessage{
		GameIdleChatMessage: &commonenums.GameIdleChatMessagePayload{
			MessageId:           event.ID,
			ChannelType:         event.ChannelType.String(),
			ChannelId:           event.ChannelID,
			SenderCharacterId:   event.SenderCharacterID,
			ReceiverCharacterId: receiverCharacterID,
			Content:             event.Content,
			SenderName:          event.SenderName,
			CreatedAt:           timestamppb.New(createdAt),
		},
	}
	message.Header["message_id"] = strconv.FormatInt(event.ID, 10)
}

func (r *GameIdleEventRepo) encodeCloseSession(
	envelope *commonenums.Event,
	message *client.Message,
	event *model.CharacterCloseSessionEvent,
) {
	reason := commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_UNSPECIFIED
	switch event.Reason {
	case enum.CharacterCloseSessionReasonOccupied:
		reason = commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_OCCUPIED
	case enum.CharacterCloseSessionReasonTimeout:
		reason = commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_TIMEOUT
	}
	envelope.Type = commonenums.EventType_EVENT_TYPE_GAME_IDLE_CLOSE_SESSION
	envelope.Payload = &commonenums.Event_GameIdleCloseSession{
		GameIdleCloseSession: &commonenums.GameIdleCloseSessionPayload{
			SessionId:       event.SessionID,
			Reason:          reason,
			Message:         event.Message,
			ShouldReconnect: event.ShouldReconnect,
		},
	}
	message.Header["session_id"] = event.SessionID
}

func (r *GameIdleEventRepo) encodeActionCompleted(
	envelope *commonenums.Event,
	message *client.Message,
	event *model.ActionCompletedEvent,
) {
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
	envelope.Type = commonenums.EventType_EVENT_TYPE_GAME_IDLE_ACTION_COMPLETED
	envelope.Timestamp = timestamppb.New(event.CompletedAt)
	envelope.Payload = &commonenums.Event_GameIdleActionCompleted{
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
	}
	message.Header["character_id"] = strconv.FormatInt(event.CharacterID, 10)
}

func (r *GameIdleEventRepo) encodeAbilityLeveledUp(
	envelope *commonenums.Event,
	message *client.Message,
	event *model.AbilityLeveledUpEvent,
) {
	envelope.Type = commonenums.EventType_EVENT_TYPE_GAME_IDLE_ABILITY_LEVELED_UP
	envelope.Payload = &commonenums.Event_GameIdleAbilityLeveledUp{
		GameIdleAbilityLeveledUp: &commonenums.GameIdleAbilityLeveledUpPayload{
			CharacterId:  event.CharacterID,
			AbilityId:    event.AbilityID,
			Level:        event.Level,
			Exp:          event.Exp,
			NextLevelExp: event.NextLevelExp,
		},
	}
	message.Header["character_id"] = strconv.FormatInt(event.CharacterID, 10)
}

func (r *GameIdleEventRepo) encodeActionQueueUpdated(
	envelope *commonenums.Event,
	message *client.Message,
	event *model.ActionQueueUpdatedEvent,
) {
	items := make([]*commonenums.GameIdleActionQueueItem, 0, len(event.Items))
	for _, item := range event.Items {
		items = append(items, &commonenums.GameIdleActionQueueItem{
			ActionId:  item.ActionID,
			Times:     item.Times,
			CreatedAt: timestamppb.New(item.CreatedAt),
		})
	}
	updatedAt := time.Now()
	if !event.UpdatedAt.IsZero() {
		updatedAt = event.UpdatedAt
	}
	envelope.Type = commonenums.EventType_EVENT_TYPE_GAME_IDLE_ACTION_QUEUE_UPDATED
	envelope.Timestamp = timestamppb.New(updatedAt)
	envelope.Payload = &commonenums.Event_GameIdleActionQueueUpdated{
		GameIdleActionQueueUpdated: &commonenums.GameIdleActionQueueUpdatedPayload{
			CharacterId: event.CharacterID,
			Items:       items,
			Reason:      event.Reason,
			UpdatedAt:   timestamppb.New(updatedAt),
		},
	}
	message.Header["character_id"] = strconv.FormatInt(event.CharacterID, 10)
}
