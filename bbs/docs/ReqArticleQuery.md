# ReqArticleQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tag_id** | Option<**String**> |  | [optional]
**domain_id** | Option<**String**> |  | [optional]
**r#type** | Option<**Type**> |  (enum: ARTICLE_TYPE_UNSPECIFIED, ARTICLE_TYPE_NORMAL, ARTICLE_TYPE_QA, ARTICLE_TYPE_LOTTERY, ARTICLE_TYPE_POLL, ARTICLE_TYPE_COLUMN) | [optional]
**order** | Option<**Order**> |  (enum: ARTICLE_ORDER_UNSPECIFIED, ARTICLE_ORDER_NEWEST, ARTICLE_ORDER_HOTTEST) | [optional]
**keyword** | Option<**String**> |  | [optional]
**author_id** | Option<**String**> |  | [optional]
**publish_status** | Option<**PublishStatus**> |  (enum: ARTICLE_PUBLISH_STATUS_UNSPECIFIED, ARTICLE_PUBLISH_STATUS_DRAFT, ARTICLE_PUBLISH_STATUS_PUBLISHED, ARTICLE_PUBLISH_STATUS_ARCHIVED, ARTICLE_PUBLISH_STATUS_SCHEDULED) | [optional]
**publish_statuses** | Option<**Vec<PublishStatuses>**> |  (enum: ARTICLE_PUBLISH_STATUS_UNSPECIFIED, ARTICLE_PUBLISH_STATUS_DRAFT, ARTICLE_PUBLISH_STATUS_PUBLISHED, ARTICLE_PUBLISH_STATUS_ARCHIVED, ARTICLE_PUBLISH_STATUS_SCHEDULED) | [optional]
**visibility** | Option<**Visibility**> |  (enum: ARTICLE_VISIBILITY_UNSPECIFIED, ARTICLE_VISIBILITY_PUBLIC, ARTICLE_VISIBILITY_PRIVATE) | [optional]
**visibilities** | Option<**Vec<Visibilities>**> |  (enum: ARTICLE_VISIBILITY_UNSPECIFIED, ARTICLE_VISIBILITY_PUBLIC, ARTICLE_VISIBILITY_PRIVATE) | [optional]
**restriction** | Option<**Restriction**> |  (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]
**restrictions** | Option<**Vec<Restrictions>**> |  (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


