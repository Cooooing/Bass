

# ArticleQuery

文章查询条件。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tagId** | **String** | 标签 ID。 |  [optional] |
|**domainId** | **String** | 板块 ID。 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 文章状态。 |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | 文章类型。 |  [optional] |
|**order** | [**OrderEnum**](#OrderEnum) | 排序方式。 |  [optional] |
|**keyword** | **String** | 搜索关键词。 |  [optional] |
|**authorId** | **String** | 作者账号 ID。 |  [optional] |
|**listable** | **Boolean** | 是否在列表中展示。 |  [optional] |



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



