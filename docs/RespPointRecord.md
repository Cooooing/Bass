# RespPointRecord

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**created_at** | **string** | 记录创建时间 | [optional] [default to undefined]
**type_label** | **string** | 流水类型展示文案 | [optional] [default to undefined]
**direction** | **string** | 收入或支出方向 | [optional] [default to undefined]
**amount** | **string** | 绝对变动积分 | [optional] [default to undefined]
**balance_after** | **string** | 变动后的积分余额 | [optional] [default to undefined]

## Example

```typescript
import { RespPointRecord } from '@bass/bbs-sdk-axios';

const instance: RespPointRecord = {
    created_at,
    type_label,
    direction,
    amount,
    balance_after,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
