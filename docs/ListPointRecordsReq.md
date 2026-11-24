# ListPointRecordsReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | 账号名，必填 | [default to undefined]
**page** | [**PageReq**](PageReq.md) | 分页参数 | [optional] [default to undefined]

## Example

```typescript
import { ListPointRecordsReq } from '@bass/bbs-sdk-axios';

const instance: ListPointRecordsReq = {
    name,
    page,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
