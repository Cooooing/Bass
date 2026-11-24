package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/platform/enum/v1"
)

type HolidayType string

const (
	HolidayTypeWorkday       HolidayType = "workday"        // 法定工作日
	HolidayTypeWeekend       HolidayType = "weekend"        // 常规周末休息日
	HolidayTypeFestival      HolidayType = "festival"       // 法定节假日
	HolidayTypeMakeupWorkday HolidayType = "makeup_workday" // 调休补班日
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
