package usecase

import (
	"context"
	"fmt"
	"game_idle/internal/enum"
)

// StateCommand 是进入角色状态机的业务命令。
type StateCommand interface {
	StateCommandType() enum.StateCommandType
}

// StateCommandHandler 同步完成核心状态变更，必须在返回前保证状态已经落到 Redis 热状态。
type StateCommandHandler interface {
	Apply(ctx context.Context, command StateCommand) (*StateChangeSet, error)
}

// StateEngine 是 game_idle 的状态机入口，负责分发命令并展开领域事件。
type StateEngine struct {
	commandHandlers  map[enum.StateCommandType]StateCommandHandler
	handlers         map[enum.StateEventType][]StateEventHandler
	maxDispatchCount int
}

func NewStateEngine(
	commandHandlers map[enum.StateCommandType]StateCommandHandler,
	handlers map[enum.StateEventType][]StateEventHandler,
) *StateEngine {
	return &StateEngine{
		commandHandlers:  commandHandlers,
		handlers:         handlers,
		maxDispatchCount: 1024,
	}
}

func (e *StateEngine) Apply(ctx context.Context, command StateCommand) (*StateChangeSet, error) {
	commandHandler := e.commandHandlers[command.StateCommandType()]
	if commandHandler == nil {
		return nil, fmt.Errorf("game idle state command handler is required: %s", command.StateCommandType())
	}
	changeSet, err := commandHandler.Apply(ctx, command)
	if err != nil || changeSet == nil {
		return changeSet, err
	}
	events, err := e.dispatch(ctx, changeSet.Events)
	if err != nil {
		return changeSet, err
	}
	changeSet.Events = events
	return changeSet, nil
}

// dispatch 按队列展开事件级联，避免 handler 直接递归造成不可控调用树。
func (e *StateEngine) dispatch(ctx context.Context, initialEvents []*StateEvent) ([]*StateEvent, error) {
	events := append([]*StateEvent(nil), initialEvents...)
	for index := 0; index < len(events); index++ {
		if len(events) > e.maxDispatchCount {
			return events, fmt.Errorf("game idle state event dispatch overflow")
		}
		for _, handler := range e.handlers[events[index].Type] {
			nextEvents, err := handler.Handle(ctx, events[index])
			if err != nil {
				return events, err
			}
			events = append(events, nextEvents...)
		}
	}
	return events, nil
}
