# ArticleSearchService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**searchArticles**](ArticleSearchService.md#searchArticles) | **POST** /v1/content/article/search |  |
| [**searchArticlesWithHttpInfo**](ArticleSearchService.md#searchArticlesWithHttpInfo) | **POST** /v1/content/article/search |  |



## searchArticles

> SearchArticlesResp searchArticles(searchArticlesRequest)



服务方法。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleSearchService;
import com.bass.bbs.api.ArticleSearchService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleSearchService apiInstance = new ArticleSearchService(defaultClient);
        SearchArticlesReq searchArticlesReq = new SearchArticlesReq(); // SearchArticlesReq | 
        try {
            APIsearchArticlesRequest request = APIsearchArticlesRequest.newBuilder()
                .searchArticlesReq(searchArticlesReq)
                .build();
            SearchArticlesResp result = apiInstance.searchArticles(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleSearchService#searchArticles");
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
| searchArticlesRequest | [**APIsearchArticlesRequest**](ArticleSearchService.md#APIsearchArticlesRequest)|-|-|

### Return type

[**SearchArticlesResp**](SearchArticlesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## searchArticlesWithHttpInfo

> ApiResponse<SearchArticlesResp> searchArticlesWithHttpInfo(searchArticlesRequest)



服务方法。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ArticleSearchService;
import com.bass.bbs.api.ArticleSearchService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ArticleSearchService apiInstance = new ArticleSearchService(defaultClient);
        SearchArticlesReq searchArticlesReq = new SearchArticlesReq(); // SearchArticlesReq | 
        try {
            APIsearchArticlesRequest request = APIsearchArticlesRequest.newBuilder()
                .searchArticlesReq(searchArticlesReq)
                .build();
            ApiResponse<SearchArticlesResp> response = apiInstance.searchArticlesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleSearchService#searchArticles");
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
| searchArticlesRequest | [**APIsearchArticlesRequest**](ArticleSearchService.md#APIsearchArticlesRequest)|-|-|

### Return type

ApiResponse<[**SearchArticlesResp**](SearchArticlesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsearchArticlesRequest"></a>
## APIsearchArticlesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **searchArticlesReq** | [**SearchArticlesReq**](SearchArticlesReq.md) |  | |


