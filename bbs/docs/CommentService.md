# CommentService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**create**](CommentService.md#create) | **POST** /v1/content/comment/create |  |
| [**like**](CommentService.md#like) | **POST** /v1/content/comment/like |  |
| [**list**](CommentService.md#list) | **POST** /v1/content/comment/list |  |
| [**listReplies**](CommentService.md#listreplies) | **POST** /v1/content/comment/list-replies |  |
| [**listThreads**](CommentService.md#listthreads) | **POST** /v1/content/comment/list-threads |  |
| [**listTimeline**](CommentService.md#listtimeline) | **POST** /v1/content/comment/list-timeline |  |
| [**thank**](CommentService.md#thank) | **POST** /v1/content/comment/thank |  |



## create

> CreateCommentReply create(createCommentRequest)



创建评论。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { CreateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // CreateCommentRequest
    createCommentRequest: ...,
  } satisfies CreateRequest;

  try {
    const data = await api.create(body);
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
| **createCommentRequest** | [CreateCommentRequest](CreateCommentRequest.md) |  | |

### Return type

[**CreateCommentReply**](CreateCommentReply.md)

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

> LikeCommentReply like(likeCommentRequest)



点赞或取消点赞评论。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { LikeRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // LikeCommentRequest
    likeCommentRequest: ...,
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
| **likeCommentRequest** | [LikeCommentRequest](LikeCommentRequest.md) |  | |

### Return type

[**LikeCommentReply**](LikeCommentReply.md)

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

> ListCommentsReply list(listCommentsRequest)



分页查询评论列表。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ListRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ListCommentsRequest
    listCommentsRequest: ...,
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
| **listCommentsRequest** | [ListCommentsRequest](ListCommentsRequest.md) |  | |

### Return type

[**ListCommentsReply**](ListCommentsReply.md)

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


## listReplies

> ListCommentRepliesReply listReplies(listCommentRepliesRequest)



分页查询评论回复。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ListRepliesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ListCommentRepliesRequest
    listCommentRepliesRequest: ...,
  } satisfies ListRepliesRequest;

  try {
    const data = await api.listReplies(body);
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
| **listCommentRepliesRequest** | [ListCommentRepliesRequest](ListCommentRepliesRequest.md) |  | |

### Return type

[**ListCommentRepliesReply**](ListCommentRepliesReply.md)

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


## listThreads

> ListCommentThreadsReply listThreads(listCommentThreadsRequest)



分页查询评论楼层。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ListThreadsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ListCommentThreadsRequest
    listCommentThreadsRequest: ...,
  } satisfies ListThreadsRequest;

  try {
    const data = await api.listThreads(body);
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
| **listCommentThreadsRequest** | [ListCommentThreadsRequest](ListCommentThreadsRequest.md) |  | |

### Return type

[**ListCommentThreadsReply**](ListCommentThreadsReply.md)

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


## listTimeline

> ListCommentTimelineReply listTimeline(listCommentTimelineRequest)



分页查询评论时间线。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ListTimelineRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ListCommentTimelineRequest
    listCommentTimelineRequest: ...,
  } satisfies ListTimelineRequest;

  try {
    const data = await api.listTimeline(body);
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
| **listCommentTimelineRequest** | [ListCommentTimelineRequest](ListCommentTimelineRequest.md) |  | |

### Return type

[**ListCommentTimelineReply**](ListCommentTimelineReply.md)

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

> ThankCommentReply thank(thankCommentRequest)



感谢或取消感谢评论。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ThankRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ThankCommentRequest
    thankCommentRequest: ...,
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
| **thankCommentRequest** | [ThankCommentRequest](ThankCommentRequest.md) |  | |

### Return type

[**ThankCommentReply**](ThankCommentReply.md)

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

