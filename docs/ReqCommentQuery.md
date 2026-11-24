

# ReqCommentQuery

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**commentId** | **String** | 评论 ID |  [optional] |
|**articleId** | **String** | 文章 ID |  [optional] |
|**parentId** | **String** | 父评论 ID |  [optional] |
|**replyId** | **String** | 回复的评论 ID |  [optional] |
|**order** | [**OrderEnum**](#OrderEnum) | 排序方式 |  [optional] |
|**userId** | **String** | 评论账号 ID |  [optional] |
|**level** | **Integer** | 评论层级 |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) | 管理限制 |  [optional] |
|**restrictions** | [**List&lt;RestrictionsEnum&gt;**](#List&lt;RestrictionsEnum&gt;) | 管理限制列表 |  [optional] |



## Enum: OrderEnum

| Name | Value |
|---- | -----|
| COMMENT_ORDER_UNSPECIFIED | &quot;COMMENT_ORDER_UNSPECIFIED&quot; |
| COMMENT_ORDER_NEWEST | &quot;COMMENT_ORDER_NEWEST&quot; |
| COMMENT_ORDER_HOTTEST | &quot;COMMENT_ORDER_HOTTEST&quot; |
| COMMENT_ORDER_OLDEST | &quot;COMMENT_ORDER_OLDEST&quot; |



## Enum: RestrictionEnum

| Name | Value |
|---- | -----|
| CONTENT_RESTRICTION_UNSPECIFIED | &quot;CONTENT_RESTRICTION_UNSPECIFIED&quot; |
| CONTENT_RESTRICTION_NONE | &quot;CONTENT_RESTRICTION_NONE&quot; |
| CONTENT_RESTRICTION_HIDDEN | &quot;CONTENT_RESTRICTION_HIDDEN&quot; |
| CONTENT_RESTRICTION_LOCKED | &quot;CONTENT_RESTRICTION_LOCKED&quot; |



## Enum: List&lt;RestrictionsEnum&gt;

| Name | Value |
|---- | -----|
| CONTENT_RESTRICTION_UNSPECIFIED | &quot;CONTENT_RESTRICTION_UNSPECIFIED&quot; |
| CONTENT_RESTRICTION_NONE | &quot;CONTENT_RESTRICTION_NONE&quot; |
| CONTENT_RESTRICTION_HIDDEN | &quot;CONTENT_RESTRICTION_HIDDEN&quot; |
| CONTENT_RESTRICTION_LOCKED | &quot;CONTENT_RESTRICTION_LOCKED&quot; |



