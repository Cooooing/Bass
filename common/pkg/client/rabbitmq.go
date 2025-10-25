package client

import (
	"common/pkg/constant"
	"common/pkg/model"
	"fmt"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	log  *log.Helper
	conf *model.RabbitmqConf
	conn *amqp.Connection
	pool *sync.Pool // *amqp091.Channel 池
}

// NewRabbitMQClient 初始化 RabbitMQ 单机客户端
func NewRabbitMQClient(log *log.Helper, conf *model.RabbitmqConf) (*RabbitMQClient, func(), error) {
	conn, err := amqp.DialConfig(conf.Url, amqp.Config{
		Heartbeat: conf.Heartbeat.AsDuration(),
		Dial:      amqp.DefaultDial(conf.DialTimeout.AsDuration()),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	client := &RabbitMQClient{
		log:  log,
		conn: conn,
		conf: conf,
		pool: &sync.Pool{
			New: func() any {
				ch, err := conn.Channel()
				if err != nil {
					log.Errorf("failed to create channel: %s", err)
					return nil
				}
				if err := ch.Qos(int(conf.PrefetchCount), 0, conf.PrefetchGlobal); err != nil {
					log.Errorf("failed to set qos: %s", err)
				}
				return ch
			},
		},
	}

	// 初始化队列和交换机
	err = client.declareResources()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to declare resources: %w", err)
	}

	log.Infof("rabbitmq: connected to [%s]", conn.RemoteAddr().String())

	// 清理函数
	cleanup := func() {
		if err := client.conn.Close(); err != nil {
			log.Errorf("failed to close RabbitMQ connection: %s", err.Error())
		} else {
			log.Infof("rabbitmq connection closed")
		}
	}

	return client, cleanup, nil
}

// 从池中获取一个 channel
func (r *RabbitMQClient) getChannel() (*amqp.Channel, error) {
	ch := r.pool.Get()
	if ch == nil {
		return nil, fmt.Errorf("failed to get channel from pool")
	}
	return ch.(*amqp.Channel), nil
}

// 放回 channel
func (r *RabbitMQClient) releaseChannel(ch *amqp.Channel) {
	if ch != nil && !ch.IsClosed() {
		r.pool.Put(ch)
	}
}

func (r *RabbitMQClient) declareResources() error {
	ch, err := r.getChannel()
	if err != nil {
		return err
	}
	defer r.releaseChannel(ch)

	// 声明 Exchange
	for _, v := range constant.ExchangeMap {
		err = ch.ExchangeDeclare(v.Name.String(), v.Kind, v.Durable, v.AutoDelete, v.Internal, v.NoWait, v.Args)
		if err != nil {
			return err
		}
	}

	// 声明 Queue
	for _, v := range constant.QueueMap {
		_, err = ch.QueueDeclare(v.Name.String(), v.Durable, v.AutoDelete, v.Exclusive, v.NoWait, v.Args)
		if err != nil {
			return err
		}
	}

	// 绑定 Queue 到 Exchange
	for _, v := range constant.QueueBindMap {
		err = ch.QueueBind(v.Name.String(), v.Key.String(), v.Exchange.String(), v.NoWait, v.Args)
		if err != nil {
			return err
		}
	}

	return nil
}

// Publish 发送消息
func (r *RabbitMQClient) Publish(exchange, routingKey string, body []byte) error {
	ch, err := r.getChannel()
	if err != nil {
		return err
	}
	defer r.releaseChannel(ch)

	if exchange == "" {
		return fmt.Errorf("exchange cannot be empty")
	}
	if routingKey == "" {
		return fmt.Errorf("routingKey cannot be empty")
	}

	return ch.Publish(
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         body,
			DeliveryMode: uint8(r.conf.DeliveryMode), // 消息持久化
		},
	)
}

// Consume 消费消息
func (r *RabbitMQClient) Consume(queue string) (<-chan amqp.Delivery, error) {
	ch, err := r.getChannel()
	if err != nil {
		return nil, err
	}
	defer r.releaseChannel(ch)

	if queue == "" {
		return nil, fmt.Errorf("queue cannot be empty")
	}
	var consumerTag string

	msgs, err := ch.Consume(
		queue,
		consumerTag,
		r.conf.AutoAck,
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // args
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume queue [%s]: %w", queue, err)
	}

	return msgs, nil
}
