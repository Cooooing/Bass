

# Article


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** |  |  [optional] |
|**title** | **String** |  |  [optional] |
|**content** | **String** |  |  [optional] |
|**contentRender** | **String** |  |  [optional] |
|**hasPostscript** | **Boolean** |  |  [optional] |
|**hasReward** | **Boolean** |  |  [optional] |
|**rewardContent** | **String** |  |  [optional] |
|**rewardContentRender** | **String** |  |  [optional] |
|**rewardPoints** | **Integer** |  |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) |  |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) |  |  [optional] |
|**statement** | **String** |  |  [optional] |
|**commentable** | **Boolean** |  |  [optional] |
|**anonymous** | **Boolean** |  |  [optional] |
|**listable** | **Boolean** |  |  [optional] |
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
|**createdBy** | **String** |  |  [optional] |
|**updatedBy** | **String** |  |  [optional] |
|**createdAt** | **String** |  |  [optional] |
|**updatedAt** | **String** |  |  [optional] |



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



