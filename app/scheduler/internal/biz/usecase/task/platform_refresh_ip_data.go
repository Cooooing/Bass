package task

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	platformv1 "common/proto/gen/platform/v1"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
	"fmt"
	"time"

	schedulerenum "scheduler/internal/enum"
)

// PlatformRefreshIPData refreshes the Platform-owned IP location data set from
// its external source. Platform validates, archives, and activates the data.
type PlatformRefreshIPData struct{ platformClient *rpc.PlatformClient }

func NewPlatformRefreshIPData(platformClient *rpc.PlatformClient) *PlatformRefreshIPData {
	return &PlatformRefreshIPData{platformClient: platformClient}
}

func (*PlatformRefreshIPData) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNamePlatformRefreshIPData
}

func (*PlatformRefreshIPData) Title() string { return "刷新 IP 地理数据" }

func (*PlatformRefreshIPData) Description() string {
	return "调用 Platform 下载 IP 地理数据、归档至对象存储并热更新本地查询器。"
}

func (t *PlatformRefreshIPData) DefaultScheduledTasks() []*DefaultScheduledTask {
	return []*DefaultScheduledTask{{
		Title:         t.Title(),
		Description:   t.Description(),
		Enabled:       true,
		CronSpec:      "0 0 2 1 * ?",
		Payload:       commonscheduler.MustMarshalPayload(&schedulerv1.RefreshIPDataTaskPayload{}),
		Timeout:       10 * time.Minute,
		MaxAttempts:   3,
		MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteLatest,
		AllowOverlap:  false,
	}}
}

func (*PlatformRefreshIPData) DefaultDelayedTasks() []*DefaultDelayedTask { return nil }

func (t *PlatformRefreshIPData) Execute(ctx context.Context, execution Execution) (*Result, error) {
	payload := &schedulerv1.RefreshIPDataTaskPayload{}
	if err := commonscheduler.UnmarshalPayload(execution.Payload, payload); err != nil {
		return nil, err
	}
	response, err := t.platformClient.IPResolution.UpdateIPDataFromSource(ctx, &platformv1.UpdateIPDataFromSource_Req{})
	if err != nil || response == nil {
		if err == nil {
			err = fmt.Errorf("IP data refresh returned an empty response")
		}
		return nil, err
	}
	return &Result{Summary: "IP location data refreshed"}, nil
}
