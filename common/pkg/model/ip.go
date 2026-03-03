package model

type IpInfo struct {
	Ip          string `json:"ip,omitempty"`           // ip
	Country     string `json:"country,omitempty"`      // 国家
	Province    string `json:"province,omitempty"`     // 省份
	City        string `json:"city,omitempty"`         // 城市
	ISP         string `json:"isp,omitempty"`          // 运营商
	CountryCode string `json:"country_code,omitempty"` // 国家代码
}
