package constant

import "fmt"

// RedisKey 定义 Redis 键模板。
var (
	RequestNonce    = "common:auth:request_nonce:{nonce:%s}"                                  // 请求防重放。
	DeadLetterAlert = "common:event:dead_letter_alert:{service:%s}:{source:%s}:{event_id:%s}" // 死信告警去重。
)

func GetKeyRequestNonce(nonce string) string {
	return fmt.Sprintf(RequestNonce, nonce)
}

func GetKeyDeadLetterAlert(service string, source string, eventID string) string {
	return fmt.Sprintf(DeadLetterAlert, service, source, eventID)
}
