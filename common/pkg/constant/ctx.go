package constant

type CtxKey string

func (k CtxKey) String() string {
	return string(k)
}

var (
	CtxToken    CtxKey = "CtxToken"    // 上下文token key
	CtxUserInfo CtxKey = "CtxUserInfo" // 上下文用户信息 key
	CtxNodeInfo CtxKey = "CtxNodeInfo" // 上下文节点信息 key
)
