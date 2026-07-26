package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type LocationService struct {
	bbsuserv1.UnimplementedLocationServiceServer
	locationUsecase *usecase.LocationUsecase
}

func NewLocationService(
	locationUsecase *usecase.LocationUsecase,
) *LocationService {
	return &LocationService{
		locationUsecase: locationUsecase,
	}
}

func (s *LocationService) RegisterGrpc(gs *grpc.Server) {
}

func (s *LocationService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterLocationServiceHTTPServer(hs, s)
}

func (s *LocationService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Req) (*bbsuserv1.GetCurrentLocation_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	location, err := s.locationUsecase.GetCurrentLocation(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentLocation_Resp{
		Location: location,
	}, nil
}

func (s *LocationService) UpsertCurrent(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Req) (*bbsuserv1.UpsertCurrentLocation_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	location, err := s.locationUsecase.UpsertCurrentLocation(ctx, &usecase.UpsertCurrentLocationReq{
		UserID:   user.ID,
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpsertCurrentLocation_Resp{
		Location: location,
	}, nil
}
