package repo

import (
	commonclient "common/pkg/client"
	"common/pkg/constant"
	"context"
	"fmt"
	"log/slog"
	bizrepo "scheduler/internal/biz/repo"
	"scheduler/internal/config"
	"time"
)

var _ bizrepo.TaskAlert = (*TaskAlert)(nil)

type TaskAlert struct {
	logger     *slog.Logger
	larkClient *commonclient.LarkWebhookClient
	conf       *config.Bootstrap
}

func NewTaskAlert(logger *slog.Logger, larkClient *commonclient.LarkWebhookClient, conf *config.Bootstrap) bizrepo.TaskAlert {
	return &TaskAlert{logger: logger, larkClient: larkClient, conf: conf}
}

func (a *TaskAlert) Alert(ctx context.Context, req *bizrepo.TaskAlertReq) (*bizrepo.TaskAlertResponse, error) {
	task := req.Task
	record := req.Record
	reason := req.Reason
	a.logger.ErrorContext(ctx, "scheduler task alert", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name, constant.LogFieldExecutionID, record.ID, constant.LogFieldTraceID, record.TraceID, constant.LogFieldStatus, record.Status, constant.LogFieldReason, reason, constant.LogFieldLastError, record.LastError)
	if a.conf.GetAlert() == nil || a.conf.GetAlert().GetLarkWebhook() == nil || a.conf.GetAlert().GetLarkWebhook().GetToken() == "" {
		return &bizrepo.TaskAlertResponse{}, nil
	}
	lark := a.conf.GetAlert().GetLarkWebhook()
	timeout := time.Duration(0)
	if lark.GetTimeout() != nil {
		timeout = lark.GetTimeout().AsDuration()
	}
	err := a.larkClient.SendText(ctx, &commonclient.LarkWebhookRequest{
		BaseURL: lark.GetBaseUrl(),
		Token:   lark.GetToken(),
		Secret:  lark.GetSecret(),
		Timeout: timeout,
		Text: fmt.Sprintf(
			"Scheduler task alert\nreason: %s\ntask_id: %d\ntask_name: %s\ntask_title: %s\nexecution_id: %d\ntrace_id: %s\nstatus: %s\nlast_error: %s",
			reason,
			task.ID,
			task.Name,
			task.Title,
			record.ID,
			record.TraceID,
			record.Status,
			record.LastError,
		),
	})
	if err != nil {
		a.logger.ErrorContext(ctx, "send scheduler task alert failed", constant.LogFieldKind, constant.LogKindLark, constant.LogFieldTaskID, task.ID, constant.LogFieldExecutionID, record.ID, constant.LogFieldReason, reason, constant.LogFieldErr, err)
		return nil, err
	}
	return &bizrepo.TaskAlertResponse{}, nil
}
