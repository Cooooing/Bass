# \AccountService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Avatar**](AccountService.md#Avatar) | **Get** /v1/user/account/avatar | 
[**GetCurrent**](AccountService.md#GetCurrent) | **Post** /v1/user/account/get-current | 
[**GetProfile**](AccountService.md#GetProfile) | **Post** /v1/user/account/get-profile | 
[**UpdateProfile**](AccountService.md#UpdateProfile) | **Post** /v1/user/account/update-profile | 



## Avatar

> ImageReply Avatar(ctx).Name(name).Execute()





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
	name := "name_example" // string | 用于生成头像的账号名。 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.Avatar(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.Avatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Avatar`: ImageReply
	fmt.Fprintf(os.Stdout, "Response from `AccountService.Avatar`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 用于生成头像的账号名。 | 

### Return type

[**ImageReply**](ImageReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrent

> GetCurrentAccountReply GetCurrent(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.AccountService.GetCurrent(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.GetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrent`: GetCurrentAccountReply
	fmt.Fprintf(os.Stdout, "Response from `AccountService.GetCurrent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRequest struct via the builder pattern


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


## GetProfile

> GetProfileAccountReply GetProfile(ctx).GetProfileAccountRequest(getProfileAccountRequest).Execute()





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
	getProfileAccountRequest := *openapiclient.NewGetProfileAccountRequest("UserId_example") // GetProfileAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountService.GetProfile(context.Background()).GetProfileAccountRequest(getProfileAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: GetProfileAccountReply
	fmt.Fprintf(os.Stdout, "Response from `AccountService.GetProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


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


## UpdateProfile

> UpdateProfileAccountReply UpdateProfile(ctx).UpdateProfileAccountRequest(updateProfileAccountRequest).Execute()





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
	resp, r, err := apiClient.AccountService.UpdateProfile(context.Background()).UpdateProfileAccountRequest(updateProfileAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountService.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProfile`: UpdateProfileAccountReply
	fmt.Fprintf(os.Stdout, "Response from `AccountService.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


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

