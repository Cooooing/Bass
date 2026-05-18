package domain

import (
	"common/pkg/client/rpc"
	"context"
	"notify/internal/biz/domain/sender"

	"github.com/go-kratos/kratos/v2/log"
)

// RPCUserResolver 通过 gRPC 获取用户联系信息
type RPCUserResolver struct {
	log        *log.Helper
	userClient *rpc.UserClient
}

func NewRPCUserResolver(logger log.Logger, userClient *rpc.UserClient) *RPCUserResolver {
	return &RPCUserResolver{
		log:        log.NewHelper(logger),
		userClient: userClient,
	}
}

func (r *RPCUserResolver) Resolve(ctx context.Context, userID int64) (*sender.UserInfo, error) {
	if r.userClient == nil {
		return &sender.UserInfo{}, nil
	}

	// TODO: 调用 user RPC 获取手机号、邮箱等
	// resp, err := r.userClient.User.GetUser(ctx, &userv1.GetUserRequest{Id: userID})
	// if err != nil {
	//     return nil, err
	// }
	// return &sender.UserInfo{
	//     Phone: resp.Phone,
	//     Email: resp.Email,
	// }, nil

	return &sender.UserInfo{}, nil
}
