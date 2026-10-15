package data

import (
	"common/pkg/client/rpc"
	platformv1 "common/proto/gen/platform/v1"
	platformv1enum "common/proto/gen/platform/v1/enum"
	"context"

	"bff_bbs/internal/biz/model"
	bizrepo "bff_bbs/internal/biz/repo"
	"bff_bbs/internal/enum"
)

var _ bizrepo.HolidayClient = (*HolidayRepo)(nil)

type HolidayRepo struct {
	platformClient *rpc.PlatformClient
}

func NewHolidayRepo(platformClient *rpc.PlatformClient) bizrepo.HolidayClient {
	return &HolidayRepo{platformClient: platformClient}
}

func (r *HolidayRepo) GetCurrent(ctx context.Context) (*model.HolidayInfo, error) {
	resp, err := r.platformClient.Holiday.GetCurrent(ctx, &platformv1.GetCurrentHoliday_Req{})
	if err != nil {
		return nil, err
	}
	return &model.HolidayInfo{
		Date:          resp.GetDate().AsTime(),
		Type:          r.holidayTypeFromPlatform(resp.GetType()),
		DayName:       resp.GetDayName(),
		HolidayName:   resp.GetHolidayName(),
		DaysToHoliday: resp.GetDaysToHoliday(),
		DaysRemaining: resp.GetDaysRemaining(),
	}, nil
}

func (r *HolidayRepo) holidayTypeFromPlatform(value platformv1enum.HolidayType) enum.HolidayType {
	switch value {
	case platformv1enum.HolidayType_HOLIDAY_TYPE_WORKDAY:
		return enum.HolidayTypeWorkday
	case platformv1enum.HolidayType_HOLIDAY_TYPE_WEEKEND:
		return enum.HolidayTypeWeekend
	case platformv1enum.HolidayType_HOLIDAY_TYPE_FESTIVAL:
		return enum.HolidayTypeFestival
	case platformv1enum.HolidayType_HOLIDAY_TYPE_MAKEUP_WORKDAY:
		return enum.HolidayTypeMakeupWorkday
	default:
		return ""
	}
}
