# WebSocketService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createSession**](#createsession) | **POST** /v1/game-idle/ws/create-session | |

# **createSession**
> CreateWebSocketSessionResp createSession(createWebSocketSessionReq)

创建角色 WS session。

### Example

```typescript
import {
    WebSocketService,
    Configuration,
    CreateWebSocketSessionReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new WebSocketService(configuration);

let createWebSocketSessionReq: CreateWebSocketSessionReq; //

const { status, data } = await apiInstance.createSession(
    createWebSocketSessionReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createWebSocketSessionReq** | **CreateWebSocketSessionReq**|  | |


### Return type

**CreateWebSocketSessionResp**

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

