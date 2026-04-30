package client

import (
	"common/api/gen/common"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

const testNatsURL = "nats://192.168.100.10:30083"

var l = log.DefaultLogger

func testNatsConf() *common.Nats {
	return &common.Nats{
		Url:  testNatsURL,
		Name: "test-client",
	}
}

func TestNewClient(t *testing.T) {
	conf := testNatsConf()
	client, cleanup, err := NewNatsClient(l, conf)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer cleanup()

	if client.Conn() == nil {
		t.Fatal("conn should not be nil")
	}
}

func TestPublishSubscribe(t *testing.T) {
	client, cleanup, err := NewNatsClient(l, testNatsConf())
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer cleanup()

	ctx := context.Background()
	received := make(chan *Message, 1)

	unsub, err := client.Subscribe(ctx, "test.topic", func(ctx context.Context, msg *Message) error {
		received <- msg
		return nil
	})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer unsub.Unsubscribe()

	time.Sleep(100 * time.Millisecond)

	err = client.Publish(ctx, "test.topic", &Message{
		Data: []byte("hello nats"),
		Header: map[string]string{
			"x-request-id": "req-123",
		},
	})
	if err != nil {
		t.Fatalf("publish: %v", err)
	}

	select {
	case msg := <-received:
		if string(msg.Data) != "hello nats" {
			t.Fatalf("expected 'hello nats', got '%s'", string(msg.Data))
		}
		if msg.Header["x-request-id"] != "req-123" {
			t.Fatalf("expected header x-request-id=req-123, got %s", msg.Header["x-request-id"])
		}
	case <-time.After(3 * time.Second):
		t.Fatal("timeout waiting for message")
	}
}

func TestQueueSubscribe(t *testing.T) {
	clients := make([]*NatsClient, 3)
	for i := range clients {
		conf := testNatsConf()
		conf.Name = fmt.Sprintf("worker-%d", i)
		c, cleanup, err := NewNatsClient(l, conf)
		if err != nil {
			t.Fatalf("new client %d: %v", i, err)
		}
		defer cleanup()
		clients[i] = c
	}

	ctx := context.Background()
	var count atomic.Int64
	var wg sync.WaitGroup
	totalMessages := 30

	wg.Add(totalMessages)

	for i, c := range clients {
		_, err := c.QueueSubscribe(ctx, "orders.new", "order-processors",
			func(ctx context.Context, msg *Message) error {
				count.Add(1)
				wg.Done()
				return nil
			},
		)
		if err != nil {
			t.Fatalf("queue subscribe client %d: %v", i, err)
		}
	}

	time.Sleep(200 * time.Millisecond)

	for i := 0; i < totalMessages; i++ {
		err := clients[0].Publish(ctx, "orders.new", &Message{
			Data: []byte(fmt.Sprintf("order-%d", i)),
		})
		if err != nil {
			t.Fatalf("publish %d: %v", i, err)
		}
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		if got := count.Load(); got != int64(totalMessages) {
			t.Fatalf("expected %d messages, got %d", totalMessages, got)
		}
	case <-time.After(5 * time.Second):
		t.Fatalf("timeout: processed %d/%d messages", count.Load(), totalMessages)
	}
}

func TestPublishAfterClose(t *testing.T) {
	conf := testNatsConf()
	client, _, err := NewNatsClient(l, conf)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	_ = client.Close()

	err = client.Publish(context.Background(), "test", &Message{Data: []byte("fail")})
	if err == nil {
		t.Fatal("expected error publishing after close")
	}
}

func TestSubscribeAfterClose(t *testing.T) {
	conf := testNatsConf()
	client, _, err := NewNatsClient(l, conf)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	_ = client.Close()

	_, err = client.Subscribe(context.Background(), "test", func(ctx context.Context, msg *Message) error {
		return nil
	})
	if err == nil {
		t.Fatal("expected error subscribing after close")
	}
}

func TestMultipleSubjects(t *testing.T) {
	conf := testNatsConf()
	client, cleanup, err := NewNatsClient(l, conf)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer cleanup()

	ctx := context.Background()
	subjects := []string{"user.created", "user.updated", "user.deleted"}
	received := make(map[string]*Message)
	var mu sync.Mutex

	for _, sub := range subjects {
		s := sub
		_, err := client.Subscribe(ctx, s, func(ctx context.Context, msg *Message) error {
			mu.Lock()
			received[s] = msg
			mu.Unlock()
			return nil
		})
		if err != nil {
			t.Fatalf("subscribe %s: %v", s, err)
		}
	}

	time.Sleep(100 * time.Millisecond)

	for _, sub := range subjects {
		err := client.Publish(ctx, sub, &Message{Data: []byte(sub)})
		if err != nil {
			t.Fatalf("publish %s: %v", sub, err)
		}
	}

	time.Sleep(500 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()
	for _, sub := range subjects {
		msg, ok := received[sub]
		if !ok {
			t.Fatalf("missing message for subject %s", sub)
		}
		if string(msg.Data) != sub {
			t.Fatalf("subject %s: expected '%s', got '%s'", sub, sub, string(msg.Data))
		}
	}
}

func TestHandlerError(t *testing.T) {
	conf := testNatsConf()
	client, cleanup, err := NewNatsClient(l, conf)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer cleanup()

	ctx := context.Background()
	var failCount atomic.Int32

	_, err = client.Subscribe(ctx, "test.error", func(ctx context.Context, msg *Message) error {
		failCount.Add(1)
		return fmt.Errorf("processing failed")
	})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	err = client.Publish(ctx, "test.error", &Message{Data: []byte("will fail")})
	if err != nil {
		t.Fatalf("publish: %v", err)
	}

	time.Sleep(300 * time.Millisecond)

	if failCount.Load() != 1 {
		t.Fatalf("expected 1 failure, got %d", failCount.Load())
	}
}

func TestUnsubscribe(t *testing.T) {
	conf := testNatsConf()
	client, cleanup, err := NewNatsClient(l, conf)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer cleanup()

	ctx := context.Background()
	var count atomic.Int32

	unsub, err := client.Subscribe(ctx, "test.unsub", func(ctx context.Context, msg *Message) error {
		count.Add(1)
		return nil
	})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	_ = client.Publish(ctx, "test.unsub", &Message{Data: []byte("msg1")})
	time.Sleep(200 * time.Millisecond)

	if err := unsub.Unsubscribe(); err != nil {
		t.Fatalf("unsubscribe: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	_ = client.Publish(ctx, "test.unsub", &Message{Data: []byte("msg2")})
	time.Sleep(300 * time.Millisecond)

	if got := count.Load(); got != 1 {
		t.Fatalf("expected 1 message, got %d", got)
	}
}

func TestConcurrentPublish(t *testing.T) {
	conf := testNatsConf()
	client, cleanup, err := NewNatsClient(l, conf)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer cleanup()

	ctx := context.Background()
	var received atomic.Int64
	totalMessages := 100

	_, err = client.Subscribe(ctx, "test.concurrent", func(ctx context.Context, msg *Message) error {
		received.Add(1)
		return nil
	})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	var wg sync.WaitGroup
	for i := 0; i < totalMessages; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			err := client.Publish(ctx, "test.concurrent", &Message{
				Data: []byte(fmt.Sprintf("msg-%d", n)),
			})
			if err != nil {
				t.Errorf("publish %d: %v", n, err)
			}
		}(i)
	}
	wg.Wait()

	time.Sleep(1 * time.Second)

	if got := received.Load(); got != int64(totalMessages) {
		t.Fatalf("expected %d, got %d", totalMessages, got)
	}
}

func BenchmarkPublish(b *testing.B) {
	conf := testNatsConf()
	client, cleanup, err := NewNatsClient(l, conf)
	if err != nil {
		b.Fatalf("new client: %v", err)
	}
	defer cleanup()

	ctx := context.Background()
	msg := &Message{Data: []byte("benchmark payload data")}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if err := client.Publish(ctx, "bench.test", msg); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func newTestNatsClient(conf *common.Nats) (func(), error) {
	_, cleanup, err := NewNatsClient(l, conf)
	return cleanup, err
}

// Ensure NatsClient satisfies the expected interfaces.
var (
	_ Publisher  = (*NatsClient)(nil)
	_ Subscriber = (*NatsClient)(nil)
)
