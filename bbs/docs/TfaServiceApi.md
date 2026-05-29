# \TfaServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**tfa_service_begin_enable**](TfaServiceApi.md#tfa_service_begin_enable) | **POST** /v1/user/tfa/begin-enable | 
[**tfa_service_confirm_enable**](TfaServiceApi.md#tfa_service_confirm_enable) | **POST** /v1/user/tfa/confirm-enable | 
[**tfa_service_disable**](TfaServiceApi.md#tfa_service_disable) | **POST** /v1/user/tfa/disable | 
[**tfa_service_get_current**](TfaServiceApi.md#tfa_service_get_current) | **POST** /v1/user/tfa/get-current | 
[**tfa_service_validate**](TfaServiceApi.md#tfa_service_validate) | **POST** /v1/user/tfa/validate | 



## tfa_service_begin_enable

> models::BeginEnableTfaReply tfa_service_begin_enable(body)


开始启用二步验证

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::BeginEnableTfaReply**](BeginEnableTfa_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## tfa_service_confirm_enable

> serde_json::Value tfa_service_confirm_enable(confirm_enable_tfa_request)


确认二步验证码并正式启用二步验证

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**confirm_enable_tfa_request** | [**ConfirmEnableTfaRequest**](ConfirmEnableTfaRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## tfa_service_disable

> serde_json::Value tfa_service_disable(disable_tfa_request)


校验二步验证码并关闭二步验证

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**disable_tfa_request** | [**DisableTfaRequest**](DisableTfaRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## tfa_service_get_current

> models::GetCurrentTfaReply tfa_service_get_current(body)


获取当前登录账号的二步验证状态

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentTfaReply**](GetCurrentTfa_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## tfa_service_validate

> models::ValidateTfaReply tfa_service_validate(validate_tfa_request)


校验当前登录账号的二步验证码

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**validate_tfa_request** | [**ValidateTfaRequest**](ValidateTfaRequest.md) |  | [required] |

### Return type

[**models::ValidateTfaReply**](ValidateTfa_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

