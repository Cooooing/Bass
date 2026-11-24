# ArticleService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**acceptArticleAnswer**](ArticleService.md#acceptarticleanswer) | **POST** /v1/content/article/answer/accept |  |
| [**archiveArticle**](ArticleService.md#archivearticle) | **POST** /v1/content/article/archive |  |
| [**cancelPublishArticle**](ArticleService.md#cancelpublisharticle) | **POST** /v1/content/article/publish/cancel |  |
| [**collectArticle**](ArticleService.md#collectarticle) | **POST** /v1/content/article/collect |  |
| [**countCollectedArticles**](ArticleService.md#countcollectedarticles) | **POST** /v1/content/article/collect/count |  |
| [**createDraftArticle**](ArticleService.md#createdraftarticle) | **POST** /v1/content/article/draft/create |  |
| [**discardDraftArticle**](ArticleService.md#discarddraftarticle) | **POST** /v1/content/article/draft/discard |  |
| [**enterArticleLottery**](ArticleService.md#enterarticlelottery) | **POST** /v1/content/article/lottery/entry |  |
| [**getArticle**](ArticleService.md#getarticle) | **POST** /v1/content/article/get |  |
| [**likeArticle**](ArticleService.md#likearticle) | **POST** /v1/content/article/like |  |
| [**listArticles**](ArticleService.md#listarticles) | **POST** /v1/content/article/list |  |
| [**listFollowingArticles**](ArticleService.md#listfollowingarticles) | **POST** /v1/content/article/following/list |  |
| [**listNearbyArticles**](ArticleService.md#listnearbyarticles) | **POST** /v1/content/article/nearby/list |  |
| [**pageArticleViewHistory**](ArticleService.md#pagearticleviewhistory) | **POST** /v1/content/article/view-history/page |  |
| [**pageCollectedArticles**](ArticleService.md#pagecollectedarticles) | **POST** /v1/content/article/collected/page |  |
| [**pageWatchedArticles**](ArticleService.md#pagewatchedarticles) | **POST** /v1/content/article/watched/page |  |
| [**publishArticle**](ArticleService.md#publisharticle) | **POST** /v1/content/article/publish |  |
| [**rewardArticle**](ArticleService.md#rewardarticle) | **POST** /v1/content/article/reward |  |
| [**setPinnedArticleComment**](ArticleService.md#setpinnedarticlecomment) | **POST** /v1/content/article/comment/pin |  |
| [**thankArticle**](ArticleService.md#thankarticle) | **POST** /v1/content/article/thank |  |
| [**updateDraftArticle**](ArticleService.md#updatedraftarticle) | **POST** /v1/content/article/draft/update |  |
| [**voteArticlePoll**](ArticleService.md#votearticlepoll) | **POST** /v1/content/article/poll/vote |  |
| [**watchArticle**](ArticleService.md#watcharticle) | **POST** /v1/content/article/watch |  |



## acceptArticleAnswer

> object acceptArticleAnswer(acceptArticleAnswerReq)



采纳文章回答

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { AcceptArticleAnswerRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // AcceptArticleAnswerReq
    acceptArticleAnswerReq: ...,
  } satisfies AcceptArticleAnswerRequest;

  try {
    const data = await api.acceptArticleAnswer(body);
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
| **acceptArticleAnswerReq** | [AcceptArticleAnswerReq](AcceptArticleAnswerReq.md) |  | |

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


## archiveArticle

> object archiveArticle(archiveArticleReq)



归档文章

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ArchiveArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ArchiveArticleReq
    archiveArticleReq: ...,
  } satisfies ArchiveArticleRequest;

  try {
    const data = await api.archiveArticle(body);
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


## cancelPublishArticle

> object cancelPublishArticle(cancelPublishArticleReq)



取消定时发布

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { CancelPublishArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // CancelPublishArticleReq
    cancelPublishArticleReq: ...,
  } satisfies CancelPublishArticleRequest;

  try {
    const data = await api.cancelPublishArticle(body);
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


## collectArticle

> CollectArticleResp collectArticle(collectArticleReq)



收藏文章

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { CollectArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // CollectArticleReq
    collectArticleReq: ...,
  } satisfies CollectArticleRequest;

  try {
    const data = await api.collectArticle(body);
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


## countCollectedArticles

> CountCollectedArticlesResp countCollectedArticles(body)



统计当前账号收藏的文章数量。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { CountCollectedArticlesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // object
    body: Object,
  } satisfies CountCollectedArticlesRequest;

  try {
    const data = await api.countCollectedArticles(body);
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

[**CountCollectedArticlesResp**](CountCollectedArticlesResp.md)

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


## createDraftArticle

> CreateDraftArticleResp createDraftArticle(createDraftArticleReq)



创建文章草稿

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { CreateDraftArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // CreateDraftArticleReq
    createDraftArticleReq: ...,
  } satisfies CreateDraftArticleRequest;

  try {
    const data = await api.createDraftArticle(body);
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


## discardDraftArticle

> object discardDraftArticle(discardDraftArticleReq)



丢弃文章草稿

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { DiscardDraftArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // DiscardDraftArticleReq
    discardDraftArticleReq: ...,
  } satisfies DiscardDraftArticleRequest;

  try {
    const data = await api.discardDraftArticle(body);
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


## enterArticleLottery

> object enterArticleLottery(enterArticleLotteryReq)



报名或撤销抽奖报名。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { EnterArticleLotteryRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // EnterArticleLotteryReq
    enterArticleLotteryReq: ...,
  } satisfies EnterArticleLotteryRequest;

  try {
    const data = await api.enterArticleLottery(body);
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
| **enterArticleLotteryReq** | [EnterArticleLotteryReq](EnterArticleLotteryReq.md) |  | |

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


## getArticle

> GetArticleResp getArticle(getArticleReq)



查询文章详情

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { GetArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // GetArticleReq
    getArticleReq: ...,
  } satisfies GetArticleRequest;

  try {
    const data = await api.getArticle(body);
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


## likeArticle

> LikeArticleResp likeArticle(likeArticleReq)



点赞文章

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { LikeArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // LikeArticleReq
    likeArticleReq: ...,
  } satisfies LikeArticleRequest;

  try {
    const data = await api.likeArticle(body);
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


## listArticles

> ListArticlesResp listArticles(listArticlesReq)



查询文章列表

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ListArticlesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ListArticlesReq
    listArticlesReq: ...,
  } satisfies ListArticlesRequest;

  try {
    const data = await api.listArticles(body);
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


## listFollowingArticles

> ListFollowingArticlesResp listFollowingArticles(listFollowingArticlesReq)



查询当前账号关注作者发布的文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ListFollowingArticlesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ListFollowingArticlesReq
    listFollowingArticlesReq: ...,
  } satisfies ListFollowingArticlesRequest;

  try {
    const data = await api.listFollowingArticles(body);
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
| **listFollowingArticlesReq** | [ListFollowingArticlesReq](ListFollowingArticlesReq.md) |  | |

### Return type

[**ListFollowingArticlesResp**](ListFollowingArticlesResp.md)

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


## listNearbyArticles

> ListNearbyArticlesResp listNearbyArticles(listNearbyArticlesReq)



查询当前账号所在城市发布的文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ListNearbyArticlesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ListNearbyArticlesReq
    listNearbyArticlesReq: ...,
  } satisfies ListNearbyArticlesRequest;

  try {
    const data = await api.listNearbyArticles(body);
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
| **listNearbyArticlesReq** | [ListNearbyArticlesReq](ListNearbyArticlesReq.md) |  | |

### Return type

[**ListNearbyArticlesResp**](ListNearbyArticlesResp.md)

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


## pageArticleViewHistory

> PageArticleViewHistoryResp pageArticleViewHistory(pageArticleViewHistoryReq)



查询当前账号的浏览历史。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { PageArticleViewHistoryRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // PageArticleViewHistoryReq
    pageArticleViewHistoryReq: ...,
  } satisfies PageArticleViewHistoryRequest;

  try {
    const data = await api.pageArticleViewHistory(body);
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
| **pageArticleViewHistoryReq** | [PageArticleViewHistoryReq](PageArticleViewHistoryReq.md) |  | |

### Return type

[**PageArticleViewHistoryResp**](PageArticleViewHistoryResp.md)

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


## pageCollectedArticles

> PageCollectedArticlesResp pageCollectedArticles(pageCollectedArticlesReq)



查询当前账号收藏的文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { PageCollectedArticlesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // PageCollectedArticlesReq
    pageCollectedArticlesReq: ...,
  } satisfies PageCollectedArticlesRequest;

  try {
    const data = await api.pageCollectedArticles(body);
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
| **pageCollectedArticlesReq** | [PageCollectedArticlesReq](PageCollectedArticlesReq.md) |  | |

### Return type

[**PageCollectedArticlesResp**](PageCollectedArticlesResp.md)

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


## pageWatchedArticles

> PageWatchedArticlesResp pageWatchedArticles(pageWatchedArticlesReq)



查询当前账号关注的文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { PageWatchedArticlesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // PageWatchedArticlesReq
    pageWatchedArticlesReq: ...,
  } satisfies PageWatchedArticlesRequest;

  try {
    const data = await api.pageWatchedArticles(body);
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
| **pageWatchedArticlesReq** | [PageWatchedArticlesReq](PageWatchedArticlesReq.md) |  | |

### Return type

[**PageWatchedArticlesResp**](PageWatchedArticlesResp.md)

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


## publishArticle

> object publishArticle(publishArticleReq)



发布文章

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { PublishArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // PublishArticleReq
    publishArticleReq: ...,
  } satisfies PublishArticleRequest;

  try {
    const data = await api.publishArticle(body);
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


## rewardArticle

> RewardArticleResp rewardArticle(rewardArticleReq)



打赏文章

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { RewardArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // RewardArticleReq
    rewardArticleReq: ...,
  } satisfies RewardArticleRequest;

  try {
    const data = await api.rewardArticle(body);
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

[**RewardArticleResp**](RewardArticleResp.md)

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


## setPinnedArticleComment

> object setPinnedArticleComment(setPinnedArticleCommentReq)



设置或取消当前账号文章的置顶评论。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { SetPinnedArticleCommentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // SetPinnedArticleCommentReq
    setPinnedArticleCommentReq: ...,
  } satisfies SetPinnedArticleCommentRequest;

  try {
    const data = await api.setPinnedArticleComment(body);
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
| **setPinnedArticleCommentReq** | [SetPinnedArticleCommentReq](SetPinnedArticleCommentReq.md) |  | |

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


## thankArticle

> ThankArticleResp thankArticle(thankArticleReq)



感谢文章，确认后不可撤销。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ThankArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ThankArticleReq
    thankArticleReq: ...,
  } satisfies ThankArticleRequest;

  try {
    const data = await api.thankArticle(body);
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


## updateDraftArticle

> UpdateDraftArticleResp updateDraftArticle(updateDraftArticleReq)



编辑文章草稿

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateDraftArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // UpdateDraftArticleReq
    updateDraftArticleReq: ...,
  } satisfies UpdateDraftArticleRequest;

  try {
    const data = await api.updateDraftArticle(body);
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


## voteArticlePoll

> object voteArticlePoll(voteArticlePollReq)



提交一次不可修改的文章投票。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { VoteArticlePollRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // VoteArticlePollReq
    voteArticlePollReq: ...,
  } satisfies VoteArticlePollRequest;

  try {
    const data = await api.voteArticlePoll(body);
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
| **voteArticlePollReq** | [VoteArticlePollReq](VoteArticlePollReq.md) |  | |

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


## watchArticle

> WatchArticleResp watchArticle(watchArticleReq)



关注文章

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { WatchArticleRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // WatchArticleReq
    watchArticleReq: ...,
  } satisfies WatchArticleRequest;

  try {
    const data = await api.watchArticle(body);
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
| **watchArticleReq** | [WatchArticleReq](WatchArticleReq.md) |  | |

### Return type

[**WatchArticleResp**](WatchArticleResp.md)

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

