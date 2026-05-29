# \PrivacySettingService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrent**](PrivacySettingService.md#GetCurrent) | **Post** /v1/user/privacy-setting/get-current | 
[**UpdateCurrent**](PrivacySettingService.md#UpdateCurrent) | **Post** /v1/user/privacy-setting/update-current | 



## GetCurrent

> GetCurrentPrivacySettingReply GetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.PrivacySettingService.GetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivacySettingService.GetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrent`: GetCurrentPrivacySettingReply
	fmt.Fprintf(os.Stdout, "Response from `PrivacySettingService.GetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentPrivacySettingReply**](GetCurrentPrivacySettingReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrent

> UpdateCurrentPrivacySettingReply UpdateCurrent(ctx).UpdateCurrentPrivacySettingRequest(updateCurrentPrivacySettingRequest).Execute()





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
	updateCurrentPrivacySettingRequest := *openapiclient.NewUpdateCurrentPrivacySettingRequest() // UpdateCurrentPrivacySettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivacySettingService.UpdateCurrent(context.Background()).UpdateCurrentPrivacySettingRequest(updateCurrentPrivacySettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivacySettingService.UpdateCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCurrent`: UpdateCurrentPrivacySettingReply
	fmt.Fprintf(os.Stdout, "Response from `PrivacySettingService.UpdateCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCurrentPrivacySettingRequest** | [**UpdateCurrentPrivacySettingRequest**](UpdateCurrentPrivacySettingRequest.md) |  | 

### Return type

[**UpdateCurrentPrivacySettingReply**](UpdateCurrentPrivacySettingReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

