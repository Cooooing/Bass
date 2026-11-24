# RefreshTokenResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_token** | **string** | 新的访问令牌 | [optional] [default to undefined]
**refresh_token** | **string** | 新的刷新令牌 | [optional] [default to undefined]
**access_token_expires_at** | **string** | 访问令牌过期时间 | [optional] [default to undefined]
**refresh_token_expires_at** | **string** | 刷新令牌空闲过期时间 | [optional] [default to undefined]
**session_expires_at** | **string** | session 绝对过期时间 | [optional] [default to undefined]

## Example

```typescript
import { RefreshTokenResp } from '@bass/bbs-sdk-axios';

const instance: RefreshTokenResp = {
    access_token,
    refresh_token,
    access_token_expires_at,
    refresh_token_expires_at,
    session_expires_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
