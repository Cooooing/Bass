# SetPinnedArticleCommentReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 文章 ID | [default to undefined]
**comment_id** | **string** | 置顶的一级评论，未传表示取消 | [optional] [default to undefined]

## Example

```typescript
import { SetPinnedArticleCommentReq } from '@bass/bbs-sdk-axios';

const instance: SetPinnedArticleCommentReq = {
    article_id,
    comment_id,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
