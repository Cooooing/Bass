# GlobalChatTopic

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 话题记录 ID。 | [optional] [default to undefined]
**user_id** | **string** | 设置话题的账号 ID。 | [optional] [default to undefined]
**content** | **string** | 话题内容。 | [optional] [default to undefined]
**points_cost** | **string** | 本次设置实际消耗的积分。 | [optional] [default to undefined]
**created_at** | **string** | 话题设置时间。 | [optional] [default to undefined]

## Example

```typescript
import { GlobalChatTopic } from '@bass/bbs-sdk-axios';

const instance: GlobalChatTopic = {
    id,
    user_id,
    content,
    points_cost,
    created_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
