# MarkPrivateChatSessionReadReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**peer_user_id** | **string** | 对端账号 ID。 | [optional] [default to undefined]
**message_id** | **string** | 最后已读消息 ID。 | [optional] [default to undefined]

## Example

```typescript
import { MarkPrivateChatSessionReadReq } from '@bass/bbs-sdk-axios';

const instance: MarkPrivateChatSessionReadReq = {
    peer_user_id,
    message_id,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
