package task

import (
	commonClient "common/pkg/client"
	"common/pkg/cutil/collections/dict"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
)

type AsynqClient struct {
	log    *log.Helper
	Client *asynq.Client
}

func NewAsynqClient(log *log.Helper, redisClient *commonClient.RedisClient) *AsynqClient {
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

func NewAsynqServer(log *log.Helper, redisClient *commonClient.RedisClient, tasks dict.Map[TaskName, Handler]) *AsynqServer {
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
	go func() {
		if err := s.Server.Run(s.mux); err != nil {
			s.log.Fatal(err)
		}
	}()
}

type GlobalErrHandler struct {
	tasks dict.Map[TaskName, Handler]
}

func NewGlobalErrHandler(tasks dict.Map[TaskName, Handler]) *GlobalErrHandler {
	return &GlobalErrHandler{
		tasks: tasks,
	}
}

func (h *GlobalErrHandler) HandleError(ctx context.Context, task *asynq.Task, err error) {
	if t, ok := h.tasks.Get(TaskName(task.Type())); ok {
		t.ErrHandler(ctx, task, err)
	}
}
