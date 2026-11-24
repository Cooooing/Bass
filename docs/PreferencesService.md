# \PreferencesService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_current_preferences**](PreferencesService.md#get_current_preferences) | **POST** /v1/user/preference/get-current | 
[**update_current_preferences**](PreferencesService.md#update_current_preferences) | **POST** /v1/user/preference/update-current | 



## get_current_preferences

> models::GetCurrentPreferencesResp get_current_preferences(body)


获取当前账号的偏好设置。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentPreferencesResp**](GetCurrentPreferences_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_current_preferences

> models::UpdateCurrentPreferencesResp update_current_preferences(update_current_preferences_req)


更新当前账号的偏好设置。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_current_preferences_req** | [**UpdateCurrentPreferencesReq**](UpdateCurrentPreferencesReq.md) |  | [required] |

### Return type

[**models::UpdateCurrentPreferencesResp**](UpdateCurrentPreferences_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

