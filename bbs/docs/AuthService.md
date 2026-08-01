# \AuthService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelAccount**](AuthService.md#CancelAccount) | **Post** /v1/user/auth/cancel-account | 
[**Login**](AuthService.md#Login) | **Post** /v1/user/auth/login | 
[**Logout**](AuthService.md#Logout) | **Post** /v1/user/auth/logout | 
[**RefreshToken**](AuthService.md#RefreshToken) | **Post** /v1/user/auth/refresh-token | 
[**StartEmailRegistration**](AuthService.md#StartEmailRegistration) | **Post** /v1/user/auth/start-email-registration | 
[**StartPhoneRegistration**](AuthService.md#StartPhoneRegistration) | **Post** /v1/user/auth/start-phone-registration | 
[**VerifyEmailRegistration**](AuthService.md#VerifyEmailRegistration) | **Post** /v1/user/auth/verify-email-registration | 
[**VerifyPhoneRegistration**](AuthService.md#VerifyPhoneRegistration) | **Post** /v1/user/auth/verify-phone-registration | 



## CancelAccount

> map[string]interface{} CancelAccount(ctx).CancelAccountReq(cancelAccountReq).Execute()





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
	cancelAccountReq := *openapiclient.NewCancelAccountReq("Password_example") // CancelAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.CancelAccount(context.Background()).CancelAccountReq(cancelAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.CancelAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthService.CancelAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cancelAccountReq** | [**CancelAccountReq**](CancelAccountReq.md) |  | 

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


## Login

> LoginResp Login(ctx).LoginReq(loginReq).Execute()





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
	loginReq := *openapiclient.NewLoginReq("Type_example") // LoginReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.Login(context.Background()).LoginReq(loginReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.Login``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Login`: LoginResp
	fmt.Fprintf(os.Stdout, "Response from `AuthService.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginReq** | [**LoginReq**](LoginReq.md) |  | 

### Return type

[**LoginResp**](LoginResp.md)

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


## RefreshToken

> RefreshTokenResp RefreshToken(ctx).RefreshTokenReq(refreshTokenReq).Execute()





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
	refreshTokenReq := *openapiclient.NewRefreshTokenReq("RefreshToken_example") // RefreshTokenReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.RefreshToken(context.Background()).RefreshTokenReq(refreshTokenReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.RefreshToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshToken`: RefreshTokenResp
	fmt.Fprintf(os.Stdout, "Response from `AuthService.RefreshToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshTokenReq** | [**RefreshTokenReq**](RefreshTokenReq.md) |  | 

### Return type

[**RefreshTokenResp**](RefreshTokenResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartEmailRegistration

> StartEmailRegistrationResp StartEmailRegistration(ctx).StartEmailRegistrationReq(startEmailRegistrationReq).Execute()





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
	startEmailRegistrationReq := *openapiclient.NewStartEmailRegistrationReq("Email_example", "Password_example", "Name_example") // StartEmailRegistrationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.StartEmailRegistration(context.Background()).StartEmailRegistrationReq(startEmailRegistrationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.StartEmailRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartEmailRegistration`: StartEmailRegistrationResp
	fmt.Fprintf(os.Stdout, "Response from `AuthService.StartEmailRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartEmailRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startEmailRegistrationReq** | [**StartEmailRegistrationReq**](StartEmailRegistrationReq.md) |  | 

### Return type

[**StartEmailRegistrationResp**](StartEmailRegistrationResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartPhoneRegistration

> StartPhoneRegistrationResp StartPhoneRegistration(ctx).StartPhoneRegistrationReq(startPhoneRegistrationReq).Execute()





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
	startPhoneRegistrationReq := *openapiclient.NewStartPhoneRegistrationReq("Phone_example", "Password_example", "Name_example") // StartPhoneRegistrationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.StartPhoneRegistration(context.Background()).StartPhoneRegistrationReq(startPhoneRegistrationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.StartPhoneRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartPhoneRegistration`: StartPhoneRegistrationResp
	fmt.Fprintf(os.Stdout, "Response from `AuthService.StartPhoneRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartPhoneRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startPhoneRegistrationReq** | [**StartPhoneRegistrationReq**](StartPhoneRegistrationReq.md) |  | 

### Return type

[**StartPhoneRegistrationResp**](StartPhoneRegistrationResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEmailRegistration

> map[string]interface{} VerifyEmailRegistration(ctx).VerifyEmailRegistrationReq(verifyEmailRegistrationReq).Execute()





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
	verifyEmailRegistrationReq := *openapiclient.NewVerifyEmailRegistrationReq("Email_example", "Code_example") // VerifyEmailRegistrationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.VerifyEmailRegistration(context.Background()).VerifyEmailRegistrationReq(verifyEmailRegistrationReq).Execute()
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
 **verifyEmailRegistrationReq** | [**VerifyEmailRegistrationReq**](VerifyEmailRegistrationReq.md) |  | 

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

> map[string]interface{} VerifyPhoneRegistration(ctx).VerifyPhoneRegistrationReq(verifyPhoneRegistrationReq).Execute()





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
	verifyPhoneRegistrationReq := *openapiclient.NewVerifyPhoneRegistrationReq("Phone_example", "Code_example") // VerifyPhoneRegistrationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.VerifyPhoneRegistration(context.Background()).VerifyPhoneRegistrationReq(verifyPhoneRegistrationReq).Execute()
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
 **verifyPhoneRegistrationReq** | [**VerifyPhoneRegistrationReq**](VerifyPhoneRegistrationReq.md) |  | 

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

