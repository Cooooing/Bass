# WebSocketService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createSession**](WebSocketService.md#createSession) | **POST** /v1/game-idle/ws/create-session |  |
| [**createSessionWithHttpInfo**](WebSocketService.md#createSessionWithHttpInfo) | **POST** /v1/game-idle/ws/create-session |  |



## createSession

> CreateWebSocketSessionResp createSession(createSessionRequest)



创建角色 WS session。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.WebSocketService;
import com.bass.bbs.api.WebSocketService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        WebSocketService apiInstance = new WebSocketService(defaultClient);
        CreateWebSocketSessionReq createWebSocketSessionReq = new CreateWebSocketSessionReq(); // CreateWebSocketSessionReq | 
        try {
            APIcreateSessionRequest request = APIcreateSessionRequest.newBuilder()
                .createWebSocketSessionReq(createWebSocketSessionReq)
                .build();
            CreateWebSocketSessionResp result = apiInstance.createSession(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling WebSocketService#createSession");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| createSessionRequest | [**APIcreateSessionRequest**](WebSocketService.md#APIcreateSessionRequest)|-|-|

### Return type

[**CreateWebSocketSessionResp**](CreateWebSocketSessionResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## createSessionWithHttpInfo

> ApiResponse<CreateWebSocketSessionResp> createSessionWithHttpInfo(createSessionRequest)



创建角色 WS session。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.WebSocketService;
import com.bass.bbs.api.WebSocketService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        WebSocketService apiInstance = new WebSocketService(defaultClient);
        CreateWebSocketSessionReq createWebSocketSessionReq = new CreateWebSocketSessionReq(); // CreateWebSocketSessionReq | 
        try {
            APIcreateSessionRequest request = APIcreateSessionRequest.newBuilder()
                .createWebSocketSessionReq(createWebSocketSessionReq)
                .build();
            ApiResponse<CreateWebSocketSessionResp> response = apiInstance.createSessionWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling WebSocketService#createSession");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| createSessionRequest | [**APIcreateSessionRequest**](WebSocketService.md#APIcreateSessionRequest)|-|-|

### Return type

ApiResponse<[**CreateWebSocketSessionResp**](CreateWebSocketSessionResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcreateSessionRequest"></a>
## APIcreateSessionRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createWebSocketSessionReq** | [**CreateWebSocketSessionReq**](CreateWebSocketSessionReq.md) |  | |


