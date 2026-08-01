# \CommentService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CommentService.md#Create) | **Post** /v1/content/comment/create | 
[**Like**](CommentService.md#Like) | **Post** /v1/content/comment/like | 
[**List**](CommentService.md#List) | **Post** /v1/content/comment/list | 
[**ListReplies**](CommentService.md#ListReplies) | **Post** /v1/content/comment/list-replies | 
[**ListThreads**](CommentService.md#ListThreads) | **Post** /v1/content/comment/list-threads | 
[**ListTimeline**](CommentService.md#ListTimeline) | **Post** /v1/content/comment/list-timeline | 
[**Thank**](CommentService.md#Thank) | **Post** /v1/content/comment/thank | 



## Create

> CreateCommentResp Create(ctx).CreateCommentReq(createCommentReq).Execute()





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
	createCommentReq := *openapiclient.NewCreateCommentReq("ArticleId_example", "Content_example") // CreateCommentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.Create(context.Background()).CreateCommentReq(createCommentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateCommentResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCommentReq** | [**CreateCommentReq**](CreateCommentReq.md) |  | 

### Return type

[**CreateCommentResp**](CreateCommentResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Like

> LikeCommentResp Like(ctx).LikeCommentReq(likeCommentReq).Execute()





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
	likeCommentReq := *openapiclient.NewLikeCommentReq("Id_example", false) // LikeCommentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.Like(context.Background()).LikeCommentReq(likeCommentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.Like``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Like`: LikeCommentResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.Like`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLikeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **likeCommentReq** | [**LikeCommentReq**](LikeCommentReq.md) |  | 

### Return type

[**LikeCommentResp**](LikeCommentResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListCommentsResp List(ctx).ListCommentsReq(listCommentsReq).Execute()





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
	listCommentsReq := *openapiclient.NewListCommentsReq() // ListCommentsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.List(context.Background()).ListCommentsReq(listCommentsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListCommentsResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCommentsReq** | [**ListCommentsReq**](ListCommentsReq.md) |  | 

### Return type

[**ListCommentsResp**](ListCommentsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReplies

> ListCommentRepliesResp ListReplies(ctx).ListCommentRepliesReq(listCommentRepliesReq).Execute()





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
	listCommentRepliesReq := *openapiclient.NewListCommentRepliesReq("ArticleId_example", "ParentId_example") // ListCommentRepliesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.ListReplies(context.Background()).ListCommentRepliesReq(listCommentRepliesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListReplies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReplies`: ListCommentRepliesResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListReplies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCommentRepliesReq** | [**ListCommentRepliesReq**](ListCommentRepliesReq.md) |  | 

### Return type

[**ListCommentRepliesResp**](ListCommentRepliesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListThreads

> ListCommentThreadsResp ListThreads(ctx).ListCommentThreadsReq(listCommentThreadsReq).Execute()





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
	listCommentThreadsReq := *openapiclient.NewListCommentThreadsReq("ArticleId_example") // ListCommentThreadsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.ListThreads(context.Background()).ListCommentThreadsReq(listCommentThreadsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListThreads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListThreads`: ListCommentThreadsResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListThreads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCommentThreadsReq** | [**ListCommentThreadsReq**](ListCommentThreadsReq.md) |  | 

### Return type

[**ListCommentThreadsResp**](ListCommentThreadsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTimeline

> ListCommentTimelineResp ListTimeline(ctx).ListCommentTimelineReq(listCommentTimelineReq).Execute()





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
	listCommentTimelineReq := *openapiclient.NewListCommentTimelineReq("ArticleId_example") // ListCommentTimelineReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.ListTimeline(context.Background()).ListCommentTimelineReq(listCommentTimelineReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTimeline`: ListCommentTimelineResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCommentTimelineReq** | [**ListCommentTimelineReq**](ListCommentTimelineReq.md) |  | 

### Return type

[**ListCommentTimelineResp**](ListCommentTimelineResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Thank

> ThankCommentResp Thank(ctx).ThankCommentReq(thankCommentReq).Execute()





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
	thankCommentReq := *openapiclient.NewThankCommentReq("Id_example", false) // ThankCommentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.Thank(context.Background()).ThankCommentReq(thankCommentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.Thank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Thank`: ThankCommentResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.Thank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thankCommentReq** | [**ThankCommentReq**](ThankCommentReq.md) |  | 

### Return type

[**ThankCommentResp**](ThankCommentResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

