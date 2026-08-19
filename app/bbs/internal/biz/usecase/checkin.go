package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CheckinUsecase struct {
	checkinClient repo.CheckinClient
}

func NewCheckinUsecase(checkinClient repo.CheckinClient) *CheckinUsecase {
	return &CheckinUsecase{checkinClient: checkinClient}
}

func (u *CheckinUsecase) CheckIn(ctx context.Context, userID int64) (*bbsuserv1.CheckIn_Resp, error) {
	checkin, err := u.checkinClient.CheckIn(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.CheckIn_Resp{
		RecordId:      checkin.RecordID,
		Date:          timestamppb.New(*checkin.Date),
		CurrentStreak: checkin.CurrentStreak,
		LongestStreak: checkin.LongestStreak,
	}, nil
}

func (u *CheckinUsecase) GetOverview(ctx context.Context, userID int64, month string) (*bbsuserv1.GetCheckinOverview_Resp, error) {
	overview, err := u.checkinClient.GetOverview(ctx, userID, month)
	if err != nil {
		return nil, err
	}
	response := &bbsuserv1.GetCheckinOverview_Resp{
		Records:       make([]*bbsuserv1.GetCheckinOverview_Resp_Record, 0, len(overview.Records)),
		CurrentStreak: overview.CurrentStreak,
		LongestStreak: overview.LongestStreak,
	}
	for _, record := range overview.Records {
		response.Records = append(response.Records, &bbsuserv1.GetCheckinOverview_Resp_Record{
			Id:      record.ID,
			Date:    timestamppb.New(*record.Date),
			Checked: record.Checked,
		})
	}
	return response, nil
}
