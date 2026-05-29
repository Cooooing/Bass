# CommentService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**callList**](CommentService.md#callList) | **POST** /v1/content/comment/list |  |
| [**callListWithHttpInfo**](CommentService.md#callListWithHttpInfo) | **POST** /v1/content/comment/list |  |
| [**create**](CommentService.md#create) | **POST** /v1/content/comment/create |  |
| [**createWithHttpInfo**](CommentService.md#createWithHttpInfo) | **POST** /v1/content/comment/create |  |
| [**like**](CommentService.md#like) | **POST** /v1/content/comment/like |  |
| [**likeWithHttpInfo**](CommentService.md#likeWithHttpInfo) | **POST** /v1/content/comment/like |  |
| [**thank**](CommentService.md#thank) | **POST** /v1/content/comment/thank |  |
| [**thankWithHttpInfo**](CommentService.md#thankWithHttpInfo) | **POST** /v1/content/comment/thank |  |



## callList

> ListCommentsReply callList(callListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentService;
import com.bass.bbs.api.CommentService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentService apiInstance = new CommentService(defaultClient);
        ListCommentsRequest listCommentsRequest = new ListCommentsRequest(); // ListCommentsRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listCommentsRequest(listCommentsRequest)
                .build();
            ListCommentsReply result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#callList");
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
| callListRequest | [**APIcallListRequest**](CommentService.md#APIcallListRequest)|-|-|

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

## callListWithHttpInfo

> ApiResponse<ListCommentsReply> callListWithHttpInfo(callListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentService;
import com.bass.bbs.api.CommentService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentService apiInstance = new CommentService(defaultClient);
        ListCommentsRequest listCommentsRequest = new ListCommentsRequest(); // ListCommentsRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listCommentsRequest(listCommentsRequest)
                .build();
            ApiResponse<ListCommentsReply> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#callList");
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
| callListRequest | [**APIcallListRequest**](CommentService.md#APIcallListRequest)|-|-|

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


<a id="APIcallListRequest"></a>
## APIcallListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listCommentsRequest** | [**ListCommentsRequest**](ListCommentsRequest.md) |  | |



## create

> CreateCommentReply create(createRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentService;
import com.bass.bbs.api.CommentService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentService apiInstance = new CommentService(defaultClient);
        CreateCommentRequest createCommentRequest = new CreateCommentRequest(); // CreateCommentRequest | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createCommentRequest(createCommentRequest)
                .build();
            CreateCommentReply result = apiInstance.create(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#create");
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
| createRequest | [**APIcreateRequest**](CommentService.md#APIcreateRequest)|-|-|

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

## createWithHttpInfo

> ApiResponse<CreateCommentReply> createWithHttpInfo(createRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentService;
import com.bass.bbs.api.CommentService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentService apiInstance = new CommentService(defaultClient);
        CreateCommentRequest createCommentRequest = new CreateCommentRequest(); // CreateCommentRequest | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createCommentRequest(createCommentRequest)
                .build();
            ApiResponse<CreateCommentReply> response = apiInstance.createWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#create");
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
| createRequest | [**APIcreateRequest**](CommentService.md#APIcreateRequest)|-|-|

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


<a id="APIcreateRequest"></a>
## APIcreateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createCommentRequest** | [**CreateCommentRequest**](CreateCommentRequest.md) |  | |



## like

> Object like(likeRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentService;
import com.bass.bbs.api.CommentService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentService apiInstance = new CommentService(defaultClient);
        LikeCommentRequest likeCommentRequest = new LikeCommentRequest(); // LikeCommentRequest | 
        try {
            APIlikeRequest request = APIlikeRequest.newBuilder()
                .likeCommentRequest(likeCommentRequest)
                .build();
            Object result = apiInstance.like(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#like");
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
| likeRequest | [**APIlikeRequest**](CommentService.md#APIlikeRequest)|-|-|

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

## likeWithHttpInfo

> ApiResponse<Object> likeWithHttpInfo(likeRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentService;
import com.bass.bbs.api.CommentService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentService apiInstance = new CommentService(defaultClient);
        LikeCommentRequest likeCommentRequest = new LikeCommentRequest(); // LikeCommentRequest | 
        try {
            APIlikeRequest request = APIlikeRequest.newBuilder()
                .likeCommentRequest(likeCommentRequest)
                .build();
            ApiResponse<Object> response = apiInstance.likeWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#like");
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
| likeRequest | [**APIlikeRequest**](CommentService.md#APIlikeRequest)|-|-|

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


<a id="APIlikeRequest"></a>
## APIlikeRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **likeCommentRequest** | [**LikeCommentRequest**](LikeCommentRequest.md) |  | |



## thank

> Object thank(thankRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentService;
import com.bass.bbs.api.CommentService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentService apiInstance = new CommentService(defaultClient);
        ThankCommentRequest thankCommentRequest = new ThankCommentRequest(); // ThankCommentRequest | 
        try {
            APIthankRequest request = APIthankRequest.newBuilder()
                .thankCommentRequest(thankCommentRequest)
                .build();
            Object result = apiInstance.thank(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#thank");
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
| thankRequest | [**APIthankRequest**](CommentService.md#APIthankRequest)|-|-|

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

## thankWithHttpInfo

> ApiResponse<Object> thankWithHttpInfo(thankRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CommentService;
import com.bass.bbs.api.CommentService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CommentService apiInstance = new CommentService(defaultClient);
        ThankCommentRequest thankCommentRequest = new ThankCommentRequest(); // ThankCommentRequest | 
        try {
            APIthankRequest request = APIthankRequest.newBuilder()
                .thankCommentRequest(thankCommentRequest)
                .build();
            ApiResponse<Object> response = apiInstance.thankWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#thank");
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
| thankRequest | [**APIthankRequest**](CommentService.md#APIthankRequest)|-|-|

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


<a id="APIthankRequest"></a>
## APIthankRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **thankCommentRequest** | [**ThankCommentRequest**](ThankCommentRequest.md) |  | |


