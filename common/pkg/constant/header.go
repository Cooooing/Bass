package constant

// 请求元信息头名称定义。
const (
	HeaderUserAgent      = "User-Agent"
	HeaderDeviceID       = "X-Device-ID"
	HeaderPlatform       = "X-Platform"
	HeaderRequestID      = "X-Request-ID"
	HeaderTraceID        = "X-Trace-ID"
	HeaderForwardedFor   = "X-Forwarded-For"
	HeaderRealIP         = "X-Real-IP"
	HeaderClientIP       = "X-Client-IP"
	HeaderTimestamp      = "X-Timestamp"   // 时间戳，防止过期请求
	HeaderNonce          = "X-Nonce"       // 随机数，防止重放攻击
	HeaderAuthentication = "Authorization" // 令牌请求头名称
	HeaderBassAppName    = "X-Bass-App-Name" // 客户端应用名称
	HeaderBassAppVersion = "X-Bass-App-Version" // 客户端应用版本
)
