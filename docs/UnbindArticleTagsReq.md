# UnbindArticleTagsReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 文章 ID | [default to undefined]
**tag_ids** | **Array&lt;string&gt;** | 标签 ID 列表 | [default to undefined]

## Example

```typescript
import { UnbindArticleTagsReq } from '@bass/bbs-sdk-axios';

const instance: UnbindArticleTagsReq = {
    article_id,
    tag_ids,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
