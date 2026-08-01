# RespNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> |  | [optional]
**event_id** | Option<**String**> |  | [optional]
**receiver_id** | Option<**String**> |  | [optional]
**event_type** | Option<**EventType**> |  (enum: EVENT_TYPE_UNSPECIFIED, EVENT_TYPE_USER_FOLLOW, EVENT_TYPE_USER_UNFOLLOW, EVENT_TYPE_USER_REGISTER, EVENT_TYPE_USER_LOGIN, EVENT_TYPE_USER_LOGOUT, EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE, EVENT_TYPE_USER_PHONE_VERIFICATION_CODE, EVENT_TYPE_USER_BLOCK, EVENT_TYPE_USER_UNBLOCK, EVENT_TYPE_USER_TOTP_ENABLE, EVENT_TYPE_USER_TOTP_DISABLE, EVENT_TYPE_USER_ACCOUNT_CANCELLED, EVENT_TYPE_USER_ACCOUNT_BANNED, EVENT_TYPE_USER_ACCOUNT_UNBANNED, EVENT_TYPE_ARTICLE_PUBLISHED, EVENT_TYPE_ARTICLE_LIKED, EVENT_TYPE_ARTICLE_THANKED, EVENT_TYPE_ARTICLE_COLLECTED, EVENT_TYPE_ARTICLE_WATCHED, EVENT_TYPE_ARTICLE_ACCEPTED_ANSWER, EVENT_TYPE_ARTICLE_POSTSCRIPT_ADDED, EVENT_TYPE_ARTICLE_STATUS_UPDATED, EVENT_TYPE_COMMENT_PUBLISHED, EVENT_TYPE_COMMENT_LIKED, EVENT_TYPE_COMMENT_THANKED, EVENT_TYPE_COMMENT_STATUS_UPDATED) | [optional]
**title** | Option<**String**> |  | [optional]
**content** | Option<**String**> |  | [optional]
**read_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


