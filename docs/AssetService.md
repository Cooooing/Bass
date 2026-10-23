# \AssetService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteMarkdownImageUpload**](AssetService.md#CompleteMarkdownImageUpload) | **Post** /v1/asset/markdown-image/complete | 
[**PrepareMarkdownImageUpload**](AssetService.md#PrepareMarkdownImageUpload) | **Post** /v1/asset/markdown-image/prepare | 



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

