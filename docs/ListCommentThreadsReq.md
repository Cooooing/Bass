# ListCommentThreadsReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageReq**](PageReq.md) | 分页参数 | [optional] [default to undefined]
**article_id** | **string** | 文章 ID | [default to undefined]
**order** | **string** | 顶层评论排序方式 | [optional] [default to undefined]
**reply_preview_limit** | **number** | 每个楼层返回的回复预览数量 | [optional] [default to undefined]

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
