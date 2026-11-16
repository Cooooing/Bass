package nats

import (
	"context"
	"time"

	"github.com/nats-io/nats.go/jetstream"
)

type Message struct {
	ID       string
	Subject  string
	Data     []byte
	Header   map[string]string
	Metadata *MessageMetadata
	raw      jetstream.Msg
	settled  bool
}

type MessageMetadata struct {
	Stream       string
	Sequence     uint64
	Consumer     string
	NumDelivered uint64
	Timestamp    time.Time
}

type MessageHandler func(context.Context, *Message) error

func (m *Message) Ack() error {
	if m.raw == nil {
		return nil
	}
	m.settled = true
	return m.raw.Ack()
}

func (m *Message) Nack() error {
	if m.raw == nil {
		return nil
	}
	m.settled = true
	return m.raw.Nak()
}

func (m *Message) NackWithDelay(delay time.Duration) error {
	if m.raw == nil {
		return nil
	}
	m.settled = true
	return m.raw.NakWithDelay(delay)
}

func (m *Message) Term() error {
	if m.raw == nil {
		return nil
	}
	m.settled = true
	return m.raw.Term()
}

func (m *Message) InProgress() error {
	if m.raw == nil {
		return nil
	}
	return m.raw.InProgress()
}
