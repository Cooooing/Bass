

# ListCommentThreadsReq

请求参数。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**page** | [**PageReq**](PageReq.md) | 分页参数 |  [optional] |
|**articleId** | **String** | 文章 ID |  |
|**order** | [**OrderEnum**](#OrderEnum) | 顶层评论排序方式 |  [optional] |
|**replyPreviewLimit** | **Integer** | 每个楼层返回的回复预览数量 |  [optional] |



## Enum: OrderEnum

| Name | Value |
|---- | -----|
| COMMENT_ORDER_UNSPECIFIED | &quot;COMMENT_ORDER_UNSPECIFIED&quot; |
| COMMENT_ORDER_NEWEST | &quot;COMMENT_ORDER_NEWEST&quot; |
| COMMENT_ORDER_HOTTEST | &quot;COMMENT_ORDER_HOTTEST&quot; |
| COMMENT_ORDER_OLDEST | &quot;COMMENT_ORDER_OLDEST&quot; |



