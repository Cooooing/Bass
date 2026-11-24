package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/enum/v1"
)

type DeviceType string

const (
	DeviceTypeUnknown DeviceType = "unknown" // 未识别设备
	DeviceTypeDesktop DeviceType = "desktop" // 桌面设备
	DeviceTypeMobile  DeviceType = "mobile"  // 移动设备
	DeviceTypeTablet  DeviceType = "tablet"  // 平板设备
	DeviceTypeBot     DeviceType = "bot"     // 自动化客户端
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
