# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 通知 ID。 | [optional]
**event_id** | Option<**String**> | 来源事件 ID。 | [optional]
**receiver_id** | Option<**String**> | 接收账号 ID。 | [optional]
**event_type** | Option<**EventType**> | 来源事件类型。 (enum: EVENT_TYPE_UNSPECIFIED, EVENT_TYPE_USER_FOLLOW, EVENT_TYPE_USER_UNFOLLOW, EVENT_TYPE_USER_REGISTER, EVENT_TYPE_USER_LOGIN, EVENT_TYPE_USER_LOGOUT, EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE, EVENT_TYPE_USER_PHONE_VERIFICATION_CODE, EVENT_TYPE_USER_BLOCK, EVENT_TYPE_USER_UNBLOCK, EVENT_TYPE_USER_TOTP_ENABLE, EVENT_TYPE_USER_TOTP_DISABLE, EVENT_TYPE_ARTICLE_PUBLISHED, EVENT_TYPE_ARTICLE_LIKED, EVENT_TYPE_ARTICLE_THANKED, EVENT_TYPE_ARTICLE_COLLECTED, EVENT_TYPE_ARTICLE_WATCHED, EVENT_TYPE_COMMENT_PUBLISHED, EVENT_TYPE_COMMENT_LIKED) | [optional]
**title** | Option<**String**> | 通知标题。 | [optional]
**content** | Option<**String**> | 通知内容。 | [optional]
**read_at** | Option<**String**> | 已读时间。 | [optional]
**created_at** | Option<**String**> | 创建时间。 | [optional]
**updated_at** | Option<**String**> | 更新时间。 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


