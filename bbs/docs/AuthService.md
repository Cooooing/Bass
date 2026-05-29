# \AuthService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**login_by_password**](AuthService.md#login_by_password) | **POST** /v1/user/auth/login-by-password | 
[**logout**](AuthService.md#logout) | **POST** /v1/user/auth/logout | 
[**start_email_registration**](AuthService.md#start_email_registration) | **POST** /v1/user/auth/start-email-registration | 
[**start_phone_registration**](AuthService.md#start_phone_registration) | **POST** /v1/user/auth/start-phone-registration | 
[**verify_email_registration**](AuthService.md#verify_email_registration) | **POST** /v1/user/auth/verify-email-registration | 
[**verify_phone_registration**](AuthService.md#verify_phone_registration) | **POST** /v1/user/auth/verify-phone-registration | 



## login_by_password

> models::LoginByPasswordReply login_by_password(login_by_password_request)


使用密码登录账号

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**login_by_password_request** | [**LoginByPasswordRequest**](LoginByPasswordRequest.md) |  | [required] |

### Return type

[**models::LoginByPasswordReply**](LoginByPassword_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## logout

> serde_json::Value logout(body)


登出当前登录账号

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


## start_email_registration

> models::StartEmailRegistrationReply start_email_registration(start_email_registration_request)


使用邮箱发起账号注册

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**start_email_registration_request** | [**StartEmailRegistrationRequest**](StartEmailRegistrationRequest.md) |  | [required] |

### Return type

[**models::StartEmailRegistrationReply**](StartEmailRegistration_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## start_phone_registration

> models::StartPhoneRegistrationReply start_phone_registration(start_phone_registration_request)


使用手机号发起账号注册

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**start_phone_registration_request** | [**StartPhoneRegistrationRequest**](StartPhoneRegistrationRequest.md) |  | [required] |

### Return type

[**models::StartPhoneRegistrationReply**](StartPhoneRegistration_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## verify_email_registration

> serde_json::Value verify_email_registration(verify_email_registration_request)


校验邮箱注册验证码

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**verify_email_registration_request** | [**VerifyEmailRegistrationRequest**](VerifyEmailRegistrationRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## verify_phone_registration

> serde_json::Value verify_phone_registration(verify_phone_registration_request)


校验手机号注册验证码

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**verify_phone_registration_request** | [**VerifyPhoneRegistrationRequest**](VerifyPhoneRegistrationRequest.md) |  | [required] |

### Return type

[**serde_json::Value**](serde_json::Value.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

