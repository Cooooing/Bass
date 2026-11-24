# \MoonbreezeService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMoonbreeze**](MoonbreezeService.md#CreateMoonbreeze) | **Post** /v1/content/moonbreeze/create | 
[**PageMemberMoonbreezes**](MoonbreezeService.md#PageMemberMoonbreezes) | **Post** /v1/content/moonbreeze/page-member | 
[**PagePublicMoonbreezes**](MoonbreezeService.md#PagePublicMoonbreezes) | **Post** /v1/content/moonbreeze/page-public | 
[**PageWatchingMoonbreezes**](MoonbreezeService.md#PageWatchingMoonbreezes) | **Post** /v1/content/moonbreeze/page-watching | 



## CreateMoonbreeze

> CreateMoonbreezeResp CreateMoonbreeze(ctx).CreateMoonbreezeReq(createMoonbreezeReq).Execute()





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
	createMoonbreezeReq := *openapiclient.NewCreateMoonbreezeReq() // CreateMoonbreezeReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoonbreezeService.CreateMoonbreeze(context.Background()).CreateMoonbreezeReq(createMoonbreezeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoonbreezeService.CreateMoonbreeze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMoonbreeze`: CreateMoonbreezeResp
	fmt.Fprintf(os.Stdout, "Response from `MoonbreezeService.CreateMoonbreeze`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMoonbreezeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMoonbreezeReq** | [**CreateMoonbreezeReq**](CreateMoonbreezeReq.md) |  | 

### Return type

[**CreateMoonbreezeResp**](CreateMoonbreezeResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageMemberMoonbreezes

> PageMemberMoonbreezesResp PageMemberMoonbreezes(ctx).PageMemberMoonbreezesReq(pageMemberMoonbreezesReq).Execute()





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
	pageMemberMoonbreezesReq := *openapiclient.NewPageMemberMoonbreezesReq() // PageMemberMoonbreezesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoonbreezeService.PageMemberMoonbreezes(context.Background()).PageMemberMoonbreezesReq(pageMemberMoonbreezesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoonbreezeService.PageMemberMoonbreezes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageMemberMoonbreezes`: PageMemberMoonbreezesResp
	fmt.Fprintf(os.Stdout, "Response from `MoonbreezeService.PageMemberMoonbreezes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageMemberMoonbreezesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageMemberMoonbreezesReq** | [**PageMemberMoonbreezesReq**](PageMemberMoonbreezesReq.md) |  | 

### Return type

[**PageMemberMoonbreezesResp**](PageMemberMoonbreezesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PagePublicMoonbreezes

> PagePublicMoonbreezesResp PagePublicMoonbreezes(ctx).PagePublicMoonbreezesReq(pagePublicMoonbreezesReq).Execute()





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
	pagePublicMoonbreezesReq := *openapiclient.NewPagePublicMoonbreezesReq() // PagePublicMoonbreezesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoonbreezeService.PagePublicMoonbreezes(context.Background()).PagePublicMoonbreezesReq(pagePublicMoonbreezesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoonbreezeService.PagePublicMoonbreezes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PagePublicMoonbreezes`: PagePublicMoonbreezesResp
	fmt.Fprintf(os.Stdout, "Response from `MoonbreezeService.PagePublicMoonbreezes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPagePublicMoonbreezesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagePublicMoonbreezesReq** | [**PagePublicMoonbreezesReq**](PagePublicMoonbreezesReq.md) |  | 

### Return type

[**PagePublicMoonbreezesResp**](PagePublicMoonbreezesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageWatchingMoonbreezes

> PageWatchingMoonbreezesResp PageWatchingMoonbreezes(ctx).PageWatchingMoonbreezesReq(pageWatchingMoonbreezesReq).Execute()





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
	pageWatchingMoonbreezesReq := *openapiclient.NewPageWatchingMoonbreezesReq() // PageWatchingMoonbreezesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoonbreezeService.PageWatchingMoonbreezes(context.Background()).PageWatchingMoonbreezesReq(pageWatchingMoonbreezesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoonbreezeService.PageWatchingMoonbreezes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageWatchingMoonbreezes`: PageWatchingMoonbreezesResp
	fmt.Fprintf(os.Stdout, "Response from `MoonbreezeService.PageWatchingMoonbreezes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageWatchingMoonbreezesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageWatchingMoonbreezesReq** | [**PageWatchingMoonbreezesReq**](PageWatchingMoonbreezesReq.md) |  | 

### Return type

[**PageWatchingMoonbreezesResp**](PageWatchingMoonbreezesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

