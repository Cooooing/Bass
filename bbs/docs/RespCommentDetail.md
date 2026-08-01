# RespCommentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | Option<**String**> |  | [optional]
**article_id** | Option<**String**> |  | [optional]
**content** | Option<**String**> |  | [optional]
**content_render** | Option<**String**> |  | [optional]
**level** | Option<**i32**> |  | [optional]
**parent_id** | Option<**String**> |  | [optional]
**reply_id** | Option<**String**> |  | [optional]
**reply_count** | Option<**i32**> |  | [optional]
**like_count** | Option<**i32**> |  | [optional]
**thank_count** | Option<**i32**> |  | [optional]
**user** | Option<[**models::RespAccountProfile**](RespAccountProfile.md)> |  | [optional]
**reply_user** | Option<[**models::RespAccountProfile**](RespAccountProfile.md)> |  | [optional]
**article** | Option<[**models::RespArticleBrief**](RespArticleBrief.md)> |  | [optional]
**viewer_action_state** | Option<[**models::RespCommentViewerActionState**](RespCommentViewerActionState.md)> |  | [optional]
**restriction** | Option<**Restriction**> |  (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]
**deleted_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**created_by** | Option<**String**> |  | [optional]
**updated_by** | Option<**String**> |  | [optional]
**created_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]
**updated_at** | Option<**chrono::DateTime<chrono::FixedOffset>**> |  | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


