# TagService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**list**](#list) | **POST** /v1/content/tag/list | |

# **list**
> ListTagsReply list(listTagsRequest)


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

