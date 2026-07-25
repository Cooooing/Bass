package client

import (
	"common/pkg/constant"
	"common/proto/gen/common"
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Message struct {
	Subject string
	Data    []byte
	Header  map[string]string
	Ack     func() error
	Nack    func() error
}

type MessageHandler func(ctx context.Context, msg *Message) error

type Publisher interface {
	Publish(ctx context.Context, subject string, msg *Message) error
	Close() error
}

type Subscriber interface {
	Subscribe(ctx context.Context, subject string, handler MessageHandler) (Unsubscriber, error)
	QueueSubscribe(ctx context.Context, subject, queue string, handler MessageHandler) (Unsubscriber, error)
	Close() error
}

type Unsubscriber interface {
	Unsubscribe() error
}

type NatsClient struct {
	conf    *common.Nats
	conn    *nats.Conn
	js      nats.JetStreamContext
	mu      sync.RWMutex
	closed  bool
	logger  *slog.Logger
	service string
	tracer  oteltrace.Tracer
	subs    []*nats.Subscription
}

func NewNatsClient(
	logger *slog.Logger,
	conf *common.Nats,
	observer *Observer,
) (*NatsClient, func(), error) {
	if observer == nil {
		observer = NewObservability(logger, nil)
	}
	if conf == nil {
		conf = &common.Nats{}
	}
	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}
	if conf.Port == 0 {
		conf.Port = 4222
	}
	if conf.Name == "" {
		conf.Name = "nats-client"
	}
	if conf.GetTimeout() == nil || conf.GetTimeout().AsDuration() == 0 {
		conf.Timeout = durationpb.New(5 * time.Second)
	}
	if conf.MaxReconnects == 0 {
		conf.MaxReconnects = 10
	}
	if conf.GetReconnectWait() == nil || conf.GetReconnectWait().AsDuration() == 0 {
		conf.ReconnectWait = durationpb.New(2 * time.Second)
	}
	if conf.GetPingInterval() == nil || conf.GetPingInterval().AsDuration() == 0 {
		conf.PingInterval = durationpb.New(30 * time.Second)
	}
	if conf.GetFlusherTimeout() == nil || conf.GetFlusherTimeout().AsDuration() == 0 {
		conf.FlusherTimeout = durationpb.New(5 * time.Second)
	}

	opts := []nats.Option{
		nats.Name(conf.Name),
		nats.MaxReconnects(int(conf.MaxReconnects)),
		nats.ReconnectWait(conf.ReconnectWait.AsDuration()),
		nats.PingInterval(conf.PingInterval.AsDuration()),
		nats.FlusherTimeout(conf.FlusherTimeout.AsDuration()),
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			logger.Warn("nats disconnected", constant.LogFieldKind, constant.LogKindMessage, constant.LogFieldErr, err)
		}),
		nats.ReconnectHandler(func(c *nats.Conn) {
			logger.Info("nats reconnected", constant.LogFieldKind, constant.LogKindMessage, constant.LogFieldAddress, c.ConnectedUrl())
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			logger.Info("nats connection closed", constant.LogFieldKind, constant.LogKindMessage)
		}),
		nats.ErrorHandler(func(_ *nats.Conn, sub *nats.Subscription, err error) {
			subject := "unknown"
			if sub != nil {
				subject = sub.Subject
			}
			logger.Error("nats error", constant.LogFieldKind, constant.LogKindMessage, constant.LogFieldSubject, subject, constant.LogFieldErr, err)
		}),
	}
	if conf.User != "" || conf.Password != "" {
		opts = append(opts, nats.UserInfo(conf.User, conf.Password))
	}

	address := fmt.Sprintf("nats://%s:%d", conf.Host, conf.Port)
	nc, err := nats.Connect(address, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("nats connect [%s]: %w", address, err)
	}

	service := observer.Service()
	if service == "" {
		service = conf.Name
	}
	if service == "" {
		service = "unknown"
	}
	c := &NatsClient{
		conf:    conf,
		conn:    nc,
		logger:  logger,
		service: service,
		tracer:  otel.Tracer(service + ".message"),
	}

	if conf.EnableJetStream {
		js, err := nc.JetStream()
		if err != nil {
			nc.Close()
			return nil, nil, fmt.Errorf("nats jetstream: %w", err)
		}
		c.js = js
		logger.Info("jetstream enabled", constant.LogFieldKind, constant.LogKindMessage, "domain", conf.JetStreamDomain, "stream", conf.StreamName)
	}

	logger.Info("nats connected", constant.LogFieldKind, constant.LogKindMessage, constant.LogFieldAddress, nc.ConnectedUrl())
	return c, func() {
		if err := c.Close(); err != nil {
			logger.Error("nats close failed", constant.LogFieldKind, constant.LogKindMessage, constant.LogFieldErr, err)
		}
	}, nil
}

