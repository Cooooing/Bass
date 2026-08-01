# DomainService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**create**](DomainService.md#create) | **POST** /v1/content/domain/create |  |
| [**list**](DomainService.md#list) | **POST** /v1/content/domain/list |  |
| [**update**](DomainService.md#update) | **POST** /v1/content/domain/update |  |



## create

> CreateDomainResp create(createDomainReq)



创建领域。

### Example

```ts
import {
  Configuration,
  DomainService,
} from '@bass/bbs-sdk-fetch';
import type { CreateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new DomainService();

  const body = {
    // CreateDomainReq
    createDomainReq: ...,
  } satisfies CreateRequest;

  try {
    const data = await api.create(body);
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


## list

> ListDomainsResp list(listDomainsReq)



查询领域列表。

### Example

```ts
import {
  Configuration,
  DomainService,
} from '@bass/bbs-sdk-fetch';
import type { ListRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new DomainService();

  const body = {
    // ListDomainsReq
    listDomainsReq: ...,
  } satisfies ListRequest;

  try {
    const data = await api.list(body);
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


## update

> UpdateDomainResp update(updateDomainReq)



更新领域。

### Example

```ts
import {
  Configuration,
  DomainService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new DomainService();

  const body = {
    // UpdateDomainReq
    updateDomainReq: ...,
  } satisfies UpdateRequest;

  try {
    const data = await api.update(body);
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

