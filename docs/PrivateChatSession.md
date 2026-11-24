

# PrivateChatSession

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**peerUserId** | **String** | 对端账号 ID。 |  [optional] |
|**lastReadMessageId** | **String** | 最后已读消息 ID。 |  [optional] |
|**unreadCount** | **Integer** | 未读消息数。 |  [optional] |
|**muted** | **Boolean** | 是否免打扰。 |  [optional] |
|**pinned** | **Boolean** | 是否置顶。 |  [optional] |
|**lastMessage** | [**PrivateChatMessage**](PrivateChatMessage.md) | 最后一条消息。 |  [optional] |
|**lastMessageAt** | **OffsetDateTime** | 最后消息时间。 |  [optional] |
|**peer** | [**PrivateChatMember**](PrivateChatMember.md) | 对端展示资料。 |  [optional] |



