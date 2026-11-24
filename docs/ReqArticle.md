

# ReqArticle

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**title** | **String** | 必填，标题 |  |
|**content** | **String** | 必填，正文 |  |
|**rewardContent** | **String** | 可选，打赏区内容 |  [optional] |
|**rewardPoints** | **Integer** | 可选，打赏积分 |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | 必填，文章类型 |  |
|**statement** | **String** | 可选，创作声明 |  [optional] |
|**commentable** | **Boolean** | 可选，是否允许评论 |  [optional] |
|**poll** | [**ArticlePollDraft**](ArticlePollDraft.md) | 可选，投票文章的草稿配置 |  [optional] |
|**lottery** | [**ArticleLotteryDraft**](ArticleLotteryDraft.md) | 可选，抽奖文章的草稿配置 |  [optional] |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| ARTICLE_TYPE_UNSPECIFIED | &quot;ARTICLE_TYPE_UNSPECIFIED&quot; |
| ARTICLE_TYPE_NORMAL | &quot;ARTICLE_TYPE_NORMAL&quot; |
| ARTICLE_TYPE_QA | &quot;ARTICLE_TYPE_QA&quot; |
| ARTICLE_TYPE_LOTTERY | &quot;ARTICLE_TYPE_LOTTERY&quot; |
| ARTICLE_TYPE_POLL | &quot;ARTICLE_TYPE_POLL&quot; |
| ARTICLE_TYPE_COLUMN | &quot;ARTICLE_TYPE_COLUMN&quot; |



