# PrivateChatMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** |  | [optional] [default to undefined]
**sender_id** | **string** |  | [optional] [default to undefined]
**receiver_id** | **string** |  | [optional] [default to undefined]
**content** | **string** |  | [optional] [default to undefined]
**message_type** | **string** |  | [optional] [default to undefined]
**revoked** | **boolean** |  | [optional] [default to undefined]
**created_at** | **string** |  | [optional] [default to undefined]
**revoked_at** | **string** |  | [optional] [default to undefined]

## Example

```typescript
import { PrivateChatMessage } from '@bass/bbs-sdk-axios';

const instance: PrivateChatMessage = {
    id,
    sender_id,
    receiver_id,
    content,
    message_type,
    revoked,
    created_at,
    revoked_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
