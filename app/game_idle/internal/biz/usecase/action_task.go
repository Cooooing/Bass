package usecase

import (
	"common/pkg/client/timewheel"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/enum"
	"time"
)

// ActionTask 按行动类型构建时间轮任务，具体实现放在 usecase/task 包中。
type ActionTask interface {
	BuildTask(ctx context.Context, req *BuildActionTaskReq) (*timewheel.Task, error)
}

type BuildActionTaskReq struct {
	CharacterID  int64
	QueueItem    *model.ActionQueueItem
	Action       *model.Action
	Now          time.Time
	PendingTasks chan<- *PendingActionTask
	OfflineTasks chan<- *OfflineActionTask
}

type PendingActionTask struct {
	CharacterID int64
	TaskID      string
	ActionID    string
	Items       map[string]int64
	StopReason  enum.ActionStopReason
}

type ActionResultEvent struct {
	CharacterID int64            `json:"character_id"`
	ActionID    string           `json:"action_id"`
	Items       map[string]int64 `json:"items"`
}

type OfflineActionTask struct {
	CharacterID  int64
	LastLogoutAt time.Time
	Now          time.Time
}
