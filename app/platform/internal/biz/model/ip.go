package model

// IpInfo 是 IP 地理解析结果的领域模型。
type IpInfo struct { //nolint:revive // 字段名沿用 platform RPC 契约的 Ip/Isp 拼写，避免响应映射时做额外转换。
	Ip          string `json:"ip,omitempty"`           //nolint:revive // 与 platform RPC 契约保持一致。
	Country     string `json:"country,omitempty"`      // 国家
	Province    string `json:"province,omitempty"`     // 省份
	City        string `json:"city,omitempty"`         // 城市
	ISP         string `json:"isp,omitempty"`          // 运营商
	CountryCode string `json:"country_code,omitempty"` // 国家代码
}
