# MoonbreezeService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**create**](#create) | **POST** /v1/content/moonbreeze/create | |
|[**pageMember**](#pagemember) | **POST** /v1/content/moonbreeze/page-member | |
|[**pagePublic**](#pagepublic) | **POST** /v1/content/moonbreeze/page-public | |
|[**pageWatching**](#pagewatching) | **POST** /v1/content/moonbreeze/page-watching | |

# **create**
> CreateMoonbreezeResp create(createMoonbreezeReq)


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

const { status, data } = await apiInstance.create(
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

# **pageMember**
> PageMoonbreezesResp pageMember(pageMemberMoonbreezesReq)


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

const { status, data } = await apiInstance.pageMember(
    pageMemberMoonbreezesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pageMemberMoonbreezesReq** | **PageMemberMoonbreezesReq**|  | |


### Return type

**PageMoonbreezesResp**

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

# **pagePublic**
> PageMoonbreezesResp pagePublic(pagePublicMoonbreezesReq)


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

const { status, data } = await apiInstance.pagePublic(
    pagePublicMoonbreezesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pagePublicMoonbreezesReq** | **PagePublicMoonbreezesReq**|  | |


### Return type

**PageMoonbreezesResp**

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

# **pageWatching**
> PageMoonbreezesResp pageWatching(pageWatchingMoonbreezesReq)


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

const { status, data } = await apiInstance.pageWatching(
    pageWatchingMoonbreezesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pageWatchingMoonbreezesReq** | **PageWatchingMoonbreezesReq**|  | |


### Return type

**PageMoonbreezesResp**

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

