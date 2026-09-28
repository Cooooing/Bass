# RespRelationStatus


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | **string** |  | [optional] [default to undefined]
**following** | **boolean** |  | [optional] [default to undefined]
**followed_by** | **boolean** |  | [optional] [default to undefined]
**blocking** | **boolean** |  | [optional] [default to undefined]
**blocked_by** | **boolean** |  | [optional] [default to undefined]

## Example

```typescript
import { RespRelationStatus } from '@bass/bbs-sdk-axios';

const instance: RespRelationStatus = {
    target_id,
    following,
    followed_by,
    blocking,
    blocked_by,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
