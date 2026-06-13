

# ArticleQuery

文章查询条件。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tagId** | **String** |  |  [optional] |
|**domainId** | **String** |  |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) |  |  [optional] |
|**order** | [**OrderEnum**](#OrderEnum) |  |  [optional] |
|**keyword** | **String** |  |  [optional] |
|**authorId** | **String** |  |  [optional] |
|**publishStatus** | [**PublishStatusEnum**](#PublishStatusEnum) |  |  [optional] |
|**publishStatuses** | [**List&lt;PublishStatusesEnum&gt;**](#List&lt;PublishStatusesEnum&gt;) |  |  [optional] |
|**visibility** | [**VisibilityEnum**](#VisibilityEnum) |  |  [optional] |
|**visibilities** | [**List&lt;VisibilitiesEnum&gt;**](#List&lt;VisibilitiesEnum&gt;) |  |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) |  |  [optional] |
|**restrictions** | [**List&lt;RestrictionsEnum&gt;**](#List&lt;RestrictionsEnum&gt;) |  |  [optional] |



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



## Enum: PublishStatusEnum

| Name | Value |
|---- | -----|
| ARTICLE_PUBLISH_STATUS_UNSPECIFIED | &quot;ARTICLE_PUBLISH_STATUS_UNSPECIFIED&quot; |
| ARTICLE_PUBLISH_STATUS_DRAFT | &quot;ARTICLE_PUBLISH_STATUS_DRAFT&quot; |
| ARTICLE_PUBLISH_STATUS_PUBLISHED | &quot;ARTICLE_PUBLISH_STATUS_PUBLISHED&quot; |
| ARTICLE_PUBLISH_STATUS_ARCHIVED | &quot;ARTICLE_PUBLISH_STATUS_ARCHIVED&quot; |



## Enum: List&lt;PublishStatusesEnum&gt;

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



## Enum: List&lt;VisibilitiesEnum&gt;

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



## Enum: List&lt;RestrictionsEnum&gt;

| Name | Value |
|---- | -----|
| CONTENT_RESTRICTION_UNSPECIFIED | &quot;CONTENT_RESTRICTION_UNSPECIFIED&quot; |
| CONTENT_RESTRICTION_NONE | &quot;CONTENT_RESTRICTION_NONE&quot; |
| CONTENT_RESTRICTION_HIDDEN | &quot;CONTENT_RESTRICTION_HIDDEN&quot; |
| CONTENT_RESTRICTION_LOCKED | &quot;CONTENT_RESTRICTION_LOCKED&quot; |



