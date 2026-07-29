package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type ClientType string

const (
	ClientTypeUnknown  ClientType = "unknown"
	ClientTypeWeb      ClientType = "web"
	ClientTypeAdminWeb ClientType = "admin_web"
	ClientTypeMobile   ClientType = "mobile"
)

var ClientTypeMap = commonenum.NewMapping[ClientType, v1.ClientType](map[ClientType]commonenum.Entry[ClientType, v1.ClientType]{
	ClientTypeUnknown:  {Proto: v1.ClientType_CLIENT_TYPE_UNKNOWN},
	ClientTypeWeb:      {Proto: v1.ClientType_CLIENT_TYPE_WEB},
	ClientTypeAdminWeb: {Proto: v1.ClientType_CLIENT_TYPE_ADMIN_WEB},
	ClientTypeMobile:   {Proto: v1.ClientType_CLIENT_TYPE_MOBILE},
})

func (e ClientType) String() string {
	return string(e)
}
