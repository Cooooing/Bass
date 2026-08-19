package service

import (
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CheckinService struct {
	v1.UnimplementedCheckinServiceServer
	checkinUsecase *usecase.CheckinUsecase
}

func NewCheckinService(
	checkinUsecase *usecase.CheckinUsecase,
) *CheckinService {
	return &CheckinService{checkinUsecase: checkinUsecase}
}

func (s *CheckinService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterCheckinServiceServer(gs, s)
}

func (s *CheckinService) RegisterHttp(hs *http.Server) {
}

func (s *CheckinService) CheckIn(ctx context.Context, req *v1.CheckIn_Req) (*v1.CheckIn_Resp, error) {
	response, err := s.checkinUsecase.CheckIn(ctx, &usecase.CheckInReq{UserID: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	return &v1.CheckIn_Resp{
		Record: &v1.CheckinRecord{
			Id:      response.Record.ID,
			UserId:  response.Record.UserID,
			Date:    timestamppb.New(*response.Record.Date),
			Checked: response.Record.Checked,
		},
		Stat: &v1.CheckinStat{
			UserId:        response.Stat.UserID,
			CurrentStreak: *response.Stat.CurrentStreak,
			LongestStreak: *response.Stat.LongestStreak,
		},
	}, nil
}

func (s *CheckinService) GetOverview(ctx context.Context, req *v1.GetCheckinOverview_Req) (*v1.GetCheckinOverview_Resp, error) {
	response, err := s.checkinUsecase.GetOverview(ctx, &usecase.GetCheckinOverviewReq{
		UserID: req.GetUserId(),
		Month:  req.GetMonth(),
	})
	if err != nil {
		return nil, err
	}
	records := make([]*v1.CheckinRecord, 0, len(response.Records))
	for _, record := range response.Records {
		records = append(records, &v1.CheckinRecord{
			Id:      record.ID,
			UserId:  record.UserID,
			Date:    timestamppb.New(*record.Date),
			Checked: record.Checked,
		})
	}
	return &v1.GetCheckinOverview_Resp{
		Records: records,
		Stat: &v1.CheckinStat{
			UserId:        response.Stat.UserID,
			CurrentStreak: *response.Stat.CurrentStreak,
			LongestStreak: *response.Stat.LongestStreak,
		},
	}, nil
}
