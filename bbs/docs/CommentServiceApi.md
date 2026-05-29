# CommentServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**commentServiceCreate**](CommentServiceApi.md#commentservicecreate) | **POST** /v1/content/comment/create |  |
| [**commentServiceLike**](CommentServiceApi.md#commentservicelike) | **POST** /v1/content/comment/like |  |
| [**commentServiceList**](CommentServiceApi.md#commentservicelist) | **POST** /v1/content/comment/list |  |
| [**commentServiceThank**](CommentServiceApi.md#commentservicethank) | **POST** /v1/content/comment/thank |  |



## commentServiceCreate

> CreateCommentReply commentServiceCreate(createCommentRequest)



### Example

```ts
import {
  Configuration,
  CommentServiceApi,
} from '@bass/bbs-sdk';
import type { CommentServiceCreateRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new CommentServiceApi();

  const body = {
    // CreateCommentRequest
    createCommentRequest: ...,
  } satisfies CommentServiceCreateRequest;

  try {
    const data = await api.commentServiceCreate(body);
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


## commentServiceLike

> object commentServiceLike(likeCommentRequest)



### Example

```ts
import {
  Configuration,
  CommentServiceApi,
} from '@bass/bbs-sdk';
import type { CommentServiceLikeRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new CommentServiceApi();

  const body = {
    // LikeCommentRequest
    likeCommentRequest: ...,
  } satisfies CommentServiceLikeRequest;

  try {
    const data = await api.commentServiceLike(body);
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


## commentServiceList

> ListCommentsReply commentServiceList(listCommentsRequest)



### Example

```ts
import {
  Configuration,
  CommentServiceApi,
} from '@bass/bbs-sdk';
import type { CommentServiceListRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new CommentServiceApi();

  const body = {
    // ListCommentsRequest
    listCommentsRequest: ...,
  } satisfies CommentServiceListRequest;

  try {
    const data = await api.commentServiceList(body);
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


## commentServiceThank

> object commentServiceThank(thankCommentRequest)



### Example

```ts
import {
  Configuration,
  CommentServiceApi,
} from '@bass/bbs-sdk';
import type { CommentServiceThankRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new CommentServiceApi();

  const body = {
    // ThankCommentRequest
    thankCommentRequest: ...,
  } satisfies CommentServiceThankRequest;

  try {
    const data = await api.commentServiceThank(body);
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

