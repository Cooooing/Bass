package enum

// WebSocketMessageType 表示摸鱼放置 BFF websocket 消息类型。
// 消息类型使用小写点分层命名：客户端发送命令使用“领域.对象.动作”，服务端发送事件使用“领域.对象.结果”。
type WebSocketMessageType string

const (
	// 服务端发送给客户端的事件。
	WebSocketMessageTypeChatMessageReceived WebSocketMessageType = "chat.message.received" // 收到聊天消息
	WebSocketMessageTypeChatMessageListed   WebSocketMessageType = "chat.message.listed"   // 聊天消息列表
	WebSocketMessageTypeActionCompleted     WebSocketMessageType = "action.completed"      // 行动完成
	WebSocketMessageTypeActionListed        WebSocketMessageType = "action.listed"         // 行动列表
	WebSocketMessageTypeBackpackUpdated     WebSocketMessageType = "backpack.updated"      // 背包变化
	WebSocketMessageTypeBackpackItemsListed WebSocketMessageType = "backpack.items.listed" // 背包物品列表
	WebSocketMessageTypeSessionClose        WebSocketMessageType = "session.close"         // 关闭会话
	WebSocketMessageTypeCommandFailed       WebSocketMessageType = "command.failed"        // 命令失败

	// 客户端发送给服务端的命令。
	WebSocketMessageTypeChatMessageSend WebSocketMessageType = "chat.message.send" // 发送聊天消息
	WebSocketMessageTypeChatMessageList WebSocketMessageType = "chat.message.list" // 查询聊天消息
	WebSocketMessageTypeBackpackGet     WebSocketMessageType = "backpack.get"      // 查询背包
	WebSocketMessageTypeActionList      WebSocketMessageType = "action.list"       // 查询行动列表
	WebSocketMessageTypeActionAdd       WebSocketMessageType = "action.add"        // 添加行动
	WebSocketMessageTypeActionMove      WebSocketMessageType = "action.move"       // 移动行动
	WebSocketMessageTypeActionRemove    WebSocketMessageType = "action.remove"     // 移除行动
	WebSocketMessageTypeActionClear     WebSocketMessageType = "action.clear"      // 清空行动
)
