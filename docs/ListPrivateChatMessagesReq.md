# ListPrivateChatMessagesReq


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**peer_user_id** | **string** |  | [optional] [default to undefined]
**before_id** | **string** |  | [optional] [default to undefined]
**page** | [**PageReq**](PageReq.md) |  | [optional] [default to undefined]

## Example

```typescript
import { ListPrivateChatMessagesReq } from '@bass/bbs-sdk-axios';

const instance: ListPrivateChatMessagesReq = {
    peer_user_id,
    before_id,
    page,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
