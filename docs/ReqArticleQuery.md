

# ReqArticleQuery

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tagId** | **String** | 可选，标签 ID |  [optional] |
|**domainId** | **String** | 可选，领域 ID |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | 可选，文章类型 |  [optional] |
|**order** | [**OrderEnum**](#OrderEnum) | 可选，排序方式 |  [optional] |
|**authorId** | **String** | 可选，作者账号 ID |  [optional] |
|**publishStatus** | [**PublishStatusEnum**](#PublishStatusEnum) | 可选，发布状态 |  [optional] |
|**publishStatuses** | [**List&lt;PublishStatusesEnum&gt;**](#List&lt;PublishStatusesEnum&gt;) | 可选，发布状态列表 |  [optional] |
|**visibility** | [**VisibilityEnum**](#VisibilityEnum) | 可选，可见范围 |  [optional] |
|**visibilities** | [**List&lt;VisibilitiesEnum&gt;**](#List&lt;VisibilitiesEnum&gt;) | 可选，可见范围列表 |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) | 可选，治理限制 |  [optional] |
|**restrictions** | [**List&lt;RestrictionsEnum&gt;**](#List&lt;RestrictionsEnum&gt;) | 可选，治理限制列表 |  [optional] |
|**scheduled** | **Boolean** | 可选，仅本人草稿：是否已设置发布时间 |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| ARTICLE_TYPE_UNSPECIFIED | &quot;ARTICLE_TYPE_UNSPECIFIED&quot; |
| ARTICLE_TYPE_NORMAL | &quot;ARTICLE_TYPE_NORMAL&quot; |
| ARTICLE_TYPE_QA | &quot;ARTICLE_TYPE_QA&quot; |
| ARTICLE_TYPE_LOTTERY | &quot;ARTICLE_TYPE_LOTTERY&quot; |
| ARTICLE_TYPE_POLL | &quot;ARTICLE_TYPE_POLL&quot; |
| ARTICLE_TYPE_COLUMN | &quot;ARTICLE_TYPE_COLUMN&quot; |



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



