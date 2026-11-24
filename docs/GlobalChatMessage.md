# GlobalChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 消息 ID。 | [optional]
**sender_id** | Option<**String**> | 发送账号 ID。 | [optional]
**content** | Option<**String**> | 原始 Markdown 内容。 | [optional]
**message_type** | Option<**MessageType**> | 消息类型。 (enum: IM_MESSAGE_TYPE_UNSPECIFIED, IM_MESSAGE_TYPE_TEXT, IM_MESSAGE_TYPE_RED_PACKET) | [optional]
**revoked** | Option<**bool**> | 是否已撤回。 | [optional]
**sender** | Option<[**models::GlobalChatMember**](GlobalChatMember.md)> | 发送者展示资料。 | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 创建时间。 | [optional]
**revoked_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 撤回时间。 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


