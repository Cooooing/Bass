

# RespArticleBrief


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** |  |  [optional] |
|**title** | **String** |  |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) |  |  [optional] |
|**authorUser** | [**RespAccountProfile**](RespAccountProfile.md) |  |  [optional] |
|**coverImageUrl** | **String** |  |  [optional] |
|**publishStatus** | [**PublishStatusEnum**](#PublishStatusEnum) |  |  [optional] |
|**visibility** | [**VisibilityEnum**](#VisibilityEnum) |  |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) |  |  [optional] |
|**createdBy** | **String** |  |  [optional] |
|**updatedBy** | **String** |  |  [optional] |
|**createdAt** | **OffsetDateTime** |  |  [optional] |
|**updatedAt** | **OffsetDateTime** |  |  [optional] |



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
| ARTICLE_PUBLISH_STATUS_SCHEDULED | &quot;ARTICLE_PUBLISH_STATUS_SCHEDULED&quot; |



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



