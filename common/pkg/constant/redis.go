package constant

import "fmt"

// RedisKey 定义 Redis 键模板。
var (
	RequestNonce        = "common:auth:request_nonce:{nonce:%s}"                                  // 请求防重放。
	AsynqTaskVersion    = "common:asynq:task_version"                                             // Asynq 任务版本映射。
	OutboxPublisherLock = "common:event:outbox_publisher_lock:{service:%s}"                       // outbox 单轮投递锁。
	DeadLetterAlert     = "common:event:dead_letter_alert:{service:%s}:{source:%s}:{event_id:%s}" // 死信告警去重。
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
