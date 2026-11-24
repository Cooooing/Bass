# CommentService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createComment**](#createcomment) | **POST** /v1/content/comment/create | |
|[**likeComment**](#likecomment) | **POST** /v1/content/comment/like | |
|[**listCommentReplies**](#listcommentreplies) | **POST** /v1/content/comment/list-replies | |
|[**listCommentThreads**](#listcommentthreads) | **POST** /v1/content/comment/list-threads | |
|[**listCommentTimeline**](#listcommenttimeline) | **POST** /v1/content/comment/list-timeline | |
|[**listComments**](#listcomments) | **POST** /v1/content/comment/list | |
|[**thankComment**](#thankcomment) | **POST** /v1/content/comment/thank | |

# **createComment**
> CreateCommentResp createComment(createCommentReq)

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

const { status, data } = await apiInstance.createComment(
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

# **likeComment**
> LikeCommentResp likeComment(likeCommentReq)

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

const { status, data } = await apiInstance.likeComment(
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

# **listCommentReplies**
> ListCommentRepliesResp listCommentReplies(listCommentRepliesReq)

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

const { status, data } = await apiInstance.listCommentReplies(
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

# **listCommentThreads**
> ListCommentThreadsResp listCommentThreads(listCommentThreadsReq)

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

const { status, data } = await apiInstance.listCommentThreads(
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

# **listCommentTimeline**
> ListCommentTimelineResp listCommentTimeline(listCommentTimelineReq)

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

const { status, data } = await apiInstance.listCommentTimeline(
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

# **listComments**
> ListCommentsResp listComments(listCommentsReq)

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

const { status, data } = await apiInstance.listComments(
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

# **thankComment**
> ThankCommentResp thankComment(thankCommentReq)

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

const { status, data } = await apiInstance.thankComment(
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

