# \AuthService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginByPassword**](AuthService.md#LoginByPassword) | **Post** /v1/user/auth/login-by-password | 
[**Logout**](AuthService.md#Logout) | **Post** /v1/user/auth/logout | 
[**StartEmailRegistration**](AuthService.md#StartEmailRegistration) | **Post** /v1/user/auth/start-email-registration | 
[**StartPhoneRegistration**](AuthService.md#StartPhoneRegistration) | **Post** /v1/user/auth/start-phone-registration | 
[**VerifyEmailRegistration**](AuthService.md#VerifyEmailRegistration) | **Post** /v1/user/auth/verify-email-registration | 
[**VerifyPhoneRegistration**](AuthService.md#VerifyPhoneRegistration) | **Post** /v1/user/auth/verify-phone-registration | 



## LoginByPassword

> LoginByPasswordReply LoginByPassword(ctx).LoginByPasswordRequest(loginByPasswordRequest).Execute()





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
	loginByPasswordRequest := *openapiclient.NewLoginByPasswordRequest("Account_example", "Password_example") // LoginByPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.LoginByPassword(context.Background()).LoginByPasswordRequest(loginByPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.LoginByPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginByPassword`: LoginByPasswordReply
	fmt.Fprintf(os.Stdout, "Response from `AuthService.LoginByPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginByPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginByPasswordRequest** | [**LoginByPasswordRequest**](LoginByPasswordRequest.md) |  | 

### Return type

[**LoginByPasswordReply**](LoginByPasswordReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> map[string]interface{} Logout(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.AuthService.Logout(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.Logout``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Logout`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthService.Logout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## StartEmailRegistration

> StartEmailRegistrationReply StartEmailRegistration(ctx).StartEmailRegistrationRequest(startEmailRegistrationRequest).Execute()





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
	startEmailRegistrationRequest := *openapiclient.NewStartEmailRegistrationRequest("Email_example", "Password_example", "Name_example") // StartEmailRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.StartEmailRegistration(context.Background()).StartEmailRegistrationRequest(startEmailRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.StartEmailRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartEmailRegistration`: StartEmailRegistrationReply
	fmt.Fprintf(os.Stdout, "Response from `AuthService.StartEmailRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartEmailRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startEmailRegistrationRequest** | [**StartEmailRegistrationRequest**](StartEmailRegistrationRequest.md) |  | 

### Return type

[**StartEmailRegistrationReply**](StartEmailRegistrationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartPhoneRegistration

> StartPhoneRegistrationReply StartPhoneRegistration(ctx).StartPhoneRegistrationRequest(startPhoneRegistrationRequest).Execute()





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
	startPhoneRegistrationRequest := *openapiclient.NewStartPhoneRegistrationRequest("Phone_example", "Password_example", "Name_example") // StartPhoneRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.StartPhoneRegistration(context.Background()).StartPhoneRegistrationRequest(startPhoneRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.StartPhoneRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartPhoneRegistration`: StartPhoneRegistrationReply
	fmt.Fprintf(os.Stdout, "Response from `AuthService.StartPhoneRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartPhoneRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startPhoneRegistrationRequest** | [**StartPhoneRegistrationRequest**](StartPhoneRegistrationRequest.md) |  | 

### Return type

[**StartPhoneRegistrationReply**](StartPhoneRegistrationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEmailRegistration

> map[string]interface{} VerifyEmailRegistration(ctx).VerifyEmailRegistrationRequest(verifyEmailRegistrationRequest).Execute()





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
	verifyEmailRegistrationRequest := *openapiclient.NewVerifyEmailRegistrationRequest("Code_example", "CodeToken_example") // VerifyEmailRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.VerifyEmailRegistration(context.Background()).VerifyEmailRegistrationRequest(verifyEmailRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.VerifyEmailRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyEmailRegistration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthService.VerifyEmailRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEmailRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyEmailRegistrationRequest** | [**VerifyEmailRegistrationRequest**](VerifyEmailRegistrationRequest.md) |  | 

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


## VerifyPhoneRegistration

> map[string]interface{} VerifyPhoneRegistration(ctx).VerifyPhoneRegistrationRequest(verifyPhoneRegistrationRequest).Execute()





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
	verifyPhoneRegistrationRequest := *openapiclient.NewVerifyPhoneRegistrationRequest("Code_example", "CodeToken_example") // VerifyPhoneRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.VerifyPhoneRegistration(context.Background()).VerifyPhoneRegistrationRequest(verifyPhoneRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.VerifyPhoneRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPhoneRegistration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthService.VerifyPhoneRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPhoneRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyPhoneRegistrationRequest** | [**VerifyPhoneRegistrationRequest**](VerifyPhoneRegistrationRequest.md) |  | 

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

