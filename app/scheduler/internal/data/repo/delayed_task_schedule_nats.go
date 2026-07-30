package repo

import (
	commonclient "common/pkg/client"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"scheduler/internal/biz/repo"
	"scheduler/internal/config"
	schedulerenum "scheduler/internal/enum"
	"time"

	"github.com/google/uuid"

	"github.com/nats-io/nats.go/jetstream"
)

type DelayedTaskScheduleNatsRepo struct {
	logger          *slog.Logger
	js              jetstream.JetStream
	streamName      string
	schedulePrefix  string
	executeSubject  string
	consumerDurable string
	ackWait         time.Duration
	fetchBatch      int
	fetchWait       time.Duration
	scheduleTTL     time.Duration
	allSubjects     []string
	consumer        jetstream.Consumer
	consume         jetstream.ConsumeContext
}

func NewDelayedTaskScheduleNatsRepo(
	logger *slog.Logger,
	conf *config.Bootstrap,
	natsClient *commonclient.NatsClient,
) (repo.DelayedTaskScheduleRepo, error) {
	js, err := jetstream.New(natsClient.Conn())
	if err != nil {
		return nil, err
	}
	s := conf.GetScheduler()
	streamName := s.GetStreamName()
	if streamName == "" {
		streamName = "SCHEDULER"
	}
	schedulePrefix := s.GetDelayedTaskScheduleSubjectPrefix()
	if schedulePrefix == "" {
		schedulePrefix = "scheduler.schedule.delayed_task_execution"
	}
	executeSubject := s.GetDelayedTaskExecuteSubject()
	if executeSubject == "" {
		executeSubject = "scheduler.execute.delayed_task"
	}
	consumerDurable := s.GetDelayedTaskConsumerDurable()
	if consumerDurable == "" {
		consumerDurable = "scheduler_delayed_task_execution"
	}
	ackWait := time.Minute
	if s.GetAckWait() != nil && s.GetAckWait().AsDuration() > 0 {
		ackWait = s.GetAckWait().AsDuration()
	}
	fetchBatch := int(s.GetFetchBatch())
	if fetchBatch <= 0 {
		fetchBatch = 16
	}
	fetchWait := 2 * time.Second
	if s.GetFetchWait() != nil && s.GetFetchWait().AsDuration() > 0 {
		fetchWait = s.GetFetchWait().AsDuration()
	}
	scheduledSchedulePrefix := s.GetScheduledTaskScheduleSubjectPrefix()
	if scheduledSchedulePrefix == "" {
		scheduledSchedulePrefix = "scheduler.schedule.scheduled_task"
	}
	scheduledExecuteSubject := s.GetScheduledTaskExecuteSubject()
	if scheduledExecuteSubject == "" {
		scheduledExecuteSubject = "scheduler.execute.scheduled_task"
	}
	return &DelayedTaskScheduleNatsRepo{
		logger:          logger,
		js:              js,
		streamName:      streamName,
		schedulePrefix:  schedulePrefix,
		executeSubject:  executeSubject,
		consumerDurable: consumerDurable,
		ackWait:         ackWait,
		fetchBatch:      fetchBatch,
		fetchWait:       fetchWait,
		scheduleTTL:     10 * 365 * 24 * time.Hour,
		allSubjects: []string{
			scheduledSchedulePrefix + ".>",
			scheduledExecuteSubject + ".>",
			schedulePrefix + ".>",
			executeSubject,
		},
	}, nil
}

func (r *DelayedTaskScheduleNatsRepo) Ensure(ctx context.Context) error {
	if _, err := r.js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:              r.streamName,
		Subjects:          r.allSubjects,
		AllowMsgSchedules: true,
		AllowMsgTTL:       true,
		AllowRollup:       true,
	}); err != nil {
		return err
	}
	if err := r.js.DeleteConsumer(ctx, r.streamName, r.consumerDurable); err != nil && !errors.Is(err, jetstream.ErrConsumerNotFound) {
		return err
	}
	consumer, err := r.js.CreateConsumer(ctx, r.streamName, jetstream.ConsumerConfig{
		Durable:       r.consumerDurable,
		Name:          r.consumerDurable,
		DeliverPolicy: jetstream.DeliverNewPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
		AckWait:       r.ackWait,
		FilterSubject: r.executeSubject,
		MaxAckPending: r.fetchBatch * 4,
	})
	if err != nil {
		return err
	}
	r.consumer = consumer
	return nil
}

