# \AuthService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelAccount**](AuthService.md#CancelAccount) | **Post** /v1/user/auth/cancel-account | 
[**Login**](AuthService.md#Login) | **Post** /v1/user/auth/login | 
[**Logout**](AuthService.md#Logout) | **Post** /v1/user/auth/logout | 
[**RefreshToken**](AuthService.md#RefreshToken) | **Post** /v1/user/auth/refresh-token | 
[**Register**](AuthService.md#Register) | **Post** /v1/user/auth/register | 



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


## Register

> map[string]interface{} Register(ctx).RegisterReq(registerReq).Execute()





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
	registerReq := *openapiclient.NewRegisterReq("Type_example", "Name_example", "Password_example") // RegisterReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.Register(context.Background()).RegisterReq(registerReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.Register``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Register`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthService.Register`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerReq** | [**RegisterReq**](RegisterReq.md) |  | 

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

