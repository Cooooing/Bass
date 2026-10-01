# \MoonbreezeService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](MoonbreezeService.md#Create) | **Post** /v1/content/moonbreeze/create | 
[**PageMember**](MoonbreezeService.md#PageMember) | **Post** /v1/content/moonbreeze/page-member | 
[**PagePublic**](MoonbreezeService.md#PagePublic) | **Post** /v1/content/moonbreeze/page-public | 
[**PageWatching**](MoonbreezeService.md#PageWatching) | **Post** /v1/content/moonbreeze/page-watching | 



## Create

> CreateMoonbreezeResp Create(ctx).CreateMoonbreezeReq(createMoonbreezeReq).Execute()



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
	resp, r, err := apiClient.MoonbreezeService.Create(context.Background()).CreateMoonbreezeReq(createMoonbreezeReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoonbreezeService.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateMoonbreezeResp
	fmt.Fprintf(os.Stdout, "Response from `MoonbreezeService.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


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


## PageMember

> PageMoonbreezesResp PageMember(ctx).PageMemberMoonbreezesReq(pageMemberMoonbreezesReq).Execute()



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
	resp, r, err := apiClient.MoonbreezeService.PageMember(context.Background()).PageMemberMoonbreezesReq(pageMemberMoonbreezesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoonbreezeService.PageMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageMember`: PageMoonbreezesResp
	fmt.Fprintf(os.Stdout, "Response from `MoonbreezeService.PageMember`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageMemberMoonbreezesReq** | [**PageMemberMoonbreezesReq**](PageMemberMoonbreezesReq.md) |  | 

### Return type

[**PageMoonbreezesResp**](PageMoonbreezesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PagePublic

> PageMoonbreezesResp PagePublic(ctx).PagePublicMoonbreezesReq(pagePublicMoonbreezesReq).Execute()



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
	resp, r, err := apiClient.MoonbreezeService.PagePublic(context.Background()).PagePublicMoonbreezesReq(pagePublicMoonbreezesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoonbreezeService.PagePublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PagePublic`: PageMoonbreezesResp
	fmt.Fprintf(os.Stdout, "Response from `MoonbreezeService.PagePublic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPagePublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagePublicMoonbreezesReq** | [**PagePublicMoonbreezesReq**](PagePublicMoonbreezesReq.md) |  | 

### Return type

[**PageMoonbreezesResp**](PageMoonbreezesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PageWatching

> PageMoonbreezesResp PageWatching(ctx).PageWatchingMoonbreezesReq(pageWatchingMoonbreezesReq).Execute()



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
	resp, r, err := apiClient.MoonbreezeService.PageWatching(context.Background()).PageWatchingMoonbreezesReq(pageWatchingMoonbreezesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoonbreezeService.PageWatching``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PageWatching`: PageMoonbreezesResp
	fmt.Fprintf(os.Stdout, "Response from `MoonbreezeService.PageWatching`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPageWatchingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageWatchingMoonbreezesReq** | [**PageWatchingMoonbreezesReq**](PageWatchingMoonbreezesReq.md) |  | 

### Return type

[**PageMoonbreezesResp**](PageMoonbreezesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

