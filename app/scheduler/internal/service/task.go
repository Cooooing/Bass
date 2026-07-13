package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/repo"
	"scheduler/internal/biz/usecase"
	schedulerenum "scheduler/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SchedulerTaskService struct {
	schedulerv1.UnimplementedSchedulerTaskServiceServer
	taskUsecase *usecase.TaskUsecase
}

func NewSchedulerTaskService(taskUsecase *usecase.TaskUsecase) *SchedulerTaskService {
	return &SchedulerTaskService{taskUsecase: taskUsecase}
}

func (s *SchedulerTaskService) RegisterGrpc(gs *grpc.Server) {
	schedulerv1.RegisterSchedulerTaskServiceServer(gs, s)
}

func (s *SchedulerTaskService) Upsert(ctx context.Context, req *schedulerv1.UpsertSchedulerTask_Request) (*schedulerv1.UpsertSchedulerTask_Reply, error) {
	if req.GetName() == "" || req.GetTitle() == "" || req.GetCronSpec() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	alertEnabled := true
	if req.AlertEnabled != nil {
		alertEnabled = req.GetAlertEnabled()
	}
	row, err := s.taskUsecase.Upsert(ctx, &model.Task{
		ID:             req.GetId(),
		Name:           req.GetName(),
		Title:          req.GetTitle(),
		Description:    req.GetDescription(),
		Enabled:        req.GetEnabled(),
		CronSpec:       req.GetCronSpec(),
		Payload:        req.GetPayload(),
		TimeoutSeconds: req.GetTimeoutSeconds(),
		AllowOverlap:   req.GetAllowOverlap(),
		AlertEnabled:   alertEnabled,
	})
	if err != nil {
		return nil, err
	}
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.UpsertSchedulerTask_Reply{Row: &schedulerv1.SchedulerTask{
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		Id:             row.ID,
		Name:           row.Name,
		Title:          row.Title,
		Description:    row.Description,
		Enabled:        row.Enabled,
		CronSpec:       row.CronSpec,
		Payload:        row.Payload,
		TimeoutSeconds: row.TimeoutSeconds,
		AllowOverlap:   row.AllowOverlap,
		AlertEnabled:   row.AlertEnabled,
		Version:        row.Version,
	}}, nil
}

func (s *SchedulerTaskService) Get(ctx context.Context, req *schedulerv1.GetSchedulerTask_Request) (*schedulerv1.GetSchedulerTask_Reply, error) {
	row, err := s.taskUsecase.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.GetSchedulerTask_Reply{Row: &schedulerv1.SchedulerTask{
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
		Id:             row.ID,
		Name:           row.Name,
		Title:          row.Title,
		Description:    row.Description,
		Enabled:        row.Enabled,
		CronSpec:       row.CronSpec,
		Payload:        row.Payload,
		TimeoutSeconds: row.TimeoutSeconds,
		AllowOverlap:   row.AllowOverlap,
		AlertEnabled:   row.AlertEnabled,
		Version:        row.Version,
	}}, nil
}

func (s *SchedulerTaskService) Page(ctx context.Context, req *schedulerv1.PageSchedulerTasks_Request) (*schedulerv1.PageSchedulerTasks_Reply, error) {
	query := &repo.TaskGetReq{}
	if req.GetQuery() != nil {
		query.IDs = req.GetQuery().GetIds()
		if req.GetQuery().Name != nil {
			query.Name = new(req.GetQuery().GetName())
		}
		if req.GetQuery().Title != nil {
			query.Title = new(req.GetQuery().GetTitle())
		}
		if req.GetQuery().Enabled != nil {
			query.Enabled = new(req.GetQuery().GetEnabled())
		}
	}
	rows, page, err := s.taskUsecase.Page(ctx, req.GetPage(), query)
	if err != nil {
		return nil, err
	}
	replyRows := make([]*schedulerv1.SchedulerTask, 0, len(rows))
	for _, row := range rows {
		var createdAt *timestamppb.Timestamp
		if row.CreatedAt != nil {
			createdAt = timestamppb.New(*row.CreatedAt)
		}
		var updatedAt *timestamppb.Timestamp
		if row.UpdatedAt != nil {
			updatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, &schedulerv1.SchedulerTask{
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			Id:             row.ID,
			Name:           row.Name,
			Title:          row.Title,
			Description:    row.Description,
			Enabled:        row.Enabled,
			CronSpec:       row.CronSpec,
			Payload:        row.Payload,
			TimeoutSeconds: row.TimeoutSeconds,
			AllowOverlap:   row.AllowOverlap,
			AlertEnabled:   row.AlertEnabled,
			Version:        row.Version,
		})
	}
	return &schedulerv1.PageSchedulerTasks_Reply{Page: page, Rows: replyRows}, nil
}

