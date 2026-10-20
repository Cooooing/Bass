# RealtimeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createTicket**](RealtimeService.md#createTicket) | **POST** /v1/realtime/ticket |  |
| [**createTicketWithHttpInfo**](RealtimeService.md#createTicketWithHttpInfo) | **POST** /v1/realtime/ticket |  |



## createTicket

> CreateRealtimeTicketResp createTicket(createTicketRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RealtimeService;
import com.bass.bbs.api.RealtimeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RealtimeService apiInstance = new RealtimeService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcreateTicketRequest request = APIcreateTicketRequest.newBuilder()
                .body(body)
                .build();
            CreateRealtimeTicketResp result = apiInstance.createTicket(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling RealtimeService#createTicket");
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
| createTicketRequest | [**APIcreateTicketRequest**](RealtimeService.md#APIcreateTicketRequest)|-|-|

### Return type

[**CreateRealtimeTicketResp**](CreateRealtimeTicketResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## createTicketWithHttpInfo

> ApiResponse<CreateRealtimeTicketResp> createTicketWithHttpInfo(createTicketRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.RealtimeService;
import com.bass.bbs.api.RealtimeService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        RealtimeService apiInstance = new RealtimeService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcreateTicketRequest request = APIcreateTicketRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<CreateRealtimeTicketResp> response = apiInstance.createTicketWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling RealtimeService#createTicket");
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
| createTicketRequest | [**APIcreateTicketRequest**](RealtimeService.md#APIcreateTicketRequest)|-|-|

### Return type

ApiResponse<[**CreateRealtimeTicketResp**](CreateRealtimeTicketResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcreateTicketRequest"></a>
## APIcreateTicketRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |


