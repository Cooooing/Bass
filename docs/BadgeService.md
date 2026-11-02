# BadgeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteDraft**](BadgeService.md#deletedraft) | **POST** /v1/badge/draft/delete |  |
| [**listDrafts**](BadgeService.md#listdrafts) | **POST** /v1/badge/draft/list |  |
| [**renderMaker**](BadgeService.md#rendermaker) | **GET** /v1/badge/maker/render |  |
| [**renderPublished**](BadgeService.md#renderpublished) | **GET** /v1/badge/render |  |
| [**upsertDraft**](BadgeService.md#upsertdraft) | **POST** /v1/badge/draft/upsert |  |



## deleteDraft

> object deleteDraft(deleteBadgeDraftReq)



DeleteDraft 删除一个未发布草稿。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { DeleteDraftRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // DeleteBadgeDraftReq
    deleteBadgeDraftReq: ...,
  } satisfies DeleteDraftRequest;

  try {
    const data = await api.deleteDraft(body);
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


## listDrafts

> ListBadgeDraftsResp listDrafts(body)



ListDrafts 返回当前登录用户的徽章草稿。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { ListDraftsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // object
    body: Object,
  } satisfies ListDraftsRequest;

  try {
    const data = await api.listDrafts(body);
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


## renderMaker

> ImageResp renderMaker(renderSpecImageAssetId, renderSpecText, renderSpecBackgroundColors, renderSpecTextColors, renderSpecBackgroundDirection, renderSpecTextDirection, renderSpecSize, renderSpecBorder, renderSpecTextBarLength, renderSpecTextSize, renderSpecTextBarRadius, renderSpecScale, renderSpecShadow, renderSpecAnimationSeconds, renderSpecFontAssetId)



RenderMaker 仅允许已登录用户预览其提交的完整 Maker 参数。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { RenderMakerRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // string (optional)
    renderSpecImageAssetId: renderSpecImageAssetId_example,
    // string (optional)
    renderSpecText: renderSpecText_example,
    // Array<string> (optional)
    renderSpecBackgroundColors: ...,
    // Array<string> (optional)
    renderSpecTextColors: ...,
    // string (optional)
    renderSpecBackgroundDirection: renderSpecBackgroundDirection_example,
    // string (optional)
    renderSpecTextDirection: renderSpecTextDirection_example,
    // number (optional)
    renderSpecSize: 56,
    // number (optional)
    renderSpecBorder: 56,
    // number (optional)
    renderSpecTextBarLength: 56,
    // number (optional)
    renderSpecTextSize: 56,
    // number (optional)
    renderSpecTextBarRadius: 56,
    // number (optional)
    renderSpecScale: 1.2,
    // number (optional)
    renderSpecShadow: 1.2,
    // number (optional)
    renderSpecAnimationSeconds: 1.2,
    // string (optional)
    renderSpecFontAssetId: renderSpecFontAssetId_example,
  } satisfies RenderMakerRequest;

  try {
    const data = await api.renderMaker(body);
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
| **renderSpecImageAssetId** | `string` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecText** | `string` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecBackgroundColors** | `Array<string>` |  | [Optional] |
| **renderSpecTextColors** | `Array<string>` |  | [Optional] |
| **renderSpecBackgroundDirection** | `string` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecTextDirection** | `string` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecSize** | `number` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecBorder** | `number` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecTextBarLength** | `number` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecTextSize** | `number` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecTextBarRadius** | `number` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecScale** | `number` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecShadow** | `number` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecAnimationSeconds** | `number` |  | [Optional] [Defaults to `undefined`] |
| **renderSpecFontAssetId** | `string` |  | [Optional] [Defaults to `undefined`] |

### Return type

[**ImageResp**](ImageResp.md)

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


## renderPublished

> ImageResp renderPublished(code)



RenderPublished 匿名返回一个已发布徽章的主站 SVG。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { RenderPublishedRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // string (optional)
    code: code_example,
  } satisfies RenderPublishedRequest;

  try {
    const data = await api.renderPublished(body);
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
| **code** | `string` |  | [Optional] [Defaults to `undefined`] |

### Return type

[**ImageResp**](ImageResp.md)

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


## upsertDraft

> UpsertBadgeDraftResp upsertDraft(upsertBadgeDraftReq)



UpsertDraft 保存当前登录用户的徽章草稿。

### Example

```ts
import {
  Configuration,
  BadgeService,
} from '@bass/bbs-sdk-fetch';
import type { UpsertDraftRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new BadgeService();

  const body = {
    // UpsertBadgeDraftReq
    upsertBadgeDraftReq: ...,
  } satisfies UpsertDraftRequest;

  try {
    const data = await api.upsertDraft(body);
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

