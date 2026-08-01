# ArticleListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> |  | [optional]
**title** | Option<**String**> |  | [optional]
**content** | Option<**String**> |  | [optional]
**content_render** | Option<**String**> |  | [optional]
**has_postscript** | Option<**bool**> |  | [optional]
**has_reward** | Option<**bool**> |  | [optional]
**r#type** | Option<**Type**> |  (enum: ARTICLE_TYPE_UNSPECIFIED, ARTICLE_TYPE_NORMAL, ARTICLE_TYPE_QA, ARTICLE_TYPE_LOTTERY, ARTICLE_TYPE_POLL, ARTICLE_TYPE_COLUMN) | [optional]
**statement** | Option<**String**> |  | [optional]
**commentable** | Option<**bool**> |  | [optional]
**view_count** | Option<**i32**> |  | [optional]
**thank_count** | Option<**i32**> |  | [optional]
**like_count** | Option<**i32**> |  | [optional]
**collect_count** | Option<**i32**> |  | [optional]
**reward_count** | Option<**i32**> |  | [optional]
**reply_count** | Option<**i32**> |  | [optional]
**author_user** | Option<[**models::AccountProfile**](AccountProfile.md)> |  | [optional]
**last_reply_user** | Option<[**models::AccountProfile**](AccountProfile.md)> |  | [optional]
**last_reply_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**cover_image_url** | Option<**String**> |  | [optional]
**viewer_action_state** | Option<[**models::ArticleViewerActionState**](ArticleViewerActionState.md)> |  | [optional]
**published_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**publish_status** | Option<**PublishStatus**> |  (enum: ARTICLE_PUBLISH_STATUS_UNSPECIFIED, ARTICLE_PUBLISH_STATUS_DRAFT, ARTICLE_PUBLISH_STATUS_PUBLISHED, ARTICLE_PUBLISH_STATUS_ARCHIVED, ARTICLE_PUBLISH_STATUS_SCHEDULED) | [optional]
**visibility** | Option<**Visibility**> |  (enum: ARTICLE_VISIBILITY_UNSPECIFIED, ARTICLE_VISIBILITY_PUBLIC, ARTICLE_VISIBILITY_PRIVATE) | [optional]
**restriction** | Option<**Restriction**> |  (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]
**edited_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**created_by** | Option<**String**> |  | [optional]
**updated_by** | Option<**String**> |  | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


