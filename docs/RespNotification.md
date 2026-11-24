# RespNotification

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 通知 ID | [optional] [default to undefined]
**event_id** | **string** | 来源事件 ID | [optional] [default to undefined]
**receiver_id** | **string** | 接收账号 ID | [optional] [default to undefined]
**notification_type** | **string** | 通知类型 | [optional] [default to undefined]
**title** | **string** | 通知标题 | [optional] [default to undefined]
**content** | **string** | 通知内容 | [optional] [default to undefined]
**read_at** | **string** | 已读时间 | [optional] [default to undefined]
**created_at** | **string** | 创建时间 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间 | [optional] [default to undefined]

## Example

```typescript
import { RespNotification } from '@bass/bbs-sdk-axios';

const instance: RespNotification = {
    id,
    event_id,
    receiver_id,
    notification_type,
    title,
    content,
    read_at,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
