# LoginResp


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_token** | **string** |  | [optional] [default to undefined]
**refresh_token** | **string** |  | [optional] [default to undefined]
**access_token_expires_at** | **string** |  | [optional] [default to undefined]
**refresh_token_expires_at** | **string** |  | [optional] [default to undefined]
**session_expires_at** | **string** |  | [optional] [default to undefined]
**account** | [**RespAccount**](RespAccount.md) |  | [optional] [default to undefined]

## Example

```typescript
import { LoginResp } from '@bass/bbs-sdk-axios';

const instance: LoginResp = {
    access_token,
    refresh_token,
    access_token_expires_at,
    refresh_token_expires_at,
    session_expires_at,
    account,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
