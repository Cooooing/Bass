package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/platform/v1/enum"
)

type HolidayType string

const (
	HolidayTypeWorkday       HolidayType = "workday"        // workday
	HolidayTypeWeekend       HolidayType = "weekend"        // weekend
	HolidayTypeFestival      HolidayType = "festival"       // festival
	HolidayTypeMakeupWorkday HolidayType = "makeup_workday" // makeup_workday
)

var HolidayTypeMap = commonenum.NewMapping[HolidayType, v1.HolidayType](map[HolidayType]commonenum.Entry[HolidayType, v1.HolidayType]{
	HolidayTypeWorkday:       {Proto: v1.HolidayType_HOLIDAY_TYPE_WORKDAY},
	HolidayTypeWeekend:       {Proto: v1.HolidayType_HOLIDAY_TYPE_WEEKEND},
	HolidayTypeFestival:      {Proto: v1.HolidayType_HOLIDAY_TYPE_FESTIVAL},
	HolidayTypeMakeupWorkday: {Proto: v1.HolidayType_HOLIDAY_TYPE_MAKEUP_WORKDAY},
})

func (e HolidayType) String() string {
	return string(e)
}
