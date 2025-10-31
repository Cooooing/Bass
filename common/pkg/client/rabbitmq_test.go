package client

import (
	"common/pkg/model"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/durationpb"
)

func newClient() (*RabbitMQClient, func(), error) {
	client, f, err := NewRabbitMQClient(log.NewHelper(log.NewStdLogger(os.Stdout)), &model.RabbitmqConf{
		Url:            "amqp://root:123456@192.168.1.6:5672/admin_vhost",
		Heartbeat:      durationpb.New(time.Second * 10),
		DialTimeout:    durationpb.New(time.Second * 5),
		PrefetchCount:  10,
		PrefetchGlobal: false,
		DeliveryMode:   2,
		AutoAck:        false,
	})
	if err != nil {
		return nil, nil, err
	}
	return client, f, err
}

func TestRabbitMQClient_Publishsh(t *testing.T) {
	client, f, err := newClient()
	if err != nil {
		t.Error(err)
	}
	defer f()

	for i := 0; i < 100; i++ {
		s := fmt.Sprintf("hello %d", i)
		err = client.Publish("user_events", "user.created", []byte(s))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestRabbitMQClient_Consume(t *testing.T) {
	client, f, err := newClient()
	if err != nil {
		t.Error(err)
	}
	defer f()

	ch, err := client.Consume("email_service_queue")
	if err != nil {
		t.Error(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			timeout := time.After(3 * time.Second)
			select {
			case msg := <-ch:
				t.Log("Received message:", string(msg.Body))
				err = msg.Ack(false)

			case <-timeout:
				return
			}
		}
	}()
	wg.Wait()
}