func (s *SchedulerTaskService) ListAvailableTasks(ctx context.Context, req *schedulerv1.ListSchedulerAvailableTasks_Request) (*schedulerv1.ListSchedulerAvailableTasks_Reply, error) {
	keyword := ""
	if req.GetQuery() != nil {
		keyword = req.GetQuery().GetKeyword()
	}
	rows := s.taskUsecase.ListAvailableTasks(keyword)
	replyRows := make([]*schedulerv1.SchedulerAvailableTask, 0, len(rows))
	for _, row := range rows {
		replyRows = append(replyRows, &schedulerv1.SchedulerAvailableTask{
			Name:        row.Name,
			Title:       row.Title,
			Description: row.Description,
		})
	}
	return &schedulerv1.ListSchedulerAvailableTasks_Reply{Rows: replyRows}, nil
}

func (s *SchedulerTaskService) PageExecutionRecords(ctx context.Context, req *schedulerv1.PageSchedulerTaskExecutionRecords_Request) (*schedulerv1.PageSchedulerTaskExecutionRecords_Reply, error) {
	query := &repo.TaskExecutionRecordGetReq{}
	if req.GetQuery() != nil {
		query.IDs = req.GetQuery().GetIds()
		if req.GetQuery().TaskId != nil {
			query.TaskID = new(req.GetQuery().GetTaskId())
		}
		if req.GetQuery().Status != nil {
			statusValue, ok := schedulerenum.TaskExecutionStatusMap.ToEnum(req.GetQuery().GetStatus())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.Status = &statusValue
		}
		if req.GetQuery().TriggerType != nil {
			triggerType, ok := schedulerenum.TaskTriggerTypeMap.ToEnum(req.GetQuery().GetTriggerType())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.TriggerType = &triggerType
		}
	}
	rows, page, err := s.taskUsecase.PageExecutionRecords(ctx, req.GetPage(), query)
	if err != nil {
		return nil, err
	}
	replyRows := make([]*schedulerv1.SchedulerTaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		var createdAt *timestamppb.Timestamp
		if row.CreatedAt != nil {
			createdAt = timestamppb.New(*row.CreatedAt)
		}
		var updatedAt *timestamppb.Timestamp
		if row.UpdatedAt != nil {
			updatedAt = timestamppb.New(*row.UpdatedAt)
		}
		var startedAt *timestamppb.Timestamp
		if row.StartedAt != nil {
			startedAt = timestamppb.New(*row.StartedAt)
		}
		var finishedAt *timestamppb.Timestamp
		if row.FinishedAt != nil {
			finishedAt = timestamppb.New(*row.FinishedAt)
		}
		statusValue, ok := schedulerenum.TaskExecutionStatusMap.ToProto(row.Status)
		if !ok {
			statusValue = schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_UNSPECIFIED
		}
		triggerType, ok := schedulerenum.TaskTriggerTypeMap.ToProto(row.TriggerType)
		if !ok {
			triggerType = schedulerv1.SchedulerTaskTriggerType_SCHEDULER_TASK_TRIGGER_TYPE_UNSPECIFIED
		}
		replyRows = append(replyRows, &schedulerv1.SchedulerTaskExecutionRecord{
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Id:          row.ID,
			TaskId:      row.TaskID,
			ScheduledAt: timestamppb.New(row.ScheduledAt),
			StartedAt:   startedAt,
			FinishedAt:  finishedAt,
			DurationMs:  row.DurationMS,
			Status:      statusValue,
			TriggerType: triggerType,
			TaskVersion: row.TaskVersion,
			WorkerId:    row.WorkerID,
			Payload:     row.Payload,
			LastError:   row.LastError,
			TraceId:     row.TraceID,
		})
	}
	return &schedulerv1.PageSchedulerTaskExecutionRecords_Reply{Page: page, Rows: replyRows}, nil
}

