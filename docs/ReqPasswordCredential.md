# ReqPasswordCredential

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**account** | **string** | 账号名、邮箱或手机号 | [optional] [default to undefined]
**password** | **string** | 密码 | [optional] [default to undefined]
**code** | **string** | 账号启用 TOTP 时作为 TOTP code | [optional] [default to undefined]

## Example

```typescript
import { ReqPasswordCredential } from '@bass/bbs-sdk-axios';

const instance: ReqPasswordCredential = {
    account,
    password,
    code,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
