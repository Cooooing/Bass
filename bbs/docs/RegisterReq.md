# RegisterReq


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **string** |  | [default to undefined]
**name** | **string** |  | [default to undefined]
**password** | **string** |  | [default to undefined]
**nickname** | **string** |  | [optional] [default to undefined]
**email_credential** | [**ReqEmailCredential**](ReqEmailCredential.md) |  | [optional] [default to undefined]
**phone_credential** | [**ReqPhoneCredential**](ReqPhoneCredential.md) |  | [optional] [default to undefined]

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
