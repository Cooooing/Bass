# CommentServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**commentServiceCreate**](CommentServiceApi.md#commentServiceCreate) | **POST** /v1/content/comment/create |  |
| [**commentServiceCreateWithHttpInfo**](CommentServiceApi.md#commentServiceCreateWithHttpInfo) | **POST** /v1/content/comment/create |  |
| [**commentServiceLike**](CommentServiceApi.md#commentServiceLike) | **POST** /v1/content/comment/like |  |
| [**commentServiceLikeWithHttpInfo**](CommentServiceApi.md#commentServiceLikeWithHttpInfo) | **POST** /v1/content/comment/like |  |
| [**commentServiceList**](CommentServiceApi.md#commentServiceList) | **POST** /v1/content/comment/list |  |
| [**commentServiceListWithHttpInfo**](CommentServiceApi.md#commentServiceListWithHttpInfo) | **POST** /v1/content/comment/list |  |
| [**commentServiceThank**](CommentServiceApi.md#commentServiceThank) | **POST** /v1/content/comment/thank |  |
| [**commentServiceThankWithHttpInfo**](CommentServiceApi.md#commentServiceThankWithHttpInfo) | **POST** /v1/content/comment/thank |  |



## commentServiceCreate

> CreateCommentReply commentServiceCreate(commentServiceCreateRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentServiceApi;
import com.bass.bbs.api.CommentServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentServiceApi apiInstance = new CommentServiceApi(defaultClient);
        CreateCommentRequest createCommentRequest = new CreateCommentRequest(); // CreateCommentRequest | 
        try {
            APIcommentServiceCreateRequest request = APIcommentServiceCreateRequest.newBuilder()
                .createCommentRequest(createCommentRequest)
                .build();
            CreateCommentReply result = apiInstance.commentServiceCreate(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentServiceApi#commentServiceCreate");
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
| commentServiceCreateRequest | [**APIcommentServiceCreateRequest**](CommentServiceApi.md#APIcommentServiceCreateRequest)|-|-|

### Return type

[**CreateCommentReply**](CreateCommentReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## commentServiceCreateWithHttpInfo

> ApiResponse<CreateCommentReply> commentServiceCreateWithHttpInfo(commentServiceCreateRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentServiceApi;
import com.bass.bbs.api.CommentServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentServiceApi apiInstance = new CommentServiceApi(defaultClient);
        CreateCommentRequest createCommentRequest = new CreateCommentRequest(); // CreateCommentRequest | 
        try {
            APIcommentServiceCreateRequest request = APIcommentServiceCreateRequest.newBuilder()
                .createCommentRequest(createCommentRequest)
                .build();
            ApiResponse<CreateCommentReply> response = apiInstance.commentServiceCreateWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentServiceApi#commentServiceCreate");
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
| commentServiceCreateRequest | [**APIcommentServiceCreateRequest**](CommentServiceApi.md#APIcommentServiceCreateRequest)|-|-|

### Return type

ApiResponse<[**CreateCommentReply**](CreateCommentReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcommentServiceCreateRequest"></a>
## APIcommentServiceCreateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createCommentRequest** | [**CreateCommentRequest**](CreateCommentRequest.md) |  | |



## commentServiceLike

> Object commentServiceLike(commentServiceLikeRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentServiceApi;
import com.bass.bbs.api.CommentServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentServiceApi apiInstance = new CommentServiceApi(defaultClient);
        LikeCommentRequest likeCommentRequest = new LikeCommentRequest(); // LikeCommentRequest | 
        try {
            APIcommentServiceLikeRequest request = APIcommentServiceLikeRequest.newBuilder()
                .likeCommentRequest(likeCommentRequest)
                .build();
            Object result = apiInstance.commentServiceLike(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentServiceApi#commentServiceLike");
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
| commentServiceLikeRequest | [**APIcommentServiceLikeRequest**](CommentServiceApi.md#APIcommentServiceLikeRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## commentServiceLikeWithHttpInfo

> ApiResponse<Object> commentServiceLikeWithHttpInfo(commentServiceLikeRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentServiceApi;
import com.bass.bbs.api.CommentServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentServiceApi apiInstance = new CommentServiceApi(defaultClient);
        LikeCommentRequest likeCommentRequest = new LikeCommentRequest(); // LikeCommentRequest | 
        try {
            APIcommentServiceLikeRequest request = APIcommentServiceLikeRequest.newBuilder()
                .likeCommentRequest(likeCommentRequest)
                .build();
            ApiResponse<Object> response = apiInstance.commentServiceLikeWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentServiceApi#commentServiceLike");
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
| commentServiceLikeRequest | [**APIcommentServiceLikeRequest**](CommentServiceApi.md#APIcommentServiceLikeRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcommentServiceLikeRequest"></a>
## APIcommentServiceLikeRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **likeCommentRequest** | [**LikeCommentRequest**](LikeCommentRequest.md) |  | |



## commentServiceList

> ListCommentsReply commentServiceList(commentServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentServiceApi;
import com.bass.bbs.api.CommentServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentServiceApi apiInstance = new CommentServiceApi(defaultClient);
        ListCommentsRequest listCommentsRequest = new ListCommentsRequest(); // ListCommentsRequest | 
        try {
            APIcommentServiceListRequest request = APIcommentServiceListRequest.newBuilder()
                .listCommentsRequest(listCommentsRequest)
                .build();
            ListCommentsReply result = apiInstance.commentServiceList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentServiceApi#commentServiceList");
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
| commentServiceListRequest | [**APIcommentServiceListRequest**](CommentServiceApi.md#APIcommentServiceListRequest)|-|-|

### Return type

[**ListCommentsReply**](ListCommentsReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## commentServiceListWithHttpInfo

> ApiResponse<ListCommentsReply> commentServiceListWithHttpInfo(commentServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentServiceApi;
import com.bass.bbs.api.CommentServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentServiceApi apiInstance = new CommentServiceApi(defaultClient);
        ListCommentsRequest listCommentsRequest = new ListCommentsRequest(); // ListCommentsRequest | 
        try {
            APIcommentServiceListRequest request = APIcommentServiceListRequest.newBuilder()
                .listCommentsRequest(listCommentsRequest)
                .build();
            ApiResponse<ListCommentsReply> response = apiInstance.commentServiceListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentServiceApi#commentServiceList");
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
| commentServiceListRequest | [**APIcommentServiceListRequest**](CommentServiceApi.md#APIcommentServiceListRequest)|-|-|

### Return type

ApiResponse<[**ListCommentsReply**](ListCommentsReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcommentServiceListRequest"></a>
## APIcommentServiceListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listCommentsRequest** | [**ListCommentsRequest**](ListCommentsRequest.md) |  | |



## commentServiceThank

> Object commentServiceThank(commentServiceThankRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentServiceApi;
import com.bass.bbs.api.CommentServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentServiceApi apiInstance = new CommentServiceApi(defaultClient);
        ThankCommentRequest thankCommentRequest = new ThankCommentRequest(); // ThankCommentRequest | 
        try {
            APIcommentServiceThankRequest request = APIcommentServiceThankRequest.newBuilder()
                .thankCommentRequest(thankCommentRequest)
                .build();
            Object result = apiInstance.commentServiceThank(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentServiceApi#commentServiceThank");
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
| commentServiceThankRequest | [**APIcommentServiceThankRequest**](CommentServiceApi.md#APIcommentServiceThankRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## commentServiceThankWithHttpInfo

> ApiResponse<Object> commentServiceThankWithHttpInfo(commentServiceThankRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentServiceApi;
import com.bass.bbs.api.CommentServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentServiceApi apiInstance = new CommentServiceApi(defaultClient);
        ThankCommentRequest thankCommentRequest = new ThankCommentRequest(); // ThankCommentRequest | 
        try {
            APIcommentServiceThankRequest request = APIcommentServiceThankRequest.newBuilder()
                .thankCommentRequest(thankCommentRequest)
                .build();
            ApiResponse<Object> response = apiInstance.commentServiceThankWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentServiceApi#commentServiceThank");
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
| commentServiceThankRequest | [**APIcommentServiceThankRequest**](CommentServiceApi.md#APIcommentServiceThankRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcommentServiceThankRequest"></a>
## APIcommentServiceThankRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **thankCommentRequest** | [**ThankCommentRequest**](ThankCommentRequest.md) |  | |


