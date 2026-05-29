

# CommentQuery


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**commentId** | **String** |  |  [optional] |
|**articleId** | **String** |  |  [optional] |
|**parentId** | **String** |  |  [optional] |
|**replyId** | **String** |  |  [optional] |
|**order** | [**OrderEnum**](#OrderEnum) |  |  [optional] |
|**userId** | **String** |  |  [optional] |
|**level** | **Integer** |  |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) |  |  [optional] |



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



