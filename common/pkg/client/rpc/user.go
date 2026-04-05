package rpc

import (
	userv1 "common/gen/user/v1"

	"google.golang.org/grpc"
)

type UserClient struct {
	User     userv1.UserUserServiceClient
	Auth     userv1.UserAuthenticationServiceClient
	Relation userv1.UserUserRelationServiceClient
	TwoFA    userv1.UserTwoFactorAuthenticationServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		User:     userv1.NewUserUserServiceClient(conn),
		Auth:     userv1.NewUserAuthenticationServiceClient(conn),
		Relation: userv1.NewUserUserRelationServiceClient(conn),
		TwoFA:    userv1.NewUserTwoFactorAuthenticationServiceClient(conn),
	}
}
