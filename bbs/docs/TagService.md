# TagService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**bindArticle**](TagService.md#bindarticle) | **POST** /v1/content/tag/bind-article |  |
| [**create**](TagService.md#create) | **POST** /v1/content/tag/create |  |
| [**list**](TagService.md#list) | **POST** /v1/content/tag/list |  |
| [**listArticleTags**](TagService.md#listarticletags) | **POST** /v1/content/tag/list-article-tags |  |
| [**unbindArticle**](TagService.md#unbindarticle) | **POST** /v1/content/tag/unbind-article |  |
| [**update**](TagService.md#update) | **POST** /v1/content/tag/update |  |



## bindArticle

> object bindArticle(bindArticleTagsReq)



绑定文章标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { BindArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // BindArticleTagsReq
    bindArticleTagsReq: ...,
  } satisfies BindArticleRequest;

  try {
    const data = await api.bindArticle(body);
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


## create

> CreateTagResp create(createTagReq)



创建标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { CreateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // CreateTagReq
    createTagReq: ...,
  } satisfies CreateRequest;

  try {
    const data = await api.create(body);
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


## list

> ListTagsResp list(listTagsReq)



查询标签列表。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { ListRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // ListTagsReq
    listTagsReq: ...,
  } satisfies ListRequest;

  try {
    const data = await api.list(body);
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


## unbindArticle

> object unbindArticle(unbindArticleTagsReq)



解绑文章标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { UnbindArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // UnbindArticleTagsReq
    unbindArticleTagsReq: ...,
  } satisfies UnbindArticleRequest;

  try {
    const data = await api.unbindArticle(body);
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


## update

> UpdateTagResp update(updateTagReq)



更新标签。

### Example

```ts
import {
  Configuration,
  TagService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TagService();

  const body = {
    // UpdateTagReq
    updateTagReq: ...,
  } satisfies UpdateRequest;

  try {
    const data = await api.update(body);
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