func (c *NatsClient) Conn() *nats.Conn {
	return c.conn
}

func (c *NatsClient) JetStream() nats.JetStreamContext {
	return c.js
}

func (c *NatsClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}
	c.closed = true

	for _, sub := range c.subs {
		if sub.IsValid() {
			if err := sub.Drain(); err != nil {
				c.logger.Error("drain subscription failed", constant.LogFieldKind, constant.LogKindMessage, constant.LogFieldErr, err)
			}
		}
	}

	if c.conn != nil {
		_ = c.conn.Drain()
		time.Sleep(500 * time.Millisecond)
		c.conn.Close()
	}

	c.logger.Info("nats client closed", constant.LogFieldKind, constant.LogKindMessage)
	return nil
}

func (c *NatsClient) Publish(ctx context.Context, subject string, msg *Message) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	start := time.Now()
	subjectLabel := subject
	if strings.HasPrefix(subjectLabel, "push.node.") {
		subjectLabel = "push.node.*"
	} else if strings.Count(subjectLabel, ".") > 2 {
		subjectLabel = "dynamic"
	}
	ctx, span := c.tracer.Start(ctx, "nats publish "+subjectLabel, oteltrace.WithSpanKind(oteltrace.SpanKindProducer), oteltrace.WithAttributes(
		attribute.String("messaging.system", "nats"),
		attribute.String("messaging.destination.name", subject),
		attribute.String("messaging.operation", "publish"),
	))
	defer func() {
		status := "ok"
		level := slog.LevelInfo
		if err != nil {
			status = "error"
			level = slog.LevelWarn
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
		latency := time.Since(start)
		MessageRequestsTotal.WithLabelValues(c.service, "publish", subjectLabel, status).Inc()
		MessageRequestDurationSeconds.WithLabelValues(c.service, "publish", subjectLabel, status).Observe(latency.Seconds())
		attrs := []slog.Attr{
			slog.String(constant.LogFieldKind, constant.LogKindMessage),
			slog.String(constant.LogFieldDirection, "publish"),
			slog.String(constant.LogFieldSubject, subject),
			slog.String(constant.LogFieldStatus, status),
			slog.Int64(constant.LogFieldLatencyMS, latency.Milliseconds()),
		}
		if err != nil {
			attrs = append(attrs, slog.Any(constant.LogFieldErr, err))
		}
		c.logger.LogAttrs(ctx, level, "nats message", attrs...)
	}()

	if msg == nil {
		err = fmt.Errorf("nats message is nil")
		return err
	}
	if msg.Header == nil {
		msg.Header = map[string]string{}
	}
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(msg.Header))

	var header nats.Header
	if len(msg.Header) > 0 {
		header = make(nats.Header, len(msg.Header))
		for k, v := range msg.Header {
			header.Set(k, v)
		}
	}
	natsMsg := &nats.Msg{
		Subject: subject,
		Data:    msg.Data,
		Header:  header,
	}

	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.closed {
		err = fmt.Errorf("client is closed")
		return err
	}

	if c.js != nil && c.conf.StreamName != "" {
		ack, publishErr := c.js.PublishMsg(natsMsg)
		if publishErr != nil {
			err = fmt.Errorf("jetstream publish to %s: %w", subject, publishErr)
			return err
		}
		c.logger.DebugContext(ctx, "jetstream ack", constant.LogFieldKind, constant.LogKindMessage, "stream", ack.Stream, "sequence", ack.Sequence)
		return nil
	}

	if publishErr := c.conn.PublishMsg(natsMsg); publishErr != nil {
		err = fmt.Errorf("nats publish to %s: %w", subject, publishErr)
		return err
	}
	if flushErr := c.conn.FlushTimeout(c.conf.Timeout.AsDuration()); flushErr != nil {
		err = fmt.Errorf("nats flush: %w", flushErr)
		return err
	}
	return nil
}

func (c *NatsClient) Subscribe(ctx context.Context, subject string, handler MessageHandler) (Unsubscriber, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil, fmt.Errorf("client is closed")
	}

	if c.js != nil && c.conf.StreamName != "" {
		sub, err := c.js.Subscribe(subject, func(m *nats.Msg) {
			c.handleMsg(ctx, m, handler, true, "")
		}, nats.DeliverAll(), nats.AckExplicit())
		if err != nil {
			return nil, fmt.Errorf("jetstream subscribe %s: %w", subject, err)
		}
		c.subs = append(c.subs, sub)
		return &subscription{
			sub: sub,
		}, nil
	}

	sub, err := c.conn.Subscribe(subject, func(m *nats.Msg) {
		c.handleMsg(ctx, m, handler, false, "")
	})
	if err != nil {
		return nil, fmt.Errorf("nats subscribe %s: %w", subject, err)
	}

	c.subs = append(c.subs, sub)
	c.logger.InfoContext(ctx, "nats subscribed", constant.LogFieldKind, constant.LogKindMessage, constant.LogFieldSubject, subject)
	return &subscription{
		sub: sub,
	}, nil
}

