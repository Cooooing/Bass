# CommentService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createComment**](CommentService.md#createcomment) | **POST** /v1/content/comment/create |  |
| [**likeComment**](CommentService.md#likecomment) | **POST** /v1/content/comment/like |  |
| [**listCommentReplies**](CommentService.md#listcommentreplies) | **POST** /v1/content/comment/list-replies |  |
| [**listCommentThreads**](CommentService.md#listcommentthreads) | **POST** /v1/content/comment/list-threads |  |
| [**listCommentTimeline**](CommentService.md#listcommenttimeline) | **POST** /v1/content/comment/list-timeline |  |
| [**listComments**](CommentService.md#listcomments) | **POST** /v1/content/comment/list |  |
| [**thankComment**](CommentService.md#thankcomment) | **POST** /v1/content/comment/thank |  |



## createComment

> CreateCommentResp createComment(createCommentReq)



创建评论。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { CreateCommentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // CreateCommentReq
    createCommentReq: ...,
  } satisfies CreateCommentRequest;

  try {
    const data = await api.createComment(body);
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


## likeComment

> LikeCommentResp likeComment(likeCommentReq)



点赞或取消点赞评论。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { LikeCommentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // LikeCommentReq
    likeCommentReq: ...,
  } satisfies LikeCommentRequest;

  try {
    const data = await api.likeComment(body);
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


## listCommentReplies

> ListCommentRepliesResp listCommentReplies(listCommentRepliesReq)



分页查询评论回复。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ListCommentRepliesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ListCommentRepliesReq
    listCommentRepliesReq: ...,
  } satisfies ListCommentRepliesRequest;

  try {
    const data = await api.listCommentReplies(body);
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


## listCommentThreads

> ListCommentThreadsResp listCommentThreads(listCommentThreadsReq)



分页查询评论楼层。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ListCommentThreadsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ListCommentThreadsReq
    listCommentThreadsReq: ...,
  } satisfies ListCommentThreadsRequest;

  try {
    const data = await api.listCommentThreads(body);
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


## listCommentTimeline

> ListCommentTimelineResp listCommentTimeline(listCommentTimelineReq)



分页查询评论时间线。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ListCommentTimelineRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ListCommentTimelineReq
    listCommentTimelineReq: ...,
  } satisfies ListCommentTimelineRequest;

  try {
    const data = await api.listCommentTimeline(body);
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


## listComments

> ListCommentsResp listComments(listCommentsReq)



分页查询评论列表。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ListCommentsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ListCommentsReq
    listCommentsReq: ...,
  } satisfies ListCommentsRequest;

  try {
    const data = await api.listComments(body);
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


## thankComment

> ThankCommentResp thankComment(thankCommentReq)



感谢或取消感谢评论。

### Example

```ts
import {
  Configuration,
  CommentService,
} from '@bass/bbs-sdk-fetch';
import type { ThankCommentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CommentService();

  const body = {
    // ThankCommentReq
    thankCommentReq: ...,
  } satisfies ThankCommentRequest;

  try {
    const data = await api.thankComment(body);
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

