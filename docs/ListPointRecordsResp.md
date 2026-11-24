# ListPointRecordsResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**rows** | [**Array&lt;RespPointRecord&gt;**](RespPointRecord.md) | 积分流水列表 | [optional] [default to undefined]
**page** | [**PageResp**](PageResp.md) | 分页结果 | [optional] [default to undefined]

## Example

```typescript
import { ListPointRecordsResp } from '@bass/bbs-sdk-axios';

const instance: ListPointRecordsResp = {
    rows,
    page,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
