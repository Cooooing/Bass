# RespTotp

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **string** | 账号 ID | [optional] [default to undefined]
**enable** | **boolean** | 是否启用 TOTP | [optional] [default to undefined]
**enable_time** | **string** | 启用时间 | [optional] [default to undefined]

## Example

```typescript
import { RespTotp } from '@bass/bbs-sdk-axios';

const instance: RespTotp = {
    user_id,
    enable,
    enable_time,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
