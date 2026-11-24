# CommentService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createComment**](CommentService.md#createComment) | **POST** /v1/content/comment/create |  |
| [**createCommentWithHttpInfo**](CommentService.md#createCommentWithHttpInfo) | **POST** /v1/content/comment/create |  |
| [**likeComment**](CommentService.md#likeComment) | **POST** /v1/content/comment/like |  |
| [**likeCommentWithHttpInfo**](CommentService.md#likeCommentWithHttpInfo) | **POST** /v1/content/comment/like |  |
| [**listCommentReplies**](CommentService.md#listCommentReplies) | **POST** /v1/content/comment/list-replies |  |
| [**listCommentRepliesWithHttpInfo**](CommentService.md#listCommentRepliesWithHttpInfo) | **POST** /v1/content/comment/list-replies |  |
| [**listCommentThreads**](CommentService.md#listCommentThreads) | **POST** /v1/content/comment/list-threads |  |
| [**listCommentThreadsWithHttpInfo**](CommentService.md#listCommentThreadsWithHttpInfo) | **POST** /v1/content/comment/list-threads |  |
| [**listCommentTimeline**](CommentService.md#listCommentTimeline) | **POST** /v1/content/comment/list-timeline |  |
| [**listCommentTimelineWithHttpInfo**](CommentService.md#listCommentTimelineWithHttpInfo) | **POST** /v1/content/comment/list-timeline |  |
| [**listComments**](CommentService.md#listComments) | **POST** /v1/content/comment/list |  |
| [**listCommentsWithHttpInfo**](CommentService.md#listCommentsWithHttpInfo) | **POST** /v1/content/comment/list |  |
| [**thankComment**](CommentService.md#thankComment) | **POST** /v1/content/comment/thank |  |
| [**thankCommentWithHttpInfo**](CommentService.md#thankCommentWithHttpInfo) | **POST** /v1/content/comment/thank |  |



## createComment

> CreateCommentResp createComment(createCommentRequest)



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
            APIcreateCommentRequest request = APIcreateCommentRequest.newBuilder()
                .createCommentReq(createCommentReq)
                .build();
            CreateCommentResp result = apiInstance.createComment(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#createComment");
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
| createCommentRequest | [**APIcreateCommentRequest**](CommentService.md#APIcreateCommentRequest)|-|-|

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

## createCommentWithHttpInfo

> ApiResponse<CreateCommentResp> createCommentWithHttpInfo(createCommentRequest)



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
            APIcreateCommentRequest request = APIcreateCommentRequest.newBuilder()
                .createCommentReq(createCommentReq)
                .build();
            ApiResponse<CreateCommentResp> response = apiInstance.createCommentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#createComment");
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
| createCommentRequest | [**APIcreateCommentRequest**](CommentService.md#APIcreateCommentRequest)|-|-|

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


<a id="APIcreateCommentRequest"></a>
## APIcreateCommentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createCommentReq** | [**CreateCommentReq**](CreateCommentReq.md) |  | |



## likeComment

> LikeCommentResp likeComment(likeCommentRequest)



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
            APIlikeCommentRequest request = APIlikeCommentRequest.newBuilder()
                .likeCommentReq(likeCommentReq)
                .build();
            LikeCommentResp result = apiInstance.likeComment(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#likeComment");
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
| likeCommentRequest | [**APIlikeCommentRequest**](CommentService.md#APIlikeCommentRequest)|-|-|

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

## likeCommentWithHttpInfo

> ApiResponse<LikeCommentResp> likeCommentWithHttpInfo(likeCommentRequest)



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
            APIlikeCommentRequest request = APIlikeCommentRequest.newBuilder()
                .likeCommentReq(likeCommentReq)
                .build();
            ApiResponse<LikeCommentResp> response = apiInstance.likeCommentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#likeComment");
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
| likeCommentRequest | [**APIlikeCommentRequest**](CommentService.md#APIlikeCommentRequest)|-|-|

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


<a id="APIlikeCommentRequest"></a>
## APIlikeCommentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **likeCommentReq** | [**LikeCommentReq**](LikeCommentReq.md) |  | |



## listCommentReplies

> ListCommentRepliesResp listCommentReplies(listCommentRepliesRequest)



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
            APIlistCommentRepliesRequest request = APIlistCommentRepliesRequest.newBuilder()
                .listCommentRepliesReq(listCommentRepliesReq)
                .build();
            ListCommentRepliesResp result = apiInstance.listCommentReplies(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listCommentReplies");
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
| listCommentRepliesRequest | [**APIlistCommentRepliesRequest**](CommentService.md#APIlistCommentRepliesRequest)|-|-|

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

## listCommentRepliesWithHttpInfo

> ApiResponse<ListCommentRepliesResp> listCommentRepliesWithHttpInfo(listCommentRepliesRequest)



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
            APIlistCommentRepliesRequest request = APIlistCommentRepliesRequest.newBuilder()
                .listCommentRepliesReq(listCommentRepliesReq)
                .build();
            ApiResponse<ListCommentRepliesResp> response = apiInstance.listCommentRepliesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listCommentReplies");
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
| listCommentRepliesRequest | [**APIlistCommentRepliesRequest**](CommentService.md#APIlistCommentRepliesRequest)|-|-|

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


<a id="APIlistCommentRepliesRequest"></a>
## APIlistCommentRepliesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listCommentRepliesReq** | [**ListCommentRepliesReq**](ListCommentRepliesReq.md) |  | |



## listCommentThreads

> ListCommentThreadsResp listCommentThreads(listCommentThreadsRequest)



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
            APIlistCommentThreadsRequest request = APIlistCommentThreadsRequest.newBuilder()
                .listCommentThreadsReq(listCommentThreadsReq)
                .build();
            ListCommentThreadsResp result = apiInstance.listCommentThreads(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listCommentThreads");
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
| listCommentThreadsRequest | [**APIlistCommentThreadsRequest**](CommentService.md#APIlistCommentThreadsRequest)|-|-|

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

## listCommentThreadsWithHttpInfo

> ApiResponse<ListCommentThreadsResp> listCommentThreadsWithHttpInfo(listCommentThreadsRequest)



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
            APIlistCommentThreadsRequest request = APIlistCommentThreadsRequest.newBuilder()
                .listCommentThreadsReq(listCommentThreadsReq)
                .build();
            ApiResponse<ListCommentThreadsResp> response = apiInstance.listCommentThreadsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listCommentThreads");
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
| listCommentThreadsRequest | [**APIlistCommentThreadsRequest**](CommentService.md#APIlistCommentThreadsRequest)|-|-|

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


<a id="APIlistCommentThreadsRequest"></a>
## APIlistCommentThreadsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listCommentThreadsReq** | [**ListCommentThreadsReq**](ListCommentThreadsReq.md) |  | |



## listCommentTimeline

> ListCommentTimelineResp listCommentTimeline(listCommentTimelineRequest)



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
            APIlistCommentTimelineRequest request = APIlistCommentTimelineRequest.newBuilder()
                .listCommentTimelineReq(listCommentTimelineReq)
                .build();
            ListCommentTimelineResp result = apiInstance.listCommentTimeline(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listCommentTimeline");
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
| listCommentTimelineRequest | [**APIlistCommentTimelineRequest**](CommentService.md#APIlistCommentTimelineRequest)|-|-|

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

## listCommentTimelineWithHttpInfo

> ApiResponse<ListCommentTimelineResp> listCommentTimelineWithHttpInfo(listCommentTimelineRequest)



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
            APIlistCommentTimelineRequest request = APIlistCommentTimelineRequest.newBuilder()
                .listCommentTimelineReq(listCommentTimelineReq)
                .build();
            ApiResponse<ListCommentTimelineResp> response = apiInstance.listCommentTimelineWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listCommentTimeline");
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
| listCommentTimelineRequest | [**APIlistCommentTimelineRequest**](CommentService.md#APIlistCommentTimelineRequest)|-|-|

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


<a id="APIlistCommentTimelineRequest"></a>
## APIlistCommentTimelineRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listCommentTimelineReq** | [**ListCommentTimelineReq**](ListCommentTimelineReq.md) |  | |



## listComments

> ListCommentsResp listComments(listCommentsRequest)



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
            APIlistCommentsRequest request = APIlistCommentsRequest.newBuilder()
                .listCommentsReq(listCommentsReq)
                .build();
            ListCommentsResp result = apiInstance.listComments(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listComments");
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
| listCommentsRequest | [**APIlistCommentsRequest**](CommentService.md#APIlistCommentsRequest)|-|-|

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

## listCommentsWithHttpInfo

> ApiResponse<ListCommentsResp> listCommentsWithHttpInfo(listCommentsRequest)



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
            APIlistCommentsRequest request = APIlistCommentsRequest.newBuilder()
                .listCommentsReq(listCommentsReq)
                .build();
            ApiResponse<ListCommentsResp> response = apiInstance.listCommentsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#listComments");
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
| listCommentsRequest | [**APIlistCommentsRequest**](CommentService.md#APIlistCommentsRequest)|-|-|

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


<a id="APIlistCommentsRequest"></a>
## APIlistCommentsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listCommentsReq** | [**ListCommentsReq**](ListCommentsReq.md) |  | |



## thankComment

> ThankCommentResp thankComment(thankCommentRequest)



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
            APIthankCommentRequest request = APIthankCommentRequest.newBuilder()
                .thankCommentReq(thankCommentReq)
                .build();
            ThankCommentResp result = apiInstance.thankComment(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#thankComment");
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
| thankCommentRequest | [**APIthankCommentRequest**](CommentService.md#APIthankCommentRequest)|-|-|

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

## thankCommentWithHttpInfo

> ApiResponse<ThankCommentResp> thankCommentWithHttpInfo(thankCommentRequest)



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
            APIthankCommentRequest request = APIthankCommentRequest.newBuilder()
                .thankCommentReq(thankCommentReq)
                .build();
            ApiResponse<ThankCommentResp> response = apiInstance.thankCommentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CommentService#thankComment");
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
| thankCommentRequest | [**APIthankCommentRequest**](CommentService.md#APIthankCommentRequest)|-|-|

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


<a id="APIthankCommentRequest"></a>
## APIthankCommentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **thankCommentReq** | [**ThankCommentReq**](ThankCommentReq.md) |  | |


