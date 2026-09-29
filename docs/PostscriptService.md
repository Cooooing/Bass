# PostscriptService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**add**](PostscriptService.md#add) | **POST** /v1/content/postscript/add |  |
| [**addWithHttpInfo**](PostscriptService.md#addWithHttpInfo) | **POST** /v1/content/postscript/add |  |
| [**callList**](PostscriptService.md#callList) | **POST** /v1/content/postscript/list |  |
| [**callListWithHttpInfo**](PostscriptService.md#callListWithHttpInfo) | **POST** /v1/content/postscript/list |  |



## add

> AddPostscriptResp add(addRequest)



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
        AddPostscriptReq addPostscriptReq = new AddPostscriptReq(); // AddPostscriptReq | 
        try {
            APIaddRequest request = APIaddRequest.newBuilder()
                .addPostscriptReq(addPostscriptReq)
                .build();
            AddPostscriptResp result = apiInstance.add(request);
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

[**AddPostscriptResp**](AddPostscriptResp.md)


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

> ApiResponse<AddPostscriptResp> addWithHttpInfo(addRequest)



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
        AddPostscriptReq addPostscriptReq = new AddPostscriptReq(); // AddPostscriptReq | 
        try {
            APIaddRequest request = APIaddRequest.newBuilder()
                .addPostscriptReq(addPostscriptReq)
                .build();
            ApiResponse<AddPostscriptResp> response = apiInstance.addWithHttpInfo(request);
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

ApiResponse<[**AddPostscriptResp**](AddPostscriptResp.md)>


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
| **addPostscriptReq** | [**AddPostscriptReq**](AddPostscriptReq.md) |  | |



## callList

> ListPostscriptsResp callList(callListRequest)



查询文章附言列表。

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
        ListPostscriptsReq listPostscriptsReq = new ListPostscriptsReq(); // ListPostscriptsReq | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listPostscriptsReq(listPostscriptsReq)
                .build();
            ListPostscriptsResp result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptService#callList");
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
| callListRequest | [**APIcallListRequest**](PostscriptService.md#APIcallListRequest)|-|-|

### Return type

[**ListPostscriptsResp**](ListPostscriptsResp.md)


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

> ApiResponse<ListPostscriptsResp> callListWithHttpInfo(callListRequest)



查询文章附言列表。

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
        ListPostscriptsReq listPostscriptsReq = new ListPostscriptsReq(); // ListPostscriptsReq | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listPostscriptsReq(listPostscriptsReq)
                .build();
            ApiResponse<ListPostscriptsResp> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptService#callList");
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
| callListRequest | [**APIcallListRequest**](PostscriptService.md#APIcallListRequest)|-|-|

### Return type

ApiResponse<[**ListPostscriptsResp**](ListPostscriptsResp.md)>


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
| **listPostscriptsReq** | [**ListPostscriptsReq**](ListPostscriptsReq.md) |  | |


