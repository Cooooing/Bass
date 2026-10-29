package model

// GlobalChatMessageCount 是一个账号在时间窗口内的有效聊天室文本消息数量。
type GlobalChatMessageCount struct {
	UserID int64
	Count  int32
}
