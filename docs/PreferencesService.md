# \PreferencesService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentPreferences**](PreferencesService.md#GetCurrentPreferences) | **Post** /v1/user/preference/get-current | 
[**UpdateCurrentPreferences**](PreferencesService.md#UpdateCurrentPreferences) | **Post** /v1/user/preference/update-current | 



## GetCurrentPreferences

> GetCurrentPreferencesResp GetCurrentPreferences(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.PreferencesService.GetCurrentPreferences(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesService.GetCurrentPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentPreferences`: GetCurrentPreferencesResp
	fmt.Fprintf(os.Stdout, "Response from `PreferencesService.GetCurrentPreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentPreferencesRequest struct via the builder pattern


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


## UpdateCurrentPreferences

> UpdateCurrentPreferencesResp UpdateCurrentPreferences(ctx).UpdateCurrentPreferencesReq(updateCurrentPreferencesReq).Execute()





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
	resp, r, err := apiClient.PreferencesService.UpdateCurrentPreferences(context.Background()).UpdateCurrentPreferencesReq(updateCurrentPreferencesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesService.UpdateCurrentPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCurrentPreferences`: UpdateCurrentPreferencesResp
	fmt.Fprintf(os.Stdout, "Response from `PreferencesService.UpdateCurrentPreferences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentPreferencesRequest struct via the builder pattern


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

