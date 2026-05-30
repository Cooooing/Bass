

# Article

文章。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 文章 ID。 |  [optional] |
|**title** | **String** | 标题。 |  [optional] |
|**content** | **String** | 原始内容。 |  [optional] |
|**contentRender** | **String** | 渲染后的内容。 |  [optional] |
|**hasPostscript** | **Boolean** | 是否有附言。 |  [optional] |
|**hasReward** | **Boolean** | 是否有打赏内容。 |  [optional] |
|**rewardContent** | **String** | 打赏后可见内容。 |  [optional] |
|**rewardContentRender** | **String** | 渲染后的打赏内容。 |  [optional] |
|**rewardPoints** | **Integer** | 打赏所需积分。 |  [optional] |
|**status** | [**StatusEnum**](#StatusEnum) | 文章状态。 |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | 文章类型。 |  [optional] |
|**statement** | **String** | 文章声明。 |  [optional] |
|**commentable** | **Boolean** | 是否允许评论。 |  [optional] |
|**anonymous** | **Boolean** | 是否匿名展示。 |  [optional] |
|**listable** | **Boolean** | 是否在列表中展示。 |  [optional] |
|**viewCount** | **Integer** | 浏览数量。 |  [optional] |
|**thankCount** | **Integer** | 感谢数量。 |  [optional] |
|**likeCount** | **Integer** | 点赞数量。 |  [optional] |
|**collectCount** | **Integer** | 收藏数量。 |  [optional] |
|**watchCount** | **Integer** | 关注数量。 |  [optional] |
|**replyCount** | **Integer** | 回复数量。 |  [optional] |
|**bountyPoints** | **Integer** | 悬赏积分。 |  [optional] |
|**acceptedAnswerId** | **String** | 已采纳答案评论 ID。 |  [optional] |
|**authorUser** | [**AccountProfile**](AccountProfile.md) | 作者账号展示资料。 |  [optional] |
|**lastReplyUser** | [**AccountProfile**](AccountProfile.md) | 最后回复账号展示资料。 |  [optional] |
|**lastReplyAt** | **String** | 最后回复时间。 |  [optional] |
|**coverImageUrl** | **String** | 封面图片 URL。 |  [optional] |
|**createdBy** | **String** | 创建账号 ID。 |  [optional] |
|**updatedBy** | **String** | 更新账号 ID。 |  [optional] |
|**createdAt** | **String** | 创建时间。 |  [optional] |
|**updatedAt** | **String** | 更新时间。 |  [optional] |



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



