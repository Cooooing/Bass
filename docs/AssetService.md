# AssetService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**completeMarkdownImageUpload**](AssetService.md#completemarkdownimageupload) | **POST** /v1/asset/markdown-image/complete |  |
| [**prepareMarkdownImageUpload**](AssetService.md#preparemarkdownimageupload) | **POST** /v1/asset/markdown-image/prepare |  |



## completeMarkdownImageUpload

> CompleteMarkdownImageUploadResp completeMarkdownImageUpload(completeMarkdownImageUploadReq)



查询已由对象存储回调确认的 Markdown 图片。

### Example

```ts
import {
  Configuration,
  AssetService,
} from '@bass/bbs-sdk-fetch';
import type { CompleteMarkdownImageUploadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AssetService();

  const body = {
    // CompleteMarkdownImageUploadReq
    completeMarkdownImageUploadReq: ...,
  } satisfies CompleteMarkdownImageUploadRequest;

  try {
    const data = await api.completeMarkdownImageUpload(body);
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
| **completeMarkdownImageUploadReq** | [CompleteMarkdownImageUploadReq](CompleteMarkdownImageUploadReq.md) |  | |

### Return type

[**CompleteMarkdownImageUploadResp**](CompleteMarkdownImageUploadResp.md)

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


## prepareMarkdownImageUpload

> PrepareMarkdownImageUploadResp prepareMarkdownImageUpload(prepareMarkdownImageUploadReq)



申请 Markdown 图片的内容寻址直传凭证。

### Example

```ts
import {
  Configuration,
  AssetService,
} from '@bass/bbs-sdk-fetch';
import type { PrepareMarkdownImageUploadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AssetService();

  const body = {
    // PrepareMarkdownImageUploadReq
    prepareMarkdownImageUploadReq: ...,
  } satisfies PrepareMarkdownImageUploadRequest;

  try {
    const data = await api.prepareMarkdownImageUpload(body);
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
| **prepareMarkdownImageUploadReq** | [PrepareMarkdownImageUploadReq](PrepareMarkdownImageUploadReq.md) |  | |

### Return type

[**PrepareMarkdownImageUploadResp**](PrepareMarkdownImageUploadResp.md)

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

