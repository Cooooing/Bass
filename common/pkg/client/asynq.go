package client

import (
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/samber/lo"
)

type Handler interface {
	Name() constant.TaskName
	Handler() asynq.HandlerFunc
	ErrHandler(ctx context.Context, task *asynq.Task, err error)
}

type AsynqClient struct {
	log    *log.Helper
	Client *asynq.Client
}

func NewAsynqClient(log *log.Helper, redisClient *RedisClient) (*AsynqClient, func()) {
	c := &AsynqClient{
		log:    log,
		Client: asynq.NewClientFromRedisClient(redisClient.Client),
	}
	cleanup := func() {
		err := c.Client.Close()
		if err != nil {
			log.Error(err)
		}
	}
	return c, cleanup
}

type AsynqServer struct {
	log    *log.Helper
	mux    *asynq.ServeMux
	Server *asynq.Server
}

func NewAsynqServer(log *log.Helper, redisClient *RedisClient, tasks map[constant.TaskName]Handler) (*AsynqServer, func()) {
	mux := asynq.NewServeMux()
	for _, t := range lo.Values(tasks) {
		log.Infof("register task: %s", t.Name().String())
		mux.HandleFunc(t.Name().String(), t.Handler())
	}

	server := asynq.NewServerFromRedisClient(redisClient.Client, asynq.Config{
		Concurrency:              10,
		TaskCheckInterval:        2 * time.Second,
		DelayedTaskCheckInterval: 2 * time.Second,
		ShutdownTimeout:          30 * time.Second,
		Logger:                   log,
		LogLevel:                 asynq.InfoLevel,
		ErrorHandler:             NewGlobalErrHandler(log, tasks),
	})
	s := &AsynqServer{
		log:    log,
		mux:    mux,
		Server: server,
	}
	return s, s.Server.Shutdown
}

func (s *AsynqServer) Run() {
	if err := s.Server.Run(s.mux); err != nil {
		s.log.Fatal(err)
	}
}

type AsynqScheduler struct {
	log       *log.Helper
	Scheduler *asynq.Scheduler
}

func NewAsynqScheduler(log *log.Helper, redisClient *RedisClient) (*AsynqScheduler, func()) {
	scheduler := asynq.NewSchedulerFromRedisClient(redisClient.Client, &asynq.SchedulerOpts{
		Logger:   log,
		LogLevel: asynq.InfoLevel,
		Location: time.Local,
	})
	s := &AsynqScheduler{
		log:       log,
		Scheduler: scheduler,
	}
	return s, s.Scheduler.Shutdown
}

func (s *AsynqScheduler) Run() {
	if err := s.Scheduler.Run(); err != nil {
		s.log.Fatal(err)
	}
}

type GlobalErrHandler struct {
	Log   *log.Helper
	tasks map[constant.TaskName]Handler
}

func NewGlobalErrHandler(log *log.Helper, tasks map[constant.TaskName]Handler) *GlobalErrHandler {
	return &GlobalErrHandler{
		Log:   log,
		tasks: tasks,
	}
}

func (h *GlobalErrHandler) HandleError(ctx context.Context, task *asynq.Task, err error) {
	h.Log.Errorf("task %s err: %s data: %s", task.Type(), err, string(task.Payload()))
	// 寻找匹配的 handler
	for _, taskName := range lo.Keys(h.tasks) {
		if strings.HasPrefix(task.Type(), string(taskName)) {
			if t, ok := h.tasks[taskName]; ok {
				t.ErrHandler(ctx, task, err)
				return
			}
		}
	}
}

type Producer struct {
	client    *asynq.Client
	scheduler *asynq.Scheduler
}

func NewProducer(client *AsynqClient, scheduler *AsynqScheduler) *Producer {
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
