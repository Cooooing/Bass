# ArticleSearchService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**search**](ArticleSearchService.md#search) | **POST** /v1/content/article/search |  |



## search

> SearchArticlesResp search(searchArticlesReq)



### Example

```ts
import {
  Configuration,
  ArticleSearchService,
} from '@bass/bbs-sdk-fetch';
import type { SearchRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleSearchService();

  const body = {
    // SearchArticlesReq
    searchArticlesReq: ...,
  } satisfies SearchRequest;

  try {
    const data = await api.search(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **searchArticlesReq** | [SearchArticlesReq](SearchArticlesReq.md) |  | |

### Return type

[**SearchArticlesResp**](SearchArticlesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

