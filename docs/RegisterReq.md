# RegisterReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **string** | 注册方式 | [default to undefined]
**name** | **string** | 账号名 | [default to undefined]
**password** | **string** | 账号密码 | [default to undefined]
**nickname** | **string** | 昵称 | [optional] [default to undefined]
**email_credential** | [**ReqEmailCredential**](ReqEmailCredential.md) | 邮箱注册凭证 | [optional] [default to undefined]
**phone_credential** | [**ReqPhoneCredential**](ReqPhoneCredential.md) | 手机注册凭证 | [optional] [default to undefined]

## Example

```typescript
import { RegisterReq } from '@bass/bbs-sdk-axios';

const instance: RegisterReq = {
    type,
    name,
    password,
    nickname,
    email_credential,
    phone_credential,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
