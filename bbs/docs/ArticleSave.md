# ArticleSave

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 文章 ID。 | [optional]
**title** | **String** | 标题。 | 
**content** | **String** | 原始内容。 | 
**reward_content** | Option<**String**> | 打赏后可见内容。 | [optional]
**reward_points** | Option<**i32**> | 打赏所需积分。 | [optional]
**status** | **Status** | 文章状态。 (enum: ARTICLE_STATUS_UNSPECIFIED, ARTICLE_STATUS_NORMAL, ARTICLE_STATUS_HIDDEN, ARTICLE_STATUS_LOCKED, ARTICLE_STATUS_DRAFTS, ARTICLE_STATUS_DELETED) | 
**r#type** | **Type** | 文章类型。 (enum: ARTICLE_TYPE_UNSPECIFIED, ARTICLE_TYPE_NORMAL, ARTICLE_TYPE_QA) | 
**bounty_points** | Option<**i32**> | 悬赏积分。 | [optional]
**statement** | Option<**String**> | 文章声明。 | [optional]
**commentable** | Option<**bool**> | 是否允许评论。 | [optional]
**anonymous** | Option<**bool**> | 是否匿名展示。 | [optional]
**listable** | Option<**bool**> | 是否在列表中展示。 | [optional]
**tags** | Option<[**Vec<models::TagSave>**](TagSave.md)> | 绑定标签。 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


