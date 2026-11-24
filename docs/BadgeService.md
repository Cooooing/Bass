# BadgeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteBadgeDraft**](BadgeService.md#deletebadgedraft) | **POST** /v1/badge/draft/delete |  |
| [**listBadgeDrafts**](BadgeService.md#listbadgedrafts) | **POST** /v1/badge/draft/list |  |
| [**renderMakerBadge**](BadgeService.md#rendermakerbadge) | **GET** /v1/badge/maker/render |  |
| [**renderPublishedBadge**](BadgeService.md#renderpublishedbadge) | **GET** /v1/badge/render |  |
| [**upsertBadgeDraft**](BadgeService.md#upsertbadgedraft) | **POST** /v1/badge/draft/upsert |  |



## deleteBadgeDraft

> object deleteBadgeDraft(deleteBadgeDraftReq)



DeleteDraft 删除一个未发布草稿。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { DeleteBadgeDraftRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // DeleteBadgeDraftReq
    deleteBadgeDraftReq: ...,
  } satisfies DeleteBadgeDraftRequest;

  try {
    const data = await api.deleteBadgeDraft(body);
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
| **deleteBadgeDraftReq** | [DeleteBadgeDraftReq](DeleteBadgeDraftReq.md) |  | |

### Return type

**object**

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


## listBadgeDrafts

> ListBadgeDraftsResp listBadgeDrafts(body)



ListDrafts 返回当前登录用户的徽章草稿。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { ListBadgeDraftsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // object
    body: Object,
  } satisfies ListBadgeDraftsRequest;

  try {
    const data = await api.listBadgeDrafts(body);
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
| **body** | `object` |  | |

### Return type

[**ListBadgeDraftsResp**](ListBadgeDraftsResp.md)

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


## renderMakerBadge

> RenderMakerBadgeResp renderMakerBadge(renderSpecImageAssetId, renderSpecText, renderSpecBackgroundColors, renderSpecTextColors, renderSpecBackgroundDirection, renderSpecTextDirection, renderSpecSize, renderSpecBorder, renderSpecTextBarLength, renderSpecTextSize, renderSpecTextBarRadius, renderSpecScale, renderSpecShadow, renderSpecAnimationSeconds, renderSpecFontAssetId)



RenderMaker 仅允许已登录用户预览其提交的完整 Maker 参数。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { RenderMakerBadgeRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // string | Platform 已确认可用的徽章图标资源 ID。 (optional)
    renderSpecImageAssetId: renderSpecImageAssetId_example,
    // string | 文字条内容。 (optional)
    renderSpecText: renderSpecText_example,
    // Array<string> | 背景颜色列表。 (optional)
    renderSpecBackgroundColors: ...,
    // Array<string> | 文字颜色列表。 (optional)
    renderSpecTextColors: ...,
    // string | 背景渐变方向或角度。 (optional)
    renderSpecBackgroundDirection: renderSpecBackgroundDirection_example,
    // string | 文字渐变方向或角度。 (optional)
    renderSpecTextDirection: renderSpecTextDirection_example,
    // number | 图标边长。 (optional)
    renderSpecSize: 56,
    // number | 图标边距。 (optional)
    renderSpecBorder: 56,
    // number | 文字条长度。 (optional)
    renderSpecTextBarLength: 56,
    // number | 文字字号。 (optional)
    renderSpecTextSize: 56,
    // number | 文字条圆角。 (optional)
    renderSpecTextBarRadius: 56,
    // number | 整体缩放比例。 (optional)
    renderSpecScale: 1.2,
    // number | 阴影强度。 (optional)
    renderSpecShadow: 1.2,
    // number | 动画时长。 (optional)
    renderSpecAnimationSeconds: 1.2,
    // string | Platform 已确认可用的自定义字体资源 ID。 (optional)
    renderSpecFontAssetId: renderSpecFontAssetId_example,
  } satisfies RenderMakerBadgeRequest;

  try {
    const data = await api.renderMakerBadge(body);
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
| **renderSpecImageAssetId** | `string` | Platform 已确认可用的徽章图标资源 ID。 | [Optional] [Defaults to `undefined`] |
| **renderSpecText** | `string` | 文字条内容。 | [Optional] [Defaults to `undefined`] |
| **renderSpecBackgroundColors** | `Array<string>` | 背景颜色列表。 | [Optional] |
| **renderSpecTextColors** | `Array<string>` | 文字颜色列表。 | [Optional] |
| **renderSpecBackgroundDirection** | `string` | 背景渐变方向或角度。 | [Optional] [Defaults to `undefined`] |
| **renderSpecTextDirection** | `string` | 文字渐变方向或角度。 | [Optional] [Defaults to `undefined`] |
| **renderSpecSize** | `number` | 图标边长。 | [Optional] [Defaults to `undefined`] |
| **renderSpecBorder** | `number` | 图标边距。 | [Optional] [Defaults to `undefined`] |
| **renderSpecTextBarLength** | `number` | 文字条长度。 | [Optional] [Defaults to `undefined`] |
| **renderSpecTextSize** | `number` | 文字字号。 | [Optional] [Defaults to `undefined`] |
| **renderSpecTextBarRadius** | `number` | 文字条圆角。 | [Optional] [Defaults to `undefined`] |
| **renderSpecScale** | `number` | 整体缩放比例。 | [Optional] [Defaults to `undefined`] |
| **renderSpecShadow** | `number` | 阴影强度。 | [Optional] [Defaults to `undefined`] |
| **renderSpecAnimationSeconds** | `number` | 动画时长。 | [Optional] [Defaults to `undefined`] |
| **renderSpecFontAssetId** | `string` | Platform 已确认可用的自定义字体资源 ID。 | [Optional] [Defaults to `undefined`] |

### Return type

[**RenderMakerBadgeResp**](RenderMakerBadgeResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## renderPublishedBadge

> RenderPublishedBadgeResp renderPublishedBadge(code)



RenderPublished 匿名返回一个已发布徽章的主站 SVG。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { RenderPublishedBadgeRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // string | 已发布徽章稳定编码。 (optional)
    code: code_example,
  } satisfies RenderPublishedBadgeRequest;

  try {
    const data = await api.renderPublishedBadge(body);
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
| **code** | `string` | 已发布徽章稳定编码。 | [Optional] [Defaults to `undefined`] |

### Return type

[**RenderPublishedBadgeResp**](RenderPublishedBadgeResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## upsertBadgeDraft

> UpsertBadgeDraftResp upsertBadgeDraft(upsertBadgeDraftReq)



UpsertDraft 保存当前登录用户的徽章草稿。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { UpsertBadgeDraftRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // UpsertBadgeDraftReq
    upsertBadgeDraftReq: ...,
  } satisfies UpsertBadgeDraftRequest;

  try {
    const data = await api.upsertBadgeDraft(body);
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
| **upsertBadgeDraftReq** | [UpsertBadgeDraftReq](UpsertBadgeDraftReq.md) |  | |

### Return type

[**UpsertBadgeDraftResp**](UpsertBadgeDraftResp.md)

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

