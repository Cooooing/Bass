# ListNotificationsResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**page** | [**PageResp**](PageResp.md) | 分页结果 | [optional] [default to undefined]
**rows** | [**Array&lt;RespNotification&gt;**](RespNotification.md) | 通知列表 | [optional] [default to undefined]

## Example

```typescript
import { ListNotificationsResp } from '@bass/bbs-sdk-axios';

const instance: ListNotificationsResp = {
    page,
    rows,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
