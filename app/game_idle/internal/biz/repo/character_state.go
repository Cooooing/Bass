package repo

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/enum"
)

// CharacterStateRepo 管理单角色运行时状态的串行执行和本地事务。
type CharacterStateRepo interface {
	Execute(ctx context.Context, characterID int64, operation CharacterStateOperation) error
	Begin(ctx context.Context, characterID int64) (CharacterStateTx, error)
	Clear(ctx context.Context, characterID int64) error
}

// CharacterStateOperation 是需要按角色串行执行的状态操作，同一角色不会并发进入该函数。
type CharacterStateOperation func(ctx context.Context) error

// CharacterStateTx 表示一次单角色状态事务，业务层修改 Draft 后显式提交或回滚。
type CharacterStateTx interface {
	Draft() *CharacterStateDraft
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

// CharacterStateDraft 是角色状态事务草稿。
type CharacterStateDraft struct {
	CharacterID int64
	// Queue 是行动队列草稿，结算和队列操作都在草稿上调整队首和剩余次数。
	Queue *model.CharacterActionQueue
	// BackpackItems 是背包草稿，key 为 itemID，value 为角色持有数量和累计变化。
	BackpackItems map[string]*model.CharacterItem
	// BackpackOperationCount 记录背包热状态变更次数，用于低频批量刷库。
	BackpackOperationCount int64
	// Abilities 是能力草稿，key 为能力枚举，value 为等级、经验和下级经验。
	Abilities map[enum.Ability]*model.CharacterAbility
	// AbilityOperationCount 记录能力热状态变更次数，用于低频批量刷库。
	AbilityOperationCount int64
}
