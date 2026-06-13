

# CommentQuery

评论查询条件。

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
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) |  |  [optional] |
|**restrictions** | [**List&lt;RestrictionsEnum&gt;**](#List&lt;RestrictionsEnum&gt;) |  |  [optional] |



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



