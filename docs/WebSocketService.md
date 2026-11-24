# WebSocketService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createWebSocketTicket**](WebSocketService.md#createWebSocketTicket) | **POST** /v1/game-idle/ws/create-ticket |  |
| [**createWebSocketTicketWithHttpInfo**](WebSocketService.md#createWebSocketTicketWithHttpInfo) | **POST** /v1/game-idle/ws/create-ticket |  |



## createWebSocketTicket

> CreateWebSocketTicketResp createWebSocketTicket(createWebSocketTicketRequest)



创建角色 WS 连接凭证。

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
        CreateWebSocketTicketReq createWebSocketTicketReq = new CreateWebSocketTicketReq(); // CreateWebSocketTicketReq | 
        try {
            APIcreateWebSocketTicketRequest request = APIcreateWebSocketTicketRequest.newBuilder()
                .createWebSocketTicketReq(createWebSocketTicketReq)
                .build();
            CreateWebSocketTicketResp result = apiInstance.createWebSocketTicket(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling WebSocketService#createWebSocketTicket");
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
| createWebSocketTicketRequest | [**APIcreateWebSocketTicketRequest**](WebSocketService.md#APIcreateWebSocketTicketRequest)|-|-|

### Return type

[**CreateWebSocketTicketResp**](CreateWebSocketTicketResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## createWebSocketTicketWithHttpInfo

> ApiResponse<CreateWebSocketTicketResp> createWebSocketTicketWithHttpInfo(createWebSocketTicketRequest)



创建角色 WS 连接凭证。

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
        CreateWebSocketTicketReq createWebSocketTicketReq = new CreateWebSocketTicketReq(); // CreateWebSocketTicketReq | 
        try {
            APIcreateWebSocketTicketRequest request = APIcreateWebSocketTicketRequest.newBuilder()
                .createWebSocketTicketReq(createWebSocketTicketReq)
                .build();
            ApiResponse<CreateWebSocketTicketResp> response = apiInstance.createWebSocketTicketWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling WebSocketService#createWebSocketTicket");
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
| createWebSocketTicketRequest | [**APIcreateWebSocketTicketRequest**](WebSocketService.md#APIcreateWebSocketTicketRequest)|-|-|

### Return type

ApiResponse<[**CreateWebSocketTicketResp**](CreateWebSocketTicketResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcreateWebSocketTicketRequest"></a>
## APIcreateWebSocketTicketRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createWebSocketTicketReq** | [**CreateWebSocketTicketReq**](CreateWebSocketTicketReq.md) |  | |


