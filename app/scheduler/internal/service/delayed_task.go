package service

import (
	"common/pkg/apperror"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	schedulerv1 "common/proto/gen/scheduler/v1"
	schedulerv1enum "common/proto/gen/scheduler/v1/enum"
	"context"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/usecase"
	schedulerenum "scheduler/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SchedulerDelayedTaskService struct {
	schedulerv1.UnimplementedSchedulerDelayedTaskServiceServer
	usecase *usecase.DelayedTaskUsecase
}

func NewSchedulerDelayedTaskService(usecase *usecase.DelayedTaskUsecase) *SchedulerDelayedTaskService {
	return &SchedulerDelayedTaskService{usecase: usecase}
}

func (s *SchedulerDelayedTaskService) RegisterGrpc(gs *grpc.Server) {
	schedulerv1.RegisterSchedulerDelayedTaskServiceServer(gs, s)
}
func (s *SchedulerDelayedTaskService) RegisterHttp(hs *http.Server) {}

func (s *SchedulerDelayedTaskService) Register(ctx context.Context, req *schedulerv1.RegisterSchedulerDelayedTask_Req) (*schedulerv1.RegisterSchedulerDelayedTask_Resp, error) {
	row, err := s.usecase.Register(ctx, &usecase.DelayedTaskRegisterReq{IdempotencyKey: req.GetIdempotencyKey(), TaskName: req.GetTaskName(), Payload: req.GetPayload(), ExecuteAt: req.GetExecuteAt().AsTime(), MaxAttempts: req.GetMaxAttempts(), TimeoutSeconds: req.GetTimeoutSeconds()})
	if err != nil {
		return nil, err
	}
	return &schedulerv1.RegisterSchedulerDelayedTask_Resp{Row: registerDelayedTaskToProto(row)}, nil
}

func (s *SchedulerDelayedTaskService) Cancel(ctx context.Context, req *schedulerv1.CancelSchedulerDelayedTask_Req) (*schedulerv1.CancelSchedulerDelayedTask_Resp, error) {
	return &schedulerv1.CancelSchedulerDelayedTask_Resp{}, s.usecase.Cancel(ctx, req.GetId(), req.GetIdempotencyKey())
}

func (s *SchedulerDelayedTaskService) Get(ctx context.Context, req *schedulerv1.GetSchedulerDelayedTask_Req) (*schedulerv1.GetSchedulerDelayedTask_Resp, error) {
	row, err := s.usecase.Get(ctx, req.GetId(), req.GetIdempotencyKey())
	if err != nil {
		return nil, err
	}
	return &schedulerv1.GetSchedulerDelayedTask_Resp{Row: getDelayedTaskToProto(row)}, nil
}

func (s *SchedulerDelayedTaskService) Page(ctx context.Context, req *schedulerv1.PageSchedulerDelayedTasks_Req) (*schedulerv1.PageSchedulerDelayedTasks_Resp, error) {
	req = util.OrDefault(req, &schedulerv1.PageSchedulerDelayedTasks_Req{})
	query := util.OrDefault(req.Query, &schedulerv1.PageSchedulerDelayedTasks_Req_Query{})
	var status *schedulerenum.DelayedTaskStatus
	if query.Status != nil && query.GetStatus() != schedulerv1enum.SchedulerDelayedTaskStatus_SCHEDULER_DELAYED_TASK_STATUS_UNSPECIFIED {
		value, ok := schedulerenum.DelayedTaskStatusMap.ToEnum(query.GetStatus())
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		status = &value
	}
	resp, err := s.usecase.Page(ctx, &usecase.DelayedTaskPageReq{Page: req.GetPage(), TaskName: query.TaskName, Status: status})
	if err != nil {
		return nil, err
	}
	rows := make([]*schedulerv1.PageSchedulerDelayedTasks_Resp_DelayedTask, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		rows = append(rows, pageDelayedTaskToProto(row))
	}
	return &schedulerv1.PageSchedulerDelayedTasks_Resp{Page: resp.Page, Rows: rows}, nil
}

func (s *SchedulerDelayedTaskService) Trigger(ctx context.Context, req *schedulerv1.TriggerSchedulerDelayedTask_Req) (*schedulerv1.TriggerSchedulerDelayedTask_Resp, error) {
	return &schedulerv1.TriggerSchedulerDelayedTask_Resp{}, s.usecase.Trigger(ctx, req.GetId())
}

