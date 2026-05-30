# Article

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 文章 ID。 | [optional]
**title** | Option<**String**> | 标题。 | [optional]
**content** | Option<**String**> | 原始内容。 | [optional]
**content_render** | Option<**String**> | 渲染后的内容。 | [optional]
**has_postscript** | Option<**bool**> | 是否有附言。 | [optional]
**has_reward** | Option<**bool**> | 是否有打赏内容。 | [optional]
**reward_content** | Option<**String**> | 打赏后可见内容。 | [optional]
**reward_content_render** | Option<**String**> | 渲染后的打赏内容。 | [optional]
**reward_points** | Option<**i32**> | 打赏所需积分。 | [optional]
**status** | Option<**Status**> | 文章状态。 (enum: ARTICLE_STATUS_UNSPECIFIED, ARTICLE_STATUS_NORMAL, ARTICLE_STATUS_HIDDEN, ARTICLE_STATUS_LOCKED, ARTICLE_STATUS_DRAFTS, ARTICLE_STATUS_DELETED) | [optional]
**r#type** | Option<**Type**> | 文章类型。 (enum: ARTICLE_TYPE_UNSPECIFIED, ARTICLE_TYPE_NORMAL, ARTICLE_TYPE_QA) | [optional]
**statement** | Option<**String**> | 文章声明。 | [optional]
**commentable** | Option<**bool**> | 是否允许评论。 | [optional]
**anonymous** | Option<**bool**> | 是否匿名展示。 | [optional]
**listable** | Option<**bool**> | 是否在列表中展示。 | [optional]
**view_count** | Option<**i32**> | 浏览数量。 | [optional]
**thank_count** | Option<**i32**> | 感谢数量。 | [optional]
**like_count** | Option<**i32**> | 点赞数量。 | [optional]
**collect_count** | Option<**i32**> | 收藏数量。 | [optional]
**watch_count** | Option<**i32**> | 关注数量。 | [optional]
**reply_count** | Option<**i32**> | 回复数量。 | [optional]
**bounty_points** | Option<**i32**> | 悬赏积分。 | [optional]
**accepted_answer_id** | Option<**String**> | 已采纳答案评论 ID。 | [optional]
**author_user** | Option<[**models::AccountProfile**](AccountProfile.md)> | 作者账号展示资料。 | [optional]
**last_reply_user** | Option<[**models::AccountProfile**](AccountProfile.md)> | 最后回复账号展示资料。 | [optional]
**last_reply_at** | Option<**String**> | 最后回复时间。 | [optional]
**cover_image_url** | Option<**String**> | 封面图片 URL。 | [optional]
**created_by** | Option<**String**> | 创建账号 ID。 | [optional]
**updated_by** | Option<**String**> | 更新账号 ID。 | [optional]
**created_at** | Option<**String**> | 创建时间。 | [optional]
**updated_at** | Option<**String**> | 更新时间。 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


