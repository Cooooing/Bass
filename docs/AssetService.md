# AssetService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**completeMarkdownImageUpload**](#completemarkdownimageupload) | **POST** /v1/asset/markdown-image/complete | |
|[**prepareMarkdownImageUpload**](#preparemarkdownimageupload) | **POST** /v1/asset/markdown-image/prepare | |

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

