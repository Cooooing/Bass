# AcceptArticleAnswerReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 问答文章 ID | [default to undefined]
**comment_id** | **string** | 采纳的一级回答 ID | [default to undefined]

## Example

```typescript
import { AcceptArticleAnswerReq } from '@bass/bbs-sdk-axios';

const instance: AcceptArticleAnswerReq = {
    article_id,
    comment_id,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
