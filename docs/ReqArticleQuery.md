# ReqArticleQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tag_id** | Option<**String**> | 可选，标签 ID | [optional]
**domain_id** | Option<**String**> | 可选，领域 ID | [optional]
**r#type** | Option<**Type**> | 可选，文章类型 (enum: ARTICLE_TYPE_UNSPECIFIED, ARTICLE_TYPE_NORMAL, ARTICLE_TYPE_QA, ARTICLE_TYPE_LOTTERY, ARTICLE_TYPE_POLL, ARTICLE_TYPE_COLUMN) | [optional]
**order** | Option<**Order**> | 可选，排序方式 (enum: ARTICLE_ORDER_UNSPECIFIED, ARTICLE_ORDER_NEWEST, ARTICLE_ORDER_HOTTEST) | [optional]
**author_id** | Option<**String**> | 可选，作者账号 ID | [optional]
**publish_status** | Option<**PublishStatus**> | 可选，发布状态 (enum: ARTICLE_PUBLISH_STATUS_UNSPECIFIED, ARTICLE_PUBLISH_STATUS_DRAFT, ARTICLE_PUBLISH_STATUS_PUBLISHED, ARTICLE_PUBLISH_STATUS_ARCHIVED) | [optional]
**publish_statuses** | Option<**Vec<PublishStatuses>**> | 可选，发布状态列表 (enum: ARTICLE_PUBLISH_STATUS_UNSPECIFIED, ARTICLE_PUBLISH_STATUS_DRAFT, ARTICLE_PUBLISH_STATUS_PUBLISHED, ARTICLE_PUBLISH_STATUS_ARCHIVED) | [optional]
**visibility** | Option<**Visibility**> | 可选，可见范围 (enum: ARTICLE_VISIBILITY_UNSPECIFIED, ARTICLE_VISIBILITY_PUBLIC, ARTICLE_VISIBILITY_PRIVATE) | [optional]
**visibilities** | Option<**Vec<Visibilities>**> | 可选，可见范围列表 (enum: ARTICLE_VISIBILITY_UNSPECIFIED, ARTICLE_VISIBILITY_PUBLIC, ARTICLE_VISIBILITY_PRIVATE) | [optional]
**restriction** | Option<**Restriction**> | 可选，治理限制 (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]
**restrictions** | Option<**Vec<Restrictions>**> | 可选，治理限制列表 (enum: CONTENT_RESTRICTION_UNSPECIFIED, CONTENT_RESTRICTION_NONE, CONTENT_RESTRICTION_HIDDEN, CONTENT_RESTRICTION_LOCKED) | [optional]
**scheduled** | Option<**bool**> | 可选，仅本人草稿：是否已设置发布时间 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