func registerDelayedTaskToProto(row *model.DelayedTask) *schedulerv1.RegisterSchedulerDelayedTask_Resp_DelayedTask {
	base := delayedTaskBase(row)
	if base == nil {
		return nil
	}
	return &schedulerv1.RegisterSchedulerDelayedTask_Resp_DelayedTask{Id: base.Id, IdempotencyKey: base.IdempotencyKey, TaskName: base.TaskName, Payload: base.Payload, ExecuteAt: base.ExecuteAt, NextRunAt: base.NextRunAt, Status: base.Status, Attempt: base.Attempt, MaxAttempts: base.MaxAttempts, TimeoutSeconds: base.TimeoutSeconds, StartedAt: base.StartedAt, FinishedAt: base.FinishedAt, LastError: base.LastError, CreatedAt: base.CreatedAt, UpdatedAt: base.UpdatedAt}
}
func getDelayedTaskToProto(row *model.DelayedTask) *schedulerv1.GetSchedulerDelayedTask_Resp_DelayedTask {
	base := delayedTaskBase(row)
	if base == nil {
		return nil
	}
	return &schedulerv1.GetSchedulerDelayedTask_Resp_DelayedTask{Id: base.Id, IdempotencyKey: base.IdempotencyKey, TaskName: base.TaskName, Payload: base.Payload, ExecuteAt: base.ExecuteAt, NextRunAt: base.NextRunAt, Status: base.Status, Attempt: base.Attempt, MaxAttempts: base.MaxAttempts, TimeoutSeconds: base.TimeoutSeconds, StartedAt: base.StartedAt, FinishedAt: base.FinishedAt, LastError: base.LastError, CreatedAt: base.CreatedAt, UpdatedAt: base.UpdatedAt}
}
func pageDelayedTaskToProto(row *model.DelayedTask) *schedulerv1.PageSchedulerDelayedTasks_Resp_DelayedTask {
	base := delayedTaskBase(row)
	if base == nil {
		return nil
	}
	return &schedulerv1.PageSchedulerDelayedTasks_Resp_DelayedTask{Id: base.Id, IdempotencyKey: base.IdempotencyKey, TaskName: base.TaskName, Payload: base.Payload, ExecuteAt: base.ExecuteAt, NextRunAt: base.NextRunAt, Status: base.Status, Attempt: base.Attempt, MaxAttempts: base.MaxAttempts, TimeoutSeconds: base.TimeoutSeconds, StartedAt: base.StartedAt, FinishedAt: base.FinishedAt, LastError: base.LastError, CreatedAt: base.CreatedAt, UpdatedAt: base.UpdatedAt}
}

type delayedTaskProtoBase struct {
	Id             int64
	IdempotencyKey string
	TaskName       string
	Payload        string
	ExecuteAt      *timestamppb.Timestamp
	NextRunAt      *timestamppb.Timestamp
	Status         schedulerv1enum.SchedulerDelayedTaskStatus
	Attempt        int32
	MaxAttempts    int32
	TimeoutSeconds int32
	StartedAt      *timestamppb.Timestamp
	FinishedAt     *timestamppb.Timestamp
	LastError      string
	CreatedAt      *timestamppb.Timestamp
	UpdatedAt      *timestamppb.Timestamp
}

func delayedTaskBase(row *model.DelayedTask) *delayedTaskProtoBase {
	if row == nil {
		return nil
	}
	base := &delayedTaskProtoBase{Id: row.ID, IdempotencyKey: row.IdempotencyKey, TaskName: row.TaskName, Payload: row.Payload, ExecuteAt: timestamppb.New(row.ExecuteAt), NextRunAt: timestamppb.New(row.NextRunAt), Status: schedulerenum.DelayedTaskStatusMap.MustToProto(row.Status), Attempt: row.Attempt, MaxAttempts: row.MaxAttempts, TimeoutSeconds: row.TimeoutSeconds, LastError: row.LastError}
	if row.StartedAt != nil {
		base.StartedAt = timestamppb.New(*row.StartedAt)
	}
	if row.FinishedAt != nil {
		base.FinishedAt = timestamppb.New(*row.FinishedAt)
	}
	if row.CreatedAt != nil {
		base.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		base.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return base
}

var _ = common.PageReq{}
