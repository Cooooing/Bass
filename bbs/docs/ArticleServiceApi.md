# ArticleServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**articleServiceAcceptAnswer**](ArticleServiceApi.md#articleServiceAcceptAnswer) | **POST** /v1/content/article/accept-answer |  |
| [**articleServiceAcceptAnswerWithHttpInfo**](ArticleServiceApi.md#articleServiceAcceptAnswerWithHttpInfo) | **POST** /v1/content/article/accept-answer |  |
| [**articleServiceCollect**](ArticleServiceApi.md#articleServiceCollect) | **POST** /v1/content/article/collect |  |
| [**articleServiceCollectWithHttpInfo**](ArticleServiceApi.md#articleServiceCollectWithHttpInfo) | **POST** /v1/content/article/collect |  |
| [**articleServiceCreate**](ArticleServiceApi.md#articleServiceCreate) | **POST** /v1/content/article/create |  |
| [**articleServiceCreateWithHttpInfo**](ArticleServiceApi.md#articleServiceCreateWithHttpInfo) | **POST** /v1/content/article/create |  |
| [**articleServiceDelete**](ArticleServiceApi.md#articleServiceDelete) | **POST** /v1/content/article/delete |  |
| [**articleServiceDeleteWithHttpInfo**](ArticleServiceApi.md#articleServiceDeleteWithHttpInfo) | **POST** /v1/content/article/delete |  |
| [**articleServiceGet**](ArticleServiceApi.md#articleServiceGet) | **POST** /v1/content/article/get |  |
| [**articleServiceGetWithHttpInfo**](ArticleServiceApi.md#articleServiceGetWithHttpInfo) | **POST** /v1/content/article/get |  |
| [**articleServiceLike**](ArticleServiceApi.md#articleServiceLike) | **POST** /v1/content/article/like |  |
| [**articleServiceLikeWithHttpInfo**](ArticleServiceApi.md#articleServiceLikeWithHttpInfo) | **POST** /v1/content/article/like |  |
| [**articleServiceList**](ArticleServiceApi.md#articleServiceList) | **POST** /v1/content/article/list |  |
| [**articleServiceListWithHttpInfo**](ArticleServiceApi.md#articleServiceListWithHttpInfo) | **POST** /v1/content/article/list |  |
| [**articleServicePublish**](ArticleServiceApi.md#articleServicePublish) | **POST** /v1/content/article/publish |  |
| [**articleServicePublishWithHttpInfo**](ArticleServiceApi.md#articleServicePublishWithHttpInfo) | **POST** /v1/content/article/publish |  |
| [**articleServiceReward**](ArticleServiceApi.md#articleServiceReward) | **POST** /v1/content/article/reward |  |
| [**articleServiceRewardWithHttpInfo**](ArticleServiceApi.md#articleServiceRewardWithHttpInfo) | **POST** /v1/content/article/reward |  |
| [**articleServiceThank**](ArticleServiceApi.md#articleServiceThank) | **POST** /v1/content/article/thank |  |
| [**articleServiceThankWithHttpInfo**](ArticleServiceApi.md#articleServiceThankWithHttpInfo) | **POST** /v1/content/article/thank |  |
| [**articleServiceUpdateDraft**](ArticleServiceApi.md#articleServiceUpdateDraft) | **POST** /v1/content/article/update-draft |  |
| [**articleServiceUpdateDraftWithHttpInfo**](ArticleServiceApi.md#articleServiceUpdateDraftWithHttpInfo) | **POST** /v1/content/article/update-draft |  |
| [**articleServiceWatch**](ArticleServiceApi.md#articleServiceWatch) | **POST** /v1/content/article/watch |  |
| [**articleServiceWatchWithHttpInfo**](ArticleServiceApi.md#articleServiceWatchWithHttpInfo) | **POST** /v1/content/article/watch |  |



## articleServiceAcceptAnswer

> Object articleServiceAcceptAnswer(articleServiceAcceptAnswerRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        AcceptAnswerArticleRequest acceptAnswerArticleRequest = new AcceptAnswerArticleRequest(); // AcceptAnswerArticleRequest | 
        try {
            APIarticleServiceAcceptAnswerRequest request = APIarticleServiceAcceptAnswerRequest.newBuilder()
                .acceptAnswerArticleRequest(acceptAnswerArticleRequest)
                .build();
            Object result = apiInstance.articleServiceAcceptAnswer(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceAcceptAnswer");
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
| articleServiceAcceptAnswerRequest | [**APIarticleServiceAcceptAnswerRequest**](ArticleServiceApi.md#APIarticleServiceAcceptAnswerRequest)|-|-|

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

## articleServiceAcceptAnswerWithHttpInfo

> ApiResponse<Object> articleServiceAcceptAnswerWithHttpInfo(articleServiceAcceptAnswerRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        AcceptAnswerArticleRequest acceptAnswerArticleRequest = new AcceptAnswerArticleRequest(); // AcceptAnswerArticleRequest | 
        try {
            APIarticleServiceAcceptAnswerRequest request = APIarticleServiceAcceptAnswerRequest.newBuilder()
                .acceptAnswerArticleRequest(acceptAnswerArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.articleServiceAcceptAnswerWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceAcceptAnswer");
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
| articleServiceAcceptAnswerRequest | [**APIarticleServiceAcceptAnswerRequest**](ArticleServiceApi.md#APIarticleServiceAcceptAnswerRequest)|-|-|

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


<a id="APIarticleServiceAcceptAnswerRequest"></a>
## APIarticleServiceAcceptAnswerRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **acceptAnswerArticleRequest** | [**AcceptAnswerArticleRequest**](AcceptAnswerArticleRequest.md) |  | |



## articleServiceCollect

> Object articleServiceCollect(articleServiceCollectRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        CollectArticleRequest collectArticleRequest = new CollectArticleRequest(); // CollectArticleRequest | 
        try {
            APIarticleServiceCollectRequest request = APIarticleServiceCollectRequest.newBuilder()
                .collectArticleRequest(collectArticleRequest)
                .build();
            Object result = apiInstance.articleServiceCollect(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceCollect");
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
| articleServiceCollectRequest | [**APIarticleServiceCollectRequest**](ArticleServiceApi.md#APIarticleServiceCollectRequest)|-|-|

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

## articleServiceCollectWithHttpInfo

> ApiResponse<Object> articleServiceCollectWithHttpInfo(articleServiceCollectRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        CollectArticleRequest collectArticleRequest = new CollectArticleRequest(); // CollectArticleRequest | 
        try {
            APIarticleServiceCollectRequest request = APIarticleServiceCollectRequest.newBuilder()
                .collectArticleRequest(collectArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.articleServiceCollectWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceCollect");
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
| articleServiceCollectRequest | [**APIarticleServiceCollectRequest**](ArticleServiceApi.md#APIarticleServiceCollectRequest)|-|-|

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


<a id="APIarticleServiceCollectRequest"></a>
## APIarticleServiceCollectRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **collectArticleRequest** | [**CollectArticleRequest**](CollectArticleRequest.md) |  | |



## articleServiceCreate

> CreateArticleReply articleServiceCreate(articleServiceCreateRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        CreateArticleRequest createArticleRequest = new CreateArticleRequest(); // CreateArticleRequest | 
        try {
            APIarticleServiceCreateRequest request = APIarticleServiceCreateRequest.newBuilder()
                .createArticleRequest(createArticleRequest)
                .build();
            CreateArticleReply result = apiInstance.articleServiceCreate(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceCreate");
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
| articleServiceCreateRequest | [**APIarticleServiceCreateRequest**](ArticleServiceApi.md#APIarticleServiceCreateRequest)|-|-|

### Return type

[**CreateArticleReply**](CreateArticleReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## articleServiceCreateWithHttpInfo

> ApiResponse<CreateArticleReply> articleServiceCreateWithHttpInfo(articleServiceCreateRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        CreateArticleRequest createArticleRequest = new CreateArticleRequest(); // CreateArticleRequest | 
        try {
            APIarticleServiceCreateRequest request = APIarticleServiceCreateRequest.newBuilder()
                .createArticleRequest(createArticleRequest)
                .build();
            ApiResponse<CreateArticleReply> response = apiInstance.articleServiceCreateWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceCreate");
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
| articleServiceCreateRequest | [**APIarticleServiceCreateRequest**](ArticleServiceApi.md#APIarticleServiceCreateRequest)|-|-|

### Return type

ApiResponse<[**CreateArticleReply**](CreateArticleReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIarticleServiceCreateRequest"></a>
## APIarticleServiceCreateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createArticleRequest** | [**CreateArticleRequest**](CreateArticleRequest.md) |  | |



## articleServiceDelete

> Object articleServiceDelete(articleServiceDeleteRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        DeleteArticleRequest deleteArticleRequest = new DeleteArticleRequest(); // DeleteArticleRequest | 
        try {
            APIarticleServiceDeleteRequest request = APIarticleServiceDeleteRequest.newBuilder()
                .deleteArticleRequest(deleteArticleRequest)
                .build();
            Object result = apiInstance.articleServiceDelete(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceDelete");
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
| articleServiceDeleteRequest | [**APIarticleServiceDeleteRequest**](ArticleServiceApi.md#APIarticleServiceDeleteRequest)|-|-|

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

## articleServiceDeleteWithHttpInfo

> ApiResponse<Object> articleServiceDeleteWithHttpInfo(articleServiceDeleteRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        DeleteArticleRequest deleteArticleRequest = new DeleteArticleRequest(); // DeleteArticleRequest | 
        try {
            APIarticleServiceDeleteRequest request = APIarticleServiceDeleteRequest.newBuilder()
                .deleteArticleRequest(deleteArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.articleServiceDeleteWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceDelete");
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
| articleServiceDeleteRequest | [**APIarticleServiceDeleteRequest**](ArticleServiceApi.md#APIarticleServiceDeleteRequest)|-|-|

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


<a id="APIarticleServiceDeleteRequest"></a>
## APIarticleServiceDeleteRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **deleteArticleRequest** | [**DeleteArticleRequest**](DeleteArticleRequest.md) |  | |



## articleServiceGet

> GetArticleReply articleServiceGet(articleServiceGetRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        GetArticleRequest getArticleRequest = new GetArticleRequest(); // GetArticleRequest | 
        try {
            APIarticleServiceGetRequest request = APIarticleServiceGetRequest.newBuilder()
                .getArticleRequest(getArticleRequest)
                .build();
            GetArticleReply result = apiInstance.articleServiceGet(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceGet");
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
| articleServiceGetRequest | [**APIarticleServiceGetRequest**](ArticleServiceApi.md#APIarticleServiceGetRequest)|-|-|

### Return type

[**GetArticleReply**](GetArticleReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## articleServiceGetWithHttpInfo

> ApiResponse<GetArticleReply> articleServiceGetWithHttpInfo(articleServiceGetRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        GetArticleRequest getArticleRequest = new GetArticleRequest(); // GetArticleRequest | 
        try {
            APIarticleServiceGetRequest request = APIarticleServiceGetRequest.newBuilder()
                .getArticleRequest(getArticleRequest)
                .build();
            ApiResponse<GetArticleReply> response = apiInstance.articleServiceGetWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceGet");
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
| articleServiceGetRequest | [**APIarticleServiceGetRequest**](ArticleServiceApi.md#APIarticleServiceGetRequest)|-|-|

### Return type

ApiResponse<[**GetArticleReply**](GetArticleReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIarticleServiceGetRequest"></a>
## APIarticleServiceGetRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getArticleRequest** | [**GetArticleRequest**](GetArticleRequest.md) |  | |



## articleServiceLike

> Object articleServiceLike(articleServiceLikeRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        LikeArticleRequest likeArticleRequest = new LikeArticleRequest(); // LikeArticleRequest | 
        try {
            APIarticleServiceLikeRequest request = APIarticleServiceLikeRequest.newBuilder()
                .likeArticleRequest(likeArticleRequest)
                .build();
            Object result = apiInstance.articleServiceLike(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceLike");
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
| articleServiceLikeRequest | [**APIarticleServiceLikeRequest**](ArticleServiceApi.md#APIarticleServiceLikeRequest)|-|-|

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

## articleServiceLikeWithHttpInfo

> ApiResponse<Object> articleServiceLikeWithHttpInfo(articleServiceLikeRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        LikeArticleRequest likeArticleRequest = new LikeArticleRequest(); // LikeArticleRequest | 
        try {
            APIarticleServiceLikeRequest request = APIarticleServiceLikeRequest.newBuilder()
                .likeArticleRequest(likeArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.articleServiceLikeWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceLike");
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
| articleServiceLikeRequest | [**APIarticleServiceLikeRequest**](ArticleServiceApi.md#APIarticleServiceLikeRequest)|-|-|

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


<a id="APIarticleServiceLikeRequest"></a>
## APIarticleServiceLikeRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **likeArticleRequest** | [**LikeArticleRequest**](LikeArticleRequest.md) |  | |



## articleServiceList

> ListArticlesReply articleServiceList(articleServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        ListArticlesRequest listArticlesRequest = new ListArticlesRequest(); // ListArticlesRequest | 
        try {
            APIarticleServiceListRequest request = APIarticleServiceListRequest.newBuilder()
                .listArticlesRequest(listArticlesRequest)
                .build();
            ListArticlesReply result = apiInstance.articleServiceList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceList");
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
| articleServiceListRequest | [**APIarticleServiceListRequest**](ArticleServiceApi.md#APIarticleServiceListRequest)|-|-|

### Return type

[**ListArticlesReply**](ListArticlesReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## articleServiceListWithHttpInfo

> ApiResponse<ListArticlesReply> articleServiceListWithHttpInfo(articleServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        ListArticlesRequest listArticlesRequest = new ListArticlesRequest(); // ListArticlesRequest | 
        try {
            APIarticleServiceListRequest request = APIarticleServiceListRequest.newBuilder()
                .listArticlesRequest(listArticlesRequest)
                .build();
            ApiResponse<ListArticlesReply> response = apiInstance.articleServiceListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceList");
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
| articleServiceListRequest | [**APIarticleServiceListRequest**](ArticleServiceApi.md#APIarticleServiceListRequest)|-|-|

### Return type

ApiResponse<[**ListArticlesReply**](ListArticlesReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIarticleServiceListRequest"></a>
## APIarticleServiceListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listArticlesRequest** | [**ListArticlesRequest**](ListArticlesRequest.md) |  | |



## articleServicePublish

> Object articleServicePublish(articleServicePublishRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        PublishArticleRequest publishArticleRequest = new PublishArticleRequest(); // PublishArticleRequest | 
        try {
            APIarticleServicePublishRequest request = APIarticleServicePublishRequest.newBuilder()
                .publishArticleRequest(publishArticleRequest)
                .build();
            Object result = apiInstance.articleServicePublish(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServicePublish");
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
| articleServicePublishRequest | [**APIarticleServicePublishRequest**](ArticleServiceApi.md#APIarticleServicePublishRequest)|-|-|

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

## articleServicePublishWithHttpInfo

> ApiResponse<Object> articleServicePublishWithHttpInfo(articleServicePublishRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        PublishArticleRequest publishArticleRequest = new PublishArticleRequest(); // PublishArticleRequest | 
        try {
            APIarticleServicePublishRequest request = APIarticleServicePublishRequest.newBuilder()
                .publishArticleRequest(publishArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.articleServicePublishWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServicePublish");
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
| articleServicePublishRequest | [**APIarticleServicePublishRequest**](ArticleServiceApi.md#APIarticleServicePublishRequest)|-|-|

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


<a id="APIarticleServicePublishRequest"></a>
## APIarticleServicePublishRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **publishArticleRequest** | [**PublishArticleRequest**](PublishArticleRequest.md) |  | |



## articleServiceReward

> Object articleServiceReward(articleServiceRewardRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        RewardArticleRequest rewardArticleRequest = new RewardArticleRequest(); // RewardArticleRequest | 
        try {
            APIarticleServiceRewardRequest request = APIarticleServiceRewardRequest.newBuilder()
                .rewardArticleRequest(rewardArticleRequest)
                .build();
            Object result = apiInstance.articleServiceReward(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceReward");
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
| articleServiceRewardRequest | [**APIarticleServiceRewardRequest**](ArticleServiceApi.md#APIarticleServiceRewardRequest)|-|-|

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

## articleServiceRewardWithHttpInfo

> ApiResponse<Object> articleServiceRewardWithHttpInfo(articleServiceRewardRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        RewardArticleRequest rewardArticleRequest = new RewardArticleRequest(); // RewardArticleRequest | 
        try {
            APIarticleServiceRewardRequest request = APIarticleServiceRewardRequest.newBuilder()
                .rewardArticleRequest(rewardArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.articleServiceRewardWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceReward");
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
| articleServiceRewardRequest | [**APIarticleServiceRewardRequest**](ArticleServiceApi.md#APIarticleServiceRewardRequest)|-|-|

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


<a id="APIarticleServiceRewardRequest"></a>
## APIarticleServiceRewardRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **rewardArticleRequest** | [**RewardArticleRequest**](RewardArticleRequest.md) |  | |



## articleServiceThank

> Object articleServiceThank(articleServiceThankRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        ThankArticleRequest thankArticleRequest = new ThankArticleRequest(); // ThankArticleRequest | 
        try {
            APIarticleServiceThankRequest request = APIarticleServiceThankRequest.newBuilder()
                .thankArticleRequest(thankArticleRequest)
                .build();
            Object result = apiInstance.articleServiceThank(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceThank");
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
| articleServiceThankRequest | [**APIarticleServiceThankRequest**](ArticleServiceApi.md#APIarticleServiceThankRequest)|-|-|

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

## articleServiceThankWithHttpInfo

> ApiResponse<Object> articleServiceThankWithHttpInfo(articleServiceThankRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        ThankArticleRequest thankArticleRequest = new ThankArticleRequest(); // ThankArticleRequest | 
        try {
            APIarticleServiceThankRequest request = APIarticleServiceThankRequest.newBuilder()
                .thankArticleRequest(thankArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.articleServiceThankWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceThank");
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
| articleServiceThankRequest | [**APIarticleServiceThankRequest**](ArticleServiceApi.md#APIarticleServiceThankRequest)|-|-|

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


<a id="APIarticleServiceThankRequest"></a>
## APIarticleServiceThankRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **thankArticleRequest** | [**ThankArticleRequest**](ThankArticleRequest.md) |  | |



## articleServiceUpdateDraft

> UpdateDraftArticleReply articleServiceUpdateDraft(articleServiceUpdateDraftRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        UpdateDraftArticleRequest updateDraftArticleRequest = new UpdateDraftArticleRequest(); // UpdateDraftArticleRequest | 
        try {
            APIarticleServiceUpdateDraftRequest request = APIarticleServiceUpdateDraftRequest.newBuilder()
                .updateDraftArticleRequest(updateDraftArticleRequest)
                .build();
            UpdateDraftArticleReply result = apiInstance.articleServiceUpdateDraft(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceUpdateDraft");
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
| articleServiceUpdateDraftRequest | [**APIarticleServiceUpdateDraftRequest**](ArticleServiceApi.md#APIarticleServiceUpdateDraftRequest)|-|-|

### Return type

[**UpdateDraftArticleReply**](UpdateDraftArticleReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## articleServiceUpdateDraftWithHttpInfo

> ApiResponse<UpdateDraftArticleReply> articleServiceUpdateDraftWithHttpInfo(articleServiceUpdateDraftRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        UpdateDraftArticleRequest updateDraftArticleRequest = new UpdateDraftArticleRequest(); // UpdateDraftArticleRequest | 
        try {
            APIarticleServiceUpdateDraftRequest request = APIarticleServiceUpdateDraftRequest.newBuilder()
                .updateDraftArticleRequest(updateDraftArticleRequest)
                .build();
            ApiResponse<UpdateDraftArticleReply> response = apiInstance.articleServiceUpdateDraftWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceUpdateDraft");
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
| articleServiceUpdateDraftRequest | [**APIarticleServiceUpdateDraftRequest**](ArticleServiceApi.md#APIarticleServiceUpdateDraftRequest)|-|-|

### Return type

ApiResponse<[**UpdateDraftArticleReply**](UpdateDraftArticleReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIarticleServiceUpdateDraftRequest"></a>
## APIarticleServiceUpdateDraftRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateDraftArticleRequest** | [**UpdateDraftArticleRequest**](UpdateDraftArticleRequest.md) |  | |



## articleServiceWatch

> Object articleServiceWatch(articleServiceWatchRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        WatchArticleRequest watchArticleRequest = new WatchArticleRequest(); // WatchArticleRequest | 
        try {
            APIarticleServiceWatchRequest request = APIarticleServiceWatchRequest.newBuilder()
                .watchArticleRequest(watchArticleRequest)
                .build();
            Object result = apiInstance.articleServiceWatch(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceWatch");
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
| articleServiceWatchRequest | [**APIarticleServiceWatchRequest**](ArticleServiceApi.md#APIarticleServiceWatchRequest)|-|-|

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

## articleServiceWatchWithHttpInfo

> ApiResponse<Object> articleServiceWatchWithHttpInfo(articleServiceWatchRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleServiceApi;
import com.bass.bbs.api.ArticleServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleServiceApi apiInstance = new ArticleServiceApi(defaultClient);
        WatchArticleRequest watchArticleRequest = new WatchArticleRequest(); // WatchArticleRequest | 
        try {
            APIarticleServiceWatchRequest request = APIarticleServiceWatchRequest.newBuilder()
                .watchArticleRequest(watchArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.articleServiceWatchWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleServiceApi#articleServiceWatch");
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
| articleServiceWatchRequest | [**APIarticleServiceWatchRequest**](ArticleServiceApi.md#APIarticleServiceWatchRequest)|-|-|

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


<a id="APIarticleServiceWatchRequest"></a>
## APIarticleServiceWatchRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **watchArticleRequest** | [**WatchArticleRequest**](WatchArticleRequest.md) |  | |


