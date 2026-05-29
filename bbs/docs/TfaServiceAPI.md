# \TfaServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TfaServiceBeginEnable**](TfaServiceAPI.md#TfaServiceBeginEnable) | **Post** /v1/user/tfa/begin-enable | 
[**TfaServiceConfirmEnable**](TfaServiceAPI.md#TfaServiceConfirmEnable) | **Post** /v1/user/tfa/confirm-enable | 
[**TfaServiceDisable**](TfaServiceAPI.md#TfaServiceDisable) | **Post** /v1/user/tfa/disable | 
[**TfaServiceGetCurrent**](TfaServiceAPI.md#TfaServiceGetCurrent) | **Post** /v1/user/tfa/get-current | 
[**TfaServiceValidate**](TfaServiceAPI.md#TfaServiceValidate) | **Post** /v1/user/tfa/validate | 



## TfaServiceBeginEnable

> BeginEnableTfaReply TfaServiceBeginEnable(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.TfaServiceAPI.TfaServiceBeginEnable(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaServiceAPI.TfaServiceBeginEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TfaServiceBeginEnable`: BeginEnableTfaReply
	fmt.Fprintf(os.Stdout, "Response from `TfaServiceAPI.TfaServiceBeginEnable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTfaServiceBeginEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**BeginEnableTfaReply**](BeginEnableTfaReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TfaServiceConfirmEnable

> map[string]interface{} TfaServiceConfirmEnable(ctx).ConfirmEnableTfaRequest(confirmEnableTfaRequest).Execute()





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
	confirmEnableTfaRequest := *openapiclient.NewConfirmEnableTfaRequest() // ConfirmEnableTfaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaServiceAPI.TfaServiceConfirmEnable(context.Background()).ConfirmEnableTfaRequest(confirmEnableTfaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaServiceAPI.TfaServiceConfirmEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TfaServiceConfirmEnable`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TfaServiceAPI.TfaServiceConfirmEnable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTfaServiceConfirmEnableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirmEnableTfaRequest** | [**ConfirmEnableTfaRequest**](ConfirmEnableTfaRequest.md) |  | 

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


## TfaServiceDisable

> map[string]interface{} TfaServiceDisable(ctx).DisableTfaRequest(disableTfaRequest).Execute()





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
	disableTfaRequest := *openapiclient.NewDisableTfaRequest() // DisableTfaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaServiceAPI.TfaServiceDisable(context.Background()).DisableTfaRequest(disableTfaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaServiceAPI.TfaServiceDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TfaServiceDisable`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TfaServiceAPI.TfaServiceDisable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTfaServiceDisableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableTfaRequest** | [**DisableTfaRequest**](DisableTfaRequest.md) |  | 

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


## TfaServiceGetCurrent

> GetCurrentTfaReply TfaServiceGetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.TfaServiceAPI.TfaServiceGetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaServiceAPI.TfaServiceGetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TfaServiceGetCurrent`: GetCurrentTfaReply
	fmt.Fprintf(os.Stdout, "Response from `TfaServiceAPI.TfaServiceGetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTfaServiceGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentTfaReply**](GetCurrentTfaReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TfaServiceValidate

> ValidateTfaReply TfaServiceValidate(ctx).ValidateTfaRequest(validateTfaRequest).Execute()





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
	validateTfaRequest := *openapiclient.NewValidateTfaRequest() // ValidateTfaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaServiceAPI.TfaServiceValidate(context.Background()).ValidateTfaRequest(validateTfaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaServiceAPI.TfaServiceValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TfaServiceValidate`: ValidateTfaReply
	fmt.Fprintf(os.Stdout, "Response from `TfaServiceAPI.TfaServiceValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTfaServiceValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateTfaRequest** | [**ValidateTfaRequest**](ValidateTfaRequest.md) |  | 

### Return type

[**ValidateTfaReply**](ValidateTfaReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

