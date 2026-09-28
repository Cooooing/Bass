# WebSocketService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createTicket**](#createticket) | **POST** /v1/game-idle/ws/create-ticket | |

# **createTicket**
> CreateWebSocketTicketResp createTicket(createWebSocketTicketReq)

创建角色 WS 连接凭证。

### Example

```typescript
import {
    WebSocketService,
    Configuration,
    CreateWebSocketTicketReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new WebSocketService(configuration);

let createWebSocketTicketReq: CreateWebSocketTicketReq; //

const { status, data } = await apiInstance.createTicket(
    createWebSocketTicketReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createWebSocketTicketReq** | **CreateWebSocketTicketReq**|  | |


### Return type

**CreateWebSocketTicketResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

