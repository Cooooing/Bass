# ArticleSearchService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**search**](ArticleSearchService.md#search) | **POST** /v1/content/article/search |  |
| [**searchWithHttpInfo**](ArticleSearchService.md#searchWithHttpInfo) | **POST** /v1/content/article/search |  |



## search

> SearchArticlesResp search(searchRequest)



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
            APIsearchRequest request = APIsearchRequest.newBuilder()
                .searchArticlesReq(searchArticlesReq)
                .build();
            SearchArticlesResp result = apiInstance.search(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleSearchService#search");
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
| searchRequest | [**APIsearchRequest**](ArticleSearchService.md#APIsearchRequest)|-|-|

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

## searchWithHttpInfo

> ApiResponse<SearchArticlesResp> searchWithHttpInfo(searchRequest)



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
            APIsearchRequest request = APIsearchRequest.newBuilder()
                .searchArticlesReq(searchArticlesReq)
                .build();
            ApiResponse<SearchArticlesResp> response = apiInstance.searchWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ArticleSearchService#search");
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
| searchRequest | [**APIsearchRequest**](ArticleSearchService.md#APIsearchRequest)|-|-|

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


<a id="APIsearchRequest"></a>
## APIsearchRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **searchArticlesReq** | [**SearchArticlesReq**](SearchArticlesReq.md) |  | |


