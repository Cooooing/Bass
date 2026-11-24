# ListCommentThreadsResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | Option<[**models::PageResp**](PageResp.md)> | 分页结果 | [optional]
**rows** | Option<[**Vec<models::RespCommentThread>**](RespCommentThread.md)> | 评论楼层列表 | [optional]
**accepted_answer** | Option<[**models::RespPinnedComment**](RespPinnedComment.md)> | 问答已采纳回答；不可见时不返回 | [optional]
**pinned_comment** | Option<[**models::RespPinnedComment**](RespPinnedComment.md)> | 作者手动置顶评论；不可见时不返回 | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


