# \AccountServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**account_service_get_current**](AccountServiceApi.md#account_service_get_current) | **POST** /v1/user/account/get-current | 
[**account_service_get_profile**](AccountServiceApi.md#account_service_get_profile) | **POST** /v1/user/account/get-profile | 
[**account_service_update_profile**](AccountServiceApi.md#account_service_update_profile) | **POST** /v1/user/account/update-profile | 



## account_service_get_current

> models::GetCurrentAccountReply account_service_get_current(body)


获取当前登录账号的完整资料

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentAccountReply**](GetCurrentAccount_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## account_service_get_profile

> models::GetProfileAccountReply account_service_get_profile(get_profile_account_request)


按账号 ID 获取账号展示资料

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**get_profile_account_request** | [**GetProfileAccountRequest**](GetProfileAccountRequest.md) |  | [required] |

### Return type

[**models::GetProfileAccountReply**](GetProfileAccount_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## account_service_update_profile

> models::UpdateProfileAccountReply account_service_update_profile(update_profile_account_request)


更新当前登录账号的展示资料

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_profile_account_request** | [**UpdateProfileAccountRequest**](UpdateProfileAccountRequest.md) |  | [required] |

### Return type

[**models::UpdateProfileAccountReply**](UpdateProfileAccount_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

