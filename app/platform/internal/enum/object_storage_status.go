package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/platform/v1/enum"
)

type ObjectStorageStatus string

const (
	ObjectStorageStatusAvailable ObjectStorageStatus = "available"
	ObjectStorageStatusBlocked   ObjectStorageStatus = "blocked"
	ObjectStorageStatusDeleted   ObjectStorageStatus = "deleted"
)

var ObjectStorageStatusMap = commonenum.NewMapping[ObjectStorageStatus, v1.ObjectStorageStatus](map[ObjectStorageStatus]commonenum.Entry[ObjectStorageStatus, v1.ObjectStorageStatus]{
	ObjectStorageStatusAvailable: {Proto: v1.ObjectStorageStatus_OBJECT_STORAGE_STATUS_AVAILABLE},
	ObjectStorageStatusBlocked:   {Proto: v1.ObjectStorageStatus_OBJECT_STORAGE_STATUS_BLOCKED},
	ObjectStorageStatusDeleted:   {Proto: v1.ObjectStorageStatus_OBJECT_STORAGE_STATUS_DELETED},
})

func (e ObjectStorageStatus) String() string {
	return string(e)
}
