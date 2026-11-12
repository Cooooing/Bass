package model

// IpInfo 是 IP 地址的地理信息。
type IpInfo struct {
	Ip          string
	Country     string
	Province    string
	City        string
	ISP         string
	CountryCode string
}
