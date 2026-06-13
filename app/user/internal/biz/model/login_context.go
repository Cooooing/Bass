package model

// LoginContext 表示由入口层显式传入的登录客户端上下文。
type LoginContext struct {
	IP          string
	Country     string
	CountryCode string
	Province    string
	City        string
	ISP         string
	UserAgent   string
	DeviceID    string
	Platform    string
	RequestID   string
}
