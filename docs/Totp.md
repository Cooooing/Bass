# Totp

TOTP 认证状态。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **string** | 账号 ID。 | [optional] [default to undefined]
**enable** | **boolean** | 是否启用 TOTP。 | [optional] [default to undefined]
**enable_time** | **string** | 启用时间。 | [optional] [default to undefined]

## Example

```typescript
import { Totp } from '@bass/bbs-sdk-axios';

const instance: Totp = {
    user_id,
    enable,
    enable_time,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
