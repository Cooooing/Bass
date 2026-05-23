# \AccountServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountServiceBatchGetProfile**](AccountServiceAPI.md#AccountServiceBatchGetProfile) | **Post** /v1/user/account/batch-get-profile | 
[**AccountServiceGetCurrent**](AccountServiceAPI.md#AccountServiceGetCurrent) | **Post** /v1/user/account/get-current | 
[**AccountServiceGetProfile**](AccountServiceAPI.md#AccountServiceGetProfile) | **Post** /v1/user/account/get-profile | 
[**AccountServiceUpdateProfile**](AccountServiceAPI.md#AccountServiceUpdateProfile) | **Post** /v1/user/account/update-profile | 



## AccountServiceBatchGetProfile

> CommonApiAppBbsV1UserBatchGetProfileAccountReply AccountServiceBatchGetProfile(ctx).CommonApiAppBbsV1UserBatchGetProfileAccountRequest(commonApiAppBbsV1UserBatchGetProfileAccountRequest).Execute()





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
	commonApiAppBbsV1UserBatchGetProfileAccountRequest := *openapiclient.NewCommonApiAppBbsV1UserBatchGetProfileAccountRequest() // CommonApiAppBbsV1UserBatchGetProfileAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountServiceAPI.AccountServiceBatchGetProfile(context.Background()).CommonApiAppBbsV1UserBatchGetProfileAccountRequest(commonApiAppBbsV1UserBatchGetProfileAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceAPI.AccountServiceBatchGetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountServiceBatchGetProfile`: CommonApiAppBbsV1UserBatchGetProfileAccountReply
	fmt.Fprintf(os.Stdout, "Response from `AccountServiceAPI.AccountServiceBatchGetProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceBatchGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserBatchGetProfileAccountRequest** | [**CommonApiAppBbsV1UserBatchGetProfileAccountRequest**](CommonApiAppBbsV1UserBatchGetProfileAccountRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserBatchGetProfileAccountReply**](CommonApiAppBbsV1UserBatchGetProfileAccountReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceGetCurrent

> CommonApiAppBbsV1UserGetCurrentAccountReply AccountServiceGetCurrent(ctx).Body(body).Execute()





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
	// response from `AccountServiceGetCurrent`: CommonApiAppBbsV1UserGetCurrentAccountReply
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

[**CommonApiAppBbsV1UserGetCurrentAccountReply**](CommonApiAppBbsV1UserGetCurrentAccountReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceGetProfile

> CommonApiAppBbsV1UserGetProfileAccountReply AccountServiceGetProfile(ctx).CommonApiAppBbsV1UserGetProfileAccountRequest(commonApiAppBbsV1UserGetProfileAccountRequest).Execute()





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
	commonApiAppBbsV1UserGetProfileAccountRequest := *openapiclient.NewCommonApiAppBbsV1UserGetProfileAccountRequest() // CommonApiAppBbsV1UserGetProfileAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountServiceAPI.AccountServiceGetProfile(context.Background()).CommonApiAppBbsV1UserGetProfileAccountRequest(commonApiAppBbsV1UserGetProfileAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceAPI.AccountServiceGetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountServiceGetProfile`: CommonApiAppBbsV1UserGetProfileAccountReply
	fmt.Fprintf(os.Stdout, "Response from `AccountServiceAPI.AccountServiceGetProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserGetProfileAccountRequest** | [**CommonApiAppBbsV1UserGetProfileAccountRequest**](CommonApiAppBbsV1UserGetProfileAccountRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserGetProfileAccountReply**](CommonApiAppBbsV1UserGetProfileAccountReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountServiceUpdateProfile

> CommonApiAppBbsV1UserUpdateProfileAccountReply AccountServiceUpdateProfile(ctx).CommonApiAppBbsV1UserUpdateProfileAccountRequest(commonApiAppBbsV1UserUpdateProfileAccountRequest).Execute()





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
	commonApiAppBbsV1UserUpdateProfileAccountRequest := *openapiclient.NewCommonApiAppBbsV1UserUpdateProfileAccountRequest() // CommonApiAppBbsV1UserUpdateProfileAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountServiceAPI.AccountServiceUpdateProfile(context.Background()).CommonApiAppBbsV1UserUpdateProfileAccountRequest(commonApiAppBbsV1UserUpdateProfileAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountServiceAPI.AccountServiceUpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountServiceUpdateProfile`: CommonApiAppBbsV1UserUpdateProfileAccountReply
	fmt.Fprintf(os.Stdout, "Response from `AccountServiceAPI.AccountServiceUpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountServiceUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserUpdateProfileAccountRequest** | [**CommonApiAppBbsV1UserUpdateProfileAccountRequest**](CommonApiAppBbsV1UserUpdateProfileAccountRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserUpdateProfileAccountReply**](CommonApiAppBbsV1UserUpdateProfileAccountReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

