# ListCommentThreadsResp


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageResp**](PageResp.md) |  | [optional] [default to undefined]
**rows** | [**Array&lt;RespCommentThread&gt;**](RespCommentThread.md) |  | [optional] [default to undefined]
**top_comment** | [**RespTopComment**](RespTopComment.md) |  | [optional] [default to undefined]

## Example

```typescript
import { ListCommentThreadsResp } from '@bass/bbs-sdk-axios';

const instance: ListCommentThreadsResp = {
    page,
    rows,
    top_comment,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
