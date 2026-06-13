# TagService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**callList**](TagService.md#callList) | **POST** /v1/content/tag/list |  |
| [**callListWithHttpInfo**](TagService.md#callListWithHttpInfo) | **POST** /v1/content/tag/list |  |
| [**create**](TagService.md#create) | **POST** /v1/content/tag/create |  |
| [**createWithHttpInfo**](TagService.md#createWithHttpInfo) | **POST** /v1/content/tag/create |  |
| [**update**](TagService.md#update) | **POST** /v1/content/tag/update |  |
| [**updateWithHttpInfo**](TagService.md#updateWithHttpInfo) | **POST** /v1/content/tag/update |  |



## callList

> ListTagsReply callList(callListRequest)



分页查询标签列表。

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



分页查询标签列表。

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



## create

> CreateTagReply create(createRequest)



创建标签。

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
        CreateTagRequest createTagRequest = new CreateTagRequest(); // CreateTagRequest | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createTagRequest(createTagRequest)
                .build();
            CreateTagReply result = apiInstance.create(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#create");
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
| createRequest | [**APIcreateRequest**](TagService.md#APIcreateRequest)|-|-|

### Return type

[**CreateTagReply**](CreateTagReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## createWithHttpInfo

> ApiResponse<CreateTagReply> createWithHttpInfo(createRequest)



创建标签。

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
        CreateTagRequest createTagRequest = new CreateTagRequest(); // CreateTagRequest | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createTagRequest(createTagRequest)
                .build();
            ApiResponse<CreateTagReply> response = apiInstance.createWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#create");
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
| createRequest | [**APIcreateRequest**](TagService.md#APIcreateRequest)|-|-|

### Return type

ApiResponse<[**CreateTagReply**](CreateTagReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcreateRequest"></a>
## APIcreateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createTagRequest** | [**CreateTagRequest**](CreateTagRequest.md) |  | |



## update

> UpdateTagReply update(updateRequest)



更新标签。

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
        UpdateTagRequest updateTagRequest = new UpdateTagRequest(); // UpdateTagRequest | 
        try {
            APIupdateRequest request = APIupdateRequest.newBuilder()
                .updateTagRequest(updateTagRequest)
                .build();
            UpdateTagReply result = apiInstance.update(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#update");
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
| updateRequest | [**APIupdateRequest**](TagService.md#APIupdateRequest)|-|-|

### Return type

[**UpdateTagReply**](UpdateTagReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## updateWithHttpInfo

> ApiResponse<UpdateTagReply> updateWithHttpInfo(updateRequest)



更新标签。

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
        UpdateTagRequest updateTagRequest = new UpdateTagRequest(); // UpdateTagRequest | 
        try {
            APIupdateRequest request = APIupdateRequest.newBuilder()
                .updateTagRequest(updateTagRequest)
                .build();
            ApiResponse<UpdateTagReply> response = apiInstance.updateWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TagService#update");
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
| updateRequest | [**APIupdateRequest**](TagService.md#APIupdateRequest)|-|-|

### Return type

ApiResponse<[**UpdateTagReply**](UpdateTagReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIupdateRequest"></a>
## APIupdateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateTagRequest** | [**UpdateTagRequest**](UpdateTagRequest.md) |  | |


