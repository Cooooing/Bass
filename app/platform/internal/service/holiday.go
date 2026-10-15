package service

import (
	v1 "common/proto/gen/platform/v1"
	"context"
	"time"

	"platform/internal/biz/usecase"
	"platform/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// HolidayService 实现 platform 节假日 gRPC 服务。
type HolidayService struct {
	v1.UnimplementedPlatformHolidayServiceServer
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
	v1.RegisterPlatformHolidayServiceServer(gs, s)
}

func (s *HolidayService) RegisterHTTP(hs *http.Server) {
}

func (s *HolidayService) GetCurrent(ctx context.Context, req *v1.GetCurrentHoliday_Req) (*v1.GetCurrentHoliday_Resp, error) {
	info, err := s.holidayUsecase.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.GetCurrentHoliday_Resp{
		Date:          timestamppb.New(info.Date),
		Type:          enum.HolidayTypeMap.MustToProto(info.Type),
		DayName:       info.DayName,
		HolidayName:   info.HolidayName,
		DaysToHoliday: info.DaysToHoliday,
		DaysRemaining: info.DaysRemaining,
	}, nil
}

func (s *HolidayService) RefreshHoliday(ctx context.Context, req *v1.RefreshHoliday_Req) (*v1.RefreshHoliday_Resp, error) {
	var start, end time.Time
	if req.GetStart() != nil {
		start = req.GetStart().AsTime()
	}
	if req.GetEnd() != nil {
		end = req.GetEnd().AsTime()
	}
	if err := s.holidayUsecase.Refresh(ctx, start, end); err != nil {
		return nil, err
	}
	return &v1.RefreshHoliday_Resp{}, nil
}
