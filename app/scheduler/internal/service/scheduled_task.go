package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	schedulerv1 "common/proto/gen/scheduler/v1"
	schedulerv1enum "common/proto/gen/scheduler/v1/enum"
	"context"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/usecase"
	schedulerenum "scheduler/internal/enum"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SchedulerScheduledTaskService struct {
	schedulerv1.UnimplementedSchedulerScheduledTaskServiceServer
	scheduledTaskUsecase *usecase.ScheduledTaskUsecase
}

func NewSchedulerScheduledTaskService(
	scheduledTaskUsecase *usecase.ScheduledTaskUsecase,
) *SchedulerScheduledTaskService {
	return &SchedulerScheduledTaskService{
		scheduledTaskUsecase: scheduledTaskUsecase,
	}
}

func (s *SchedulerScheduledTaskService) RegisterGrpc(gs *grpc.Server) {
	schedulerv1.RegisterSchedulerScheduledTaskServiceServer(gs, s)
}

func (s *SchedulerScheduledTaskService) RegisterHttp(hs *http.Server) {
}

func (s *SchedulerScheduledTaskService) Upsert(ctx context.Context, req *schedulerv1.UpsertSchedulerScheduledTask_Req) (*schedulerv1.UpsertSchedulerScheduledTask_Resp, error) {
	if req == nil || req.GetName() == "" || req.GetTitle() == "" || req.GetCronSpec() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	misfirePolicy, ok := schedulerenum.TaskMisfirePolicyMap.ToEnum(req.GetMisfirePolicy())
	if req.GetMisfirePolicy() != schedulerv1enum.SchedulerTaskMisfirePolicy_SCHEDULER_TASK_MISFIRE_POLICY_UNSPECIFIED && !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var staleAfter *time.Duration
	if req.StaleAfterSeconds != nil {
		staleAfter = new(time.Duration)
		*staleAfter = time.Duration(req.GetStaleAfterSeconds()) * time.Second
	}
	row, err := s.scheduledTaskUsecase.Upsert(ctx, &model.ScheduledTask{
		ID:            req.GetId(),
		Name:          req.GetName(),
		Title:         req.GetTitle(),
		Description:   req.GetDescription(),
		Enabled:       req.GetEnabled(),
		CronSpec:      req.GetCronSpec(),
		Payload:       req.GetPayload(),
		Timeout:       time.Duration(req.GetTimeoutSeconds()) * time.Second,
		StaleAfter:    staleAfter,
		MaxAttempts:   req.GetMaxAttempts(),
		MisfirePolicy: misfirePolicy,
		AllowOverlap:  req.GetAllowOverlap(),
	})
	if err != nil {
		return nil, err
	}
	resp := &schedulerv1.UpsertSchedulerScheduledTask_Resp{
		Row: &schedulerv1.UpsertSchedulerScheduledTask_Resp_ScheduledTask{
			Id:             row.ID,
			Name:           row.Name,
			Title:          row.Title,
			Description:    row.Description,
			Enabled:        row.Enabled,
			CronSpec:       row.CronSpec,
			Payload:        row.Payload,
			TimeoutSeconds: int32(row.Timeout / time.Second),
			MaxAttempts:    row.MaxAttempts,
			AllowOverlap:   row.AllowOverlap,
			Version:        row.Version,
			MisfirePolicy:  schedulerenum.TaskMisfirePolicyMap.MustToProto(row.MisfirePolicy),
		},
	}
	if row.CreatedAt != nil {
		resp.Row.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		resp.Row.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return resp, nil
}

func (s *SchedulerScheduledTaskService) Get(ctx context.Context, req *schedulerv1.GetSchedulerScheduledTask_Req) (*schedulerv1.GetSchedulerScheduledTask_Resp, error) {
	row, err := s.scheduledTaskUsecase.Get(ctx, &usecase.ScheduledTaskGetReq{ID: req.GetId()})
	if err != nil {
		return nil, err
	}
	resp := &schedulerv1.GetSchedulerScheduledTask_Resp{
		Row: &schedulerv1.GetSchedulerScheduledTask_Resp_ScheduledTask{
			Id:             row.ID,
			Name:           row.Name,
			Title:          row.Title,
			Description:    row.Description,
			Enabled:        row.Enabled,
			CronSpec:       row.CronSpec,
			Payload:        row.Payload,
			TimeoutSeconds: int32(row.Timeout / time.Second),
			MaxAttempts:    row.MaxAttempts,
			AllowOverlap:   row.AllowOverlap,
			Version:        row.Version,
			MisfirePolicy:  schedulerenum.TaskMisfirePolicyMap.MustToProto(row.MisfirePolicy),
		},
	}
	if row.CreatedAt != nil {
		resp.Row.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		resp.Row.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return resp, nil
}

func (s *SchedulerScheduledTaskService) Page(ctx context.Context, req *schedulerv1.PageSchedulerScheduledTasks_Req) (*schedulerv1.PageSchedulerScheduledTasks_Resp, error) {
	query := &usecase.ScheduledTaskPageReq{}
	if req.GetQuery() != nil {
		query.IDs = req.GetQuery().GetIds()
		query.Name = req.GetQuery().Name
		query.Title = req.GetQuery().Title
		query.Enabled = req.GetQuery().Enabled
	}
	query.Page = req.GetPage()
	pageResp, err := s.scheduledTaskUsecase.Page(ctx, query)
	if err != nil {
		return nil, err
	}
	rows := make([]*schedulerv1.PageSchedulerScheduledTasks_Resp_ScheduledTask, 0, len(pageResp.Rows))
	for _, row := range pageResp.Rows {
		item := &schedulerv1.PageSchedulerScheduledTasks_Resp_ScheduledTask{
			Id:             row.ID,
			Name:           row.Name,
			Title:          row.Title,
			Description:    row.Description,
			Enabled:        row.Enabled,
			CronSpec:       row.CronSpec,
			Payload:        row.Payload,
			TimeoutSeconds: int32(row.Timeout / time.Second),
			MaxAttempts:    row.MaxAttempts,
			AllowOverlap:   row.AllowOverlap,
			Version:        row.Version,
			MisfirePolicy:  schedulerenum.TaskMisfirePolicyMap.MustToProto(row.MisfirePolicy),
		}
		if row.CreatedAt != nil {
			item.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		rows = append(rows, item)
	}
	return &schedulerv1.PageSchedulerScheduledTasks_Resp{
		Page: pageResp.Page,
		Rows: rows,
	}, nil
}

func (s *SchedulerScheduledTaskService) ListAvailableTasks(
	ctx context.Context,
	req *schedulerv1.ListSchedulerAvailableScheduledTasks_Req,
) (*schedulerv1.ListSchedulerAvailableScheduledTasks_Resp, error) {
	keyword := ""
	if req.GetQuery() != nil {
		keyword = req.GetQuery().GetKeyword()
	}
	rows, err := s.scheduledTaskUsecase.ListAvailableTasks(ctx, keyword)
	if err != nil {
		return nil, err
	}
	replyRows := make([]*schedulerv1.ListSchedulerAvailableScheduledTasks_Resp_AvailableTask, 0, len(rows))
	for _, row := range rows {
		replyRows = append(replyRows, &schedulerv1.ListSchedulerAvailableScheduledTasks_Resp_AvailableTask{
			Name:        row.Name,
			Title:       row.Title,
			Description: row.Description,
		})
	}
	return &schedulerv1.ListSchedulerAvailableScheduledTasks_Resp{
		Rows: replyRows,
	}, nil
}

func (s *SchedulerScheduledTaskService) PageExecutionRecords(
	ctx context.Context,
	req *schedulerv1.PageSchedulerScheduledTaskExecutionRecords_Req,
) (*schedulerv1.PageSchedulerScheduledTaskExecutionRecords_Resp, error) {
	query := &usecase.ScheduledTaskExecutionRecordPageReq{}
	if req.GetQuery() != nil {
		query.IDs = req.GetQuery().GetIds()
		query.ScheduledTaskID = req.GetQuery().ScheduledTaskId
		if req.GetQuery().Status != nil && req.GetQuery().GetStatus() != schedulerv1enum.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_UNSPECIFIED {
			statusValue, ok := schedulerenum.TaskExecutionStatusMap.ToEnum(req.GetQuery().GetStatus())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.Status = &statusValue
		}
		if req.GetQuery().TriggerType != nil && req.GetQuery().GetTriggerType() != schedulerv1enum.SchedulerTaskTriggerType_SCHEDULER_TASK_TRIGGER_TYPE_UNSPECIFIED {
			triggerType, ok := schedulerenum.TaskTriggerTypeMap.ToEnum(req.GetQuery().GetTriggerType())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.TriggerType = &triggerType
		}
	}
	query.Page = req.GetPage()
	pageResp, err := s.scheduledTaskUsecase.PageExecutionRecords(ctx, query)
	if err != nil {
		return nil, err
	}
	rows := make([]*schedulerv1.PageSchedulerScheduledTaskExecutionRecords_Resp_ExecutionRecord, 0, len(pageResp.Rows))
	for _, row := range pageResp.Rows {
		item := &schedulerv1.PageSchedulerScheduledTaskExecutionRecords_Resp_ExecutionRecord{
			Id:                   row.ID,
			ScheduledTaskId:      row.ScheduledTaskID,
			ScheduledTaskVersion: row.ScheduledTaskVersion,
			TriggerType:          schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType),
			ScheduleKey:          row.ScheduleKey,
			ScheduledAt:          timestamppb.New(row.ScheduledAt),
			Status:               schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status),
			Attempt:              row.Attempt,
			MaxAttempts:          row.MaxAttempts,
			TimeoutSeconds:       int32(row.Timeout / time.Second),
			WorkerId:             row.WorkerID,
			Payload:              row.Payload,
			LastError:            row.LastError,
			TraceId:              row.TraceID,
		}
		if row.StartedAt != nil {
			item.StartedAt = timestamppb.New(*row.StartedAt)
		}
		if row.FinishedAt != nil {
			item.FinishedAt = timestamppb.New(*row.FinishedAt)
		}
		if row.Duration != nil {
			item.DurationMs = new(row.Duration.Milliseconds())
		}
		if row.CreatedAt != nil {
			item.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		rows = append(rows, item)
	}
	return &schedulerv1.PageSchedulerScheduledTaskExecutionRecords_Resp{
		Page: pageResp.Page,
		Rows: rows,
	}, nil
}

func (s *SchedulerScheduledTaskService) Trigger(ctx context.Context, req *schedulerv1.TriggerSchedulerScheduledTask_Req) (*schedulerv1.TriggerSchedulerScheduledTask_Resp, error) {
	row, err := s.scheduledTaskUsecase.Trigger(ctx, &usecase.TaskTriggerReq{
		ID:      req.GetId(),
		Payload: req.GetPayload(),
	})
	if err != nil {
		return nil, err
	}
	item := &schedulerv1.TriggerSchedulerScheduledTask_Resp_ExecutionRecord{
		Id:                   row.ID,
		ScheduledTaskId:      row.ScheduledTaskID,
		ScheduledTaskVersion: row.ScheduledTaskVersion,
		TriggerType:          schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType),
		ScheduleKey:          row.ScheduleKey,
		ScheduledAt:          timestamppb.New(row.ScheduledAt),
		Status:               schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status),
		Attempt:              row.Attempt,
		MaxAttempts:          row.MaxAttempts,
		TimeoutSeconds:       int32(row.Timeout / time.Second),
		WorkerId:             row.WorkerID,
		Payload:              row.Payload,
		LastError:            row.LastError,
		TraceId:              row.TraceID,
	}
	if row.StartedAt != nil {
		item.StartedAt = timestamppb.New(*row.StartedAt)
	}
	if row.FinishedAt != nil {
		item.FinishedAt = timestamppb.New(*row.FinishedAt)
	}
	if row.Duration != nil {
		item.DurationMs = new(row.Duration.Milliseconds())
	}
	if row.CreatedAt != nil {
		item.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.TriggerSchedulerScheduledTask_Resp{
		Row: item,
	}, nil
}

func (s *SchedulerScheduledTaskService) CancelExecution(
	ctx context.Context,
	req *schedulerv1.CancelSchedulerScheduledTaskExecution_Req,
) (*schedulerv1.CancelSchedulerScheduledTaskExecution_Resp, error) {
	row, err := s.scheduledTaskUsecase.CancelExecution(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	item := &schedulerv1.CancelSchedulerScheduledTaskExecution_Resp_ExecutionRecord{
		Id:                   row.ID,
		ScheduledTaskId:      row.ScheduledTaskID,
		ScheduledTaskVersion: row.ScheduledTaskVersion,
		TriggerType:          schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType),
		ScheduleKey:          row.ScheduleKey,
		ScheduledAt:          timestamppb.New(row.ScheduledAt),
		Status:               schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status),
		Attempt:              row.Attempt,
		MaxAttempts:          row.MaxAttempts,
		TimeoutSeconds:       int32(row.Timeout / time.Second),
		WorkerId:             row.WorkerID,
		Payload:              row.Payload,
		LastError:            row.LastError,
		TraceId:              row.TraceID,
	}
	return &schedulerv1.CancelSchedulerScheduledTaskExecution_Resp{
		Row: item,
	}, nil
}
