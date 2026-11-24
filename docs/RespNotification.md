# RespNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 通知 ID | [optional]
**event_id** | Option<**String**> | 来源事件 ID | [optional]
**receiver_id** | Option<**String**> | 接收账号 ID | [optional]
**notification_type** | Option<**NotificationType**> | 通知类型 (enum: NOTIFICATION_TYPE_UNSPECIFIED, NOTIFICATION_TYPE_USER_FOLLOW, NOTIFICATION_TYPE_USER_REGISTER, NOTIFICATION_TYPE_USER_EMAIL_VERIFICATION_CODE, NOTIFICATION_TYPE_USER_PHONE_VERIFICATION_CODE, NOTIFICATION_TYPE_ARTICLE_PUBLISHED, NOTIFICATION_TYPE_ARTICLE_LIKED, NOTIFICATION_TYPE_ARTICLE_THANKED, NOTIFICATION_TYPE_ARTICLE_COLLECTED, NOTIFICATION_TYPE_ARTICLE_WATCHED, NOTIFICATION_TYPE_ARTICLE_ACCEPTED_ANSWER, NOTIFICATION_TYPE_ARTICLE_LOTTERY_WON, NOTIFICATION_TYPE_ARTICLE_WATCHED_COMMENT_PUBLISHED, NOTIFICATION_TYPE_ARTICLE_WATCHED_POSTSCRIPT_ADDED, NOTIFICATION_TYPE_ARTICLE_WATCHED_ANSWER_ACCEPTED, NOTIFICATION_TYPE_ARTICLE_WATCHED_POLL_ENDED, NOTIFICATION_TYPE_ARTICLE_WATCHED_LOTTERY_DRAWN, NOTIFICATION_TYPE_ARTICLE_MENTIONED, NOTIFICATION_TYPE_POSTSCRIPT_MENTIONED, NOTIFICATION_TYPE_ARTICLE_REWARDED, NOTIFICATION_TYPE_COMMENT_PUBLISHED, NOTIFICATION_TYPE_COMMENT_LIKED, NOTIFICATION_TYPE_COMMENT_THANKED, NOTIFICATION_TYPE_COMMENT_MENTIONED, NOTIFICATION_TYPE_IM_GLOBAL_CHAT_MENTIONED) | [optional]
**title** | Option<**String**> | 通知标题 | [optional]
**content** | Option<**String**> | 通知内容 | [optional]
**read_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 已读时间 | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 创建时间 | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 更新时间 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


