# GlobalChatMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> |  | [optional]
**sender_id** | Option<**String**> |  | [optional]
**content** | Option<**String**> |  | [optional]
**message_type** | Option<**MessageType**> |  (enum: IM_MESSAGE_TYPE_UNSPECIFIED, IM_MESSAGE_TYPE_TEXT, IM_MESSAGE_TYPE_RED_PACKET) | [optional]
**revoked** | Option<**bool**> |  | [optional]
**sender** | Option<[**models::GlobalChatMember**](GlobalChatMember.md)> |  | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**revoked_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


