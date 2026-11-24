# VoteArticlePollReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**article_id** | **string** | 必填，文章 ID | [default to undefined]
**option_ids** | **Array&lt;string&gt;** | 必填，所选选项 ID | [default to undefined]

## Example

```typescript
import { VoteArticlePollReq } from '@bass/bbs-sdk-axios';

const instance: VoteArticlePollReq = {
    article_id,
    option_ids,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
