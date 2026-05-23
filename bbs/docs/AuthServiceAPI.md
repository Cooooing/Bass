# \AuthServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthServiceLoginPassword**](AuthServiceAPI.md#AuthServiceLoginPassword) | **Post** /v1/user/auth/login-password | 
[**AuthServiceLogout**](AuthServiceAPI.md#AuthServiceLogout) | **Post** /v1/user/auth/logout | 
[**AuthServiceRegisterEmail**](AuthServiceAPI.md#AuthServiceRegisterEmail) | **Post** /v1/user/auth/register-email | 
[**AuthServiceRegisterPhone**](AuthServiceAPI.md#AuthServiceRegisterPhone) | **Post** /v1/user/auth/register-phone | 
[**AuthServiceVerifyEmailRegister**](AuthServiceAPI.md#AuthServiceVerifyEmailRegister) | **Post** /v1/user/auth/verify-email-register | 
[**AuthServiceVerifyPhoneRegister**](AuthServiceAPI.md#AuthServiceVerifyPhoneRegister) | **Post** /v1/user/auth/verify-phone-register | 



## AuthServiceLoginPassword

> CommonApiAppBbsV1UserLoginPasswordReply AuthServiceLoginPassword(ctx).CommonApiAppBbsV1UserLoginPasswordRequest(commonApiAppBbsV1UserLoginPasswordRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	commonApiAppBbsV1UserLoginPasswordRequest := *openapiclient.NewCommonApiAppBbsV1UserLoginPasswordRequest() // CommonApiAppBbsV1UserLoginPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceLoginPassword(context.Background()).CommonApiAppBbsV1UserLoginPasswordRequest(commonApiAppBbsV1UserLoginPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceLoginPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceLoginPassword`: CommonApiAppBbsV1UserLoginPasswordReply
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceLoginPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceLoginPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserLoginPasswordRequest** | [**CommonApiAppBbsV1UserLoginPasswordRequest**](CommonApiAppBbsV1UserLoginPasswordRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserLoginPasswordReply**](CommonApiAppBbsV1UserLoginPasswordReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceLogout

> map[string]interface{} AuthServiceLogout(ctx).Body(body).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceLogout(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceLogout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceLogout`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceLogout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceRegisterEmail

> CommonApiAppBbsV1UserRegisterEmailReply AuthServiceRegisterEmail(ctx).CommonApiAppBbsV1UserRegisterEmailRequest(commonApiAppBbsV1UserRegisterEmailRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	commonApiAppBbsV1UserRegisterEmailRequest := *openapiclient.NewCommonApiAppBbsV1UserRegisterEmailRequest() // CommonApiAppBbsV1UserRegisterEmailRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceRegisterEmail(context.Background()).CommonApiAppBbsV1UserRegisterEmailRequest(commonApiAppBbsV1UserRegisterEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceRegisterEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceRegisterEmail`: CommonApiAppBbsV1UserRegisterEmailReply
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceRegisterEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceRegisterEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserRegisterEmailRequest** | [**CommonApiAppBbsV1UserRegisterEmailRequest**](CommonApiAppBbsV1UserRegisterEmailRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserRegisterEmailReply**](CommonApiAppBbsV1UserRegisterEmailReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceRegisterPhone

> CommonApiAppBbsV1UserRegisterPhoneReply AuthServiceRegisterPhone(ctx).CommonApiAppBbsV1UserRegisterPhoneRequest(commonApiAppBbsV1UserRegisterPhoneRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	commonApiAppBbsV1UserRegisterPhoneRequest := *openapiclient.NewCommonApiAppBbsV1UserRegisterPhoneRequest() // CommonApiAppBbsV1UserRegisterPhoneRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceRegisterPhone(context.Background()).CommonApiAppBbsV1UserRegisterPhoneRequest(commonApiAppBbsV1UserRegisterPhoneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceRegisterPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceRegisterPhone`: CommonApiAppBbsV1UserRegisterPhoneReply
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceRegisterPhone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceRegisterPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserRegisterPhoneRequest** | [**CommonApiAppBbsV1UserRegisterPhoneRequest**](CommonApiAppBbsV1UserRegisterPhoneRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserRegisterPhoneReply**](CommonApiAppBbsV1UserRegisterPhoneReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceVerifyEmailRegister

> map[string]interface{} AuthServiceVerifyEmailRegister(ctx).CommonApiAppBbsV1UserVerifyEmailRegisterRequest(commonApiAppBbsV1UserVerifyEmailRegisterRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	commonApiAppBbsV1UserVerifyEmailRegisterRequest := *openapiclient.NewCommonApiAppBbsV1UserVerifyEmailRegisterRequest() // CommonApiAppBbsV1UserVerifyEmailRegisterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceVerifyEmailRegister(context.Background()).CommonApiAppBbsV1UserVerifyEmailRegisterRequest(commonApiAppBbsV1UserVerifyEmailRegisterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceVerifyEmailRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceVerifyEmailRegister`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceVerifyEmailRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceVerifyEmailRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserVerifyEmailRegisterRequest** | [**CommonApiAppBbsV1UserVerifyEmailRegisterRequest**](CommonApiAppBbsV1UserVerifyEmailRegisterRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthServiceVerifyPhoneRegister

> map[string]interface{} AuthServiceVerifyPhoneRegister(ctx).CommonApiAppBbsV1UserVerifyPhoneRegisterRequest(commonApiAppBbsV1UserVerifyPhoneRegisterRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	commonApiAppBbsV1UserVerifyPhoneRegisterRequest := *openapiclient.NewCommonApiAppBbsV1UserVerifyPhoneRegisterRequest() // CommonApiAppBbsV1UserVerifyPhoneRegisterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceVerifyPhoneRegister(context.Background()).CommonApiAppBbsV1UserVerifyPhoneRegisterRequest(commonApiAppBbsV1UserVerifyPhoneRegisterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceVerifyPhoneRegister``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceVerifyPhoneRegister`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceVerifyPhoneRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceVerifyPhoneRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserVerifyPhoneRegisterRequest** | [**CommonApiAppBbsV1UserVerifyPhoneRegisterRequest**](CommonApiAppBbsV1UserVerifyPhoneRegisterRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

