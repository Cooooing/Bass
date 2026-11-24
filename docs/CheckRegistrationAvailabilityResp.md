# CheckRegistrationAvailabilityResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name_available** | **boolean** | 账号名是否可用 | [optional] [default to undefined]
**email_can_register** | **boolean** | 邮箱是否可继续用于注册 | [optional] [default to undefined]
**phone_can_register** | **boolean** | 手机号是否可继续用于注册 | [optional] [default to undefined]

## Example

```typescript
import { CheckRegistrationAvailabilityResp } from '@bass/bbs-sdk-axios';

const instance: CheckRegistrationAvailabilityResp = {
    name_available,
    email_can_register,
    phone_can_register,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
