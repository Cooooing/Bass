# GlobalChatBarrage

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 弹幕 ID。 | [optional] [default to undefined]
**user_id** | **string** | 发送账号 ID。 | [optional] [default to undefined]
**content** | **string** | 纯文本弹幕内容。 | [optional] [default to undefined]
**points_cost** | **string** | 本次实际消耗的积分。 | [optional] [default to undefined]
**sender** | [**GlobalChatMember**](GlobalChatMember.md) | 发送者展示资料。 | [optional] [default to undefined]
**created_at** | **string** | 创建时间。 | [optional] [default to undefined]

## Example

```typescript
import { GlobalChatBarrage } from '@bass/bbs-sdk-axios';

const instance: GlobalChatBarrage = {
    id,
    user_id,
    content,
    points_cost,
    sender,
    created_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
