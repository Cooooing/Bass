# \PrivacySettingService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_current**](PrivacySettingService.md#get_current) | **POST** /v1/user/privacy-setting/get-current | 
[**update_current**](PrivacySettingService.md#update_current) | **POST** /v1/user/privacy-setting/update-current | 



## get_current

> models::GetCurrentPrivacySettingReply get_current(body)


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


## update_current

> models::UpdateCurrentPrivacySettingReply update_current(update_current_privacy_setting_request)


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

