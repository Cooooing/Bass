

# ListCommentRepliesReq

请求参数。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**page** | [**PageReq**](PageReq.md) | 分页参数 |  [optional] |
|**articleId** | **String** | 文章 ID |  |
|**parentId** | **String** | 顶层评论 ID |  |
|**order** | [**OrderEnum**](#OrderEnum) | 排序方式 |  [optional] |



## Enum: OrderEnum

| Name | Value |
|---- | -----|
| COMMENT_ORDER_UNSPECIFIED | &quot;COMMENT_ORDER_UNSPECIFIED&quot; |
| COMMENT_ORDER_NEWEST | &quot;COMMENT_ORDER_NEWEST&quot; |
| COMMENT_ORDER_HOTTEST | &quot;COMMENT_ORDER_HOTTEST&quot; |
| COMMENT_ORDER_OLDEST | &quot;COMMENT_ORDER_OLDEST&quot; |



