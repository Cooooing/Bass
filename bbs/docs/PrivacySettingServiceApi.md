# PrivacySettingServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**privacySettingServiceGetCurrent**](PrivacySettingServiceApi.md#privacysettingservicegetcurrent) | **POST** /v1/user/privacy-setting/get-current |  |
| [**privacySettingServiceUpdate**](PrivacySettingServiceApi.md#privacysettingserviceupdate) | **POST** /v1/user/privacy-setting/update-current |  |



## privacySettingServiceGetCurrent

> CommonApiAppBbsV1UserGetCurrentPrivacySettingReply privacySettingServiceGetCurrent(body)



获取当前登录账号的隐私设置

### Example

```ts
import {
  Configuration,
  PrivacySettingServiceApi,
} from '@bass/bbs-sdk';
import type { PrivacySettingServiceGetCurrentRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new PrivacySettingServiceApi();

  const body = {
    // object
    body: Object,
  } satisfies PrivacySettingServiceGetCurrentRequest;

  try {
    const data = await api.privacySettingServiceGetCurrent(body);
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

[**CommonApiAppBbsV1UserGetCurrentPrivacySettingReply**](CommonApiAppBbsV1UserGetCurrentPrivacySettingReply.md)

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


## privacySettingServiceUpdate

> CommonApiAppBbsV1UserUpdatePrivacySettingReply privacySettingServiceUpdate(commonApiAppBbsV1UserUpdatePrivacySettingRequest)



更新当前登录账号的隐私设置

### Example

```ts
import {
  Configuration,
  PrivacySettingServiceApi,
} from '@bass/bbs-sdk';
import type { PrivacySettingServiceUpdateRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new PrivacySettingServiceApi();

  const body = {
    // CommonApiAppBbsV1UserUpdatePrivacySettingRequest
    commonApiAppBbsV1UserUpdatePrivacySettingRequest: ...,
  } satisfies PrivacySettingServiceUpdateRequest;

  try {
    const data = await api.privacySettingServiceUpdate(body);
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
| **commonApiAppBbsV1UserUpdatePrivacySettingRequest** | [CommonApiAppBbsV1UserUpdatePrivacySettingRequest](CommonApiAppBbsV1UserUpdatePrivacySettingRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserUpdatePrivacySettingReply**](CommonApiAppBbsV1UserUpdatePrivacySettingReply.md)

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

