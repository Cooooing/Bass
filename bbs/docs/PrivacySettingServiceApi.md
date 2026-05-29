# \PrivacySettingServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**privacy_setting_service_get_current**](PrivacySettingServiceApi.md#privacy_setting_service_get_current) | **POST** /v1/user/privacy-setting/get-current | 
[**privacy_setting_service_update_current**](PrivacySettingServiceApi.md#privacy_setting_service_update_current) | **POST** /v1/user/privacy-setting/update-current | 



## privacy_setting_service_get_current

> models::GetCurrentPrivacySettingReply privacy_setting_service_get_current(body)


获取当前登录账号的隐私设置

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentPrivacySettingReply**](GetCurrentPrivacySetting_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## privacy_setting_service_update_current

> models::UpdateCurrentPrivacySettingReply privacy_setting_service_update_current(update_current_privacy_setting_request)


更新当前登录账号的隐私设置

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_current_privacy_setting_request** | [**UpdateCurrentPrivacySettingRequest**](UpdateCurrentPrivacySettingRequest.md) |  | [required] |

### Return type

[**models::UpdateCurrentPrivacySettingReply**](UpdateCurrentPrivacySetting_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

