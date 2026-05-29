# ArticleService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**_delete**](#_delete) | **POST** /v1/content/article/delete | |
|[**acceptAnswer**](#acceptanswer) | **POST** /v1/content/article/accept-answer | |
|[**collect**](#collect) | **POST** /v1/content/article/collect | |
|[**create**](#create) | **POST** /v1/content/article/create | |
|[**get**](#get) | **POST** /v1/content/article/get | |
|[**like**](#like) | **POST** /v1/content/article/like | |
|[**list**](#list) | **POST** /v1/content/article/list | |
|[**publish**](#publish) | **POST** /v1/content/article/publish | |
|[**reward**](#reward) | **POST** /v1/content/article/reward | |
|[**thank**](#thank) | **POST** /v1/content/article/thank | |
|[**updateDraft**](#updatedraft) | **POST** /v1/content/article/update-draft | |
|[**watch**](#watch) | **POST** /v1/content/article/watch | |

# **_delete**
> object _delete(deleteArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    DeleteArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let deleteArticleRequest: DeleteArticleRequest; //

const { status, data } = await apiInstance._delete(
    deleteArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **deleteArticleRequest** | **DeleteArticleRequest**|  | |


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

# **acceptAnswer**
> object acceptAnswer(acceptAnswerArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    AcceptAnswerArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let acceptAnswerArticleRequest: AcceptAnswerArticleRequest; //

const { status, data } = await apiInstance.acceptAnswer(
    acceptAnswerArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **acceptAnswerArticleRequest** | **AcceptAnswerArticleRequest**|  | |


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
> object collect(collectArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    CollectArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let collectArticleRequest: CollectArticleRequest; //

const { status, data } = await apiInstance.collect(
    collectArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **collectArticleRequest** | **CollectArticleRequest**|  | |


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

# **create**
> CreateArticleReply create(createArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    CreateArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let createArticleRequest: CreateArticleRequest; //

const { status, data } = await apiInstance.create(
    createArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createArticleRequest** | **CreateArticleRequest**|  | |


### Return type

**CreateArticleReply**

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
> GetArticleReply get(getArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    GetArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let getArticleRequest: GetArticleRequest; //

const { status, data } = await apiInstance.get(
    getArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **getArticleRequest** | **GetArticleRequest**|  | |


### Return type

**GetArticleReply**

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
> object like(likeArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    LikeArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let likeArticleRequest: LikeArticleRequest; //

const { status, data } = await apiInstance.like(
    likeArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **likeArticleRequest** | **LikeArticleRequest**|  | |


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

# **list**
> ListArticlesReply list(listArticlesRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    ListArticlesRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let listArticlesRequest: ListArticlesRequest; //

const { status, data } = await apiInstance.list(
    listArticlesRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listArticlesRequest** | **ListArticlesRequest**|  | |


### Return type

**ListArticlesReply**

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
> object publish(publishArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    PublishArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let publishArticleRequest: PublishArticleRequest; //

const { status, data } = await apiInstance.publish(
    publishArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **publishArticleRequest** | **PublishArticleRequest**|  | |


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
> object reward(rewardArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    RewardArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let rewardArticleRequest: RewardArticleRequest; //

const { status, data } = await apiInstance.reward(
    rewardArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **rewardArticleRequest** | **RewardArticleRequest**|  | |


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
> object thank(thankArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    ThankArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let thankArticleRequest: ThankArticleRequest; //

const { status, data } = await apiInstance.thank(
    thankArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **thankArticleRequest** | **ThankArticleRequest**|  | |


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

# **updateDraft**
> UpdateDraftArticleReply updateDraft(updateDraftArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    UpdateDraftArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let updateDraftArticleRequest: UpdateDraftArticleRequest; //

const { status, data } = await apiInstance.updateDraft(
    updateDraftArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateDraftArticleRequest** | **UpdateDraftArticleRequest**|  | |


### Return type

**UpdateDraftArticleReply**

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

# **watch**
> object watch(watchArticleRequest)


### Example

```typescript
import {
    ArticleService,
    Configuration,
    WatchArticleRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new ArticleService(configuration);

let watchArticleRequest: WatchArticleRequest; //

const { status, data } = await apiInstance.watch(
    watchArticleRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **watchArticleRequest** | **WatchArticleRequest**|  | |


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

