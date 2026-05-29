# \CommentServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommentServiceCreate**](CommentServiceAPI.md#CommentServiceCreate) | **Post** /v1/content/comment/create | 
[**CommentServiceLike**](CommentServiceAPI.md#CommentServiceLike) | **Post** /v1/content/comment/like | 
[**CommentServiceList**](CommentServiceAPI.md#CommentServiceList) | **Post** /v1/content/comment/list | 
[**CommentServiceThank**](CommentServiceAPI.md#CommentServiceThank) | **Post** /v1/content/comment/thank | 



## CommentServiceCreate

> CreateCommentReply CommentServiceCreate(ctx).CreateCommentRequest(createCommentRequest).Execute()



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
	createCommentRequest := *openapiclient.NewCreateCommentRequest() // CreateCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentServiceAPI.CommentServiceCreate(context.Background()).CreateCommentRequest(createCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentServiceAPI.CommentServiceCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentServiceCreate`: CreateCommentReply
	fmt.Fprintf(os.Stdout, "Response from `CommentServiceAPI.CommentServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommentServiceCreateRequest struct via the builder pattern


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


## CommentServiceLike

> map[string]interface{} CommentServiceLike(ctx).LikeCommentRequest(likeCommentRequest).Execute()



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
	likeCommentRequest := *openapiclient.NewLikeCommentRequest() // LikeCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentServiceAPI.CommentServiceLike(context.Background()).LikeCommentRequest(likeCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentServiceAPI.CommentServiceLike``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentServiceLike`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommentServiceAPI.CommentServiceLike`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommentServiceLikeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **likeCommentRequest** | [**LikeCommentRequest**](LikeCommentRequest.md) |  | 

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


## CommentServiceList

> ListCommentsReply CommentServiceList(ctx).ListCommentsRequest(listCommentsRequest).Execute()



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
	resp, r, err := apiClient.CommentServiceAPI.CommentServiceList(context.Background()).ListCommentsRequest(listCommentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentServiceAPI.CommentServiceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentServiceList`: ListCommentsReply
	fmt.Fprintf(os.Stdout, "Response from `CommentServiceAPI.CommentServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommentServiceListRequest struct via the builder pattern


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


## CommentServiceThank

> map[string]interface{} CommentServiceThank(ctx).ThankCommentRequest(thankCommentRequest).Execute()



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
	thankCommentRequest := *openapiclient.NewThankCommentRequest() // ThankCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentServiceAPI.CommentServiceThank(context.Background()).ThankCommentRequest(thankCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentServiceAPI.CommentServiceThank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CommentServiceThank`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommentServiceAPI.CommentServiceThank`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommentServiceThankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thankCommentRequest** | [**ThankCommentRequest**](ThankCommentRequest.md) |  | 

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

