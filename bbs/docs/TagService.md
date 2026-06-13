# TagService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**create**](#create) | **POST** /v1/content/tag/create | |
|[**list**](#list) | **POST** /v1/content/tag/list | |
|[**update**](#update) | **POST** /v1/content/tag/update | |

# **create**
> CreateTagReply create(createTagRequest)

创建标签。

### Example

```typescript
import {
    TagService,
    Configuration,
    CreateTagRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TagService(configuration);

let createTagRequest: CreateTagRequest; //

const { status, data } = await apiInstance.create(
    createTagRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createTagRequest** | **CreateTagRequest**|  | |


### Return type

**CreateTagReply**

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
> ListTagsReply list(listTagsRequest)

分页查询标签列表。

### Example

```typescript
import {
    TagService,
    Configuration,
    ListTagsRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TagService(configuration);

let listTagsRequest: ListTagsRequest; //

const { status, data } = await apiInstance.list(
    listTagsRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listTagsRequest** | **ListTagsRequest**|  | |


### Return type

**ListTagsReply**

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

# **update**
> UpdateTagReply update(updateTagRequest)

更新标签。

### Example

```typescript
import {
    TagService,
    Configuration,
    UpdateTagRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TagService(configuration);

let updateTagRequest: UpdateTagRequest; //

const { status, data } = await apiInstance.update(
    updateTagRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateTagRequest** | **UpdateTagRequest**|  | |


### Return type

**UpdateTagReply**

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

