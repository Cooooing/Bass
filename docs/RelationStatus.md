# RelationStatus

当前账号与目标账号的关系状态。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | **string** | 目标账号 ID。 | [optional] [default to undefined]
**following** | **boolean** | 当前账号是否关注目标账号。 | [optional] [default to undefined]
**followed_by** | **boolean** | 目标账号是否关注当前账号。 | [optional] [default to undefined]
**blocking** | **boolean** | 当前账号是否拉黑目标账号。 | [optional] [default to undefined]
**blocked_by** | **boolean** | 目标账号是否拉黑当前账号。 | [optional] [default to undefined]

## Example

```typescript
import { RelationStatus } from '@bass/bbs-sdk-axios';

const instance: RelationStatus = {
    target_id,
    following,
    followed_by,
    blocking,
    blocked_by,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