func (c *NatsClient) QueueSubscribe(ctx context.Context, subject, queue string, handler MessageHandler) (Unsubscriber, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil, fmt.Errorf("client is closed")
	}

	if c.js != nil && c.conf.StreamName != "" {
		sub, err := c.js.QueueSubscribe(subject, queue, func(m *nats.Msg) {
			c.handleMsg(ctx, m, handler, true, queue)
		}, nats.DeliverAll(), nats.AckExplicit())
		if err != nil {
			return nil, fmt.Errorf("jetstream queue subscribe %s[%s]: %w", subject, queue, err)
		}
		c.subs = append(c.subs, sub)
		return &subscription{
			sub: sub,
		}, nil
	}

	sub, err := c.conn.QueueSubscribe(subject, queue, func(m *nats.Msg) {
		c.handleMsg(ctx, m, handler, false, queue)
	})
	if err != nil {
		return nil, fmt.Errorf("nats queue subscribe %s[%s]: %w", subject, queue, err)
	}

	c.subs = append(c.subs, sub)
	c.logger.InfoContext(ctx, "nats queue subscribed", constant.LogFieldKind, constant.LogKindMessage, constant.LogFieldSubject, subject, constant.LogFieldQueue, queue)
	return &subscription{
		sub: sub,
	}, nil
}

func (c *NatsClient) handleMsg(ctx context.Context, m *nats.Msg, handler MessageHandler, isJetStream bool, queue string) {
	if ctx == nil {
		ctx = context.Background()
	}
	start := time.Now()
	subject := "unknown"
	if m != nil && m.Subject != "" {
		subject = m.Subject
	}
	subjectLabel := subject
	if strings.HasPrefix(subjectLabel, "push.node.") {
		subjectLabel = "push.node.*"
	} else if strings.Count(subjectLabel, ".") > 2 {
		subjectLabel = "dynamic"
	}
	msg := &Message{
		Subject: subject,
	}
	if m != nil {
		msg.Data = m.Data
		if len(m.Header) > 0 {
			msg.Header = make(map[string]string, len(m.Header))
			for k, v := range m.Header {
				if len(v) > 0 {
					msg.Header[k] = v[0]
				}
			}
		}
	}
	ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.MapCarrier(msg.Header))
	ctx, span := c.tracer.Start(ctx, "nats consume "+subjectLabel, oteltrace.WithSpanKind(oteltrace.SpanKindConsumer), oteltrace.WithAttributes(
		attribute.String("messaging.system", "nats"),
		attribute.String("messaging.destination.name", subject),
		attribute.String("messaging.operation", "consume"),
		attribute.String("messaging.consumer.group.name", queue),
	))

	var err error
	defer func() {
		status := "ok"
		level := slog.LevelInfo
		if err != nil {
			status = "error"
			level = slog.LevelWarn
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
		latency := time.Since(start)
		MessageRequestsTotal.WithLabelValues(c.service, "consume", subjectLabel, status).Inc()
		MessageRequestDurationSeconds.WithLabelValues(c.service, "consume", subjectLabel, status).Observe(latency.Seconds())
		attrs := []slog.Attr{
			slog.String(constant.LogFieldKind, constant.LogKindMessage),
			slog.String(constant.LogFieldDirection, "consume"),
			slog.String(constant.LogFieldSubject, subject),
			slog.String(constant.LogFieldQueue, queue),
			slog.String(constant.LogFieldStatus, status),
			slog.Int64(constant.LogFieldLatencyMS, latency.Milliseconds()),
		}
		if err != nil {
			attrs = append(attrs, slog.Any(constant.LogFieldErr, err))
		}
		c.logger.LogAttrs(ctx, level, "nats message", attrs...)
	}()

	if m == nil {
		err = fmt.Errorf("nats message is nil")
		return
	}
	if isJetStream {
		msg.Ack = func() error { return m.Ack() }
		msg.Nack = func() error { return m.Nak() }
	}
	if handler == nil {
		err = fmt.Errorf("nats handler is nil")
		return
	}
	if err = handler(ctx, msg); err != nil {
		if msg.Nack != nil {
			if nackErr := msg.Nack(); nackErr != nil {
				err = fmt.Errorf("%w; nack: %v", err, nackErr)
			}
		}
		return
	}
	if msg.Ack != nil {
		if ackErr := msg.Ack(); ackErr != nil {
			err = fmt.Errorf("ack: %w", ackErr)
		}
	}
}

type subscription struct {
	sub *nats.Subscription
}

func (s *subscription) Unsubscribe() error {
	if s.sub.IsValid() {
		return s.sub.Unsubscribe()
	}
	return nil
}
