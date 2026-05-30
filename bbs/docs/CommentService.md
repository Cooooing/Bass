# \CommentService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CommentService.md#Create) | **Post** /v1/content/comment/create | 
[**Like**](CommentService.md#Like) | **Post** /v1/content/comment/like | 
[**List**](CommentService.md#List) | **Post** /v1/content/comment/list | 
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

> map[string]interface{} Like(ctx).LikeCommentRequest(likeCommentRequest).Execute()





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
	// response from `Like`: map[string]interface{}
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

**map[string]interface{}**

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


## Thank

> map[string]interface{} Thank(ctx).ThankCommentRequest(thankCommentRequest).Execute()





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
	// response from `Thank`: map[string]interface{}
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

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

