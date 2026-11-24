# RespArticleBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> | 文章 ID | [optional]
**title** | Option<**String**> | 标题 | [optional]
**r#type** | Option<**Type**> | 文章类型 (enum: ARTICLE_TYPE_UNSPECIFIED, ARTICLE_TYPE_NORMAL, ARTICLE_TYPE_QA, ARTICLE_TYPE_LOTTERY, ARTICLE_TYPE_POLL, ARTICLE_TYPE_COLUMN) | [optional]
**author_user** | Option<[**models::RespAccountProfile**](RespAccountProfile.md)> | 作者展示资料 | [optional]
**cover_image_url** | Option<**String**> | 封面图片 URL | [optional]
**publish_status** | Option<**PublishStatus**> | 发布状态 (enum: ARTICLE_PUBLISH_STATUS_UNSPECIFIED, ARTICLE_PUBLISH_STATUS_DRAFT, ARTICLE_PUBLISH_STATUS_PUBLISHED, ARTICLE_PUBLISH_STATUS_ARCHIVED) | [optional]
**visibility** | Option<**Visibility**> | 可见范围 (enum: ARTICLE_VISIBILITY_UNSPECIFIED, ARTICLE_VISIBILITY_PUBLIC, ARTICLE_VISIBILITY_PRIVATE) | [optional]
**restriction** | Option<**Restriction**> | 管理限制 (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]
**created_by** | Option<**String**> | 创建账号 ID | [optional]
**updated_by** | Option<**String**> | 更新账号 ID | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 创建时间 | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> | 更新时间 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


