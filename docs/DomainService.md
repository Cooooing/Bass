# DomainService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**create**](#create) | **POST** /v1/content/domain/create | |
|[**list**](#list) | **POST** /v1/content/domain/list | |
|[**update**](#update) | **POST** /v1/content/domain/update | |

# **create**
> CreateDomainResp create(createDomainReq)

创建领域。

### Example

```typescript
import {
    DomainService,
    Configuration,
    CreateDomainReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new DomainService(configuration);

let createDomainReq: CreateDomainReq; //

const { status, data } = await apiInstance.create(
    createDomainReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createDomainReq** | **CreateDomainReq**|  | |


### Return type

**CreateDomainResp**

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
> ListDomainsResp list(listDomainsReq)

查询领域列表。

### Example

```typescript
import {
    DomainService,
    Configuration,
    ListDomainsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new DomainService(configuration);

let listDomainsReq: ListDomainsReq; //

const { status, data } = await apiInstance.list(
    listDomainsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listDomainsReq** | **ListDomainsReq**|  | |


### Return type

**ListDomainsResp**

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
> UpdateDomainResp update(updateDomainReq)

更新领域。

### Example

```typescript
import {
    DomainService,
    Configuration,
    UpdateDomainReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new DomainService(configuration);

let updateDomainReq: UpdateDomainReq; //

const { status, data } = await apiInstance.update(
    updateDomainReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateDomainReq** | **UpdateDomainReq**|  | |


### Return type

**UpdateDomainResp**

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

