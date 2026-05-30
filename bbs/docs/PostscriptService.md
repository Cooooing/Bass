# PostscriptService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**add**](PostscriptService.md#add) | **POST** /v1/content/postscript/add |  |
| [**addWithHttpInfo**](PostscriptService.md#addWithHttpInfo) | **POST** /v1/content/postscript/add |  |



## add

> AddPostscriptReply add(addRequest)



添加文章附言。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PostscriptService;
import com.bass.bbs.api.PostscriptService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PostscriptService apiInstance = new PostscriptService(defaultClient);
        AddPostscriptRequest addPostscriptRequest = new AddPostscriptRequest(); // AddPostscriptRequest | 
        try {
            APIaddRequest request = APIaddRequest.newBuilder()
                .addPostscriptRequest(addPostscriptRequest)
                .build();
            AddPostscriptReply result = apiInstance.add(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptService#add");
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
| addRequest | [**APIaddRequest**](PostscriptService.md#APIaddRequest)|-|-|

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

## addWithHttpInfo

> ApiResponse<AddPostscriptReply> addWithHttpInfo(addRequest)



添加文章附言。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PostscriptService;
import com.bass.bbs.api.PostscriptService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PostscriptService apiInstance = new PostscriptService(defaultClient);
        AddPostscriptRequest addPostscriptRequest = new AddPostscriptRequest(); // AddPostscriptRequest | 
        try {
            APIaddRequest request = APIaddRequest.newBuilder()
                .addPostscriptRequest(addPostscriptRequest)
                .build();
            ApiResponse<AddPostscriptReply> response = apiInstance.addWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptService#add");
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
| addRequest | [**APIaddRequest**](PostscriptService.md#APIaddRequest)|-|-|

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


<a id="APIaddRequest"></a>
## APIaddRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **addPostscriptRequest** | [**AddPostscriptRequest**](AddPostscriptRequest.md) |  | |


