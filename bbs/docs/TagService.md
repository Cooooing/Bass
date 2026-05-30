# \TagService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](TagService.md#List) | **Post** /v1/content/tag/list | 



## List

> ListTagsReply List(ctx).ListTagsRequest(listTagsRequest).Execute()





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
	listTagsRequest := *openapiclient.NewListTagsRequest() // ListTagsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagService.List(context.Background()).ListTagsRequest(listTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListTagsReply
	fmt.Fprintf(os.Stdout, "Response from `TagService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listTagsRequest** | [**ListTagsRequest**](ListTagsRequest.md) |  | 

### Return type

[**ListTagsReply**](ListTagsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

