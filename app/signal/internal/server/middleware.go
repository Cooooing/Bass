package server

import (
	v1 "common/api/common/v1"
	"common/pkg"
	"common/pkg/constant"
	"common/pkg/util"
	"context"
	"fmt"
	"signal/internal/biz/repo"
	"signal/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/middleware"
)

func SignalAuthMiddleware(db *gen.Client, nodeRepo repo.NodeRepo) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// 验证 signal node
			nodeKey := pkg.GetHeader(ctx, constant.HeaderSignalNode)
			fmt.Printf("nodeKey: %s\n", nodeKey)
			// nodeSignature := pkg.GetHeader(ctx, constant.HeaderSignalNodeSignature)
			// if nodeKey == "" || nodeSignature == "" {
			//	return nil, v1.ErrorUnauthorized("node is required")
			// }

			node, err := nodeRepo.GetByKey(ctx, db, nodeKey)
			fmt.Printf("node: %+v\n", node)
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
