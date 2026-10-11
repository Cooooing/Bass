package repo

import (
	commonclient "common/pkg/client"
	"time"

	"scheduler/internal/config"

	"github.com/nats-io/nats.go/jetstream"
)

type natsScheduleBootstrap struct {
	js         jetstream.JetStream
	streamName string
	ackWait    time.Duration
	fetchBatch int
	fetchWait  time.Duration
}

func (*natsScheduleBootstrap) initialize(conf *config.Bootstrap, natsClient *commonclient.NatsClient) (*natsScheduleBootstrap, error) {
	js, err := jetstream.New(natsClient.Conn())
	if err != nil {
		return nil, err
	}
	scheduler := conf.GetScheduler()
	streamName := scheduler.GetStreamName()
	if streamName == "" {
		streamName = "SCHEDULER"
	}
	ackWait := time.Minute
	if scheduler.GetAckWait() != nil && scheduler.GetAckWait().AsDuration() > 0 {
		ackWait = scheduler.GetAckWait().AsDuration()
	}
	fetchBatch := int(scheduler.GetFetchBatch())
	if fetchBatch <= 0 {
		fetchBatch = 16
	}
	fetchWait := 2 * time.Second
	if scheduler.GetFetchWait() != nil && scheduler.GetFetchWait().AsDuration() > 0 {
		fetchWait = scheduler.GetFetchWait().AsDuration()
	}
	return &natsScheduleBootstrap{js: js, streamName: streamName, ackWait: ackWait, fetchBatch: fetchBatch, fetchWait: fetchWait}, nil
}
