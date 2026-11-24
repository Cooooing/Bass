# PostscriptService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**addPostscript**](#addpostscript) | **POST** /v1/content/postscript/add | |
|[**listPostscripts**](#listpostscripts) | **POST** /v1/content/postscript/list | |

# **addPostscript**
> AddPostscriptResp addPostscript(addPostscriptReq)

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

const { status, data } = await apiInstance.addPostscript(
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

# **listPostscripts**
> ListPostscriptsResp listPostscripts(listPostscriptsReq)

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

const { status, data } = await apiInstance.listPostscripts(
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

