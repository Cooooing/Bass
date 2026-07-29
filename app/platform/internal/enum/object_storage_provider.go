package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/platform/v1/enum"
)

type ObjectStorageProvider string

const (
	ObjectStorageProviderMinio ObjectStorageProvider = "minio"
	ObjectStorageProviderQiniu ObjectStorageProvider = "qiniu"
)

var ObjectStorageProviderMap = commonenum.NewMapping[ObjectStorageProvider, v1.ObjectStorageProvider](map[ObjectStorageProvider]commonenum.Entry[ObjectStorageProvider, v1.ObjectStorageProvider]{
	ObjectStorageProviderMinio: {Proto: v1.ObjectStorageProvider_OBJECT_STORAGE_PROVIDER_MINIO},
	ObjectStorageProviderQiniu: {Proto: v1.ObjectStorageProvider_OBJECT_STORAGE_PROVIDER_QINIU},
})

func (e ObjectStorageProvider) String() string {
	return string(e)
}
