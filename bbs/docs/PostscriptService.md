# \PostscriptService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](PostscriptService.md#Add) | **Post** /v1/content/postscript/add | 



## Add

> AddPostscriptReply Add(ctx).AddPostscriptRequest(addPostscriptRequest).Execute()





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
	addPostscriptRequest := *openapiclient.NewAddPostscriptRequest() // AddPostscriptRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostscriptService.Add(context.Background()).AddPostscriptRequest(addPostscriptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostscriptService.Add``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Add`: AddPostscriptReply
	fmt.Fprintf(os.Stdout, "Response from `PostscriptService.Add`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addPostscriptRequest** | [**AddPostscriptRequest**](AddPostscriptRequest.md) |  | 

### Return type

[**AddPostscriptReply**](AddPostscriptReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

