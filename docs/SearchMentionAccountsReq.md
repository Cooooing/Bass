# SearchMentionAccountsReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**query** | **string** | 用户名或昵称搜索词，可为空。 | [optional] [default to undefined]
**page** | [**PageReq**](PageReq.md) | 分页参数。 | [optional] [default to undefined]

## Example

```typescript
import { SearchMentionAccountsReq } from '@bass/bbs-sdk-axios';

const instance: SearchMentionAccountsReq = {
    query,
    page,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
