package client

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

func TestNewClient(t *testing.T) {
	url := "nats://192.168.100.10:30083"

	client, err := NewClient(
		WithURL(url),
		WithName("test-client"),
		WithTimeout(3*time.Second),
	)
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer client.Close()

	if client.Conn() == nil {
		t.Fatal("conn should not be nil")
	}

	if client.Conn().Status() != nats.CONNECTED {
		t.Fatalf("expected CONNECTED, got %s", client.Conn().Status())
	}
}

func TestPublishSubscribe(t *testing.T) {
	url := "nats://192.168.100.10:30083"

	client, err := NewClient(WithURL(url))
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	received := make(chan *Message, 1)

	// 订阅
	unsub, err := client.Subscribe(ctx, "test.topic", func(ctx context.Context, msg *Message) error {
		received <- msg
		return nil
	})
	if err != nil {
		t.Fatalf("subscribe: %v", err)
	}
	defer unsub.Unsubscribe()

	// 等待订阅生效
	time.Sleep(100 * time.Millisecond)

	// 发布
	err = client.Publish(ctx, "test.topic", &Message{
		Data: []byte("hello nats"),
		Header: map[string]string{
			"x-request-id": "req-123",
		},
	})
	if err != nil {
		t.Fatalf("publish: %v", err)
	}

	// 等待消息
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
	url := "nats://192.168.100.10:30083"

	// 创建多个客户端模拟多个消费者
	clients := make([]*Client, 3)
	for i := range clients {
		c, err := NewClient(WithURL(url), WithName(fmt.Sprintf("worker-%d", i)))
		if err != nil {
			t.Fatalf("new client %d: %v", i, err)
		}
		defer c.Close()
		clients[i] = c
	}

	ctx := context.Background()
	var count atomic.Int64
	var wg sync.WaitGroup
	totalMessages := 30

	wg.Add(totalMessages)

	// 所有客户端订阅同一队列组
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

	// 发布消息
	for i := 0; i < totalMessages; i++ {
		err := clients[0].Publish(ctx, "orders.new", &Message{
			Data: []byte(fmt.Sprintf("order-%d", i)),
		})
		if err != nil {
			t.Fatalf("publish %d: %v", i, err)
		}
	}

	// 等待所有消息被消费
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
	url := "nats://192.168.100.10:30083"

	client, err := NewClient(WithURL(url))
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	client.Close()

	err = client.Publish(context.Background(), "test", &Message{Data: []byte("fail")})
	if err == nil {
		t.Fatal("expected error publishing after close")
	}
}

func TestSubscribeAfterClose(t *testing.T) {
	url := "nats://192.168.100.10:30083"

	client, err := NewClient(WithURL(url))
	if err != nil {
		t.Fatalf("new client: %v", err)
	}

	client.Close()

	_, err = client.Subscribe(context.Background(), "test", func(ctx context.Context, msg *Message) error {
		return nil
	})
	if err == nil {
		t.Fatal("expected error subscribing after close")
	}
}

func TestMultipleSubjects(t *testing.T) {
	url := "nats://192.168.100.10:30083"

	client, err := NewClient(WithURL(url))
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	subjects := []string{"user.created", "user.updated", "user.deleted"}
	received := make(map[string]*Message)
	var mu sync.Mutex

	for _, sub := range subjects {
		s := sub // capture
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

	// 发布到每个 subject
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
	url := "nats://192.168.100.10:30083"

	client, err := NewClient(WithURL(url))
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer client.Close()

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

	// 发布消息 — handler 返回错误不应 panic
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
	url := "nats://192.168.100.10:30083"

	client, err := NewClient(WithURL(url))
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer client.Close()

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

	// 第一条消息应该收到
	_ = client.Publish(ctx, "test.unsub", &Message{Data: []byte("msg1")})
	time.Sleep(200 * time.Millisecond)

	// 取消订阅
	if err := unsub.Unsubscribe(); err != nil {
		t.Fatalf("unsubscribe: %v", err)
	}

	time.Sleep(100 * time.Millisecond)

	// 第二条消息不应该收到
	_ = client.Publish(ctx, "test.unsub", &Message{Data: []byte("msg2")})
	time.Sleep(300 * time.Millisecond)

	if got := count.Load(); got != 1 {
		t.Fatalf("expected 1 message, got %d", got)
	}
}

func TestConcurrentPublish(t *testing.T) {
	url := "nats://192.168.100.10:30083"

	client, err := NewClient(WithURL(url))
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	defer client.Close()

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

	// 并发发布
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

// BenchmarkPublish 基准测试
func BenchmarkPublish(b *testing.B) {
	url := "nats://192.168.100.10:30083"

	client, err := NewClient(WithURL(url))
	if err != nil {
		b.Fatalf("new client: %v", err)
	}
	defer client.Close()

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
