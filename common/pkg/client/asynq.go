package client

import (
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/hibiken/asynq"
	"github.com/samber/lo"
)

type Handler interface {
	Name() constant.TaskName
	Handler() asynq.HandlerFunc
	ErrHandler(ctx context.Context, task *asynq.Task, err error)
}

type AsynqClient struct {
	logger *slog.Logger
	Client *asynq.Client
}

func NewAsynqClient(
	logger *slog.Logger,
	redisClient *RedisClient,
) (*AsynqClient, func()) {
	c := &AsynqClient{
		logger: logger,
		Client: asynq.NewClientFromRedisClient(redisClient.Client),
	}
	cleanup := func() {
		if err := c.Client.Close(); err != nil {
			c.logger.Error("asynq client close failed", constant.LogFieldKind, constant.LogKindAsynq, constant.LogFieldErr, err)
		}
	}
	return c, cleanup
}

type AsynqServer struct {
	logger *slog.Logger
	mux    *asynq.ServeMux
	Server *asynq.Server
}

func NewAsynqServer(
	logger *slog.Logger,
	redisClient *RedisClient,
	tasks map[constant.TaskName]Handler,
) (*AsynqServer, func()) {
	mux := asynq.NewServeMux()
	for _, t := range lo.Values(tasks) {
		logger.Info("register task", constant.LogFieldKind, constant.LogKindAsynq, "task", t.Name().String())
		mux.HandleFunc(t.Name().String(), t.Handler())
	}

	adapter := &asynqSlogLogger{
		logger: logger,
	}
	server := asynq.NewServerFromRedisClient(redisClient.Client, asynq.Config{
		Concurrency:              10,
		TaskCheckInterval:        2 * time.Second,
		DelayedTaskCheckInterval: 2 * time.Second,
		ShutdownTimeout:          30 * time.Second,
		Logger:                   adapter,
		LogLevel:                 asynq.InfoLevel,
		ErrorHandler:             NewGlobalErrHandler(logger, tasks),
	})
	s := &AsynqServer{
		logger: logger,
		mux:    mux,
		Server: server,
	}
	return s, s.Server.Shutdown
}

func (s *AsynqServer) Run() {
	if err := s.Server.Run(s.mux); err != nil {
		s.logger.Error("asynq server run failed", constant.LogFieldKind, constant.LogKindAsynq, constant.LogFieldErr, err)
		os.Exit(1)
	}
}

type AsynqScheduler struct {
	logger    *slog.Logger
	Scheduler *asynq.Scheduler
}

func NewAsynqScheduler(
	logger *slog.Logger,
	redisClient *RedisClient,
) (*AsynqScheduler, func()) {
	scheduler := asynq.NewSchedulerFromRedisClient(redisClient.Client, &asynq.SchedulerOpts{
		Logger: &asynqSlogLogger{
			logger: logger,
		},
		LogLevel: asynq.InfoLevel,
		Location: time.Local,
	})
	s := &AsynqScheduler{
		logger:    logger,
		Scheduler: scheduler,
	}
	return s, s.Scheduler.Shutdown
}

func (s *AsynqScheduler) Run() {
	if err := s.Scheduler.Run(); err != nil {
		s.logger.Error("asynq scheduler run failed", constant.LogFieldKind, constant.LogKindAsynq, constant.LogFieldErr, err)
		os.Exit(1)
	}
}

type GlobalErrHandler struct {
	logger *slog.Logger
	tasks  map[constant.TaskName]Handler
}

func NewGlobalErrHandler(
	logger *slog.Logger,
	tasks map[constant.TaskName]Handler,
) *GlobalErrHandler {
	return &GlobalErrHandler{
		logger: logger,
		tasks:  tasks,
	}
}

func (h *GlobalErrHandler) HandleError(ctx context.Context, task *asynq.Task, err error) {
	h.logger.ErrorContext(ctx, "task failed", constant.LogFieldKind, constant.LogKindAsynq, "task", task.Type(), "payload", string(task.Payload()), constant.LogFieldErr, err)
	for _, taskName := range lo.Keys(h.tasks) {
		if strings.HasPrefix(task.Type(), string(taskName)) {
			if t, ok := h.tasks[taskName]; ok {
				t.ErrHandler(ctx, task, err)
				return
			}
		}
	}
}

type asynqSlogLogger struct {
	logger *slog.Logger
}

func (l *asynqSlogLogger) Debug(args ...interface{}) {
	l.logger.Debug(fmt.Sprint(args...))
}

func (l *asynqSlogLogger) Info(args ...interface{}) {
	l.logger.Info(fmt.Sprint(args...))
}

func (l *asynqSlogLogger) Warn(args ...interface{}) {
	l.logger.Warn(fmt.Sprint(args...))
}

func (l *asynqSlogLogger) Error(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
}

func (l *asynqSlogLogger) Fatal(args ...interface{}) {
	l.logger.Error(fmt.Sprint(args...))
	os.Exit(1)
}

type Producer struct {
	client    *asynq.Client
	scheduler *asynq.Scheduler
}

func NewProducer(
	client *AsynqClient,
	scheduler *AsynqScheduler,
) *Producer {
	return &Producer{
		client:    client.Client,
		scheduler: scheduler.Scheduler,
	}
}

func (p *Producer) EnqueueContextTask(ctx context.Context, data *model.Task) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	opts := make([]asynq.Option, 0)
	if data.TaskId != "" {
		opts = append(opts, asynq.TaskID(data.TaskId))
	}
	_, err = p.client.EnqueueContext(
		ctx,
		asynq.NewTask(
			data.TaskName,
			marshal,
			opts...,
		),
		asynq.MaxRetry(data.MaxRetry),
		asynq.ProcessIn(util.If(data.Delay, data.Interval, 0)),
		asynq.Retention(data.Retention),
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) EnqueueContextTasks(ctx context.Context, data []*model.Task) error {
	for _, d := range data {
		err := p.EnqueueContextTask(ctx, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Producer) RegisterTask(data *model.Task) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	opts := make([]asynq.Option, 0)
	if data.TaskId != "" {
		opts = append(opts, asynq.TaskID(data.TaskId))
	}
	_, err = p.scheduler.Register(
		data.Cronspec,
		asynq.NewTask(
			data.TaskName,
			marshal,
			opts...,
		),
		asynq.MaxRetry(data.MaxRetry),
		asynq.Retention(data.Retention),
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) RegisterTasks(data []*model.Task) error {
	for _, d := range data {
		err := p.RegisterTask(d)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Producer) UnregisterTask(taskName string) error {
	err := p.scheduler.Unregister(taskName)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) UnregisterTasks(taskNames []string) error {
	for _, taskName := range taskNames {
		err := p.UnregisterTask(taskName)
		if err != nil {
			return err
		}
	}
	return nil
}
