package model

// EmojiGroup 是 BFF 已解析公开图片地址的表情分组视图。
type EmojiGroup struct {
	ID    int64
	Name  string
	Items []*EmojiItem
}

// EmojiItem 是 BFF 返回给编辑器的可插入图片表情。
type EmojiItem struct {
	ID       int64
	AssetID  int64
	Name     string
	ImageURL string
}