func (s *SchedulerTaskService) Trigger(ctx context.Context, req *schedulerv1.TriggerSchedulerTask_Request) (*schedulerv1.TriggerSchedulerTask_Reply, error) {
	row, err := s.taskUsecase.Trigger(ctx, req.GetId(), req.GetPayload())
	if err != nil {
		return nil, err
	}
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	var startedAt *timestamppb.Timestamp
	if row.StartedAt != nil {
		startedAt = timestamppb.New(*row.StartedAt)
	}
	var finishedAt *timestamppb.Timestamp
	if row.FinishedAt != nil {
		finishedAt = timestamppb.New(*row.FinishedAt)
	}
	statusValue, ok := schedulerenum.TaskExecutionStatusMap.ToProto(row.Status)
	if !ok {
		statusValue = schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_UNSPECIFIED
	}
	triggerType, ok := schedulerenum.TaskTriggerTypeMap.ToProto(row.TriggerType)
	if !ok {
		triggerType = schedulerv1.SchedulerTaskTriggerType_SCHEDULER_TASK_TRIGGER_TYPE_UNSPECIFIED
	}
	return &schedulerv1.TriggerSchedulerTask_Reply{Row: &schedulerv1.SchedulerTaskExecutionRecord{
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Id:          row.ID,
		TaskId:      row.TaskID,
		ScheduledAt: timestamppb.New(row.ScheduledAt),
		StartedAt:   startedAt,
		FinishedAt:  finishedAt,
		DurationMs:  row.DurationMS,
		Status:      statusValue,
		TriggerType: triggerType,
		TaskVersion: row.TaskVersion,
		WorkerId:    row.WorkerID,
		Payload:     row.Payload,
		LastError:   row.LastError,
		TraceId:     row.TraceID,
	}}, nil
}

