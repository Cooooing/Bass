package constant

import (
	v1 "common/gen/common/v1"
	"math"
)

// 服务启动模式定义
var (
	Dev  = "dev"
	Test = "test"
	Prod = "prod"
)

// 请求头 key 定义
var (
	HeaderTimestamp           = "X-Timestamp"            // 时间戳，防止过期请求
	HeaderNonce               = "X-Nonce"                // 随机数，防止重放攻击
	HeaderAuthentication      = "Authorization"          // token 请求头名称
	HeaderSignalNode          = "X-Signal-NodeKey"       // 信令服务节点 key
	HeaderSignalNodeSignature = "X-Signal-NodeSignature" // 节点签名
)

var page uint32 = 1
var size uint32 = 10

func PageValid(p *v1.PageRequest) *v1.PageRequest {
	if p == nil {
		return GetPageDefault()
	}
	if p.Page <= 0 {
		p.Page = page
	}
	if p.Size <= 0 {
		p.Size = size
	}
	return p
}
func GetPageDefault() *v1.PageRequest {
	return &v1.PageRequest{
		Page: page,
		Size: size,
	}
}
func GetPageMax() *v1.PageRequest {
	return &v1.PageRequest{
		Page: page,
		Size: math.MaxUint32,
	}
}
