# BadgeService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**deleteDraft**](#deletedraft) | **POST** /v1/badge/draft/delete | |
|[**listDrafts**](#listdrafts) | **POST** /v1/badge/draft/list | |
|[**renderMaker**](#rendermaker) | **GET** /v1/badge/maker/render | |
|[**renderPublished**](#renderpublished) | **GET** /v1/badge/render | |
|[**upsertDraft**](#upsertdraft) | **POST** /v1/badge/draft/upsert | |

# **deleteDraft**
> object deleteDraft(deleteBadgeDraftReq)

DeleteDraft 删除一个未发布草稿。

### Example

```typescript
import {
    BadgeService,
    Configuration,
    DeleteBadgeDraftReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new BadgeService(configuration);

let deleteBadgeDraftReq: DeleteBadgeDraftReq; //

const { status, data } = await apiInstance.deleteDraft(
    deleteBadgeDraftReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **deleteBadgeDraftReq** | **DeleteBadgeDraftReq**|  | |


### Return type

**object**

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

# **listDrafts**
> ListBadgeDraftsResp listDrafts(body)

ListDrafts 返回当前登录用户的徽章草稿。

### Example

```typescript
import {
    BadgeService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new BadgeService(configuration);

let body: object; //

const { status, data } = await apiInstance.listDrafts(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**ListBadgeDraftsResp**

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

# **renderMaker**
> ImageResp renderMaker()

RenderMaker 仅允许已登录用户预览其提交的完整 Maker 参数。

### Example

```typescript
import {
    BadgeService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new BadgeService(configuration);

let renderSpecImageAssetId: string; // (optional) (default to undefined)
let renderSpecText: string; // (optional) (default to undefined)
let renderSpecBackgroundColors: Array<string>; // (optional) (default to undefined)
let renderSpecTextColors: Array<string>; // (optional) (default to undefined)
let renderSpecBackgroundDirection: string; // (optional) (default to undefined)
let renderSpecTextDirection: string; // (optional) (default to undefined)
let renderSpecSize: number; // (optional) (default to undefined)
let renderSpecBorder: number; // (optional) (default to undefined)
let renderSpecTextBarLength: number; // (optional) (default to undefined)
let renderSpecTextSize: number; // (optional) (default to undefined)
let renderSpecTextBarRadius: number; // (optional) (default to undefined)
let renderSpecScale: number; // (optional) (default to undefined)
let renderSpecShadow: number; // (optional) (default to undefined)
let renderSpecAnimationSeconds: number; // (optional) (default to undefined)
let renderSpecFontAssetId: string; // (optional) (default to undefined)

const { status, data } = await apiInstance.renderMaker(
    renderSpecImageAssetId,
    renderSpecText,
    renderSpecBackgroundColors,
    renderSpecTextColors,
    renderSpecBackgroundDirection,
    renderSpecTextDirection,
    renderSpecSize,
    renderSpecBorder,
    renderSpecTextBarLength,
    renderSpecTextSize,
    renderSpecTextBarRadius,
    renderSpecScale,
    renderSpecShadow,
    renderSpecAnimationSeconds,
    renderSpecFontAssetId
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **renderSpecImageAssetId** | [**string**] |  | (optional) defaults to undefined|
| **renderSpecText** | [**string**] |  | (optional) defaults to undefined|
| **renderSpecBackgroundColors** | **Array&lt;string&gt;** |  | (optional) defaults to undefined|
| **renderSpecTextColors** | **Array&lt;string&gt;** |  | (optional) defaults to undefined|
| **renderSpecBackgroundDirection** | [**string**] |  | (optional) defaults to undefined|
| **renderSpecTextDirection** | [**string**] |  | (optional) defaults to undefined|
| **renderSpecSize** | [**number**] |  | (optional) defaults to undefined|
| **renderSpecBorder** | [**number**] |  | (optional) defaults to undefined|
| **renderSpecTextBarLength** | [**number**] |  | (optional) defaults to undefined|
| **renderSpecTextSize** | [**number**] |  | (optional) defaults to undefined|
| **renderSpecTextBarRadius** | [**number**] |  | (optional) defaults to undefined|
| **renderSpecScale** | [**number**] |  | (optional) defaults to undefined|
| **renderSpecShadow** | [**number**] |  | (optional) defaults to undefined|
| **renderSpecAnimationSeconds** | [**number**] |  | (optional) defaults to undefined|
| **renderSpecFontAssetId** | [**string**] |  | (optional) defaults to undefined|


### Return type

**ImageResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **renderPublished**
> ImageResp renderPublished()

RenderPublished 匿名返回一个已发布徽章的主站 SVG。

### Example

```typescript
import {
    BadgeService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new BadgeService(configuration);

let code: string; // (optional) (default to undefined)

const { status, data } = await apiInstance.renderPublished(
    code
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **code** | [**string**] |  | (optional) defaults to undefined|


### Return type

**ImageResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upsertDraft**
> UpsertBadgeDraftResp upsertDraft(upsertBadgeDraftReq)

UpsertDraft 保存当前登录用户的徽章草稿。

### Example

```typescript
import {
    BadgeService,
    Configuration,
    UpsertBadgeDraftReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new BadgeService(configuration);

let upsertBadgeDraftReq: UpsertBadgeDraftReq; //

const { status, data } = await apiInstance.upsertDraft(
    upsertBadgeDraftReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **upsertBadgeDraftReq** | **UpsertBadgeDraftReq**|  | |


### Return type

**UpsertBadgeDraftResp**

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

