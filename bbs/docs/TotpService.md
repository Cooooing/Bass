# \TotpService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**begin_enable**](TotpService.md#begin_enable) | **POST** /v1/user/totp/begin-enable | 
[**confirm_enable**](TotpService.md#confirm_enable) | **POST** /v1/user/totp/confirm-enable | 
[**disable**](TotpService.md#disable) | **POST** /v1/user/totp/disable | 
[**get_current**](TotpService.md#get_current) | **POST** /v1/user/totp/get-current | 



## begin_enable

> models::BeginEnableTotpReply begin_enable(body)


开始启用 TOTP。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::BeginEnableTotpReply**](BeginEnableTotp_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## confirm_enable

> serde_json::Value confirm_enable(confirm_enable_totp_request)


确认 TOTP 验证码并正式启用 TOTP。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**confirm_enable_totp_request** | [**ConfirmEnableTotpRequest**](ConfirmEnableTotpRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## disable

> serde_json::Value disable(disable_totp_request)


校验 TOTP 验证码并关闭 TOTP。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**disable_totp_request** | [**DisableTotpRequest**](DisableTotpRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_current

> models::GetCurrentTotpReply get_current(body)


获取当前账号的 TOTP 状态。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentTotpReply**](GetCurrentTotp_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

