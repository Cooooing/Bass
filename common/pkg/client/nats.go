package client

import (
	"common/api/gen/common"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Message 封装消息体，解耦外部依赖
type Message struct {
	Subject string
	Data    []byte
	Header  map[string]string
	Ack     func() error // 手动确认（JetStream 模式）
	Nack    func() error
}

// MessageHandler 是消息处理函数签名
type MessageHandler func(ctx context.Context, msg *Message) error

// Publisher 定义发布接口
type Publisher interface {
	Publish(ctx context.Context, subject string, msg *Message) error
	Close() error
}

// Subscriber 定义订阅接口
type Subscriber interface {
	Subscribe(ctx context.Context, subject string, handler MessageHandler) (Unsubscriber, error)
	QueueSubscribe(ctx context.Context, subject, queue string, handler MessageHandler) (Unsubscriber, error)
	Close() error
}

// Unsubscriber 取消订阅
type Unsubscriber interface {
	Unsubscribe() error
}

// NatsClient 封装 NATS 客户端
type NatsClient struct {
	conf   *common.Nats
	conn   *nats.Conn
	js     nats.JetStreamContext
	mu     sync.RWMutex
	closed bool
	log    *log.Helper
	subs   []*nats.Subscription
}

// NewNatsClient 初始化 NATS 客户端
func NewNatsClient(logger log.Logger, conf *common.Nats) (*NatsClient, func(), error) {
	helper := log.NewHelper(logger)

	// 默认值
	if conf == nil {
		conf = &common.Nats{}
	}
	if conf.Url == "" {
		conf.Url = nats.DefaultURL
	}
	if conf.Name == "" {
		conf.Name = "nats-client"
	}
	if conf.Timeout.AsDuration() == 0 {
		conf.Timeout = durationpb.New(5 * time.Second)
	}
	if conf.MaxReconnects == 0 {
		conf.MaxReconnects = 10
	}
	if conf.ReconnectWait.AsDuration() == 0 {
		conf.ReconnectWait = durationpb.New(2 * time.Second)
	}
	if conf.PingInterval.AsDuration() == 0 {
		conf.PingInterval = durationpb.New(30 * time.Second)
	}
	if conf.FlusherTimeout.AsDuration() == 0 {
		conf.FlusherTimeout = durationpb.New(5 * time.Second)
	}

	opts := []nats.Option{
		nats.Name(conf.Name),
		nats.MaxReconnects(int(conf.MaxReconnects)),
		nats.ReconnectWait(conf.ReconnectWait.AsDuration()),
		nats.PingInterval(conf.PingInterval.AsDuration()),
		nats.FlusherTimeout(conf.FlusherTimeout.AsDuration()),
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			helper.Warnf("nats disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(c *nats.Conn) {
			helper.Infof("nats reconnected to: %s", c.ConnectedUrl())
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			helper.Info("nats connection closed")
		}),
		nats.ErrorHandler(func(_ *nats.Conn, sub *nats.Subscription, err error) {
			helper.Errorf("nats error: subject=%s err=%v", sub.Subject, err)
		}),
	}

	nc, err := nats.Connect(conf.Url, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("nats connect [%s]: %w", conf.Url, err)
	}

	c := &NatsClient{
		conf: conf,
		conn: nc,
		log:  helper,
	}

	if conf.EnableJetStream {
		js, err := nc.JetStream()
		if err != nil {
			nc.Close()
			return nil, nil, fmt.Errorf("nats jetstream: %w", err)
		}
		c.js = js
		helper.Infof("jetstream enabled: domain=%s stream=%s", conf.JetStreamDomain, conf.StreamName)
	}

	helper.Infof("nats connected: %s", nc.ConnectedUrl())
	return c, func() {
		if err := c.Close(); err != nil {
			helper.Errorf("nats close: %s", err)
		}
	}, nil
}

// Conn 返回底层 nats.Conn，用于高级场景
func (c *NatsClient) Conn() *nats.Conn {
	return c.conn
}

// JetStream 返回 JetStream 上下文
func (c *NatsClient) JetStream() nats.JetStreamContext {
	return c.js
}

// Close 关闭客户端
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
				c.log.Errorf("drain subscription: %v", err)
			}
		}
	}

	if c.conn != nil {
		_ = c.conn.Drain()
		time.Sleep(500 * time.Millisecond)
		c.conn.Close()
	}

	c.log.Info("nats client closed")
	return nil
}

// Publish 发布消息到指定主题。
func (c *NatsClient) Publish(ctx context.Context, subject string, msg *Message) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.closed {
		return fmt.Errorf("client is closed")
	}

	if c.js != nil && c.conf.StreamName != "" {
		return c.publishJetStream(subject, msg)
	}

	natsMsg := &nats.Msg{
		Subject: subject,
		Data:    msg.Data,
		Header:  buildNatsHeader(msg.Header),
	}

	if err := c.conn.PublishMsg(natsMsg); err != nil {
		return fmt.Errorf("nats publish to %s: %w", subject, err)
	}

	return c.conn.FlushTimeout(c.conf.Timeout.AsDuration())
}

