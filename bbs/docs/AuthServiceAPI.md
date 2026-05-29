# \AuthServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthServiceLoginByPassword**](AuthServiceAPI.md#AuthServiceLoginByPassword) | **Post** /v1/user/auth/login-by-password | 
[**AuthServiceLogout**](AuthServiceAPI.md#AuthServiceLogout) | **Post** /v1/user/auth/logout | 
[**AuthServiceStartEmailRegistration**](AuthServiceAPI.md#AuthServiceStartEmailRegistration) | **Post** /v1/user/auth/start-email-registration | 
[**AuthServiceStartPhoneRegistration**](AuthServiceAPI.md#AuthServiceStartPhoneRegistration) | **Post** /v1/user/auth/start-phone-registration | 
[**AuthServiceVerifyEmailRegistration**](AuthServiceAPI.md#AuthServiceVerifyEmailRegistration) | **Post** /v1/user/auth/verify-email-registration | 
[**AuthServiceVerifyPhoneRegistration**](AuthServiceAPI.md#AuthServiceVerifyPhoneRegistration) | **Post** /v1/user/auth/verify-phone-registration | 



## AuthServiceLoginByPassword

> LoginByPasswordReply AuthServiceLoginByPassword(ctx).LoginByPasswordRequest(loginByPasswordRequest).Execute()





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
	loginByPasswordRequest := *openapiclient.NewLoginByPasswordRequest() // LoginByPasswordRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceLoginByPassword(context.Background()).LoginByPasswordRequest(loginByPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceLoginByPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceLoginByPassword`: LoginByPasswordReply
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceLoginByPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceLoginByPasswordRequest struct via the builder pattern


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


## AuthServiceStartEmailRegistration

> StartEmailRegistrationReply AuthServiceStartEmailRegistration(ctx).StartEmailRegistrationRequest(startEmailRegistrationRequest).Execute()





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
	startEmailRegistrationRequest := *openapiclient.NewStartEmailRegistrationRequest() // StartEmailRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceStartEmailRegistration(context.Background()).StartEmailRegistrationRequest(startEmailRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceStartEmailRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceStartEmailRegistration`: StartEmailRegistrationReply
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceStartEmailRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceStartEmailRegistrationRequest struct via the builder pattern


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


## AuthServiceStartPhoneRegistration

> StartPhoneRegistrationReply AuthServiceStartPhoneRegistration(ctx).StartPhoneRegistrationRequest(startPhoneRegistrationRequest).Execute()





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
	startPhoneRegistrationRequest := *openapiclient.NewStartPhoneRegistrationRequest() // StartPhoneRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceStartPhoneRegistration(context.Background()).StartPhoneRegistrationRequest(startPhoneRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceStartPhoneRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceStartPhoneRegistration`: StartPhoneRegistrationReply
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceStartPhoneRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceStartPhoneRegistrationRequest struct via the builder pattern


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


## AuthServiceVerifyEmailRegistration

> map[string]interface{} AuthServiceVerifyEmailRegistration(ctx).VerifyEmailRegistrationRequest(verifyEmailRegistrationRequest).Execute()





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
	verifyEmailRegistrationRequest := *openapiclient.NewVerifyEmailRegistrationRequest() // VerifyEmailRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceVerifyEmailRegistration(context.Background()).VerifyEmailRegistrationRequest(verifyEmailRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceVerifyEmailRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceVerifyEmailRegistration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceVerifyEmailRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceVerifyEmailRegistrationRequest struct via the builder pattern


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


## AuthServiceVerifyPhoneRegistration

> map[string]interface{} AuthServiceVerifyPhoneRegistration(ctx).VerifyPhoneRegistrationRequest(verifyPhoneRegistrationRequest).Execute()





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
	verifyPhoneRegistrationRequest := *openapiclient.NewVerifyPhoneRegistrationRequest() // VerifyPhoneRegistrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthServiceAPI.AuthServiceVerifyPhoneRegistration(context.Background()).VerifyPhoneRegistrationRequest(verifyPhoneRegistrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthServiceAPI.AuthServiceVerifyPhoneRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthServiceVerifyPhoneRegistration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthServiceAPI.AuthServiceVerifyPhoneRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthServiceVerifyPhoneRegistrationRequest struct via the builder pattern


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

