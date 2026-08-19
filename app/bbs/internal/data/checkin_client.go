package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.CheckinClient = (*CheckinClient)(nil)

type CheckinClient struct {
	userClient *rpc.UserClient
}

func NewCheckinClient(userClient *rpc.UserClient) repo.CheckinClient {
	return &CheckinClient{userClient: userClient}
}

func (c *CheckinClient) CheckIn(ctx context.Context, userID int64) (*repo.Checkin, error) {
	response, err := c.userClient.Checkin.CheckIn(ctx, &userv1.CheckIn_Req{UserId: userID})
	if err != nil {
		return nil, err
	}
	date := response.GetRecord().GetDate().AsTime().UTC()
	return &repo.Checkin{
		RecordID:      response.GetRecord().GetId(),
		Date:          &date,
		CurrentStreak: response.GetStat().GetCurrentStreak(),
		LongestStreak: response.GetStat().GetLongestStreak(),
	}, nil
}

func (c *CheckinClient) GetOverview(ctx context.Context, userID int64, month string) (*repo.CheckinOverview, error) {
	response, err := c.userClient.Checkin.GetOverview(ctx, &userv1.GetCheckinOverview_Req{
		UserId: userID,
		Month:  month,
	})
	if err != nil {
		return nil, err
	}
	overview := &repo.CheckinOverview{
		Records:       make([]*repo.CheckinRecord, 0, len(response.GetRecords())),
		CurrentStreak: response.GetStat().GetCurrentStreak(),
		LongestStreak: response.GetStat().GetLongestStreak(),
	}
	for _, record := range response.GetRecords() {
		date := record.GetDate().AsTime().UTC()
		overview.Records = append(overview.Records, &repo.CheckinRecord{
			ID:      record.GetId(),
			Date:    &date,
			Checked: record.GetChecked(),
		})
	}
	return overview, nil
}
