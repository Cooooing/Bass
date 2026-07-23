package repo

import (
	"context"
	"time"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type AgentJobRepo interface {
	Save(ctx context.Context, job *model.AgentJob) (*model.AgentJob, error)
	MarkRunning(ctx context.Context, jobID int64, startedAt time.Time) (*model.AgentJob, error)
	MarkSucceeded(ctx context.Context, jobID int64, finishedAt time.Time) (*model.AgentJob, error)
	Retry(ctx context.Context, req *AgentJobRetryReq) (*model.AgentJob, error)
	MarkFailed(ctx context.Context, req *AgentJobMarkFailedReq) (*model.AgentJob, error)
	Get(ctx context.Context, req *AgentJobQuery) (*model.AgentJob, error)
	List(ctx context.Context, req *AgentJobQuery) ([]*model.AgentJob, error)
	Map(ctx context.Context, req *AgentJobQuery) (map[int64]*model.AgentJob, error)
	Count(ctx context.Context, req *AgentJobQuery) (int, error)
	Page(ctx context.Context, req *AgentJobPageReq) (*AgentJobPageResp, error)
}

type AgentJobRetryReq struct {
	JobID        int64
	AttemptCount int32
	AvailableAt  time.Time
	ErrorSummary string
}

type AgentJobMarkFailedReq struct {
	JobID        int64
	FinishedAt   time.Time
	ErrorSummary string
}

type AgentJobQuery struct {
	ID              *int64
	IDs             []int64
	WorldID         *int64
	SourceEventID   *int64
	Type            *enum.AgentJobType
	Priority        *enum.AgentJobPriority
	Status          *enum.AgentJobStatus
	Statuses        []enum.AgentJobStatus
	LaneKey         *string
	AvailableBefore *time.Time
	StartedBefore   *time.Time
}

type AgentJobPageReq struct {
	Page  base.PageRequest
	Query AgentJobQuery
}

type AgentJobPageResp struct {
	Rows []*model.AgentJob
	Page base.PageResp
}
