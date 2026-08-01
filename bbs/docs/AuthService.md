# \AuthService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**cancel_account**](AuthService.md#cancel_account) | **POST** /v1/user/auth/cancel-account | 
[**login**](AuthService.md#login) | **POST** /v1/user/auth/login | 
[**logout**](AuthService.md#logout) | **POST** /v1/user/auth/logout | 
[**refresh_token**](AuthService.md#refresh_token) | **POST** /v1/user/auth/refresh-token | 
[**start_email_registration**](AuthService.md#start_email_registration) | **POST** /v1/user/auth/start-email-registration | 
[**start_phone_registration**](AuthService.md#start_phone_registration) | **POST** /v1/user/auth/start-phone-registration | 
[**verify_email_registration**](AuthService.md#verify_email_registration) | **POST** /v1/user/auth/verify-email-registration | 
[**verify_phone_registration**](AuthService.md#verify_phone_registration) | **POST** /v1/user/auth/verify-phone-registration | 



## cancel_account

> serde_json::Value cancel_account(cancel_account_req)


注销账号。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**cancel_account_req** | [**CancelAccountReq**](CancelAccountReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## login

> models::LoginResp login(login_req)


登录账号。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**login_req** | [**LoginReq**](LoginReq.md) |  | [required] |

### Return type

[**models::LoginResp**](Login_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## logout

> serde_json::Value logout(body)


退出登录。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## refresh_token

> models::RefreshTokenResp refresh_token(refresh_token_req)


刷新登录令牌。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**refresh_token_req** | [**RefreshTokenReq**](RefreshTokenReq.md) |  | [required] |

### Return type

[**models::RefreshTokenResp**](RefreshToken_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## start_email_registration

> models::StartEmailRegistrationResp start_email_registration(start_email_registration_req)


开始邮箱注册。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**start_email_registration_req** | [**StartEmailRegistrationReq**](StartEmailRegistrationReq.md) |  | [required] |

### Return type

[**models::StartEmailRegistrationResp**](StartEmailRegistration_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## start_phone_registration

> models::StartPhoneRegistrationResp start_phone_registration(start_phone_registration_req)


开始手机注册。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**start_phone_registration_req** | [**StartPhoneRegistrationReq**](StartPhoneRegistrationReq.md) |  | [required] |

### Return type

[**models::StartPhoneRegistrationResp**](StartPhoneRegistration_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## verify_email_registration

> serde_json::Value verify_email_registration(verify_email_registration_req)


校验邮箱注册验证码。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**verify_email_registration_req** | [**VerifyEmailRegistrationReq**](VerifyEmailRegistrationReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## verify_phone_registration

> serde_json::Value verify_phone_registration(verify_phone_registration_req)


校验手机注册验证码。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**verify_phone_registration_req** | [**VerifyPhoneRegistrationReq**](VerifyPhoneRegistrationReq.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

