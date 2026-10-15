package service

import (
	platformv1 "common/proto/gen/bff_bbs/v1/platform"
	"context"

	"bff_bbs/internal/biz/usecase"
	"bff_bbs/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// HolidayService 实现 BFF 节假日 HTTP 服务。
type HolidayService struct {
	platformv1.UnimplementedHolidayServiceServer
	holidayUsecase *usecase.HolidayUsecase
}

func NewHolidayService(
	holidayUsecase *usecase.HolidayUsecase,
) *HolidayService {
	return &HolidayService{
		holidayUsecase: holidayUsecase,
	}
}

func (s *HolidayService) RegisterGrpc(gs *grpc.Server) {
}

func (s *HolidayService) RegisterHTTP(hs *http.Server) {
	platformv1.RegisterHolidayServiceHTTPServer(hs, s)
}

func (s *HolidayService) GetCurrent(ctx context.Context, req *platformv1.GetCurrentHoliday_Req) (*platformv1.GetCurrentHoliday_Resp, error) {
	info, err := s.holidayUsecase.GetCurrent(ctx)
	if err != nil {
		return nil, err
	}
	return &platformv1.GetCurrentHoliday_Resp{
		Date:          timestamppb.New(info.Date),
		Type:          enum.HolidayTypeMap.MustToProto(info.Type),
		DayName:       info.DayName,
		HolidayName:   info.HolidayName,
		DaysToHoliday: info.DaysToHoliday,
		DaysRemaining: info.DaysRemaining,
	}, nil
}
