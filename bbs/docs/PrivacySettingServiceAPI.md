# \PrivacySettingServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivacySettingServiceGetCurrent**](PrivacySettingServiceAPI.md#PrivacySettingServiceGetCurrent) | **Post** /v1/user/privacy-setting/get-current | 
[**PrivacySettingServiceUpdateCurrent**](PrivacySettingServiceAPI.md#PrivacySettingServiceUpdateCurrent) | **Post** /v1/user/privacy-setting/update-current | 



## PrivacySettingServiceGetCurrent

> GetCurrentPrivacySettingReply PrivacySettingServiceGetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.PrivacySettingServiceAPI.PrivacySettingServiceGetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivacySettingServiceAPI.PrivacySettingServiceGetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivacySettingServiceGetCurrent`: GetCurrentPrivacySettingReply
	fmt.Fprintf(os.Stdout, "Response from `PrivacySettingServiceAPI.PrivacySettingServiceGetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySettingServiceGetCurrentRequest struct via the builder pattern


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


## PrivacySettingServiceUpdateCurrent

> UpdateCurrentPrivacySettingReply PrivacySettingServiceUpdateCurrent(ctx).UpdateCurrentPrivacySettingRequest(updateCurrentPrivacySettingRequest).Execute()





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
	resp, r, err := apiClient.PrivacySettingServiceAPI.PrivacySettingServiceUpdateCurrent(context.Background()).UpdateCurrentPrivacySettingRequest(updateCurrentPrivacySettingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivacySettingServiceAPI.PrivacySettingServiceUpdateCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrivacySettingServiceUpdateCurrent`: UpdateCurrentPrivacySettingReply
	fmt.Fprintf(os.Stdout, "Response from `PrivacySettingServiceAPI.PrivacySettingServiceUpdateCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivacySettingServiceUpdateCurrentRequest struct via the builder pattern


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

