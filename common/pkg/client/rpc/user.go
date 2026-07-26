package rpc

import (
	userv1 "common/proto/gen/user/v1"

	"google.golang.org/grpc"
)

type UserClient struct {
	Account        userv1.AccountServiceClient
	Auth           userv1.AuthServiceClient
	Rbac           userv1.RbacServiceClient
	Relation       userv1.RelationServiceClient
	Preferences    userv1.PreferencesServiceClient
	PrivacySetting userv1.PrivacySettingServiceClient
	Location       userv1.LocationServiceClient
	Otp            userv1.OtpServiceClient
}

func NewUserClient(
	conn *grpc.ClientConn,
) *UserClient {
	return &UserClient{
		Account:        userv1.NewAccountServiceClient(conn),
		Auth:           userv1.NewAuthServiceClient(conn),
		Rbac:           userv1.NewRbacServiceClient(conn),
		Relation:       userv1.NewRelationServiceClient(conn),
		Preferences:    userv1.NewPreferencesServiceClient(conn),
		PrivacySetting: userv1.NewPrivacySettingServiceClient(conn),
		Location:       userv1.NewLocationServiceClient(conn),
		Otp:            userv1.NewOtpServiceClient(conn),
	}
}
