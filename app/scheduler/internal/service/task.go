package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/usecase"
	schedulerenum "scheduler/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
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

func (s *SchedulerTaskService) RegisterHttp(hs *http.Server) {}

func (s *SchedulerTaskService) Upsert(ctx context.Context, req *schedulerv1.UpsertSchedulerTask_Request) (*schedulerv1.UpsertSchedulerTask_Response, error) {
	if req.GetName() == "" || req.GetTitle() == "" || req.GetCronSpec() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	alertEnabled := true
	if req.AlertEnabled != nil {
		alertEnabled = req.GetAlertEnabled()
	}
	upsertResponse, err := s.taskUsecase.Upsert(ctx, &usecase.TaskUpsertReq{Row: &model.Task{
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
	}})
	if err != nil {
		return nil, err
	}
	row := upsertResponse.Row
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.UpsertSchedulerTask_Response{Row: &schedulerv1.UpsertSchedulerTask_Response_SchedulerTask{
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

func (s *SchedulerTaskService) Get(ctx context.Context, req *schedulerv1.GetSchedulerTask_Request) (*schedulerv1.GetSchedulerTask_Response, error) {
	getResponse, err := s.taskUsecase.Get(ctx, &usecase.TaskGetReq{ID: req.GetId()})
	if err != nil {
		return nil, err
	}
	row := getResponse.Row
	var createdAt *timestamppb.Timestamp
	if row.CreatedAt != nil {
		createdAt = timestamppb.New(*row.CreatedAt)
	}
	var updatedAt *timestamppb.Timestamp
	if row.UpdatedAt != nil {
		updatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &schedulerv1.GetSchedulerTask_Response{Row: &schedulerv1.GetSchedulerTask_Response_SchedulerTask{
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

func (s *SchedulerTaskService) Page(ctx context.Context, req *schedulerv1.PageSchedulerTasks_Request) (*schedulerv1.PageSchedulerTasks_Response, error) {
	query := &usecase.TaskPageReq{}
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
	query.Page = req.GetPage()
	pageResponse, err := s.taskUsecase.Page(ctx, query)
	if err != nil {
		return nil, err
	}
	rows := pageResponse.Rows
	replyRows := make([]*schedulerv1.PageSchedulerTasks_Response_SchedulerTask, 0, len(rows))
	for _, row := range rows {
		var createdAt *timestamppb.Timestamp
		if row.CreatedAt != nil {
			createdAt = timestamppb.New(*row.CreatedAt)
		}
		var updatedAt *timestamppb.Timestamp
		if row.UpdatedAt != nil {
			updatedAt = timestamppb.New(*row.UpdatedAt)
		}
		replyRows = append(replyRows, &schedulerv1.PageSchedulerTasks_Response_SchedulerTask{
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
	return &schedulerv1.PageSchedulerTasks_Response{Page: pageResponse.Page, Rows: replyRows}, nil
}

func (s *SchedulerTaskService) ListAvailableTasks(ctx context.Context, req *schedulerv1.ListSchedulerAvailableTasks_Request) (*schedulerv1.ListSchedulerAvailableTasks_Response, error) {
	keyword := ""
	if req.GetQuery() != nil {
		keyword = req.GetQuery().GetKeyword()
	}
	listAvailableResponse, err := s.taskUsecase.ListAvailableTasks(ctx, &usecase.TaskListAvailableTasksReq{Keyword: keyword})
	if err != nil {
		return nil, err
	}
	rows := listAvailableResponse.Rows
	replyRows := make([]*schedulerv1.ListSchedulerAvailableTasks_Response_SchedulerAvailableTask, 0, len(rows))
	for _, row := range rows {
		replyRows = append(replyRows, &schedulerv1.ListSchedulerAvailableTasks_Response_SchedulerAvailableTask{
			Name:        row.Name,
			Title:       row.Title,
			Description: row.Description,
		})
	}
	return &schedulerv1.ListSchedulerAvailableTasks_Response{Rows: replyRows}, nil
}

func (s *SchedulerTaskService) PageExecutionRecords(ctx context.Context, req *schedulerv1.PageSchedulerTaskExecutionRecords_Request) (*schedulerv1.PageSchedulerTaskExecutionRecords_Response, error) {
	query := &usecase.TaskExecutionRecordPageReq{}
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
	query.Page = req.GetPage()
	pageResponse, err := s.taskUsecase.PageExecutionRecords(ctx, query)
	if err != nil {
		return nil, err
	}
	rows := pageResponse.Rows
	replyRows := make([]*schedulerv1.PageSchedulerTaskExecutionRecords_Response_SchedulerTaskExecutionRecord, 0, len(rows))
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
		statusValue := schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status)
		triggerType := schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType)
		replyRows = append(replyRows, &schedulerv1.PageSchedulerTaskExecutionRecords_Response_SchedulerTaskExecutionRecord{
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
	return &schedulerv1.PageSchedulerTaskExecutionRecords_Response{Page: pageResponse.Page, Rows: replyRows}, nil
}

func (s *SchedulerTaskService) Trigger(ctx context.Context, req *schedulerv1.TriggerSchedulerTask_Request) (*schedulerv1.TriggerSchedulerTask_Response, error) {
	triggerResponse, err := s.taskUsecase.Trigger(ctx, &usecase.TaskTriggerReq{ID: req.GetId(), Payload: req.GetPayload()})
	if err != nil {
		return nil, err
	}
	row := triggerResponse.Row
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
	statusValue := schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status)
	triggerType := schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType)
	return &schedulerv1.TriggerSchedulerTask_Response{Row: &schedulerv1.TriggerSchedulerTask_Response_SchedulerTaskExecutionRecord{
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

func (s *SchedulerTaskService) CancelExecution(ctx context.Context, req *schedulerv1.CancelSchedulerTaskExecution_Request) (*schedulerv1.CancelSchedulerTaskExecution_Response, error) {
	cancelResponse, err := s.taskUsecase.CancelExecution(ctx, &usecase.TaskCancelExecutionReq{ID: req.GetId()})
	if err != nil {
		return nil, err
	}
	row := cancelResponse.Row
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
	statusValue := schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status)
	triggerType := schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType)
	return &schedulerv1.CancelSchedulerTaskExecution_Response{Row: &schedulerv1.CancelSchedulerTaskExecution_Response_SchedulerTaskExecutionRecord{
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

func (s *SchedulerTaskService) CheckExecutionRuntimes(ctx context.Context, req *schedulerv1.CheckSchedulerTaskExecutionRuntimes_Request) (*schedulerv1.CheckSchedulerTaskExecutionRuntimes_Response, error) {
	checkResponse, err := s.taskUsecase.CheckExecutionRuntimes(ctx, &usecase.TaskCheckExecutionRuntimesReq{IDs: req.GetIds()})
	if err != nil {
		return nil, err
	}
	rows := checkResponse.Rows
	replyRows := make([]*schedulerv1.CheckSchedulerTaskExecutionRuntimes_Response_SchedulerTaskExecutionRuntime, 0, len(rows))
	for _, row := range rows {
		stateValue := schedulerenum.TaskExecutionRuntimeStateMap.MustToProto(row.State)
		replyRows = append(replyRows, &schedulerv1.CheckSchedulerTaskExecutionRuntimes_Response_SchedulerTaskExecutionRuntime{
			ExecutionRecordId: row.ExecutionRecordID,
			TaskId:            row.TaskID,
			State:             stateValue,
		})
	}
	return &schedulerv1.CheckSchedulerTaskExecutionRuntimes_Response{Rows: replyRows}, nil
}

func (s *SchedulerTaskService) MarkExecutionsUnknown(ctx context.Context, req *schedulerv1.MarkSchedulerTaskExecutionsUnknown_Request) (*schedulerv1.MarkSchedulerTaskExecutionsUnknown_Response, error) {
	markResponse, err := s.taskUsecase.MarkExecutionsUnknown(ctx, &usecase.TaskMarkExecutionsUnknownReq{IDs: req.GetIds()})
	if err != nil {
		return nil, err
	}
	rows := markResponse.Rows
	replyRows := make([]*schedulerv1.MarkSchedulerTaskExecutionsUnknown_Response_SchedulerTaskExecutionRecord, 0, len(rows))
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
		statusValue := schedulerenum.TaskExecutionStatusMap.MustToProto(row.Status)
		triggerType := schedulerenum.TaskTriggerTypeMap.MustToProto(row.TriggerType)
		replyRows = append(replyRows, &schedulerv1.MarkSchedulerTaskExecutionsUnknown_Response_SchedulerTaskExecutionRecord{
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
	return &schedulerv1.MarkSchedulerTaskExecutionsUnknown_Response{Rows: replyRows, UpdatedCount: uint32(len(replyRows))}, nil
}
