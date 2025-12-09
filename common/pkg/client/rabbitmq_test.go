package client

import (
	"common/pkg/constant"
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
		Url:            "amqp://admin:123456@127.0.0.1:5672/dev_vhost",
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

	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("hello %d", i)
		err = client.Publish(constant.ExchangeUser.String(), "user.create", []byte(s))
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

	msgs, ch, err := client.Consume(constant.QueueNotify.String())
	if err != nil {
		t.Error(err)
	}
	defer client.ReleaseChannel(ch)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			timeout := time.After(3 * time.Second)
			select {
			case msg := <-msgs:
				t.Log("Received message:", string(msg.Body))
				err = msg.Ack(false)

			case <-timeout:
				return
			}
		}
	}()
	wg.Wait()
}
