# \PreferencesServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreferencesServiceGetCurrent**](PreferencesServiceAPI.md#PreferencesServiceGetCurrent) | **Post** /v1/user/preference/get-current | 
[**PreferencesServiceUpdate**](PreferencesServiceAPI.md#PreferencesServiceUpdate) | **Post** /v1/user/preference/update-current | 



## PreferencesServiceGetCurrent

> CommonApiAppBbsV1UserGetCurrentPreferencesReply PreferencesServiceGetCurrent(ctx).Body(body).Execute()





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
	// response from `PreferencesServiceGetCurrent`: CommonApiAppBbsV1UserGetCurrentPreferencesReply
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

[**CommonApiAppBbsV1UserGetCurrentPreferencesReply**](CommonApiAppBbsV1UserGetCurrentPreferencesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreferencesServiceUpdate

> CommonApiAppBbsV1UserUpdatePreferencesReply PreferencesServiceUpdate(ctx).CommonApiAppBbsV1UserUpdatePreferencesRequest(commonApiAppBbsV1UserUpdatePreferencesRequest).Execute()





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
	commonApiAppBbsV1UserUpdatePreferencesRequest := *openapiclient.NewCommonApiAppBbsV1UserUpdatePreferencesRequest() // CommonApiAppBbsV1UserUpdatePreferencesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PreferencesServiceAPI.PreferencesServiceUpdate(context.Background()).CommonApiAppBbsV1UserUpdatePreferencesRequest(commonApiAppBbsV1UserUpdatePreferencesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PreferencesServiceAPI.PreferencesServiceUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreferencesServiceUpdate`: CommonApiAppBbsV1UserUpdatePreferencesReply
	fmt.Fprintf(os.Stdout, "Response from `PreferencesServiceAPI.PreferencesServiceUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreferencesServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserUpdatePreferencesRequest** | [**CommonApiAppBbsV1UserUpdatePreferencesRequest**](CommonApiAppBbsV1UserUpdatePreferencesRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserUpdatePreferencesReply**](CommonApiAppBbsV1UserUpdatePreferencesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

