# LoginReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **string** | 登录方式 | [default to undefined]
**password_credential** | [**ReqPasswordCredential**](ReqPasswordCredential.md) | 密码登录凭证 | [optional] [default to undefined]
**email_credential** | [**ReqEmailCredential**](ReqEmailCredential.md) | 邮箱登录凭证 | [optional] [default to undefined]
**phone_credential** | [**ReqPhoneCredential**](ReqPhoneCredential.md) | 手机登录凭证 | [optional] [default to undefined]

## Example

```typescript
import { LoginReq } from '@bass/bbs-sdk-axios';

const instance: LoginReq = {
    type,
    password_credential,
    email_credential,
    phone_credential,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
