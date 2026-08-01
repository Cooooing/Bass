# OtpService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**beginEnableTotp**](OtpService.md#beginenabletotp) | **POST** /v1/user/otp/totp/begin-enable |  |
| [**confirmEnableTotp**](OtpService.md#confirmenabletotp) | **POST** /v1/user/otp/totp/confirm-enable |  |
| [**disableTotp**](OtpService.md#disabletotp) | **POST** /v1/user/otp/totp/disable |  |
| [**getCurrentTotp**](OtpService.md#getcurrenttotp) | **POST** /v1/user/otp/totp/get-current |  |
| [**sendEmailOtp**](OtpService.md#sendemailotp) | **POST** /v1/user/otp/email/send |  |
| [**sendPhoneOtp**](OtpService.md#sendphoneotp) | **POST** /v1/user/otp/phone/send |  |



## beginEnableTotp

> BeginEnableTotpResp beginEnableTotp(body)



开始启用 TOTP。

### Example

```ts
import {
  Configuration,
  OtpService,
} from '@bass/bbs-sdk-fetch';
import type { BeginEnableTotpRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new OtpService();

  const body = {
    // object
    body: Object,
  } satisfies BeginEnableTotpRequest;

  try {
    const data = await api.beginEnableTotp(body);
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

[**BeginEnableTotpResp**](BeginEnableTotpResp.md)

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


## confirmEnableTotp

> object confirmEnableTotp(confirmEnableTotpReq)



确认启用 TOTP。

### Example

```ts
import {
  Configuration,
  OtpService,
} from '@bass/bbs-sdk-fetch';
import type { ConfirmEnableTotpRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new OtpService();

  const body = {
    // ConfirmEnableTotpReq
    confirmEnableTotpReq: ...,
  } satisfies ConfirmEnableTotpRequest;

  try {
    const data = await api.confirmEnableTotp(body);
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
| **confirmEnableTotpReq** | [ConfirmEnableTotpReq](ConfirmEnableTotpReq.md) |  | |

### Return type

**object**

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


## disableTotp

> object disableTotp(disableTotpReq)



关闭 TOTP。

### Example

```ts
import {
  Configuration,
  OtpService,
} from '@bass/bbs-sdk-fetch';
import type { DisableTotpRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new OtpService();

  const body = {
    // DisableTotpReq
    disableTotpReq: ...,
  } satisfies DisableTotpRequest;

  try {
    const data = await api.disableTotp(body);
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
| **disableTotpReq** | [DisableTotpReq](DisableTotpReq.md) |  | |

### Return type

**object**

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


## getCurrentTotp

> GetCurrentTotpResp getCurrentTotp(body)



获取当前账号 TOTP 状态。

### Example

```ts
import {
  Configuration,
  OtpService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentTotpRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new OtpService();

  const body = {
    // object
    body: Object,
  } satisfies GetCurrentTotpRequest;

  try {
    const data = await api.getCurrentTotp(body);
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

[**GetCurrentTotpResp**](GetCurrentTotpResp.md)

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


## sendEmailOtp

> SendEmailOtpResp sendEmailOtp(sendEmailOtpReq)



发送邮箱 OTP。

### Example

```ts
import {
  Configuration,
  OtpService,
} from '@bass/bbs-sdk-fetch';
import type { SendEmailOtpRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new OtpService();

  const body = {
    // SendEmailOtpReq
    sendEmailOtpReq: ...,
  } satisfies SendEmailOtpRequest;

  try {
    const data = await api.sendEmailOtp(body);
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
| **sendEmailOtpReq** | [SendEmailOtpReq](SendEmailOtpReq.md) |  | |

### Return type

[**SendEmailOtpResp**](SendEmailOtpResp.md)

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


## sendPhoneOtp

> SendPhoneOtpResp sendPhoneOtp(sendPhoneOtpReq)



发送手机 OTP。

### Example

```ts
import {
  Configuration,
  OtpService,
} from '@bass/bbs-sdk-fetch';
import type { SendPhoneOtpRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new OtpService();

  const body = {
    // SendPhoneOtpReq
    sendPhoneOtpReq: ...,
  } satisfies SendPhoneOtpRequest;

  try {
    const data = await api.sendPhoneOtp(body);
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
| **sendPhoneOtpReq** | [SendPhoneOtpReq](SendPhoneOtpReq.md) |  | |

### Return type

[**SendPhoneOtpResp**](SendPhoneOtpResp.md)

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

