package server

import (
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"
	"context"
	"strings"
)

// ClientIP 返回请求的客户端 IP，优先使用上下文中已解析的 IP 信息。
func ClientIP(
	ctx context.Context,
) string {
	if ipInfo, ok := util.GetContextValue[*model.IpInfo](ctx, constant.CtxIpInfo); ok && ipInfo != nil {
		if ip := strings.TrimSpace(ipInfo.Ip); ip != "" {
			return ip
		}
	}
	for _, header := range [...]string{constant.HeaderForwardedFor, constant.HeaderRealIP, constant.HeaderClientIP} {
		for _, item := range strings.Split(GetHeader(ctx, header), ",") {
			if ip := strings.TrimSpace(item); ip != "" {
				return ip
			}
		}
	}
	return ""
}
