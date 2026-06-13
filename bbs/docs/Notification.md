

# Notification

通知记录。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 通知 ID。 |  [optional] |
|**eventId** | **String** | 来源事件 ID。 |  [optional] |
|**receiverId** | **String** | 接收账号 ID。 |  [optional] |
|**eventType** | [**EventTypeEnum**](#EventTypeEnum) | 来源事件类型。 |  [optional] |
|**title** | **String** | 通知标题。 |  [optional] |
|**content** | **String** | 通知内容。 |  [optional] |
|**readAt** | **String** | 已读时间。 |  [optional] |
|**createdAt** | **String** | 创建时间。 |  [optional] |
|**updatedAt** | **String** | 更新时间。 |  [optional] |



## Enum: EventTypeEnum

| Name | Value |
|---- | -----|
| EVENT_TYPE_UNSPECIFIED | &quot;EVENT_TYPE_UNSPECIFIED&quot; |
| EVENT_TYPE_USER_FOLLOW | &quot;EVENT_TYPE_USER_FOLLOW&quot; |
| EVENT_TYPE_USER_UNFOLLOW | &quot;EVENT_TYPE_USER_UNFOLLOW&quot; |
| EVENT_TYPE_USER_REGISTER | &quot;EVENT_TYPE_USER_REGISTER&quot; |
| EVENT_TYPE_USER_LOGIN | &quot;EVENT_TYPE_USER_LOGIN&quot; |
| EVENT_TYPE_USER_LOGOUT | &quot;EVENT_TYPE_USER_LOGOUT&quot; |
| EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE | &quot;EVENT_TYPE_USER_EMAIL_VERIFICATION_CODE&quot; |
| EVENT_TYPE_USER_PHONE_VERIFICATION_CODE | &quot;EVENT_TYPE_USER_PHONE_VERIFICATION_CODE&quot; |
| EVENT_TYPE_USER_BLOCK | &quot;EVENT_TYPE_USER_BLOCK&quot; |
| EVENT_TYPE_USER_UNBLOCK | &quot;EVENT_TYPE_USER_UNBLOCK&quot; |
| EVENT_TYPE_USER_TOTP_ENABLE | &quot;EVENT_TYPE_USER_TOTP_ENABLE&quot; |
| EVENT_TYPE_USER_TOTP_DISABLE | &quot;EVENT_TYPE_USER_TOTP_DISABLE&quot; |
| EVENT_TYPE_ARTICLE_PUBLISHED | &quot;EVENT_TYPE_ARTICLE_PUBLISHED&quot; |
| EVENT_TYPE_ARTICLE_LIKED | &quot;EVENT_TYPE_ARTICLE_LIKED&quot; |
| EVENT_TYPE_ARTICLE_THANKED | &quot;EVENT_TYPE_ARTICLE_THANKED&quot; |
| EVENT_TYPE_ARTICLE_COLLECTED | &quot;EVENT_TYPE_ARTICLE_COLLECTED&quot; |
| EVENT_TYPE_ARTICLE_WATCHED | &quot;EVENT_TYPE_ARTICLE_WATCHED&quot; |
| EVENT_TYPE_ARTICLE_ACCEPTED_ANSWER | &quot;EVENT_TYPE_ARTICLE_ACCEPTED_ANSWER&quot; |
| EVENT_TYPE_ARTICLE_POSTSCRIPT_ADDED | &quot;EVENT_TYPE_ARTICLE_POSTSCRIPT_ADDED&quot; |
| EVENT_TYPE_ARTICLE_STATUS_UPDATED | &quot;EVENT_TYPE_ARTICLE_STATUS_UPDATED&quot; |
| EVENT_TYPE_COMMENT_PUBLISHED | &quot;EVENT_TYPE_COMMENT_PUBLISHED&quot; |
| EVENT_TYPE_COMMENT_LIKED | &quot;EVENT_TYPE_COMMENT_LIKED&quot; |
| EVENT_TYPE_COMMENT_THANKED | &quot;EVENT_TYPE_COMMENT_THANKED&quot; |
| EVENT_TYPE_COMMENT_STATUS_UPDATED | &quot;EVENT_TYPE_COMMENT_STATUS_UPDATED&quot; |



