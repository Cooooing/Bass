package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"fmt"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"time"
)

// ActionQueueAddCommand 表示向角色行动队列添加一个行动。
type ActionQueueAddCommand struct {
	CharacterID int64
	ActionID    string
	Times       int64
	Capacity    int32
	Position    *int32
	Now         time.Time
}

func (c *ActionQueueAddCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueAdd
}

// ActionQueueMoveCommand 表示调整队列中已有行动的位置。
type ActionQueueMoveCommand struct {
	CharacterID     int64
	CurrentPosition int32
	TargetPosition  int32
}

func (c *ActionQueueMoveCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueMove
}

// ActionQueueRemoveCommand 表示移除队列指定位置的行动。
type ActionQueueRemoveCommand struct {
	CharacterID int64
	Position    int32
}

func (c *ActionQueueRemoveCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueRemove
}

// ActionQueueClearCommand 表示清空角色行动队列。
type ActionQueueClearCommand struct {
	CharacterID int64
}

func (c *ActionQueueClearCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueClear
}

// ActionQueueCompleteHeadCommand 表示队首行动完成一次结算后推进队列。
type ActionQueueCompleteHeadCommand struct {
	CharacterID int64
	TaskID      string
	ActionID    string
	RemoveHead  bool
}

func (c *ActionQueueCompleteHeadCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueCompleteHead
}

// ActionQueueCommandHandler 负责行动队列热状态变更，调用方只根据变更结果调度时间轮。
type ActionQueueCommandHandler struct {
	actionQueueRepo repo.ActionQueueRepo
}

func NewActionQueueCommandHandler(actionQueueRepo repo.ActionQueueRepo) *ActionQueueCommandHandler {
	return &ActionQueueCommandHandler{actionQueueRepo: actionQueueRepo}
}

func (h *ActionQueueCommandHandler) Apply(ctx context.Context, command StateCommand) (*StateChangeSet, error) {
	switch req := command.(type) {
	case *ActionQueueAddCommand:
		return h.add(ctx, req)
	case *ActionQueueMoveCommand:
		return h.move(ctx, req)
	case *ActionQueueRemoveCommand:
		return h.remove(ctx, req)
	case *ActionQueueClearCommand:
		return h.clear(ctx, req)
	case *ActionQueueCompleteHeadCommand:
		return h.completeHead(ctx, req)
	default:
		return nil, fmt.Errorf("game idle action queue command invalid: %s", command.StateCommandType())
	}
}

func (h *ActionQueueCommandHandler) add(ctx context.Context, req *ActionQueueAddCommand) (*StateChangeSet, error) {
	queue, err := h.actionQueueRepo.Load(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	oldHeadID := h.headID(queue)
	if len(queue.Items) >= int(req.Capacity) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_QUEUE_FULL)
	}
	position := len(queue.Items)
	if req.Position != nil {
		position = int(*req.Position)
	}
	if position < 0 || position > len(queue.Items) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}
	now := req.Now
	if now.IsZero() {
		now = time.Now()
	}
	queueItem := &model.ActionQueueItem{
		ID:        fmt.Sprintf("character:%d:action:%s:at:%d", req.CharacterID, req.ActionID, now.UnixNano()),
		ActionID:  req.ActionID,
		Times:     req.Times,
		CreatedAt: now,
	}
	queue.Items = append(queue.Items, nil)
	copy(queue.Items[position+1:], queue.Items[position:])
	queue.Items[position] = queueItem
	return h.save(ctx, queue, oldHeadID)
}

