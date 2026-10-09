package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type DeviceType string

const (
	DeviceTypeUnknown DeviceType = "unknown" // unknown
	DeviceTypeDesktop DeviceType = "desktop" // desktop
	DeviceTypeMobile  DeviceType = "mobile"  // mobile
	DeviceTypeTablet  DeviceType = "tablet"  // tablet
	DeviceTypeBot     DeviceType = "bot"     // bot
)

var DeviceTypeMap = commonenum.NewMapping[DeviceType, v1.DeviceType](map[DeviceType]commonenum.Entry[DeviceType, v1.DeviceType]{
	DeviceTypeUnknown: {Proto: v1.DeviceType_DEVICE_TYPE_UNKNOWN},
	DeviceTypeDesktop: {Proto: v1.DeviceType_DEVICE_TYPE_DESKTOP},
	DeviceTypeMobile:  {Proto: v1.DeviceType_DEVICE_TYPE_MOBILE},
	DeviceTypeTablet:  {Proto: v1.DeviceType_DEVICE_TYPE_TABLET},
	DeviceTypeBot:     {Proto: v1.DeviceType_DEVICE_TYPE_BOT},
})

func (e DeviceType) String() string {
	return string(e)
}
