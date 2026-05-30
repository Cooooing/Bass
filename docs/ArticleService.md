# ArticleService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**_delete**](ArticleService.md#_delete) | **POST** /v1/content/article/delete |  |
| [**acceptAnswer**](ArticleService.md#acceptanswer) | **POST** /v1/content/article/accept-answer |  |
| [**collect**](ArticleService.md#collect) | **POST** /v1/content/article/collect |  |
| [**create**](ArticleService.md#create) | **POST** /v1/content/article/create |  |
| [**get**](ArticleService.md#get) | **POST** /v1/content/article/get |  |
| [**like**](ArticleService.md#like) | **POST** /v1/content/article/like |  |
| [**list**](ArticleService.md#list) | **POST** /v1/content/article/list |  |
| [**publish**](ArticleService.md#publish) | **POST** /v1/content/article/publish |  |
| [**reward**](ArticleService.md#reward) | **POST** /v1/content/article/reward |  |
| [**thank**](ArticleService.md#thank) | **POST** /v1/content/article/thank |  |
| [**updateDraft**](ArticleService.md#updatedraft) | **POST** /v1/content/article/update-draft |  |
| [**watch**](ArticleService.md#watch) | **POST** /v1/content/article/watch |  |



## _delete

> object _delete(deleteArticleRequest)



删除文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { DeleteRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // DeleteArticleRequest
    deleteArticleRequest: ...,
  } satisfies DeleteRequest;

  try {
    const data = await api._delete(body);
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
| **deleteArticleRequest** | [DeleteArticleRequest](DeleteArticleRequest.md) |  | |

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


## acceptAnswer

> object acceptAnswer(acceptAnswerArticleRequest)



采纳文章评论为答案。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { AcceptAnswerRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // AcceptAnswerArticleRequest
    acceptAnswerArticleRequest: ...,
  } satisfies AcceptAnswerRequest;

  try {
    const data = await api.acceptAnswer(body);
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
| **acceptAnswerArticleRequest** | [AcceptAnswerArticleRequest](AcceptAnswerArticleRequest.md) |  | |

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


## collect

> object collect(collectArticleRequest)



收藏或取消收藏文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { CollectRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // CollectArticleRequest
    collectArticleRequest: ...,
  } satisfies CollectRequest;

  try {
    const data = await api.collect(body);
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
| **collectArticleRequest** | [CollectArticleRequest](CollectArticleRequest.md) |  | |

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

> CreateArticleReply create(createArticleRequest)



创建文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { CreateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // CreateArticleRequest
    createArticleRequest: ...,
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
| **createArticleRequest** | [CreateArticleRequest](CreateArticleRequest.md) |  | |

### Return type

[**CreateArticleReply**](CreateArticleReply.md)

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


## get

> GetArticleReply get(getArticleRequest)



获取文章详情。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { GetRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // GetArticleRequest
    getArticleRequest: ...,
  } satisfies GetRequest;

  try {
    const data = await api.get(body);
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
| **getArticleRequest** | [GetArticleRequest](GetArticleRequest.md) |  | |

### Return type

[**GetArticleReply**](GetArticleReply.md)

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


## like

> object like(likeArticleRequest)



点赞或取消点赞文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { LikeRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // LikeArticleRequest
    likeArticleRequest: ...,
  } satisfies LikeRequest;

  try {
    const data = await api.like(body);
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
| **likeArticleRequest** | [LikeArticleRequest](LikeArticleRequest.md) |  | |

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


## list

> ListArticlesReply list(listArticlesRequest)



分页查询文章列表。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ListRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ListArticlesRequest
    listArticlesRequest: ...,
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
| **listArticlesRequest** | [ListArticlesRequest](ListArticlesRequest.md) |  | |

### Return type

[**ListArticlesReply**](ListArticlesReply.md)

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


## publish

> object publish(publishArticleRequest)



发布文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { PublishRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // PublishArticleRequest
    publishArticleRequest: ...,
  } satisfies PublishRequest;

  try {
    const data = await api.publish(body);
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
| **publishArticleRequest** | [PublishArticleRequest](PublishArticleRequest.md) |  | |

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


## reward

> object reward(rewardArticleRequest)



打赏文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { RewardRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // RewardArticleRequest
    rewardArticleRequest: ...,
  } satisfies RewardRequest;

  try {
    const data = await api.reward(body);
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
| **rewardArticleRequest** | [RewardArticleRequest](RewardArticleRequest.md) |  | |

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


## thank

> object thank(thankArticleRequest)



感谢或取消感谢文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ThankRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ThankArticleRequest
    thankArticleRequest: ...,
  } satisfies ThankRequest;

  try {
    const data = await api.thank(body);
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
| **thankArticleRequest** | [ThankArticleRequest](ThankArticleRequest.md) |  | |

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


## updateDraft

> UpdateDraftArticleReply updateDraft(updateDraftArticleRequest)



更新文章草稿。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateDraftRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // UpdateDraftArticleRequest
    updateDraftArticleRequest: ...,
  } satisfies UpdateDraftRequest;

  try {
    const data = await api.updateDraft(body);
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
| **updateDraftArticleRequest** | [UpdateDraftArticleRequest](UpdateDraftArticleRequest.md) |  | |

### Return type

[**UpdateDraftArticleReply**](UpdateDraftArticleReply.md)

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


## watch

> object watch(watchArticleRequest)



关注或取消关注文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { WatchRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // WatchArticleRequest
    watchArticleRequest: ...,
  } satisfies WatchRequest;

  try {
    const data = await api.watch(body);
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
| **watchArticleRequest** | [WatchArticleRequest](WatchArticleRequest.md) |  | |

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

