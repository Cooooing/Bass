# AccountService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**avatar**](AccountService.md#avatar) | **GET** /v1/user/account/avatar |  |
| [**getCurrent**](AccountService.md#getcurrent) | **POST** /v1/user/account/get-current |  |
| [**getProfile**](AccountService.md#getprofile) | **POST** /v1/user/account/get-profile |  |
| [**updateProfile**](AccountService.md#updateprofile) | **POST** /v1/user/account/update-profile |  |



## avatar

> ImageReply avatar(name)



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
    // string | 用于生成头像的账号名。 (optional)
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
| **name** | `string` | 用于生成头像的账号名。 | [Optional] [Defaults to `undefined`] |

### Return type

[**ImageReply**](ImageReply.md)

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

> GetCurrentAccountReply getCurrent(body)



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

[**GetCurrentAccountReply**](GetCurrentAccountReply.md)

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

> GetProfileAccountReply getProfile(getProfileAccountRequest)



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
    // GetProfileAccountRequest
    getProfileAccountRequest: ...,
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
| **getProfileAccountRequest** | [GetProfileAccountRequest](GetProfileAccountRequest.md) |  | |

### Return type

[**GetProfileAccountReply**](GetProfileAccountReply.md)

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

> UpdateProfileAccountReply updateProfile(updateProfileAccountRequest)



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
    // UpdateProfileAccountRequest
    updateProfileAccountRequest: ...,
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
| **updateProfileAccountRequest** | [UpdateProfileAccountRequest](UpdateProfileAccountRequest.md) |  | |

### Return type

[**UpdateProfileAccountReply**](UpdateProfileAccountReply.md)

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

