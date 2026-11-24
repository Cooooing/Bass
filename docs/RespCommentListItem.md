# RespCommentListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 评论 ID | [optional]
**article_id** | Option<**String**> | 文章 ID | [optional]
**content** | Option<**String**> | 原始内容 | [optional]
**content_render** | Option<**String**> | 渲染后的内容 | [optional]
**level** | Option<**i32**> | 评论层级 | [optional]
**parent_id** | Option<**String**> | 父评论 ID | [optional]
**reply_id** | Option<**String**> | 回复的评论 ID | [optional]
**reply_count** | Option<**i32**> | 回复数量 | [optional]
**like_count** | Option<**i32**> | 点赞数量 | [optional]
**thank_count** | Option<**i32**> | 感谢数量 | [optional]
**user** | Option<[**models::RespAccountProfile**](RespAccountProfile.md)> | 评论账号展示资料 | [optional]
**reply_user** | Option<[**models::RespAccountProfile**](RespAccountProfile.md)> | 被回复账号展示资料 | [optional]
**article** | Option<[**models::RespArticleBrief**](RespArticleBrief.md)> | 所属文章摘要 | [optional]
**viewer_action_state** | Option<[**models::RespCommentViewerActionState**](RespCommentViewerActionState.md)> | 当前查看账号的评论行为状态 | [optional]
**restriction** | Option<**Restriction**> | 管理限制 (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]
**deleted_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 逻辑删除时间 | [optional]
**created_by** | Option<**String**> | 创建账号 ID | [optional]
**updated_by** | Option<**String**> | 更新账号 ID | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 创建时间 | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 更新时间 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


