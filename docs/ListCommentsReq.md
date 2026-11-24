# ListCommentsReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageReq**](PageReq.md) | 分页参数 | [optional] [default to undefined]
**query** | [**ReqCommentQuery**](ReqCommentQuery.md) | 查询条件 | [optional] [default to undefined]

## Example

```typescript
import { ListCommentsReq } from '@bass/bbs-sdk-axios';

const instance: ListCommentsReq = {
    page,
    query,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
