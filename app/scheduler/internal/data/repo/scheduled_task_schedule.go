package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	nats "common/pkg/client/nats"
	commonenum "common/pkg/enum"

	"scheduler/internal/biz/repo"
	"scheduler/internal/config"
	schedulerenum "scheduler/internal/enum"

	"github.com/google/uuid"

	"github.com/nats-io/nats.go/jetstream"
)

type ScheduledTaskScheduleRepo struct {
	logger          *slog.Logger
	natsClient      *nats.Client
	js              jetstream.JetStream
	streamName      string
	executeSubject  string
	consumerDurable string
	scheduleTTL     time.Duration
	subscription    *nats.Subscription
}

func NewScheduledTaskScheduleRepo(
	logger *slog.Logger,
	natsClient *nats.Client,
	conf *config.Bootstrap,
) (repo.ScheduledTaskScheduleRepo, error) {
	if conf == nil || conf.GetBusiness() == nil {
		return nil, fmt.Errorf("scheduler business config is required")
	}
	scheduleTTL := 10 * 365 * 24 * time.Hour
	if conf != nil && conf.GetBusiness() != nil && conf.GetBusiness().GetScheduleTtl() != nil && conf.GetBusiness().GetScheduleTtl().AsDuration() > 0 {
		scheduleTTL = conf.GetBusiness().GetScheduleTtl().AsDuration()
	}
	stream := commonenum.EventStreamConfigFor(commonenum.EventStreamSchedulerTasks)
	return &ScheduledTaskScheduleRepo{
		logger:          logger,
		natsClient:      natsClient,
		js:              natsClient.JetStream(),
		streamName:      stream.Name.String(),
		executeSubject:  commonenum.SchedulerScheduledTaskExecuteSubject,
		consumerDurable: commonenum.EventConsumerSchedulerScheduledTask.String(),
		scheduleTTL:     scheduleTTL,
	}, nil
}

func (r *ScheduledTaskScheduleRepo) Ensure(ctx context.Context) error {
	stream := commonenum.EventStreamConfigFor(commonenum.EventStreamSchedulerTasks)
	return r.natsClient.EnsureStream(ctx, nats.StreamSpec{
		Name:              stream.Name.String(),
		Subjects:          stream.Subjects,
		Retention:         stream.Retention,
		AllowMsgSchedules: stream.AllowMsgSchedules,
		AllowMsgTTL:       stream.AllowMsgTTL,
		AllowRollup:       stream.AllowRollup,
	})
}
func (r *ScheduledTaskScheduleRepo) Schedule(ctx context.Context, req *repo.ScheduledTaskScheduleReq) error {
	if req == nil || req.ScheduledTask == nil || req.Subject == "" || req.Target == "" {
		return fmt.Errorf("scheduled task schedule request is invalid")
	}
	stream, err := r.js.Stream(ctx, r.streamName)
	if err != nil {
		return err
	}
	// 同一个调度源 subject 只保留一份调度定义，更新时先清理旧调度。
	if err = stream.Purge(ctx, jetstream.WithPurgeSubject(req.Subject)); err != nil {
		return err
	}
	body, err := json.Marshal(repo.ScheduledTaskScheduleMessage{
		ScheduledTaskID:      req.ScheduledTask.ID,
		ScheduledTaskVersion: req.ScheduledTask.Version,
		TriggerType:          schedulerenum.TaskTriggerTypeSchedule,
	})
	if err != nil {
		return err
	}
	_, err = r.js.Publish(ctx, req.Subject, body,
		jetstream.WithScheduleCron(req.ScheduledTask.CronSpec),
		jetstream.WithScheduleTarget(fmt.Sprintf("%s.%d", req.Target, req.ScheduledTask.ID)),
		jetstream.WithScheduleTTL(r.scheduleTTL),
	)
	return err
}

func (r *ScheduledTaskScheduleRepo) Cancel(ctx context.Context, subject string) error {
	if subject == "" {
		return nil
	}
	stream, err := r.js.Stream(ctx, r.streamName)
	if err != nil {
		return err
	}
	return stream.Purge(ctx, jetstream.WithPurgeSubject(subject))
}

func (r *ScheduledTaskScheduleRepo) Consume(ctx context.Context, handler func(context.Context, *repo.ScheduledTaskScheduleMessage) (*repo.MessageHandleResult, error)) error {
	if err := r.Ensure(ctx); err != nil {
		return err
	}
	subscription, err := r.natsClient.Consume(ctx, nats.ConsumerSpec{
		Stream:        r.streamName,
		Durable:       r.consumerDurable,
		FilterSubject: r.executeSubject + ".>",
		DeliverPolicy: jetstream.DeliverAllPolicy,
	}, func(ctx context.Context, msg *nats.Message) error {
		payload := &repo.ScheduledTaskScheduleMessage{}
		_ = json.Unmarshal(msg.Data, payload)
		if meta := msg.Metadata; meta != nil {
			if payload.ScheduleKey == "" {
				payload.ScheduleKey = uuid.NewSHA1(
					uuid.NameSpaceOID,
					[]byte(fmt.Sprintf("%s:%d", meta.Stream, meta.Sequence)),
				).String()
			}
			if payload.ScheduledAt.IsZero() {
				payload.ScheduledAt = meta.Timestamp.Truncate(time.Second)
			}
			payload.NumDelivered = meta.NumDelivered
			payload.StreamSequence = meta.Sequence
			stream, streamErr := r.js.Stream(ctx, r.streamName)
			if streamErr == nil {
				// 只补最新策略依赖 subject 最新序列判断是否只补最近一次。
				last, lastErr := stream.GetLastMsgForSubject(ctx, msg.Subject)
				payload.LatestForSubject = lastErr == nil && last != nil && last.Sequence == meta.Sequence
			}
		}
		if payload.ScheduledAt.IsZero() {
			payload.ScheduledAt = time.Now().Truncate(time.Second)
		}
		result, handleErr := handler(ctx, payload)
		if handleErr != nil {
			r.logger.WarnContext(
				ctx,
				"定时任务消息处理失败",
				"error", handleErr,
				"scheduled_task_id", payload.ScheduledTaskID,
				"schedule_key", payload.ScheduleKey,
			)
			return handleErr
		}
		if result == nil || result.Action == schedulerenum.MessageHandleActionComplete {
			return nil
		}
		if result.Action == schedulerenum.MessageHandleActionRetry {
			delay := result.RetryAfter
			if delay <= 0 {
				delay = time.Second
			}
			return msg.NackWithDelay(delay)
		}
		if result.Action == schedulerenum.MessageHandleActionDiscard {
			return msg.Term()
		}
		return nil
	})
	if err != nil {
		return err
	}
	r.subscription = subscription
	return nil
}

func (r *ScheduledTaskScheduleRepo) Stop(ctx context.Context) error {
	if r.subscription != nil {
		if err := r.subscription.Unsubscribe(); err != nil {
			return err
		}
		r.subscription = nil
	}
	return nil
}
