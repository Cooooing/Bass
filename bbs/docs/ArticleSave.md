

# ArticleSave

文章保存内容。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 文章 ID。 |  [optional] |
|**title** | **String** | 标题。 |  |
|**content** | **String** | 原始内容。 |  |
|**rewardContent** | **String** | 打赏后可见内容。 |  [optional] |
|**rewardPoints** | **Integer** | 打赏所需积分。 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 文章状态。 |  |
|**type** | [**TypeEnum**](#TypeEnum) | 文章类型。 |  |
|**bountyPoints** | **Integer** | 悬赏积分。 |  [optional] |
|**statement** | **String** | 文章声明。 |  [optional] |
|**commentable** | **Boolean** | 是否允许评论。 |  [optional] |
|**anonymous** | **Boolean** | 是否匿名展示。 |  [optional] |
|**listable** | **Boolean** | 是否在列表中展示。 |  [optional] |
|**tags** | [**List&lt;TagSave&gt;**](TagSave.md) | 绑定标签。 |  [optional] |



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



