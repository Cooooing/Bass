

# CommentQuery

评论查询条件。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**commentId** | **String** | 评论 ID。 |  [optional] |
|**articleId** | **String** | 文章 ID。 |  [optional] |
|**parentId** | **String** | 父评论 ID。 |  [optional] |
|**replyId** | **String** | 回复的评论 ID。 |  [optional] |
|**order** | [**OrderEnum**](#OrderEnum) | 排序方式。 |  [optional] |
|**userId** | **String** | 评论账号 ID。 |  [optional] |
|**level** | **Integer** | 评论层级。 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 评论状态。 |  [optional] |



## Enum: OrderEnum

| Name | Value |
|---- | -----|
| COMMENT_ORDER_UNSPECIFIED | &quot;COMMENT_ORDER_UNSPECIFIED&quot; |
| COMMENT_ORDER_NEWEST | &quot;COMMENT_ORDER_NEWEST&quot; |
| COMMENT_ORDER_HOTTEST | &quot;COMMENT_ORDER_HOTTEST&quot; |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| COMMENT_STATUS_UNSPECIFIED | &quot;COMMENT_STATUS_UNSPECIFIED&quot; |
| COMMENT_STATUS_NORMAL | &quot;COMMENT_STATUS_NORMAL&quot; |
| COMMENT_STATUS_HIDDEN | &quot;COMMENT_STATUS_HIDDEN&quot; |



