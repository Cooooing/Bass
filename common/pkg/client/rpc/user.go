package rpc

import (
	userv1 "common/api/gen/user/v1"

	"google.golang.org/grpc"
)

type UserClient struct {
	User     userv1.UserUserServiceClient
	Auth     userv1.UserAuthServiceClient
	Relation userv1.UserUserRelationServiceClient
	TwoFA    userv1.UserTwoFactorAuthServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		User:     userv1.NewUserUserServiceClient(conn),
		Auth:     userv1.NewUserAuthServiceClient(conn),
		Relation: userv1.NewUserUserRelationServiceClient(conn),
		TwoFA:    userv1.NewUserTwoFactorAuthServiceClient(conn),
	}
}
