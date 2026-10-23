

# GlobalChatMessage


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** |  |  [optional] |
|**senderId** | **String** |  |  [optional] |
|**content** | **String** |  |  [optional] |
|**messageType** | [**MessageTypeEnum**](#MessageTypeEnum) |  |  [optional] |
|**revoked** | **Boolean** |  |  [optional] |
|**sender** | [**GlobalChatMember**](GlobalChatMember.md) |  |  [optional] |
|**createdAt** | **OffsetDateTime** |  |  [optional] |
|**revokedAt** | **OffsetDateTime** |  |  [optional] |



## Enum: MessageTypeEnum

| Name | Value |
|---- | -----|
| IM_MESSAGE_TYPE_UNSPECIFIED | &quot;IM_MESSAGE_TYPE_UNSPECIFIED&quot; |
| IM_MESSAGE_TYPE_TEXT | &quot;IM_MESSAGE_TYPE_TEXT&quot; |
| IM_MESSAGE_TYPE_RED_PACKET | &quot;IM_MESSAGE_TYPE_RED_PACKET&quot; |



