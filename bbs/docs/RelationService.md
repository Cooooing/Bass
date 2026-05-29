# \RelationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Block**](RelationService.md#Block) | **Post** /v1/user/relation/block | 
[**Follow**](RelationService.md#Follow) | **Post** /v1/user/relation/follow | 
[**GetStatus**](RelationService.md#GetStatus) | **Post** /v1/user/relation/get-status | 
[**ListBlocked**](RelationService.md#ListBlocked) | **Post** /v1/user/relation/list-blocked | 
[**ListFollowers**](RelationService.md#ListFollowers) | **Post** /v1/user/relation/list-followers | 
[**ListFollowing**](RelationService.md#ListFollowing) | **Post** /v1/user/relation/list-following | 
[**Unblock**](RelationService.md#Unblock) | **Post** /v1/user/relation/unblock | 
[**Unfollow**](RelationService.md#Unfollow) | **Post** /v1/user/relation/unfollow | 



## Block

> map[string]interface{} Block(ctx).BlockRelationRequest(blockRelationRequest).Execute()





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
	blockRelationRequest := *openapiclient.NewBlockRelationRequest() // BlockRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.Block(context.Background()).BlockRelationRequest(blockRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.Block``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Block`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationService.Block`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockRelationRequest** | [**BlockRelationRequest**](BlockRelationRequest.md) |  | 

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


## Follow

> map[string]interface{} Follow(ctx).FollowRelationRequest(followRelationRequest).Execute()





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
	followRelationRequest := *openapiclient.NewFollowRelationRequest() // FollowRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.Follow(context.Background()).FollowRelationRequest(followRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.Follow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Follow`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationService.Follow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFollowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **followRelationRequest** | [**FollowRelationRequest**](FollowRelationRequest.md) |  | 

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


## GetStatus

> GetStatusRelationReply GetStatus(ctx).GetStatusRelationRequest(getStatusRelationRequest).Execute()





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
	getStatusRelationRequest := *openapiclient.NewGetStatusRelationRequest() // GetStatusRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.GetStatus(context.Background()).GetStatusRelationRequest(getStatusRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: GetStatusRelationReply
	fmt.Fprintf(os.Stdout, "Response from `RelationService.GetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStatusRelationRequest** | [**GetStatusRelationRequest**](GetStatusRelationRequest.md) |  | 

### Return type

[**GetStatusRelationReply**](GetStatusRelationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlocked

> ListBlockedRelationsReply ListBlocked(ctx).ListBlockedRelationsRequest(listBlockedRelationsRequest).Execute()





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
	listBlockedRelationsRequest := *openapiclient.NewListBlockedRelationsRequest() // ListBlockedRelationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.ListBlocked(context.Background()).ListBlockedRelationsRequest(listBlockedRelationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.ListBlocked``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlocked`: ListBlockedRelationsReply
	fmt.Fprintf(os.Stdout, "Response from `RelationService.ListBlocked`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listBlockedRelationsRequest** | [**ListBlockedRelationsRequest**](ListBlockedRelationsRequest.md) |  | 

### Return type

[**ListBlockedRelationsReply**](ListBlockedRelationsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFollowers

> ListFollowersRelationsReply ListFollowers(ctx).ListFollowersRelationsRequest(listFollowersRelationsRequest).Execute()





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
	listFollowersRelationsRequest := *openapiclient.NewListFollowersRelationsRequest() // ListFollowersRelationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.ListFollowers(context.Background()).ListFollowersRelationsRequest(listFollowersRelationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.ListFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowers`: ListFollowersRelationsReply
	fmt.Fprintf(os.Stdout, "Response from `RelationService.ListFollowers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listFollowersRelationsRequest** | [**ListFollowersRelationsRequest**](ListFollowersRelationsRequest.md) |  | 

### Return type

[**ListFollowersRelationsReply**](ListFollowersRelationsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFollowing

> ListFollowingRelationsReply ListFollowing(ctx).ListFollowingRelationsRequest(listFollowingRelationsRequest).Execute()





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
	listFollowingRelationsRequest := *openapiclient.NewListFollowingRelationsRequest() // ListFollowingRelationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.ListFollowing(context.Background()).ListFollowingRelationsRequest(listFollowingRelationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.ListFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowing`: ListFollowingRelationsReply
	fmt.Fprintf(os.Stdout, "Response from `RelationService.ListFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listFollowingRelationsRequest** | [**ListFollowingRelationsRequest**](ListFollowingRelationsRequest.md) |  | 

### Return type

[**ListFollowingRelationsReply**](ListFollowingRelationsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unblock

> map[string]interface{} Unblock(ctx).UnblockRelationRequest(unblockRelationRequest).Execute()





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
	unblockRelationRequest := *openapiclient.NewUnblockRelationRequest() // UnblockRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.Unblock(context.Background()).UnblockRelationRequest(unblockRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.Unblock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Unblock`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationService.Unblock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnblockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unblockRelationRequest** | [**UnblockRelationRequest**](UnblockRelationRequest.md) |  | 

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


## Unfollow

> map[string]interface{} Unfollow(ctx).UnfollowRelationRequest(unfollowRelationRequest).Execute()





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
	unfollowRelationRequest := *openapiclient.NewUnfollowRelationRequest() // UnfollowRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.Unfollow(context.Background()).UnfollowRelationRequest(unfollowRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.Unfollow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Unfollow`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationService.Unfollow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unfollowRelationRequest** | [**UnfollowRelationRequest**](UnfollowRelationRequest.md) |  | 

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

