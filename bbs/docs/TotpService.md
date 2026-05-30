# \TotpService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BeginEnable**](TotpService.md#BeginEnable) | **Post** /v1/user/totp/begin-enable | 
[**ConfirmEnable**](TotpService.md#ConfirmEnable) | **Post** /v1/user/totp/confirm-enable | 
[**Disable**](TotpService.md#Disable) | **Post** /v1/user/totp/disable | 
[**GetCurrent**](TotpService.md#GetCurrent) | **Post** /v1/user/totp/get-current | 



## BeginEnable

> BeginEnableTotpReply BeginEnable(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.TotpService.BeginEnable(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TotpService.BeginEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BeginEnable`: BeginEnableTotpReply
	fmt.Fprintf(os.Stdout, "Response from `TotpService.BeginEnable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBeginEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**BeginEnableTotpReply**](BeginEnableTotpReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmEnable

> map[string]interface{} ConfirmEnable(ctx).ConfirmEnableTotpRequest(confirmEnableTotpRequest).Execute()





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
	confirmEnableTotpRequest := *openapiclient.NewConfirmEnableTotpRequest("Code_example") // ConfirmEnableTotpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TotpService.ConfirmEnable(context.Background()).ConfirmEnableTotpRequest(confirmEnableTotpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TotpService.ConfirmEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmEnable`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TotpService.ConfirmEnable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirmEnableTotpRequest** | [**ConfirmEnableTotpRequest**](ConfirmEnableTotpRequest.md) |  | 

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


## Disable

> map[string]interface{} Disable(ctx).DisableTotpRequest(disableTotpRequest).Execute()





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
	disableTotpRequest := *openapiclient.NewDisableTotpRequest("Code_example") // DisableTotpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TotpService.Disable(context.Background()).DisableTotpRequest(disableTotpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TotpService.Disable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Disable`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TotpService.Disable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableTotpRequest** | [**DisableTotpRequest**](DisableTotpRequest.md) |  | 

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


## GetCurrent

> GetCurrentTotpReply GetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.TotpService.GetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TotpService.GetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrent`: GetCurrentTotpReply
	fmt.Fprintf(os.Stdout, "Response from `TotpService.GetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentTotpReply**](GetCurrentTotpReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

