# \AssetService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteBadgeFontUpload**](AssetService.md#CompleteBadgeFontUpload) | **Post** /v1/asset/badge-font/complete | 
[**CompleteBadgeImageUpload**](AssetService.md#CompleteBadgeImageUpload) | **Post** /v1/asset/badge-image/complete | 
[**CompleteMarkdownImageUpload**](AssetService.md#CompleteMarkdownImageUpload) | **Post** /v1/asset/markdown-image/complete | 
[**PrepareBadgeFontUpload**](AssetService.md#PrepareBadgeFontUpload) | **Post** /v1/asset/badge-font/prepare | 
[**PrepareBadgeImageUpload**](AssetService.md#PrepareBadgeImageUpload) | **Post** /v1/asset/badge-image/prepare | 
[**PrepareMarkdownImageUpload**](AssetService.md#PrepareMarkdownImageUpload) | **Post** /v1/asset/markdown-image/prepare | 



## CompleteBadgeFontUpload

> CompleteBadgeFontUploadResp CompleteBadgeFontUpload(ctx).CompleteBadgeFontUploadReq(completeBadgeFontUploadReq).Execute()





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
	completeBadgeFontUploadReq := *openapiclient.NewCompleteBadgeFontUploadReq("Hash_example") // CompleteBadgeFontUploadReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetService.CompleteBadgeFontUpload(context.Background()).CompleteBadgeFontUploadReq(completeBadgeFontUploadReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetService.CompleteBadgeFontUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteBadgeFontUpload`: CompleteBadgeFontUploadResp
	fmt.Fprintf(os.Stdout, "Response from `AssetService.CompleteBadgeFontUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteBadgeFontUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completeBadgeFontUploadReq** | [**CompleteBadgeFontUploadReq**](CompleteBadgeFontUploadReq.md) |  | 

### Return type

[**CompleteBadgeFontUploadResp**](CompleteBadgeFontUploadResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteBadgeImageUpload

> CompleteBadgeImageUploadResp CompleteBadgeImageUpload(ctx).CompleteBadgeImageUploadReq(completeBadgeImageUploadReq).Execute()





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
	completeBadgeImageUploadReq := *openapiclient.NewCompleteBadgeImageUploadReq("Hash_example") // CompleteBadgeImageUploadReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetService.CompleteBadgeImageUpload(context.Background()).CompleteBadgeImageUploadReq(completeBadgeImageUploadReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetService.CompleteBadgeImageUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteBadgeImageUpload`: CompleteBadgeImageUploadResp
	fmt.Fprintf(os.Stdout, "Response from `AssetService.CompleteBadgeImageUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteBadgeImageUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completeBadgeImageUploadReq** | [**CompleteBadgeImageUploadReq**](CompleteBadgeImageUploadReq.md) |  | 

### Return type

[**CompleteBadgeImageUploadResp**](CompleteBadgeImageUploadResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteMarkdownImageUpload

> CompleteMarkdownImageUploadResp CompleteMarkdownImageUpload(ctx).CompleteMarkdownImageUploadReq(completeMarkdownImageUploadReq).Execute()





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
	completeMarkdownImageUploadReq := *openapiclient.NewCompleteMarkdownImageUploadReq("Hash_example") // CompleteMarkdownImageUploadReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetService.CompleteMarkdownImageUpload(context.Background()).CompleteMarkdownImageUploadReq(completeMarkdownImageUploadReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetService.CompleteMarkdownImageUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteMarkdownImageUpload`: CompleteMarkdownImageUploadResp
	fmt.Fprintf(os.Stdout, "Response from `AssetService.CompleteMarkdownImageUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteMarkdownImageUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completeMarkdownImageUploadReq** | [**CompleteMarkdownImageUploadReq**](CompleteMarkdownImageUploadReq.md) |  | 

### Return type

[**CompleteMarkdownImageUploadResp**](CompleteMarkdownImageUploadResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrepareBadgeFontUpload

> PrepareBadgeFontUploadResp PrepareBadgeFontUpload(ctx).PrepareBadgeFontUploadReq(prepareBadgeFontUploadReq).Execute()





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
	prepareBadgeFontUploadReq := *openapiclient.NewPrepareBadgeFontUploadReq("Hash_example", "MimeType_example", "Size_example") // PrepareBadgeFontUploadReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetService.PrepareBadgeFontUpload(context.Background()).PrepareBadgeFontUploadReq(prepareBadgeFontUploadReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetService.PrepareBadgeFontUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrepareBadgeFontUpload`: PrepareBadgeFontUploadResp
	fmt.Fprintf(os.Stdout, "Response from `AssetService.PrepareBadgeFontUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrepareBadgeFontUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prepareBadgeFontUploadReq** | [**PrepareBadgeFontUploadReq**](PrepareBadgeFontUploadReq.md) |  | 

### Return type

[**PrepareBadgeFontUploadResp**](PrepareBadgeFontUploadResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrepareBadgeImageUpload

> PrepareBadgeImageUploadResp PrepareBadgeImageUpload(ctx).PrepareBadgeImageUploadReq(prepareBadgeImageUploadReq).Execute()





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
	prepareBadgeImageUploadReq := *openapiclient.NewPrepareBadgeImageUploadReq("Hash_example", "MimeType_example", "Size_example") // PrepareBadgeImageUploadReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetService.PrepareBadgeImageUpload(context.Background()).PrepareBadgeImageUploadReq(prepareBadgeImageUploadReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetService.PrepareBadgeImageUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrepareBadgeImageUpload`: PrepareBadgeImageUploadResp
	fmt.Fprintf(os.Stdout, "Response from `AssetService.PrepareBadgeImageUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrepareBadgeImageUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prepareBadgeImageUploadReq** | [**PrepareBadgeImageUploadReq**](PrepareBadgeImageUploadReq.md) |  | 

### Return type

[**PrepareBadgeImageUploadResp**](PrepareBadgeImageUploadResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrepareMarkdownImageUpload

> PrepareMarkdownImageUploadResp PrepareMarkdownImageUpload(ctx).PrepareMarkdownImageUploadReq(prepareMarkdownImageUploadReq).Execute()





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
	prepareMarkdownImageUploadReq := *openapiclient.NewPrepareMarkdownImageUploadReq("Hash_example", "MimeType_example", "Size_example") // PrepareMarkdownImageUploadReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AssetService.PrepareMarkdownImageUpload(context.Background()).PrepareMarkdownImageUploadReq(prepareMarkdownImageUploadReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetService.PrepareMarkdownImageUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrepareMarkdownImageUpload`: PrepareMarkdownImageUploadResp
	fmt.Fprintf(os.Stdout, "Response from `AssetService.PrepareMarkdownImageUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrepareMarkdownImageUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prepareMarkdownImageUploadReq** | [**PrepareMarkdownImageUploadReq**](PrepareMarkdownImageUploadReq.md) |  | 

### Return type

[**PrepareMarkdownImageUploadResp**](PrepareMarkdownImageUploadResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

