package usecase

import (
	"common/pkg/constant"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"log/slog"
	"time"
)

// GameIdleEventUsecase 提供类型化事件发布入口，业务用例不直接组装统一事件载体。
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
	return u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{CloseSession: event})
}

func (u *GameIdleEventUsecase) PublishChatMessage(ctx context.Context, message *model.ChatMessage) error {
	return u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{ChatMessage: message})
}

func (u *GameIdleEventUsecase) PublishActionCompleted(ctx context.Context, event *model.ActionCompletedEvent) {
	if err := u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{ActionCompleted: event}); err != nil {
		u.logger.ErrorContext(ctx, "game idle action completed event publish failed", constant.LogFieldErr, err, "character_id", event.CharacterID)
	}
}

func (u *GameIdleEventUsecase) PublishAbilityLeveledUp(ctx context.Context, event *model.AbilityLeveledUpEvent) {
	if event == nil {
		return
	}
	if err := u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{AbilityLeveledUp: event}); err != nil {
		u.logger.ErrorContext(ctx, "game idle ability leveled up event publish failed", constant.LogFieldErr, err, "character_id", event.CharacterID)
	}
}

func (u *GameIdleEventUsecase) PublishActionQueueUpdated(ctx context.Context, queue *model.ActionQueue) {
	if err := u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{
		ActionQueueUpdated: &model.ActionQueueUpdatedEvent{
			CharacterID: queue.CharacterID,
			Items:       queue.Items,
			UpdatedAt:   time.Now(),
		},
	}); err != nil {
		u.logger.ErrorContext(ctx, "game idle action queue updated event publish failed", constant.LogFieldErr, err, "character_id", queue.CharacterID)
	}
}
