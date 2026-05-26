package constant

import (
	"common/api/gen/common"
)

// 服务启动模式定义
const (
	Dev  = "dev"
	Test = "test"
	Prod = "prod"
)

// 请求头 key 定义
const (
	HeaderTimestamp      = "X-Timestamp"   // 时间戳，防止过期请求
	HeaderNonce          = "X-Nonce"       // 随机数，防止重放攻击
	HeaderAuthentication = "Authorization" // 令牌请求头名称
)

const (
	defaultPage uint32 = 1
	defaultSize uint32 = 10
	maxPageSize uint32 = 1000
)

func PageValid(p *common.PageRequest) *common.PageRequest {
	if p == nil {
		return GetPageDefault()
	}
	if p.Page <= 0 {
		p.Page = defaultPage
	}
	if p.Size <= 0 {
		p.Size = defaultSize
	}
	if p.Size > maxPageSize {
		p.Size = maxPageSize
	}
	return p
}
func GetPageDefault() *common.PageRequest {
	return &common.PageRequest{
		Page: defaultPage,
		Size: defaultSize,
	}
}
func GetPageMax() *common.PageRequest {
	return &common.PageRequest{
		Page: defaultPage,
		Size: maxPageSize,
	}
}
