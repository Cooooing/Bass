package server

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// GrpcService 接口，每个 service 实现它
type GrpcService interface {
	RegisterGrpc(gs *grpc.Server)
}

// HttpService 接口，每个 service 实现它
type HttpService interface {
	RegisterHttp(hs *http.Server)
}
