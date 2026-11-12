package model

// IpInfo 是登录审计所需的 IP 地理信息。
type IpInfo struct {
	Ip          string
	Country     string
	Province    string
	City        string
	ISP         string
	CountryCode string
}
