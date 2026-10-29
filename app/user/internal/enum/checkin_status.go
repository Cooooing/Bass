package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

// CheckinStatus 表示自动签到积分事务的状态。
type CheckinStatus string

const (
	CheckinStatusPending   CheckinStatus = "pending"   // pending
	CheckinStatusCompleted CheckinStatus = "completed" // completed
	CheckinStatusCanceled  CheckinStatus = "canceled"  // canceled
)

var CheckinStatusMap = commonenum.NewMapping[CheckinStatus, v1.CheckinStatus](map[CheckinStatus]commonenum.Entry[CheckinStatus, v1.CheckinStatus]{
	CheckinStatusPending:   {Proto: v1.CheckinStatus_CHECKIN_STATUS_PENDING},
	CheckinStatusCompleted: {Proto: v1.CheckinStatus_CHECKIN_STATUS_COMPLETED},
	CheckinStatusCanceled:  {Proto: v1.CheckinStatus_CHECKIN_STATUS_CANCELED},
})

func (e CheckinStatus) String() string {
	return string(e)
}
