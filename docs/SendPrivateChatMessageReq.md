# SendPrivateChatMessageReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**receiver_id** | **string** | 接收账号 ID。 | [optional] [default to undefined]
**client_message_id** | **string** | 客户端幂等消息 ID。 | [optional] [default to undefined]
**content** | **string** | 原始 Markdown 内容。 | [optional] [default to undefined]

## Example

```typescript
import { SendPrivateChatMessageReq } from '@bass/bbs-sdk-axios';

const instance: SendPrivateChatMessageReq = {
    receiver_id,
    client_message_id,
    content,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
