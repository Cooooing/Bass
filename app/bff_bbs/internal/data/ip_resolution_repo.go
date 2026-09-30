package data

import (
	"bff_bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	commonmodel "common/pkg/model"
	platformv1 "common/proto/gen/platform/v1"
	"context"
)

var _ repo.IPResolutionClient = (*IPResolutionClient)(nil)

type IPResolutionClient struct {
	platformClient *rpc.PlatformClient
}

func NewIPResolutionClient(platformClient *rpc.PlatformClient) repo.IPResolutionClient {
	return &IPResolutionClient{platformClient: platformClient}
}

func (r *IPResolutionClient) Resolve(ctx context.Context, ip string) (*commonmodel.IpInfo, error) {
	resp, err := r.platformClient.IpResolution.ResolveIp(ctx, &platformv1.ResolveIp_Req{Ip: ip})
	if err != nil {
		return nil, err
	}
	return &commonmodel.IpInfo{
		Ip:          ip,
		Country:     resp.GetCountry(),
		Province:    resp.GetProvince(),
		City:        resp.GetCity(),
		ISP:         resp.GetIsp(),
		CountryCode: resp.GetCountryCode(),
	}, nil
}
