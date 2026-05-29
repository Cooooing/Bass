# ArticleService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**acceptAnswer**](ArticleService.md#acceptAnswer) | **POST** /v1/content/article/accept-answer |  |
| [**acceptAnswerWithHttpInfo**](ArticleService.md#acceptAnswerWithHttpInfo) | **POST** /v1/content/article/accept-answer |  |
| [**callList**](ArticleService.md#callList) | **POST** /v1/content/article/list |  |
| [**callListWithHttpInfo**](ArticleService.md#callListWithHttpInfo) | **POST** /v1/content/article/list |  |
| [**collect**](ArticleService.md#collect) | **POST** /v1/content/article/collect |  |
| [**collectWithHttpInfo**](ArticleService.md#collectWithHttpInfo) | **POST** /v1/content/article/collect |  |
| [**create**](ArticleService.md#create) | **POST** /v1/content/article/create |  |
| [**createWithHttpInfo**](ArticleService.md#createWithHttpInfo) | **POST** /v1/content/article/create |  |
| [**delete**](ArticleService.md#delete) | **POST** /v1/content/article/delete |  |
| [**deleteWithHttpInfo**](ArticleService.md#deleteWithHttpInfo) | **POST** /v1/content/article/delete |  |
| [**get**](ArticleService.md#get) | **POST** /v1/content/article/get |  |
| [**getWithHttpInfo**](ArticleService.md#getWithHttpInfo) | **POST** /v1/content/article/get |  |
| [**like**](ArticleService.md#like) | **POST** /v1/content/article/like |  |
| [**likeWithHttpInfo**](ArticleService.md#likeWithHttpInfo) | **POST** /v1/content/article/like |  |
| [**publish**](ArticleService.md#publish) | **POST** /v1/content/article/publish |  |
| [**publishWithHttpInfo**](ArticleService.md#publishWithHttpInfo) | **POST** /v1/content/article/publish |  |
| [**reward**](ArticleService.md#reward) | **POST** /v1/content/article/reward |  |
| [**rewardWithHttpInfo**](ArticleService.md#rewardWithHttpInfo) | **POST** /v1/content/article/reward |  |
| [**thank**](ArticleService.md#thank) | **POST** /v1/content/article/thank |  |
| [**thankWithHttpInfo**](ArticleService.md#thankWithHttpInfo) | **POST** /v1/content/article/thank |  |
| [**updateDraft**](ArticleService.md#updateDraft) | **POST** /v1/content/article/update-draft |  |
| [**updateDraftWithHttpInfo**](ArticleService.md#updateDraftWithHttpInfo) | **POST** /v1/content/article/update-draft |  |
| [**watch**](ArticleService.md#watch) | **POST** /v1/content/article/watch |  |
| [**watchWithHttpInfo**](ArticleService.md#watchWithHttpInfo) | **POST** /v1/content/article/watch |  |



## acceptAnswer

> Object acceptAnswer(acceptAnswerRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        AcceptAnswerArticleRequest acceptAnswerArticleRequest = new AcceptAnswerArticleRequest(); // AcceptAnswerArticleRequest | 
        try {
            APIacceptAnswerRequest request = APIacceptAnswerRequest.newBuilder()
                .acceptAnswerArticleRequest(acceptAnswerArticleRequest)
                .build();
            Object result = apiInstance.acceptAnswer(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#acceptAnswer");
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
| acceptAnswerRequest | [**APIacceptAnswerRequest**](ArticleService.md#APIacceptAnswerRequest)|-|-|

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

## acceptAnswerWithHttpInfo

> ApiResponse<Object> acceptAnswerWithHttpInfo(acceptAnswerRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        AcceptAnswerArticleRequest acceptAnswerArticleRequest = new AcceptAnswerArticleRequest(); // AcceptAnswerArticleRequest | 
        try {
            APIacceptAnswerRequest request = APIacceptAnswerRequest.newBuilder()
                .acceptAnswerArticleRequest(acceptAnswerArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.acceptAnswerWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#acceptAnswer");
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
| acceptAnswerRequest | [**APIacceptAnswerRequest**](ArticleService.md#APIacceptAnswerRequest)|-|-|

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


<a id="APIacceptAnswerRequest"></a>
## APIacceptAnswerRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **acceptAnswerArticleRequest** | [**AcceptAnswerArticleRequest**](AcceptAnswerArticleRequest.md) |  | |



## callList

> ListArticlesReply callList(callListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        ListArticlesRequest listArticlesRequest = new ListArticlesRequest(); // ListArticlesRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listArticlesRequest(listArticlesRequest)
                .build();
            ListArticlesReply result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#callList");
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
| callListRequest | [**APIcallListRequest**](ArticleService.md#APIcallListRequest)|-|-|

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

## callListWithHttpInfo

> ApiResponse<ListArticlesReply> callListWithHttpInfo(callListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        ListArticlesRequest listArticlesRequest = new ListArticlesRequest(); // ListArticlesRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listArticlesRequest(listArticlesRequest)
                .build();
            ApiResponse<ListArticlesReply> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#callList");
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
| callListRequest | [**APIcallListRequest**](ArticleService.md#APIcallListRequest)|-|-|

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


<a id="APIcallListRequest"></a>
## APIcallListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listArticlesRequest** | [**ListArticlesRequest**](ListArticlesRequest.md) |  | |



## collect

> Object collect(collectRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        CollectArticleRequest collectArticleRequest = new CollectArticleRequest(); // CollectArticleRequest | 
        try {
            APIcollectRequest request = APIcollectRequest.newBuilder()
                .collectArticleRequest(collectArticleRequest)
                .build();
            Object result = apiInstance.collect(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#collect");
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
| collectRequest | [**APIcollectRequest**](ArticleService.md#APIcollectRequest)|-|-|

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

## collectWithHttpInfo

> ApiResponse<Object> collectWithHttpInfo(collectRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        CollectArticleRequest collectArticleRequest = new CollectArticleRequest(); // CollectArticleRequest | 
        try {
            APIcollectRequest request = APIcollectRequest.newBuilder()
                .collectArticleRequest(collectArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.collectWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#collect");
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
| collectRequest | [**APIcollectRequest**](ArticleService.md#APIcollectRequest)|-|-|

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


<a id="APIcollectRequest"></a>
## APIcollectRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **collectArticleRequest** | [**CollectArticleRequest**](CollectArticleRequest.md) |  | |



## create

> CreateArticleReply create(createRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        CreateArticleRequest createArticleRequest = new CreateArticleRequest(); // CreateArticleRequest | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createArticleRequest(createArticleRequest)
                .build();
            CreateArticleReply result = apiInstance.create(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#create");
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
| createRequest | [**APIcreateRequest**](ArticleService.md#APIcreateRequest)|-|-|

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

## createWithHttpInfo

> ApiResponse<CreateArticleReply> createWithHttpInfo(createRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        CreateArticleRequest createArticleRequest = new CreateArticleRequest(); // CreateArticleRequest | 
        try {
            APIcreateRequest request = APIcreateRequest.newBuilder()
                .createArticleRequest(createArticleRequest)
                .build();
            ApiResponse<CreateArticleReply> response = apiInstance.createWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#create");
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
| createRequest | [**APIcreateRequest**](ArticleService.md#APIcreateRequest)|-|-|

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


<a id="APIcreateRequest"></a>
## APIcreateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createArticleRequest** | [**CreateArticleRequest**](CreateArticleRequest.md) |  | |



## delete

> Object delete(deleteRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        DeleteArticleRequest deleteArticleRequest = new DeleteArticleRequest(); // DeleteArticleRequest | 
        try {
            APIdeleteRequest request = APIdeleteRequest.newBuilder()
                .deleteArticleRequest(deleteArticleRequest)
                .build();
            Object result = apiInstance.delete(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#delete");
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
| deleteRequest | [**APIdeleteRequest**](ArticleService.md#APIdeleteRequest)|-|-|

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

## deleteWithHttpInfo

> ApiResponse<Object> deleteWithHttpInfo(deleteRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        DeleteArticleRequest deleteArticleRequest = new DeleteArticleRequest(); // DeleteArticleRequest | 
        try {
            APIdeleteRequest request = APIdeleteRequest.newBuilder()
                .deleteArticleRequest(deleteArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.deleteWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#delete");
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
| deleteRequest | [**APIdeleteRequest**](ArticleService.md#APIdeleteRequest)|-|-|

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


<a id="APIdeleteRequest"></a>
## APIdeleteRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **deleteArticleRequest** | [**DeleteArticleRequest**](DeleteArticleRequest.md) |  | |



## get

> GetArticleReply get(getRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        GetArticleRequest getArticleRequest = new GetArticleRequest(); // GetArticleRequest | 
        try {
            APIgetRequest request = APIgetRequest.newBuilder()
                .getArticleRequest(getArticleRequest)
                .build();
            GetArticleReply result = apiInstance.get(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#get");
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
| getRequest | [**APIgetRequest**](ArticleService.md#APIgetRequest)|-|-|

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

## getWithHttpInfo

> ApiResponse<GetArticleReply> getWithHttpInfo(getRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        GetArticleRequest getArticleRequest = new GetArticleRequest(); // GetArticleRequest | 
        try {
            APIgetRequest request = APIgetRequest.newBuilder()
                .getArticleRequest(getArticleRequest)
                .build();
            ApiResponse<GetArticleReply> response = apiInstance.getWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#get");
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
| getRequest | [**APIgetRequest**](ArticleService.md#APIgetRequest)|-|-|

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


<a id="APIgetRequest"></a>
## APIgetRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getArticleRequest** | [**GetArticleRequest**](GetArticleRequest.md) |  | |



## like

> Object like(likeRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        LikeArticleRequest likeArticleRequest = new LikeArticleRequest(); // LikeArticleRequest | 
        try {
            APIlikeRequest request = APIlikeRequest.newBuilder()
                .likeArticleRequest(likeArticleRequest)
                .build();
            Object result = apiInstance.like(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#like");
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
| likeRequest | [**APIlikeRequest**](ArticleService.md#APIlikeRequest)|-|-|

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
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        LikeArticleRequest likeArticleRequest = new LikeArticleRequest(); // LikeArticleRequest | 
        try {
            APIlikeRequest request = APIlikeRequest.newBuilder()
                .likeArticleRequest(likeArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.likeWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#like");
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
| likeRequest | [**APIlikeRequest**](ArticleService.md#APIlikeRequest)|-|-|

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
| **likeArticleRequest** | [**LikeArticleRequest**](LikeArticleRequest.md) |  | |



## publish

> Object publish(publishRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        PublishArticleRequest publishArticleRequest = new PublishArticleRequest(); // PublishArticleRequest | 
        try {
            APIpublishRequest request = APIpublishRequest.newBuilder()
                .publishArticleRequest(publishArticleRequest)
                .build();
            Object result = apiInstance.publish(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#publish");
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
| publishRequest | [**APIpublishRequest**](ArticleService.md#APIpublishRequest)|-|-|

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

## publishWithHttpInfo

> ApiResponse<Object> publishWithHttpInfo(publishRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        PublishArticleRequest publishArticleRequest = new PublishArticleRequest(); // PublishArticleRequest | 
        try {
            APIpublishRequest request = APIpublishRequest.newBuilder()
                .publishArticleRequest(publishArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.publishWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#publish");
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
| publishRequest | [**APIpublishRequest**](ArticleService.md#APIpublishRequest)|-|-|

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


<a id="APIpublishRequest"></a>
## APIpublishRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **publishArticleRequest** | [**PublishArticleRequest**](PublishArticleRequest.md) |  | |



## reward

> Object reward(rewardRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        RewardArticleRequest rewardArticleRequest = new RewardArticleRequest(); // RewardArticleRequest | 
        try {
            APIrewardRequest request = APIrewardRequest.newBuilder()
                .rewardArticleRequest(rewardArticleRequest)
                .build();
            Object result = apiInstance.reward(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#reward");
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
| rewardRequest | [**APIrewardRequest**](ArticleService.md#APIrewardRequest)|-|-|

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

## rewardWithHttpInfo

> ApiResponse<Object> rewardWithHttpInfo(rewardRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        RewardArticleRequest rewardArticleRequest = new RewardArticleRequest(); // RewardArticleRequest | 
        try {
            APIrewardRequest request = APIrewardRequest.newBuilder()
                .rewardArticleRequest(rewardArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.rewardWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#reward");
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
| rewardRequest | [**APIrewardRequest**](ArticleService.md#APIrewardRequest)|-|-|

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


<a id="APIrewardRequest"></a>
## APIrewardRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **rewardArticleRequest** | [**RewardArticleRequest**](RewardArticleRequest.md) |  | |



## thank

> Object thank(thankRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        ThankArticleRequest thankArticleRequest = new ThankArticleRequest(); // ThankArticleRequest | 
        try {
            APIthankRequest request = APIthankRequest.newBuilder()
                .thankArticleRequest(thankArticleRequest)
                .build();
            Object result = apiInstance.thank(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#thank");
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
| thankRequest | [**APIthankRequest**](ArticleService.md#APIthankRequest)|-|-|

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
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        ThankArticleRequest thankArticleRequest = new ThankArticleRequest(); // ThankArticleRequest | 
        try {
            APIthankRequest request = APIthankRequest.newBuilder()
                .thankArticleRequest(thankArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.thankWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#thank");
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
| thankRequest | [**APIthankRequest**](ArticleService.md#APIthankRequest)|-|-|

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
| **thankArticleRequest** | [**ThankArticleRequest**](ThankArticleRequest.md) |  | |



## updateDraft

> UpdateDraftArticleReply updateDraft(updateDraftRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        UpdateDraftArticleRequest updateDraftArticleRequest = new UpdateDraftArticleRequest(); // UpdateDraftArticleRequest | 
        try {
            APIupdateDraftRequest request = APIupdateDraftRequest.newBuilder()
                .updateDraftArticleRequest(updateDraftArticleRequest)
                .build();
            UpdateDraftArticleReply result = apiInstance.updateDraft(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#updateDraft");
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
| updateDraftRequest | [**APIupdateDraftRequest**](ArticleService.md#APIupdateDraftRequest)|-|-|

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

## updateDraftWithHttpInfo

> ApiResponse<UpdateDraftArticleReply> updateDraftWithHttpInfo(updateDraftRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        UpdateDraftArticleRequest updateDraftArticleRequest = new UpdateDraftArticleRequest(); // UpdateDraftArticleRequest | 
        try {
            APIupdateDraftRequest request = APIupdateDraftRequest.newBuilder()
                .updateDraftArticleRequest(updateDraftArticleRequest)
                .build();
            ApiResponse<UpdateDraftArticleReply> response = apiInstance.updateDraftWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#updateDraft");
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
| updateDraftRequest | [**APIupdateDraftRequest**](ArticleService.md#APIupdateDraftRequest)|-|-|

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


<a id="APIupdateDraftRequest"></a>
## APIupdateDraftRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateDraftArticleRequest** | [**UpdateDraftArticleRequest**](UpdateDraftArticleRequest.md) |  | |



## watch

> Object watch(watchRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        WatchArticleRequest watchArticleRequest = new WatchArticleRequest(); // WatchArticleRequest | 
        try {
            APIwatchRequest request = APIwatchRequest.newBuilder()
                .watchArticleRequest(watchArticleRequest)
                .build();
            Object result = apiInstance.watch(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#watch");
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
| watchRequest | [**APIwatchRequest**](ArticleService.md#APIwatchRequest)|-|-|

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

## watchWithHttpInfo

> ApiResponse<Object> watchWithHttpInfo(watchRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleService;
import com.bass.bbs.api.ArticleService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleService apiInstance = new ArticleService(defaultClient);
        WatchArticleRequest watchArticleRequest = new WatchArticleRequest(); // WatchArticleRequest | 
        try {
            APIwatchRequest request = APIwatchRequest.newBuilder()
                .watchArticleRequest(watchArticleRequest)
                .build();
            ApiResponse<Object> response = apiInstance.watchWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#watch");
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
| watchRequest | [**APIwatchRequest**](ArticleService.md#APIwatchRequest)|-|-|

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


<a id="APIwatchRequest"></a>
## APIwatchRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **watchArticleRequest** | [**WatchArticleRequest**](WatchArticleRequest.md) |  | |


