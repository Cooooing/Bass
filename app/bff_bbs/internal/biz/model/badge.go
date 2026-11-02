package model

// BadgePreview 是 BFF 向客户端返回的 SVG 二进制内容。
type BadgePreview struct {
	Data        []byte
	ContentType string
}
