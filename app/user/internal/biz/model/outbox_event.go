package model

import commonenum "common/pkg/enum"

type OutboxEvent struct {
	ID      int64
	EventID string
	Subject commonenum.EventSubject
	Payload string
	Headers map[string]string
}
