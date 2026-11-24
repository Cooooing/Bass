# TagService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**bindArticleTags**](#bindarticletags) | **POST** /v1/content/tag/bind-article | |
|[**createTag**](#createtag) | **POST** /v1/content/tag/create | |
|[**listArticleTags**](#listarticletags) | **POST** /v1/content/tag/list-article-tags | |
|[**listTags**](#listtags) | **POST** /v1/content/tag/list | |
|[**unbindArticleTags**](#unbindarticletags) | **POST** /v1/content/tag/unbind-article | |
|[**updateTag**](#updatetag) | **POST** /v1/content/tag/update | |

# **bindArticleTags**
> object bindArticleTags(bindArticleTagsReq)

绑定文章标签。

### Example

```typescript
import {
    TagService,
    Configuration,
    BindArticleTagsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TagService(configuration);

let bindArticleTagsReq: BindArticleTagsReq; //

const { status, data } = await apiInstance.bindArticleTags(
    bindArticleTagsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **bindArticleTagsReq** | **BindArticleTagsReq**|  | |


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

# **createTag**
> CreateTagResp createTag(createTagReq)

创建标签。

### Example

```typescript
import {
    TagService,
    Configuration,
    CreateTagReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TagService(configuration);

let createTagReq: CreateTagReq; //

const { status, data } = await apiInstance.createTag(
    createTagReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createTagReq** | **CreateTagReq**|  | |


### Return type

**CreateTagResp**

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

# **listArticleTags**
> ListArticleTagsResp listArticleTags(listArticleTagsReq)

查询文章标签列表。

### Example

```typescript
import {
    TagService,
    Configuration,
    ListArticleTagsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TagService(configuration);

let listArticleTagsReq: ListArticleTagsReq; //

const { status, data } = await apiInstance.listArticleTags(
    listArticleTagsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listArticleTagsReq** | **ListArticleTagsReq**|  | |


### Return type

**ListArticleTagsResp**

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

# **listTags**
> ListTagsResp listTags(listTagsReq)

查询标签列表。

### Example

```typescript
import {
    TagService,
    Configuration,
    ListTagsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TagService(configuration);

let listTagsReq: ListTagsReq; //

const { status, data } = await apiInstance.listTags(
    listTagsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listTagsReq** | **ListTagsReq**|  | |


### Return type

**ListTagsResp**

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

# **unbindArticleTags**
> object unbindArticleTags(unbindArticleTagsReq)

解绑文章标签。

### Example

```typescript
import {
    TagService,
    Configuration,
    UnbindArticleTagsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TagService(configuration);

let unbindArticleTagsReq: UnbindArticleTagsReq; //

const { status, data } = await apiInstance.unbindArticleTags(
    unbindArticleTagsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **unbindArticleTagsReq** | **UnbindArticleTagsReq**|  | |


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

# **updateTag**
> UpdateTagResp updateTag(updateTagReq)

更新标签。

### Example

```typescript
import {
    TagService,
    Configuration,
    UpdateTagReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TagService(configuration);

let updateTagReq: UpdateTagReq; //

const { status, data } = await apiInstance.updateTag(
    updateTagReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateTagReq** | **UpdateTagReq**|  | |


### Return type

**UpdateTagResp**

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

