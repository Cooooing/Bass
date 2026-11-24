# TransferAccountPointsReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_name** | **string** | 收款账号名，必填 | [default to undefined]
**amount** | **string** | 转账积分数量，必填 | [default to undefined]
**remark** | **string** | 转账备注 | [optional] [default to undefined]
**idempotency_key** | **string** | 本次转账的技术幂等键，必填 | [default to undefined]

## Example

```typescript
import { TransferAccountPointsReq } from '@bass/bbs-sdk-axios';

const instance: TransferAccountPointsReq = {
    target_name,
    amount,
    remark,
    idempotency_key,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
