# PostscriptService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**add**](PostscriptService.md#add) | **POST** /v1/content/postscript/add |  |



## add

> AddPostscriptReply add(addPostscriptRequest)



添加文章附言。

### Example

```ts
import {
  Configuration,
  PostscriptService,
} from '@bass/bbs-sdk-fetch';
import type { AddRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PostscriptService();

  const body = {
    // AddPostscriptRequest
    addPostscriptRequest: ...,
  } satisfies AddRequest;

  try {
    const data = await api.add(body);
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
| **addPostscriptRequest** | [AddPostscriptRequest](AddPostscriptRequest.md) |  | |

### Return type

[**AddPostscriptReply**](AddPostscriptReply.md)

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

