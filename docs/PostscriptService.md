# PostscriptService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**add**](#add) | **POST** /v1/content/postscript/add | |

# **add**
> AddPostscriptReply add(addPostscriptRequest)

添加文章附言

### Example

```typescript
import {
    PostscriptService,
    Configuration,
    AddPostscriptRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new PostscriptService(configuration);

let addPostscriptRequest: AddPostscriptRequest; //

const { status, data } = await apiInstance.add(
    addPostscriptRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **addPostscriptRequest** | **AddPostscriptRequest**|  | |


### Return type

**AddPostscriptReply**

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

