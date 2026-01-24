package service

import (
	v1 "common/api/connector/v1"
	"connector/internal/biz"
	"connector/internal/data"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type HealthService struct {
	v1.UnimplementedConnectorHealthServiceServer
	*BaseService
}

func NewHealthService(baseService *BaseService, baseDomain *biz.BaseDomain, repo *data.BaseRepo) *HealthService {
	return &HealthService{
		BaseService: baseService,
	}
}

func (s *HealthService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterConnectorHealthServiceServer(gs, s)
}

func (s *HealthService) RegisterHttp(hs *http.Server) {
	v1.RegisterConnectorHealthServiceHTTPServer(hs, s)
}

func (s *HealthService) Ping(ctx context.Context, req *v1.PingRequest) (rsp *v1.PingResponse, err error) {
	return &v1.PingResponse{}, nil
}

func (s *HealthService) Pow(ctx context.Context, req *v1.PowRequest) (rsp *v1.PowResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for n := 0; ; n++ {
		msg := req.Challenge + ":" + strconv.Itoa(n)
		sum := sha256.Sum256([]byte(msg))
		h := hex.EncodeToString(sum[:])

		if strings.HasPrefix(h, strings.Repeat("0", int(req.Difficulty))) {
			return &v1.PowResponse{
				Nonce:   strconv.Itoa(n),
				HashHex: h,
			}, nil
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}
}
