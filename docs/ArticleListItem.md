# ArticleListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 文章 ID | [optional]
**title** | Option<**String**> | 标题 | [optional]
**content** | Option<**String**> | 摘要原文 | [optional]
**content_render** | Option<**String**> | 摘要渲染内容 | [optional]
**has_postscript** | Option<**bool**> | 是否有附言 | [optional]
**has_reward** | Option<**bool**> | 是否有打赏区 | [optional]
**r#type** | Option<**Type**> | 文章类型 (enum: ARTICLE_TYPE_UNSPECIFIED, ARTICLE_TYPE_NORMAL, ARTICLE_TYPE_QA, ARTICLE_TYPE_LOTTERY, ARTICLE_TYPE_POLL, ARTICLE_TYPE_COLUMN) | [optional]
**statement** | Option<**String**> | 创作声明 | [optional]
**commentable** | Option<**bool**> | 是否允许评论 | [optional]
**view_count** | Option<**i32**> | 浏览数 | [optional]
**thank_count** | Option<**i32**> | 感谢数 | [optional]
**like_count** | Option<**i32**> | 点赞数 | [optional]
**collect_count** | Option<**i32**> | 收藏数 | [optional]
**watch_count** | Option<**i32**> | 关注数 | [optional]
**reward_count** | Option<**i32**> | 打赏数 | [optional]
**reply_count** | Option<**i32**> | 回复数 | [optional]
**author_user** | Option<[**models::AccountProfile**](AccountProfile.md)> | 作者资料 | [optional]
**last_reply_user** | Option<[**models::AccountProfile**](AccountProfile.md)> | 最后回复用户 | [optional]
**last_reply_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 最后回复时间 | [optional]
**cover_image_url** | Option<**String**> | 封面图片 URL | [optional]
**viewer_action_state** | Option<[**models::ArticleViewerActionState**](ArticleViewerActionState.md)> | 当前用户互动状态 | [optional]
**published_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 发布时间 | [optional]
**publish_status** | Option<**PublishStatus**> | 发布状态 (enum: ARTICLE_PUBLISH_STATUS_UNSPECIFIED, ARTICLE_PUBLISH_STATUS_DRAFT, ARTICLE_PUBLISH_STATUS_PUBLISHED, ARTICLE_PUBLISH_STATUS_ARCHIVED) | [optional]
**visibility** | Option<**Visibility**> | 可见范围 (enum: ARTICLE_VISIBILITY_UNSPECIFIED, ARTICLE_VISIBILITY_PUBLIC, ARTICLE_VISIBILITY_PRIVATE) | [optional]
**restriction** | Option<**Restriction**> | 治理限制 (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]
**edited_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 编辑时间 | [optional]
**tags** | Option<[**Vec<models::ArticleTag>**](ArticleTag.md)> | 文章标签 | [optional]
**domains** | Option<[**Vec<models::ArticleDomain>**](ArticleDomain.md)> | 文章所属领域 | [optional]
**created_by** | Option<**String**> | 创建账号 ID | [optional]
**updated_by** | Option<**String**> | 更新账号 ID | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 创建时间 | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 更新时间 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


