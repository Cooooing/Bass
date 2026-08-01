# OtpService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**beginEnableTotp**](#beginenabletotp) | **POST** /v1/user/otp/totp/begin-enable | |
|[**confirmEnableTotp**](#confirmenabletotp) | **POST** /v1/user/otp/totp/confirm-enable | |
|[**disableTotp**](#disabletotp) | **POST** /v1/user/otp/totp/disable | |
|[**getCurrentTotp**](#getcurrenttotp) | **POST** /v1/user/otp/totp/get-current | |
|[**sendEmailOtp**](#sendemailotp) | **POST** /v1/user/otp/email/send | |
|[**sendPhoneOtp**](#sendphoneotp) | **POST** /v1/user/otp/phone/send | |

# **beginEnableTotp**
> BeginEnableTotpResp beginEnableTotp(body)

开始启用 TOTP。

### Example

```typescript
import {
    OtpService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new OtpService(configuration);

let body: object; //

const { status, data } = await apiInstance.beginEnableTotp(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**BeginEnableTotpResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **confirmEnableTotp**
> object confirmEnableTotp(confirmEnableTotpReq)

确认启用 TOTP。

### Example

```typescript
import {
    OtpService,
    Configuration,
    ConfirmEnableTotpReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new OtpService(configuration);

let confirmEnableTotpReq: ConfirmEnableTotpReq; //

const { status, data } = await apiInstance.confirmEnableTotp(
    confirmEnableTotpReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **confirmEnableTotpReq** | **ConfirmEnableTotpReq**|  | |


### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **disableTotp**
> object disableTotp(disableTotpReq)

关闭 TOTP。

### Example

```typescript
import {
    OtpService,
    Configuration,
    DisableTotpReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new OtpService(configuration);

let disableTotpReq: DisableTotpReq; //

const { status, data } = await apiInstance.disableTotp(
    disableTotpReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **disableTotpReq** | **DisableTotpReq**|  | |


### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCurrentTotp**
> GetCurrentTotpResp getCurrentTotp(body)

获取当前账号 TOTP 状态。

### Example

```typescript
import {
    OtpService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new OtpService(configuration);

let body: object; //

const { status, data } = await apiInstance.getCurrentTotp(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**GetCurrentTotpResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sendEmailOtp**
> SendEmailOtpResp sendEmailOtp(sendEmailOtpReq)

发送邮箱 OTP。

### Example

```typescript
import {
    OtpService,
    Configuration,
    SendEmailOtpReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new OtpService(configuration);

let sendEmailOtpReq: SendEmailOtpReq; //

const { status, data } = await apiInstance.sendEmailOtp(
    sendEmailOtpReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **sendEmailOtpReq** | **SendEmailOtpReq**|  | |


### Return type

**SendEmailOtpResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sendPhoneOtp**
> SendPhoneOtpResp sendPhoneOtp(sendPhoneOtpReq)

发送手机 OTP。

### Example

```typescript
import {
    OtpService,
    Configuration,
    SendPhoneOtpReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new OtpService(configuration);

let sendPhoneOtpReq: SendPhoneOtpReq; //

const { status, data } = await apiInstance.sendPhoneOtp(
    sendPhoneOtpReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **sendPhoneOtpReq** | **SendPhoneOtpReq**|  | |


### Return type

**SendPhoneOtpResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

