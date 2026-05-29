# \PreferencesServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**preferences_service_get_current**](PreferencesServiceApi.md#preferences_service_get_current) | **POST** /v1/user/preference/get-current | 
[**preferences_service_update_current**](PreferencesServiceApi.md#preferences_service_update_current) | **POST** /v1/user/preference/update-current | 



## preferences_service_get_current

> models::GetCurrentPreferencesReply preferences_service_get_current(body)


获取当前登录账号的偏好设置

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentPreferencesReply**](GetCurrentPreferences_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## preferences_service_update_current

> models::UpdateCurrentPreferencesReply preferences_service_update_current(update_current_preferences_request)


更新当前登录账号的偏好设置

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_current_preferences_request** | [**UpdateCurrentPreferencesRequest**](UpdateCurrentPreferencesRequest.md) |  | [required] |

### Return type

[**models::UpdateCurrentPreferencesReply**](UpdateCurrentPreferences_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

