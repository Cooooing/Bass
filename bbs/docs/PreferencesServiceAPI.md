# \PreferencesServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreferencesServiceGetCurrent**](PreferencesServiceAPI.md#PreferencesServiceGetCurrent) | **Post** /v1/user/preference/get-current | 
[**PreferencesServiceUpdateCurrent**](PreferencesServiceAPI.md#PreferencesServiceUpdateCurrent) | **Post** /v1/user/preference/update-current | 



## PreferencesServiceGetCurrent

> GetCurrentPreferencesReply PreferencesServiceGetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.PreferencesServiceAPI.PreferencesServiceGetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesServiceAPI.PreferencesServiceGetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreferencesServiceGetCurrent`: GetCurrentPreferencesReply
	fmt.Fprintf(os.Stdout, "Response from `PreferencesServiceAPI.PreferencesServiceGetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreferencesServiceGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentPreferencesReply**](GetCurrentPreferencesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreferencesServiceUpdateCurrent

> UpdateCurrentPreferencesReply PreferencesServiceUpdateCurrent(ctx).UpdateCurrentPreferencesRequest(updateCurrentPreferencesRequest).Execute()





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
	updateCurrentPreferencesRequest := *openapiclient.NewUpdateCurrentPreferencesRequest() // UpdateCurrentPreferencesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreferencesServiceAPI.PreferencesServiceUpdateCurrent(context.Background()).UpdateCurrentPreferencesRequest(updateCurrentPreferencesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesServiceAPI.PreferencesServiceUpdateCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreferencesServiceUpdateCurrent`: UpdateCurrentPreferencesReply
	fmt.Fprintf(os.Stdout, "Response from `PreferencesServiceAPI.PreferencesServiceUpdateCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreferencesServiceUpdateCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCurrentPreferencesRequest** | [**UpdateCurrentPreferencesRequest**](UpdateCurrentPreferencesRequest.md) |  | 

### Return type

[**UpdateCurrentPreferencesReply**](UpdateCurrentPreferencesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

