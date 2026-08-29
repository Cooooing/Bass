package usecase

import (
	"common/pkg/constant"
	"context"
	"game_idle/internal/biz/repo"
	"log/slog"
)

// CharacterStateUsecase 负责角色热状态与数据库快照之间的刷盘和清理。
type CharacterStateUsecase struct {
	logger          *slog.Logger
	actionQueueRepo repo.ActionQueueRepo
	backpackRepo    repo.BackpackRepo
	abilityRepo     repo.CharacterAbilityRepo
}

func NewCharacterStateUsecase(
	logger *slog.Logger,
	actionQueueRepo repo.ActionQueueRepo,
	backpackRepo repo.BackpackRepo,
	abilityRepo repo.CharacterAbilityRepo,
) *CharacterStateUsecase {
	return &CharacterStateUsecase{
		logger:          logger,
		actionQueueRepo: actionQueueRepo,
		backpackRepo:    backpackRepo,
		abilityRepo:     abilityRepo,
	}
}

// Persist 将角色当前 Redis 热状态写入数据库快照。
func (u *CharacterStateUsecase) Persist(ctx context.Context, characterID int64) error {
	if err := u.actionQueueRepo.Persist(ctx, characterID); err != nil {
		return err
	}
	if err := u.backpackRepo.PersistItems(ctx, characterID); err != nil {
		return err
	}
	return u.abilityRepo.Persist(ctx, characterID)
}

// ClearCache 清理角色 Redis 热状态，通常只在确认刷盘成功后调用。
func (u *CharacterStateUsecase) ClearCache(ctx context.Context, characterID int64) error {
	if err := u.actionQueueRepo.Clear(ctx, characterID); err != nil {
		return err
	}
	if err := u.backpackRepo.Clear(ctx, characterID); err != nil {
		return err
	}
	return u.abilityRepo.Clear(ctx, characterID)
}

// FlushAndClear 先刷盘再清理缓存；失败只记录日志，保留 Redis 数据等待后续重试。
func (u *CharacterStateUsecase) FlushAndClear(ctx context.Context, characterID int64) {
	if err := u.Persist(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle persist character state failed", constant.LogFieldErr, err, "character_id", characterID)
		return
	}
	if err := u.ClearCache(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle clear character state cache failed", constant.LogFieldErr, err, "character_id", characterID)
	}
}
