# \TagService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindArticleTags**](TagService.md#BindArticleTags) | **Post** /v1/content/tag/bind-article | 
[**CreateTag**](TagService.md#CreateTag) | **Post** /v1/content/tag/create | 
[**ListArticleTags**](TagService.md#ListArticleTags) | **Post** /v1/content/tag/list-article-tags | 
[**ListTags**](TagService.md#ListTags) | **Post** /v1/content/tag/list | 
[**UnbindArticleTags**](TagService.md#UnbindArticleTags) | **Post** /v1/content/tag/unbind-article | 
[**UpdateTag**](TagService.md#UpdateTag) | **Post** /v1/content/tag/update | 



## BindArticleTags

> map[string]interface{} BindArticleTags(ctx).BindArticleTagsReq(bindArticleTagsReq).Execute()





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
	resp, r, err := apiClient.TagService.BindArticleTags(context.Background()).BindArticleTagsReq(bindArticleTagsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.BindArticleTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BindArticleTags`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TagService.BindArticleTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBindArticleTagsRequest struct via the builder pattern


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


## CreateTag

> CreateTagResp CreateTag(ctx).CreateTagReq(createTagReq).Execute()





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
	createTagReq := *openapiclient.NewCreateTagReq(*openapiclient.NewReqTag("Code_example", "Name_example")) // CreateTagReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagService.CreateTag(context.Background()).CreateTagReq(createTagReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: CreateTagResp
	fmt.Fprintf(os.Stdout, "Response from `TagService.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


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


## ListTags

> ListTagsResp ListTags(ctx).ListTagsReq(listTagsReq).Execute()





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
	resp, r, err := apiClient.TagService.ListTags(context.Background()).ListTagsReq(listTagsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.ListTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTags`: ListTagsResp
	fmt.Fprintf(os.Stdout, "Response from `TagService.ListTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTagsRequest struct via the builder pattern


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


## UnbindArticleTags

> map[string]interface{} UnbindArticleTags(ctx).UnbindArticleTagsReq(unbindArticleTagsReq).Execute()





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
	resp, r, err := apiClient.TagService.UnbindArticleTags(context.Background()).UnbindArticleTagsReq(unbindArticleTagsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.UnbindArticleTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnbindArticleTags`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TagService.UnbindArticleTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnbindArticleTagsRequest struct via the builder pattern


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


## UpdateTag

> UpdateTagResp UpdateTag(ctx).UpdateTagReq(updateTagReq).Execute()





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
	updateTagReq := *openapiclient.NewUpdateTagReq("TagId_example", *openapiclient.NewReqTag("Code_example", "Name_example")) // UpdateTagReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagService.UpdateTag(context.Background()).UpdateTagReq(updateTagReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagService.UpdateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTag`: UpdateTagResp
	fmt.Fprintf(os.Stdout, "Response from `TagService.UpdateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


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

