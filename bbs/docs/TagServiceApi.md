# TagServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**tagServiceList**](TagServiceApi.md#tagServiceList) | **POST** /v1/content/tag/list |  |
| [**tagServiceListWithHttpInfo**](TagServiceApi.md#tagServiceListWithHttpInfo) | **POST** /v1/content/tag/list |  |



## tagServiceList

> ListTagsReply tagServiceList(tagServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagServiceApi;
import com.bass.bbs.api.TagServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagServiceApi apiInstance = new TagServiceApi(defaultClient);
        ListTagsRequest listTagsRequest = new ListTagsRequest(); // ListTagsRequest | 
        try {
            APItagServiceListRequest request = APItagServiceListRequest.newBuilder()
                .listTagsRequest(listTagsRequest)
                .build();
            ListTagsReply result = apiInstance.tagServiceList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagServiceApi#tagServiceList");
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
| tagServiceListRequest | [**APItagServiceListRequest**](TagServiceApi.md#APItagServiceListRequest)|-|-|

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

## tagServiceListWithHttpInfo

> ApiResponse<ListTagsReply> tagServiceListWithHttpInfo(tagServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TagServiceApi;
import com.bass.bbs.api.TagServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TagServiceApi apiInstance = new TagServiceApi(defaultClient);
        ListTagsRequest listTagsRequest = new ListTagsRequest(); // ListTagsRequest | 
        try {
            APItagServiceListRequest request = APItagServiceListRequest.newBuilder()
                .listTagsRequest(listTagsRequest)
                .build();
            ApiResponse<ListTagsReply> response = apiInstance.tagServiceListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagServiceApi#tagServiceList");
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
| tagServiceListRequest | [**APItagServiceListRequest**](TagServiceApi.md#APItagServiceListRequest)|-|-|

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


<a id="APItagServiceListRequest"></a>
## APItagServiceListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listTagsRequest** | [**ListTagsRequest**](ListTagsRequest.md) |  | |


