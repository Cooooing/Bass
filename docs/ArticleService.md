# ArticleService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**acceptAnswer**](ArticleService.md#acceptanswer) | **POST** /v1/content/article/answer/accept |  |
| [**archive**](ArticleService.md#archive) | **POST** /v1/content/article/archive |  |
| [**cancelPublish**](ArticleService.md#cancelpublish) | **POST** /v1/content/article/publish/cancel |  |
| [**collect**](ArticleService.md#collect) | **POST** /v1/content/article/collect |  |
| [**createDraft**](ArticleService.md#createdraft) | **POST** /v1/content/article/draft/create |  |
| [**discardDraft**](ArticleService.md#discarddraft) | **POST** /v1/content/article/draft/discard |  |
| [**enterLottery**](ArticleService.md#enterlottery) | **POST** /v1/content/article/lottery/entry |  |
| [**get**](ArticleService.md#get) | **POST** /v1/content/article/get |  |
| [**like**](ArticleService.md#like) | **POST** /v1/content/article/like |  |
| [**list**](ArticleService.md#list) | **POST** /v1/content/article/list |  |
| [**listFollowing**](ArticleService.md#listfollowing) | **POST** /v1/content/article/following/list |  |
| [**listNearby**](ArticleService.md#listnearby) | **POST** /v1/content/article/nearby/list |  |
| [**pageViewHistory**](ArticleService.md#pageviewhistory) | **POST** /v1/content/article/view-history/page |  |
| [**publish**](ArticleService.md#publish) | **POST** /v1/content/article/publish |  |
| [**reward**](ArticleService.md#reward) | **POST** /v1/content/article/reward |  |
| [**setPinnedComment**](ArticleService.md#setpinnedcomment) | **POST** /v1/content/article/comment/pin |  |
| [**thank**](ArticleService.md#thank) | **POST** /v1/content/article/thank |  |
| [**updateDraft**](ArticleService.md#updatedraft) | **POST** /v1/content/article/draft/update |  |
| [**votePoll**](ArticleService.md#votepoll) | **POST** /v1/content/article/poll/vote |  |
| [**watch**](ArticleService.md#watch) | **POST** /v1/content/article/watch |  |



## acceptAnswer

> object acceptAnswer(acceptArticleAnswerReq)



采纳文章回答

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
    // AcceptArticleAnswerReq
    acceptArticleAnswerReq: ...,
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


## enterLottery

> object enterLottery(enterArticleLotteryReq)



报名或撤销抽奖报名。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { EnterLotteryRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // EnterArticleLotteryReq
    enterArticleLotteryReq: ...,
  } satisfies EnterLotteryRequest;

  try {
    const data = await api.enterLottery(body);
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


## listFollowing

> ListFollowingArticlesResp listFollowing(listFollowingArticlesReq)



查询当前账号关注作者发布的文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ListFollowingRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ListFollowingArticlesReq
    listFollowingArticlesReq: ...,
  } satisfies ListFollowingRequest;

  try {
    const data = await api.listFollowing(body);
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


## listNearby

> ListNearbyArticlesResp listNearby(listNearbyArticlesReq)



查询当前账号所在城市发布的文章。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { ListNearbyRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // ListNearbyArticlesReq
    listNearbyArticlesReq: ...,
  } satisfies ListNearbyRequest;

  try {
    const data = await api.listNearby(body);
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


## pageViewHistory

> PageArticleViewHistoryResp pageViewHistory(pageArticleViewHistoryReq)



查询当前账号的浏览历史。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { PageViewHistoryRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // PageArticleViewHistoryReq
    pageArticleViewHistoryReq: ...,
  } satisfies PageViewHistoryRequest;

  try {
    const data = await api.pageViewHistory(body);
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


## setPinnedComment

> object setPinnedComment(setPinnedArticleCommentReq)



设置或取消当前账号文章的置顶评论。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { SetPinnedCommentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // SetPinnedArticleCommentReq
    setPinnedArticleCommentReq: ...,
  } satisfies SetPinnedCommentRequest;

  try {
    const data = await api.setPinnedComment(body);
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


## votePoll

> object votePoll(voteArticlePollReq)



提交一次不可修改的文章投票。

### Example

```ts
import {
  Configuration,
  ArticleService,
} from '@bass/bbs-sdk-fetch';
import type { VotePollRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new ArticleService();

  const body = {
    // VoteArticlePollReq
    voteArticlePollReq: ...,
  } satisfies VotePollRequest;

  try {
    const data = await api.votePoll(body);
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


## watch

> WatchArticleResp watch(watchArticleReq)



关注文章

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
    // WatchArticleReq
    watchArticleReq: ...,
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

