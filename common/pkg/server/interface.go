package server

import (
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// Service 注册一个服务的 HTTP 和 gRPC 传输入口。
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(hs *http.Server)
}
