package usecase

import (
	"common/pkg/client/timewheel"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/enum"
	"time"
)

// CharacterActionTaskBuilder 按行动类型构建时间轮任务；任务只负责结算，不直接修改队列。
type CharacterActionTaskBuilder interface {
	BuildTask(ctx context.Context, req *BuildCharacterActionTaskReq) (*timewheel.Task, error)
}

// BuildCharacterActionTaskReq 是调度器传给具体行动任务的运行时上下文。
type BuildCharacterActionTaskReq struct {
	CharacterID  int64
	QueueItem    *model.CharacterActionQueueItem
	Action       *model.MetaAction
	Now          time.Time
	PendingTasks chan<- *PendingCharacterActionTask
}

// PendingCharacterActionTask 表示一次时间轮任务执行后的队列推进请求。
type PendingCharacterActionTask struct {
	CharacterID int64
	TaskID      string
	ActionID    string
	StopReason  enum.ActionStopReason
	Items       []*model.CharacterBackpackItemChange
	AbilityID   enum.Ability
	ExpReward   int64
	StartedAt   time.Time
	CompletedAt time.Time
}

type OfflineCharacterActionTask struct {
	CharacterID  int64
	LastLogoutAt time.Time
	Now          time.Time
}
