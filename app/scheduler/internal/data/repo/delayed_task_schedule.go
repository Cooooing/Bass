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

type DelayedTaskScheduleRepo struct {
	logger          *slog.Logger
	natsClient      *nats.Client
	js              jetstream.JetStream
	streamName      string
	executeSubject  string
	consumerDurable string
	scheduleTTL     time.Duration
	subscription    *nats.Subscription
}

func NewDelayedTaskScheduleRepo(
	logger *slog.Logger,
	natsClient *nats.Client,
	conf *config.Bootstrap,
) (repo.DelayedTaskScheduleRepo, error) {
	if conf == nil || conf.GetBusiness() == nil {
		return nil, fmt.Errorf("scheduler business config is required")
	}
	scheduleTTL := 10 * 365 * 24 * time.Hour
	if conf != nil && conf.GetBusiness() != nil && conf.GetBusiness().GetScheduleTtl() != nil && conf.GetBusiness().GetScheduleTtl().AsDuration() > 0 {
		scheduleTTL = conf.GetBusiness().GetScheduleTtl().AsDuration()
	}
	stream := commonenum.EventStreamConfigFor(commonenum.EventStreamSchedulerTasks)
	return &DelayedTaskScheduleRepo{
		logger:          logger,
		natsClient:      natsClient,
		js:              natsClient.JetStream(),
		streamName:      stream.Name.String(),
		executeSubject:  commonenum.SchedulerDelayedTaskExecuteSubject,
		consumerDurable: commonenum.EventConsumerSchedulerDelayedTask.String(),
		scheduleTTL:     scheduleTTL,
	}, nil
}

func (r *DelayedTaskScheduleRepo) Ensure(ctx context.Context) error {
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

func (r *DelayedTaskScheduleRepo) Schedule(ctx context.Context, req *repo.DelayedTaskScheduleReq) error {
	if req == nil || req.Record == nil || req.Subject == "" || req.Target == "" {
		return fmt.Errorf("delayed task schedule request is invalid")
	}
	stream, err := r.js.Stream(ctx, r.streamName)
	if err != nil {
		return err
	}
	// 延迟执行记录按独立调度源 subject 注册，幂等重建时先清理旧调度。
	if err = stream.Purge(ctx, jetstream.WithPurgeSubject(req.Subject)); err != nil {
		return err
	}
	body, err := json.Marshal(repo.DelayedTaskScheduleMessage{
		ExecutionRecordID: req.Record.ID,
		TriggerType:       req.Record.TriggerType,
	})
	if err != nil {
		return err
	}
	opts := []jetstream.PublishOpt{
		jetstream.WithScheduleAt(req.Record.ScheduledAt),
		jetstream.WithScheduleTarget(req.Target),
		jetstream.WithScheduleTTL(r.scheduleTTL),
	}
	_, err = r.js.Publish(ctx, req.Subject, body, opts...)
	return err
}

func (r *DelayedTaskScheduleRepo) Publish(ctx context.Context, req *repo.DelayedTaskScheduleReq) error {
	if req == nil || req.Record == nil || req.Target == "" {
		return fmt.Errorf("delayed task publish request is invalid")
	}
	body, err := json.Marshal(repo.DelayedTaskScheduleMessage{
		ExecutionRecordID: req.Record.ID,
		TriggerType:       req.Record.TriggerType,
	})
	if err != nil {
		return err
	}
	_, err = r.js.Publish(ctx, req.Target, body)
	return err
}

func (r *DelayedTaskScheduleRepo) Cancel(ctx context.Context, subject string) error {
	if subject == "" {
		return nil
	}
	stream, err := r.js.Stream(ctx, r.streamName)
	if err != nil {
		return err
	}
	return stream.Purge(ctx, jetstream.WithPurgeSubject(subject))
}

func (r *DelayedTaskScheduleRepo) Consume(ctx context.Context, handler func(context.Context, *repo.DelayedTaskScheduleMessage) (*repo.MessageHandleResult, error)) error {
	if err := r.Ensure(ctx); err != nil {
		return err
	}
	subscription, err := r.natsClient.Consume(ctx, nats.ConsumerSpec{
		Stream:        r.streamName,
		Durable:       r.consumerDurable,
		FilterSubject: r.executeSubject,
		DeliverPolicy: jetstream.DeliverAllPolicy,
	}, func(ctx context.Context, msg *nats.Message) error {
		payload := &repo.DelayedTaskScheduleMessage{}
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
		}
		if payload.ScheduledAt.IsZero() {
			payload.ScheduledAt = time.Now().Truncate(time.Second)
		}
		result, handleErr := handler(ctx, payload)
		if handleErr != nil {
			r.logger.WarnContext(
				ctx,
				"延迟任务消息处理失败",
				"error", handleErr,
				"execution_record_id", payload.ExecutionRecordID,
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

func (r *DelayedTaskScheduleRepo) Stop(ctx context.Context) error {
	if r.subscription != nil {
		if err := r.subscription.Unsubscribe(); err != nil {
			return err
		}
		r.subscription = nil
	}
	return nil
}
