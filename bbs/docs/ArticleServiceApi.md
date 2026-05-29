# ArticleServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**articleServiceAcceptAnswer**](ArticleServiceApi.md#articleserviceacceptanswer) | **POST** /v1/content/article/accept-answer |  |
| [**articleServiceCollect**](ArticleServiceApi.md#articleservicecollect) | **POST** /v1/content/article/collect |  |
| [**articleServiceCreate**](ArticleServiceApi.md#articleservicecreate) | **POST** /v1/content/article/create |  |
| [**articleServiceDelete**](ArticleServiceApi.md#articleservicedelete) | **POST** /v1/content/article/delete |  |
| [**articleServiceGet**](ArticleServiceApi.md#articleserviceget) | **POST** /v1/content/article/get |  |
| [**articleServiceLike**](ArticleServiceApi.md#articleservicelike) | **POST** /v1/content/article/like |  |
| [**articleServiceList**](ArticleServiceApi.md#articleservicelist) | **POST** /v1/content/article/list |  |
| [**articleServicePublish**](ArticleServiceApi.md#articleservicepublish) | **POST** /v1/content/article/publish |  |
| [**articleServiceReward**](ArticleServiceApi.md#articleservicereward) | **POST** /v1/content/article/reward |  |
| [**articleServiceThank**](ArticleServiceApi.md#articleservicethank) | **POST** /v1/content/article/thank |  |
| [**articleServiceUpdateDraft**](ArticleServiceApi.md#articleserviceupdatedraft) | **POST** /v1/content/article/update-draft |  |
| [**articleServiceWatch**](ArticleServiceApi.md#articleservicewatch) | **POST** /v1/content/article/watch |  |



## articleServiceAcceptAnswer

> object articleServiceAcceptAnswer(acceptAnswerArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceAcceptAnswerRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // AcceptAnswerArticleRequest
    acceptAnswerArticleRequest: ...,
  } satisfies ArticleServiceAcceptAnswerRequest;

  try {
    const data = await api.articleServiceAcceptAnswer(body);
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


## articleServiceCollect

> object articleServiceCollect(collectArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceCollectRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // CollectArticleRequest
    collectArticleRequest: ...,
  } satisfies ArticleServiceCollectRequest;

  try {
    const data = await api.articleServiceCollect(body);
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


## articleServiceCreate

> CreateArticleReply articleServiceCreate(createArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceCreateRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // CreateArticleRequest
    createArticleRequest: ...,
  } satisfies ArticleServiceCreateRequest;

  try {
    const data = await api.articleServiceCreate(body);
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


## articleServiceDelete

> object articleServiceDelete(deleteArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceDeleteRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // DeleteArticleRequest
    deleteArticleRequest: ...,
  } satisfies ArticleServiceDeleteRequest;

  try {
    const data = await api.articleServiceDelete(body);
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


## articleServiceGet

> GetArticleReply articleServiceGet(getArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceGetRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // GetArticleRequest
    getArticleRequest: ...,
  } satisfies ArticleServiceGetRequest;

  try {
    const data = await api.articleServiceGet(body);
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


## articleServiceLike

> object articleServiceLike(likeArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceLikeRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // LikeArticleRequest
    likeArticleRequest: ...,
  } satisfies ArticleServiceLikeRequest;

  try {
    const data = await api.articleServiceLike(body);
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


## articleServiceList

> ListArticlesReply articleServiceList(listArticlesRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceListRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // ListArticlesRequest
    listArticlesRequest: ...,
  } satisfies ArticleServiceListRequest;

  try {
    const data = await api.articleServiceList(body);
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


## articleServicePublish

> object articleServicePublish(publishArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServicePublishRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // PublishArticleRequest
    publishArticleRequest: ...,
  } satisfies ArticleServicePublishRequest;

  try {
    const data = await api.articleServicePublish(body);
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


## articleServiceReward

> object articleServiceReward(rewardArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceRewardRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // RewardArticleRequest
    rewardArticleRequest: ...,
  } satisfies ArticleServiceRewardRequest;

  try {
    const data = await api.articleServiceReward(body);
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


## articleServiceThank

> object articleServiceThank(thankArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceThankRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // ThankArticleRequest
    thankArticleRequest: ...,
  } satisfies ArticleServiceThankRequest;

  try {
    const data = await api.articleServiceThank(body);
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


## articleServiceUpdateDraft

> UpdateDraftArticleReply articleServiceUpdateDraft(updateDraftArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceUpdateDraftRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // UpdateDraftArticleRequest
    updateDraftArticleRequest: ...,
  } satisfies ArticleServiceUpdateDraftRequest;

  try {
    const data = await api.articleServiceUpdateDraft(body);
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


## articleServiceWatch

> object articleServiceWatch(watchArticleRequest)



### Example

```ts
import {
  Configuration,
  ArticleServiceApi,
} from '@bass/bbs-sdk';
import type { ArticleServiceWatchRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new ArticleServiceApi();

  const body = {
    // WatchArticleRequest
    watchArticleRequest: ...,
  } satisfies ArticleServiceWatchRequest;

  try {
    const data = await api.articleServiceWatch(body);
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

