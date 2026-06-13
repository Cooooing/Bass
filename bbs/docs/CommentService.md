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
| [**listReplies**](CommentService.md#listReplies) | **POST** /v1/content/comment/list-replies |  |
| [**listRepliesWithHttpInfo**](CommentService.md#listRepliesWithHttpInfo) | **POST** /v1/content/comment/list-replies |  |
| [**listThreads**](CommentService.md#listThreads) | **POST** /v1/content/comment/list-threads |  |
| [**listThreadsWithHttpInfo**](CommentService.md#listThreadsWithHttpInfo) | **POST** /v1/content/comment/list-threads |  |
| [**listTimeline**](CommentService.md#listTimeline) | **POST** /v1/content/comment/list-timeline |  |
| [**listTimelineWithHttpInfo**](CommentService.md#listTimelineWithHttpInfo) | **POST** /v1/content/comment/list-timeline |  |
| [**thank**](CommentService.md#thank) | **POST** /v1/content/comment/thank |  |
| [**thankWithHttpInfo**](CommentService.md#thankWithHttpInfo) | **POST** /v1/content/comment/thank |  |



## callList

> ListCommentsReply callList(callListRequest)



分页查询评论列表。

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



分页查询评论列表。

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



创建评论。

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



创建评论。

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

> LikeCommentReply like(likeRequest)



点赞或取消点赞评论。

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
            LikeCommentReply result = apiInstance.like(request);
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

[**LikeCommentReply**](LikeCommentReply.md)


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

> ApiResponse<LikeCommentReply> likeWithHttpInfo(likeRequest)



点赞或取消点赞评论。

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
            ApiResponse<LikeCommentReply> response = apiInstance.likeWithHttpInfo(request);
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

ApiResponse<[**LikeCommentReply**](LikeCommentReply.md)>


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



## listReplies

> ListCommentRepliesReply listReplies(listRepliesRequest)



分页查询评论回复。

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
        ListCommentRepliesRequest listCommentRepliesRequest = new ListCommentRepliesRequest(); // ListCommentRepliesRequest | 
        try {
            APIlistRepliesRequest request = APIlistRepliesRequest.newBuilder()
                .listCommentRepliesRequest(listCommentRepliesRequest)
                .build();
            ListCommentRepliesReply result = apiInstance.listReplies(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listReplies");
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
| listRepliesRequest | [**APIlistRepliesRequest**](CommentService.md#APIlistRepliesRequest)|-|-|

### Return type

[**ListCommentRepliesReply**](ListCommentRepliesReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listRepliesWithHttpInfo

> ApiResponse<ListCommentRepliesReply> listRepliesWithHttpInfo(listRepliesRequest)



分页查询评论回复。

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
        ListCommentRepliesRequest listCommentRepliesRequest = new ListCommentRepliesRequest(); // ListCommentRepliesRequest | 
        try {
            APIlistRepliesRequest request = APIlistRepliesRequest.newBuilder()
                .listCommentRepliesRequest(listCommentRepliesRequest)
                .build();
            ApiResponse<ListCommentRepliesReply> response = apiInstance.listRepliesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listReplies");
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
| listRepliesRequest | [**APIlistRepliesRequest**](CommentService.md#APIlistRepliesRequest)|-|-|

### Return type

ApiResponse<[**ListCommentRepliesReply**](ListCommentRepliesReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistRepliesRequest"></a>
## APIlistRepliesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listCommentRepliesRequest** | [**ListCommentRepliesRequest**](ListCommentRepliesRequest.md) |  | |



## listThreads

> ListCommentThreadsReply listThreads(listThreadsRequest)



分页查询评论楼层。

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
        ListCommentThreadsRequest listCommentThreadsRequest = new ListCommentThreadsRequest(); // ListCommentThreadsRequest | 
        try {
            APIlistThreadsRequest request = APIlistThreadsRequest.newBuilder()
                .listCommentThreadsRequest(listCommentThreadsRequest)
                .build();
            ListCommentThreadsReply result = apiInstance.listThreads(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listThreads");
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
| listThreadsRequest | [**APIlistThreadsRequest**](CommentService.md#APIlistThreadsRequest)|-|-|

### Return type

[**ListCommentThreadsReply**](ListCommentThreadsReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listThreadsWithHttpInfo

> ApiResponse<ListCommentThreadsReply> listThreadsWithHttpInfo(listThreadsRequest)



分页查询评论楼层。

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
        ListCommentThreadsRequest listCommentThreadsRequest = new ListCommentThreadsRequest(); // ListCommentThreadsRequest | 
        try {
            APIlistThreadsRequest request = APIlistThreadsRequest.newBuilder()
                .listCommentThreadsRequest(listCommentThreadsRequest)
                .build();
            ApiResponse<ListCommentThreadsReply> response = apiInstance.listThreadsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listThreads");
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
| listThreadsRequest | [**APIlistThreadsRequest**](CommentService.md#APIlistThreadsRequest)|-|-|

### Return type

ApiResponse<[**ListCommentThreadsReply**](ListCommentThreadsReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistThreadsRequest"></a>
## APIlistThreadsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listCommentThreadsRequest** | [**ListCommentThreadsRequest**](ListCommentThreadsRequest.md) |  | |



## listTimeline

> ListCommentTimelineReply listTimeline(listTimelineRequest)



分页查询评论时间线。

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
        ListCommentTimelineRequest listCommentTimelineRequest = new ListCommentTimelineRequest(); // ListCommentTimelineRequest | 
        try {
            APIlistTimelineRequest request = APIlistTimelineRequest.newBuilder()
                .listCommentTimelineRequest(listCommentTimelineRequest)
                .build();
            ListCommentTimelineReply result = apiInstance.listTimeline(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listTimeline");
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
| listTimelineRequest | [**APIlistTimelineRequest**](CommentService.md#APIlistTimelineRequest)|-|-|

### Return type

[**ListCommentTimelineReply**](ListCommentTimelineReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listTimelineWithHttpInfo

> ApiResponse<ListCommentTimelineReply> listTimelineWithHttpInfo(listTimelineRequest)



分页查询评论时间线。

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
        ListCommentTimelineRequest listCommentTimelineRequest = new ListCommentTimelineRequest(); // ListCommentTimelineRequest | 
        try {
            APIlistTimelineRequest request = APIlistTimelineRequest.newBuilder()
                .listCommentTimelineRequest(listCommentTimelineRequest)
                .build();
            ApiResponse<ListCommentTimelineReply> response = apiInstance.listTimelineWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listTimeline");
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
| listTimelineRequest | [**APIlistTimelineRequest**](CommentService.md#APIlistTimelineRequest)|-|-|

### Return type

ApiResponse<[**ListCommentTimelineReply**](ListCommentTimelineReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistTimelineRequest"></a>
## APIlistTimelineRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listCommentTimelineRequest** | [**ListCommentTimelineRequest**](ListCommentTimelineRequest.md) |  | |



## thank

> ThankCommentReply thank(thankRequest)



感谢或取消感谢评论。

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
            ThankCommentReply result = apiInstance.thank(request);
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

[**ThankCommentReply**](ThankCommentReply.md)


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

> ApiResponse<ThankCommentReply> thankWithHttpInfo(thankRequest)



感谢或取消感谢评论。

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
            ApiResponse<ThankCommentReply> response = apiInstance.thankWithHttpInfo(request);
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

ApiResponse<[**ThankCommentReply**](ThankCommentReply.md)>


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


