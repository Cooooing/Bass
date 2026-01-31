package service

import (
	v1 "common/api/connector/v1"
	"connector/internal/biz"
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
	v1.UnimplementedConnectorServiceServer
	*BaseService
	*biz.SessionDomain
}

func NewCallbackService(baseService *BaseService, sessionDomain *biz.SessionDomain) *CallbackService {
	return &CallbackService{
		BaseService:   baseService,
		SessionDomain: sessionDomain,
	}
}

func (s *CallbackService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterConnectorServiceServer(gs, s)
}

func (s *CallbackService) RegisterHttp(hs *http.Server) {
	v1.RegisterConnectorServiceHTTPServer(hs, s)
}

func (s *CallbackService) Ping(ctx context.Context, req *v1.PingRequest) (rsp *v1.PingReply, err error) {
	return &v1.PingReply{}, nil
}

func (s *CallbackService) Pow(ctx context.Context, req *v1.PowRequest) (rsp *v1.PowReply, err error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for n := 0; ; n++ {
		msg := req.Challenge + ":" + strconv.Itoa(n)
		sum := sha256.Sum256([]byte(msg))
		h := hex.EncodeToString(sum[:])

		if strings.HasPrefix(h, strings.Repeat("0", int(req.Difficulty))) {
			return &v1.PowReply{
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

func (s *CallbackService) Session(ctx context.Context, req *v1.SessionRequest) (rsp *v1.SessionReply, err error) {
	return &v1.SessionReply{
		SessionIds: s.GetSessionIds(),
	}, nil
}

func (s *CallbackService) Send(ctx context.Context, req *v1.SendRequest) (rsp *v1.SendReply, err error) {
	for _, m := range req.Messages {
		err := s.SessionDomain.SendMessage(m.SessionId, m.Payload)
		if err != nil {
			return nil, err
		}
	}
	return &v1.SendReply{}, nil
}
