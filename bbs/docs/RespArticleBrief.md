# RespArticleBrief

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> |  | [optional]
**title** | Option<**String**> |  | [optional]
**r#type** | Option<**Type**> |  (enum: ARTICLE_TYPE_UNSPECIFIED, ARTICLE_TYPE_NORMAL, ARTICLE_TYPE_QA, ARTICLE_TYPE_LOTTERY, ARTICLE_TYPE_POLL, ARTICLE_TYPE_COLUMN) | [optional]
**author_user** | Option<[**models::RespAccountProfile**](RespAccountProfile.md)> |  | [optional]
**cover_image_url** | Option<**String**> |  | [optional]
**publish_status** | Option<**PublishStatus**> |  (enum: ARTICLE_PUBLISH_STATUS_UNSPECIFIED, ARTICLE_PUBLISH_STATUS_DRAFT, ARTICLE_PUBLISH_STATUS_PUBLISHED, ARTICLE_PUBLISH_STATUS_ARCHIVED, ARTICLE_PUBLISH_STATUS_SCHEDULED) | [optional]
**visibility** | Option<**Visibility**> |  (enum: ARTICLE_VISIBILITY_UNSPECIFIED, ARTICLE_VISIBILITY_PUBLIC, ARTICLE_VISIBILITY_PRIVATE) | [optional]
**restriction** | Option<**Restriction**> |  (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]
**created_by** | Option<**String**> |  | [optional]
**updated_by** | Option<**String**> |  | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