func (r *DelayedTaskScheduleNatsRepo) Schedule(ctx context.Context, req *repo.DelayedTaskScheduleReq) error {
	if req == nil || req.DelayedTask == nil || req.Record == nil || req.Subject == "" || req.Target == "" {
		return fmt.Errorf("delayed task schedule request is invalid")
	}
	stream, err := r.js.Stream(ctx, r.streamName)
	if err != nil {
		return err
	}
	// 延迟执行记录按独立调度源主题注册，幂等重建时先清理旧调度。
	if err = stream.Purge(ctx, jetstream.WithPurgeSubject(req.Subject)); err != nil {
		return err
	}
	body, err := json.Marshal(repo.DelayedTaskScheduleMessage{
		ExecutionRecordID: req.Record.ID,
		DelayedTaskKey:    req.DelayedTask.TaskKey,
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

func (r *DelayedTaskScheduleNatsRepo) Cancel(ctx context.Context, subject string) error {
	if subject == "" {
		return nil
	}
	stream, err := r.js.Stream(ctx, r.streamName)
	if err != nil {
		return err
	}
	return stream.Purge(ctx, jetstream.WithPurgeSubject(subject))
}

func (r *DelayedTaskScheduleNatsRepo) Consume(ctx context.Context, handler func(context.Context, *repo.DelayedTaskScheduleMessage) (*repo.MessageHandleResult, error)) error {
	if r.consumer == nil {
		if err := r.Ensure(ctx); err != nil {
			return err
		}
	}
	consume, err := r.consumer.Consume(func(msg jetstream.Msg) {
		payload := &repo.DelayedTaskScheduleMessage{}
		_ = json.Unmarshal(msg.Data(), payload)
		meta, _ := msg.Metadata()
		if meta != nil {
			payload.ScheduleKey = uuid.NewSHA1(
				uuid.NameSpaceOID,
				[]byte(fmt.Sprintf("%s:%d", meta.Stream, meta.Sequence.Stream)),
			).String()
			payload.ScheduledAt = meta.Timestamp.Truncate(time.Second)
			payload.NumDelivered = meta.NumDelivered
		}
		if payload.ScheduledAt.IsZero() {
			payload.ScheduledAt = time.Now().Truncate(time.Second)
		}
		// 长任务处理期间定期续约 NATS 确认等待时间，避免确认超时后误重投。
		done := make(chan struct{})
		go func() {
			interval := r.ackWait / 2
			if interval <= 0 {
				interval = time.Second
			}
			ticker := time.NewTicker(interval)
			defer ticker.Stop()
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					_ = msg.InProgress()
				}
			}
		}()
		result, handleErr := handler(ctx, payload)
		close(done)
		if handleErr != nil {
			r.logger.WarnContext(
				ctx,
				"延迟任务消息处理失败",
				"error", handleErr,
				"execution_record_id", payload.ExecutionRecordID,
				"delayed_task_key", payload.DelayedTaskKey,
				"schedule_key", payload.ScheduleKey,
			)
			delay := r.ackWait / 2
			if delay <= 0 {
				delay = time.Second
			}
			_ = msg.NakWithDelay(delay)
			return
		}
		if result == nil || result.Action == schedulerenum.MessageHandleActionComplete {
			_ = msg.Ack()
			return
		}
		if result.Action == schedulerenum.MessageHandleActionRetry {
			delay := result.RetryAfter
			if delay <= 0 {
				delay = time.Second
			}
			_ = msg.NakWithDelay(delay)
			return
		}
		if result.Action == schedulerenum.MessageHandleActionDiscard {
			_ = msg.Term()
			return
		}
		_ = msg.Ack()
	}, jetstream.PullMaxMessages(r.fetchBatch), jetstream.PullExpiry(r.fetchWait))
	if err != nil {
		return err
	}
	r.consume = consume
	return nil
}

func (r *DelayedTaskScheduleNatsRepo) Stop(ctx context.Context) error {
	if r.consume != nil {
		r.consume.Stop()
	}
	return nil
}
