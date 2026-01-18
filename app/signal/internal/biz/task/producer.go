package task

import (
	"time"

	"github.com/hibiken/asynq"
)

const (
	TaskTypeNodePing = "node_ping"
)

type TaskProducer struct {
	client *asynq.Client
}

func NewNodeTaskProducer(c *asynq.Client) *TaskProducer {
	return &TaskProducer{client: c}
}

func (p *TaskProducer) EnqueueNodePing(node string, delay time.Duration) error {

	task := asynq.NewTask(
		TaskTypeNodePing,
		[]byte(node),
		asynq.TaskID("node-health:"+node),
	)

	_, err := p.client.Enqueue(task, asynq.ProcessIn(delay))
	return err
}
