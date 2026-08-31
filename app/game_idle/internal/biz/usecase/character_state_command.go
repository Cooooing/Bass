package usecase

import (
	"context"
	"game_idle/internal/enum"
)

// CharacterStateCommand 是进入角色状态通道的业务命令，所有角色状态变更都应抽象成命令。
type CharacterStateCommand interface {
	// CharacterStateID 返回命令所属角色，底层会按该值串行化执行。
	CharacterStateID() int64
	// StateCommandType 返回命令类型，用于路由到具体处理器。
	StateCommandType() enum.StateCommandType
}

// CharacterStateCommandHandler 同步完成核心状态变更，必须在返回前保证状态已经落到本地热状态。
type CharacterStateCommandHandler interface {
	Apply(ctx context.Context, command CharacterStateCommand) (*CharacterStateChangeSet, error)
}
