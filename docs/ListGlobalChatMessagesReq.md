# ListGlobalChatMessagesReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**before_id** | **string** | 向前分页游标。 | [optional] [default to undefined]
**page** | [**PageReq**](PageReq.md) | 分页参数。 | [optional] [default to undefined]

## Example

```typescript
import { ListGlobalChatMessagesReq } from '@bass/bbs-sdk-axios';

const instance: ListGlobalChatMessagesReq = {
    before_id,
    page,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
