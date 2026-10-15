package usecase

import (
	"context"
	"strings"

	platformv1 "common/proto/gen/bff_bbs/v1/platform"

	"bff_bbs/internal/biz/repo"
)

type LocationUsecase struct {
	ipResolutionClient repo.IPResolutionClient
}

func NewLocationUsecase(
	ipResolutionClient repo.IPResolutionClient,
) *LocationUsecase {
	return &LocationUsecase{
		ipResolutionClient: ipResolutionClient,
	}
}

func (u *LocationUsecase) GetCurrent(ctx context.Context, ip string) (*platformv1.GetCurrentLocation_Resp_Location, error) {
	if u.ipResolutionClient == nil || strings.TrimSpace(ip) == "" {
		return nil, nil
	}
	resolved, err := u.ipResolutionClient.Resolve(ctx, strings.TrimSpace(ip))
	if err != nil {
		return nil, err
	}
	if resolved == nil {
		return nil, nil
	}

	return &platformv1.GetCurrentLocation_Resp_Location{
		Country:  new(strings.TrimSpace(resolved.Country)),
		Province: new(strings.TrimSpace(resolved.Province)),
		City:     new(strings.TrimSpace(resolved.City)),
	}, nil
}
