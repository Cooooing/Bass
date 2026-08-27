# WebSocketService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createSession**](WebSocketService.md#createsession) | **POST** /v1/game-idle/ws/create-session |  |



## createSession

> CreateWebSocketSessionResp createSession(createWebSocketSessionReq)



创建角色 WS session。

### Example

```ts
import {
  Configuration,
  WebSocketService,
} from '@bass/bbs-sdk-fetch';
import type { CreateSessionRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new WebSocketService();

  const body = {
    // CreateWebSocketSessionReq
    createWebSocketSessionReq: ...,
  } satisfies CreateSessionRequest;

  try {
    const data = await api.createSession(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **createWebSocketSessionReq** | [CreateWebSocketSessionReq](CreateWebSocketSessionReq.md) |  | |

### Return type

[**CreateWebSocketSessionResp**](CreateWebSocketSessionResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

