# GlobalChatMessage

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 消息 ID。 | [optional] [default to undefined]
**sender_id** | **string** | 发送账号 ID。 | [optional] [default to undefined]
**content** | **string** | 原始 Markdown 内容。 | [optional] [default to undefined]
**message_type** | **string** | 消息类型。 | [optional] [default to undefined]
**revoked** | **boolean** | 是否已撤回。 | [optional] [default to undefined]
**sender** | [**GlobalChatMember**](GlobalChatMember.md) | 发送者展示资料。 | [optional] [default to undefined]
**created_at** | **string** | 创建时间。 | [optional] [default to undefined]
**revoked_at** | **string** | 撤回时间。 | [optional] [default to undefined]

## Example

```typescript
import { GlobalChatMessage } from '@bass/bbs-sdk-axios';

const instance: GlobalChatMessage = {
    id,
    sender_id,
    content,
    message_type,
    revoked,
    sender,
    created_at,
    revoked_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
