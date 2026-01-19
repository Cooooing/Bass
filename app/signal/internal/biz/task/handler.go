package task

import (
	"github.com/hibiken/asynq"
)

type Handler interface {
	Name() TaskName
	Handler() asynq.HandlerFunc
}
