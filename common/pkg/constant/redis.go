package constant

import "fmt"

// RedisKey 定义 Redis 键模板。
var (
	RequestNonce        = "Auth:RequestNonce:{%s}"               // 请求防重放。
	AsynqTaskVersion    = "Asynq:TaskVersion"                    // Asynq 任务版本映射。
	OutboxPublisherLock = "Event:OutboxPublisherLock:{%s}"       // outbox 单轮投递锁。
	DeadLetterAlert     = "Event:DeadLetterAlert:{%s}:{%s}:{%s}" // 死信告警去重。
)

func GetKeyRequestNonce(nonce string) string {
	return fmt.Sprintf(RequestNonce, nonce)
}

func GetKeyOutboxPublisherLock(service string) string {
	return fmt.Sprintf(OutboxPublisherLock, service)
}

func GetKeyDeadLetterAlert(service string, source string, eventID string) string {
	return fmt.Sprintf(DeadLetterAlert, service, source, eventID)
}
