# PostscriptServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**postscriptServiceAdd**](PostscriptServiceApi.md#postscriptserviceadd) | **POST** /v1/content/postscript/add |  |



## postscriptServiceAdd

> AddPostscriptReply postscriptServiceAdd(addPostscriptRequest)



添加文章附言

### Example

```ts
import {
  Configuration,
  PostscriptServiceApi,
} from '@bass/bbs-sdk';
import type { PostscriptServiceAddRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new PostscriptServiceApi();

  const body = {
    // AddPostscriptRequest
    addPostscriptRequest: ...,
  } satisfies PostscriptServiceAddRequest;

  try {
    const data = await api.postscriptServiceAdd(body);
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

