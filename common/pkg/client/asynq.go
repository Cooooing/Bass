package client

import (
	"common/pkg/constant"
	"common/pkg/cutil/base"
	"common/pkg/cutil/collections/dict"
	"common/pkg/model"
	"context"
	"encoding/json"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
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

func NewAsynqClient(log *log.Helper, redisClient *RedisClient) *AsynqClient {
	return &AsynqClient{
		log:    log,
		Client: asynq.NewClientFromRedisClient(redisClient.Client),
	}
}

type AsynqServer struct {
	log    *log.Helper
	mux    *asynq.ServeMux
	Server *asynq.Server
}

func NewAsynqServer(log *log.Helper, redisClient *RedisClient, tasks dict.Map[constant.TaskName, Handler]) *AsynqServer {
	mux := asynq.NewServeMux()
	for _, t := range tasks.Values() {
		log.Infof("register task: %s", t.Name().String())
		mux.HandleFunc(t.Name().String(), t.Handler())
	}

	server := asynq.NewServerFromRedisClient(redisClient.Client, asynq.Config{
		TaskCheckInterval:        1 * time.Second,
		Logger:                   log,
		LogLevel:                 asynq.InfoLevel,
		ErrorHandler:             NewGlobalErrHandler(tasks),
		DelayedTaskCheckInterval: 1 * time.Second,
	})
	s := &AsynqServer{
		log:    log,
		mux:    mux,
		Server: server,
	}
	return s
}

func (s *AsynqServer) Run() {
	if err := s.Server.Run(s.mux); err != nil {
		s.log.Fatal(err)
	}
}

type GlobalErrHandler struct {
	tasks dict.Map[constant.TaskName, Handler]
}

func NewGlobalErrHandler(tasks dict.Map[constant.TaskName, Handler]) *GlobalErrHandler {
	return &GlobalErrHandler{
		tasks: tasks,
	}
}

func (h *GlobalErrHandler) HandleError(ctx context.Context, task *asynq.Task, err error) {
	if t, ok := h.tasks.Get(constant.TaskName(task.Type())); ok {
		t.ErrHandler(ctx, task, err)
	}
}

type Producer struct {
	client *asynq.Client
}

func NewProducer(client *AsynqClient) *Producer {
	return &Producer{
		client: client.Client,
	}
}

func (p *Producer) EnqueueTask(data *model.Task) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return err
	}
	opts := make([]asynq.Option, 0)
	if data.TaskId != "" {
		opts = append(opts, asynq.TaskID(data.TaskId))
	}
	_, err = p.client.Enqueue(
		asynq.NewTask(
			data.TaskName,
			marshal,
			opts...,
		),
		asynq.MaxRetry(data.MaxRetry),
		asynq.ProcessIn(base.If(data.Delay, data.Interval, 0)),
	)
	if err != nil {
		return err
	}
	return nil
}

func (p *Producer) EnqueueTasks(data []*model.Task) error {
	for _, d := range data {
		err := p.EnqueueTask(d)
		if err != nil {
			return err
		}
	}
	return nil
}
