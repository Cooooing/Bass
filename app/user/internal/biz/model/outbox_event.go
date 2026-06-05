package model

import commonenum "common/pkg/enum"

type OutboxEvent struct {
	ID      int64
	EventID string
	Subject commonenum.EventSubject
	Payload []byte
	Headers map[string]string
}
