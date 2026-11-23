package nats

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"slices"
	"sync"
	"time"

	"common/proto/gen/common/v1"

	"github.com/nats-io/nats.go/jetstream"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

const defaultNackDelay = time.Second

type StreamSpec struct {
	Name              string
	Subjects          []string
	Retention         jetstream.RetentionPolicy
	MaxBytes          int64
	AllowMsgSchedules bool
	AllowMsgTTL       bool
	AllowRollup       bool
}

type ConsumerSpec struct {
	Stream         string
	Durable        string
	FilterSubject  string
	FilterSubjects []string
	DeliverPolicy  jetstream.DeliverPolicy
}

type Acknowledgment struct {
	Stream    string
	Sequence  uint64
	Duplicate bool
}

func (c *Client) JSPublish(ctx context.Context, subject string, message *Message) (_ *Acknowledgment, err error) {
	ctx, finish := c.observer.Message(ctx, "publish", subject)
	defer func() { finish(err) }()
	ack, err := c.js.PublishMsg(ctx, c.publishMessage(ctx, subject, message))
	if err != nil {
		return nil, fmt.Errorf("jetstream publish %s: %w", subject, err)
	}
	return &Acknowledgment{Stream: ack.Stream, Sequence: ack.Sequence, Duplicate: ack.Duplicate}, nil
}

func (c *Client) VerifyStream(ctx context.Context, spec StreamSpec) error {
	want, err := c.streamConfig(spec)
	if err != nil {
		return err
	}
	return c.verifyStream(ctx, want)
}

func (c *Client) EnsureStream(ctx context.Context, spec StreamSpec) error {
	want, err := c.streamConfig(spec)
	if err != nil {
		return err
	}
	_, err = c.js.Stream(ctx, spec.Name)
	if errors.Is(err, jetstream.ErrStreamNotFound) {
		_, err = c.js.CreateStream(ctx, want)
		if err == nil {
			return nil
		}
		if _, lookupErr := c.js.Stream(ctx, spec.Name); lookupErr != nil {
			return err
		}
		return c.verifyStream(ctx, want)
	}
	if err != nil {
		return err
	}
	return c.verifyStream(ctx, want)
}

func (c *Client) Consume(ctx context.Context, spec ConsumerSpec, handler MessageHandler) (*Subscription, error) {
	want, consumerConfig, err := c.consumerConfig(spec)
	if err != nil {
		return nil, err
	}
	consumer, err := c.js.Consumer(ctx, spec.Stream, spec.Durable)
	if errors.Is(err, jetstream.ErrConsumerNotFound) {
		consumer, err = c.js.CreateConsumer(ctx, spec.Stream, want)
		if err != nil {
			createErr := err
			consumer, err = c.js.Consumer(ctx, spec.Stream, spec.Durable)
			if err != nil {
				return nil, createErr
			}
		}
	}
	if err != nil {
		return nil, err
	}
	info, err := consumer.Info(ctx)
	if err != nil {
		return nil, err
	}
	if err = c.verifyConsumerConfig(info.Config, want); err != nil {
		return nil, err
	}
	opts := []jetstream.PullConsumeOpt{
		jetstream.PullMaxMessages(int(consumerConfig.GetPullMaxMessages())),
		jetstream.PullExpiry(consumerConfig.GetPullExpiry().AsDuration()),
	}
	pull, err := consumer.Consume(func(raw jetstream.Msg) {
		carrier := propagation.MapCarrier{}
		for key, value := range raw.Headers() {
			if len(value) > 0 {
				carrier.Set(key, value[0])
			}
		}
		message := &Message{Subject: raw.Subject(), Data: raw.Data(), raw: raw}
		if len(raw.Headers()) > 0 {
			message.Header = make(map[string]string, len(raw.Headers()))
			for key, value := range raw.Headers() {
				if len(value) > 0 {
					message.Header[key] = value[0]
				}
			}
		}
		if metadata, metadataErr := raw.Metadata(); metadataErr == nil && metadata != nil {
			message.Metadata = &MessageMetadata{Stream: metadata.Stream, Sequence: metadata.Sequence.Stream, Consumer: metadata.Consumer, NumDelivered: metadata.NumDelivered, Timestamp: metadata.Timestamp}
		}
		go func() {
			messageCtx := otel.GetTextMapPropagator().Extract(ctx, carrier)
			metricSubject := spec.FilterSubject
			if metricSubject == "" {
				metricSubject = message.Subject
			}
			messageCtx, finish := c.observer.Message(messageCtx, "consume", metricSubject)
			finished := make(chan struct{})
			var heartbeat sync.WaitGroup
			heartbeat.Add(1)
			go func() {
				defer heartbeat.Done()
				interval := consumerConfig.GetAckWait().AsDuration() / 2
				ticker := time.NewTicker(interval)
				defer ticker.Stop()
				for {
					select {
					case <-finished:
						return
					case <-ticker.C:
						_ = message.InProgress()
					}
				}
			}()
			var handlerErr error
			defer func() {
				if recovered := recover(); recovered != nil {
					handlerErr = fmt.Errorf("JetStream handler panic: %v\n%s", recovered, debug.Stack())
				}
				close(finished)
				heartbeat.Wait()
				if handlerErr != nil && !message.settled {
					handlerErr = errors.Join(handlerErr, message.NackWithDelay(defaultNackDelay))
				}
				finish(handlerErr)
			}()
			if handler == nil {
				handlerErr = message.Term()
				return
			}
			handlerErr = handler(messageCtx, message)
			if handlerErr == nil && !message.settled {
				handlerErr = message.Ack()
			}
		}()
	}, opts...)
	if err != nil {
		return nil, err
	}
	result := &Subscription{pull: pull}
	c.track(result)
	return result, nil
}

func (*Client) streamConfig(spec StreamSpec) (jetstream.StreamConfig, error) {
	if spec.Name == "" || len(spec.Subjects) == 0 {
		return jetstream.StreamConfig{}, errors.New("stream requires name and subjects")
	}
	maxBytes := spec.MaxBytes
	if maxBytes == 0 {
		maxBytes = 1 << 30
	}
	return jetstream.StreamConfig{
		Name:              spec.Name,
		Subjects:          slices.Clone(spec.Subjects),
		Retention:         spec.Retention,
		MaxBytes:          maxBytes,
		AllowMsgSchedules: spec.AllowMsgSchedules,
		AllowMsgTTL:       spec.AllowMsgTTL,
		AllowRollup:       spec.AllowRollup,
	}, nil
}

func (c *Client) verifyStream(ctx context.Context, want jetstream.StreamConfig) error {
	stream, err := c.js.Stream(ctx, want.Name)
	if err != nil {
		return err
	}
	info, err := stream.Info(ctx)
	if err != nil {
		return err
	}
	actual := info.Config
	switch {
	case actual.Retention != want.Retention:
		return fmt.Errorf("stream %s retention differs", want.Name)
	case actual.MaxBytes != want.MaxBytes:
		return fmt.Errorf("stream %s max bytes differs", want.Name)
	case actual.AllowMsgSchedules != want.AllowMsgSchedules:
		return fmt.Errorf("stream %s scheduled-message setting differs", want.Name)
	case actual.AllowMsgTTL != want.AllowMsgTTL:
		return fmt.Errorf("stream %s message TTL setting differs", want.Name)
	case actual.AllowRollup != want.AllowRollup:
		return fmt.Errorf("stream %s rollup setting differs", want.Name)
	case !slices.Equal(actual.Subjects, want.Subjects):
		return fmt.Errorf("stream %s subjects differ", want.Name)
	}
	return nil
}

func (c *Client) consumerConfig(spec ConsumerSpec) (jetstream.ConsumerConfig, *common.NatsConsumer, error) {
	if spec.Stream == "" || spec.Durable == "" {
		return jetstream.ConsumerConfig{}, nil, errors.New("consumer requires stream and durable")
	}
	if spec.FilterSubject != "" && len(spec.FilterSubjects) > 0 {
		return jetstream.ConsumerConfig{}, nil, errors.New("consumer filters are mutually exclusive")
	}
	if spec.FilterSubject == "" && len(spec.FilterSubjects) == 0 {
		return jetstream.ConsumerConfig{}, nil, errors.New("consumer requires filters")
	}
	consumer, ok := c.conf.GetConsumers()[spec.Durable]
	if !ok || consumer == nil {
		return jetstream.ConsumerConfig{}, nil, fmt.Errorf("consumer %s configuration is required", spec.Durable)
	}
	switch {
	case consumer.GetAckWait() == nil || consumer.GetAckWait().AsDuration() <= 0:
		return jetstream.ConsumerConfig{}, nil, fmt.Errorf("consumer %s ack wait is required", spec.Durable)
	case consumer.GetMaxAckPending() <= 0:
		return jetstream.ConsumerConfig{}, nil, fmt.Errorf("consumer %s max ack pending is required", spec.Durable)
	case consumer.GetPullMaxMessages() <= 0:
		return jetstream.ConsumerConfig{}, nil, fmt.Errorf("consumer %s pull max messages is required", spec.Durable)
	case consumer.GetPullMaxMessages() > consumer.GetMaxAckPending():
		return jetstream.ConsumerConfig{}, nil, fmt.Errorf("consumer %s pull max messages exceeds max ack pending", spec.Durable)
	case consumer.GetPullExpiry() == nil || consumer.GetPullExpiry().AsDuration() <= 0:
		return jetstream.ConsumerConfig{}, nil, fmt.Errorf("consumer %s pull expiry is required", spec.Durable)
	}
	return jetstream.ConsumerConfig{
		Durable:        spec.Durable,
		AckPolicy:      jetstream.AckExplicitPolicy,
		FilterSubject:  spec.FilterSubject,
		FilterSubjects: slices.Clone(spec.FilterSubjects),
		AckWait:        consumer.GetAckWait().AsDuration(),
		MaxAckPending:  int(consumer.GetMaxAckPending()),
		DeliverPolicy:  spec.DeliverPolicy,
	}, consumer, nil
}

func (*Client) verifyConsumerConfig(actual, want jetstream.ConsumerConfig) error {
	switch {
	case actual.AckPolicy != want.AckPolicy:
		return fmt.Errorf("consumer %s ack policy differs", want.Durable)
	case actual.DeliverPolicy != want.DeliverPolicy:
		return fmt.Errorf("consumer %s deliver policy differs", want.Durable)
	case actual.AckWait != want.AckWait:
		return fmt.Errorf("consumer %s ack wait differs", want.Durable)
	case actual.MaxAckPending != want.MaxAckPending:
		return fmt.Errorf("consumer %s max ack pending differs", want.Durable)
	case actual.FilterSubject != want.FilterSubject:
		return fmt.Errorf("consumer %s subject filter differs", want.Durable)
	case !slices.Equal(actual.FilterSubjects, want.FilterSubjects):
		return fmt.Errorf("consumer %s subject filters differ", want.Durable)
	}
	return nil
}
