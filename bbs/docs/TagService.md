# TagService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**callList**](TagService.md#callList) | **POST** /v1/content/tag/list |  |
| [**callListWithHttpInfo**](TagService.md#callListWithHttpInfo) | **POST** /v1/content/tag/list |  |



## callList

> ListTagsReply callList(callListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        ListTagsRequest listTagsRequest = new ListTagsRequest(); // ListTagsRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listTagsRequest(listTagsRequest)
                .build();
            ListTagsReply result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#callList");
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
| callListRequest | [**APIcallListRequest**](TagService.md#APIcallListRequest)|-|-|

### Return type

[**ListTagsReply**](ListTagsReply.md)


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

> ApiResponse<ListTagsReply> callListWithHttpInfo(callListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagService;
import com.bass.bbs.api.TagService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagService apiInstance = new TagService(defaultClient);
        ListTagsRequest listTagsRequest = new ListTagsRequest(); // ListTagsRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listTagsRequest(listTagsRequest)
                .build();
            ApiResponse<ListTagsReply> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#callList");
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
| callListRequest | [**APIcallListRequest**](TagService.md#APIcallListRequest)|-|-|

### Return type

ApiResponse<[**ListTagsReply**](ListTagsReply.md)>


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
| **listTagsRequest** | [**ListTagsRequest**](ListTagsRequest.md) |  | |


