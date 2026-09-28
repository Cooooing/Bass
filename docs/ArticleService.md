# ArticleService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**archive**](ArticleService.md#archive) | **POST** /v1/content/article/archive |  |
| [**cancelPublish**](ArticleService.md#cancelpublish) | **POST** /v1/content/article/publish/cancel |  |
| [**collect**](ArticleService.md#collect) | **POST** /v1/content/article/collect |  |
| [**createDraft**](ArticleService.md#createdraft) | **POST** /v1/content/article/draft/create |  |
| [**discardDraft**](ArticleService.md#discarddraft) | **POST** /v1/content/article/draft/discard |  |
| [**get**](ArticleService.md#get) | **POST** /v1/content/article/get |  |
| [**like**](ArticleService.md#like) | **POST** /v1/content/article/like |  |
| [**list**](ArticleService.md#list) | **POST** /v1/content/article/list |  |
| [**publish**](ArticleService.md#publish) | **POST** /v1/content/article/publish |  |
| [**reward**](ArticleService.md#reward) | **POST** /v1/content/article/reward |  |
| [**schedulePublish**](ArticleService.md#schedulepublish) | **POST** /v1/content/article/publish/schedule |  |
| [**thank**](ArticleService.md#thank) | **POST** /v1/content/article/thank |  |
| [**updateDraft**](ArticleService.md#updatedraft) | **POST** /v1/content/article/draft/update |  |



## archive

> object archive(archiveArticleReq)



归档文章

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ArchiveRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ArchiveArticleReq
    archiveArticleReq: ...,
  } satisfies ArchiveRequest;

  try {
    const data = await api.archive(body);
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
| **archiveArticleReq** | [ArchiveArticleReq](ArchiveArticleReq.md) |  | |

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


## cancelPublish

> object cancelPublish(cancelPublishArticleReq)



取消定时发布

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { CancelPublishRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // CancelPublishArticleReq
    cancelPublishArticleReq: ...,
  } satisfies CancelPublishRequest;

  try {
    const data = await api.cancelPublish(body);
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
| **cancelPublishArticleReq** | [CancelPublishArticleReq](CancelPublishArticleReq.md) |  | |

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

> CollectArticleResp collect(collectArticleReq)



收藏文章

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
    // CollectArticleReq
    collectArticleReq: ...,
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
| **collectArticleReq** | [CollectArticleReq](CollectArticleReq.md) |  | |

### Return type

[**CollectArticleResp**](CollectArticleResp.md)

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


## createDraft

> CreateDraftArticleResp createDraft(createDraftArticleReq)



创建文章草稿

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { CreateDraftRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // CreateDraftArticleReq
    createDraftArticleReq: ...,
  } satisfies CreateDraftRequest;

  try {
    const data = await api.createDraft(body);
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
| **createDraftArticleReq** | [CreateDraftArticleReq](CreateDraftArticleReq.md) |  | |

### Return type

[**CreateDraftArticleResp**](CreateDraftArticleResp.md)

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


## discardDraft

> object discardDraft(discardDraftArticleReq)



丢弃文章草稿

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { DiscardDraftRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // DiscardDraftArticleReq
    discardDraftArticleReq: ...,
  } satisfies DiscardDraftRequest;

  try {
    const data = await api.discardDraft(body);
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
| **discardDraftArticleReq** | [DiscardDraftArticleReq](DiscardDraftArticleReq.md) |  | |

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


## get

> GetArticleResp get(getArticleReq)



查询文章详情

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
    // GetArticleReq
    getArticleReq: ...,
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
| **getArticleReq** | [GetArticleReq](GetArticleReq.md) |  | |

### Return type

[**GetArticleResp**](GetArticleResp.md)

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

> LikeArticleResp like(likeArticleReq)



点赞文章

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
    // LikeArticleReq
    likeArticleReq: ...,
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
| **likeArticleReq** | [LikeArticleReq](LikeArticleReq.md) |  | |

### Return type

[**LikeArticleResp**](LikeArticleResp.md)

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

> ListArticlesResp list(listArticlesReq)



查询文章列表

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
    // ListArticlesReq
    listArticlesReq: ...,
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
| **listArticlesReq** | [ListArticlesReq](ListArticlesReq.md) |  | |

### Return type

[**ListArticlesResp**](ListArticlesResp.md)

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

> object publish(publishArticleReq)



发布文章

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
    // PublishArticleReq
    publishArticleReq: ...,
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
| **publishArticleReq** | [PublishArticleReq](PublishArticleReq.md) |  | |

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

> object reward(rewardArticleReq)



打赏文章

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
    // RewardArticleReq
    rewardArticleReq: ...,
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
| **rewardArticleReq** | [RewardArticleReq](RewardArticleReq.md) |  | |

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


## schedulePublish

> object schedulePublish(schedulePublishArticleReq)



设置定时发布

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { SchedulePublishRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // SchedulePublishArticleReq
    schedulePublishArticleReq: ...,
  } satisfies SchedulePublishRequest;

  try {
    const data = await api.schedulePublish(body);
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
| **schedulePublishArticleReq** | [SchedulePublishArticleReq](SchedulePublishArticleReq.md) |  | |

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

> ThankArticleResp thank(thankArticleReq)



感谢文章

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
    // ThankArticleReq
    thankArticleReq: ...,
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
| **thankArticleReq** | [ThankArticleReq](ThankArticleReq.md) |  | |

### Return type

[**ThankArticleResp**](ThankArticleResp.md)

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

> UpdateDraftArticleResp updateDraft(updateDraftArticleReq)



编辑文章草稿

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
    // UpdateDraftArticleReq
    updateDraftArticleReq: ...,
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
| **updateDraftArticleReq** | [UpdateDraftArticleReq](UpdateDraftArticleReq.md) |  | |

### Return type

[**UpdateDraftArticleResp**](UpdateDraftArticleResp.md)

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

