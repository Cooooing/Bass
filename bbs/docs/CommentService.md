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

> CreateCommentResp create(createCommentReq)



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
    // CreateCommentReq
    createCommentReq: ...,
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
| **createCommentReq** | [CreateCommentReq](CreateCommentReq.md) |  | |

### Return type

[**CreateCommentResp**](CreateCommentResp.md)

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

> LikeCommentResp like(likeCommentReq)



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
    // LikeCommentReq
    likeCommentReq: ...,
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
| **likeCommentReq** | [LikeCommentReq](LikeCommentReq.md) |  | |

### Return type

[**LikeCommentResp**](LikeCommentResp.md)

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

> ListCommentsResp list(listCommentsReq)



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
    // ListCommentsReq
    listCommentsReq: ...,
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
| **listCommentsReq** | [ListCommentsReq](ListCommentsReq.md) |  | |

### Return type

[**ListCommentsResp**](ListCommentsResp.md)

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

> ListCommentRepliesResp listReplies(listCommentRepliesReq)



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
    // ListCommentRepliesReq
    listCommentRepliesReq: ...,
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
| **listCommentRepliesReq** | [ListCommentRepliesReq](ListCommentRepliesReq.md) |  | |

### Return type

[**ListCommentRepliesResp**](ListCommentRepliesResp.md)

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

> ListCommentThreadsResp listThreads(listCommentThreadsReq)



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
    // ListCommentThreadsReq
    listCommentThreadsReq: ...,
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
| **listCommentThreadsReq** | [ListCommentThreadsReq](ListCommentThreadsReq.md) |  | |

### Return type

[**ListCommentThreadsResp**](ListCommentThreadsResp.md)

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

> ListCommentTimelineResp listTimeline(listCommentTimelineReq)



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
    // ListCommentTimelineReq
    listCommentTimelineReq: ...,
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
| **listCommentTimelineReq** | [ListCommentTimelineReq](ListCommentTimelineReq.md) |  | |

### Return type

[**ListCommentTimelineResp**](ListCommentTimelineResp.md)

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

> ThankCommentResp thank(thankCommentReq)



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
    // ThankCommentReq
    thankCommentReq: ...,
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
| **thankCommentReq** | [ThankCommentReq](ThankCommentReq.md) |  | |

### Return type

[**ThankCommentResp**](ThankCommentResp.md)

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

