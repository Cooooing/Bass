# ListCommentThreadsReq


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageReq**](PageReq.md) |  | [optional] [default to undefined]
**article_id** | **string** |  | [default to undefined]
**order** | **string** |  | [optional] [default to undefined]
**reply_preview_limit** | **number** |  | [optional] [default to undefined]

## Example

```typescript
import { ListCommentThreadsReq } from '@bass/bbs-sdk-axios';

const instance: ListCommentThreadsReq = {
    page,
    article_id,
    order,
    reply_preview_limit,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
