package client

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/nats-io/nats.go"
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

// Option 是 NATS 客户端的函数选项
type Option func(*options)

type options struct {
	name    string
	url     string
	logger  log.Logger
	timeout time.Duration

	// 连接选项
	maxReconnects  int
	reconnectWait  time.Duration
	pingInterval   time.Duration
	flusherTimeout time.Duration

	// JetStream 选项（可选）
	enableJetStream bool
	streamName      string
}

func defaultOptions() *options {
	return &options{
		name:           "nats-client",
		url:            nats.DefaultURL,
		timeout:        5 * time.Second,
		maxReconnects:  10,
		reconnectWait:  2 * time.Second,
		pingInterval:   30 * time.Second,
		flusherTimeout: 5 * time.Second,
	}
}

func WithName(name string) Option {
	return func(o *options) { o.name = name }
}

func WithURL(url string) Option {
	return func(o *options) { o.url = url }
}

func WithLogger(l log.Logger) Option {
	return func(o *options) { o.logger = l }
}

func WithTimeout(d time.Duration) Option {
	return func(o *options) { o.timeout = d }
}

func WithMaxReconnects(n int) Option {
	return func(o *options) { o.maxReconnects = n }
}

func WithReconnectWait(d time.Duration) Option {
	return func(o *options) { o.reconnectWait = d }
}

func WithPingInterval(d time.Duration) Option {
	return func(o *options) { o.pingInterval = d }
}

func WithFlusherTimeout(d time.Duration) Option {
	return func(o *options) { o.flusherTimeout = d }
}

func WithJetStream(streamName string) Option {
	return func(o *options) {
		o.enableJetStream = true
		o.streamName = streamName
	}
}

// Client 是 NATS 客户端，同时实现 Publisher 和 Subscriber
type Client struct {
	opts   *options
	conn   *nats.Conn
	js     nats.JetStreamContext
	mu     sync.RWMutex
	closed bool
	log    *log.Helper
	subs   []*nats.Subscription
}

// NewClient 创建并连接 NATS 客户端
func NewClient(opts ...Option) (*Client, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	logger := o.logger
	if logger == nil {
		logger = log.DefaultLogger
	}

	c := &Client{
		opts: o,
		log:  log.NewHelper(log.With(logger, "module", "nats/"+o.name)),
	}

	if err := c.connect(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) connect() error {
	c.log.Infof("connecting to nats: %s", c.opts.url)

	nc, err := nats.Connect(c.opts.url,
		nats.Name(c.opts.name),
		nats.MaxReconnects(c.opts.maxReconnects),
		nats.ReconnectWait(c.opts.reconnectWait),
		nats.PingInterval(c.opts.pingInterval),
		nats.FlusherTimeout(c.opts.flusherTimeout),
		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			c.log.Warnf("nats disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			c.log.Infof("nats reconnected to: %s", conn.ConnectedUrl())
		}),
		nats.ClosedHandler(func(conn *nats.Conn) {
			c.log.Info("nats connection closed")
		}),
		nats.ErrorHandler(func(conn *nats.Conn, sub *nats.Subscription, err error) {
			c.log.Errorf("nats error: subject=%s err=%v", sub.Subject, err)
		}),
	)
	if err != nil {
		return fmt.Errorf("nats connect: %w", err)
	}

	c.conn = nc

	if c.opts.enableJetStream {
		js, err := nc.JetStream()
		if err != nil {
			nc.Close()
			return fmt.Errorf("nats jetstream: %w", err)
		}
		c.js = js
		c.log.Infof("jetstream enabled, stream=%s", c.opts.streamName)
	}

	c.log.Infof("nats connected: %s", nc.ConnectedUrl())
	return nil
}

// Publish 发布消息到指定 subject
func (c *Client) Publish(ctx context.Context, subject string, msg *Message) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if c.closed {
		return fmt.Errorf("client is closed")
	}

	if c.js != nil && c.opts.streamName != "" {
		return c.publishJetStream(subject, msg)
	}

	natsMsg := &nats.Msg{
		Subject: subject,
		Data:    msg.Data,
		Header:  toNatsHeader(msg.Header),
	}

	if err := c.conn.PublishMsg(natsMsg); err != nil {
		return fmt.Errorf("nats publish to %s: %w", subject, err)
	}

	return c.conn.FlushTimeout(c.opts.timeout)
}

