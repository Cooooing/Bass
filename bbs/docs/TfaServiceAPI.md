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

> CommonApiAppBbsV1UserBeginEnableTfaReply TfaServiceBeginEnable(ctx).Body(body).Execute()





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
	// response from `TfaServiceBeginEnable`: CommonApiAppBbsV1UserBeginEnableTfaReply
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

[**CommonApiAppBbsV1UserBeginEnableTfaReply**](CommonApiAppBbsV1UserBeginEnableTfaReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TfaServiceConfirmEnable

> map[string]interface{} TfaServiceConfirmEnable(ctx).CommonApiAppBbsV1UserConfirmEnableTfaRequest(commonApiAppBbsV1UserConfirmEnableTfaRequest).Execute()





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
	commonApiAppBbsV1UserConfirmEnableTfaRequest := *openapiclient.NewCommonApiAppBbsV1UserConfirmEnableTfaRequest() // CommonApiAppBbsV1UserConfirmEnableTfaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaServiceAPI.TfaServiceConfirmEnable(context.Background()).CommonApiAppBbsV1UserConfirmEnableTfaRequest(commonApiAppBbsV1UserConfirmEnableTfaRequest).Execute()
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
 **commonApiAppBbsV1UserConfirmEnableTfaRequest** | [**CommonApiAppBbsV1UserConfirmEnableTfaRequest**](CommonApiAppBbsV1UserConfirmEnableTfaRequest.md) |  | 

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

> map[string]interface{} TfaServiceDisable(ctx).CommonApiAppBbsV1UserDisableTfaRequest(commonApiAppBbsV1UserDisableTfaRequest).Execute()





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
	commonApiAppBbsV1UserDisableTfaRequest := *openapiclient.NewCommonApiAppBbsV1UserDisableTfaRequest() // CommonApiAppBbsV1UserDisableTfaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaServiceAPI.TfaServiceDisable(context.Background()).CommonApiAppBbsV1UserDisableTfaRequest(commonApiAppBbsV1UserDisableTfaRequest).Execute()
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
 **commonApiAppBbsV1UserDisableTfaRequest** | [**CommonApiAppBbsV1UserDisableTfaRequest**](CommonApiAppBbsV1UserDisableTfaRequest.md) |  | 

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

> CommonApiAppBbsV1UserGetCurrentTfaReply TfaServiceGetCurrent(ctx).Body(body).Execute()





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
	// response from `TfaServiceGetCurrent`: CommonApiAppBbsV1UserGetCurrentTfaReply
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

[**CommonApiAppBbsV1UserGetCurrentTfaReply**](CommonApiAppBbsV1UserGetCurrentTfaReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TfaServiceValidate

> CommonApiAppBbsV1UserValidateTfaReply TfaServiceValidate(ctx).CommonApiAppBbsV1UserValidateTfaRequest(commonApiAppBbsV1UserValidateTfaRequest).Execute()





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
	commonApiAppBbsV1UserValidateTfaRequest := *openapiclient.NewCommonApiAppBbsV1UserValidateTfaRequest() // CommonApiAppBbsV1UserValidateTfaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TfaServiceAPI.TfaServiceValidate(context.Background()).CommonApiAppBbsV1UserValidateTfaRequest(commonApiAppBbsV1UserValidateTfaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TfaServiceAPI.TfaServiceValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TfaServiceValidate`: CommonApiAppBbsV1UserValidateTfaReply
	fmt.Fprintf(os.Stdout, "Response from `TfaServiceAPI.TfaServiceValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTfaServiceValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserValidateTfaRequest** | [**CommonApiAppBbsV1UserValidateTfaRequest**](CommonApiAppBbsV1UserValidateTfaRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserValidateTfaReply**](CommonApiAppBbsV1UserValidateTfaReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

