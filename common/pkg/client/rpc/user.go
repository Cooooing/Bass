package rpc

import (
	"common/pkg/client/localrpc"
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
	Outbox         userv1.OutboxServiceClient
	Checkin        userv1.CheckinServiceClient
}

func NewUserClient(
	conn grpc.ClientConnInterface,
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
		Outbox:         userv1.NewOutboxServiceClient(conn),
		Checkin:        userv1.NewCheckinServiceClient(conn),
	}
}

func NewLocalUserClient[T any](services []T) *UserClient {
	conn := localrpc.NewConn()
	MountUserServices(conn, services)
	return NewUserClient(conn)
}

func MountUserServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&userv1.AuthService_ServiceDesc, service)
		conn.RegisterMatching(&userv1.AccountService_ServiceDesc, service)
		conn.RegisterMatching(&userv1.RbacService_ServiceDesc, service)
		conn.RegisterMatching(&userv1.RelationService_ServiceDesc, service)
		conn.RegisterMatching(&userv1.PreferencesService_ServiceDesc, service)
		conn.RegisterMatching(&userv1.PrivacySettingService_ServiceDesc, service)
		conn.RegisterMatching(&userv1.LocationService_ServiceDesc, service)
		conn.RegisterMatching(&userv1.OtpService_ServiceDesc, service)
		conn.RegisterMatching(&userv1.OutboxService_ServiceDesc, service)
		conn.RegisterMatching(&userv1.CheckinService_ServiceDesc, service)
	}
}
