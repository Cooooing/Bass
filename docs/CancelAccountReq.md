# CancelAccountReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**password** | **string** | 当前密码 | [default to undefined]
**code** | **string** | 启用 TOTP 时作为 TOTP code | [optional] [default to undefined]

## Example

```typescript
import { CancelAccountReq } from '@bass/bbs-sdk-axios';

const instance: CancelAccountReq = {
    password,
    code,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
