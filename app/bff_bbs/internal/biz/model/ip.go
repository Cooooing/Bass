package model

// IpInfo 是 IP 地址的地理信息。
type IpInfo struct { //nolint:revive // 字段名沿用 platform RPC 契约的 Ip/Isp 拼写，避免响应映射时做额外转换。
	Ip          string //nolint:revive // 与 platform RPC 契约保持一致。
	Country     string
	Province    string
	City        string
	ISP         string
	CountryCode string
}
