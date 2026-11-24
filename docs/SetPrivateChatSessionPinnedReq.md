# SetPrivateChatSessionPinnedReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**peer_user_id** | **string** | 对端账号 ID。 | [optional] [default to undefined]
**pinned** | **boolean** | 是否置顶。 | [optional] [default to undefined]

## Example

```typescript
import { SetPrivateChatSessionPinnedReq } from '@bass/bbs-sdk-axios';

const instance: SetPrivateChatSessionPinnedReq = {
    peer_user_id,
    pinned,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
