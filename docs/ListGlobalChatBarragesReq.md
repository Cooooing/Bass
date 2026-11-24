# ListGlobalChatBarragesReq

请求参数。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**user_id** | **string** | 按发送账号筛选；仅管理端可用。 | [optional] [default to undefined]
**page** | [**PageReq**](PageReq.md) | 分页参数。 | [optional] [default to undefined]

## Example

```typescript
import { ListGlobalChatBarragesReq } from '@bass/bbs-sdk-axios';

const instance: ListGlobalChatBarragesReq = {
    user_id,
    page,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
