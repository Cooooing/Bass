# \OtpService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**begin_enable_totp**](OtpService.md#begin_enable_totp) | **POST** /v1/user/otp/totp/begin-enable | 
[**confirm_enable_totp**](OtpService.md#confirm_enable_totp) | **POST** /v1/user/otp/totp/confirm-enable | 
[**disable_totp**](OtpService.md#disable_totp) | **POST** /v1/user/otp/totp/disable | 
[**get_current_totp**](OtpService.md#get_current_totp) | **POST** /v1/user/otp/totp/get-current | 
[**send_email_otp**](OtpService.md#send_email_otp) | **POST** /v1/user/otp/email/send | 
[**send_phone_otp**](OtpService.md#send_phone_otp) | **POST** /v1/user/otp/phone/send | 



## begin_enable_totp

> models::BeginEnableTotpResp begin_enable_totp(body)


开始启用 TOTP。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::BeginEnableTotpResp**](BeginEnableTotp_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## confirm_enable_totp

> serde_json::Value confirm_enable_totp(confirm_enable_totp_req)


确认启用 TOTP。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**confirm_enable_totp_req** | [**ConfirmEnableTotpReq**](ConfirmEnableTotpReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## disable_totp

> serde_json::Value disable_totp(disable_totp_req)


关闭 TOTP。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**disable_totp_req** | [**DisableTotpReq**](DisableTotpReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_current_totp

> models::GetCurrentTotpResp get_current_totp(body)


获取当前账号 TOTP 状态。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentTotpResp**](GetCurrentTotp_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## send_email_otp

> models::SendEmailOtpResp send_email_otp(send_email_otp_req)


发送邮箱 OTP。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**send_email_otp_req** | [**SendEmailOtpReq**](SendEmailOtpReq.md) |  | [required] |

### Return type

[**models::SendEmailOtpResp**](SendEmailOtp_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## send_phone_otp

> models::SendPhoneOtpResp send_phone_otp(send_phone_otp_req)


发送手机 OTP。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**send_phone_otp_req** | [**SendPhoneOtpReq**](SendPhoneOtpReq.md) |  | [required] |

### Return type

[**models::SendPhoneOtpResp**](SendPhoneOtp_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

