

# RespCommentDetail

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 评论 ID |  [optional] |
|**articleId** | **String** | 文章 ID |  [optional] |
|**content** | **String** | 原始内容 |  [optional] |
|**contentRender** | **String** | 渲染后的内容 |  [optional] |
|**level** | **Integer** | 评论层级 |  [optional] |
|**parentId** | **String** | 父评论 ID |  [optional] |
|**replyId** | **String** | 回复的评论 ID |  [optional] |
|**replyCount** | **Integer** | 回复数量 |  [optional] |
|**likeCount** | **Integer** | 点赞数量 |  [optional] |
|**thankCount** | **Integer** | 感谢数量 |  [optional] |
|**user** | [**RespAccountProfile**](RespAccountProfile.md) | 评论账号展示资料 |  [optional] |
|**replyUser** | [**RespAccountProfile**](RespAccountProfile.md) | 被回复账号展示资料 |  [optional] |
|**article** | [**RespArticleBrief**](RespArticleBrief.md) | 所属文章摘要 |  [optional] |
|**viewerActionState** | [**RespCommentViewerActionState**](RespCommentViewerActionState.md) | 当前查看账号的评论行为状态 |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) | 管理限制 |  [optional] |
|**deletedAt** | **OffsetDateTime** | 逻辑删除时间 |  [optional] |
|**createdBy** | **String** | 创建账号 ID |  [optional] |
|**updatedBy** | **String** | 更新账号 ID |  [optional] |
|**createdAt** | **OffsetDateTime** | 创建时间 |  [optional] |
|**updatedAt** | **OffsetDateTime** | 更新时间 |  [optional] |



## Enum: RestrictionEnum

| Name | Value |
|---- | -----|
| CONTENT_RESTRICTION_UNSPECIFIED | &quot;CONTENT_RESTRICTION_UNSPECIFIED&quot; |
| CONTENT_RESTRICTION_NONE | &quot;CONTENT_RESTRICTION_NONE&quot; |
| CONTENT_RESTRICTION_HIDDEN | &quot;CONTENT_RESTRICTION_HIDDEN&quot; |
| CONTENT_RESTRICTION_LOCKED | &quot;CONTENT_RESTRICTION_LOCKED&quot; |



