# \AccountServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountServiceGetCurrent**](AccountServiceAPI.md#AccountServiceGetCurrent) | **Post** /v1/user/account/get-current | 
[**AccountServiceGetProfile**](AccountServiceAPI.md#AccountServiceGetProfile) | **Post** /v1/user/account/get-profile | 
[**AccountServiceUpdateProfile**](AccountServiceAPI.md#AccountServiceUpdateProfile) | **Post** /v1/user/account/update-profile | 



## AccountServiceGetCurrent

> GetCurrentAccountReply AccountServiceGetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.AccountServiceAPI.AccountServiceGetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceAPI.AccountServiceGetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountServiceGetCurrent`: GetCurrentAccountReply
	fmt.Fprintf(os.Stdout, "Response from `AccountServiceAPI.AccountServiceGetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceGetCurrentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetCurrentAccountReply**](GetCurrentAccountReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceGetProfile

> GetProfileAccountReply AccountServiceGetProfile(ctx).GetProfileAccountRequest(getProfileAccountRequest).Execute()





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
	getProfileAccountRequest := *openapiclient.NewGetProfileAccountRequest() // GetProfileAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountServiceAPI.AccountServiceGetProfile(context.Background()).GetProfileAccountRequest(getProfileAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceAPI.AccountServiceGetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountServiceGetProfile`: GetProfileAccountReply
	fmt.Fprintf(os.Stdout, "Response from `AccountServiceAPI.AccountServiceGetProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getProfileAccountRequest** | [**GetProfileAccountRequest**](GetProfileAccountRequest.md) |  | 

### Return type

[**GetProfileAccountReply**](GetProfileAccountReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceUpdateProfile

> UpdateProfileAccountReply AccountServiceUpdateProfile(ctx).UpdateProfileAccountRequest(updateProfileAccountRequest).Execute()





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
	updateProfileAccountRequest := *openapiclient.NewUpdateProfileAccountRequest() // UpdateProfileAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountServiceAPI.AccountServiceUpdateProfile(context.Background()).UpdateProfileAccountRequest(updateProfileAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceAPI.AccountServiceUpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountServiceUpdateProfile`: UpdateProfileAccountReply
	fmt.Fprintf(os.Stdout, "Response from `AccountServiceAPI.AccountServiceUpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProfileAccountRequest** | [**UpdateProfileAccountRequest**](UpdateProfileAccountRequest.md) |  | 

### Return type

[**UpdateProfileAccountReply**](UpdateProfileAccountReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

