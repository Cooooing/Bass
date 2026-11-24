# ArticleService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**acceptArticleAnswer**](#acceptarticleanswer) | **POST** /v1/content/article/answer/accept | |
|[**archiveArticle**](#archivearticle) | **POST** /v1/content/article/archive | |
|[**cancelPublishArticle**](#cancelpublisharticle) | **POST** /v1/content/article/publish/cancel | |
|[**collectArticle**](#collectarticle) | **POST** /v1/content/article/collect | |
|[**countCollectedArticles**](#countcollectedarticles) | **POST** /v1/content/article/collect/count | |
|[**createDraftArticle**](#createdraftarticle) | **POST** /v1/content/article/draft/create | |
|[**discardDraftArticle**](#discarddraftarticle) | **POST** /v1/content/article/draft/discard | |
|[**enterArticleLottery**](#enterarticlelottery) | **POST** /v1/content/article/lottery/entry | |
|[**getArticle**](#getarticle) | **POST** /v1/content/article/get | |
|[**likeArticle**](#likearticle) | **POST** /v1/content/article/like | |
|[**listArticles**](#listarticles) | **POST** /v1/content/article/list | |
|[**listFollowingArticles**](#listfollowingarticles) | **POST** /v1/content/article/following/list | |
|[**listNearbyArticles**](#listnearbyarticles) | **POST** /v1/content/article/nearby/list | |
|[**pageArticleViewHistory**](#pagearticleviewhistory) | **POST** /v1/content/article/view-history/page | |
|[**pageCollectedArticles**](#pagecollectedarticles) | **POST** /v1/content/article/collected/page | |
|[**pageWatchedArticles**](#pagewatchedarticles) | **POST** /v1/content/article/watched/page | |
|[**publishArticle**](#publisharticle) | **POST** /v1/content/article/publish | |
|[**rewardArticle**](#rewardarticle) | **POST** /v1/content/article/reward | |
|[**setPinnedArticleComment**](#setpinnedarticlecomment) | **POST** /v1/content/article/comment/pin | |
|[**thankArticle**](#thankarticle) | **POST** /v1/content/article/thank | |
|[**updateDraftArticle**](#updatedraftarticle) | **POST** /v1/content/article/draft/update | |
|[**voteArticlePoll**](#votearticlepoll) | **POST** /v1/content/article/poll/vote | |
|[**watchArticle**](#watcharticle) | **POST** /v1/content/article/watch | |

# **acceptArticleAnswer**
> object acceptArticleAnswer(acceptArticleAnswerReq)

采纳文章回答

### Example

```typescript
import {
    ArticleService,
    Configuration,
    AcceptArticleAnswerReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let acceptArticleAnswerReq: AcceptArticleAnswerReq; //

const { status, data } = await apiInstance.acceptArticleAnswer(
    acceptArticleAnswerReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **acceptArticleAnswerReq** | **AcceptArticleAnswerReq**|  | |


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

# **archiveArticle**
> object archiveArticle(archiveArticleReq)

归档文章

### Example

```typescript
import {
    ArticleService,
    Configuration,
    ArchiveArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let archiveArticleReq: ArchiveArticleReq; //

const { status, data } = await apiInstance.archiveArticle(
    archiveArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **archiveArticleReq** | **ArchiveArticleReq**|  | |


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

# **cancelPublishArticle**
> object cancelPublishArticle(cancelPublishArticleReq)

取消定时发布

### Example

```typescript
import {
    ArticleService,
    Configuration,
    CancelPublishArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let cancelPublishArticleReq: CancelPublishArticleReq; //

const { status, data } = await apiInstance.cancelPublishArticle(
    cancelPublishArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **cancelPublishArticleReq** | **CancelPublishArticleReq**|  | |


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

# **collectArticle**
> CollectArticleResp collectArticle(collectArticleReq)

收藏文章

### Example

```typescript
import {
    ArticleService,
    Configuration,
    CollectArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let collectArticleReq: CollectArticleReq; //

const { status, data } = await apiInstance.collectArticle(
    collectArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **collectArticleReq** | **CollectArticleReq**|  | |


### Return type

**CollectArticleResp**

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

# **countCollectedArticles**
> CountCollectedArticlesResp countCollectedArticles(body)

统计当前账号收藏的文章数量。

### Example

```typescript
import {
    ArticleService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let body: object; //

const { status, data } = await apiInstance.countCollectedArticles(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**CountCollectedArticlesResp**

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

# **createDraftArticle**
> CreateDraftArticleResp createDraftArticle(createDraftArticleReq)

创建文章草稿

### Example

```typescript
import {
    ArticleService,
    Configuration,
    CreateDraftArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let createDraftArticleReq: CreateDraftArticleReq; //

const { status, data } = await apiInstance.createDraftArticle(
    createDraftArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createDraftArticleReq** | **CreateDraftArticleReq**|  | |


### Return type

**CreateDraftArticleResp**

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

# **discardDraftArticle**
> object discardDraftArticle(discardDraftArticleReq)

丢弃文章草稿

### Example

```typescript
import {
    ArticleService,
    Configuration,
    DiscardDraftArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let discardDraftArticleReq: DiscardDraftArticleReq; //

const { status, data } = await apiInstance.discardDraftArticle(
    discardDraftArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **discardDraftArticleReq** | **DiscardDraftArticleReq**|  | |


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

# **enterArticleLottery**
> object enterArticleLottery(enterArticleLotteryReq)

报名或撤销抽奖报名。

### Example

```typescript
import {
    ArticleService,
    Configuration,
    EnterArticleLotteryReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let enterArticleLotteryReq: EnterArticleLotteryReq; //

const { status, data } = await apiInstance.enterArticleLottery(
    enterArticleLotteryReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **enterArticleLotteryReq** | **EnterArticleLotteryReq**|  | |


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

# **getArticle**
> GetArticleResp getArticle(getArticleReq)

查询文章详情

### Example

```typescript
import {
    ArticleService,
    Configuration,
    GetArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let getArticleReq: GetArticleReq; //

const { status, data } = await apiInstance.getArticle(
    getArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **getArticleReq** | **GetArticleReq**|  | |


### Return type

**GetArticleResp**

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

# **likeArticle**
> LikeArticleResp likeArticle(likeArticleReq)

点赞文章

### Example

```typescript
import {
    ArticleService,
    Configuration,
    LikeArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let likeArticleReq: LikeArticleReq; //

const { status, data } = await apiInstance.likeArticle(
    likeArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **likeArticleReq** | **LikeArticleReq**|  | |


### Return type

**LikeArticleResp**

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

# **listArticles**
> ListArticlesResp listArticles(listArticlesReq)

查询文章列表

### Example

```typescript
import {
    ArticleService,
    Configuration,
    ListArticlesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let listArticlesReq: ListArticlesReq; //

const { status, data } = await apiInstance.listArticles(
    listArticlesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listArticlesReq** | **ListArticlesReq**|  | |


### Return type

**ListArticlesResp**

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

# **listFollowingArticles**
> ListFollowingArticlesResp listFollowingArticles(listFollowingArticlesReq)

查询当前账号关注作者发布的文章。

### Example

```typescript
import {
    ArticleService,
    Configuration,
    ListFollowingArticlesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let listFollowingArticlesReq: ListFollowingArticlesReq; //

const { status, data } = await apiInstance.listFollowingArticles(
    listFollowingArticlesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listFollowingArticlesReq** | **ListFollowingArticlesReq**|  | |


### Return type

**ListFollowingArticlesResp**

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

# **listNearbyArticles**
> ListNearbyArticlesResp listNearbyArticles(listNearbyArticlesReq)

查询当前账号所在城市发布的文章。

### Example

```typescript
import {
    ArticleService,
    Configuration,
    ListNearbyArticlesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let listNearbyArticlesReq: ListNearbyArticlesReq; //

const { status, data } = await apiInstance.listNearbyArticles(
    listNearbyArticlesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listNearbyArticlesReq** | **ListNearbyArticlesReq**|  | |


### Return type

**ListNearbyArticlesResp**

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

# **pageArticleViewHistory**
> PageArticleViewHistoryResp pageArticleViewHistory(pageArticleViewHistoryReq)

查询当前账号的浏览历史。

### Example

```typescript
import {
    ArticleService,
    Configuration,
    PageArticleViewHistoryReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let pageArticleViewHistoryReq: PageArticleViewHistoryReq; //

const { status, data } = await apiInstance.pageArticleViewHistory(
    pageArticleViewHistoryReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pageArticleViewHistoryReq** | **PageArticleViewHistoryReq**|  | |


### Return type

**PageArticleViewHistoryResp**

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

# **pageCollectedArticles**
> PageCollectedArticlesResp pageCollectedArticles(pageCollectedArticlesReq)

查询当前账号收藏的文章。

### Example

```typescript
import {
    ArticleService,
    Configuration,
    PageCollectedArticlesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let pageCollectedArticlesReq: PageCollectedArticlesReq; //

const { status, data } = await apiInstance.pageCollectedArticles(
    pageCollectedArticlesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pageCollectedArticlesReq** | **PageCollectedArticlesReq**|  | |


### Return type

**PageCollectedArticlesResp**

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

# **pageWatchedArticles**
> PageWatchedArticlesResp pageWatchedArticles(pageWatchedArticlesReq)

查询当前账号关注的文章。

### Example

```typescript
import {
    ArticleService,
    Configuration,
    PageWatchedArticlesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let pageWatchedArticlesReq: PageWatchedArticlesReq; //

const { status, data } = await apiInstance.pageWatchedArticles(
    pageWatchedArticlesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pageWatchedArticlesReq** | **PageWatchedArticlesReq**|  | |


### Return type

**PageWatchedArticlesResp**

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

# **publishArticle**
> object publishArticle(publishArticleReq)

发布文章

### Example

```typescript
import {
    ArticleService,
    Configuration,
    PublishArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let publishArticleReq: PublishArticleReq; //

const { status, data } = await apiInstance.publishArticle(
    publishArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **publishArticleReq** | **PublishArticleReq**|  | |


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

# **rewardArticle**
> RewardArticleResp rewardArticle(rewardArticleReq)

打赏文章

### Example

```typescript
import {
    ArticleService,
    Configuration,
    RewardArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let rewardArticleReq: RewardArticleReq; //

const { status, data } = await apiInstance.rewardArticle(
    rewardArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rewardArticleReq** | **RewardArticleReq**|  | |


### Return type

**RewardArticleResp**

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

# **setPinnedArticleComment**
> object setPinnedArticleComment(setPinnedArticleCommentReq)

设置或取消当前账号文章的置顶评论。

### Example

```typescript
import {
    ArticleService,
    Configuration,
    SetPinnedArticleCommentReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let setPinnedArticleCommentReq: SetPinnedArticleCommentReq; //

const { status, data } = await apiInstance.setPinnedArticleComment(
    setPinnedArticleCommentReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **setPinnedArticleCommentReq** | **SetPinnedArticleCommentReq**|  | |


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

# **thankArticle**
> ThankArticleResp thankArticle(thankArticleReq)

感谢文章，确认后不可撤销。

### Example

```typescript
import {
    ArticleService,
    Configuration,
    ThankArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let thankArticleReq: ThankArticleReq; //

const { status, data } = await apiInstance.thankArticle(
    thankArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **thankArticleReq** | **ThankArticleReq**|  | |


### Return type

**ThankArticleResp**

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

# **updateDraftArticle**
> UpdateDraftArticleResp updateDraftArticle(updateDraftArticleReq)

编辑文章草稿

### Example

```typescript
import {
    ArticleService,
    Configuration,
    UpdateDraftArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let updateDraftArticleReq: UpdateDraftArticleReq; //

const { status, data } = await apiInstance.updateDraftArticle(
    updateDraftArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateDraftArticleReq** | **UpdateDraftArticleReq**|  | |


### Return type

**UpdateDraftArticleResp**

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

# **voteArticlePoll**
> object voteArticlePoll(voteArticlePollReq)

提交一次不可修改的文章投票。

### Example

```typescript
import {
    ArticleService,
    Configuration,
    VoteArticlePollReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let voteArticlePollReq: VoteArticlePollReq; //

const { status, data } = await apiInstance.voteArticlePoll(
    voteArticlePollReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **voteArticlePollReq** | **VoteArticlePollReq**|  | |


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

# **watchArticle**
> WatchArticleResp watchArticle(watchArticleReq)

关注文章

### Example

```typescript
import {
    ArticleService,
    Configuration,
    WatchArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let watchArticleReq: WatchArticleReq; //

const { status, data } = await apiInstance.watchArticle(
    watchArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **watchArticleReq** | **WatchArticleReq**|  | |


### Return type

**WatchArticleResp**

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

