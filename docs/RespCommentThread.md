# RespCommentThread

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**root** | [**RespCommentListItem**](RespCommentListItem.md) | 顶层评论 | [optional] [default to undefined]
**preview_replies** | [**Array&lt;RespCommentListItem&gt;**](RespCommentListItem.md) | 回复预览列表 | [optional] [default to undefined]
**reply_count** | **number** | 回复数量 | [optional] [default to undefined]
**has_more_replies** | **boolean** | 是否还有更多回复 | [optional] [default to undefined]

## Example

```typescript
import { RespCommentThread } from '@bass/bbs-sdk-axios';

const instance: RespCommentThread = {
    root,
    preview_replies,
    reply_count,
    has_more_replies,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
