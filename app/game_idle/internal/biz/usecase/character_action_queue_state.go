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

// CharacterActionQueueAddCommand 表示向角色行动队列添加一个行动。
type CharacterActionQueueAddCommand struct {
	CharacterID int64
	ActionID    string
	Times       int64
	Capacity    int32
	// Position 为空时追加到队尾，非空时插入到指定位置。
	Position *int32
	Now      time.Time
}

func (c *CharacterActionQueueAddCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueAdd
}

func (c *CharacterActionQueueAddCommand) CharacterStateID() int64 {
	return c.CharacterID
}

// CharacterActionQueueMoveCommand 表示调整队列中已有行动的位置。
type CharacterActionQueueMoveCommand struct {
	CharacterID     int64
	CurrentPosition int32
	TargetPosition  int32
}

func (c *CharacterActionQueueMoveCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueMove
}

func (c *CharacterActionQueueMoveCommand) CharacterStateID() int64 {
	return c.CharacterID
}

// CharacterActionQueueRemoveCommand 表示移除队列指定位置的行动。
type CharacterActionQueueRemoveCommand struct {
	CharacterID int64
	Position    int32
}

func (c *CharacterActionQueueRemoveCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueRemove
}

func (c *CharacterActionQueueRemoveCommand) CharacterStateID() int64 {
	return c.CharacterID
}

// CharacterActionQueueClearCommand 表示清空角色行动队列。
type CharacterActionQueueClearCommand struct {
	CharacterID int64
}

func (c *CharacterActionQueueClearCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueClear
}

func (c *CharacterActionQueueClearCommand) CharacterStateID() int64 {
	return c.CharacterID
}

// CharacterActionQueueCompleteHeadCommand 表示队首行动完成一次结算后推进队列。
type CharacterActionQueueCompleteHeadCommand struct {
	CharacterID int64
	// TaskID 是时间轮任务对应的队列项 ID，用于避免旧任务误删新队首。
	TaskID   string
	ActionID string
	// RemoveHead 为 true 时直接移除队首，常用于物品不足或配置不可执行。
	RemoveHead bool
}

func (c *CharacterActionQueueCompleteHeadCommand) StateCommandType() enum.StateCommandType {
	return enum.StateCommandTypeActionQueueCompleteHead
}

func (c *CharacterActionQueueCompleteHeadCommand) CharacterStateID() int64 {
	return c.CharacterID
}

// CharacterActionQueueCommandHandler 负责行动队列热状态变更，调用方只根据变更结果调度时间轮。
type CharacterActionQueueCommandHandler struct {
	actionQueueRepo repo.CharacterActionQueueRepo
}

func NewCharacterActionQueueCommandHandler(actionQueueRepo repo.CharacterActionQueueRepo) *CharacterActionQueueCommandHandler {
	return &CharacterActionQueueCommandHandler{actionQueueRepo: actionQueueRepo}
}

func (h *CharacterActionQueueCommandHandler) Apply(ctx context.Context, command CharacterStateCommand) (*CharacterStateChangeSet, error) {
	queue, err := h.actionQueueRepo.Load(ctx, command.CharacterStateID())
	if err != nil {
		return nil, err
	}
	oldHeadID := ""
	if len(queue.Items) > 0 {
		oldHeadID = queue.Items[0].ID
	}
	timesRemaining := int64(0)
	shouldSetTimesRemaining := false

	switch req := command.(type) {
	case *CharacterActionQueueAddCommand:
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
		queueItem := &model.CharacterActionQueueItem{
			ID:        fmt.Sprintf("character:%d:action:%s:at:%d", req.CharacterID, req.ActionID, now.UnixNano()),
			ActionID:  req.ActionID,
			Times:     req.Times,
			CreatedAt: now,
		}
		queue.Items = append(queue.Items, nil)
		// 插入动作保持队列有序，当前位置及之后的行动整体后移。
		copy(queue.Items[position+1:], queue.Items[position:])
		queue.Items[position] = queueItem
	case *CharacterActionQueueMoveCommand:
		currentPosition := int(req.CurrentPosition)
		targetPosition := int(req.TargetPosition)
		if currentPosition < 0 || targetPosition < 0 || currentPosition >= len(queue.Items) || targetPosition >= len(queue.Items) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
		}
		queueItem := queue.Items[currentPosition]
		queue.Items = append(queue.Items[:currentPosition], queue.Items[currentPosition+1:]...)
		if targetPosition >= len(queue.Items) {
			queue.Items = append(queue.Items, queueItem)
			break
		}
		queue.Items = append(queue.Items, nil)
		// 移动动作先取出原项，再插入目标位置，避免暴露内部队列项 ID 给外部接口。
		copy(queue.Items[targetPosition+1:], queue.Items[targetPosition:])
		queue.Items[targetPosition] = queueItem
	case *CharacterActionQueueRemoveCommand:
		if len(queue.Items) == 0 || req.Position < 0 || int(req.Position) >= len(queue.Items) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
		}
		position := int(req.Position)
		queue.Items = append(queue.Items[:position], queue.Items[position+1:]...)
	case *CharacterActionQueueClearCommand:
		queue.Items = make([]*model.CharacterActionQueueItem, 0)
	case *CharacterActionQueueCompleteHeadCommand:
		if len(queue.Items) == 0 || queue.Items[0].ID != req.TaskID || queue.Items[0].ActionID != req.ActionID {
			return &CharacterStateChangeSet{CharacterID: req.CharacterID}, nil
		}
		current := queue.Items[0]
		finishCurrent := req.RemoveHead
		// 有限次数行动每完成一次扣减一次；无限行动保持队列不变并继续调度下一轮。
		if !finishCurrent && current.Times != -1 {
			current.Times--
			finishCurrent = current.Times <= 0
		}
		timesRemaining = current.Times
		shouldSetTimesRemaining = true
		if finishCurrent {
			queue.Items = queue.Items[1:]
			timesRemaining = 0
		}
		if !finishCurrent && current.Times == -1 {
			return &CharacterStateChangeSet{
				CharacterID:          req.CharacterID,
				QueueCommandApplied:  true,
				ActionTimesRemaining: -1,
			}, nil
		}
	default:
		return nil, fmt.Errorf("game idle action queue command invalid: %s", command.StateCommandType())
	}

	newHeadID := ""
	if len(queue.Items) > 0 {
		newHeadID = queue.Items[0].ID
	}
	if err := h.actionQueueRepo.Save(ctx, queue); err != nil {
		return nil, err
	}
	// 统一生成队列变化事件，调度器据此同步时间轮，BFF 据此推送队列快照。
	changedAt := time.Now()
	changed := &CharacterActionQueueChangedStateEvent{
		CharacterID: queue.CharacterID,
		Queue:       queue,
		OldHeadID:   oldHeadID,
		NewHeadID:   newHeadID,
		ChangedAt:   changedAt,
	}
	changeSet := &CharacterStateChangeSet{
		CharacterID:         queue.CharacterID,
		QueueCommandApplied: true,
		QueueChanged:        changed,
	}
	if shouldSetTimesRemaining {
		changeSet.ActionTimesRemaining = timesRemaining
	}
	return changeSet, nil
}
