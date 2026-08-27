package data

import (
	"common/pkg/client"
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
	"time"

	"google.golang.org/protobuf/proto"
)

var _ repo.WebSocketEventRepo = (*WebSocketEventRepo)(nil)

type WebSocketEventRepo struct {
	natsClient *client.NatsClient
}

type WebSocketEventSubscription struct {
	subscriptions []client.Unsubscriber
}

func NewWebSocketEventRepo(natsClient *client.NatsClient) repo.WebSocketEventRepo {
	return &WebSocketEventRepo{
		natsClient: natsClient,
	}
}

func (r *WebSocketEventRepo) Subscribe(ctx context.Context, handler repo.WebSocketEventHandler) (repo.WebSocketEventSubscription, error) {
	subscription, err := r.natsClient.Subscribe(ctx, commonenum.EventSubjectGameIdle.String(), func(ctx context.Context, msg *client.Message) error {
		event := &commonenums.Event{}
		if err := proto.Unmarshal(msg.Data, event); err != nil {
			return err
		}
		switch event.GetType() {
		case commonenums.EventType_EVENT_TYPE_GAME_IDLE_CHAT_MESSAGE:
			payload := event.GetGameIdleChatMessage()
			return handler(ctx, &model.WebSocketEvent{
				Type: commonenum.EventTypeGameIdleChatMessage,
				ChatMessage: &model.WebSocketChatMessage{
					MessageID:           payload.GetMessageId(),
					ChannelType:         payload.GetChannelType(),
					ChannelID:           payload.GetChannelId(),
					SenderCharacterID:   payload.GetSenderCharacterId(),
					ReceiverCharacterID: payload.GetReceiverCharacterId(),
					Content:             payload.GetContent(),
				},
			})
		case commonenums.EventType_EVENT_TYPE_GAME_IDLE_CLOSE_SESSION:
			payload := event.GetGameIdleCloseSession()
			reason := "unknown"
			switch payload.GetReason() {
			case commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_OCCUPIED:
				reason = "occupied"
			case commonenums.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_TIMEOUT:
				reason = "timeout"
			}
			return handler(ctx, &model.WebSocketEvent{
				Type: commonenum.EventTypeGameIdleCloseSession,
				CloseSession: &model.WebSocketCloseSession{
					SessionID:       payload.GetSessionId(),
					Reason:          reason,
					Message:         payload.GetMessage(),
					ShouldReconnect: payload.GetShouldReconnect(),
				},
			})
		case commonenums.EventType_EVENT_TYPE_GAME_IDLE_ACTION_COMPLETED:
			payload := event.GetGameIdleActionCompleted()
			action := payload.GetAction()
			var startedAt *time.Time
			if action.GetStartedAt() != nil {
				startedAt = new(action.GetStartedAt().AsTime())
			}
			var completedAt *time.Time
			if action.GetCompletedAt() != nil {
				completedAt = new(action.GetCompletedAt().AsTime())
			}
			itemChanges := make([]*model.WebSocketItemChange, 0, len(payload.GetItemChanges()))
			for _, item := range payload.GetItemChanges() {
				itemChanges = append(itemChanges, &model.WebSocketItemChange{
					ItemID:        item.GetItemId(),
					QuantityDelta: item.GetQuantityDelta(),
					QuantityAfter: item.GetQuantityAfter(),
				})
			}
			abilityChanges := make([]*model.WebSocketAbilityChange, 0, len(payload.GetAbilityChanges()))
			for _, ability := range payload.GetAbilityChanges() {
				abilityChanges = append(abilityChanges, &model.WebSocketAbilityChange{
					AbilityID: ability.GetAbilityId(),
					ExpDelta:  ability.GetExpDelta(),
					ExpAfter:  ability.GetExpAfter(),
				})
			}
			return handler(ctx, &model.WebSocketEvent{
				Type: commonenum.EventTypeGameIdleActionCompleted,
				ActionCompleted: &model.WebSocketActionCompleted{
					CharacterID: payload.GetCharacterId(),
					Action: &model.WebSocketCompletedAction{
						ActionID:       action.GetActionId(),
						TimesFinished:  action.GetTimesFinished(),
						TimesRemaining: action.GetTimesRemaining(),
						StartedAt:      startedAt,
						CompletedAt:    completedAt,
					},
					ItemChanges:    itemChanges,
					AbilityChanges: abilityChanges,
				},
			})
		case commonenums.EventType_EVENT_TYPE_GAME_IDLE_ABILITY_LEVELED_UP:
			payload := event.GetGameIdleAbilityLeveledUp()
			return handler(ctx, &model.WebSocketEvent{
				Type: commonenum.EventTypeGameIdleAbilityLeveledUp,
				AbilityLeveledUp: &model.WebSocketAbilityLeveledUp{
					CharacterID:  payload.GetCharacterId(),
					AbilityID:    payload.GetAbilityId(),
					Level:        payload.GetLevel(),
					Exp:          payload.GetExp(),
					NextLevelExp: payload.GetNextLevelExp(),
				},
			})
		default:
			return nil
		}
	})
	if err != nil {
		return nil, err
	}
	return &WebSocketEventSubscription{
		subscriptions: []client.Unsubscriber{subscription},
	}, nil
}

func (s *WebSocketEventSubscription) Unsubscribe() error {
	var err error
	for _, subscription := range s.subscriptions {
		if currentErr := subscription.Unsubscribe(); currentErr != nil {
			err = currentErr
		}
	}
	return err
}
