# AccountService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**avatar**](AccountService.md#avatar) | **GET** /v1/user/account/avatar |  |
| [**getCurrent**](AccountService.md#getcurrent) | **POST** /v1/user/account/get-current |  |
| [**getProfile**](AccountService.md#getprofile) | **POST** /v1/user/account/get-profile |  |
| [**updateProfile**](AccountService.md#updateprofile) | **POST** /v1/user/account/update-profile |  |



## avatar

> ImageResp avatar(name)



生成默认账号头像。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { AvatarRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // string (optional)
    name: name_example,
  } satisfies AvatarRequest;

  try {
    const data = await api.avatar(body);
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
| **name** | `string` |  | [Optional] [Defaults to `undefined`] |

### Return type

[**ImageResp**](ImageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## getCurrent

> GetCurrentAccountResp getCurrent(body)



获取当前账号的完整资料。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // object
    body: Object,
  } satisfies GetCurrentRequest;

  try {
    const data = await api.getCurrent(body);
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
| **body** | `object` |  | |

### Return type

[**GetCurrentAccountResp**](GetCurrentAccountResp.md)

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


## getProfile

> GetProfileAccountResp getProfile(getProfileAccountReq)



按账号 ID 获取账号展示资料。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { GetProfileRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // GetProfileAccountReq
    getProfileAccountReq: ...,
  } satisfies GetProfileRequest;

  try {
    const data = await api.getProfile(body);
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
| **getProfileAccountReq** | [GetProfileAccountReq](GetProfileAccountReq.md) |  | |

### Return type

[**GetProfileAccountResp**](GetProfileAccountResp.md)

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


## updateProfile

> UpdateProfileAccountResp updateProfile(updateProfileAccountReq)



更新当前账号的展示资料。

### Example

```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateProfileRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // UpdateProfileAccountReq
    updateProfileAccountReq: ...,
  } satisfies UpdateProfileRequest;

  try {
    const data = await api.updateProfile(body);
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
| **updateProfileAccountReq** | [UpdateProfileAccountReq](UpdateProfileAccountReq.md) |  | |

### Return type

[**UpdateProfileAccountResp**](UpdateProfileAccountResp.md)

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

