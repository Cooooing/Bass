# AssetService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**completeBadgeFontUpload**](#completebadgefontupload) | **POST** /v1/asset/badge-font/complete | |
|[**completeBadgeImageUpload**](#completebadgeimageupload) | **POST** /v1/asset/badge-image/complete | |
|[**completeMarkdownImageUpload**](#completemarkdownimageupload) | **POST** /v1/asset/markdown-image/complete | |
|[**prepareBadgeFontUpload**](#preparebadgefontupload) | **POST** /v1/asset/badge-font/prepare | |
|[**prepareBadgeImageUpload**](#preparebadgeimageupload) | **POST** /v1/asset/badge-image/prepare | |
|[**prepareMarkdownImageUpload**](#preparemarkdownimageupload) | **POST** /v1/asset/markdown-image/prepare | |

# **completeBadgeFontUpload**
> CompleteBadgeFontUploadResp completeBadgeFontUpload(completeBadgeFontUploadReq)

查询已由对象存储回调确认的徽章 Maker 自定义字体。

### Example

```typescript
import {
    AssetService,
    Configuration,
    CompleteBadgeFontUploadReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AssetService(configuration);

let completeBadgeFontUploadReq: CompleteBadgeFontUploadReq; //

const { status, data } = await apiInstance.completeBadgeFontUpload(
    completeBadgeFontUploadReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **completeBadgeFontUploadReq** | **CompleteBadgeFontUploadReq**|  | |


### Return type

**CompleteBadgeFontUploadResp**

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

# **completeBadgeImageUpload**
> CompleteBadgeImageUploadResp completeBadgeImageUpload(completeBadgeImageUploadReq)

查询已由对象存储回调确认的徽章 Maker 图片。

### Example

```typescript
import {
    AssetService,
    Configuration,
    CompleteBadgeImageUploadReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AssetService(configuration);

let completeBadgeImageUploadReq: CompleteBadgeImageUploadReq; //

const { status, data } = await apiInstance.completeBadgeImageUpload(
    completeBadgeImageUploadReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **completeBadgeImageUploadReq** | **CompleteBadgeImageUploadReq**|  | |


### Return type

**CompleteBadgeImageUploadResp**

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

# **completeMarkdownImageUpload**
> CompleteMarkdownImageUploadResp completeMarkdownImageUpload(completeMarkdownImageUploadReq)

查询已由对象存储回调确认的 Markdown 图片。

### Example

```typescript
import {
    AssetService,
    Configuration,
    CompleteMarkdownImageUploadReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AssetService(configuration);

let completeMarkdownImageUploadReq: CompleteMarkdownImageUploadReq; //

const { status, data } = await apiInstance.completeMarkdownImageUpload(
    completeMarkdownImageUploadReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **completeMarkdownImageUploadReq** | **CompleteMarkdownImageUploadReq**|  | |


### Return type

**CompleteMarkdownImageUploadResp**

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

# **prepareBadgeFontUpload**
> PrepareBadgeFontUploadResp prepareBadgeFontUpload(prepareBadgeFontUploadReq)

申请徽章 Maker 自定义字体的内容寻址直传凭证。

### Example

```typescript
import {
    AssetService,
    Configuration,
    PrepareBadgeFontUploadReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AssetService(configuration);

let prepareBadgeFontUploadReq: PrepareBadgeFontUploadReq; //

const { status, data } = await apiInstance.prepareBadgeFontUpload(
    prepareBadgeFontUploadReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **prepareBadgeFontUploadReq** | **PrepareBadgeFontUploadReq**|  | |


### Return type

**PrepareBadgeFontUploadResp**

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

# **prepareBadgeImageUpload**
> PrepareBadgeImageUploadResp prepareBadgeImageUpload(prepareBadgeImageUploadReq)

申请徽章 Maker 图片的内容寻址直传凭证。

### Example

```typescript
import {
    AssetService,
    Configuration,
    PrepareBadgeImageUploadReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AssetService(configuration);

let prepareBadgeImageUploadReq: PrepareBadgeImageUploadReq; //

const { status, data } = await apiInstance.prepareBadgeImageUpload(
    prepareBadgeImageUploadReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **prepareBadgeImageUploadReq** | **PrepareBadgeImageUploadReq**|  | |


### Return type

**PrepareBadgeImageUploadResp**

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

# **prepareMarkdownImageUpload**
> PrepareMarkdownImageUploadResp prepareMarkdownImageUpload(prepareMarkdownImageUploadReq)

申请 Markdown 图片的内容寻址直传凭证。

### Example

```typescript
import {
    AssetService,
    Configuration,
    PrepareMarkdownImageUploadReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AssetService(configuration);

let prepareMarkdownImageUploadReq: PrepareMarkdownImageUploadReq; //

const { status, data } = await apiInstance.prepareMarkdownImageUpload(
    prepareMarkdownImageUploadReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **prepareMarkdownImageUploadReq** | **PrepareMarkdownImageUploadReq**|  | |


### Return type

**PrepareMarkdownImageUploadResp**

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

