# PrivateChatSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**peer_user_id** | Option<**String**> | 对端账号 ID。 | [optional]
**last_read_message_id** | Option<**String**> | 最后已读消息 ID。 | [optional]
**unread_count** | Option<**i32**> | 未读消息数。 | [optional]
**muted** | Option<**bool**> | 是否免打扰。 | [optional]
**pinned** | Option<**bool**> | 是否置顶。 | [optional]
**last_message** | Option<[**models::PrivateChatMessage**](PrivateChatMessage.md)> | 最后一条消息。 | [optional]
**last_message_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 最后消息时间。 | [optional]
**peer** | Option<[**models::PrivateChatMember**](PrivateChatMember.md)> | 对端展示资料。 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


