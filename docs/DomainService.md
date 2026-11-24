# DomainService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createDomain**](DomainService.md#createdomain) | **POST** /v1/content/domain/create |  |
| [**listDomains**](DomainService.md#listdomains) | **POST** /v1/content/domain/list |  |
| [**updateDomain**](DomainService.md#updatedomain) | **POST** /v1/content/domain/update |  |



## createDomain

> CreateDomainResp createDomain(createDomainReq)



创建领域。

### Example

```ts
import {
  Configuration,
  DomainService,
} from '@bass/bbs-sdk-fetch';
import type { CreateDomainRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new DomainService();

  const body = {
    // CreateDomainReq
    createDomainReq: ...,
  } satisfies CreateDomainRequest;

  try {
    const data = await api.createDomain(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **createDomainReq** | [CreateDomainReq](CreateDomainReq.md) |  | |

### Return type

[**CreateDomainResp**](CreateDomainResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listDomains

> ListDomainsResp listDomains(listDomainsReq)



查询领域列表。

### Example

```ts
import {
  Configuration,
  DomainService,
} from '@bass/bbs-sdk-fetch';
import type { ListDomainsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new DomainService();

  const body = {
    // ListDomainsReq
    listDomainsReq: ...,
  } satisfies ListDomainsRequest;

  try {
    const data = await api.listDomains(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **listDomainsReq** | [ListDomainsReq](ListDomainsReq.md) |  | |

### Return type

[**ListDomainsResp**](ListDomainsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## updateDomain

> UpdateDomainResp updateDomain(updateDomainReq)



更新领域。

### Example

```ts
import {
  Configuration,
  DomainService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateDomainRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new DomainService();

  const body = {
    // UpdateDomainReq
    updateDomainReq: ...,
  } satisfies UpdateDomainRequest;

  try {
    const data = await api.updateDomain(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **updateDomainReq** | [UpdateDomainReq](UpdateDomainReq.md) |  | |

### Return type

[**UpdateDomainResp**](UpdateDomainResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

