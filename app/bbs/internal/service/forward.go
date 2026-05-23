package service

import (
	"common/pkg/constant"
	"context"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// forwardAuth 将入口请求中的认证头透传给内部 RPC。
func forwardAuth(ctx context.Context) context.Context {
	tr, ok := transport.FromServerContext(ctx)
	if !ok || tr.RequestHeader() == nil {
		return ctx
	}
	auth := tr.RequestHeader().Get(constant.HeaderAuthentication)
	if auth == "" {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, strings.ToLower(constant.HeaderAuthentication), auth)
}

// formatProtoTime 将内部 protobuf 时间转换为 BFF 对外稳定字符串。
func formatProtoTime(ts *timestamppb.Timestamp) string {
	if ts == nil {
		return ""
	}
	return ts.AsTime().Format(time.RFC3339Nano)
}
