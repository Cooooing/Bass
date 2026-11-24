# PrivateChatSession

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**peer_user_id** | **string** | 对端账号 ID。 | [optional] [default to undefined]
**last_read_message_id** | **string** | 最后已读消息 ID。 | [optional] [default to undefined]
**unread_count** | **number** | 未读消息数。 | [optional] [default to undefined]
**muted** | **boolean** | 是否免打扰。 | [optional] [default to undefined]
**pinned** | **boolean** | 是否置顶。 | [optional] [default to undefined]
**last_message** | [**PrivateChatMessage**](PrivateChatMessage.md) | 最后一条消息。 | [optional] [default to undefined]
**last_message_at** | **string** | 最后消息时间。 | [optional] [default to undefined]
**peer** | [**PrivateChatMember**](PrivateChatMember.md) | 对端展示资料。 | [optional] [default to undefined]

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
