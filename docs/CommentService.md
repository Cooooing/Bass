# CommentService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**create**](#create) | **POST** /v1/content/comment/create | |
|[**like**](#like) | **POST** /v1/content/comment/like | |
|[**list**](#list) | **POST** /v1/content/comment/list | |
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
> object like(likeCommentRequest)

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

# **thank**
> object thank(thankCommentRequest)

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

