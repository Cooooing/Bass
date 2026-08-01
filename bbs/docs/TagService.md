# \TagService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindArticle**](TagService.md#BindArticle) | **Post** /v1/content/tag/bind-article | 
[**Create**](TagService.md#Create) | **Post** /v1/content/tag/create | 
[**List**](TagService.md#List) | **Post** /v1/content/tag/list | 
[**ListArticleTags**](TagService.md#ListArticleTags) | **Post** /v1/content/tag/list-article-tags | 
[**UnbindArticle**](TagService.md#UnbindArticle) | **Post** /v1/content/tag/unbind-article | 
[**Update**](TagService.md#Update) | **Post** /v1/content/tag/update | 



## BindArticle

> map[string]interface{} BindArticle(ctx).BindArticleTagsReq(bindArticleTagsReq).Execute()





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
	bindArticleTagsReq := *openapiclient.NewBindArticleTagsReq("ArticleId_example", []string{"TagIds_example"}) // BindArticleTagsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagService.BindArticle(context.Background()).BindArticleTagsReq(bindArticleTagsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.BindArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindArticle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TagService.BindArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBindArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindArticleTagsReq** | [**BindArticleTagsReq**](BindArticleTagsReq.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> CreateTagResp Create(ctx).CreateTagReq(createTagReq).Execute()





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
	createTagReq := *openapiclient.NewCreateTagReq(*openapiclient.NewReqTag("Code_example", "Name_example", "DomainId_example")) // CreateTagReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagService.Create(context.Background()).CreateTagReq(createTagReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateTagResp
	fmt.Fprintf(os.Stdout, "Response from `TagService.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagReq** | [**CreateTagReq**](CreateTagReq.md) |  | 

### Return type

[**CreateTagResp**](CreateTagResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListTagsResp List(ctx).ListTagsReq(listTagsReq).Execute()





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
	listTagsReq := *openapiclient.NewListTagsReq() // ListTagsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagService.List(context.Background()).ListTagsReq(listTagsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListTagsResp
	fmt.Fprintf(os.Stdout, "Response from `TagService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listTagsReq** | [**ListTagsReq**](ListTagsReq.md) |  | 

### Return type

[**ListTagsResp**](ListTagsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArticleTags

> ListArticleTagsResp ListArticleTags(ctx).ListArticleTagsReq(listArticleTagsReq).Execute()





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
	listArticleTagsReq := *openapiclient.NewListArticleTagsReq("ArticleId_example") // ListArticleTagsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagService.ListArticleTags(context.Background()).ListArticleTagsReq(listArticleTagsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.ListArticleTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListArticleTags`: ListArticleTagsResp
	fmt.Fprintf(os.Stdout, "Response from `TagService.ListArticleTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListArticleTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listArticleTagsReq** | [**ListArticleTagsReq**](ListArticleTagsReq.md) |  | 

### Return type

[**ListArticleTagsResp**](ListArticleTagsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnbindArticle

> map[string]interface{} UnbindArticle(ctx).UnbindArticleTagsReq(unbindArticleTagsReq).Execute()





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
	unbindArticleTagsReq := *openapiclient.NewUnbindArticleTagsReq("ArticleId_example", []string{"TagIds_example"}) // UnbindArticleTagsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagService.UnbindArticle(context.Background()).UnbindArticleTagsReq(unbindArticleTagsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.UnbindArticle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindArticle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TagService.UnbindArticle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnbindArticleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unbindArticleTagsReq** | [**UnbindArticleTagsReq**](UnbindArticleTagsReq.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateTagResp Update(ctx).UpdateTagReq(updateTagReq).Execute()





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
	updateTagReq := *openapiclient.NewUpdateTagReq("TagId_example", *openapiclient.NewReqTag("Code_example", "Name_example", "DomainId_example")) // UpdateTagReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagService.Update(context.Background()).UpdateTagReq(updateTagReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateTagResp
	fmt.Fprintf(os.Stdout, "Response from `TagService.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTagReq** | [**UpdateTagReq**](UpdateTagReq.md) |  | 

### Return type

[**UpdateTagResp**](UpdateTagResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

