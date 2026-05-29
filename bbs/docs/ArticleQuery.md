

# ArticleQuery


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tagId** | **String** |  |  [optional] |
|**domainId** | **String** |  |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) |  |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) |  |  [optional] |
|**order** | [**OrderEnum**](#OrderEnum) |  |  [optional] |
|**keyword** | **String** |  |  [optional] |
|**authorId** | **String** |  |  [optional] |
|**listable** | **Boolean** |  |  [optional] |



## Enum: StatusEnum

| Name | Value |
|---- | -----|
| ARTICLE_STATUS_UNSPECIFIED | &quot;ARTICLE_STATUS_UNSPECIFIED&quot; |
| ARTICLE_STATUS_NORMAL | &quot;ARTICLE_STATUS_NORMAL&quot; |
| ARTICLE_STATUS_HIDDEN | &quot;ARTICLE_STATUS_HIDDEN&quot; |
| ARTICLE_STATUS_LOCKED | &quot;ARTICLE_STATUS_LOCKED&quot; |
| ARTICLE_STATUS_DRAFTS | &quot;ARTICLE_STATUS_DRAFTS&quot; |
| ARTICLE_STATUS_DELETED | &quot;ARTICLE_STATUS_DELETED&quot; |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| ARTICLE_TYPE_UNSPECIFIED | &quot;ARTICLE_TYPE_UNSPECIFIED&quot; |
| ARTICLE_TYPE_NORMAL | &quot;ARTICLE_TYPE_NORMAL&quot; |
| ARTICLE_TYPE_QA | &quot;ARTICLE_TYPE_QA&quot; |



## Enum: OrderEnum

| Name | Value |
|---- | -----|
| ARTICLE_ORDER_UNSPECIFIED | &quot;ARTICLE_ORDER_UNSPECIFIED&quot; |
| ARTICLE_ORDER_NEWEST | &quot;ARTICLE_ORDER_NEWEST&quot; |
| ARTICLE_ORDER_HOTTEST | &quot;ARTICLE_ORDER_HOTTEST&quot; |



