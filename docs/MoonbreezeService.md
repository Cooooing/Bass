# MoonbreezeService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**createMoonbreeze**](#createmoonbreeze) | **POST** /v1/content/moonbreeze/create | |
|[**pageMemberMoonbreezes**](#pagemembermoonbreezes) | **POST** /v1/content/moonbreeze/page-member | |
|[**pagePublicMoonbreezes**](#pagepublicmoonbreezes) | **POST** /v1/content/moonbreeze/page-public | |
|[**pageWatchingMoonbreezes**](#pagewatchingmoonbreezes) | **POST** /v1/content/moonbreeze/page-watching | |

# **createMoonbreeze**
> CreateMoonbreezeResp createMoonbreeze(createMoonbreezeReq)

服务方法。

### Example

```typescript
import {
    MoonbreezeService,
    Configuration,
    CreateMoonbreezeReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new MoonbreezeService(configuration);

let createMoonbreezeReq: CreateMoonbreezeReq; //

const { status, data } = await apiInstance.createMoonbreeze(
    createMoonbreezeReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createMoonbreezeReq** | **CreateMoonbreezeReq**|  | |


### Return type

**CreateMoonbreezeResp**

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

# **pageMemberMoonbreezes**
> PageMemberMoonbreezesResp pageMemberMoonbreezes(pageMemberMoonbreezesReq)

服务方法。

### Example

```typescript
import {
    MoonbreezeService,
    Configuration,
    PageMemberMoonbreezesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new MoonbreezeService(configuration);

let pageMemberMoonbreezesReq: PageMemberMoonbreezesReq; //

const { status, data } = await apiInstance.pageMemberMoonbreezes(
    pageMemberMoonbreezesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pageMemberMoonbreezesReq** | **PageMemberMoonbreezesReq**|  | |


### Return type

**PageMemberMoonbreezesResp**

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

# **pagePublicMoonbreezes**
> PagePublicMoonbreezesResp pagePublicMoonbreezes(pagePublicMoonbreezesReq)

服务方法。

### Example

```typescript
import {
    MoonbreezeService,
    Configuration,
    PagePublicMoonbreezesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new MoonbreezeService(configuration);

let pagePublicMoonbreezesReq: PagePublicMoonbreezesReq; //

const { status, data } = await apiInstance.pagePublicMoonbreezes(
    pagePublicMoonbreezesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pagePublicMoonbreezesReq** | **PagePublicMoonbreezesReq**|  | |


### Return type

**PagePublicMoonbreezesResp**

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

# **pageWatchingMoonbreezes**
> PageWatchingMoonbreezesResp pageWatchingMoonbreezes(pageWatchingMoonbreezesReq)

服务方法。

### Example

```typescript
import {
    MoonbreezeService,
    Configuration,
    PageWatchingMoonbreezesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new MoonbreezeService(configuration);

let pageWatchingMoonbreezesReq: PageWatchingMoonbreezesReq; //

const { status, data } = await apiInstance.pageWatchingMoonbreezes(
    pageWatchingMoonbreezesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pageWatchingMoonbreezesReq** | **PageWatchingMoonbreezesReq**|  | |


### Return type

**PageWatchingMoonbreezesResp**

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

