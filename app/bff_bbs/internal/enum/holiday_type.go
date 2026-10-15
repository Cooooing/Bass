package enum

import (
	commonenum "common/pkg/enum"
	platformv1enum "common/proto/gen/bff_bbs/v1/platform/enum"
)

type HolidayType string

const (
	HolidayTypeWorkday       HolidayType = "workday"        // workday
	HolidayTypeWeekend       HolidayType = "weekend"        // weekend
	HolidayTypeFestival      HolidayType = "festival"       // festival
	HolidayTypeMakeupWorkday HolidayType = "makeup_workday" // makeup_workday
)

var HolidayTypeMap = commonenum.NewMapping[HolidayType, platformv1enum.HolidayType](map[HolidayType]commonenum.Entry[HolidayType, platformv1enum.HolidayType]{
	HolidayTypeWorkday:       {Proto: platformv1enum.HolidayType_HOLIDAY_TYPE_WORKDAY},
	HolidayTypeWeekend:       {Proto: platformv1enum.HolidayType_HOLIDAY_TYPE_WEEKEND},
	HolidayTypeFestival:      {Proto: platformv1enum.HolidayType_HOLIDAY_TYPE_FESTIVAL},
	HolidayTypeMakeupWorkday: {Proto: platformv1enum.HolidayType_HOLIDAY_TYPE_MAKEUP_WORKDAY},
})

func (e HolidayType) String() string {
	return string(e)
}
