

# RespCommentListItem


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** |  |  [optional] |
|**articleId** | **String** |  |  [optional] |
|**content** | **String** |  |  [optional] |
|**contentRender** | **String** |  |  [optional] |
|**level** | **Integer** |  |  [optional] |
|**parentId** | **String** |  |  [optional] |
|**replyId** | **String** |  |  [optional] |
|**replyCount** | **Integer** |  |  [optional] |
|**likeCount** | **Integer** |  |  [optional] |
|**thankCount** | **Integer** |  |  [optional] |
|**user** | [**RespAccountProfile**](RespAccountProfile.md) |  |  [optional] |
|**replyUser** | [**RespAccountProfile**](RespAccountProfile.md) |  |  [optional] |
|**article** | [**RespArticleBrief**](RespArticleBrief.md) |  |  [optional] |
|**viewerActionState** | [**RespCommentViewerActionState**](RespCommentViewerActionState.md) |  |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) |  |  [optional] |
|**deletedAt** | **OffsetDateTime** |  |  [optional] |
|**createdBy** | **String** |  |  [optional] |
|**updatedBy** | **String** |  |  [optional] |
|**createdAt** | **OffsetDateTime** |  |  [optional] |
|**updatedAt** | **OffsetDateTime** |  |  [optional] |



## Enum: RestrictionEnum

| Name | Value |
|---- | -----|
| CONTENT_RESTRICTION_UNSPECIFIED | &quot;CONTENT_RESTRICTION_UNSPECIFIED&quot; |
| CONTENT_RESTRICTION_NONE | &quot;CONTENT_RESTRICTION_NONE&quot; |
| CONTENT_RESTRICTION_HIDDEN | &quot;CONTENT_RESTRICTION_HIDDEN&quot; |
| CONTENT_RESTRICTION_LOCKED | &quot;CONTENT_RESTRICTION_LOCKED&quot; |



