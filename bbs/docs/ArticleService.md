# ArticleService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**archive**](ArticleService.md#archive) | **POST** /v1/content/article/archive |  |
| [**archiveWithHttpInfo**](ArticleService.md#archiveWithHttpInfo) | **POST** /v1/content/article/archive |  |
| [**callList**](ArticleService.md#callList) | **POST** /v1/content/article/list |  |
| [**callListWithHttpInfo**](ArticleService.md#callListWithHttpInfo) | **POST** /v1/content/article/list |  |
| [**cancelPublish**](ArticleService.md#cancelPublish) | **POST** /v1/content/article/publish/cancel |  |
| [**cancelPublishWithHttpInfo**](ArticleService.md#cancelPublishWithHttpInfo) | **POST** /v1/content/article/publish/cancel |  |
| [**collect**](ArticleService.md#collect) | **POST** /v1/content/article/collect |  |
| [**collectWithHttpInfo**](ArticleService.md#collectWithHttpInfo) | **POST** /v1/content/article/collect |  |
| [**createDraft**](ArticleService.md#createDraft) | **POST** /v1/content/article/draft/create |  |
| [**createDraftWithHttpInfo**](ArticleService.md#createDraftWithHttpInfo) | **POST** /v1/content/article/draft/create |  |
| [**discardDraft**](ArticleService.md#discardDraft) | **POST** /v1/content/article/draft/discard |  |
| [**discardDraftWithHttpInfo**](ArticleService.md#discardDraftWithHttpInfo) | **POST** /v1/content/article/draft/discard |  |
| [**get**](ArticleService.md#get) | **POST** /v1/content/article/get |  |
| [**getWithHttpInfo**](ArticleService.md#getWithHttpInfo) | **POST** /v1/content/article/get |  |
| [**like**](ArticleService.md#like) | **POST** /v1/content/article/like |  |
| [**likeWithHttpInfo**](ArticleService.md#likeWithHttpInfo) | **POST** /v1/content/article/like |  |
| [**publish**](ArticleService.md#publish) | **POST** /v1/content/article/publish |  |
| [**publishWithHttpInfo**](ArticleService.md#publishWithHttpInfo) | **POST** /v1/content/article/publish |  |
| [**reward**](ArticleService.md#reward) | **POST** /v1/content/article/reward |  |
| [**rewardWithHttpInfo**](ArticleService.md#rewardWithHttpInfo) | **POST** /v1/content/article/reward |  |
| [**schedulePublish**](ArticleService.md#schedulePublish) | **POST** /v1/content/article/publish/schedule |  |
| [**schedulePublishWithHttpInfo**](ArticleService.md#schedulePublishWithHttpInfo) | **POST** /v1/content/article/publish/schedule |  |
| [**thank**](ArticleService.md#thank) | **POST** /v1/content/article/thank |  |
| [**thankWithHttpInfo**](ArticleService.md#thankWithHttpInfo) | **POST** /v1/content/article/thank |  |
| [**updateDraft**](ArticleService.md#updateDraft) | **POST** /v1/content/article/draft/update |  |
| [**updateDraftWithHttpInfo**](ArticleService.md#updateDraftWithHttpInfo) | **POST** /v1/content/article/draft/update |  |



## archive

> Object archive(archiveRequest)



归档文章

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
        ArchiveArticleReq archiveArticleReq = new ArchiveArticleReq(); // ArchiveArticleReq | 
        try {
            APIarchiveRequest request = APIarchiveRequest.newBuilder()
                .archiveArticleReq(archiveArticleReq)
                .build();
            Object result = apiInstance.archive(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#archive");
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
| archiveRequest | [**APIarchiveRequest**](ArticleService.md#APIarchiveRequest)|-|-|

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

## archiveWithHttpInfo

> ApiResponse<Object> archiveWithHttpInfo(archiveRequest)



归档文章

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
        ArchiveArticleReq archiveArticleReq = new ArchiveArticleReq(); // ArchiveArticleReq | 
        try {
            APIarchiveRequest request = APIarchiveRequest.newBuilder()
                .archiveArticleReq(archiveArticleReq)
                .build();
            ApiResponse<Object> response = apiInstance.archiveWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#archive");
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
| archiveRequest | [**APIarchiveRequest**](ArticleService.md#APIarchiveRequest)|-|-|

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


<a id="APIarchiveRequest"></a>
## APIarchiveRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **archiveArticleReq** | [**ArchiveArticleReq**](ArchiveArticleReq.md) |  | |



## callList

> ListArticlesResp callList(callListRequest)



查询文章列表

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
        ListArticlesReq listArticlesReq = new ListArticlesReq(); // ListArticlesReq | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listArticlesReq(listArticlesReq)
                .build();
            ListArticlesResp result = apiInstance.callList(request);
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

[**ListArticlesResp**](ListArticlesResp.md)


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

> ApiResponse<ListArticlesResp> callListWithHttpInfo(callListRequest)



查询文章列表

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
        ListArticlesReq listArticlesReq = new ListArticlesReq(); // ListArticlesReq | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listArticlesReq(listArticlesReq)
                .build();
            ApiResponse<ListArticlesResp> response = apiInstance.callListWithHttpInfo(request);
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

ApiResponse<[**ListArticlesResp**](ListArticlesResp.md)>


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
| **listArticlesReq** | [**ListArticlesReq**](ListArticlesReq.md) |  | |



## cancelPublish

> Object cancelPublish(cancelPublishRequest)



取消定时发布

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
        CancelPublishArticleReq cancelPublishArticleReq = new CancelPublishArticleReq(); // CancelPublishArticleReq | 
        try {
            APIcancelPublishRequest request = APIcancelPublishRequest.newBuilder()
                .cancelPublishArticleReq(cancelPublishArticleReq)
                .build();
            Object result = apiInstance.cancelPublish(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#cancelPublish");
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
| cancelPublishRequest | [**APIcancelPublishRequest**](ArticleService.md#APIcancelPublishRequest)|-|-|

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

## cancelPublishWithHttpInfo

> ApiResponse<Object> cancelPublishWithHttpInfo(cancelPublishRequest)



取消定时发布

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
        CancelPublishArticleReq cancelPublishArticleReq = new CancelPublishArticleReq(); // CancelPublishArticleReq | 
        try {
            APIcancelPublishRequest request = APIcancelPublishRequest.newBuilder()
                .cancelPublishArticleReq(cancelPublishArticleReq)
                .build();
            ApiResponse<Object> response = apiInstance.cancelPublishWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#cancelPublish");
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
| cancelPublishRequest | [**APIcancelPublishRequest**](ArticleService.md#APIcancelPublishRequest)|-|-|

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


<a id="APIcancelPublishRequest"></a>
## APIcancelPublishRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **cancelPublishArticleReq** | [**CancelPublishArticleReq**](CancelPublishArticleReq.md) |  | |



## collect

> CollectArticleResp collect(collectRequest)



收藏文章

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
        CollectArticleReq collectArticleReq = new CollectArticleReq(); // CollectArticleReq | 
        try {
            APIcollectRequest request = APIcollectRequest.newBuilder()
                .collectArticleReq(collectArticleReq)
                .build();
            CollectArticleResp result = apiInstance.collect(request);
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

[**CollectArticleResp**](CollectArticleResp.md)


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

> ApiResponse<CollectArticleResp> collectWithHttpInfo(collectRequest)



收藏文章

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
        CollectArticleReq collectArticleReq = new CollectArticleReq(); // CollectArticleReq | 
        try {
            APIcollectRequest request = APIcollectRequest.newBuilder()
                .collectArticleReq(collectArticleReq)
                .build();
            ApiResponse<CollectArticleResp> response = apiInstance.collectWithHttpInfo(request);
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

ApiResponse<[**CollectArticleResp**](CollectArticleResp.md)>


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
| **collectArticleReq** | [**CollectArticleReq**](CollectArticleReq.md) |  | |



## createDraft

> CreateDraftArticleResp createDraft(createDraftRequest)



创建文章草稿

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
        CreateDraftArticleReq createDraftArticleReq = new CreateDraftArticleReq(); // CreateDraftArticleReq | 
        try {
            APIcreateDraftRequest request = APIcreateDraftRequest.newBuilder()
                .createDraftArticleReq(createDraftArticleReq)
                .build();
            CreateDraftArticleResp result = apiInstance.createDraft(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#createDraft");
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
| createDraftRequest | [**APIcreateDraftRequest**](ArticleService.md#APIcreateDraftRequest)|-|-|

### Return type

[**CreateDraftArticleResp**](CreateDraftArticleResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## createDraftWithHttpInfo

> ApiResponse<CreateDraftArticleResp> createDraftWithHttpInfo(createDraftRequest)



创建文章草稿

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
        CreateDraftArticleReq createDraftArticleReq = new CreateDraftArticleReq(); // CreateDraftArticleReq | 
        try {
            APIcreateDraftRequest request = APIcreateDraftRequest.newBuilder()
                .createDraftArticleReq(createDraftArticleReq)
                .build();
            ApiResponse<CreateDraftArticleResp> response = apiInstance.createDraftWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#createDraft");
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
| createDraftRequest | [**APIcreateDraftRequest**](ArticleService.md#APIcreateDraftRequest)|-|-|

### Return type

ApiResponse<[**CreateDraftArticleResp**](CreateDraftArticleResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcreateDraftRequest"></a>
## APIcreateDraftRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createDraftArticleReq** | [**CreateDraftArticleReq**](CreateDraftArticleReq.md) |  | |



## discardDraft

> Object discardDraft(discardDraftRequest)



丢弃文章草稿

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
        DiscardDraftArticleReq discardDraftArticleReq = new DiscardDraftArticleReq(); // DiscardDraftArticleReq | 
        try {
            APIdiscardDraftRequest request = APIdiscardDraftRequest.newBuilder()
                .discardDraftArticleReq(discardDraftArticleReq)
                .build();
            Object result = apiInstance.discardDraft(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#discardDraft");
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
| discardDraftRequest | [**APIdiscardDraftRequest**](ArticleService.md#APIdiscardDraftRequest)|-|-|

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

## discardDraftWithHttpInfo

> ApiResponse<Object> discardDraftWithHttpInfo(discardDraftRequest)



丢弃文章草稿

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
        DiscardDraftArticleReq discardDraftArticleReq = new DiscardDraftArticleReq(); // DiscardDraftArticleReq | 
        try {
            APIdiscardDraftRequest request = APIdiscardDraftRequest.newBuilder()
                .discardDraftArticleReq(discardDraftArticleReq)
                .build();
            ApiResponse<Object> response = apiInstance.discardDraftWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#discardDraft");
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
| discardDraftRequest | [**APIdiscardDraftRequest**](ArticleService.md#APIdiscardDraftRequest)|-|-|

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


<a id="APIdiscardDraftRequest"></a>
## APIdiscardDraftRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **discardDraftArticleReq** | [**DiscardDraftArticleReq**](DiscardDraftArticleReq.md) |  | |



## get

> GetArticleResp get(getRequest)



查询文章详情

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
        GetArticleReq getArticleReq = new GetArticleReq(); // GetArticleReq | 
        try {
            APIgetRequest request = APIgetRequest.newBuilder()
                .getArticleReq(getArticleReq)
                .build();
            GetArticleResp result = apiInstance.get(request);
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

[**GetArticleResp**](GetArticleResp.md)


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

> ApiResponse<GetArticleResp> getWithHttpInfo(getRequest)



查询文章详情

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
        GetArticleReq getArticleReq = new GetArticleReq(); // GetArticleReq | 
        try {
            APIgetRequest request = APIgetRequest.newBuilder()
                .getArticleReq(getArticleReq)
                .build();
            ApiResponse<GetArticleResp> response = apiInstance.getWithHttpInfo(request);
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

ApiResponse<[**GetArticleResp**](GetArticleResp.md)>


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
| **getArticleReq** | [**GetArticleReq**](GetArticleReq.md) |  | |



## like

> LikeArticleResp like(likeRequest)



点赞文章

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
        LikeArticleReq likeArticleReq = new LikeArticleReq(); // LikeArticleReq | 
        try {
            APIlikeRequest request = APIlikeRequest.newBuilder()
                .likeArticleReq(likeArticleReq)
                .build();
            LikeArticleResp result = apiInstance.like(request);
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

[**LikeArticleResp**](LikeArticleResp.md)


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

> ApiResponse<LikeArticleResp> likeWithHttpInfo(likeRequest)



点赞文章

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
        LikeArticleReq likeArticleReq = new LikeArticleReq(); // LikeArticleReq | 
        try {
            APIlikeRequest request = APIlikeRequest.newBuilder()
                .likeArticleReq(likeArticleReq)
                .build();
            ApiResponse<LikeArticleResp> response = apiInstance.likeWithHttpInfo(request);
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

ApiResponse<[**LikeArticleResp**](LikeArticleResp.md)>


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
| **likeArticleReq** | [**LikeArticleReq**](LikeArticleReq.md) |  | |



## publish

> Object publish(publishRequest)



发布文章

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
        PublishArticleReq publishArticleReq = new PublishArticleReq(); // PublishArticleReq | 
        try {
            APIpublishRequest request = APIpublishRequest.newBuilder()
                .publishArticleReq(publishArticleReq)
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



发布文章

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
        PublishArticleReq publishArticleReq = new PublishArticleReq(); // PublishArticleReq | 
        try {
            APIpublishRequest request = APIpublishRequest.newBuilder()
                .publishArticleReq(publishArticleReq)
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
| **publishArticleReq** | [**PublishArticleReq**](PublishArticleReq.md) |  | |



## reward

> Object reward(rewardRequest)



打赏文章

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
        RewardArticleReq rewardArticleReq = new RewardArticleReq(); // RewardArticleReq | 
        try {
            APIrewardRequest request = APIrewardRequest.newBuilder()
                .rewardArticleReq(rewardArticleReq)
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



打赏文章

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
        RewardArticleReq rewardArticleReq = new RewardArticleReq(); // RewardArticleReq | 
        try {
            APIrewardRequest request = APIrewardRequest.newBuilder()
                .rewardArticleReq(rewardArticleReq)
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
| **rewardArticleReq** | [**RewardArticleReq**](RewardArticleReq.md) |  | |



## schedulePublish

> Object schedulePublish(schedulePublishRequest)



设置定时发布

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
        SchedulePublishArticleReq schedulePublishArticleReq = new SchedulePublishArticleReq(); // SchedulePublishArticleReq | 
        try {
            APIschedulePublishRequest request = APIschedulePublishRequest.newBuilder()
                .schedulePublishArticleReq(schedulePublishArticleReq)
                .build();
            Object result = apiInstance.schedulePublish(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#schedulePublish");
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
| schedulePublishRequest | [**APIschedulePublishRequest**](ArticleService.md#APIschedulePublishRequest)|-|-|

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

## schedulePublishWithHttpInfo

> ApiResponse<Object> schedulePublishWithHttpInfo(schedulePublishRequest)



设置定时发布

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
        SchedulePublishArticleReq schedulePublishArticleReq = new SchedulePublishArticleReq(); // SchedulePublishArticleReq | 
        try {
            APIschedulePublishRequest request = APIschedulePublishRequest.newBuilder()
                .schedulePublishArticleReq(schedulePublishArticleReq)
                .build();
            ApiResponse<Object> response = apiInstance.schedulePublishWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#schedulePublish");
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
| schedulePublishRequest | [**APIschedulePublishRequest**](ArticleService.md#APIschedulePublishRequest)|-|-|

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


<a id="APIschedulePublishRequest"></a>
## APIschedulePublishRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **schedulePublishArticleReq** | [**SchedulePublishArticleReq**](SchedulePublishArticleReq.md) |  | |



## thank

> ThankArticleResp thank(thankRequest)



感谢文章

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
        ThankArticleReq thankArticleReq = new ThankArticleReq(); // ThankArticleReq | 
        try {
            APIthankRequest request = APIthankRequest.newBuilder()
                .thankArticleReq(thankArticleReq)
                .build();
            ThankArticleResp result = apiInstance.thank(request);
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

[**ThankArticleResp**](ThankArticleResp.md)


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

> ApiResponse<ThankArticleResp> thankWithHttpInfo(thankRequest)



感谢文章

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
        ThankArticleReq thankArticleReq = new ThankArticleReq(); // ThankArticleReq | 
        try {
            APIthankRequest request = APIthankRequest.newBuilder()
                .thankArticleReq(thankArticleReq)
                .build();
            ApiResponse<ThankArticleResp> response = apiInstance.thankWithHttpInfo(request);
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

ApiResponse<[**ThankArticleResp**](ThankArticleResp.md)>


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
| **thankArticleReq** | [**ThankArticleReq**](ThankArticleReq.md) |  | |



## updateDraft

> UpdateDraftArticleResp updateDraft(updateDraftRequest)



编辑文章草稿

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
        UpdateDraftArticleReq updateDraftArticleReq = new UpdateDraftArticleReq(); // UpdateDraftArticleReq | 
        try {
            APIupdateDraftRequest request = APIupdateDraftRequest.newBuilder()
                .updateDraftArticleReq(updateDraftArticleReq)
                .build();
            UpdateDraftArticleResp result = apiInstance.updateDraft(request);
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

[**UpdateDraftArticleResp**](UpdateDraftArticleResp.md)


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

> ApiResponse<UpdateDraftArticleResp> updateDraftWithHttpInfo(updateDraftRequest)



编辑文章草稿

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
        UpdateDraftArticleReq updateDraftArticleReq = new UpdateDraftArticleReq(); // UpdateDraftArticleReq | 
        try {
            APIupdateDraftRequest request = APIupdateDraftRequest.newBuilder()
                .updateDraftArticleReq(updateDraftArticleReq)
                .build();
            ApiResponse<UpdateDraftArticleResp> response = apiInstance.updateDraftWithHttpInfo(request);
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

ApiResponse<[**UpdateDraftArticleResp**](UpdateDraftArticleResp.md)>


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
| **updateDraftArticleReq** | [**UpdateDraftArticleReq**](UpdateDraftArticleReq.md) |  | |