func (s *SchedulerTaskService) CancelExecution(ctx context.Context, req *schedulerv1.CancelSchedulerTaskExecution_Request) (*schedulerv1.CancelSchedulerTaskExecution_Reply, error) {
	row, err := s.taskUsecase.CancelExecution(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	var startedAt *timestamppb.Timestamp
	if row.StartedAt != nil {
		startedAt = timestamppb.New(*row.StartedAt)
	}
	var finishedAt *timestamppb.Timestamp
	if row.FinishedAt != nil {
		finishedAt = timestamppb.New(*row.FinishedAt)
	}
	statusValue, ok := schedulerenum.TaskExecutionStatusMap.ToProto(row.Status)
	if !ok {
		statusValue = schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_UNSPECIFIED
	}
	triggerType, ok := schedulerenum.TaskTriggerTypeMap.ToProto(row.TriggerType)
	if !ok {
		triggerType = schedulerv1.SchedulerTaskTriggerType_SCHEDULER_TASK_TRIGGER_TYPE_UNSPECIFIED
	}
	return &schedulerv1.CancelSchedulerTaskExecution_Reply{Row: &schedulerv1.SchedulerTaskExecutionRecord{
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Id:          row.ID,
		TaskId:      row.TaskID,
		ScheduledAt: timestamppb.New(row.ScheduledAt),
		StartedAt:   startedAt,
		FinishedAt:  finishedAt,
		DurationMs:  row.DurationMS,
		Status:      statusValue,
		TriggerType: triggerType,
		TaskVersion: row.TaskVersion,
		WorkerId:    row.WorkerID,
		Payload:     row.Payload,
		LastError:   row.LastError,
		TraceId:     row.TraceID,
	}}, nil
}

func (s *SchedulerTaskService) CheckExecutionRuntimes(ctx context.Context, req *schedulerv1.CheckSchedulerTaskExecutionRuntimes_Request) (*schedulerv1.CheckSchedulerTaskExecutionRuntimes_Reply, error) {
	rows, err := s.taskUsecase.CheckExecutionRuntimes(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	replyRows := make([]*schedulerv1.SchedulerTaskExecutionRuntime, 0, len(rows))
	for _, row := range rows {
		stateValue, ok := schedulerenum.TaskExecutionRuntimeStateMap.ToProto(row.State)
		if !ok {
			stateValue = schedulerv1.SchedulerTaskExecutionRuntimeState_SCHEDULER_TASK_EXECUTION_RUNTIME_STATE_UNSPECIFIED
		}
		replyRows = append(replyRows, &schedulerv1.SchedulerTaskExecutionRuntime{
			ExecutionRecordId: row.ExecutionRecordID,
			TaskId:            row.TaskID,
			State:             stateValue,
		})
	}
	return &schedulerv1.CheckSchedulerTaskExecutionRuntimes_Reply{Rows: replyRows}, nil
}

func (s *SchedulerTaskService) MarkExecutionsUnknown(ctx context.Context, req *schedulerv1.MarkSchedulerTaskExecutionsUnknown_Request) (*schedulerv1.MarkSchedulerTaskExecutionsUnknown_Reply, error) {
	rows, err := s.taskUsecase.MarkExecutionsUnknown(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	replyRows := make([]*schedulerv1.SchedulerTaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		var createdAt *timestamppb.Timestamp
		if row.CreatedAt != nil {
			createdAt = timestamppb.New(*row.CreatedAt)
		}
		var updatedAt *timestamppb.Timestamp
		if row.UpdatedAt != nil {
			updatedAt = timestamppb.New(*row.UpdatedAt)
		}
		var startedAt *timestamppb.Timestamp
		if row.StartedAt != nil {
			startedAt = timestamppb.New(*row.StartedAt)
		}
		var finishedAt *timestamppb.Timestamp
		if row.FinishedAt != nil {
			finishedAt = timestamppb.New(*row.FinishedAt)
		}
		statusValue, ok := schedulerenum.TaskExecutionStatusMap.ToProto(row.Status)
		if !ok {
			statusValue = schedulerv1.SchedulerTaskExecutionStatus_SCHEDULER_TASK_EXECUTION_STATUS_UNSPECIFIED
		}
		triggerType, ok := schedulerenum.TaskTriggerTypeMap.ToProto(row.TriggerType)
		if !ok {
			triggerType = schedulerv1.SchedulerTaskTriggerType_SCHEDULER_TASK_TRIGGER_TYPE_UNSPECIFIED
		}
		replyRows = append(replyRows, &schedulerv1.SchedulerTaskExecutionRecord{
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			Id:          row.ID,
			TaskId:      row.TaskID,
			ScheduledAt: timestamppb.New(row.ScheduledAt),
			StartedAt:   startedAt,
			FinishedAt:  finishedAt,
			DurationMs:  row.DurationMS,
			Status:      statusValue,
			TriggerType: triggerType,
			TaskVersion: row.TaskVersion,
			WorkerId:    row.WorkerID,
			Payload:     row.Payload,
			LastError:   row.LastError,
			TraceId:     row.TraceID,
		})
	}
	return &schedulerv1.MarkSchedulerTaskExecutionsUnknown_Reply{Rows: replyRows, UpdatedCount: uint32(len(replyRows))}, nil
}
