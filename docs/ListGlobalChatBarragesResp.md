# ListGlobalChatBarragesResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageResp**](PageResp.md) | 分页信息。 | [optional] [default to undefined]
**rows** | [**Array&lt;GlobalChatBarrage&gt;**](GlobalChatBarrage.md) | 已发布弹幕。 | [optional] [default to undefined]

## Example

```typescript
import { ListGlobalChatBarragesResp } from '@bass/bbs-sdk-axios';

const instance: ListGlobalChatBarragesResp = {
    page,
    rows,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
