# \PrivacySettingService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentPrivacySetting**](PrivacySettingService.md#GetCurrentPrivacySetting) | **Post** /v1/user/privacy-setting/get-current | 
[**UpdateCurrentPrivacySetting**](PrivacySettingService.md#UpdateCurrentPrivacySetting) | **Post** /v1/user/privacy-setting/update-current | 



## GetCurrentPrivacySetting

> GetCurrentPrivacySettingResp GetCurrentPrivacySetting(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.PrivacySettingService.GetCurrentPrivacySetting(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivacySettingService.GetCurrentPrivacySetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentPrivacySetting`: GetCurrentPrivacySettingResp
	fmt.Fprintf(os.Stdout, "Response from `PrivacySettingService.GetCurrentPrivacySetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentPrivacySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentPrivacySettingResp**](GetCurrentPrivacySettingResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrentPrivacySetting

> UpdateCurrentPrivacySettingResp UpdateCurrentPrivacySetting(ctx).UpdateCurrentPrivacySettingReq(updateCurrentPrivacySettingReq).Execute()





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
	updateCurrentPrivacySettingReq := *openapiclient.NewUpdateCurrentPrivacySettingReq() // UpdateCurrentPrivacySettingReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivacySettingService.UpdateCurrentPrivacySetting(context.Background()).UpdateCurrentPrivacySettingReq(updateCurrentPrivacySettingReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivacySettingService.UpdateCurrentPrivacySetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCurrentPrivacySetting`: UpdateCurrentPrivacySettingResp
	fmt.Fprintf(os.Stdout, "Response from `PrivacySettingService.UpdateCurrentPrivacySetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentPrivacySettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCurrentPrivacySettingReq** | [**UpdateCurrentPrivacySettingReq**](UpdateCurrentPrivacySettingReq.md) |  | 

### Return type

[**UpdateCurrentPrivacySettingResp**](UpdateCurrentPrivacySettingResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

