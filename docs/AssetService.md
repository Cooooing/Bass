# AssetService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**completeBadgeFontUpload**](AssetService.md#completebadgefontupload) | **POST** /v1/asset/badge-font/complete |  |
| [**completeBadgeImageUpload**](AssetService.md#completebadgeimageupload) | **POST** /v1/asset/badge-image/complete |  |
| [**completeMarkdownImageUpload**](AssetService.md#completemarkdownimageupload) | **POST** /v1/asset/markdown-image/complete |  |
| [**prepareBadgeFontUpload**](AssetService.md#preparebadgefontupload) | **POST** /v1/asset/badge-font/prepare |  |
| [**prepareBadgeImageUpload**](AssetService.md#preparebadgeimageupload) | **POST** /v1/asset/badge-image/prepare |  |
| [**prepareMarkdownImageUpload**](AssetService.md#preparemarkdownimageupload) | **POST** /v1/asset/markdown-image/prepare |  |



## completeBadgeFontUpload

> CompleteBadgeFontUploadResp completeBadgeFontUpload(completeBadgeFontUploadReq)



查询已由对象存储回调确认的徽章 Maker 自定义字体。

### Example

```ts
import {
  Configuration,
  AssetService,
} from '@bass/bbs-sdk-fetch';
import type { CompleteBadgeFontUploadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AssetService();

  const body = {
    // CompleteBadgeFontUploadReq
    completeBadgeFontUploadReq: ...,
  } satisfies CompleteBadgeFontUploadRequest;

  try {
    const data = await api.completeBadgeFontUpload(body);
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
| **completeBadgeFontUploadReq** | [CompleteBadgeFontUploadReq](CompleteBadgeFontUploadReq.md) |  | |

### Return type

[**CompleteBadgeFontUploadResp**](CompleteBadgeFontUploadResp.md)

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


## completeBadgeImageUpload

> CompleteBadgeImageUploadResp completeBadgeImageUpload(completeBadgeImageUploadReq)



查询已由对象存储回调确认的徽章 Maker 图片。

### Example

```ts
import {
  Configuration,
  AssetService,
} from '@bass/bbs-sdk-fetch';
import type { CompleteBadgeImageUploadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AssetService();

  const body = {
    // CompleteBadgeImageUploadReq
    completeBadgeImageUploadReq: ...,
  } satisfies CompleteBadgeImageUploadRequest;

  try {
    const data = await api.completeBadgeImageUpload(body);
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
| **completeBadgeImageUploadReq** | [CompleteBadgeImageUploadReq](CompleteBadgeImageUploadReq.md) |  | |

### Return type

[**CompleteBadgeImageUploadResp**](CompleteBadgeImageUploadResp.md)

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


## prepareBadgeFontUpload

> PrepareBadgeFontUploadResp prepareBadgeFontUpload(prepareBadgeFontUploadReq)



申请徽章 Maker 自定义字体的内容寻址直传凭证。

### Example

```ts
import {
  Configuration,
  AssetService,
} from '@bass/bbs-sdk-fetch';
import type { PrepareBadgeFontUploadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AssetService();

  const body = {
    // PrepareBadgeFontUploadReq
    prepareBadgeFontUploadReq: ...,
  } satisfies PrepareBadgeFontUploadRequest;

  try {
    const data = await api.prepareBadgeFontUpload(body);
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
| **prepareBadgeFontUploadReq** | [PrepareBadgeFontUploadReq](PrepareBadgeFontUploadReq.md) |  | |

### Return type

[**PrepareBadgeFontUploadResp**](PrepareBadgeFontUploadResp.md)

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


## prepareBadgeImageUpload

> PrepareBadgeImageUploadResp prepareBadgeImageUpload(prepareBadgeImageUploadReq)



申请徽章 Maker 图片的内容寻址直传凭证。

### Example

```ts
import {
  Configuration,
  AssetService,
} from '@bass/bbs-sdk-fetch';
import type { PrepareBadgeImageUploadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AssetService();

  const body = {
    // PrepareBadgeImageUploadReq
    prepareBadgeImageUploadReq: ...,
  } satisfies PrepareBadgeImageUploadRequest;

  try {
    const data = await api.prepareBadgeImageUpload(body);
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
| **prepareBadgeImageUploadReq** | [PrepareBadgeImageUploadReq](PrepareBadgeImageUploadReq.md) |  | |

### Return type

[**PrepareBadgeImageUploadResp**](PrepareBadgeImageUploadResp.md)

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

