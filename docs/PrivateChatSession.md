# PrivateChatSession


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**peer_user_id** | **string** |  | [optional] [default to undefined]
**last_read_message_id** | **string** |  | [optional] [default to undefined]
**unread_count** | **number** |  | [optional] [default to undefined]
**muted** | **boolean** |  | [optional] [default to undefined]
**pinned** | **boolean** |  | [optional] [default to undefined]
**last_message** | [**PrivateChatMessage**](PrivateChatMessage.md) |  | [optional] [default to undefined]
**last_message_at** | **string** |  | [optional] [default to undefined]
**peer** | [**PrivateChatMember**](PrivateChatMember.md) |  | [optional] [default to undefined]

## Example

```typescript
import { PrivateChatSession } from '@bass/bbs-sdk-axios';

const instance: PrivateChatSession = {
    peer_user_id,
    last_read_message_id,
    unread_count,
    muted,
    pinned,
    last_message,
    last_message_at,
    peer,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
