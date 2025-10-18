package client

import (
	"common/pkg/model"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	log      *log.Helper
	conf     *model.RabbitmqConf
	conn     *amqp.Connection
	channels map[string]*amqp.Channel
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
		log:      log,
		conn:     conn,
		channels: make(map[string]*amqp.Channel),
		conf:     conf,
	}

	log.Infof("rabbitmq: connected to [%s]", conn.RemoteAddr().String())

	// 清理函数
	cleanup := func() {
		for name, ch := range client.channels {
			if err := ch.Close(); err != nil {
				log.Errorf("failed to close channel [%s]: %s", name, err.Error())
			}
		}
		if err := client.conn.Close(); err != nil {
			log.Errorf("failed to close RabbitMQ connection: %s", err.Error())
		} else {
			log.Infof("rabbitmq connection closed")
		}
	}

	return client, cleanup, nil
}

// GetChannel 获取指定名称的 channel，如果不存在则创建
func (r *RabbitMQClient) GetChannel(name string) (*amqp.Channel, error) {
	if ch, ok := r.channels[name]; ok {
		return ch, nil
	}

	ch, err := r.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create channel [%s]: %w", name, err)
	}

	// 设置 QoS
	if err := ch.Qos(int(r.conf.PrefetchCount), 0, r.conf.PrefetchGlobal); err != nil {
		return nil, fmt.Errorf("failed to set Qos for channel [%s]: %w", name, err)
	}

	r.channels[name] = ch
	r.log.Infof("rabbitmq: created channel [%s]", name)
	return ch, nil
}

// Publish 发送消息
func (r *RabbitMQClient) Publish(channelName, exchange, routingKey string, body []byte) error {
	ch, err := r.GetChannel(channelName)
	if err != nil {
		return err
	}

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
func (r *RabbitMQClient) Consume(channelName, queue string) (<-chan amqp.Delivery, error) {
	ch, err := r.GetChannel(channelName)
	if err != nil {
		return nil, err
	}

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
