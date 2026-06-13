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
> CreateCommentReply create(createCommentRequest)

创建评论。

### Example

```typescript
import {
    CommentService,
    Configuration,
    CreateCommentRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let createCommentRequest: CreateCommentRequest; //

const { status, data } = await apiInstance.create(
    createCommentRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createCommentRequest** | **CreateCommentRequest**|  | |


### Return type

**CreateCommentReply**

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
> LikeCommentReply like(likeCommentRequest)

点赞或取消点赞评论。

### Example

```typescript
import {
    CommentService,
    Configuration,
    LikeCommentRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let likeCommentRequest: LikeCommentRequest; //

const { status, data } = await apiInstance.like(
    likeCommentRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **likeCommentRequest** | **LikeCommentRequest**|  | |


### Return type

**LikeCommentReply**

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
> ListCommentsReply list(listCommentsRequest)

分页查询评论列表。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ListCommentsRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let listCommentsRequest: ListCommentsRequest; //

const { status, data } = await apiInstance.list(
    listCommentsRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listCommentsRequest** | **ListCommentsRequest**|  | |


### Return type

**ListCommentsReply**

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
> ListCommentRepliesReply listReplies(listCommentRepliesRequest)

分页查询评论回复。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ListCommentRepliesRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let listCommentRepliesRequest: ListCommentRepliesRequest; //

const { status, data } = await apiInstance.listReplies(
    listCommentRepliesRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listCommentRepliesRequest** | **ListCommentRepliesRequest**|  | |


### Return type

**ListCommentRepliesReply**

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
> ListCommentThreadsReply listThreads(listCommentThreadsRequest)

分页查询评论楼层。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ListCommentThreadsRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let listCommentThreadsRequest: ListCommentThreadsRequest; //

const { status, data } = await apiInstance.listThreads(
    listCommentThreadsRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listCommentThreadsRequest** | **ListCommentThreadsRequest**|  | |


### Return type

**ListCommentThreadsReply**

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
> ListCommentTimelineReply listTimeline(listCommentTimelineRequest)

分页查询评论时间线。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ListCommentTimelineRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let listCommentTimelineRequest: ListCommentTimelineRequest; //

const { status, data } = await apiInstance.listTimeline(
    listCommentTimelineRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listCommentTimelineRequest** | **ListCommentTimelineRequest**|  | |


### Return type

**ListCommentTimelineReply**

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
> ThankCommentReply thank(thankCommentRequest)

感谢或取消感谢评论。

### Example

```typescript
import {
    CommentService,
    Configuration,
    ThankCommentRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CommentService(configuration);

let thankCommentRequest: ThankCommentRequest; //

const { status, data } = await apiInstance.thank(
    thankCommentRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **thankCommentRequest** | **ThankCommentRequest**|  | |


### Return type

**ThankCommentReply**

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

