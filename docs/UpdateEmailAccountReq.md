# UpdateEmailAccountReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email** | **string** | 新邮箱地址，必填 | [default to undefined]
**code** | **string** | 邮箱 OTP，必填 | [default to undefined]

## Example

```typescript
import { UpdateEmailAccountReq } from '@bass/bbs-sdk-axios';

const instance: UpdateEmailAccountReq = {
    email,
    code,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
