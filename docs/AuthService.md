# \AuthService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginAccount**](AuthService.md#LoginAccount) | **Post** /v1/game-idle/auth/login | 
[**RegisterAccount**](AuthService.md#RegisterAccount) | **Post** /v1/game-idle/auth/register | 



## LoginAccount

> LoginAccountResp LoginAccount(ctx).LoginAccountReq(loginAccountReq).Execute()





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
	loginAccountReq := *openapiclient.NewLoginAccountReq("Email_example", "Password_example") // LoginAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.LoginAccount(context.Background()).LoginAccountReq(loginAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.LoginAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginAccount`: LoginAccountResp
	fmt.Fprintf(os.Stdout, "Response from `AuthService.LoginAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginAccountReq** | [**LoginAccountReq**](LoginAccountReq.md) |  | 

### Return type

[**LoginAccountResp**](LoginAccountResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterAccount

> map[string]interface{} RegisterAccount(ctx).RegisterAccountReq(registerAccountReq).Execute()





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
	registerAccountReq := *openapiclient.NewRegisterAccountReq("Email_example", "Password_example") // RegisterAccountReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthService.RegisterAccount(context.Background()).RegisterAccountReq(registerAccountReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthService.RegisterAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterAccount`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AuthService.RegisterAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerAccountReq** | [**RegisterAccountReq**](RegisterAccountReq.md) |  | 

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

