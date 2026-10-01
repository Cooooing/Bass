# MoonbreezeService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**create**](MoonbreezeService.md#create) | **POST** /v1/content/moonbreeze/create |  |
| [**pageMember**](MoonbreezeService.md#pagemember) | **POST** /v1/content/moonbreeze/page-member |  |
| [**pagePublic**](MoonbreezeService.md#pagepublic) | **POST** /v1/content/moonbreeze/page-public |  |
| [**pageWatching**](MoonbreezeService.md#pagewatching) | **POST** /v1/content/moonbreeze/page-watching |  |



## create

> CreateMoonbreezeResp create(createMoonbreezeReq)



### Example

```ts
import {
  Configuration,
  MoonbreezeService,
} from '@bass/bbs-sdk-fetch';
import type { CreateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new MoonbreezeService();

  const body = {
    // CreateMoonbreezeReq
    createMoonbreezeReq: ...,
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


## pageMember

> PageMoonbreezesResp pageMember(pageMemberMoonbreezesReq)



### Example

```ts
import {
  Configuration,
  MoonbreezeService,
} from '@bass/bbs-sdk-fetch';
import type { PageMemberRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new MoonbreezeService();

  const body = {
    // PageMemberMoonbreezesReq
    pageMemberMoonbreezesReq: ...,
  } satisfies PageMemberRequest;

  try {
    const data = await api.pageMember(body);
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

[**PageMoonbreezesResp**](PageMoonbreezesResp.md)

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


## pagePublic

> PageMoonbreezesResp pagePublic(pagePublicMoonbreezesReq)



### Example

```ts
import {
  Configuration,
  MoonbreezeService,
} from '@bass/bbs-sdk-fetch';
import type { PagePublicRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new MoonbreezeService();

  const body = {
    // PagePublicMoonbreezesReq
    pagePublicMoonbreezesReq: ...,
  } satisfies PagePublicRequest;

  try {
    const data = await api.pagePublic(body);
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

[**PageMoonbreezesResp**](PageMoonbreezesResp.md)

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


## pageWatching

> PageMoonbreezesResp pageWatching(pageWatchingMoonbreezesReq)



### Example

```ts
import {
  Configuration,
  MoonbreezeService,
} from '@bass/bbs-sdk-fetch';
import type { PageWatchingRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new MoonbreezeService();

  const body = {
    // PageWatchingMoonbreezesReq
    pageWatchingMoonbreezesReq: ...,
  } satisfies PageWatchingRequest;

  try {
    const data = await api.pageWatching(body);
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

[**PageMoonbreezesResp**](PageMoonbreezesResp.md)

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

