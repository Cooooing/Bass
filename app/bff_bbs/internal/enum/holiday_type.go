package enum

import (
	commonenum "common/pkg/enum"
	platformv1enum "common/proto/gen/bff_bbs/platform/enum/v1"
)

type HolidayType string

const (
	HolidayTypeWorkday       HolidayType = "workday"        // 法定工作日
	HolidayTypeWeekend       HolidayType = "weekend"        // 常规周末休息日
	HolidayTypeFestival      HolidayType = "festival"       // 法定节假日
	HolidayTypeMakeupWorkday HolidayType = "makeup_workday" // 调休补班日
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
