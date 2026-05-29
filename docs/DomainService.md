# DomainService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**list**](#list) | **POST** /v1/content/domain/list | |

# **list**
> ListDomainsReply list(listDomainsRequest)


### Example

```typescript
import {
    DomainService,
    Configuration,
    ListDomainsRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new DomainService(configuration);

let listDomainsRequest: ListDomainsRequest; //

const { status, data } = await apiInstance.list(
    listDomainsRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listDomainsRequest** | **ListDomainsRequest**|  | |


### Return type

**ListDomainsReply**

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

