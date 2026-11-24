# \AuthService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**login_account**](AuthService.md#login_account) | **POST** /v1/game-idle/auth/login | 
[**register_account**](AuthService.md#register_account) | **POST** /v1/game-idle/auth/register | 



## login_account

> models::LoginAccountResp login_account(login_account_req)


登录账号。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**login_account_req** | [**LoginAccountReq**](LoginAccountReq.md) |  | [required] |

### Return type

[**models::LoginAccountResp**](LoginAccount_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## register_account

> serde_json::Value register_account(register_account_req)


注册账号。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**register_account_req** | [**RegisterAccountReq**](RegisterAccountReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

