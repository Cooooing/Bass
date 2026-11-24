# CreateRealtimeTicketResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ticket** | **string** | 一次性 SSE 建连凭证。 | [optional] [default to undefined]
**expires_in_seconds** | **string** | 有效期秒数。 | [optional] [default to undefined]
**path** | **string** | realtime node SSE 路径。 | [optional] [default to undefined]

## Example

```typescript
import { CreateRealtimeTicketResp } from '@bass/bbs-sdk-axios';

const instance: CreateRealtimeTicketResp = {
    ticket,
    expires_in_seconds,
    path,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
