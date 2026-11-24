# SearchMentionAccountsResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**rows** | [**Array&lt;AccountProfile&gt;**](AccountProfile.md) | 最多五个可提及账号候选。 | [optional] [default to undefined]
**page** | [**PageResp**](PageResp.md) | 分页结果。 | [optional] [default to undefined]

## Example

```typescript
import { SearchMentionAccountsResp } from '@bass/bbs-sdk-axios';

const instance: SearchMentionAccountsResp = {
    rows,
    page,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
