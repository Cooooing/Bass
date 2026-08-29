package usecase

import (
	"common/pkg/client/timewheel"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/enum"
	"time"
)

// ActionTask 按行动类型构建时间轮任务；任务只负责结算，不直接修改队列。
type ActionTask interface {
	BuildTask(ctx context.Context, req *BuildActionTaskReq) (*timewheel.Task, error)
}

// BuildActionTaskReq 是调度器传给具体行动任务的运行时上下文。
type BuildActionTaskReq struct {
	CharacterID  int64
	QueueItem    *model.ActionQueueItem
	Action       *model.Action
	Now          time.Time
	PendingTasks chan<- *PendingActionTask
}

// PendingActionTask 表示一次时间轮任务执行后的队列推进请求。
type PendingActionTask struct {
	CharacterID      int64
	TaskID           string
	ActionID         string
	StopReason       enum.ActionStopReason
	StartedAt        time.Time
	CompletedAt      time.Time
	ItemChanges      []*model.ActionCompletedItemChange
	AbilityChanges   []*model.ActionCompletedAbilityChange
	AbilityLeveledUp *model.AbilityLeveledUpEvent
}

type OfflineActionTask struct {
	CharacterID  int64
	LastLogoutAt time.Time
	Now          time.Time
}