func (h *ActionQueueCommandHandler) move(ctx context.Context, req *ActionQueueMoveCommand) (*StateChangeSet, error) {
	queue, err := h.actionQueueRepo.Load(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	oldHeadID := h.headID(queue)
	currentPosition := int(req.CurrentPosition)
	targetPosition := int(req.TargetPosition)
	if currentPosition < 0 || targetPosition < 0 || currentPosition >= len(queue.Items) || targetPosition >= len(queue.Items) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}
	queueItem := queue.Items[currentPosition]
	queue.Items = append(queue.Items[:currentPosition], queue.Items[currentPosition+1:]...)
	if targetPosition >= len(queue.Items) {
		queue.Items = append(queue.Items, queueItem)
		return h.save(ctx, queue, oldHeadID)
	}
	queue.Items = append(queue.Items, nil)
	copy(queue.Items[targetPosition+1:], queue.Items[targetPosition:])
	queue.Items[targetPosition] = queueItem
	return h.save(ctx, queue, oldHeadID)
}

func (h *ActionQueueCommandHandler) remove(ctx context.Context, req *ActionQueueRemoveCommand) (*StateChangeSet, error) {
	queue, err := h.actionQueueRepo.Load(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	oldHeadID := h.headID(queue)
	if len(queue.Items) == 0 || req.Position < 0 || int(req.Position) >= len(queue.Items) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}
	position := int(req.Position)
	queue.Items = append(queue.Items[:position], queue.Items[position+1:]...)
	return h.save(ctx, queue, oldHeadID)
}

func (h *ActionQueueCommandHandler) clear(ctx context.Context, req *ActionQueueClearCommand) (*StateChangeSet, error) {
	queue, err := h.actionQueueRepo.Load(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	oldHeadID := h.headID(queue)
	queue.Items = make([]*model.ActionQueueItem, 0)
	return h.save(ctx, queue, oldHeadID)
}

func (h *ActionQueueCommandHandler) completeHead(ctx context.Context, req *ActionQueueCompleteHeadCommand) (*StateChangeSet, error) {
	queue, err := h.actionQueueRepo.Load(ctx, req.CharacterID)
	if err != nil {
		return nil, err
	}
	oldHeadID := h.headID(queue)
	if len(queue.Items) == 0 || queue.Items[0].ID != req.TaskID || queue.Items[0].ActionID != req.ActionID {
		return &StateChangeSet{CharacterID: req.CharacterID}, nil
	}
	current := queue.Items[0]
	finishCurrent := req.RemoveHead
	if !finishCurrent && current.Times != -1 {
		current.Times--
		finishCurrent = current.Times <= 0
	}
	timesRemaining := current.Times
	if finishCurrent {
		queue.Items = queue.Items[1:]
		timesRemaining = 0
	}
	if !finishCurrent && current.Times == -1 {
		return &StateChangeSet{
			CharacterID:          req.CharacterID,
			QueueCommandApplied:  true,
			ActionTimesRemaining: -1,
		}, nil
	}
	changeSet, err := h.save(ctx, queue, oldHeadID)
	if err != nil {
		return nil, err
	}
	changeSet.ActionTimesRemaining = timesRemaining
	return changeSet, nil
}

func (h *ActionQueueCommandHandler) save(ctx context.Context, queue *model.ActionQueue, oldHeadID string) (*StateChangeSet, error) {
	newHeadID := h.headID(queue)
	if err := h.actionQueueRepo.Save(ctx, queue); err != nil {
		return nil, err
	}
	changedAt := time.Now()
	changed := &ActionQueueChangedStateEvent{
		CharacterID: queue.CharacterID,
		Queue:       queue,
		OldHeadID:   oldHeadID,
		NewHeadID:   newHeadID,
		ChangedAt:   changedAt,
	}
	return &StateChangeSet{
		CharacterID:         queue.CharacterID,
		QueueCommandApplied: true,
		QueueChanged:        changed,
		Events: []*StateEvent{{
			Type:               enum.StateEventTypeActionQueueChanged,
			ActionQueueChanged: changed,
		}},
	}, nil
}

func (h *ActionQueueCommandHandler) headID(queue *model.ActionQueue) string {
	if len(queue.Items) == 0 {
		return ""
	}
	return queue.Items[0].ID
}
