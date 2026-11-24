# SendGlobalChatMessageReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**client_message_id** | **string** | 客户端幂等消息 ID。 | [optional] [default to undefined]
**content** | **string** | 原始 Markdown 内容。 | [optional] [default to undefined]

## Example

```typescript
import { SendGlobalChatMessageReq } from '@bass/bbs-sdk-axios';

const instance: SendGlobalChatMessageReq = {
    client_message_id,
    content,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
