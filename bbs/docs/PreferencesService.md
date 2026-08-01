# \PreferencesService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrent**](PreferencesService.md#GetCurrent) | **Post** /v1/user/preference/get-current | 
[**UpdateCurrent**](PreferencesService.md#UpdateCurrent) | **Post** /v1/user/preference/update-current | 



## GetCurrent

> GetCurrentPreferencesResp GetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.PreferencesService.GetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesService.GetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrent`: GetCurrentPreferencesResp
	fmt.Fprintf(os.Stdout, "Response from `PreferencesService.GetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentPreferencesResp**](GetCurrentPreferencesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrent

> UpdateCurrentPreferencesResp UpdateCurrent(ctx).UpdateCurrentPreferencesReq(updateCurrentPreferencesReq).Execute()





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
	updateCurrentPreferencesReq := *openapiclient.NewUpdateCurrentPreferencesReq() // UpdateCurrentPreferencesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreferencesService.UpdateCurrent(context.Background()).UpdateCurrentPreferencesReq(updateCurrentPreferencesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesService.UpdateCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCurrent`: UpdateCurrentPreferencesResp
	fmt.Fprintf(os.Stdout, "Response from `PreferencesService.UpdateCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCurrentPreferencesReq** | [**UpdateCurrentPreferencesReq**](UpdateCurrentPreferencesReq.md) |  | 

### Return type

[**UpdateCurrentPreferencesResp**](UpdateCurrentPreferencesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

