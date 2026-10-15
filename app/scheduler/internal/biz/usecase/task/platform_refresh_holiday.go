package task

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	platformv1 "common/proto/gen/platform/v1"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
	"time"

	schedulerenum "scheduler/internal/enum"
)

type PlatformRefreshHoliday struct {
	platformClient *rpc.PlatformClient
	title          string
	description    string
}

func NewPlatformRefreshHoliday(
	platformClient *rpc.PlatformClient,
) *PlatformRefreshHoliday {
	return &PlatformRefreshHoliday{
		platformClient: platformClient,
		title:          "刷新节假日缓存",
		description:    "调用 platform.HolidayService.RefreshHoliday 将外部节假日数据写入平台表。",
	}
}

func (t *PlatformRefreshHoliday) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNamePlatformRefreshHoliday
}

func (t *PlatformRefreshHoliday) Title() string {
	return t.title
}

func (t *PlatformRefreshHoliday) Description() string {
	return t.description
}

func (t *PlatformRefreshHoliday) DefaultScheduledTasks() []*DefaultScheduledTask {
	return []*DefaultScheduledTask{
		{
			Title:         t.Title(),
			Description:   t.Description(),
			Enabled:       true,
			CronSpec:      "0 0 * * * ?",
			Payload:       commonscheduler.MustMarshalPayload(&schedulerv1.RefreshHolidayTaskPayload{}),
			Timeout:       30 * time.Second,
			MaxAttempts:   1,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteLatest,
			AllowOverlap:  false,
		},
	}
}

func (t *PlatformRefreshHoliday) DefaultDelayedTasks() []*DefaultDelayedTask {
	return nil
}

func (t *PlatformRefreshHoliday) Execute(ctx context.Context, payload string) error {
	data := &schedulerv1.RefreshHolidayTaskPayload{}
	if err := commonscheduler.UnmarshalPayload(payload, data); err != nil {
		return err
	}
	_, err := t.platformClient.Holiday.RefreshHoliday(ctx, &platformv1.RefreshHoliday_Req{Date: data.GetDate()})
	return err
}
