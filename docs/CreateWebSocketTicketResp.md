# CreateWebSocketTicketResp

响应结果。

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**character_id** | **string** | 角色 ID | [optional] [default to undefined]
**ticket** | **string** | WS 连接凭证 | [optional] [default to undefined]
**expires_in_seconds** | **string** | 剩余有效秒数 | [optional] [default to undefined]
**path** | **string** | WS 连接路径 | [optional] [default to undefined]

## Example

```typescript
import { CreateWebSocketTicketResp } from '@bass/bbs-sdk-axios';

const instance: CreateWebSocketTicketResp = {
    character_id,
    ticket,
    expires_in_seconds,
    path,
};
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
