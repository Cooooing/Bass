# AccountServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**accountServiceBatchGetProfile**](AccountServiceApi.md#accountservicebatchgetprofile) | **POST** /v1/user/account/batch-get-profile |  |
| [**accountServiceGetCurrent**](AccountServiceApi.md#accountservicegetcurrent) | **POST** /v1/user/account/get-current |  |
| [**accountServiceGetProfile**](AccountServiceApi.md#accountservicegetprofile) | **POST** /v1/user/account/get-profile |  |
| [**accountServiceUpdateProfile**](AccountServiceApi.md#accountserviceupdateprofile) | **POST** /v1/user/account/update-profile |  |



## accountServiceBatchGetProfile

> CommonApiAppBbsV1UserBatchGetProfileAccountReply accountServiceBatchGetProfile(commonApiAppBbsV1UserBatchGetProfileAccountRequest)



批量获取账号展示资料

### Example

```ts
import {
  Configuration,
  AccountServiceApi,
} from '@bass/bbs-sdk';
import type { AccountServiceBatchGetProfileRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AccountServiceApi();

  const body = {
    // CommonApiAppBbsV1UserBatchGetProfileAccountRequest
    commonApiAppBbsV1UserBatchGetProfileAccountRequest: ...,
  } satisfies AccountServiceBatchGetProfileRequest;

  try {
    const data = await api.accountServiceBatchGetProfile(body);
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
| **commonApiAppBbsV1UserBatchGetProfileAccountRequest** | [CommonApiAppBbsV1UserBatchGetProfileAccountRequest](CommonApiAppBbsV1UserBatchGetProfileAccountRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserBatchGetProfileAccountReply**](CommonApiAppBbsV1UserBatchGetProfileAccountReply.md)

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


## accountServiceGetCurrent

> CommonApiAppBbsV1UserGetCurrentAccountReply accountServiceGetCurrent(body)



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

[**CommonApiAppBbsV1UserGetCurrentAccountReply**](CommonApiAppBbsV1UserGetCurrentAccountReply.md)

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

> CommonApiAppBbsV1UserGetProfileAccountReply accountServiceGetProfile(commonApiAppBbsV1UserGetProfileAccountRequest)



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
    // CommonApiAppBbsV1UserGetProfileAccountRequest
    commonApiAppBbsV1UserGetProfileAccountRequest: ...,
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
| **commonApiAppBbsV1UserGetProfileAccountRequest** | [CommonApiAppBbsV1UserGetProfileAccountRequest](CommonApiAppBbsV1UserGetProfileAccountRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserGetProfileAccountReply**](CommonApiAppBbsV1UserGetProfileAccountReply.md)

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

> CommonApiAppBbsV1UserUpdateProfileAccountReply accountServiceUpdateProfile(commonApiAppBbsV1UserUpdateProfileAccountRequest)



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
    // CommonApiAppBbsV1UserUpdateProfileAccountRequest
    commonApiAppBbsV1UserUpdateProfileAccountRequest: ...,
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
| **commonApiAppBbsV1UserUpdateProfileAccountRequest** | [CommonApiAppBbsV1UserUpdateProfileAccountRequest](CommonApiAppBbsV1UserUpdateProfileAccountRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserUpdateProfileAccountReply**](CommonApiAppBbsV1UserUpdateProfileAccountReply.md)

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

