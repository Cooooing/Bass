package server

import (
	v1 "common/api/common/v1"
	"common/pkg/constant"
	"common/pkg/util"
	"common/pkg/util/server"

	"context"
	"signal/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
)

var NodeEndpoints = map[string]struct{}{
	"/common.api.signal.v1.SignalNodeService/Register":   {},
	"/common.api.signal.v1.SignalNodeService/Unregister": {},
	"/common.api.signal.v1.SignalNodeService/Online":     {},
	"/common.api.signal.v1.SignalNodeService/Offline":    {},
}

// NodeEndpointsMatch 节点接口鉴权匹配
func NodeEndpointsMatch() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		_, exist := NodeEndpoints[operation]
		return exist
	}
}

func SignalAuthMiddleware(nodeDomain *domain.NodeDomain) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// 验证 signal node
			nodeKey := server.GetHeader(ctx, constant.HeaderSignalNode)
			// nodeSignature := pkg.GetHeader(ctx, constant.HeaderSignalNodeSignature)
			// if nodeKey == "" || nodeSignature == "" {
			//	return nil, v1.ErrorUnauthorized("node is required")
			// }

			node, err := nodeDomain.GetByKey(ctx, nodeKey)
			if err != nil {
				return nil, v1.ErrorUnauthorized("node not found")
			}

			// Todo 通过 secret 进行节点认证
			_ = node

			ctx = util.SetContextValue(ctx, constant.CtxNodeInfo, node)
			return handler(ctx, req)
		}
	}
}
