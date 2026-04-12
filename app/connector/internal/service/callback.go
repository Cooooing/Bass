package service

import (
	v1 "common/api/gen/connector/v1"
	"connector/internal/biz/domain"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type CallbackService struct {
	v1.UnimplementedConnectorCallbackServiceServer
	*BaseService
	*domain.SessionDomain
}

func NewCallbackService(baseService *BaseService, sessionDomain *domain.SessionDomain) *CallbackService {
	return &CallbackService{
		BaseService:   baseService,
		SessionDomain: sessionDomain,
	}
}

func (s *CallbackService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterConnectorCallbackServiceServer(gs, s)
}

func (s *CallbackService) RegisterHttp(hs *http.Server) {
	v1.RegisterConnectorCallbackServiceHTTPServer(hs, s)
}

func (s *CallbackService) Ping(ctx context.Context, req *v1.PingConnector_Request) (rsp *v1.PingConnector_Reply, err error) {
	return &v1.PingConnector_Reply{}, nil
}

func (s *CallbackService) Pow(ctx context.Context, req *v1.PowConnector_Request) (rsp *v1.PowConnector_Reply, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for n := 0; ; n++ {
		msg := req.Challenge + ":" + strconv.Itoa(n)
		sum := sha256.Sum256([]byte(msg))
		h := hex.EncodeToString(sum[:])

		if strings.HasPrefix(h, strings.Repeat("0", int(req.Difficulty))) {
			return &v1.PowConnector_Reply{
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

func (s *CallbackService) Session(ctx context.Context, req *v1.SessionConnector_Request) (rsp *v1.SessionConnector_Reply, err error) {
	return &v1.SessionConnector_Reply{
		SessionIds: s.GetSessionIds(),
	}, nil
}

func (s *CallbackService) Send(ctx context.Context, req *v1.SendConnector_Request) (rsp *v1.SendConnector_Reply, err error) {
	for _, m := range req.Messages {
		err := s.SessionDomain.SendMessage(ctx, m.SessionId, m.Payload)
		if err != nil {
			return nil, err
		}
	}
	return &v1.SendConnector_Reply{}, nil
}
