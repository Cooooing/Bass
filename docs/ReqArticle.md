# ReqArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**title** | **String** | 必填，标题 | 
**content** | **String** | 必填，正文 | 
**reward_content** | Option<**String**> | 可选，打赏区内容 | [optional]
**reward_points** | Option<**i32**> | 可选，打赏积分 | [optional]
**r#type** | **Type** | 必填，文章类型 (enum: ARTICLE_TYPE_UNSPECIFIED, ARTICLE_TYPE_NORMAL, ARTICLE_TYPE_QA, ARTICLE_TYPE_LOTTERY, ARTICLE_TYPE_POLL, ARTICLE_TYPE_COLUMN) | 
**statement** | Option<**String**> | 可选，创作声明 | [optional]
**commentable** | Option<**bool**> | 可选，是否允许评论 | [optional]
**poll** | Option<[**models::ArticlePollDraft**](ArticlePollDraft.md)> | 可选，投票文章的草稿配置 | [optional]
**lottery** | Option<[**models::ArticleLotteryDraft**](ArticleLotteryDraft.md)> | 可选，抽奖文章的草稿配置 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


