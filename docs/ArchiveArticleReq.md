# ArchiveArticleReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 必填，文章 ID | [default to undefined]
**reason** | **string** | 可选，归档原因 | [optional] [default to undefined]

## Example

```typescript
import { ArchiveArticleReq } from '@bass/bbs-sdk-axios';

const instance: ArchiveArticleReq = {
    article_id,
    reason,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
