# \ArticleSearchService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Search**](ArticleSearchService.md#Search) | **Post** /v1/content/article/search | 



## Search

> SearchArticlesResp Search(ctx).SearchArticlesReq(searchArticlesReq).Execute()



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
	searchArticlesReq := *openapiclient.NewSearchArticlesReq("Query_example") // SearchArticlesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArticleSearchService.Search(context.Background()).SearchArticlesReq(searchArticlesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArticleSearchService.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: SearchArticlesResp
	fmt.Fprintf(os.Stdout, "Response from `ArticleSearchService.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchArticlesReq** | [**SearchArticlesReq**](SearchArticlesReq.md) |  | 

### Return type

[**SearchArticlesResp**](SearchArticlesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

