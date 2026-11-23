package data

import (
	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"
	commonevent "common/proto/gen/common/enums/event/v1"
	"context"
	"time"

	"bff_game_idle/internal/biz/model"
	"bff_game_idle/internal/biz/repo"

	"google.golang.org/protobuf/proto"
)

var _ repo.WebSocketEventRepo = (*WebSocketEventRepo)(nil)

type WebSocketEventRepo struct {
	natsClient *nats.Client
}

type WebSocketEventSubscription struct {
	subscriptions []*nats.Subscription
}

func NewWebSocketEventRepo(natsClient *nats.Client) repo.WebSocketEventRepo {
	return &WebSocketEventRepo{
		natsClient: natsClient,
	}
}

func (r *WebSocketEventRepo) Subscribe(ctx context.Context, handler repo.WebSocketEventHandler) (repo.WebSocketEventSubscription, error) {
	subscription, err := r.natsClient.Subscribe(ctx, commonenum.EventSubjectGameIdle.String()+".>", func(ctx context.Context, msg *nats.Message) error {
		event := &commonevent.Event{}
		if err := proto.Unmarshal(msg.Data, event); err != nil {
			return err
		}
		switch event.GetSubject() {
		case commonevent.EventSubject_EVENT_SUBJECT_GAME_IDLE_CHAT_MESSAGE:
			payload := event.GetGameIdleChatMessage()
			var createdAt *time.Time
			if payload.GetCreatedAt() != nil {
				createdAt = new(payload.GetCreatedAt().AsTime())
			}
			return handler(ctx, &model.WebSocketEvent{
				ChatMessage: &model.WebSocketChatMessage{
					MessageID:           payload.GetMessageId(),
					ChannelType:         payload.GetChannelType(),
					ChannelID:           payload.GetChannelId(),
					SenderCharacterID:   payload.GetSenderCharacterId(),
					SenderName:          payload.GetSenderName(),
					ReceiverCharacterID: payload.GetReceiverCharacterId(),
					Content:             payload.GetContent(),
					CreatedAt:           createdAt,
				},
			})
		case commonevent.EventSubject_EVENT_SUBJECT_GAME_IDLE_CLOSE_SESSION:
			payload := event.GetGameIdleCloseSession()
			reason := "unknown"
			switch payload.GetReason() {
			case commonevent.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_OCCUPIED:
				reason = "occupied"
			case commonevent.GameIdleCloseSessionReason_GAME_IDLE_CLOSE_SESSION_REASON_TIMEOUT:
				reason = "timeout"
			}
			return handler(ctx, &model.WebSocketEvent{
				CloseSession: &model.WebSocketCloseSession{
					SessionID:       payload.GetSessionId(),
					Reason:          reason,
					Message:         payload.GetMessage(),
					ShouldReconnect: payload.GetShouldReconnect(),
				},
			})
		case commonevent.EventSubject_EVENT_SUBJECT_GAME_IDLE_ACTION_COMPLETED:
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
		case commonevent.EventSubject_EVENT_SUBJECT_GAME_IDLE_ABILITY_LEVELED_UP:
			payload := event.GetGameIdleAbilityLeveledUp()
			return handler(ctx, &model.WebSocketEvent{
				AbilityLeveledUp: &model.WebSocketAbilityLeveledUp{
					CharacterID:  payload.GetCharacterId(),
					AbilityID:    payload.GetAbilityId(),
					Level:        payload.GetLevel(),
					Exp:          payload.GetExp(),
					NextLevelExp: payload.GetNextLevelExp(),
				},
			})
		case commonevent.EventSubject_EVENT_SUBJECT_GAME_IDLE_ACTION_QUEUE_UPDATED:
			payload := event.GetGameIdleActionQueueUpdated()
			items := make([]*model.ActionQueueItem, 0, len(payload.GetItems()))
			for _, item := range payload.GetItems() {
				items = append(items, &model.ActionQueueItem{
					ActionID:  item.GetActionId(),
					Times:     item.GetTimes(),
					CreatedAt: item.GetCreatedAt().AsTime(),
				})
			}
			return handler(ctx, &model.WebSocketEvent{
				ActionQueueUpdated: &model.ActionQueue{
					CharacterID: payload.GetCharacterId(),
					Items:       items,
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
		subscriptions: []*nats.Subscription{subscription},
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
