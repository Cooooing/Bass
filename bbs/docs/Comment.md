# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 评论 ID。 | [optional]
**article_id** | Option<**String**> | 文章 ID。 | [optional]
**content** | Option<**String**> | 原始内容。 | [optional]
**content_render** | Option<**String**> | 渲染后的内容。 | [optional]
**level** | Option<**i32**> | 评论层级。 | [optional]
**parent_id** | Option<**String**> | 父评论 ID。 | [optional]
**reply_id** | Option<**String**> | 回复的评论 ID。 | [optional]
**status** | Option<**Status**> | 评论状态。 (enum: COMMENT_STATUS_UNSPECIFIED, COMMENT_STATUS_NORMAL, COMMENT_STATUS_HIDDEN) | [optional]
**reply_count** | Option<**i32**> | 回复数量。 | [optional]
**like_count** | Option<**i32**> | 点赞数量。 | [optional]
**thank_count** | Option<**i32**> | 感谢数量。 | [optional]
**user** | Option<[**models::AccountProfile**](AccountProfile.md)> | 评论账号展示资料。 | [optional]
**reply_user** | Option<[**models::AccountProfile**](AccountProfile.md)> | 被回复账号展示资料。 | [optional]
**article** | Option<[**models::Article**](Article.md)> | 所属文章。 | [optional]
**created_by** | Option<**String**> | 创建账号 ID。 | [optional]
**updated_by** | Option<**String**> | 更新账号 ID。 | [optional]
**created_at** | Option<**String**> | 创建时间。 | [optional]
**updated_at** | Option<**String**> | 更新时间。 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


