package enum

import (
	event "common/proto/gen/common/enums/event/v1"
	"slices"

	"github.com/nats-io/nats.go/jetstream"
)

// EventStream identifies an application-managed JetStream stream.
type EventStream string

const (
	EventStreamUser                 EventStream = "user_events"            // 用户领域事件流
	EventStreamContent              EventStream = "content_events"         // 内容领域事件流
	EventStreamIM                   EventStream = "im_events"              // 即时通信领域事件流
	EventStreamSchedulerTasks       EventStream = "scheduler_tasks"        // 调度任务流
	EventStreamPlatformAssetUploads EventStream = "platform_asset_uploads" // 平台资产上传流
)

var EventStreamMap = NewMapping[EventStream, event.EventStream](map[EventStream]Entry[EventStream, event.EventStream]{
	EventStreamUser:                 {Proto: event.EventStream_EVENT_STREAM_USER},
	EventStreamContent:              {Proto: event.EventStream_EVENT_STREAM_CONTENT},
	EventStreamIM:                   {Proto: event.EventStream_EVENT_STREAM_IM},
	EventStreamSchedulerTasks:       {Proto: event.EventStream_EVENT_STREAM_SCHEDULER_TASKS},
	EventStreamPlatformAssetUploads: {Proto: event.EventStream_EVENT_STREAM_PLATFORM_ASSET_UPLOADS},
})

type EventStreamConfig struct {
	Name              EventStream
	Subjects          []string
	Retention         jetstream.RetentionPolicy
	AllowMsgSchedules bool
	AllowMsgTTL       bool
	AllowRollup       bool
}

var eventStreams = map[EventStream]EventStreamConfig{
	EventStreamUser: {
		Name: EventStreamUser,
		Subjects: []string{
			EventSubjectUserFollow.String(),
			EventSubjectUserUnfollow.String(),
			EventSubjectUserRegister.String(),
			EventSubjectUserLogin.String(),
			EventSubjectUserLogout.String(),
			EventSubjectUserEmailVerificationCode.String(),
			EventSubjectUserPhoneVerificationCode.String(),
			EventSubjectUserBlock.String(),
			EventSubjectUserUnblock.String(),
			EventSubjectUserTotpEnable.String(),
			EventSubjectUserTotpDisable.String(),
			EventSubjectUserAccountCancelled.String(),
			EventSubjectUserAccountBanned.String(),
			EventSubjectUserAccountUnbanned.String(),
			EventSubjectUserCheckinCompleted.String(),
		},
	},
	EventStreamContent: {
		Name: EventStreamContent,
		Subjects: []string{
			EventSubjectContentArticlePublish.String(),
			EventSubjectContentArticleLike.String(),
			EventSubjectContentArticleThank.String(),
			EventSubjectContentArticleReward.String(),
			EventSubjectContentArticleCollect.String(),
			EventSubjectContentArticleWatch.String(),
			EventSubjectContentArticleAcceptAnswer.String(),
			EventSubjectContentArticlePostscriptAdd.String(),
			EventSubjectContentArticleStatusUpdate.String(),
			EventSubjectContentArticleLotteryWon.String(),
			EventSubjectContentArticlePollEnded.String(),
			EventSubjectContentArticleLotteryDrawn.String(),
			EventSubjectContentCommentPublish.String(),
			EventSubjectContentCommentLike.String(),
			EventSubjectContentCommentThank.String(),
			EventSubjectContentCommentStatusUpdate.String(),
		},
	},
	EventStreamIM: {
		Name:     EventStreamIM,
		Subjects: []string{EventSubjectIMGlobalChatMessageCreated.String()},
	},
	EventStreamSchedulerTasks: {
		Name: EventStreamSchedulerTasks,
		Subjects: []string{
			SchedulerScheduledTaskScheduleSubjectPrefix + ".>",
			SchedulerScheduledTaskExecuteSubject + ".>",
			SchedulerDelayedTaskScheduleSubjectPrefix + ".>",
			SchedulerDelayedTaskExecuteSubject,
		},
		Retention:         jetstream.WorkQueuePolicy,
		AllowMsgSchedules: true,
		AllowMsgTTL:       true,
		AllowRollup:       true,
	},
	EventStreamPlatformAssetUploads: {
		Name:      EventStreamPlatformAssetUploads,
		Subjects:  []string{PlatformAssetUploadSubject},
		Retention: jetstream.WorkQueuePolicy,
	},
}

// EventStreamConfigFor 返回调用方可安全读取的 stream 拓扑副本。
func EventStreamConfigFor(stream EventStream) EventStreamConfig {
	config := eventStreams[stream]
	config.Subjects = slices.Clone(config.Subjects)
	return config
}

func (e EventStream) String() string {
	return string(e)
}
