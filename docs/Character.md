# Character

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 角色 ID | [optional] [default to undefined]
**name** | **string** | 角色名称 | [optional] [default to undefined]
**status** | **string** | 角色状态 | [optional] [default to undefined]
**slot** | **number** | 角色槽位 | [optional] [default to undefined]
**action_queue_capacity** | **number** | 行动队列容量上限 | [optional] [default to undefined]
**max_offline_seconds** | **string** | 最大离线收益结算秒数 | [optional] [default to undefined]
**created_at** | **string** | 创建时间 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间 | [optional] [default to undefined]
**last_offline_at** | **string** | 最近离线时间 | [optional] [default to undefined]

## Example

```typescript
import { Character } from '@bass/bbs-sdk-axios';

const instance: Character = {
    id,
    name,
    status,
    slot,
    action_queue_capacity,
    max_offline_seconds,
    created_at,
    updated_at,
    last_offline_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
