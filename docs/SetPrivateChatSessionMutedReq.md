# SetPrivateChatSessionMutedReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**peer_user_id** | **string** | 对端账号 ID。 | [optional] [default to undefined]
**muted** | **boolean** | 是否免打扰。 | [optional] [default to undefined]

## Example

```typescript
import { SetPrivateChatSessionMutedReq } from '@bass/bbs-sdk-axios';

const instance: SetPrivateChatSessionMutedReq = {
    peer_user_id,
    muted,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
