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

> ListCommentsResp callList(callListRequest)



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
        ListCommentsReq listCommentsReq = new ListCommentsReq(); // ListCommentsReq | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listCommentsReq(listCommentsReq)
                .build();
            ListCommentsResp result = apiInstance.callList(request);
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

[**ListCommentsResp**](ListCommentsResp.md)


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

> ApiResponse<ListCommentsResp> callListWithHttpInfo(callListRequest)



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
        ListCommentsReq listCommentsReq = new ListCommentsReq(); // ListCommentsReq | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listCommentsReq(listCommentsReq)
                .build();
            ApiResponse<ListCommentsResp> response = apiInstance.callListWithHttpInfo(request);
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

ApiResponse<[**ListCommentsResp**](ListCommentsResp.md)>


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
| **listCommentsReq** | [**ListCommentsReq**](ListCommentsReq.md) |  | |



## create

> CreateCommentResp create(createRequest)



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
        CreateCommentReq createCommentReq = new CreateCommentReq(); // CreateCommentReq | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createCommentReq(createCommentReq)
                .build();
            CreateCommentResp result = apiInstance.create(request);
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

[**CreateCommentResp**](CreateCommentResp.md)


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

> ApiResponse<CreateCommentResp> createWithHttpInfo(createRequest)



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
        CreateCommentReq createCommentReq = new CreateCommentReq(); // CreateCommentReq | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createCommentReq(createCommentReq)
                .build();
            ApiResponse<CreateCommentResp> response = apiInstance.createWithHttpInfo(request);
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

ApiResponse<[**CreateCommentResp**](CreateCommentResp.md)>


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
| **createCommentReq** | [**CreateCommentReq**](CreateCommentReq.md) |  | |



## like

> LikeCommentResp like(likeRequest)



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
        LikeCommentReq likeCommentReq = new LikeCommentReq(); // LikeCommentReq | 
        try {
            APIlikeRequest request = APIlikeRequest.newBuilder()
                .likeCommentReq(likeCommentReq)
                .build();
            LikeCommentResp result = apiInstance.like(request);
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

[**LikeCommentResp**](LikeCommentResp.md)


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

> ApiResponse<LikeCommentResp> likeWithHttpInfo(likeRequest)



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
        LikeCommentReq likeCommentReq = new LikeCommentReq(); // LikeCommentReq | 
        try {
            APIlikeRequest request = APIlikeRequest.newBuilder()
                .likeCommentReq(likeCommentReq)
                .build();
            ApiResponse<LikeCommentResp> response = apiInstance.likeWithHttpInfo(request);
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

ApiResponse<[**LikeCommentResp**](LikeCommentResp.md)>


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
| **likeCommentReq** | [**LikeCommentReq**](LikeCommentReq.md) |  | |



## listReplies

> ListCommentRepliesResp listReplies(listRepliesRequest)



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
        ListCommentRepliesReq listCommentRepliesReq = new ListCommentRepliesReq(); // ListCommentRepliesReq | 
        try {
            APIlistRepliesRequest request = APIlistRepliesRequest.newBuilder()
                .listCommentRepliesReq(listCommentRepliesReq)
                .build();
            ListCommentRepliesResp result = apiInstance.listReplies(request);
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

[**ListCommentRepliesResp**](ListCommentRepliesResp.md)


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

> ApiResponse<ListCommentRepliesResp> listRepliesWithHttpInfo(listRepliesRequest)



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
        ListCommentRepliesReq listCommentRepliesReq = new ListCommentRepliesReq(); // ListCommentRepliesReq | 
        try {
            APIlistRepliesRequest request = APIlistRepliesRequest.newBuilder()
                .listCommentRepliesReq(listCommentRepliesReq)
                .build();
            ApiResponse<ListCommentRepliesResp> response = apiInstance.listRepliesWithHttpInfo(request);
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

ApiResponse<[**ListCommentRepliesResp**](ListCommentRepliesResp.md)>


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
| **listCommentRepliesReq** | [**ListCommentRepliesReq**](ListCommentRepliesReq.md) |  | |



## listThreads

> ListCommentThreadsResp listThreads(listThreadsRequest)



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
        ListCommentThreadsReq listCommentThreadsReq = new ListCommentThreadsReq(); // ListCommentThreadsReq | 
        try {
            APIlistThreadsRequest request = APIlistThreadsRequest.newBuilder()
                .listCommentThreadsReq(listCommentThreadsReq)
                .build();
            ListCommentThreadsResp result = apiInstance.listThreads(request);
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

[**ListCommentThreadsResp**](ListCommentThreadsResp.md)


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

> ApiResponse<ListCommentThreadsResp> listThreadsWithHttpInfo(listThreadsRequest)



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
        ListCommentThreadsReq listCommentThreadsReq = new ListCommentThreadsReq(); // ListCommentThreadsReq | 
        try {
            APIlistThreadsRequest request = APIlistThreadsRequest.newBuilder()
                .listCommentThreadsReq(listCommentThreadsReq)
                .build();
            ApiResponse<ListCommentThreadsResp> response = apiInstance.listThreadsWithHttpInfo(request);
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

ApiResponse<[**ListCommentThreadsResp**](ListCommentThreadsResp.md)>


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
| **listCommentThreadsReq** | [**ListCommentThreadsReq**](ListCommentThreadsReq.md) |  | |



## listTimeline

> ListCommentTimelineResp listTimeline(listTimelineRequest)



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
        ListCommentTimelineReq listCommentTimelineReq = new ListCommentTimelineReq(); // ListCommentTimelineReq | 
        try {
            APIlistTimelineRequest request = APIlistTimelineRequest.newBuilder()
                .listCommentTimelineReq(listCommentTimelineReq)
                .build();
            ListCommentTimelineResp result = apiInstance.listTimeline(request);
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

[**ListCommentTimelineResp**](ListCommentTimelineResp.md)


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

> ApiResponse<ListCommentTimelineResp> listTimelineWithHttpInfo(listTimelineRequest)



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
        ListCommentTimelineReq listCommentTimelineReq = new ListCommentTimelineReq(); // ListCommentTimelineReq | 
        try {
            APIlistTimelineRequest request = APIlistTimelineRequest.newBuilder()
                .listCommentTimelineReq(listCommentTimelineReq)
                .build();
            ApiResponse<ListCommentTimelineResp> response = apiInstance.listTimelineWithHttpInfo(request);
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

ApiResponse<[**ListCommentTimelineResp**](ListCommentTimelineResp.md)>


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
| **listCommentTimelineReq** | [**ListCommentTimelineReq**](ListCommentTimelineReq.md) |  | |



## thank

> ThankCommentResp thank(thankRequest)



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
        ThankCommentReq thankCommentReq = new ThankCommentReq(); // ThankCommentReq | 
        try {
            APIthankRequest request = APIthankRequest.newBuilder()
                .thankCommentReq(thankCommentReq)
                .build();
            ThankCommentResp result = apiInstance.thank(request);
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

[**ThankCommentResp**](ThankCommentResp.md)


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

> ApiResponse<ThankCommentResp> thankWithHttpInfo(thankRequest)



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
        ThankCommentReq thankCommentReq = new ThankCommentReq(); // ThankCommentReq | 
        try {
            APIthankRequest request = APIthankRequest.newBuilder()
                .thankCommentReq(thankCommentReq)
                .build();
            ApiResponse<ThankCommentResp> response = apiInstance.thankWithHttpInfo(request);
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

ApiResponse<[**ThankCommentResp**](ThankCommentResp.md)>


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
| **thankCommentReq** | [**ThankCommentReq**](ThankCommentReq.md) |  | |


