# CommentService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**create**](#create) | **POST** /v1/content/comment/create | |
|[**like**](#like) | **POST** /v1/content/comment/like | |
|[**list**](#list) | **POST** /v1/content/comment/list | |
|[**listReplies**](#listreplies) | **POST** /v1/content/comment/list-replies | |
|[**listThreads**](#listthreads) | **POST** /v1/content/comment/list-threads | |
|[**listTimeline**](#listtimeline) | **POST** /v1/content/comment/list-timeline | |
|[**thank**](#thank) | **POST** /v1/content/comment/thank | |

# **create**
> CreateCommentResp create(createCommentReq)

创建评论。

### Example

```typescript
import {
    CommentService,
    Configuration,
    CreateCommentReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let createCommentReq: CreateCommentReq; //

const { status, data } = await apiInstance.create(
    createCommentReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createCommentReq** | **CreateCommentReq**|  | |


### Return type

**CreateCommentResp**

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
> LikeCommentResp like(likeCommentReq)

点赞或取消点赞评论。

### Example

```typescript
import {
    CommentService,
    Configuration,
    LikeCommentReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let likeCommentReq: LikeCommentReq; //

const { status, data } = await apiInstance.like(
    likeCommentReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **likeCommentReq** | **LikeCommentReq**|  | |


### Return type

**LikeCommentResp**

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
> ListCommentsResp list(listCommentsReq)

分页查询评论列表。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ListCommentsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let listCommentsReq: ListCommentsReq; //

const { status, data } = await apiInstance.list(
    listCommentsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listCommentsReq** | **ListCommentsReq**|  | |


### Return type

**ListCommentsResp**

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

# **listReplies**
> ListCommentRepliesResp listReplies(listCommentRepliesReq)

分页查询评论回复。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ListCommentRepliesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let listCommentRepliesReq: ListCommentRepliesReq; //

const { status, data } = await apiInstance.listReplies(
    listCommentRepliesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listCommentRepliesReq** | **ListCommentRepliesReq**|  | |


### Return type

**ListCommentRepliesResp**

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

# **listThreads**
> ListCommentThreadsResp listThreads(listCommentThreadsReq)

分页查询评论楼层。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ListCommentThreadsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let listCommentThreadsReq: ListCommentThreadsReq; //

const { status, data } = await apiInstance.listThreads(
    listCommentThreadsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listCommentThreadsReq** | **ListCommentThreadsReq**|  | |


### Return type

**ListCommentThreadsResp**

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

# **listTimeline**
> ListCommentTimelineResp listTimeline(listCommentTimelineReq)

分页查询评论时间线。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ListCommentTimelineReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let listCommentTimelineReq: ListCommentTimelineReq; //

const { status, data } = await apiInstance.listTimeline(
    listCommentTimelineReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listCommentTimelineReq** | **ListCommentTimelineReq**|  | |


### Return type

**ListCommentTimelineResp**

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
> ThankCommentResp thank(thankCommentReq)

感谢或取消感谢评论。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ThankCommentReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let thankCommentReq: ThankCommentReq; //

const { status, data } = await apiInstance.thank(
    thankCommentReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **thankCommentReq** | **ThankCommentReq**|  | |


### Return type

**ThankCommentResp**

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

