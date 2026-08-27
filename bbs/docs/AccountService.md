# \AccountService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**avatar**](AccountService.md#avatar) | **GET** /v1/user/account/avatar | 
[**get_current**](AccountService.md#get_current) | **POST** /v1/user/account/get-current | 
[**get_profile**](AccountService.md#get_profile) | **POST** /v1/user/account/get-profile | 
[**update_email**](AccountService.md#update_email) | **POST** /v1/user/account/update-email | 
[**update_password**](AccountService.md#update_password) | **POST** /v1/user/account/update-password | 
[**update_phone**](AccountService.md#update_phone) | **POST** /v1/user/account/update-phone | 
[**update_profile**](AccountService.md#update_profile) | **POST** /v1/user/account/update-profile | 



## avatar

> models::ImageResp avatar(name)


生成默认账号头像

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**name** | Option<**String**> |  |  |

### Return type

[**models::ImageResp**](ImageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_current

> models::GetCurrentAccountResp get_current(body)


获取当前账号完整资料

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::GetCurrentAccountResp**](GetCurrentAccount_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_profile

> models::GetProfileAccountResp get_profile(get_profile_account_req)


按账号 ID 获取展示资料

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**get_profile_account_req** | [**GetProfileAccountReq**](GetProfileAccountReq.md) |  | [required] |

### Return type

[**models::GetProfileAccountResp**](GetProfileAccount_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_email

> serde_json::Value update_email(update_email_account_req)


更新当前账号邮箱

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_email_account_req** | [**UpdateEmailAccountReq**](UpdateEmailAccountReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_password

> serde_json::Value update_password(update_password_account_req)


更新当前账号密码

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_password_account_req** | [**UpdatePasswordAccountReq**](UpdatePasswordAccountReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_phone

> serde_json::Value update_phone(update_phone_account_req)


更新当前账号手机号

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_phone_account_req** | [**UpdatePhoneAccountReq**](UpdatePhoneAccountReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_profile

> models::UpdateProfileAccountResp update_profile(update_profile_account_req)


更新当前账号展示资料

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**update_profile_account_req** | [**UpdateProfileAccountReq**](UpdateProfileAccountReq.md) |  | [required] |

### Return type

[**models::UpdateProfileAccountResp**](UpdateProfileAccount_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

