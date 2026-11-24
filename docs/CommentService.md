# \CommentService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComment**](CommentService.md#CreateComment) | **Post** /v1/content/comment/create | 
[**LikeComment**](CommentService.md#LikeComment) | **Post** /v1/content/comment/like | 
[**ListCommentReplies**](CommentService.md#ListCommentReplies) | **Post** /v1/content/comment/list-replies | 
[**ListCommentThreads**](CommentService.md#ListCommentThreads) | **Post** /v1/content/comment/list-threads | 
[**ListCommentTimeline**](CommentService.md#ListCommentTimeline) | **Post** /v1/content/comment/list-timeline | 
[**ListComments**](CommentService.md#ListComments) | **Post** /v1/content/comment/list | 
[**ThankComment**](CommentService.md#ThankComment) | **Post** /v1/content/comment/thank | 



## CreateComment

> CreateCommentResp CreateComment(ctx).CreateCommentReq(createCommentReq).Execute()





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
	resp, r, err := apiClient.CommentService.CreateComment(context.Background()).CreateCommentReq(createCommentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.CreateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComment`: CreateCommentResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.CreateComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentRequest struct via the builder pattern


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


## LikeComment

> LikeCommentResp LikeComment(ctx).LikeCommentReq(likeCommentReq).Execute()





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
	resp, r, err := apiClient.CommentService.LikeComment(context.Background()).LikeCommentReq(likeCommentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.LikeComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LikeComment`: LikeCommentResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.LikeComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLikeCommentRequest struct via the builder pattern


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


## ListCommentReplies

> ListCommentRepliesResp ListCommentReplies(ctx).ListCommentRepliesReq(listCommentRepliesReq).Execute()





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
	resp, r, err := apiClient.CommentService.ListCommentReplies(context.Background()).ListCommentRepliesReq(listCommentRepliesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListCommentReplies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCommentReplies`: ListCommentRepliesResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListCommentReplies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCommentRepliesRequest struct via the builder pattern


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


## ListCommentThreads

> ListCommentThreadsResp ListCommentThreads(ctx).ListCommentThreadsReq(listCommentThreadsReq).Execute()





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
	resp, r, err := apiClient.CommentService.ListCommentThreads(context.Background()).ListCommentThreadsReq(listCommentThreadsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListCommentThreads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCommentThreads`: ListCommentThreadsResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListCommentThreads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCommentThreadsRequest struct via the builder pattern


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


## ListCommentTimeline

> ListCommentTimelineResp ListCommentTimeline(ctx).ListCommentTimelineReq(listCommentTimelineReq).Execute()





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
	resp, r, err := apiClient.CommentService.ListCommentTimeline(context.Background()).ListCommentTimelineReq(listCommentTimelineReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListCommentTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCommentTimeline`: ListCommentTimelineResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListCommentTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCommentTimelineRequest struct via the builder pattern


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


## ListComments

> ListCommentsResp ListComments(ctx).ListCommentsReq(listCommentsReq).Execute()





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
	resp, r, err := apiClient.CommentService.ListComments(context.Background()).ListCommentsReq(listCommentsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListComments`: ListCommentsResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCommentsRequest struct via the builder pattern


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


## ThankComment

> ThankCommentResp ThankComment(ctx).ThankCommentReq(thankCommentReq).Execute()





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
	thankCommentReq := *openapiclient.NewThankCommentReq("Id_example") // ThankCommentReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.ThankComment(context.Background()).ThankCommentReq(thankCommentReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ThankComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThankComment`: ThankCommentResp
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ThankComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThankCommentRequest struct via the builder pattern


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

