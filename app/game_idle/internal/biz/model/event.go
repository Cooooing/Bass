package model

// GameIdleEventMessage 是已经编码完成、可直接发布到消息队列的事件消息。
type GameIdleEventMessage struct {
	Subject string
	Data    []byte
	Header  map[string]string
}
