# MoonbreezeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createMoonbreeze**](MoonbreezeService.md#createmoonbreeze) | **POST** /v1/content/moonbreeze/create |  |
| [**pageMemberMoonbreezes**](MoonbreezeService.md#pagemembermoonbreezes) | **POST** /v1/content/moonbreeze/page-member |  |
| [**pagePublicMoonbreezes**](MoonbreezeService.md#pagepublicmoonbreezes) | **POST** /v1/content/moonbreeze/page-public |  |
| [**pageWatchingMoonbreezes**](MoonbreezeService.md#pagewatchingmoonbreezes) | **POST** /v1/content/moonbreeze/page-watching |  |



## createMoonbreeze

> CreateMoonbreezeResp createMoonbreeze(createMoonbreezeReq)



服务方法。

### Example

```ts
import {
  Configuration,
  MoonbreezeService,
} from '@bass/bbs-sdk-fetch';
import type { CreateMoonbreezeRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new MoonbreezeService();

  const body = {
    // CreateMoonbreezeReq
    createMoonbreezeReq: ...,
  } satisfies CreateMoonbreezeRequest;

  try {
    const data = await api.createMoonbreeze(body);
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
| **createMoonbreezeReq** | [CreateMoonbreezeReq](CreateMoonbreezeReq.md) |  | |

### Return type

[**CreateMoonbreezeResp**](CreateMoonbreezeResp.md)

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


## pageMemberMoonbreezes

> PageMemberMoonbreezesResp pageMemberMoonbreezes(pageMemberMoonbreezesReq)



服务方法。

### Example

```ts
import {
  Configuration,
  MoonbreezeService,
} from '@bass/bbs-sdk-fetch';
import type { PageMemberMoonbreezesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new MoonbreezeService();

  const body = {
    // PageMemberMoonbreezesReq
    pageMemberMoonbreezesReq: ...,
  } satisfies PageMemberMoonbreezesRequest;

  try {
    const data = await api.pageMemberMoonbreezes(body);
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
| **pageMemberMoonbreezesReq** | [PageMemberMoonbreezesReq](PageMemberMoonbreezesReq.md) |  | |

### Return type

[**PageMemberMoonbreezesResp**](PageMemberMoonbreezesResp.md)

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


## pagePublicMoonbreezes

> PagePublicMoonbreezesResp pagePublicMoonbreezes(pagePublicMoonbreezesReq)



服务方法。

### Example

```ts
import {
  Configuration,
  MoonbreezeService,
} from '@bass/bbs-sdk-fetch';
import type { PagePublicMoonbreezesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new MoonbreezeService();

  const body = {
    // PagePublicMoonbreezesReq
    pagePublicMoonbreezesReq: ...,
  } satisfies PagePublicMoonbreezesRequest;

  try {
    const data = await api.pagePublicMoonbreezes(body);
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
| **pagePublicMoonbreezesReq** | [PagePublicMoonbreezesReq](PagePublicMoonbreezesReq.md) |  | |

### Return type

[**PagePublicMoonbreezesResp**](PagePublicMoonbreezesResp.md)

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


## pageWatchingMoonbreezes

> PageWatchingMoonbreezesResp pageWatchingMoonbreezes(pageWatchingMoonbreezesReq)



服务方法。

### Example

```ts
import {
  Configuration,
  MoonbreezeService,
} from '@bass/bbs-sdk-fetch';
import type { PageWatchingMoonbreezesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new MoonbreezeService();

  const body = {
    // PageWatchingMoonbreezesReq
    pageWatchingMoonbreezesReq: ...,
  } satisfies PageWatchingMoonbreezesRequest;

  try {
    const data = await api.pageWatchingMoonbreezes(body);
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
| **pageWatchingMoonbreezesReq** | [PageWatchingMoonbreezesReq](PageWatchingMoonbreezesReq.md) |  | |

### Return type

[**PageWatchingMoonbreezesResp**](PageWatchingMoonbreezesResp.md)

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

