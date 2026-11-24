# ArticleSearchService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**searchArticles**](#searcharticles) | **POST** /v1/content/article/search | |

# **searchArticles**
> SearchArticlesResp searchArticles(searchArticlesReq)

服务方法。

### Example

```typescript
import {
    ArticleSearchService,
    Configuration,
    SearchArticlesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleSearchService(configuration);

let searchArticlesReq: SearchArticlesReq; //

const { status, data } = await apiInstance.searchArticles(
    searchArticlesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **searchArticlesReq** | **SearchArticlesReq**|  | |


### Return type

**SearchArticlesResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

