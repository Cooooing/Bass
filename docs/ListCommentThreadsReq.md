# ListCommentThreadsReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | Option<[**models::PageReq**](PageReq.md)> | 分页参数 | [optional]
**article_id** | **String** | 文章 ID | 
**order** | Option<**Order**> | 顶层评论排序方式 (enum: COMMENT_ORDER_UNSPECIFIED, COMMENT_ORDER_NEWEST, COMMENT_ORDER_HOTTEST, COMMENT_ORDER_OLDEST) | [optional]
**reply_preview_limit** | Option<**i32**> | 每个楼层返回的回复预览数量 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


