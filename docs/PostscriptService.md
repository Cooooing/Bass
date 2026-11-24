# PostscriptService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addPostscript**](PostscriptService.md#addpostscript) | **POST** /v1/content/postscript/add |  |
| [**listPostscripts**](PostscriptService.md#listpostscripts) | **POST** /v1/content/postscript/list |  |



## addPostscript

> AddPostscriptResp addPostscript(addPostscriptReq)



添加文章附言。

### Example

```ts
import {
  Configuration,
  PostscriptService,
} from '@bass/bbs-sdk-fetch';
import type { AddPostscriptRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PostscriptService();

  const body = {
    // AddPostscriptReq
    addPostscriptReq: ...,
  } satisfies AddPostscriptRequest;

  try {
    const data = await api.addPostscript(body);
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
| **addPostscriptReq** | [AddPostscriptReq](AddPostscriptReq.md) |  | |

### Return type

[**AddPostscriptResp**](AddPostscriptResp.md)

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


## listPostscripts

> ListPostscriptsResp listPostscripts(listPostscriptsReq)



查询文章附言列表。

### Example

```ts
import {
  Configuration,
  PostscriptService,
} from '@bass/bbs-sdk-fetch';
import type { ListPostscriptsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PostscriptService();

  const body = {
    // ListPostscriptsReq
    listPostscriptsReq: ...,
  } satisfies ListPostscriptsRequest;

  try {
    const data = await api.listPostscripts(body);
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
| **listPostscriptsReq** | [ListPostscriptsReq](ListPostscriptsReq.md) |  | |

### Return type

[**ListPostscriptsResp**](ListPostscriptsResp.md)

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

