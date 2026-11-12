package constant

import "fmt"

// RedisKey 定义 Redis 键模板。
var (
	RequestNonce = "common:auth:request_nonce:{nonce:%s}" // 请求防重放。
)

func GetKeyRequestNonce(nonce string) string {
	return fmt.Sprintf(RequestNonce, nonce)
}
