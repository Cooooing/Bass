

# ArticleDetail

协议消息。

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**id** | **String** | 文章 ID |  [optional] |
|**title** | **String** | 标题 |  [optional] |
|**content** | **String** | 原始内容 |  [optional] |
|**contentRender** | **String** | 渲染内容 |  [optional] |
|**hasPostscript** | **Boolean** | 是否有附言 |  [optional] |
|**hasReward** | **Boolean** | 是否有打赏区 |  [optional] |
|**rewardContent** | **String** | 打赏后可见内容 |  [optional] |
|**rewardContentRender** | **String** | 渲染打赏内容 |  [optional] |
|**rewardPoints** | **Integer** | 打赏积分 |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | 文章类型 |  [optional] |
|**statement** | **String** | 创作声明 |  [optional] |
|**commentable** | **Boolean** | 是否允许评论 |  [optional] |
|**viewCount** | **Integer** | 浏览数 |  [optional] |
|**thankCount** | **Integer** | 感谢数 |  [optional] |
|**likeCount** | **Integer** | 点赞数 |  [optional] |
|**collectCount** | **Integer** | 收藏数 |  [optional] |
|**watchCount** | **Integer** | 关注数 |  [optional] |
|**rewardCount** | **Integer** | 打赏数 |  [optional] |
|**replyCount** | **Integer** | 回复数 |  [optional] |
|**authorUser** | [**AccountProfile**](AccountProfile.md) | 作者资料 |  [optional] |
|**lastReplyUser** | [**AccountProfile**](AccountProfile.md) | 最后回复用户 |  [optional] |
|**lastReplyAt** | **OffsetDateTime** | 最后回复时间 |  [optional] |
|**coverImageUrl** | **String** | 封面图片 URL |  [optional] |
|**viewerActionState** | [**ArticleViewerActionState**](ArticleViewerActionState.md) | 当前用户互动状态 |  [optional] |
|**publishedAt** | **OffsetDateTime** | 发布时间 |  [optional] |
|**postscripts** | [**List&lt;ArticlePostscript&gt;**](ArticlePostscript.md) | 附言列表 |  [optional] |
|**publishStatus** | [**PublishStatusEnum**](#PublishStatusEnum) | 发布状态 |  [optional] |
|**visibility** | [**VisibilityEnum**](#VisibilityEnum) | 可见范围 |  [optional] |
|**restriction** | [**RestrictionEnum**](#RestrictionEnum) | 治理限制 |  [optional] |
|**editedAt** | **OffsetDateTime** | 编辑时间 |  [optional] |
|**tags** | [**List&lt;ArticleTag&gt;**](ArticleTag.md) | 文章标签 |  [optional] |
|**domains** | [**List&lt;ArticleDomain&gt;**](ArticleDomain.md) | 文章所属领域 |  [optional] |
|**city** | **String** | 发布时城市快照 |  [optional] |
|**acceptedCommentId** | **String** | 问答已采纳评论 ID |  [optional] |
|**poll** | [**ArticlePoll**](ArticlePoll.md) | 投票详情 |  [optional] |
|**lottery** | [**ArticleLottery**](ArticleLottery.md) | 抽奖详情 |  [optional] |
|**createdBy** | **String** | 创建账号 ID |  [optional] |
|**updatedBy** | **String** | 更新账号 ID |  [optional] |
|**createdAt** | **OffsetDateTime** | 创建时间 |  [optional] |
|**updatedAt** | **OffsetDateTime** | 更新时间 |  [optional] |



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



