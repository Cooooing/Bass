# UpdatePhoneAccountReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**phone** | **string** | 新手机号，必填 | [default to undefined]
**code** | **string** | 手机 OTP，必填 | [default to undefined]

## Example

```typescript
import { UpdatePhoneAccountReq } from '@bass/bbs-sdk-axios';

const instance: UpdatePhoneAccountReq = {
    phone,
    code,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
