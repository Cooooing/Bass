

# RespArticleBrief

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 文章 ID |  [optional] |
|**title** | **String** | 标题 |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | 文章类型 |  [optional] |
|**authorUser** | [**RespAccountProfile**](RespAccountProfile.md) | 作者展示资料 |  [optional] |
|**coverImageUrl** | **String** | 封面图片 URL |  [optional] |
|**publishStatus** | [**PublishStatusEnum**](#PublishStatusEnum) | 发布状态 |  [optional] |
|**visibility** | [**VisibilityEnum**](#VisibilityEnum) | 可见范围 |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) | 管理限制 |  [optional] |
|**createdBy** | **String** | 创建账号 ID |  [optional] |
|**updatedBy** | **String** | 更新账号 ID |  [optional] |
|**createdAt** | **OffsetDateTime** | 创建时间 |  [optional] |
|**updatedAt** | **OffsetDateTime** | 更新时间 |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| ARTICLE_TYPE_UNSPECIFIED | &quot;ARTICLE_TYPE_UNSPECIFIED&quot; |
| ARTICLE_TYPE_NORMAL | &quot;ARTICLE_TYPE_NORMAL&quot; |
| ARTICLE_TYPE_QA | &quot;ARTICLE_TYPE_QA&quot; |
| ARTICLE_TYPE_LOTTERY | &quot;ARTICLE_TYPE_LOTTERY&quot; |
| ARTICLE_TYPE_POLL | &quot;ARTICLE_TYPE_POLL&quot; |
| ARTICLE_TYPE_COLUMN | &quot;ARTICLE_TYPE_COLUMN&quot; |



## Enum: PublishStatusEnum

| Name | Value |
|---- | -----|
| ARTICLE_PUBLISH_STATUS_UNSPECIFIED | &quot;ARTICLE_PUBLISH_STATUS_UNSPECIFIED&quot; |
| ARTICLE_PUBLISH_STATUS_DRAFT | &quot;ARTICLE_PUBLISH_STATUS_DRAFT&quot; |
| ARTICLE_PUBLISH_STATUS_PUBLISHED | &quot;ARTICLE_PUBLISH_STATUS_PUBLISHED&quot; |
| ARTICLE_PUBLISH_STATUS_ARCHIVED | &quot;ARTICLE_PUBLISH_STATUS_ARCHIVED&quot; |



## Enum: VisibilityEnum

| Name | Value |
|---- | -----|
| ARTICLE_VISIBILITY_UNSPECIFIED | &quot;ARTICLE_VISIBILITY_UNSPECIFIED&quot; |
| ARTICLE_VISIBILITY_PUBLIC | &quot;ARTICLE_VISIBILITY_PUBLIC&quot; |
| ARTICLE_VISIBILITY_PRIVATE | &quot;ARTICLE_VISIBILITY_PRIVATE&quot; |



## Enum: RestrictionEnum

| Name | Value |
|---- | -----|
| CONTENT_RESTRICTION_UNSPECIFIED | &quot;CONTENT_RESTRICTION_UNSPECIFIED&quot; |
| CONTENT_RESTRICTION_NONE | &quot;CONTENT_RESTRICTION_NONE&quot; |
| CONTENT_RESTRICTION_HIDDEN | &quot;CONTENT_RESTRICTION_HIDDEN&quot; |
| CONTENT_RESTRICTION_LOCKED | &quot;CONTENT_RESTRICTION_LOCKED&quot; |



