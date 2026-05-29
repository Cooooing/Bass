

# ArticleSave


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** |  |  [optional] |
|**title** | **String** |  |  [optional] |
|**content** | **String** |  |  [optional] |
|**rewardContent** | **String** |  |  [optional] |
|**rewardPoints** | **Integer** |  |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) |  |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) |  |  [optional] |
|**bountyPoints** | **Integer** |  |  [optional] |
|**statement** | **String** |  |  [optional] |
|**commentable** | **Boolean** |  |  [optional] |
|**anonymous** | **Boolean** |  |  [optional] |
|**listable** | **Boolean** |  |  [optional] |
|**tags** | [**List&lt;TagSave&gt;**](TagSave.md) |  |  [optional] |



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



