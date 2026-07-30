package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	schedulerv1 "common/proto/gen/scheduler/v1"
	schedulerv1enum "common/proto/gen/scheduler/v1/enum"
	"context"
	"encoding/json"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/usecase"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SchedulerDelayedTaskService struct {
	schedulerv1.UnimplementedSchedulerDelayedTaskServiceServer
	usecase *usecase.DelayedTaskUsecase
}

func NewSchedulerDelayedTaskService(
	usecase *usecase.DelayedTaskUsecase,
) *SchedulerDelayedTaskService {
	return &SchedulerDelayedTaskService{
		usecase: usecase,
	}
}

func (s *SchedulerDelayedTaskService) RegisterGrpc(gs *grpc.Server) {
	schedulerv1.RegisterSchedulerDelayedTaskServiceServer(gs, s)
}

func (s *SchedulerDelayedTaskService) RegisterHttp(hs *http.Server) {
}

func (s *SchedulerDelayedTaskService) Upsert(ctx context.Context, req *schedulerv1.UpsertSchedulerDelayedTask_Req) (*schedulerv1.UpsertSchedulerDelayedTask_Resp, error) {
	if req == nil || strings.TrimSpace(req.GetTitle()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	handlerName, ok := schedulerenum.TaskHandlerNameMap.ToEnum(req.GetHandlerName())
	if req.GetHandlerName() == schedulerv1enum.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_UNSPECIFIED || !ok {
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
	row, err := s.usecase.Upsert(ctx, &model.DelayedTask{
		ID:            req.GetId(),
		TaskKey:       req.GetTaskKey(),
		HandlerName:   handlerName,
		Title:         req.GetTitle(),
		Description:   req.GetDescription(),
		Enabled:       req.GetEnabled(),
		Timeout:       time.Duration(req.GetTimeoutSeconds()) * time.Second,
		StaleAfter:    staleAfter,
		MaxAttempts:   req.GetMaxAttempts(),
		MisfirePolicy: misfirePolicy,
	})
	if err != nil {
		return nil, err
	}
	item := &schedulerv1.DelayedTask{
		Id:             row.ID,
		TaskKey:        row.TaskKey,
		HandlerName:    schedulerenum.TaskHandlerNameMap.MustToProto(row.HandlerName),
		Title:          row.Title,
		Description:    row.Description,
		Enabled:        row.Enabled,
		TimeoutSeconds: int32(row.Timeout / time.Second),
		MaxAttempts:    row.MaxAttempts,
		MisfirePolicy:  schedulerenum.TaskMisfirePolicyMap.MustToProto(row.MisfirePolicy),
		Version:        row.Version,
	}
	if row.CreatedAt != nil {
		item.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.UpsertSchedulerDelayedTask_Resp{
		Row: item,
	}, nil
}

func (s *SchedulerDelayedTaskService) Get(ctx context.Context, req *schedulerv1.GetSchedulerDelayedTask_Req) (*schedulerv1.GetSchedulerDelayedTask_Resp, error) {
	if req == nil || req.GetId() == 0 && strings.TrimSpace(req.GetTaskKey()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.usecase.Get(ctx, &usecase.DelayedTaskGetReq{ID: req.GetId(), TaskKey: req.GetTaskKey()})
	if err != nil {
		return nil, err
	}
	item := &schedulerv1.DelayedTask{
		Id:             row.ID,
		TaskKey:        row.TaskKey,
		HandlerName:    schedulerenum.TaskHandlerNameMap.MustToProto(row.HandlerName),
		Title:          row.Title,
		Description:    row.Description,
		Enabled:        row.Enabled,
		TimeoutSeconds: int32(row.Timeout / time.Second),
		MaxAttempts:    row.MaxAttempts,
		MisfirePolicy:  schedulerenum.TaskMisfirePolicyMap.MustToProto(row.MisfirePolicy),
		Version:        row.Version,
	}
	if row.CreatedAt != nil {
		item.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.GetSchedulerDelayedTask_Resp{
		Row: item,
	}, nil
}

func (s *SchedulerDelayedTaskService) Page(ctx context.Context, req *schedulerv1.PageSchedulerDelayedTasks_Req) (*schedulerv1.PageSchedulerDelayedTasks_Resp, error) {
	query := &usecase.DelayedTaskPageReq{}
	if req != nil {
		query.Page = req.GetPage()
		if req.GetQuery() != nil {
			query.IDs = req.GetQuery().GetIds()
			query.TaskKey = req.GetQuery().TaskKey
			if req.GetQuery().HandlerName != nil && req.GetQuery().GetHandlerName() != schedulerv1enum.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_UNSPECIFIED {
				handlerName, ok := schedulerenum.TaskHandlerNameMap.ToEnum(req.GetQuery().GetHandlerName())
				if !ok {
					return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
				}
				query.HandlerName = &handlerName
			}
			query.Title = req.GetQuery().Title
			query.Enabled = req.GetQuery().Enabled
		}
	}
	resp, err := s.usecase.Page(ctx, query)
	if err != nil {
		return nil, err
	}
	rows := make([]*schedulerv1.DelayedTask, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		item := &schedulerv1.DelayedTask{
			Id:             row.ID,
			TaskKey:        row.TaskKey,
			HandlerName:    schedulerenum.TaskHandlerNameMap.MustToProto(row.HandlerName),
			Title:          row.Title,
			Description:    row.Description,
			Enabled:        row.Enabled,
			TimeoutSeconds: int32(row.Timeout / time.Second),
			MaxAttempts:    row.MaxAttempts,
			MisfirePolicy:  schedulerenum.TaskMisfirePolicyMap.MustToProto(row.MisfirePolicy),
			Version:        row.Version,
		}
		if row.CreatedAt != nil {
			item.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		rows = append(rows, item)
	}
	return &schedulerv1.PageSchedulerDelayedTasks_Resp{
		Page: resp.Page,
		Rows: rows,
	}, nil
}

func (s *SchedulerDelayedTaskService) ListAvailableTasks(
	ctx context.Context,
	req *schedulerv1.ListSchedulerAvailableDelayedTasks_Req,
) (*schedulerv1.ListSchedulerAvailableDelayedTasks_Resp, error) {
	keyword := ""
	if req.GetQuery() != nil {
		keyword = req.GetQuery().GetKeyword()
	}
	rows, err := s.usecase.ListAvailableTasks(ctx, keyword)
	if err != nil {
		return nil, err
	}
	replyRows := make([]*schedulerv1.ListSchedulerAvailableDelayedTasks_Resp_AvailableTask, 0, len(rows))
	for _, row := range rows {
		replyRows = append(replyRows, &schedulerv1.ListSchedulerAvailableDelayedTasks_Resp_AvailableTask{
			HandlerName: schedulerenum.TaskHandlerNameMap.MustToProto(row.HandlerName),
			Title:       row.Title,
			Description: row.Description,
		})
	}
	return &schedulerv1.ListSchedulerAvailableDelayedTasks_Resp{
		Rows: replyRows,
	}, nil
}

func (s *SchedulerDelayedTaskService) Schedule(ctx context.Context, req *schedulerv1.ScheduleSchedulerDelayedTask_Req) (*schedulerv1.ScheduleSchedulerDelayedTask_Resp, error) {
	if req == nil || strings.TrimSpace(req.GetTaskKey()) == "" || req.GetScheduledAt() == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if req.GetPayload() != "" && !json.Valid([]byte(req.GetPayload())) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.usecase.Schedule(ctx, &usecase.DelayedTaskScheduleReq{
		TaskKey:     req.GetTaskKey(),
		Payload:     req.GetPayload(),
		ScheduledAt: req.GetScheduledAt().AsTime(),
	})
	if err != nil {
		return nil, err
	}
	item := &schedulerv1.DelayedTaskExecutionRecord{
		Id:                 row.ID,
		DelayedTaskId:      row.DelayedTaskID,
		DelayedTaskVersion: row.DelayedTaskVersion,
		IdempotencyKey:     row.IdempotencyKey,
		TriggerType:        schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType),
		ScheduleKey:        row.ScheduleKey,
		ScheduledAt:        timestamppb.New(row.ScheduledAt),
		Status:             schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status),
		Attempt:            row.Attempt,
		MaxAttempts:        row.MaxAttempts,
		TimeoutSeconds:     int32(row.Timeout / time.Second),
		WorkerId:           row.WorkerID,
		Payload:            row.Payload,
		LastError:          row.LastError,
		TraceId:            row.TraceID,
	}
	if row.CreatedAt != nil {
		item.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.ScheduleSchedulerDelayedTask_Resp{
		Row: item,
	}, nil
}

func (s *SchedulerDelayedTaskService) Trigger(ctx context.Context, req *schedulerv1.TriggerSchedulerDelayedTask_Req) (*schedulerv1.TriggerSchedulerDelayedTask_Resp, error) {
	if req == nil || strings.TrimSpace(req.GetTaskKey()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.usecase.Trigger(ctx, req.GetTaskKey(), req.GetPayload())
	if err != nil {
		return nil, err
	}
	item := &schedulerv1.DelayedTaskExecutionRecord{
		Id:                 row.ID,
		DelayedTaskId:      row.DelayedTaskID,
		DelayedTaskVersion: row.DelayedTaskVersion,
		IdempotencyKey:     row.IdempotencyKey,
		TriggerType:        schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType),
		ScheduleKey:        row.ScheduleKey,
		ScheduledAt:        timestamppb.New(row.ScheduledAt),
		Status:             schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status),
		Attempt:            row.Attempt,
		MaxAttempts:        row.MaxAttempts,
		TimeoutSeconds:     int32(row.Timeout / time.Second),
		WorkerId:           row.WorkerID,
		Payload:            row.Payload,
		LastError:          row.LastError,
		TraceId:            row.TraceID,
	}
	return &schedulerv1.TriggerSchedulerDelayedTask_Resp{
		Row: item,
	}, nil
}

func (s *SchedulerDelayedTaskService) CancelExecution(
	ctx context.Context,
	req *schedulerv1.CancelSchedulerDelayedTaskExecution_Req,
) (*schedulerv1.CancelSchedulerDelayedTaskExecution_Resp, error) {
	if req == nil || req.GetId() == 0 && strings.TrimSpace(req.GetIdempotencyKey()) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.usecase.CancelExecution(ctx, req.GetId(), req.GetIdempotencyKey())
	if err != nil {
		return nil, err
	}
	item := &schedulerv1.DelayedTaskExecutionRecord{
		Id:                 row.ID,
		DelayedTaskId:      row.DelayedTaskID,
		DelayedTaskVersion: row.DelayedTaskVersion,
		IdempotencyKey:     row.IdempotencyKey,
		TriggerType:        schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType),
		ScheduleKey:        row.ScheduleKey,
		ScheduledAt:        timestamppb.New(row.ScheduledAt),
		Status:             schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status),
		Attempt:            row.Attempt,
		MaxAttempts:        row.MaxAttempts,
		TimeoutSeconds:     int32(row.Timeout / time.Second),
		WorkerId:           row.WorkerID,
		Payload:            row.Payload,
		LastError:          row.LastError,
		TraceId:            row.TraceID,
	}
	return &schedulerv1.CancelSchedulerDelayedTaskExecution_Resp{
		Row: item,
	}, nil
}

func (s *SchedulerDelayedTaskService) PageExecutionRecords(
	ctx context.Context,
	req *schedulerv1.PageSchedulerDelayedTaskExecutionRecords_Req,
) (*schedulerv1.PageSchedulerDelayedTaskExecutionRecords_Resp, error) {
	query := &usecase.DelayedTaskExecutionRecordPageReq{}
	if req != nil {
		query.Page = req.GetPage()
		if req.GetQuery() != nil {
			query.DelayedTaskID = req.GetQuery().DelayedTaskId
			query.IdempotencyKey = req.GetQuery().IdempotencyKey
			if req.GetQuery().Status != nil && req.GetQuery().GetStatus() != schedulerv1enum.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_UNSPECIFIED {
				value, ok := schedulerenum.TaskExecutionStatusMap.ToEnum(req.GetQuery().GetStatus())
				if !ok {
					return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
				}
				query.Status = &value
			}
			if req.GetQuery().TriggerType != nil && req.GetQuery().GetTriggerType() != schedulerv1enum.SchedulerTaskTriggerType_SCHEDULER_TASK_TRIGGER_TYPE_UNSPECIFIED {
				value, ok := schedulerenum.TaskTriggerTypeMap.ToEnum(req.GetQuery().GetTriggerType())
				if !ok {
					return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
				}
				query.TriggerType = &value
			}
		}
	}
	resp, err := s.usecase.PageExecutionRecords(ctx, query)
	if err != nil {
		return nil, err
	}
	rows := make([]*schedulerv1.DelayedTaskExecutionRecord, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		item := &schedulerv1.DelayedTaskExecutionRecord{
			Id:                 row.ID,
			DelayedTaskId:      row.DelayedTaskID,
			DelayedTaskVersion: row.DelayedTaskVersion,
			IdempotencyKey:     row.IdempotencyKey,
			TriggerType:        schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType),
			ScheduleKey:        row.ScheduleKey,
			ScheduledAt:        timestamppb.New(row.ScheduledAt),
			Status:             schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status),
			Attempt:            row.Attempt,
			MaxAttempts:        row.MaxAttempts,
			TimeoutSeconds:     int32(row.Timeout / time.Second),
			WorkerId:           row.WorkerID,
			Payload:            row.Payload,
			LastError:          row.LastError,
			TraceId:            row.TraceID,
		}
		if row.Duration != nil {
			item.DurationMs = new(row.Duration.Milliseconds())
		}
		if row.StartedAt != nil {
			item.StartedAt = timestamppb.New(*row.StartedAt)
		}
		if row.FinishedAt != nil {
			item.FinishedAt = timestamppb.New(*row.FinishedAt)
		}
		if row.CreatedAt != nil {
			item.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		rows = append(rows, item)
	}
	return &schedulerv1.PageSchedulerDelayedTaskExecutionRecords_Resp{
		Page: resp.Page,
		Rows: rows,
	}, nil
}
