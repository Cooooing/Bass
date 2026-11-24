# TagService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**bindArticleTags**](TagService.md#bindarticletags) | **POST** /v1/content/tag/bind-article |  |
| [**createTag**](TagService.md#createtag) | **POST** /v1/content/tag/create |  |
| [**listArticleTags**](TagService.md#listarticletags) | **POST** /v1/content/tag/list-article-tags |  |
| [**listTags**](TagService.md#listtags) | **POST** /v1/content/tag/list |  |
| [**unbindArticleTags**](TagService.md#unbindarticletags) | **POST** /v1/content/tag/unbind-article |  |
| [**updateTag**](TagService.md#updatetag) | **POST** /v1/content/tag/update |  |



## bindArticleTags

> object bindArticleTags(bindArticleTagsReq)



绑定文章标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { BindArticleTagsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // BindArticleTagsReq
    bindArticleTagsReq: ...,
  } satisfies BindArticleTagsRequest;

  try {
    const data = await api.bindArticleTags(body);
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
| **bindArticleTagsReq** | [BindArticleTagsReq](BindArticleTagsReq.md) |  | |

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


## createTag

> CreateTagResp createTag(createTagReq)



创建标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { CreateTagRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // CreateTagReq
    createTagReq: ...,
  } satisfies CreateTagRequest;

  try {
    const data = await api.createTag(body);
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
| **createTagReq** | [CreateTagReq](CreateTagReq.md) |  | |

### Return type

[**CreateTagResp**](CreateTagResp.md)

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


## listArticleTags

> ListArticleTagsResp listArticleTags(listArticleTagsReq)



查询文章标签列表。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { ListArticleTagsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // ListArticleTagsReq
    listArticleTagsReq: ...,
  } satisfies ListArticleTagsRequest;

  try {
    const data = await api.listArticleTags(body);
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
| **listArticleTagsReq** | [ListArticleTagsReq](ListArticleTagsReq.md) |  | |

### Return type

[**ListArticleTagsResp**](ListArticleTagsResp.md)

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


## listTags

> ListTagsResp listTags(listTagsReq)



查询标签列表。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { ListTagsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // ListTagsReq
    listTagsReq: ...,
  } satisfies ListTagsRequest;

  try {
    const data = await api.listTags(body);
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
| **listTagsReq** | [ListTagsReq](ListTagsReq.md) |  | |

### Return type

[**ListTagsResp**](ListTagsResp.md)

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


## unbindArticleTags

> object unbindArticleTags(unbindArticleTagsReq)



解绑文章标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { UnbindArticleTagsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // UnbindArticleTagsReq
    unbindArticleTagsReq: ...,
  } satisfies UnbindArticleTagsRequest;

  try {
    const data = await api.unbindArticleTags(body);
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
| **unbindArticleTagsReq** | [UnbindArticleTagsReq](UnbindArticleTagsReq.md) |  | |

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


## updateTag

> UpdateTagResp updateTag(updateTagReq)



更新标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateTagRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // UpdateTagReq
    updateTagReq: ...,
  } satisfies UpdateTagRequest;

  try {
    const data = await api.updateTag(body);
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
| **updateTagReq** | [UpdateTagReq](UpdateTagReq.md) |  | |

### Return type

[**UpdateTagResp**](UpdateTagResp.md)

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

