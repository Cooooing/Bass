

# ListCommentThreadsReq


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**page** | [**PageReq**](PageReq.md) |  |  [optional] |
|**articleId** | **String** |  |  |
|**order** | [**OrderEnum**](#OrderEnum) |  |  [optional] |
|**replyPreviewLimit** | **Integer** |  |  [optional] |



## Enum: OrderEnum

| Name | Value |
|---- | -----|
| COMMENT_ORDER_UNSPECIFIED | &quot;COMMENT_ORDER_UNSPECIFIED&quot; |
| COMMENT_ORDER_NEWEST | &quot;COMMENT_ORDER_NEWEST&quot; |
| COMMENT_ORDER_HOTTEST | &quot;COMMENT_ORDER_HOTTEST&quot; |
| COMMENT_ORDER_OLDEST | &quot;COMMENT_ORDER_OLDEST&quot; |



