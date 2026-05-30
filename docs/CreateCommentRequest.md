# CreateCommentRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 文章 ID。 | [default to undefined]
**content** | **string** | 评论内容。 | [default to undefined]
**reply_id** | **string** | 回复的评论 ID。 | [optional] [default to undefined]

## Example

```typescript
import { CreateCommentRequest } from '@bass/bbs-sdk-axios';

const instance: CreateCommentRequest = {
    article_id,
    content,
    reply_id,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
