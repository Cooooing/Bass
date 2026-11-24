

# ArticlePostscript

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 附言 ID |  [optional] |
|**articleId** | **String** | 文章 ID |  [optional] |
|**content** | **String** | 原始内容 |  [optional] |
|**contentRender** | **String** | 渲染内容 |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) | 治理限制 |  [optional] |
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



