# ArticleService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**acceptArticleAnswer**](ArticleService.md#acceptArticleAnswer) | **POST** /v1/content/article/answer/accept |  |
| [**acceptArticleAnswerWithHttpInfo**](ArticleService.md#acceptArticleAnswerWithHttpInfo) | **POST** /v1/content/article/answer/accept |  |
| [**archiveArticle**](ArticleService.md#archiveArticle) | **POST** /v1/content/article/archive |  |
| [**archiveArticleWithHttpInfo**](ArticleService.md#archiveArticleWithHttpInfo) | **POST** /v1/content/article/archive |  |
| [**cancelPublishArticle**](ArticleService.md#cancelPublishArticle) | **POST** /v1/content/article/publish/cancel |  |
| [**cancelPublishArticleWithHttpInfo**](ArticleService.md#cancelPublishArticleWithHttpInfo) | **POST** /v1/content/article/publish/cancel |  |
| [**collectArticle**](ArticleService.md#collectArticle) | **POST** /v1/content/article/collect |  |
| [**collectArticleWithHttpInfo**](ArticleService.md#collectArticleWithHttpInfo) | **POST** /v1/content/article/collect |  |
| [**countCollectedArticles**](ArticleService.md#countCollectedArticles) | **POST** /v1/content/article/collect/count |  |
| [**countCollectedArticlesWithHttpInfo**](ArticleService.md#countCollectedArticlesWithHttpInfo) | **POST** /v1/content/article/collect/count |  |
| [**createDraftArticle**](ArticleService.md#createDraftArticle) | **POST** /v1/content/article/draft/create |  |
| [**createDraftArticleWithHttpInfo**](ArticleService.md#createDraftArticleWithHttpInfo) | **POST** /v1/content/article/draft/create |  |
| [**discardDraftArticle**](ArticleService.md#discardDraftArticle) | **POST** /v1/content/article/draft/discard |  |
| [**discardDraftArticleWithHttpInfo**](ArticleService.md#discardDraftArticleWithHttpInfo) | **POST** /v1/content/article/draft/discard |  |
| [**enterArticleLottery**](ArticleService.md#enterArticleLottery) | **POST** /v1/content/article/lottery/entry |  |
| [**enterArticleLotteryWithHttpInfo**](ArticleService.md#enterArticleLotteryWithHttpInfo) | **POST** /v1/content/article/lottery/entry |  |
| [**getArticle**](ArticleService.md#getArticle) | **POST** /v1/content/article/get |  |
| [**getArticleWithHttpInfo**](ArticleService.md#getArticleWithHttpInfo) | **POST** /v1/content/article/get |  |
| [**likeArticle**](ArticleService.md#likeArticle) | **POST** /v1/content/article/like |  |
| [**likeArticleWithHttpInfo**](ArticleService.md#likeArticleWithHttpInfo) | **POST** /v1/content/article/like |  |
| [**listArticles**](ArticleService.md#listArticles) | **POST** /v1/content/article/list |  |
| [**listArticlesWithHttpInfo**](ArticleService.md#listArticlesWithHttpInfo) | **POST** /v1/content/article/list |  |
| [**listFollowingArticles**](ArticleService.md#listFollowingArticles) | **POST** /v1/content/article/following/list |  |
| [**listFollowingArticlesWithHttpInfo**](ArticleService.md#listFollowingArticlesWithHttpInfo) | **POST** /v1/content/article/following/list |  |
| [**listNearbyArticles**](ArticleService.md#listNearbyArticles) | **POST** /v1/content/article/nearby/list |  |
| [**listNearbyArticlesWithHttpInfo**](ArticleService.md#listNearbyArticlesWithHttpInfo) | **POST** /v1/content/article/nearby/list |  |
| [**pageArticleViewHistory**](ArticleService.md#pageArticleViewHistory) | **POST** /v1/content/article/view-history/page |  |
| [**pageArticleViewHistoryWithHttpInfo**](ArticleService.md#pageArticleViewHistoryWithHttpInfo) | **POST** /v1/content/article/view-history/page |  |
| [**pageCollectedArticles**](ArticleService.md#pageCollectedArticles) | **POST** /v1/content/article/collected/page |  |
| [**pageCollectedArticlesWithHttpInfo**](ArticleService.md#pageCollectedArticlesWithHttpInfo) | **POST** /v1/content/article/collected/page |  |
| [**pageWatchedArticles**](ArticleService.md#pageWatchedArticles) | **POST** /v1/content/article/watched/page |  |
| [**pageWatchedArticlesWithHttpInfo**](ArticleService.md#pageWatchedArticlesWithHttpInfo) | **POST** /v1/content/article/watched/page |  |
| [**publishArticle**](ArticleService.md#publishArticle) | **POST** /v1/content/article/publish |  |
| [**publishArticleWithHttpInfo**](ArticleService.md#publishArticleWithHttpInfo) | **POST** /v1/content/article/publish |  |
| [**rewardArticle**](ArticleService.md#rewardArticle) | **POST** /v1/content/article/reward |  |
| [**rewardArticleWithHttpInfo**](ArticleService.md#rewardArticleWithHttpInfo) | **POST** /v1/content/article/reward |  |
| [**setPinnedArticleComment**](ArticleService.md#setPinnedArticleComment) | **POST** /v1/content/article/comment/pin |  |
| [**setPinnedArticleCommentWithHttpInfo**](ArticleService.md#setPinnedArticleCommentWithHttpInfo) | **POST** /v1/content/article/comment/pin |  |
| [**thankArticle**](ArticleService.md#thankArticle) | **POST** /v1/content/article/thank |  |
| [**thankArticleWithHttpInfo**](ArticleService.md#thankArticleWithHttpInfo) | **POST** /v1/content/article/thank |  |
| [**updateDraftArticle**](ArticleService.md#updateDraftArticle) | **POST** /v1/content/article/draft/update |  |
| [**updateDraftArticleWithHttpInfo**](ArticleService.md#updateDraftArticleWithHttpInfo) | **POST** /v1/content/article/draft/update |  |
| [**voteArticlePoll**](ArticleService.md#voteArticlePoll) | **POST** /v1/content/article/poll/vote |  |
| [**voteArticlePollWithHttpInfo**](ArticleService.md#voteArticlePollWithHttpInfo) | **POST** /v1/content/article/poll/vote |  |
| [**watchArticle**](ArticleService.md#watchArticle) | **POST** /v1/content/article/watch |  |
| [**watchArticleWithHttpInfo**](ArticleService.md#watchArticleWithHttpInfo) | **POST** /v1/content/article/watch |  |



## acceptArticleAnswer

> Object acceptArticleAnswer(acceptArticleAnswerRequest)



采纳文章回答

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
        AcceptArticleAnswerReq acceptArticleAnswerReq = new AcceptArticleAnswerReq(); // AcceptArticleAnswerReq | 
        try {
            APIacceptArticleAnswerRequest request = APIacceptArticleAnswerRequest.newBuilder()
                .acceptArticleAnswerReq(acceptArticleAnswerReq)
                .build();
            Object result = apiInstance.acceptArticleAnswer(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#acceptArticleAnswer");
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
| acceptArticleAnswerRequest | [**APIacceptArticleAnswerRequest**](ArticleService.md#APIacceptArticleAnswerRequest)|-|-|

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

## acceptArticleAnswerWithHttpInfo

> ApiResponse<Object> acceptArticleAnswerWithHttpInfo(acceptArticleAnswerRequest)



采纳文章回答

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
        AcceptArticleAnswerReq acceptArticleAnswerReq = new AcceptArticleAnswerReq(); // AcceptArticleAnswerReq | 
        try {
            APIacceptArticleAnswerRequest request = APIacceptArticleAnswerRequest.newBuilder()
                .acceptArticleAnswerReq(acceptArticleAnswerReq)
                .build();
            ApiResponse<Object> response = apiInstance.acceptArticleAnswerWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#acceptArticleAnswer");
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
| acceptArticleAnswerRequest | [**APIacceptArticleAnswerRequest**](ArticleService.md#APIacceptArticleAnswerRequest)|-|-|

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


<a id="APIacceptArticleAnswerRequest"></a>
## APIacceptArticleAnswerRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **acceptArticleAnswerReq** | [**AcceptArticleAnswerReq**](AcceptArticleAnswerReq.md) |  | |



## archiveArticle

> Object archiveArticle(archiveArticleRequest)



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
            APIarchiveArticleRequest request = APIarchiveArticleRequest.newBuilder()
                .archiveArticleReq(archiveArticleReq)
                .build();
            Object result = apiInstance.archiveArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#archiveArticle");
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
| archiveArticleRequest | [**APIarchiveArticleRequest**](ArticleService.md#APIarchiveArticleRequest)|-|-|

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

## archiveArticleWithHttpInfo

> ApiResponse<Object> archiveArticleWithHttpInfo(archiveArticleRequest)



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
            APIarchiveArticleRequest request = APIarchiveArticleRequest.newBuilder()
                .archiveArticleReq(archiveArticleReq)
                .build();
            ApiResponse<Object> response = apiInstance.archiveArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#archiveArticle");
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
| archiveArticleRequest | [**APIarchiveArticleRequest**](ArticleService.md#APIarchiveArticleRequest)|-|-|

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


<a id="APIarchiveArticleRequest"></a>
## APIarchiveArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **archiveArticleReq** | [**ArchiveArticleReq**](ArchiveArticleReq.md) |  | |



## cancelPublishArticle

> Object cancelPublishArticle(cancelPublishArticleRequest)



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
            APIcancelPublishArticleRequest request = APIcancelPublishArticleRequest.newBuilder()
                .cancelPublishArticleReq(cancelPublishArticleReq)
                .build();
            Object result = apiInstance.cancelPublishArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#cancelPublishArticle");
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
| cancelPublishArticleRequest | [**APIcancelPublishArticleRequest**](ArticleService.md#APIcancelPublishArticleRequest)|-|-|

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

## cancelPublishArticleWithHttpInfo

> ApiResponse<Object> cancelPublishArticleWithHttpInfo(cancelPublishArticleRequest)



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
            APIcancelPublishArticleRequest request = APIcancelPublishArticleRequest.newBuilder()
                .cancelPublishArticleReq(cancelPublishArticleReq)
                .build();
            ApiResponse<Object> response = apiInstance.cancelPublishArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#cancelPublishArticle");
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
| cancelPublishArticleRequest | [**APIcancelPublishArticleRequest**](ArticleService.md#APIcancelPublishArticleRequest)|-|-|

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


<a id="APIcancelPublishArticleRequest"></a>
## APIcancelPublishArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **cancelPublishArticleReq** | [**CancelPublishArticleReq**](CancelPublishArticleReq.md) |  | |



## collectArticle

> CollectArticleResp collectArticle(collectArticleRequest)



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
            APIcollectArticleRequest request = APIcollectArticleRequest.newBuilder()
                .collectArticleReq(collectArticleReq)
                .build();
            CollectArticleResp result = apiInstance.collectArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#collectArticle");
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
| collectArticleRequest | [**APIcollectArticleRequest**](ArticleService.md#APIcollectArticleRequest)|-|-|

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

## collectArticleWithHttpInfo

> ApiResponse<CollectArticleResp> collectArticleWithHttpInfo(collectArticleRequest)



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
            APIcollectArticleRequest request = APIcollectArticleRequest.newBuilder()
                .collectArticleReq(collectArticleReq)
                .build();
            ApiResponse<CollectArticleResp> response = apiInstance.collectArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#collectArticle");
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
| collectArticleRequest | [**APIcollectArticleRequest**](ArticleService.md#APIcollectArticleRequest)|-|-|

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


<a id="APIcollectArticleRequest"></a>
## APIcollectArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **collectArticleReq** | [**CollectArticleReq**](CollectArticleReq.md) |  | |



## countCollectedArticles

> CountCollectedArticlesResp countCollectedArticles(countCollectedArticlesRequest)



统计当前账号收藏的文章数量。

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
        Object body = null; // Object | 
        try {
            APIcountCollectedArticlesRequest request = APIcountCollectedArticlesRequest.newBuilder()
                .body(body)
                .build();
            CountCollectedArticlesResp result = apiInstance.countCollectedArticles(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#countCollectedArticles");
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
| countCollectedArticlesRequest | [**APIcountCollectedArticlesRequest**](ArticleService.md#APIcountCollectedArticlesRequest)|-|-|

### Return type

[**CountCollectedArticlesResp**](CountCollectedArticlesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## countCollectedArticlesWithHttpInfo

> ApiResponse<CountCollectedArticlesResp> countCollectedArticlesWithHttpInfo(countCollectedArticlesRequest)



统计当前账号收藏的文章数量。

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
        Object body = null; // Object | 
        try {
            APIcountCollectedArticlesRequest request = APIcountCollectedArticlesRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<CountCollectedArticlesResp> response = apiInstance.countCollectedArticlesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#countCollectedArticles");
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
| countCollectedArticlesRequest | [**APIcountCollectedArticlesRequest**](ArticleService.md#APIcountCollectedArticlesRequest)|-|-|

### Return type

ApiResponse<[**CountCollectedArticlesResp**](CountCollectedArticlesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcountCollectedArticlesRequest"></a>
## APIcountCollectedArticlesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## createDraftArticle

> CreateDraftArticleResp createDraftArticle(createDraftArticleRequest)



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
            APIcreateDraftArticleRequest request = APIcreateDraftArticleRequest.newBuilder()
                .createDraftArticleReq(createDraftArticleReq)
                .build();
            CreateDraftArticleResp result = apiInstance.createDraftArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#createDraftArticle");
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
| createDraftArticleRequest | [**APIcreateDraftArticleRequest**](ArticleService.md#APIcreateDraftArticleRequest)|-|-|

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

## createDraftArticleWithHttpInfo

> ApiResponse<CreateDraftArticleResp> createDraftArticleWithHttpInfo(createDraftArticleRequest)



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
            APIcreateDraftArticleRequest request = APIcreateDraftArticleRequest.newBuilder()
                .createDraftArticleReq(createDraftArticleReq)
                .build();
            ApiResponse<CreateDraftArticleResp> response = apiInstance.createDraftArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#createDraftArticle");
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
| createDraftArticleRequest | [**APIcreateDraftArticleRequest**](ArticleService.md#APIcreateDraftArticleRequest)|-|-|

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


<a id="APIcreateDraftArticleRequest"></a>
## APIcreateDraftArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **createDraftArticleReq** | [**CreateDraftArticleReq**](CreateDraftArticleReq.md) |  | |



## discardDraftArticle

> Object discardDraftArticle(discardDraftArticleRequest)



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
            APIdiscardDraftArticleRequest request = APIdiscardDraftArticleRequest.newBuilder()
                .discardDraftArticleReq(discardDraftArticleReq)
                .build();
            Object result = apiInstance.discardDraftArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#discardDraftArticle");
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
| discardDraftArticleRequest | [**APIdiscardDraftArticleRequest**](ArticleService.md#APIdiscardDraftArticleRequest)|-|-|

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

## discardDraftArticleWithHttpInfo

> ApiResponse<Object> discardDraftArticleWithHttpInfo(discardDraftArticleRequest)



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
            APIdiscardDraftArticleRequest request = APIdiscardDraftArticleRequest.newBuilder()
                .discardDraftArticleReq(discardDraftArticleReq)
                .build();
            ApiResponse<Object> response = apiInstance.discardDraftArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#discardDraftArticle");
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
| discardDraftArticleRequest | [**APIdiscardDraftArticleRequest**](ArticleService.md#APIdiscardDraftArticleRequest)|-|-|

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


<a id="APIdiscardDraftArticleRequest"></a>
## APIdiscardDraftArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **discardDraftArticleReq** | [**DiscardDraftArticleReq**](DiscardDraftArticleReq.md) |  | |



## enterArticleLottery

> Object enterArticleLottery(enterArticleLotteryRequest)



报名或撤销抽奖报名。

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
        EnterArticleLotteryReq enterArticleLotteryReq = new EnterArticleLotteryReq(); // EnterArticleLotteryReq | 
        try {
            APIenterArticleLotteryRequest request = APIenterArticleLotteryRequest.newBuilder()
                .enterArticleLotteryReq(enterArticleLotteryReq)
                .build();
            Object result = apiInstance.enterArticleLottery(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#enterArticleLottery");
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
| enterArticleLotteryRequest | [**APIenterArticleLotteryRequest**](ArticleService.md#APIenterArticleLotteryRequest)|-|-|

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

## enterArticleLotteryWithHttpInfo

> ApiResponse<Object> enterArticleLotteryWithHttpInfo(enterArticleLotteryRequest)



报名或撤销抽奖报名。

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
        EnterArticleLotteryReq enterArticleLotteryReq = new EnterArticleLotteryReq(); // EnterArticleLotteryReq | 
        try {
            APIenterArticleLotteryRequest request = APIenterArticleLotteryRequest.newBuilder()
                .enterArticleLotteryReq(enterArticleLotteryReq)
                .build();
            ApiResponse<Object> response = apiInstance.enterArticleLotteryWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#enterArticleLottery");
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
| enterArticleLotteryRequest | [**APIenterArticleLotteryRequest**](ArticleService.md#APIenterArticleLotteryRequest)|-|-|

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


<a id="APIenterArticleLotteryRequest"></a>
## APIenterArticleLotteryRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **enterArticleLotteryReq** | [**EnterArticleLotteryReq**](EnterArticleLotteryReq.md) |  | |



## getArticle

> GetArticleResp getArticle(getArticleRequest)



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
            APIgetArticleRequest request = APIgetArticleRequest.newBuilder()
                .getArticleReq(getArticleReq)
                .build();
            GetArticleResp result = apiInstance.getArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#getArticle");
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
| getArticleRequest | [**APIgetArticleRequest**](ArticleService.md#APIgetArticleRequest)|-|-|

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

## getArticleWithHttpInfo

> ApiResponse<GetArticleResp> getArticleWithHttpInfo(getArticleRequest)



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
            APIgetArticleRequest request = APIgetArticleRequest.newBuilder()
                .getArticleReq(getArticleReq)
                .build();
            ApiResponse<GetArticleResp> response = apiInstance.getArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#getArticle");
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
| getArticleRequest | [**APIgetArticleRequest**](ArticleService.md#APIgetArticleRequest)|-|-|

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


<a id="APIgetArticleRequest"></a>
## APIgetArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getArticleReq** | [**GetArticleReq**](GetArticleReq.md) |  | |



## likeArticle

> LikeArticleResp likeArticle(likeArticleRequest)



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
            APIlikeArticleRequest request = APIlikeArticleRequest.newBuilder()
                .likeArticleReq(likeArticleReq)
                .build();
            LikeArticleResp result = apiInstance.likeArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#likeArticle");
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
| likeArticleRequest | [**APIlikeArticleRequest**](ArticleService.md#APIlikeArticleRequest)|-|-|

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

## likeArticleWithHttpInfo

> ApiResponse<LikeArticleResp> likeArticleWithHttpInfo(likeArticleRequest)



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
            APIlikeArticleRequest request = APIlikeArticleRequest.newBuilder()
                .likeArticleReq(likeArticleReq)
                .build();
            ApiResponse<LikeArticleResp> response = apiInstance.likeArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#likeArticle");
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
| likeArticleRequest | [**APIlikeArticleRequest**](ArticleService.md#APIlikeArticleRequest)|-|-|

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


<a id="APIlikeArticleRequest"></a>
## APIlikeArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **likeArticleReq** | [**LikeArticleReq**](LikeArticleReq.md) |  | |



## listArticles

> ListArticlesResp listArticles(listArticlesRequest)



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
            APIlistArticlesRequest request = APIlistArticlesRequest.newBuilder()
                .listArticlesReq(listArticlesReq)
                .build();
            ListArticlesResp result = apiInstance.listArticles(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#listArticles");
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
| listArticlesRequest | [**APIlistArticlesRequest**](ArticleService.md#APIlistArticlesRequest)|-|-|

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

## listArticlesWithHttpInfo

> ApiResponse<ListArticlesResp> listArticlesWithHttpInfo(listArticlesRequest)



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
            APIlistArticlesRequest request = APIlistArticlesRequest.newBuilder()
                .listArticlesReq(listArticlesReq)
                .build();
            ApiResponse<ListArticlesResp> response = apiInstance.listArticlesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#listArticles");
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
| listArticlesRequest | [**APIlistArticlesRequest**](ArticleService.md#APIlistArticlesRequest)|-|-|

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


<a id="APIlistArticlesRequest"></a>
## APIlistArticlesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listArticlesReq** | [**ListArticlesReq**](ListArticlesReq.md) |  | |



## listFollowingArticles

> ListFollowingArticlesResp listFollowingArticles(listFollowingArticlesRequest)



查询当前账号关注作者发布的文章。

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
        ListFollowingArticlesReq listFollowingArticlesReq = new ListFollowingArticlesReq(); // ListFollowingArticlesReq | 
        try {
            APIlistFollowingArticlesRequest request = APIlistFollowingArticlesRequest.newBuilder()
                .listFollowingArticlesReq(listFollowingArticlesReq)
                .build();
            ListFollowingArticlesResp result = apiInstance.listFollowingArticles(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#listFollowingArticles");
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
| listFollowingArticlesRequest | [**APIlistFollowingArticlesRequest**](ArticleService.md#APIlistFollowingArticlesRequest)|-|-|

### Return type

[**ListFollowingArticlesResp**](ListFollowingArticlesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listFollowingArticlesWithHttpInfo

> ApiResponse<ListFollowingArticlesResp> listFollowingArticlesWithHttpInfo(listFollowingArticlesRequest)



查询当前账号关注作者发布的文章。

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
        ListFollowingArticlesReq listFollowingArticlesReq = new ListFollowingArticlesReq(); // ListFollowingArticlesReq | 
        try {
            APIlistFollowingArticlesRequest request = APIlistFollowingArticlesRequest.newBuilder()
                .listFollowingArticlesReq(listFollowingArticlesReq)
                .build();
            ApiResponse<ListFollowingArticlesResp> response = apiInstance.listFollowingArticlesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#listFollowingArticles");
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
| listFollowingArticlesRequest | [**APIlistFollowingArticlesRequest**](ArticleService.md#APIlistFollowingArticlesRequest)|-|-|

### Return type

ApiResponse<[**ListFollowingArticlesResp**](ListFollowingArticlesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistFollowingArticlesRequest"></a>
## APIlistFollowingArticlesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowingArticlesReq** | [**ListFollowingArticlesReq**](ListFollowingArticlesReq.md) |  | |



## listNearbyArticles

> ListNearbyArticlesResp listNearbyArticles(listNearbyArticlesRequest)



查询当前账号所在城市发布的文章。

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
        ListNearbyArticlesReq listNearbyArticlesReq = new ListNearbyArticlesReq(); // ListNearbyArticlesReq | 
        try {
            APIlistNearbyArticlesRequest request = APIlistNearbyArticlesRequest.newBuilder()
                .listNearbyArticlesReq(listNearbyArticlesReq)
                .build();
            ListNearbyArticlesResp result = apiInstance.listNearbyArticles(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#listNearbyArticles");
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
| listNearbyArticlesRequest | [**APIlistNearbyArticlesRequest**](ArticleService.md#APIlistNearbyArticlesRequest)|-|-|

### Return type

[**ListNearbyArticlesResp**](ListNearbyArticlesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listNearbyArticlesWithHttpInfo

> ApiResponse<ListNearbyArticlesResp> listNearbyArticlesWithHttpInfo(listNearbyArticlesRequest)



查询当前账号所在城市发布的文章。

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
        ListNearbyArticlesReq listNearbyArticlesReq = new ListNearbyArticlesReq(); // ListNearbyArticlesReq | 
        try {
            APIlistNearbyArticlesRequest request = APIlistNearbyArticlesRequest.newBuilder()
                .listNearbyArticlesReq(listNearbyArticlesReq)
                .build();
            ApiResponse<ListNearbyArticlesResp> response = apiInstance.listNearbyArticlesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#listNearbyArticles");
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
| listNearbyArticlesRequest | [**APIlistNearbyArticlesRequest**](ArticleService.md#APIlistNearbyArticlesRequest)|-|-|

### Return type

ApiResponse<[**ListNearbyArticlesResp**](ListNearbyArticlesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistNearbyArticlesRequest"></a>
## APIlistNearbyArticlesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listNearbyArticlesReq** | [**ListNearbyArticlesReq**](ListNearbyArticlesReq.md) |  | |



## pageArticleViewHistory

> PageArticleViewHistoryResp pageArticleViewHistory(pageArticleViewHistoryRequest)



查询当前账号的浏览历史。

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
        PageArticleViewHistoryReq pageArticleViewHistoryReq = new PageArticleViewHistoryReq(); // PageArticleViewHistoryReq | 
        try {
            APIpageArticleViewHistoryRequest request = APIpageArticleViewHistoryRequest.newBuilder()
                .pageArticleViewHistoryReq(pageArticleViewHistoryReq)
                .build();
            PageArticleViewHistoryResp result = apiInstance.pageArticleViewHistory(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#pageArticleViewHistory");
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
| pageArticleViewHistoryRequest | [**APIpageArticleViewHistoryRequest**](ArticleService.md#APIpageArticleViewHistoryRequest)|-|-|

### Return type

[**PageArticleViewHistoryResp**](PageArticleViewHistoryResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## pageArticleViewHistoryWithHttpInfo

> ApiResponse<PageArticleViewHistoryResp> pageArticleViewHistoryWithHttpInfo(pageArticleViewHistoryRequest)



查询当前账号的浏览历史。

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
        PageArticleViewHistoryReq pageArticleViewHistoryReq = new PageArticleViewHistoryReq(); // PageArticleViewHistoryReq | 
        try {
            APIpageArticleViewHistoryRequest request = APIpageArticleViewHistoryRequest.newBuilder()
                .pageArticleViewHistoryReq(pageArticleViewHistoryReq)
                .build();
            ApiResponse<PageArticleViewHistoryResp> response = apiInstance.pageArticleViewHistoryWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#pageArticleViewHistory");
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
| pageArticleViewHistoryRequest | [**APIpageArticleViewHistoryRequest**](ArticleService.md#APIpageArticleViewHistoryRequest)|-|-|

### Return type

ApiResponse<[**PageArticleViewHistoryResp**](PageArticleViewHistoryResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpageArticleViewHistoryRequest"></a>
## APIpageArticleViewHistoryRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pageArticleViewHistoryReq** | [**PageArticleViewHistoryReq**](PageArticleViewHistoryReq.md) |  | |



## pageCollectedArticles

> PageCollectedArticlesResp pageCollectedArticles(pageCollectedArticlesRequest)



查询当前账号收藏的文章。

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
        PageCollectedArticlesReq pageCollectedArticlesReq = new PageCollectedArticlesReq(); // PageCollectedArticlesReq | 
        try {
            APIpageCollectedArticlesRequest request = APIpageCollectedArticlesRequest.newBuilder()
                .pageCollectedArticlesReq(pageCollectedArticlesReq)
                .build();
            PageCollectedArticlesResp result = apiInstance.pageCollectedArticles(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#pageCollectedArticles");
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
| pageCollectedArticlesRequest | [**APIpageCollectedArticlesRequest**](ArticleService.md#APIpageCollectedArticlesRequest)|-|-|

### Return type

[**PageCollectedArticlesResp**](PageCollectedArticlesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## pageCollectedArticlesWithHttpInfo

> ApiResponse<PageCollectedArticlesResp> pageCollectedArticlesWithHttpInfo(pageCollectedArticlesRequest)



查询当前账号收藏的文章。

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
        PageCollectedArticlesReq pageCollectedArticlesReq = new PageCollectedArticlesReq(); // PageCollectedArticlesReq | 
        try {
            APIpageCollectedArticlesRequest request = APIpageCollectedArticlesRequest.newBuilder()
                .pageCollectedArticlesReq(pageCollectedArticlesReq)
                .build();
            ApiResponse<PageCollectedArticlesResp> response = apiInstance.pageCollectedArticlesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#pageCollectedArticles");
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
| pageCollectedArticlesRequest | [**APIpageCollectedArticlesRequest**](ArticleService.md#APIpageCollectedArticlesRequest)|-|-|

### Return type

ApiResponse<[**PageCollectedArticlesResp**](PageCollectedArticlesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpageCollectedArticlesRequest"></a>
## APIpageCollectedArticlesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pageCollectedArticlesReq** | [**PageCollectedArticlesReq**](PageCollectedArticlesReq.md) |  | |



## pageWatchedArticles

> PageWatchedArticlesResp pageWatchedArticles(pageWatchedArticlesRequest)



查询当前账号关注的文章。

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
        PageWatchedArticlesReq pageWatchedArticlesReq = new PageWatchedArticlesReq(); // PageWatchedArticlesReq | 
        try {
            APIpageWatchedArticlesRequest request = APIpageWatchedArticlesRequest.newBuilder()
                .pageWatchedArticlesReq(pageWatchedArticlesReq)
                .build();
            PageWatchedArticlesResp result = apiInstance.pageWatchedArticles(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#pageWatchedArticles");
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
| pageWatchedArticlesRequest | [**APIpageWatchedArticlesRequest**](ArticleService.md#APIpageWatchedArticlesRequest)|-|-|

### Return type

[**PageWatchedArticlesResp**](PageWatchedArticlesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## pageWatchedArticlesWithHttpInfo

> ApiResponse<PageWatchedArticlesResp> pageWatchedArticlesWithHttpInfo(pageWatchedArticlesRequest)



查询当前账号关注的文章。

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
        PageWatchedArticlesReq pageWatchedArticlesReq = new PageWatchedArticlesReq(); // PageWatchedArticlesReq | 
        try {
            APIpageWatchedArticlesRequest request = APIpageWatchedArticlesRequest.newBuilder()
                .pageWatchedArticlesReq(pageWatchedArticlesReq)
                .build();
            ApiResponse<PageWatchedArticlesResp> response = apiInstance.pageWatchedArticlesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#pageWatchedArticles");
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
| pageWatchedArticlesRequest | [**APIpageWatchedArticlesRequest**](ArticleService.md#APIpageWatchedArticlesRequest)|-|-|

### Return type

ApiResponse<[**PageWatchedArticlesResp**](PageWatchedArticlesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpageWatchedArticlesRequest"></a>
## APIpageWatchedArticlesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pageWatchedArticlesReq** | [**PageWatchedArticlesReq**](PageWatchedArticlesReq.md) |  | |



## publishArticle

> Object publishArticle(publishArticleRequest)



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
            APIpublishArticleRequest request = APIpublishArticleRequest.newBuilder()
                .publishArticleReq(publishArticleReq)
                .build();
            Object result = apiInstance.publishArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#publishArticle");
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
| publishArticleRequest | [**APIpublishArticleRequest**](ArticleService.md#APIpublishArticleRequest)|-|-|

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

## publishArticleWithHttpInfo

> ApiResponse<Object> publishArticleWithHttpInfo(publishArticleRequest)



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
            APIpublishArticleRequest request = APIpublishArticleRequest.newBuilder()
                .publishArticleReq(publishArticleReq)
                .build();
            ApiResponse<Object> response = apiInstance.publishArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#publishArticle");
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
| publishArticleRequest | [**APIpublishArticleRequest**](ArticleService.md#APIpublishArticleRequest)|-|-|

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


<a id="APIpublishArticleRequest"></a>
## APIpublishArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **publishArticleReq** | [**PublishArticleReq**](PublishArticleReq.md) |  | |



## rewardArticle

> RewardArticleResp rewardArticle(rewardArticleRequest)



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
            APIrewardArticleRequest request = APIrewardArticleRequest.newBuilder()
                .rewardArticleReq(rewardArticleReq)
                .build();
            RewardArticleResp result = apiInstance.rewardArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#rewardArticle");
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
| rewardArticleRequest | [**APIrewardArticleRequest**](ArticleService.md#APIrewardArticleRequest)|-|-|

### Return type

[**RewardArticleResp**](RewardArticleResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## rewardArticleWithHttpInfo

> ApiResponse<RewardArticleResp> rewardArticleWithHttpInfo(rewardArticleRequest)



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
            APIrewardArticleRequest request = APIrewardArticleRequest.newBuilder()
                .rewardArticleReq(rewardArticleReq)
                .build();
            ApiResponse<RewardArticleResp> response = apiInstance.rewardArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#rewardArticle");
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
| rewardArticleRequest | [**APIrewardArticleRequest**](ArticleService.md#APIrewardArticleRequest)|-|-|

### Return type

ApiResponse<[**RewardArticleResp**](RewardArticleResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrewardArticleRequest"></a>
## APIrewardArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **rewardArticleReq** | [**RewardArticleReq**](RewardArticleReq.md) |  | |



## setPinnedArticleComment

> Object setPinnedArticleComment(setPinnedArticleCommentRequest)



设置或取消当前账号文章的置顶评论。

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
        SetPinnedArticleCommentReq setPinnedArticleCommentReq = new SetPinnedArticleCommentReq(); // SetPinnedArticleCommentReq | 
        try {
            APIsetPinnedArticleCommentRequest request = APIsetPinnedArticleCommentRequest.newBuilder()
                .setPinnedArticleCommentReq(setPinnedArticleCommentReq)
                .build();
            Object result = apiInstance.setPinnedArticleComment(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#setPinnedArticleComment");
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
| setPinnedArticleCommentRequest | [**APIsetPinnedArticleCommentRequest**](ArticleService.md#APIsetPinnedArticleCommentRequest)|-|-|

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

## setPinnedArticleCommentWithHttpInfo

> ApiResponse<Object> setPinnedArticleCommentWithHttpInfo(setPinnedArticleCommentRequest)



设置或取消当前账号文章的置顶评论。

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
        SetPinnedArticleCommentReq setPinnedArticleCommentReq = new SetPinnedArticleCommentReq(); // SetPinnedArticleCommentReq | 
        try {
            APIsetPinnedArticleCommentRequest request = APIsetPinnedArticleCommentRequest.newBuilder()
                .setPinnedArticleCommentReq(setPinnedArticleCommentReq)
                .build();
            ApiResponse<Object> response = apiInstance.setPinnedArticleCommentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#setPinnedArticleComment");
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
| setPinnedArticleCommentRequest | [**APIsetPinnedArticleCommentRequest**](ArticleService.md#APIsetPinnedArticleCommentRequest)|-|-|

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


<a id="APIsetPinnedArticleCommentRequest"></a>
## APIsetPinnedArticleCommentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **setPinnedArticleCommentReq** | [**SetPinnedArticleCommentReq**](SetPinnedArticleCommentReq.md) |  | |



## thankArticle

> ThankArticleResp thankArticle(thankArticleRequest)



感谢文章，确认后不可撤销。

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
            APIthankArticleRequest request = APIthankArticleRequest.newBuilder()
                .thankArticleReq(thankArticleReq)
                .build();
            ThankArticleResp result = apiInstance.thankArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#thankArticle");
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
| thankArticleRequest | [**APIthankArticleRequest**](ArticleService.md#APIthankArticleRequest)|-|-|

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

## thankArticleWithHttpInfo

> ApiResponse<ThankArticleResp> thankArticleWithHttpInfo(thankArticleRequest)



感谢文章，确认后不可撤销。

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
            APIthankArticleRequest request = APIthankArticleRequest.newBuilder()
                .thankArticleReq(thankArticleReq)
                .build();
            ApiResponse<ThankArticleResp> response = apiInstance.thankArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#thankArticle");
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
| thankArticleRequest | [**APIthankArticleRequest**](ArticleService.md#APIthankArticleRequest)|-|-|

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


<a id="APIthankArticleRequest"></a>
## APIthankArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **thankArticleReq** | [**ThankArticleReq**](ThankArticleReq.md) |  | |



## updateDraftArticle

> UpdateDraftArticleResp updateDraftArticle(updateDraftArticleRequest)



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
            APIupdateDraftArticleRequest request = APIupdateDraftArticleRequest.newBuilder()
                .updateDraftArticleReq(updateDraftArticleReq)
                .build();
            UpdateDraftArticleResp result = apiInstance.updateDraftArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#updateDraftArticle");
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
| updateDraftArticleRequest | [**APIupdateDraftArticleRequest**](ArticleService.md#APIupdateDraftArticleRequest)|-|-|

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

## updateDraftArticleWithHttpInfo

> ApiResponse<UpdateDraftArticleResp> updateDraftArticleWithHttpInfo(updateDraftArticleRequest)



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
            APIupdateDraftArticleRequest request = APIupdateDraftArticleRequest.newBuilder()
                .updateDraftArticleReq(updateDraftArticleReq)
                .build();
            ApiResponse<UpdateDraftArticleResp> response = apiInstance.updateDraftArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#updateDraftArticle");
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
| updateDraftArticleRequest | [**APIupdateDraftArticleRequest**](ArticleService.md#APIupdateDraftArticleRequest)|-|-|

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


<a id="APIupdateDraftArticleRequest"></a>
## APIupdateDraftArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateDraftArticleReq** | [**UpdateDraftArticleReq**](UpdateDraftArticleReq.md) |  | |



## voteArticlePoll

> Object voteArticlePoll(voteArticlePollRequest)



提交一次不可修改的文章投票。

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
        VoteArticlePollReq voteArticlePollReq = new VoteArticlePollReq(); // VoteArticlePollReq | 
        try {
            APIvoteArticlePollRequest request = APIvoteArticlePollRequest.newBuilder()
                .voteArticlePollReq(voteArticlePollReq)
                .build();
            Object result = apiInstance.voteArticlePoll(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#voteArticlePoll");
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
| voteArticlePollRequest | [**APIvoteArticlePollRequest**](ArticleService.md#APIvoteArticlePollRequest)|-|-|

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

## voteArticlePollWithHttpInfo

> ApiResponse<Object> voteArticlePollWithHttpInfo(voteArticlePollRequest)



提交一次不可修改的文章投票。

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
        VoteArticlePollReq voteArticlePollReq = new VoteArticlePollReq(); // VoteArticlePollReq | 
        try {
            APIvoteArticlePollRequest request = APIvoteArticlePollRequest.newBuilder()
                .voteArticlePollReq(voteArticlePollReq)
                .build();
            ApiResponse<Object> response = apiInstance.voteArticlePollWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#voteArticlePoll");
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
| voteArticlePollRequest | [**APIvoteArticlePollRequest**](ArticleService.md#APIvoteArticlePollRequest)|-|-|

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


<a id="APIvoteArticlePollRequest"></a>
## APIvoteArticlePollRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **voteArticlePollReq** | [**VoteArticlePollReq**](VoteArticlePollReq.md) |  | |



## watchArticle

> WatchArticleResp watchArticle(watchArticleRequest)



关注文章

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
        WatchArticleReq watchArticleReq = new WatchArticleReq(); // WatchArticleReq | 
        try {
            APIwatchArticleRequest request = APIwatchArticleRequest.newBuilder()
                .watchArticleReq(watchArticleReq)
                .build();
            WatchArticleResp result = apiInstance.watchArticle(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#watchArticle");
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
| watchArticleRequest | [**APIwatchArticleRequest**](ArticleService.md#APIwatchArticleRequest)|-|-|

### Return type

[**WatchArticleResp**](WatchArticleResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## watchArticleWithHttpInfo

> ApiResponse<WatchArticleResp> watchArticleWithHttpInfo(watchArticleRequest)



关注文章

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
        WatchArticleReq watchArticleReq = new WatchArticleReq(); // WatchArticleReq | 
        try {
            APIwatchArticleRequest request = APIwatchArticleRequest.newBuilder()
                .watchArticleReq(watchArticleReq)
                .build();
            ApiResponse<WatchArticleResp> response = apiInstance.watchArticleWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleService#watchArticle");
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
| watchArticleRequest | [**APIwatchArticleRequest**](ArticleService.md#APIwatchArticleRequest)|-|-|

### Return type

ApiResponse<[**WatchArticleResp**](WatchArticleResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIwatchArticleRequest"></a>
## APIwatchArticleRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **watchArticleReq** | [**WatchArticleReq**](WatchArticleReq.md) |  | |


