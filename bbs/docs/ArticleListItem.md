

# ArticleListItem

文章列表项。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** |  |  [optional] |
|**title** | **String** |  |  [optional] |
|**content** | **String** |  |  [optional] |
|**contentRender** | **String** |  |  [optional] |
|**hasPostscript** | **Boolean** |  |  [optional] |
|**hasReward** | **Boolean** |  |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) |  |  [optional] |
|**statement** | **String** |  |  [optional] |
|**commentable** | **Boolean** |  |  [optional] |
|**anonymous** | **Boolean** |  |  [optional] |
|**viewCount** | **Integer** |  |  [optional] |
|**thankCount** | **Integer** |  |  [optional] |
|**likeCount** | **Integer** |  |  [optional] |
|**collectCount** | **Integer** |  |  [optional] |
|**watchCount** | **Integer** |  |  [optional] |
|**replyCount** | **Integer** |  |  [optional] |
|**bountyPoints** | **Integer** |  |  [optional] |
|**acceptedAnswerId** | **String** |  |  [optional] |
|**authorUser** | [**AccountProfile**](AccountProfile.md) |  |  [optional] |
|**lastReplyUser** | [**AccountProfile**](AccountProfile.md) |  |  [optional] |
|**lastReplyAt** | **String** |  |  [optional] |
|**coverImageUrl** | **String** |  |  [optional] |
|**viewerActionState** | [**ArticleViewerActionState**](ArticleViewerActionState.md) |  |  [optional] |
|**publishedAt** | **String** |  |  [optional] |
|**publishStatus** | [**PublishStatusEnum**](#PublishStatusEnum) |  |  [optional] |
|**visibility** | [**VisibilityEnum**](#VisibilityEnum) |  |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) |  |  [optional] |
|**editedAt** | **String** |  |  [optional] |
|**createdBy** | **String** |  |  [optional] |
|**updatedBy** | **String** |  |  [optional] |
|**createdAt** | **String** |  |  [optional] |
|**updatedAt** | **String** |  |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| ARTICLE_TYPE_UNSPECIFIED | &quot;ARTICLE_TYPE_UNSPECIFIED&quot; |
| ARTICLE_TYPE_NORMAL | &quot;ARTICLE_TYPE_NORMAL&quot; |
| ARTICLE_TYPE_QA | &quot;ARTICLE_TYPE_QA&quot; |



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



