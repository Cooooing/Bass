# ListCommentThreadsResp


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageResp**](PageResp.md) |  | [optional] [default to undefined]
**rows** | [**Array&lt;RespCommentThread&gt;**](RespCommentThread.md) |  | [optional] [default to undefined]
**accepted_answer** | [**RespPinnedComment**](RespPinnedComment.md) |  | [optional] [default to undefined]
**pinned_comment** | [**RespPinnedComment**](RespPinnedComment.md) |  | [optional] [default to undefined]

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
