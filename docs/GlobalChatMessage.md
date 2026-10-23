# GlobalChatMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **string** |  | [optional] [default to undefined]
**sender_id** | **string** |  | [optional] [default to undefined]
**content** | **string** |  | [optional] [default to undefined]
**message_type** | **string** |  | [optional] [default to undefined]
**revoked** | **boolean** |  | [optional] [default to undefined]
**sender** | [**GlobalChatMember**](GlobalChatMember.md) |  | [optional] [default to undefined]
**created_at** | **string** |  | [optional] [default to undefined]
**revoked_at** | **string** |  | [optional] [default to undefined]

## Example

```typescript
import { GlobalChatMessage } from '@bass/bbs-sdk-axios';

const instance: GlobalChatMessage = {
    id,
    sender_id,
    content,
    message_type,
    revoked,
    sender,
    created_at,
    revoked_at,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
