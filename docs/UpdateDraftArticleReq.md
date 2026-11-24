# UpdateDraftArticleReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 必填，文章 ID | [default to undefined]
**article** | [**ReqArticle**](ReqArticle.md) | 必填，草稿内容 | [default to undefined]

## Example

```typescript
import { UpdateDraftArticleReq } from '@bass/bbs-sdk-axios';

const instance: UpdateDraftArticleReq = {
    article_id,
    article,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
