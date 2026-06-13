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

> CreateCommentReply Create(ctx).CreateCommentRequest(createCommentRequest).Execute()





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
	createCommentRequest := *openapiclient.NewCreateCommentRequest("ArticleId_example", "Content_example") // CreateCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.Create(context.Background()).CreateCommentRequest(createCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateCommentReply
	fmt.Fprintf(os.Stdout, "Response from `CommentService.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCommentRequest** | [**CreateCommentRequest**](CreateCommentRequest.md) |  | 

### Return type

[**CreateCommentReply**](CreateCommentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Like

> LikeCommentReply Like(ctx).LikeCommentRequest(likeCommentRequest).Execute()





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
	likeCommentRequest := *openapiclient.NewLikeCommentRequest("Id_example", false) // LikeCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.Like(context.Background()).LikeCommentRequest(likeCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.Like``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Like`: LikeCommentReply
	fmt.Fprintf(os.Stdout, "Response from `CommentService.Like`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLikeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **likeCommentRequest** | [**LikeCommentRequest**](LikeCommentRequest.md) |  | 

### Return type

[**LikeCommentReply**](LikeCommentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListCommentsReply List(ctx).ListCommentsRequest(listCommentsRequest).Execute()





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
	listCommentsRequest := *openapiclient.NewListCommentsRequest() // ListCommentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.List(context.Background()).ListCommentsRequest(listCommentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListCommentsReply
	fmt.Fprintf(os.Stdout, "Response from `CommentService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCommentsRequest** | [**ListCommentsRequest**](ListCommentsRequest.md) |  | 

### Return type

[**ListCommentsReply**](ListCommentsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReplies

> ListCommentRepliesReply ListReplies(ctx).ListCommentRepliesRequest(listCommentRepliesRequest).Execute()





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
	listCommentRepliesRequest := *openapiclient.NewListCommentRepliesRequest("ArticleId_example", "ParentId_example") // ListCommentRepliesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.ListReplies(context.Background()).ListCommentRepliesRequest(listCommentRepliesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListReplies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReplies`: ListCommentRepliesReply
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListReplies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRepliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCommentRepliesRequest** | [**ListCommentRepliesRequest**](ListCommentRepliesRequest.md) |  | 

### Return type

[**ListCommentRepliesReply**](ListCommentRepliesReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListThreads

> ListCommentThreadsReply ListThreads(ctx).ListCommentThreadsRequest(listCommentThreadsRequest).Execute()





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
	listCommentThreadsRequest := *openapiclient.NewListCommentThreadsRequest("ArticleId_example") // ListCommentThreadsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.ListThreads(context.Background()).ListCommentThreadsRequest(listCommentThreadsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListThreads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListThreads`: ListCommentThreadsReply
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListThreads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListThreadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCommentThreadsRequest** | [**ListCommentThreadsRequest**](ListCommentThreadsRequest.md) |  | 

### Return type

[**ListCommentThreadsReply**](ListCommentThreadsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTimeline

> ListCommentTimelineReply ListTimeline(ctx).ListCommentTimelineRequest(listCommentTimelineRequest).Execute()





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
	listCommentTimelineRequest := *openapiclient.NewListCommentTimelineRequest("ArticleId_example") // ListCommentTimelineRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.ListTimeline(context.Background()).ListCommentTimelineRequest(listCommentTimelineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.ListTimeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTimeline`: ListCommentTimelineReply
	fmt.Fprintf(os.Stdout, "Response from `CommentService.ListTimeline`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTimelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCommentTimelineRequest** | [**ListCommentTimelineRequest**](ListCommentTimelineRequest.md) |  | 

### Return type

[**ListCommentTimelineReply**](ListCommentTimelineReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Thank

> ThankCommentReply Thank(ctx).ThankCommentRequest(thankCommentRequest).Execute()





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
	thankCommentRequest := *openapiclient.NewThankCommentRequest("Id_example", false) // ThankCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentService.Thank(context.Background()).ThankCommentRequest(thankCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentService.Thank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Thank`: ThankCommentReply
	fmt.Fprintf(os.Stdout, "Response from `CommentService.Thank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thankCommentRequest** | [**ThankCommentRequest**](ThankCommentRequest.md) |  | 

### Return type

[**ThankCommentReply**](ThankCommentReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

