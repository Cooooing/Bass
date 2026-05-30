

# Comment

评论。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 评论 ID。 |  [optional] |
|**articleId** | **String** | 文章 ID。 |  [optional] |
|**content** | **String** | 原始内容。 |  [optional] |
|**contentRender** | **String** | 渲染后的内容。 |  [optional] |
|**level** | **Integer** | 评论层级。 |  [optional] |
|**parentId** | **String** | 父评论 ID。 |  [optional] |
|**replyId** | **String** | 回复的评论 ID。 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 评论状态。 |  [optional] |
|**replyCount** | **Integer** | 回复数量。 |  [optional] |
|**likeCount** | **Integer** | 点赞数量。 |  [optional] |
|**thankCount** | **Integer** | 感谢数量。 |  [optional] |
|**user** | [**AccountProfile**](AccountProfile.md) | 评论账号展示资料。 |  [optional] |
|**replyUser** | [**AccountProfile**](AccountProfile.md) | 被回复账号展示资料。 |  [optional] |
|**article** | [**Article**](Article.md) | 所属文章。 |  [optional] |
|**createdBy** | **String** | 创建账号 ID。 |  [optional] |
|**updatedBy** | **String** | 更新账号 ID。 |  [optional] |
|**createdAt** | **String** | 创建时间。 |  [optional] |
|**updatedAt** | **String** | 更新时间。 |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| COMMENT_STATUS_UNSPECIFIED | &quot;COMMENT_STATUS_UNSPECIFIED&quot; |
| COMMENT_STATUS_NORMAL | &quot;COMMENT_STATUS_NORMAL&quot; |
| COMMENT_STATUS_HIDDEN | &quot;COMMENT_STATUS_HIDDEN&quot; |



