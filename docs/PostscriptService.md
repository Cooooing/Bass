# \PostscriptService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPostscript**](PostscriptService.md#AddPostscript) | **Post** /v1/content/postscript/add | 
[**ListPostscripts**](PostscriptService.md#ListPostscripts) | **Post** /v1/content/postscript/list | 



## AddPostscript

> AddPostscriptResp AddPostscript(ctx).AddPostscriptReq(addPostscriptReq).Execute()





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
	addPostscriptReq := *openapiclient.NewAddPostscriptReq("ArticleId_example", "Content_example") // AddPostscriptReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostscriptService.AddPostscript(context.Background()).AddPostscriptReq(addPostscriptReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostscriptService.AddPostscript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPostscript`: AddPostscriptResp
	fmt.Fprintf(os.Stdout, "Response from `PostscriptService.AddPostscript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPostscriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPostscriptReq** | [**AddPostscriptReq**](AddPostscriptReq.md) |  | 

### Return type

[**AddPostscriptResp**](AddPostscriptResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostscripts

> ListPostscriptsResp ListPostscripts(ctx).ListPostscriptsReq(listPostscriptsReq).Execute()





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
	listPostscriptsReq := *openapiclient.NewListPostscriptsReq("ArticleId_example") // ListPostscriptsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostscriptService.ListPostscripts(context.Background()).ListPostscriptsReq(listPostscriptsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostscriptService.ListPostscripts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPostscripts`: ListPostscriptsResp
	fmt.Fprintf(os.Stdout, "Response from `PostscriptService.ListPostscripts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPostscriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listPostscriptsReq** | [**ListPostscriptsReq**](ListPostscriptsReq.md) |  | 

### Return type

[**ListPostscriptsResp**](ListPostscriptsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

