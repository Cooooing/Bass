# \PreferencesService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_current**](PreferencesService.md#get_current) | **POST** /v1/user/preference/get-current | 
[**update_current**](PreferencesService.md#update_current) | **POST** /v1/user/preference/update-current | 



## get_current

> models::GetCurrentPreferencesReply get_current(body)


获取当前账号的偏好设置。

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


## update_current

> models::UpdateCurrentPreferencesReply update_current(update_current_preferences_request)


更新当前账号的偏好设置。

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

