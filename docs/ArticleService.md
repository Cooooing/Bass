# ArticleService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**archive**](#archive) | **POST** /v1/content/article/archive | |
|[**cancelPublish**](#cancelpublish) | **POST** /v1/content/article/publish/cancel | |
|[**collect**](#collect) | **POST** /v1/content/article/collect | |
|[**createDraft**](#createdraft) | **POST** /v1/content/article/draft/create | |
|[**discardDraft**](#discarddraft) | **POST** /v1/content/article/draft/discard | |
|[**get**](#get) | **POST** /v1/content/article/get | |
|[**like**](#like) | **POST** /v1/content/article/like | |
|[**list**](#list) | **POST** /v1/content/article/list | |
|[**publish**](#publish) | **POST** /v1/content/article/publish | |
|[**reward**](#reward) | **POST** /v1/content/article/reward | |
|[**schedulePublish**](#schedulepublish) | **POST** /v1/content/article/publish/schedule | |
|[**thank**](#thank) | **POST** /v1/content/article/thank | |
|[**updateDraft**](#updatedraft) | **POST** /v1/content/article/draft/update | |

# **archive**
> object archive(archiveArticleReq)

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

const { status, data } = await apiInstance.archive(
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

# **cancelPublish**
> object cancelPublish(cancelPublishArticleReq)

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

const { status, data } = await apiInstance.cancelPublish(
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

# **collect**
> CollectArticleResp collect(collectArticleReq)

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

const { status, data } = await apiInstance.collect(
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

# **createDraft**
> CreateDraftArticleResp createDraft(createDraftArticleReq)

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

const { status, data } = await apiInstance.createDraft(
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

# **discardDraft**
> object discardDraft(discardDraftArticleReq)

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

const { status, data } = await apiInstance.discardDraft(
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

# **get**
> GetArticleResp get(getArticleReq)

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

const { status, data } = await apiInstance.get(
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

# **like**
> LikeArticleResp like(likeArticleReq)

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

const { status, data } = await apiInstance.like(
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

# **list**
> ListArticlesResp list(listArticlesReq)

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

const { status, data } = await apiInstance.list(
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

# **publish**
> object publish(publishArticleReq)

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

const { status, data } = await apiInstance.publish(
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

# **reward**
> object reward(rewardArticleReq)

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

const { status, data } = await apiInstance.reward(
    rewardArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rewardArticleReq** | **RewardArticleReq**|  | |


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

# **schedulePublish**
> object schedulePublish(schedulePublishArticleReq)

设置定时发布

### Example

```typescript
import {
    ArticleService,
    Configuration,
    SchedulePublishArticleReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let schedulePublishArticleReq: SchedulePublishArticleReq; //

const { status, data } = await apiInstance.schedulePublish(
    schedulePublishArticleReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **schedulePublishArticleReq** | **SchedulePublishArticleReq**|  | |


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

# **thank**
> ThankArticleResp thank(thankArticleReq)

感谢文章

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

const { status, data } = await apiInstance.thank(
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

# **updateDraft**
> UpdateDraftArticleResp updateDraft(updateDraftArticleReq)

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

const { status, data } = await apiInstance.updateDraft(
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

