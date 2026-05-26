package constant

type CtxKey string

func (k CtxKey) String() string {
	return string(k)
}

var (
	CtxToken    CtxKey = "CtxToken"    // 上下文令牌键
	CtxUserInfo CtxKey = "CtxUserInfo" // 上下文用户信息 key
	CtxIpInfo   CtxKey = "CtxIpInfo"   // 上下文 IP 信息键
)
