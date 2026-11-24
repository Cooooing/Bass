# ProfileRelation

协议消息。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**following** | **boolean** | 协议字段。 | [optional] [default to undefined]
**followed_by** | **boolean** | 协议字段。 | [optional] [default to undefined]
**blocking** | **boolean** | 协议字段。 | [optional] [default to undefined]
**blocked_by** | **boolean** | 协议字段。 | [optional] [default to undefined]

## Example

```typescript
import { ProfileRelation } from '@bass/bbs-sdk-axios';

const instance: ProfileRelation = {
    following,
    followed_by,
    blocking,
    blocked_by,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
