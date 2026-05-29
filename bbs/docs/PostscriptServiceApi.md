# PostscriptServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**postscriptServiceAdd**](PostscriptServiceApi.md#postscriptServiceAdd) | **POST** /v1/content/postscript/add |  |
| [**postscriptServiceAddWithHttpInfo**](PostscriptServiceApi.md#postscriptServiceAddWithHttpInfo) | **POST** /v1/content/postscript/add |  |



## postscriptServiceAdd

> AddPostscriptReply postscriptServiceAdd(postscriptServiceAddRequest)



添加文章附言

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PostscriptServiceApi;
import com.bass.bbs.api.PostscriptServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PostscriptServiceApi apiInstance = new PostscriptServiceApi(defaultClient);
        AddPostscriptRequest addPostscriptRequest = new AddPostscriptRequest(); // AddPostscriptRequest | 
        try {
            APIpostscriptServiceAddRequest request = APIpostscriptServiceAddRequest.newBuilder()
                .addPostscriptRequest(addPostscriptRequest)
                .build();
            AddPostscriptReply result = apiInstance.postscriptServiceAdd(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptServiceApi#postscriptServiceAdd");
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
| postscriptServiceAddRequest | [**APIpostscriptServiceAddRequest**](PostscriptServiceApi.md#APIpostscriptServiceAddRequest)|-|-|

### Return type

[**AddPostscriptReply**](AddPostscriptReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## postscriptServiceAddWithHttpInfo

> ApiResponse<AddPostscriptReply> postscriptServiceAddWithHttpInfo(postscriptServiceAddRequest)



添加文章附言

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PostscriptServiceApi;
import com.bass.bbs.api.PostscriptServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PostscriptServiceApi apiInstance = new PostscriptServiceApi(defaultClient);
        AddPostscriptRequest addPostscriptRequest = new AddPostscriptRequest(); // AddPostscriptRequest | 
        try {
            APIpostscriptServiceAddRequest request = APIpostscriptServiceAddRequest.newBuilder()
                .addPostscriptRequest(addPostscriptRequest)
                .build();
            ApiResponse<AddPostscriptReply> response = apiInstance.postscriptServiceAddWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptServiceApi#postscriptServiceAdd");
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
| postscriptServiceAddRequest | [**APIpostscriptServiceAddRequest**](PostscriptServiceApi.md#APIpostscriptServiceAddRequest)|-|-|

### Return type

ApiResponse<[**AddPostscriptReply**](AddPostscriptReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpostscriptServiceAddRequest"></a>
## APIpostscriptServiceAddRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **addPostscriptRequest** | [**AddPostscriptRequest**](AddPostscriptRequest.md) |  | |


