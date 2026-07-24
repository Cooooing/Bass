package model

import "user/internal/enum"

// LoginContext 表示由入口层传入并由 user 补充的登录客户端上下文。
type LoginContext struct {
	IP             string
	Country        string
	CountryCode    string
	Province       string
	City           string
	ISP            string
	UserAgent      string
	ClientType     enum.ClientType
	DeviceType     enum.DeviceType
	OSName         string
	OSVersion      string
	BrowserName    string
	BrowserVersion string
	AppName        string
	AppVersion     string
}