func (c *Client) publishJetStream(subject string, msg *Message) error {
	natsMsg := &nats.Msg{
		Subject: subject,
		Data:    msg.Data,
		Header:  toNatsHeader(msg.Header),
	}

	ack, err := c.js.PublishMsg(natsMsg)
	if err != nil {
		return fmt.Errorf("jetstream publish to %s: %w", subject, err)
	}

	c.log.Debugf("jetstream ack: stream=%s seq=%d", ack.Stream, ack.Sequence)
	return nil
}

// Subscribe 订阅 subject，返回 Unsubscriber 用于取消
func (c *Client) Subscribe(ctx context.Context, subject string, handler MessageHandler) (Unsubscriber, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil, fmt.Errorf("client is closed")
	}

	if c.js != nil && c.opts.streamName != "" {
		return c.subscribeJetStream(subject, handler)
	}

	sub, err := c.conn.Subscribe(subject, func(m *nats.Msg) {
		c.handleMsg(ctx, m, handler, false) // ← core NATS，不需要 ack
	})
	if err != nil {
		return nil, fmt.Errorf("nats subscribe %s: %w", subject, err)
	}

	c.subs = append(c.subs, sub)
	c.log.Infof("subscribed to: %s", subject)

	return &subscription{sub: sub}, nil
}

// QueueSubscribe 使用队列组订阅（负载均衡）
func (c *Client) QueueSubscribe(ctx context.Context, subject, queue string, handler MessageHandler) (Unsubscriber, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil, fmt.Errorf("client is closed")
	}

	if c.js != nil && c.opts.streamName != "" {
		return c.queueSubscribeJetStream(subject, queue, handler)
	}

	sub, err := c.conn.QueueSubscribe(subject, queue, func(m *nats.Msg) {
		c.handleMsg(ctx, m, handler, false) // ← core NATS
	})
	if err != nil {
		return nil, fmt.Errorf("nats queue subscribe %s[%s]: %w", subject, queue, err)
	}

	c.subs = append(c.subs, sub)
	c.log.Infof("queue subscribed to: %s [%s]", subject, queue)

	return &subscription{sub: sub}, nil
}

func (c *Client) subscribeJetStream(subject string, handler MessageHandler) (Unsubscriber, error) {
	sub, err := c.js.Subscribe(subject, func(m *nats.Msg) {
		c.handleMsg(context.Background(), m, handler, true)
	}, nats.DeliverAll(), nats.AckExplicit())
	if err != nil {
		return nil, fmt.Errorf("jetstream subscribe %s: %w", subject, err)
	}

	c.subs = append(c.subs, sub)
	return &subscription{sub: sub}, nil
}

func (c *Client) queueSubscribeJetStream(subject, queue string, handler MessageHandler) (Unsubscriber, error) {
	sub, err := c.js.QueueSubscribe(subject, queue, func(m *nats.Msg) {
		c.handleMsg(context.Background(), m, handler, true)
	}, nats.DeliverAll(), nats.AckExplicit())
	if err != nil {
		return nil, fmt.Errorf("jetstream queue subscribe %s[%s]: %w", subject, queue, err)
	}

	c.subs = append(c.subs, sub)
	return &subscription{sub: sub}, nil
}

func (c *Client) handleMsg(ctx context.Context, m *nats.Msg, handler MessageHandler, isJetStream bool) {
	msg := &Message{
		Subject: m.Subject,
		Data:    m.Data,
		Header:  fromNatsHeader(m.Header),
	}

	// JetStream 消息需要 Ack/Nack，包装为统一签名
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

// Conn 返回底层 nats.Conn，用于高级场景
func (c *Client) Conn() *nats.Conn {
	return c.conn
}

// JetStream 返回 JetStream 上下文
func (c *Client) JetStream() nats.JetStreamContext {
	return c.js
}

// Close 关闭客户端
func (c *Client) Close() error {
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
		c.conn.Drain()
		// 等待 Drain 完成
		time.Sleep(500 * time.Millisecond)
		c.conn.Close()
	}

	c.log.Info("nats client closed")
	return nil
}

// subscription 实现 Unsubscriber
type subscription struct {
	sub *nats.Subscription
}

func (s *subscription) Unsubscribe() error {
	if s.sub.IsValid() {
		return s.sub.Unsubscribe()
	}
	return nil
}

// 辅助函数：nats.Header <-> map[string]string
func toNatsHeader(h map[string]string) nats.Header {
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
