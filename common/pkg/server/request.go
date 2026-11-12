package server

import (
	"common/pkg/constant"
	"context"
	"strings"
)

// ClientIP 返回请求的客户端 IP。
func ClientIP(ctx context.Context) string {
	for _, header := range [...]string{constant.HeaderForwardedFor, constant.HeaderRealIP, constant.HeaderClientIP} {
		for item := range strings.SplitSeq(GetHeader(ctx, header), ",") {
			if ip := strings.TrimSpace(item); ip != "" {
				return ip
			}
		}
	}
	return ""
}
