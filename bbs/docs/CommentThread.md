# CommentThread

评论楼层项。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**root** | [**CommentListItem**](CommentListItem.md) |  | [optional] [default to undefined]
**preview_replies** | [**Array&lt;CommentListItem&gt;**](CommentListItem.md) |  | [optional] [default to undefined]
**reply_count** | **number** |  | [optional] [default to undefined]
**has_more_replies** | **boolean** |  | [optional] [default to undefined]

## Example

```typescript
import { CommentThread } from '@bass/bbs-sdk-axios';

const instance: CommentThread = {
    root,
    preview_replies,
    reply_count,
    has_more_replies,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
