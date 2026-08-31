package usecase

import (
	"common/pkg/constant"
	"context"
	"fmt"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"log/slog"
)

// CharacterStateUsecase 负责角色本地热状态与数据库快照之间的刷盘和清理。
type CharacterStateUsecase struct {
	logger          *slog.Logger
	stateRepo       repo.CharacterStateRepo
	actionQueueRepo repo.CharacterActionQueueRepo
	backpackRepo    repo.CharacterBackpackRepo
	abilityRepo     repo.CharacterAbilityRepo
	commandHandlers map[enum.StateCommandType]CharacterStateCommandHandler
}

func NewCharacterStateUsecase(
	logger *slog.Logger,
	stateRepo repo.CharacterStateRepo,
	actionQueueRepo repo.CharacterActionQueueRepo,
	backpackRepo repo.CharacterBackpackRepo,
	abilityRepo repo.CharacterAbilityRepo,
	commandHandlers map[enum.StateCommandType]CharacterStateCommandHandler,
) *CharacterStateUsecase {
	return &CharacterStateUsecase{
		logger:          logger,
		stateRepo:       stateRepo,
		actionQueueRepo: actionQueueRepo,
		backpackRepo:    backpackRepo,
		abilityRepo:     abilityRepo,
		commandHandlers: commandHandlers,
	}
}

// ExecuteCommand 在角色串行上下文内执行命令，并允许调用方同步处理调度副作用。
func (u *CharacterStateUsecase) ExecuteCommand(
	ctx context.Context,
	command CharacterStateCommand,
	after func(context.Context, *CharacterStateChangeSet) error,
) (*CharacterStateChangeSet, error) {
	var changeSet *CharacterStateChangeSet
	err := u.Execute(ctx, command.CharacterStateID(), func(runCtx context.Context) error {
		var err error
		changeSet, err = u.apply(runCtx, command)
		if err != nil || after == nil {
			return err
		}
		return after(runCtx, changeSet)
	})
	return changeSet, err
}

// Execute 将同一角色的状态操作交给底层内存管理者排队执行。
func (u *CharacterStateUsecase) Execute(ctx context.Context, characterID int64, operation repo.CharacterStateOperation) error {
	return u.stateRepo.Execute(ctx, characterID, operation)
}

func (u *CharacterStateUsecase) apply(ctx context.Context, command CharacterStateCommand) (*CharacterStateChangeSet, error) {
	commandHandler := u.commandHandlers[command.StateCommandType()]
	if commandHandler == nil {
		return nil, fmt.Errorf("game idle character state command handler is required: %s", command.StateCommandType())
	}
	return commandHandler.Apply(ctx, command)
}

// Persist 将角色当前本地热状态写入数据库快照。
func (u *CharacterStateUsecase) Persist(ctx context.Context, characterID int64) error {
	return u.Execute(ctx, characterID, func(runCtx context.Context) error {
		return u.persist(runCtx, characterID)
	})
}

func (u *CharacterStateUsecase) persist(ctx context.Context, characterID int64) error {
	if err := u.actionQueueRepo.Persist(ctx, characterID); err != nil {
		return err
	}
	if err := u.backpackRepo.PersistItems(ctx, characterID); err != nil {
		return err
	}
	return u.abilityRepo.Persist(ctx, characterID)
}

// FlushAndClear 先刷盘再清理缓存；失败只记录日志，保留本地数据等待后续重试。
func (u *CharacterStateUsecase) FlushAndClear(ctx context.Context, characterID int64) {
	_ = u.Execute(ctx, characterID, func(runCtx context.Context) error {
		u.flushAndClear(runCtx, characterID)
		return nil
	})
}

func (u *CharacterStateUsecase) flushAndClear(ctx context.Context, characterID int64) {
	if err := u.persist(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle persist character state failed", constant.LogFieldErr, err, "character_id", characterID)
		return
	}
	if err := u.stateRepo.Clear(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle clear character state cache failed", constant.LogFieldErr, err, "character_id", characterID)
	}
}