func (c *NatsClient) publishJetStream(subject string, msg *Message) error {
	natsMsg := &nats.Msg{
		Subject: subject,
		Data:    msg.Data,
		Header:  buildNatsHeader(msg.Header),
	}

	ack, err := c.js.PublishMsg(natsMsg)
	if err != nil {
		return fmt.Errorf("jetstream publish to %s: %w", subject, err)
	}

	c.log.Debugf("jetstream ack: stream=%s seq=%d", ack.Stream, ack.Sequence)
	return nil
}

// Subscribe 订阅主题，返回 Unsubscriber 用于取消。
func (c *NatsClient) Subscribe(ctx context.Context, subject string, handler MessageHandler) (Unsubscriber, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil, fmt.Errorf("client is closed")
	}

	if c.js != nil && c.conf.StreamName != "" {
		return c.subscribeJetStream(subject, handler)
	}

	sub, err := c.conn.Subscribe(subject, func(m *nats.Msg) {
		c.handleMsg(ctx, m, handler, false)
	})
	if err != nil {
		return nil, fmt.Errorf("nats subscribe %s: %w", subject, err)
	}

	c.subs = append(c.subs, sub)
	c.log.Infof("subscribed to: %s", subject)

	return &subscription{sub: sub}, nil
}

// QueueSubscribe 使用队列组订阅（负载均衡）
func (c *NatsClient) QueueSubscribe(ctx context.Context, subject, queue string, handler MessageHandler) (Unsubscriber, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil, fmt.Errorf("client is closed")
	}

	if c.js != nil && c.conf.StreamName != "" {
		return c.queueSubscribeJetStream(subject, queue, handler)
	}

	sub, err := c.conn.QueueSubscribe(subject, queue, func(m *nats.Msg) {
		c.handleMsg(ctx, m, handler, false)
	})
	if err != nil {
		return nil, fmt.Errorf("nats queue subscribe %s[%s]: %w", subject, queue, err)
	}

	c.subs = append(c.subs, sub)
	c.log.Infof("queue subscribed to: %s [%s]", subject, queue)

	return &subscription{sub: sub}, nil
}

func (c *NatsClient) subscribeJetStream(subject string, handler MessageHandler) (Unsubscriber, error) {
	sub, err := c.js.Subscribe(subject, func(m *nats.Msg) {
		c.handleMsg(context.Background(), m, handler, true)
	}, nats.DeliverAll(), nats.AckExplicit())
	if err != nil {
		return nil, fmt.Errorf("jetstream subscribe %s: %w", subject, err)
	}

	c.subs = append(c.subs, sub)
	return &subscription{sub: sub}, nil
}

func (c *NatsClient) queueSubscribeJetStream(subject, queue string, handler MessageHandler) (Unsubscriber, error) {
	sub, err := c.js.QueueSubscribe(subject, queue, func(m *nats.Msg) {
		c.handleMsg(context.Background(), m, handler, true)
	}, nats.DeliverAll(), nats.AckExplicit())
	if err != nil {
		return nil, fmt.Errorf("jetstream queue subscribe %s[%s]: %w", subject, queue, err)
	}

	c.subs = append(c.subs, sub)
	return &subscription{sub: sub}, nil
}

func (c *NatsClient) handleMsg(ctx context.Context, m *nats.Msg, handler MessageHandler, isJetStream bool) {
	msg := &Message{
		Subject: m.Subject,
		Data:    m.Data,
		Header:  fromNatsHeader(m.Header),
	}

	if isJetStream {
		msg.Ack = func() error { return m.Ack() }
		msg.Nack = func() error { return m.Nak() }
	}

	if err := handler(ctx, msg); err != nil {
		c.log.Errorf("handler error: subject=%s err=%v", m.Subject, err)
		if msg.Nack != nil {
			if nackErr := msg.Nack(); nackErr != nil {
				c.log.Errorf("nack error: %v", nackErr)
			}
		}
		return
	}

	if msg.Ack != nil {
		if ackErr := msg.Ack(); ackErr != nil {
			c.log.Errorf("ack error: %v", ackErr)
		}
	}
}

// subscription 实现 Unsubscriber。
type subscription struct {
	sub *nats.Subscription
}

func (s *subscription) Unsubscribe() error {
	if s.sub.IsValid() {
		return s.sub.Unsubscribe()
	}
	return nil
}

func buildNatsHeader(h map[string]string) nats.Header {
	if len(h) == 0 {
		return nil
	}
	nh := make(nats.Header, len(h))
	for k, v := range h {
		nh.Set(k, v)
	}
	return nh
}

func fromNatsHeader(h nats.Header) map[string]string {
	if len(h) == 0 {
		return nil
	}
	m := make(map[string]string, len(h))
	for k, v := range h {
		if len(v) > 0 {
			m[k] = v[0]
		}
	}
	return m
}
