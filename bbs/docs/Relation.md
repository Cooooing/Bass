# Relation

账号关系记录。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** | 关系记录 ID。 | [optional] [default to undefined]
**type** | **number** | 关系类型。 | [optional] [default to undefined]
**actor_id** | **string** | 发起账号 ID。 | [optional] [default to undefined]
**target_id** | **string** | 目标账号 ID。 | [optional] [default to undefined]
**created_at** | **string** | 创建时间。 | [optional] [default to undefined]
**updated_at** | **string** | 更新时间。 | [optional] [default to undefined]

## Example

```typescript
import { Relation } from '@bass/bbs-sdk-axios';

const instance: Relation = {
    id,
    type,
    actor_id,
    target_id,
    created_at,
    updated_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
