# LoginReq


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **string** |  | [default to undefined]
**password_credential** | [**ReqPasswordCredential**](ReqPasswordCredential.md) |  | [optional] [default to undefined]
**email_credential** | [**ReqEmailCredential**](ReqEmailCredential.md) |  | [optional] [default to undefined]
**phone_credential** | [**ReqPhoneCredential**](ReqPhoneCredential.md) |  | [optional] [default to undefined]

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
