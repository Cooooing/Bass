# \PostscriptService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](PostscriptService.md#Add) | **Post** /v1/content/postscript/add | 
[**List**](PostscriptService.md#List) | **Post** /v1/content/postscript/list | 



## Add

> AddPostscriptResp Add(ctx).AddPostscriptReq(addPostscriptReq).Execute()





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
	resp, r, err := apiClient.PostscriptService.Add(context.Background()).AddPostscriptReq(addPostscriptReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostscriptService.Add``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Add`: AddPostscriptResp
	fmt.Fprintf(os.Stdout, "Response from `PostscriptService.Add`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRequest struct via the builder pattern


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


## List

> ListPostscriptsResp List(ctx).ListPostscriptsReq(listPostscriptsReq).Execute()





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
	resp, r, err := apiClient.PostscriptService.List(context.Background()).ListPostscriptsReq(listPostscriptsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostscriptService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListPostscriptsResp
	fmt.Fprintf(os.Stdout, "Response from `PostscriptService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


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

