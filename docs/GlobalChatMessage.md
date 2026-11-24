

# GlobalChatMessage

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 消息 ID。 |  [optional] |
|**senderId** | **String** | 发送账号 ID。 |  [optional] |
|**content** | **String** | 原始 Markdown 内容。 |  [optional] |
|**messageType** | [**MessageTypeEnum**](#MessageTypeEnum) | 消息类型。 |  [optional] |
|**revoked** | **Boolean** | 是否已撤回。 |  [optional] |
|**sender** | [**GlobalChatMember**](GlobalChatMember.md) | 发送者展示资料。 |  [optional] |
|**createdAt** | **OffsetDateTime** | 创建时间。 |  [optional] |
|**revokedAt** | **OffsetDateTime** | 撤回时间。 |  [optional] |



## Enum: MessageTypeEnum

| Name | Value |
|---- | -----|
| IM_MESSAGE_TYPE_UNSPECIFIED | &quot;IM_MESSAGE_TYPE_UNSPECIFIED&quot; |
| IM_MESSAGE_TYPE_TEXT | &quot;IM_MESSAGE_TYPE_TEXT&quot; |
| IM_MESSAGE_TYPE_RED_PACKET | &quot;IM_MESSAGE_TYPE_RED_PACKET&quot; |



