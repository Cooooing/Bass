# WebSocketService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createTicket**](WebSocketService.md#createticket) | **POST** /v1/game-idle/ws/create-ticket |  |



## createTicket

> CreateWebSocketTicketResp createTicket(createWebSocketTicketReq)



创建角色 WS 连接凭证。

### Example

```ts
import {
  Configuration,
  WebSocketService,
} from '@bass/bbs-sdk-fetch';
import type { CreateTicketRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new WebSocketService();

  const body = {
    // CreateWebSocketTicketReq
    createWebSocketTicketReq: ...,
  } satisfies CreateTicketRequest;

  try {
    const data = await api.createTicket(body);
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
| **createWebSocketTicketReq** | [CreateWebSocketTicketReq](CreateWebSocketTicketReq.md) |  | |

### Return type

[**CreateWebSocketTicketResp**](CreateWebSocketTicketResp.md)

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

