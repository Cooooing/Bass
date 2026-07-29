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

type ScheduledTaskScheduleNatsRepo struct {
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

func NewScheduledTaskScheduleNatsRepo(
	logger *slog.Logger,
	conf *config.Bootstrap,
	natsClient *commonclient.NatsClient,
) (repo.ScheduledTaskScheduleRepo, error) {
	js, err := jetstream.New(natsClient.Conn())
	if err != nil {
		return nil, err
	}
	s := conf.GetScheduler()
	streamName := s.GetStreamName()
	if streamName == "" {
		streamName = "SCHEDULER"
	}
	schedulePrefix := s.GetScheduledTaskScheduleSubjectPrefix()
	if schedulePrefix == "" {
		schedulePrefix = "scheduler.schedule.scheduled_task"
	}
	executeSubject := s.GetScheduledTaskExecuteSubject()
	if executeSubject == "" {
		executeSubject = "scheduler.execute.scheduled_task"
	}
	consumerDurable := s.GetScheduledTaskConsumerDurable()
	if consumerDurable == "" {
		consumerDurable = "scheduler_scheduled_task_execution"
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
	delayedSchedulePrefix := s.GetDelayedTaskScheduleSubjectPrefix()
	if delayedSchedulePrefix == "" {
		delayedSchedulePrefix = "scheduler.schedule.delayed_task_execution"
	}
	delayedExecuteSubject := s.GetDelayedTaskExecuteSubject()
	if delayedExecuteSubject == "" {
		delayedExecuteSubject = "scheduler.execute.delayed_task"
	}
	return &ScheduledTaskScheduleNatsRepo{
		logger:          logger,
		js:              js,
		streamName:      streamName,
		schedulePrefix:  schedulePrefix,
		executeSubject:  executeSubject,
		consumerDurable: consumerDurable,
		ackWait:         ackWait,
		fetchBatch:      fetchBatch,
		fetchWait:       fetchWait,
		scheduleTTL:     24 * 365 * 10 * time.Hour,
		allSubjects: []string{
			schedulePrefix + ".>",
			executeSubject + ".>",
			delayedSchedulePrefix + ".>",
			delayedExecuteSubject,
		},
	}, nil
}

func (r *ScheduledTaskScheduleNatsRepo) Ensure(ctx context.Context) error {
	stream, err := r.js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:              r.streamName,
		Subjects:          r.allSubjects,
		AllowMsgSchedules: true,
		AllowMsgTTL:       true,
		AllowRollup:       true,
	})
	if err != nil {
		return err
	}
	// 消费者重建后只接收新投递，避免旧执行消息干扰当前调度配置。
	if err = stream.Purge(ctx, jetstream.WithPurgeSubject(r.executeSubject+".>")); err != nil {
		return err
	}
	if err = r.js.DeleteConsumer(ctx, r.streamName, r.consumerDurable); err != nil && !errors.Is(err, jetstream.ErrConsumerNotFound) {
		return err
	}
	consumer, err := r.js.CreateConsumer(ctx, r.streamName, jetstream.ConsumerConfig{
		Durable:       r.consumerDurable,
		Name:          r.consumerDurable,
		DeliverPolicy: jetstream.DeliverNewPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
		AckWait:       r.ackWait,
		FilterSubject: r.executeSubject + ".>",
		MaxAckPending: r.fetchBatch * 4,
	})
	if err != nil {
		return err
	}
	r.consumer = consumer
	return nil
}

func (r *ScheduledTaskScheduleNatsRepo) Schedule(ctx context.Context, req *repo.ScheduledTaskScheduleReq) error {
	if req == nil || req.ScheduledTask == nil || req.Subject == "" || req.Target == "" {
		return fmt.Errorf("scheduled task schedule request is invalid")
	}
	stream, err := r.js.Stream(ctx, r.streamName)
	if err != nil {
		return err
	}
	// 同一个调度源主题只保留一份调度定义，更新时先清理旧调度。
	if err = stream.Purge(ctx, jetstream.WithPurgeSubject(req.Subject)); err != nil {
		return err
	}
	body, err := json.Marshal(repo.ScheduledTaskScheduleMessage{
		ScheduledTaskID:      req.ScheduledTask.ID,
		ScheduledTaskTitle:   req.ScheduledTask.Title,
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

func (r *ScheduledTaskScheduleNatsRepo) Cancel(ctx context.Context, subject string) error {
	if subject == "" {
		return nil
	}
	stream, err := r.js.Stream(ctx, r.streamName)
	if err != nil {
		return err
	}
	return stream.Purge(ctx, jetstream.WithPurgeSubject(subject))
}

func (r *ScheduledTaskScheduleNatsRepo) Consume(ctx context.Context, handler func(context.Context, *repo.ScheduledTaskScheduleMessage) (*repo.MessageHandleResult, error)) error {
	if r.consumer == nil {
		if err := r.Ensure(ctx); err != nil {
			return err
		}
	}
	consume, err := r.consumer.Consume(func(msg jetstream.Msg) {
		payload := &repo.ScheduledTaskScheduleMessage{}
		_ = json.Unmarshal(msg.Data(), payload)
		meta, _ := msg.Metadata()
		if meta != nil {
			payload.ScheduleKey = uuid.NewSHA1(
				uuid.NameSpaceOID,
				[]byte(fmt.Sprintf("%s:%d", meta.Stream, meta.Sequence.Stream)),
			).String()
			payload.ScheduledAt = meta.Timestamp.Truncate(time.Second)
			payload.NumDelivered = meta.NumDelivered
			payload.StreamSequence = meta.Sequence.Stream
			stream, streamErr := r.js.Stream(ctx, r.streamName)
			if streamErr == nil {
				// 只补最新策略依赖主题最新序列判断是否只补最近一次。
				last, lastErr := stream.GetLastMsgForSubject(ctx, msg.Subject())
				payload.LatestForSubject = lastErr == nil && last != nil && last.Sequence == meta.Sequence.Stream
			}
		}
		if payload.ScheduledAt.IsZero() {
			payload.ScheduledAt = time.Now().Truncate(time.Second)
		}
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
				"定时任务消息处理失败",
				"error", handleErr,
				"scheduled_task_id", payload.ScheduledTaskID,
				"scheduled_task_title", payload.ScheduledTaskTitle,
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

func (r *ScheduledTaskScheduleNatsRepo) Stop(ctx context.Context) error {
	if r.consume != nil {
		r.consume.Stop()
	}
	return nil
}
