# BadgeService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**deleteBadgeDraft**](#deletebadgedraft) | **POST** /v1/badge/draft/delete | |
|[**listBadgeDrafts**](#listbadgedrafts) | **POST** /v1/badge/draft/list | |
|[**renderMakerBadge**](#rendermakerbadge) | **GET** /v1/badge/maker/render | |
|[**renderPublishedBadge**](#renderpublishedbadge) | **GET** /v1/badge/render | |
|[**upsertBadgeDraft**](#upsertbadgedraft) | **POST** /v1/badge/draft/upsert | |

# **deleteBadgeDraft**
> object deleteBadgeDraft(deleteBadgeDraftReq)

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

const { status, data } = await apiInstance.deleteBadgeDraft(
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

# **listBadgeDrafts**
> ListBadgeDraftsResp listBadgeDrafts(body)

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

const { status, data } = await apiInstance.listBadgeDrafts(
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

# **renderMakerBadge**
> RenderMakerBadgeResp renderMakerBadge()

RenderMaker 仅允许已登录用户预览其提交的完整 Maker 参数。

### Example

```typescript
import {
    BadgeService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new BadgeService(configuration);

let renderSpecImageAssetId: string; //Platform 已确认可用的徽章图标资源 ID。 (optional) (default to undefined)
let renderSpecText: string; //文字条内容。 (optional) (default to undefined)
let renderSpecBackgroundColors: Array<string>; //背景颜色列表。 (optional) (default to undefined)
let renderSpecTextColors: Array<string>; //文字颜色列表。 (optional) (default to undefined)
let renderSpecBackgroundDirection: string; //背景渐变方向或角度。 (optional) (default to undefined)
let renderSpecTextDirection: string; //文字渐变方向或角度。 (optional) (default to undefined)
let renderSpecSize: number; //图标边长。 (optional) (default to undefined)
let renderSpecBorder: number; //图标边距。 (optional) (default to undefined)
let renderSpecTextBarLength: number; //文字条长度。 (optional) (default to undefined)
let renderSpecTextSize: number; //文字字号。 (optional) (default to undefined)
let renderSpecTextBarRadius: number; //文字条圆角。 (optional) (default to undefined)
let renderSpecScale: number; //整体缩放比例。 (optional) (default to undefined)
let renderSpecShadow: number; //阴影强度。 (optional) (default to undefined)
let renderSpecAnimationSeconds: number; //动画时长。 (optional) (default to undefined)
let renderSpecFontAssetId: string; //Platform 已确认可用的自定义字体资源 ID。 (optional) (default to undefined)

const { status, data } = await apiInstance.renderMakerBadge(
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
| **renderSpecImageAssetId** | [**string**] | Platform 已确认可用的徽章图标资源 ID。 | (optional) defaults to undefined|
| **renderSpecText** | [**string**] | 文字条内容。 | (optional) defaults to undefined|
| **renderSpecBackgroundColors** | **Array&lt;string&gt;** | 背景颜色列表。 | (optional) defaults to undefined|
| **renderSpecTextColors** | **Array&lt;string&gt;** | 文字颜色列表。 | (optional) defaults to undefined|
| **renderSpecBackgroundDirection** | [**string**] | 背景渐变方向或角度。 | (optional) defaults to undefined|
| **renderSpecTextDirection** | [**string**] | 文字渐变方向或角度。 | (optional) defaults to undefined|
| **renderSpecSize** | [**number**] | 图标边长。 | (optional) defaults to undefined|
| **renderSpecBorder** | [**number**] | 图标边距。 | (optional) defaults to undefined|
| **renderSpecTextBarLength** | [**number**] | 文字条长度。 | (optional) defaults to undefined|
| **renderSpecTextSize** | [**number**] | 文字字号。 | (optional) defaults to undefined|
| **renderSpecTextBarRadius** | [**number**] | 文字条圆角。 | (optional) defaults to undefined|
| **renderSpecScale** | [**number**] | 整体缩放比例。 | (optional) defaults to undefined|
| **renderSpecShadow** | [**number**] | 阴影强度。 | (optional) defaults to undefined|
| **renderSpecAnimationSeconds** | [**number**] | 动画时长。 | (optional) defaults to undefined|
| **renderSpecFontAssetId** | [**string**] | Platform 已确认可用的自定义字体资源 ID。 | (optional) defaults to undefined|


### Return type

**RenderMakerBadgeResp**

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

# **renderPublishedBadge**
> RenderPublishedBadgeResp renderPublishedBadge()

RenderPublished 匿名返回一个已发布徽章的主站 SVG。

### Example

```typescript
import {
    BadgeService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new BadgeService(configuration);

let code: string; //已发布徽章稳定编码。 (optional) (default to undefined)

const { status, data } = await apiInstance.renderPublishedBadge(
    code
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **code** | [**string**] | 已发布徽章稳定编码。 | (optional) defaults to undefined|


### Return type

**RenderPublishedBadgeResp**

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

# **upsertBadgeDraft**
> UpsertBadgeDraftResp upsertBadgeDraft(upsertBadgeDraftReq)

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

const { status, data } = await apiInstance.upsertBadgeDraft(
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

