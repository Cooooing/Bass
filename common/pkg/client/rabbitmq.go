package client

import (
	"common/api/gen/common"
	"common/pkg/constant"
	"encoding/json"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	log    *log.Helper
	conf   *common.RabbitMQ
	conn   *amqp.Connection
	chPool chan *amqp.Channel // 固定大小 channel 池
}

// NewRabbitMQClient 初始化 RabbitMQ 单机客户端
func NewRabbitMQClient(log *log.Helper, conf *common.RabbitMQ) (*RabbitMQClient, func(), error) {
	conn, err := amqp.DialConfig(conf.Url, amqp.Config{
		Heartbeat: conf.Heartbeat.AsDuration(),
		Dial:      amqp.DefaultDial(conf.DialTimeout.AsDuration()),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	client := &RabbitMQClient{
		log:    log,
		conn:   conn,
		conf:   conf,
		chPool: make(chan *amqp.Channel, 16), // 固定大小 channel 池
	}

	// 初始化队列和交换机
	err = client.declareResources()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to declare resources: %w", err)
	}

	log.Infof("rabbitmq: connected to [%s]", conn.RemoteAddr().String())

	// 清理函数
	cleanup := func() {
		// 关闭所有 channel
		close(client.chPool)
		for ch := range client.chPool {
			if ch != nil && !ch.IsClosed() {
				_ = ch.Close()
			}
		}
		// 关闭连接
		if err := conn.Close(); err != nil {
			log.Errorf("failed to close RabbitMQ connection: %v", err)
		} else {
			log.Infof("rabbitmq connection closed")
		}
	}

	return client, cleanup, nil
}

// newChannel 创建一个新的 channel 并设置 Qos
func (r *RabbitMQClient) newChannel() (*amqp.Channel, error) {
	ch, err := r.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %s", err)
	}
	// 设置 Qos
	if err := ch.Qos(int(r.conf.PrefetchCount), 0, r.conf.PrefetchGlobal); err != nil {
		return nil, fmt.Errorf("failed to set qos: %s", err)
	}
	return ch, nil
}

// GetChannel 从池中获取一个 channel
func (r *RabbitMQClient) GetChannel() (*amqp.Channel, error) {
	// 从池中获取 channel
	select {
	case ch := <-r.chPool:
		// 如果有空闲通道，直接返回
		return ch, nil
	default:
		// 如果池已满，尝试创建新的通道
		ch, err := r.newChannel()
		if err != nil {
			return nil, err
		}
		return ch, nil
	}
}

// ReleaseChannel 放回 channel 到池中
func (r *RabbitMQClient) ReleaseChannel(ch *amqp.Channel) {
	if ch != nil && !ch.IsClosed() {
		// 如果池中有空闲空间，则将通道放回池中
		select {
		case r.chPool <- ch:
		default:
			// 如果池已满，关闭通道
			_ = ch.Close()
		}
	}
}

func (r *RabbitMQClient) declareResources() error {
	ch, err := r.newChannel()
	if err != nil {
		return err
	}
	defer func(ch *amqp.Channel) {
		_ = ch.Close()
	}(ch)

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
func (r *RabbitMQClient) Publish(exchange, routingKey string, body any) error {
	if exchange == "" {
		return fmt.Errorf("exchange cannot be empty")
	}
	if routingKey == "" {
		return fmt.Errorf("routingKey cannot be empty")
	}

	// 获取一个通道
	ch, err := r.GetChannel()
	if err != nil {
		return err
	}
	defer r.ReleaseChannel(ch)

	// 序列化
	marshal, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return ch.Publish(
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         marshal,
			DeliveryMode: uint8(r.conf.DeliveryMode), // 消息持久化
		},
	)
}

// Consume 消费消息
func (r *RabbitMQClient) Consume(queue string) (<-chan amqp.Delivery, *amqp.Channel, error) {
	// 检查队列是否为空
	if queue == "" {
		return nil, nil, fmt.Errorf("queue cannot be empty")
	}

	// 获取一个通道
	ch, err := r.GetChannel()
	if err != nil {
		return nil, nil, err
	}

	// 开始消费
	msgs, err := ch.Consume(
		queue,
		"", // consumerTag
		r.conf.AutoAck,
		false, // exclusive
		false, // noLocal
		false, // noWait
		nil,   // args
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to consume queue [%s]: %w", queue, err)
	}

	return msgs, ch, nil
}
