# AccountServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**accountServiceGetCurrent**](AccountServiceApi.md#accountservicegetcurrent) | **POST** /v1/user/account/get-current |  |
| [**accountServiceGetProfile**](AccountServiceApi.md#accountservicegetprofile) | **POST** /v1/user/account/get-profile |  |
| [**accountServiceUpdateProfile**](AccountServiceApi.md#accountserviceupdateprofile) | **POST** /v1/user/account/update-profile |  |



## accountServiceGetCurrent

> GetCurrentAccountReply accountServiceGetCurrent(body)



获取当前登录账号的完整资料

### Example

```ts
import {
  Configuration,
  AccountServiceApi,
} from '@bass/bbs-sdk';
import type { AccountServiceGetCurrentRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AccountServiceApi();

  const body = {
    // object
    body: Object,
  } satisfies AccountServiceGetCurrentRequest;

  try {
    const data = await api.accountServiceGetCurrent(body);
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


## accountServiceGetProfile

> GetProfileAccountReply accountServiceGetProfile(getProfileAccountRequest)



按账号 ID 获取账号展示资料

### Example

```ts
import {
  Configuration,
  AccountServiceApi,
} from '@bass/bbs-sdk';
import type { AccountServiceGetProfileRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AccountServiceApi();

  const body = {
    // GetProfileAccountRequest
    getProfileAccountRequest: ...,
  } satisfies AccountServiceGetProfileRequest;

  try {
    const data = await api.accountServiceGetProfile(body);
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


## accountServiceUpdateProfile

> UpdateProfileAccountReply accountServiceUpdateProfile(updateProfileAccountRequest)



更新当前登录账号的展示资料

### Example

```ts
import {
  Configuration,
  AccountServiceApi,
} from '@bass/bbs-sdk';
import type { AccountServiceUpdateProfileRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AccountServiceApi();

  const body = {
    // UpdateProfileAccountRequest
    updateProfileAccountRequest: ...,
  } satisfies AccountServiceUpdateProfileRequest;

  try {
    const data = await api.accountServiceUpdateProfile(body);
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

