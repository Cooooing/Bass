package nats

import (
	commonclient "common/pkg/client"
	"fmt"
	"net"
	"strconv"
	"sync"

	"common/proto/gen/common/v1"

	gonats "github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// Client 是 NATS Go SDK 的薄适配器。
type Client struct {
	observer *commonclient.Observer
	conf     *common.Nats
	conn     *gonats.Conn
	js       jetstream.JetStream
	mu       sync.Mutex
	subs     []*Subscription
}

func NewClient(conf *common.Nats, observer *commonclient.Observer) (*Client, func(), error) {
	if conf == nil {
		conf = &common.Nats{}
	}
	host, port := conf.GetHost(), conf.GetPort()
	if host == "" {
		host = "127.0.0.1"
	}
	if port == 0 {
		port = 4222
	}
	opts := []gonats.Option{gonats.Name(conf.GetName())}
	if conf.GetUser() != "" || conf.GetPassword() != "" {
		opts = append(opts, gonats.UserInfo(conf.GetUser(), conf.GetPassword()))
	}
	if value := conf.GetTimeout(); value != nil && value.AsDuration() > 0 {
		opts = append(opts, gonats.Timeout(value.AsDuration()))
	}
	if conf.GetMaxReconnects() != 0 {
		opts = append(opts, gonats.MaxReconnects(int(conf.GetMaxReconnects())))
	}
	if value := conf.GetReconnectWait(); value != nil && value.AsDuration() > 0 {
		opts = append(opts, gonats.ReconnectWait(value.AsDuration()))
	}
	if value := conf.GetPingInterval(); value != nil && value.AsDuration() > 0 {
		opts = append(opts, gonats.PingInterval(value.AsDuration()))
	}
	if value := conf.GetFlusherTimeout(); value != nil && value.AsDuration() > 0 {
		opts = append(opts, gonats.FlusherTimeout(value.AsDuration()))
	}
	conn, err := gonats.Connect("nats://"+net.JoinHostPort(host, strconv.Itoa(int(port))), opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("connect nats: %w", err)
	}
	js, err := jetstream.New(conn)
	if err != nil {
		conn.Close()
		return nil, nil, fmt.Errorf("create jetstream client: %w", err)
	}
	c := &Client{observer: observer, conf: conf, conn: conn, js: js}
	return c, func() { _ = c.Close() }, nil
}

func (c *Client) Conn() *gonats.Conn {
	return c.conn
}

func (c *Client) JetStream() jetstream.JetStream {
	return c.js
}

func (c *Client) Close() error {
	c.mu.Lock()
	subs := c.subs
	c.subs = nil
	c.mu.Unlock()
	for _, sub := range subs {
		_ = sub.Unsubscribe()
	}
	if c.conn == nil {
		return nil
	}
	return c.conn.Drain()
}

func (c *Client) track(sub *Subscription) {
	c.mu.Lock()
	defer c.mu.Unlock()
	sub.client = c
	c.subs = append(c.subs, sub)
}

func (c *Client) untrack(sub *Subscription) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for index, item := range c.subs {
		if item != sub {
			continue
		}
		c.subs = append(c.subs[:index], c.subs[index+1:]...)
		return
	}
}
