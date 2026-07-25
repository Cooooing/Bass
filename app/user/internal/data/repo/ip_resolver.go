package repo

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	platformv1 "common/proto/gen/platform/v1"
	"context"
	bizrepo "user/internal/biz/repo"
)

var _ bizrepo.IPClient = (*IPClient)(nil)

type IPClient struct{ consul *commonClient.ConsulClient }

func NewIPClient(consul *commonClient.ConsulClient) bizrepo.IPClient {
	return &IPClient{consul: consul}
}

func (r *IPClient) Resolve(ctx context.Context, ip string) (*commonModel.IpInfo, error) {
	conn, err := r.consul.GetGrpcConn(constant.PlatformServiceName.String())
	if err != nil {
		return nil, err
	}
	reply, err := platformv1.NewPlatformIpResolutionServiceClient(conn).ResolveIp(ctx, &platformv1.ResolveIp_Req{Ip: ip})
	if err != nil {
		return nil, err
	}
	return &commonModel.IpInfo{
		Ip:          ip,
		Country:     reply.GetCountry(),
		Province:    reply.GetProvince(),
		City:        reply.GetCity(),
		ISP:         reply.GetIsp(),
		CountryCode: reply.GetCountryCode(),
	}, nil
}
