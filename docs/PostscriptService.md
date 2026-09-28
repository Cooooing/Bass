# PostscriptService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**add**](#add) | **POST** /v1/content/postscript/add | |
|[**list**](#list) | **POST** /v1/content/postscript/list | |

# **add**
> AddPostscriptResp add(addPostscriptReq)

添加文章附言。

### Example

```typescript
import {
    PostscriptService,
    Configuration,
    AddPostscriptReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new PostscriptService(configuration);

let addPostscriptReq: AddPostscriptReq; //

const { status, data } = await apiInstance.add(
    addPostscriptReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **addPostscriptReq** | **AddPostscriptReq**|  | |


### Return type

**AddPostscriptResp**

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
> ListPostscriptsResp list(listPostscriptsReq)

查询文章附言列表。

### Example

```typescript
import {
    PostscriptService,
    Configuration,
    ListPostscriptsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new PostscriptService(configuration);

let listPostscriptsReq: ListPostscriptsReq; //

const { status, data } = await apiInstance.list(
    listPostscriptsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listPostscriptsReq** | **ListPostscriptsReq**|  | |


### Return type

**ListPostscriptsResp**

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

