# ListCommentRepliesReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageReq**](PageReq.md) | 分页参数 | [optional] [default to undefined]
**article_id** | **string** | 文章 ID | [default to undefined]
**parent_id** | **string** | 顶层评论 ID | [default to undefined]
**order** | **string** | 排序方式 | [optional] [default to undefined]

## Example

```typescript
import { ListCommentRepliesReq } from '@bass/bbs-sdk-axios';

const instance: ListCommentRepliesReq = {
    page,
    article_id,
    parent_id,
    order,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
