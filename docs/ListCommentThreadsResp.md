# ListCommentThreadsResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageResp**](PageResp.md) | 分页结果 | [optional] [default to undefined]
**rows** | [**Array&lt;RespCommentThread&gt;**](RespCommentThread.md) | 评论楼层列表 | [optional] [default to undefined]
**accepted_answer** | [**RespPinnedComment**](RespPinnedComment.md) | 问答已采纳回答；不可见时不返回 | [optional] [default to undefined]
**pinned_comment** | [**RespPinnedComment**](RespPinnedComment.md) | 作者手动置顶评论；不可见时不返回 | [optional] [default to undefined]

## Example

```typescript
import { ListCommentThreadsResp } from '@bass/bbs-sdk-axios';

const instance: ListCommentThreadsResp = {
    page,
    rows,
    accepted_answer,
    pinned_comment,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
