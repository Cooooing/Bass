# DomainService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**callList**](DomainService.md#callList) | **POST** /v1/content/domain/list |  |
| [**callListWithHttpInfo**](DomainService.md#callListWithHttpInfo) | **POST** /v1/content/domain/list |  |



## callList

> ListDomainsReply callList(callListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainService;
import com.bass.bbs.api.DomainService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainService apiInstance = new DomainService(defaultClient);
        ListDomainsRequest listDomainsRequest = new ListDomainsRequest(); // ListDomainsRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listDomainsRequest(listDomainsRequest)
                .build();
            ListDomainsReply result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#callList");
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
| callListRequest | [**APIcallListRequest**](DomainService.md#APIcallListRequest)|-|-|

### Return type

[**ListDomainsReply**](ListDomainsReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## callListWithHttpInfo

> ApiResponse<ListDomainsReply> callListWithHttpInfo(callListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.DomainService;
import com.bass.bbs.api.DomainService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        DomainService apiInstance = new DomainService(defaultClient);
        ListDomainsRequest listDomainsRequest = new ListDomainsRequest(); // ListDomainsRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listDomainsRequest(listDomainsRequest)
                .build();
            ApiResponse<ListDomainsReply> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling DomainService#callList");
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
| callListRequest | [**APIcallListRequest**](DomainService.md#APIcallListRequest)|-|-|

### Return type

ApiResponse<[**ListDomainsReply**](ListDomainsReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcallListRequest"></a>
## APIcallListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listDomainsRequest** | [**ListDomainsRequest**](ListDomainsRequest.md) |  | |


