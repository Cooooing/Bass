package nats

import (
	"common/proto/gen/common"
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"sync"

	gonats "github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type Subscription struct {
	client *Client
	core   *gonats.Subscription
	pull   jetstream.ConsumeContext
	once   sync.Once
	err    error
}

func (s *Subscription) Unsubscribe() error {
	s.once.Do(func() {
		if s.core != nil {
			s.err = s.core.Unsubscribe()
		}
		if s.pull != nil {
			s.pull.Stop()
		}
		if s.client != nil {
			s.client.untrack(s)
		}
	})
	return s.err
}

func (c *Client) Publish(ctx context.Context, subject string, message *Message) (err error) {
	ctx, finish := c.observer.Message(ctx, "publish", subject)
	defer func() { finish(err) }()
	raw := c.publishMessage(ctx, subject, message)
	if err := c.conn.PublishMsg(raw); err != nil {
		return fmt.Errorf("publish %s: %w", subject, err)
	}
	return c.conn.FlushWithContext(ctx)
}

func (c *Client) Subscribe(ctx context.Context, subject string, handler MessageHandler) (*Subscription, error) {
	return c.coreSubscription(c.conn.Subscribe(subject, c.coreMessageHandler(ctx, subject, handler)))
}

func (c *Client) QueueSubscribe(ctx context.Context, subject, queue string, handler MessageHandler) (*Subscription, error) {
	return c.coreSubscription(c.conn.QueueSubscribe(subject, queue, c.coreMessageHandler(ctx, subject, handler)))
}

func (c *Client) publishMessage(ctx context.Context, subject string, message *Message) *gonats.Msg {
	if message == nil {
		message = &Message{}
	}
	raw := &gonats.Msg{Subject: subject, Data: message.Data, Header: make(gonats.Header, len(message.Header)+1)}
	for key, value := range message.Header {
		raw.Header.Set(key, value)
	}
	if message.ID != "" {
		raw.Header.Set(gonats.MsgIdHdr, message.ID)
	}
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	for key, value := range carrier {
		raw.Header.Set(key, value)
	}
	return raw
}

func (c *Client) coreMessageHandler(ctx context.Context, subject string, handler MessageHandler) gonats.MsgHandler {
	return func(raw *gonats.Msg) {
		carrier := propagation.MapCarrier{}
		for key, value := range raw.Header {
			if len(value) > 0 {
				carrier.Set(key, value[0])
			}
		}
		messageCtx := otel.GetTextMapPropagator().Extract(ctx, carrier)
		messageCtx, finish := c.observer.Message(messageCtx, "consume", subject)
		var handlerErr error
		defer func() {
			if recovered := recover(); recovered != nil {
				handlerErr = fmt.Errorf("core NATS handler panic: %v\n%s", recovered, debug.Stack())
			}
			finish(handlerErr)
		}()
		if handler == nil {
			return
		}
		message := &Message{Subject: raw.Subject, Data: raw.Data}
		if len(raw.Header) > 0 {
			message.Header = make(map[string]string, len(raw.Header))
			for key, value := range raw.Header {
				if len(value) > 0 {
					message.Header[key] = value[0]
				}
			}
		}
		handlerErr = handler(messageCtx, message)
	}
}

func (c *Client) coreSubscription(sub *gonats.Subscription, err error) (*Subscription, error) {
	if err != nil {
		return nil, err
	}
	if err = configureCoreSubscription(sub, c.conf.GetCoreSubscription()); err != nil {
		_ = sub.Unsubscribe()
		return nil, err
	}
	result := &Subscription{core: sub}
	c.track(result)
	return result, nil
}

func configureCoreSubscription(sub *gonats.Subscription, config *common.NatsCoreSubscription) error {
	if config == nil {
		return errors.New("core subscription pending limits are required")
	}
	if config.GetMaxPendingMessages() <= 0 || config.GetMaxPendingBytes() <= 0 {
		return errors.New("core subscription pending limits are required")
	}
	return sub.SetPendingLimits(int(config.GetMaxPendingMessages()), int(config.GetMaxPendingBytes()))
}
