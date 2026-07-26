package repo

import (
	"common/pkg/client/rpc"
	commonModel "common/pkg/model"
	platformv1 "common/proto/gen/platform/v1"
	"context"
	bizrepo "user/internal/biz/repo"
)

var _ bizrepo.IPClient = (*IPClient)(nil)

type IPClient struct{ platformClient *rpc.PlatformClient }

func NewIPClient(
	platformClient *rpc.PlatformClient,
) bizrepo.IPClient {
	return &IPClient{
		platformClient: platformClient,
	}
}

func (r *IPClient) Resolve(ctx context.Context, ip string) (*commonModel.IpInfo, error) {
	reply, err := r.platformClient.IpResolution.ResolveIp(ctx, &platformv1.ResolveIp_Req{
		Ip: ip,
	})
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
