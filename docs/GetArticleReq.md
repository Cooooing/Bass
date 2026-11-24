

# GetArticleReq

请求参数。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**articleId** | **String** | 必填，文章 ID |  |
|**publishStatus** | [**PublishStatusEnum**](#PublishStatusEnum) | 可选，预期发布状态；未传按已发布查询 |  [optional] |



## Enum: PublishStatusEnum

| Name | Value |
|---- | -----|
| ARTICLE_PUBLISH_STATUS_UNSPECIFIED | &quot;ARTICLE_PUBLISH_STATUS_UNSPECIFIED&quot; |
| ARTICLE_PUBLISH_STATUS_DRAFT | &quot;ARTICLE_PUBLISH_STATUS_DRAFT&quot; |
| ARTICLE_PUBLISH_STATUS_PUBLISHED | &quot;ARTICLE_PUBLISH_STATUS_PUBLISHED&quot; |
| ARTICLE_PUBLISH_STATUS_ARCHIVED | &quot;ARTICLE_PUBLISH_STATUS_ARCHIVED&quot; |



