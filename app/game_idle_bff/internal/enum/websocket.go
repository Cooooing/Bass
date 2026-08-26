package enum

// WebSocketMessageType 表示摸鱼放置 BFF websocket 消息类型。
type WebSocketMessageType string

const (
	WebSocketMessageTypeChatMessage  WebSocketMessageType = "chat.message"  // 聊天消息
	WebSocketMessageTypeCloseSession WebSocketMessageType = "close_session" // 关闭会话
	WebSocketMessageTypeError        WebSocketMessageType = "error"         // 命令失败
	WebSocketMessageTypeChatSend     WebSocketMessageType = "chat.send"     // 发送聊天
	WebSocketMessageTypeChatList     WebSocketMessageType = "chat.list"     // 查询聊天
	WebSocketMessageTypeBackpackGet  WebSocketMessageType = "backpack.get"  // 查询背包
	WebSocketMessageTypeQueueList    WebSocketMessageType = "queue.list"    // 查询行动队列
	WebSocketMessageTypeQueueAdd     WebSocketMessageType = "queue.add"     // 添加行动
	WebSocketMessageTypeQueueMove    WebSocketMessageType = "queue.move"    // 移动行动
	WebSocketMessageTypeQueueRemove  WebSocketMessageType = "queue.remove"  // 移除行动
	WebSocketMessageTypeQueueClear   WebSocketMessageType = "queue.clear"   // 清空行动队列
)
