package task

import (
	commonClient "common/pkg/client"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
)

func ProvideTasks(
	ping *NodePingTaskHandler,
) []Handler {
	return []Handler{
		ping,
	}
}

type AsynqServer struct {
	log    *log.Helper
	server *asynq.Server
	mux    *asynq.ServeMux
}

func NewAsynqServer(log *log.Helper, redisClient *commonClient.RedisClient, tasks []Handler) *AsynqServer {
	mux := asynq.NewServeMux()
	for _, t := range tasks {
		mux.HandleFunc(t.Name().String(), t.Handler())
	}

	server := asynq.NewServerFromRedisClient(redisClient.Client, asynq.Config{
		TaskCheckInterval:        1 * time.Second,
		Logger:                   log,
		LogLevel:                 asynq.InfoLevel,
		ErrorHandler:             nil,
		DelayedTaskCheckInterval: 1 * time.Second,
	})
	return &AsynqServer{
		log:    log,
		mux:    mux,
		server: server,
	}
}

func (s *AsynqServer) Run() {
	go func() {
		if err := s.server.Run(s.mux); err != nil {
			s.log.Fatal(err)
		}
	}()
}
