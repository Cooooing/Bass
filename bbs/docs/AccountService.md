# \AccountService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_current**](AccountService.md#get_current) | **POST** /v1/user/account/get-current | 
[**get_profile**](AccountService.md#get_profile) | **POST** /v1/user/account/get-profile | 
[**update_profile**](AccountService.md#update_profile) | **POST** /v1/user/account/update-profile | 



## get_current

> models::GetCurrentAccountReply get_current(body)


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


## get_profile

> models::GetProfileAccountReply get_profile(get_profile_account_request)


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


## update_profile

> models::UpdateProfileAccountReply update_profile(update_profile_account_request)


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

