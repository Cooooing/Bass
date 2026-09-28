# \OtpService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BeginEnableTotp**](OtpService.md#BeginEnableTotp) | **Post** /v1/user/otp/totp/begin-enable | 
[**ConfirmEnableTotp**](OtpService.md#ConfirmEnableTotp) | **Post** /v1/user/otp/totp/confirm-enable | 
[**DisableTotp**](OtpService.md#DisableTotp) | **Post** /v1/user/otp/totp/disable | 
[**GetCurrentTotp**](OtpService.md#GetCurrentTotp) | **Post** /v1/user/otp/totp/get-current | 
[**SendEmailOtp**](OtpService.md#SendEmailOtp) | **Post** /v1/user/otp/email/send | 
[**SendPhoneOtp**](OtpService.md#SendPhoneOtp) | **Post** /v1/user/otp/phone/send | 



## BeginEnableTotp

> BeginEnableTotpResp BeginEnableTotp(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.OtpService.BeginEnableTotp(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtpService.BeginEnableTotp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BeginEnableTotp`: BeginEnableTotpResp
	fmt.Fprintf(os.Stdout, "Response from `OtpService.BeginEnableTotp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBeginEnableTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**BeginEnableTotpResp**](BeginEnableTotpResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmEnableTotp

> map[string]interface{} ConfirmEnableTotp(ctx).ConfirmEnableTotpReq(confirmEnableTotpReq).Execute()





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
	confirmEnableTotpReq := *openapiclient.NewConfirmEnableTotpReq("Code_example") // ConfirmEnableTotpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtpService.ConfirmEnableTotp(context.Background()).ConfirmEnableTotpReq(confirmEnableTotpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtpService.ConfirmEnableTotp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfirmEnableTotp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OtpService.ConfirmEnableTotp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmEnableTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirmEnableTotpReq** | [**ConfirmEnableTotpReq**](ConfirmEnableTotpReq.md) |  | 

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


## DisableTotp

> map[string]interface{} DisableTotp(ctx).DisableTotpReq(disableTotpReq).Execute()





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
	disableTotpReq := *openapiclient.NewDisableTotpReq("Code_example") // DisableTotpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtpService.DisableTotp(context.Background()).DisableTotpReq(disableTotpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtpService.DisableTotp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisableTotp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OtpService.DisableTotp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disableTotpReq** | [**DisableTotpReq**](DisableTotpReq.md) |  | 

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


## GetCurrentTotp

> GetCurrentTotpResp GetCurrentTotp(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.OtpService.GetCurrentTotp(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtpService.GetCurrentTotp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentTotp`: GetCurrentTotpResp
	fmt.Fprintf(os.Stdout, "Response from `OtpService.GetCurrentTotp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentTotpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentTotpResp**](GetCurrentTotpResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendEmailOtp

> SendEmailOtpResp SendEmailOtp(ctx).SendEmailOtpReq(sendEmailOtpReq).Execute()





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
	sendEmailOtpReq := *openapiclient.NewSendEmailOtpReq("Email_example") // SendEmailOtpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtpService.SendEmailOtp(context.Background()).SendEmailOtpReq(sendEmailOtpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtpService.SendEmailOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmailOtp`: SendEmailOtpResp
	fmt.Fprintf(os.Stdout, "Response from `OtpService.SendEmailOtp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendEmailOtpReq** | [**SendEmailOtpReq**](SendEmailOtpReq.md) |  | 

### Return type

[**SendEmailOtpResp**](SendEmailOtpResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPhoneOtp

> SendPhoneOtpResp SendPhoneOtp(ctx).SendPhoneOtpReq(sendPhoneOtpReq).Execute()





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
	sendPhoneOtpReq := *openapiclient.NewSendPhoneOtpReq("Phone_example") // SendPhoneOtpReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OtpService.SendPhoneOtp(context.Background()).SendPhoneOtpReq(sendPhoneOtpReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OtpService.SendPhoneOtp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPhoneOtp`: SendPhoneOtpResp
	fmt.Fprintf(os.Stdout, "Response from `OtpService.SendPhoneOtp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPhoneOtpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendPhoneOtpReq** | [**SendPhoneOtpReq**](SendPhoneOtpReq.md) |  | 

### Return type

[**SendPhoneOtpResp**](SendPhoneOtpResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

