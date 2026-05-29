# \RelationServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RelationServiceBlock**](RelationServiceAPI.md#RelationServiceBlock) | **Post** /v1/user/relation/block | 
[**RelationServiceFollow**](RelationServiceAPI.md#RelationServiceFollow) | **Post** /v1/user/relation/follow | 
[**RelationServiceGetStatus**](RelationServiceAPI.md#RelationServiceGetStatus) | **Post** /v1/user/relation/get-status | 
[**RelationServiceListBlocked**](RelationServiceAPI.md#RelationServiceListBlocked) | **Post** /v1/user/relation/list-blocked | 
[**RelationServiceListFollowers**](RelationServiceAPI.md#RelationServiceListFollowers) | **Post** /v1/user/relation/list-followers | 
[**RelationServiceListFollowing**](RelationServiceAPI.md#RelationServiceListFollowing) | **Post** /v1/user/relation/list-following | 
[**RelationServiceUnblock**](RelationServiceAPI.md#RelationServiceUnblock) | **Post** /v1/user/relation/unblock | 
[**RelationServiceUnfollow**](RelationServiceAPI.md#RelationServiceUnfollow) | **Post** /v1/user/relation/unfollow | 



## RelationServiceBlock

> map[string]interface{} RelationServiceBlock(ctx).BlockRelationRequest(blockRelationRequest).Execute()





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
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceBlock(context.Background()).BlockRelationRequest(blockRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServiceBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServiceBlock`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServiceBlock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServiceBlockRequest struct via the builder pattern


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


## RelationServiceFollow

> map[string]interface{} RelationServiceFollow(ctx).FollowRelationRequest(followRelationRequest).Execute()





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
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceFollow(context.Background()).FollowRelationRequest(followRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServiceFollow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServiceFollow`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServiceFollow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServiceFollowRequest struct via the builder pattern


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


## RelationServiceGetStatus

> GetStatusRelationReply RelationServiceGetStatus(ctx).GetStatusRelationRequest(getStatusRelationRequest).Execute()





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
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceGetStatus(context.Background()).GetStatusRelationRequest(getStatusRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServiceGetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServiceGetStatus`: GetStatusRelationReply
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServiceGetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServiceGetStatusRequest struct via the builder pattern


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


## RelationServiceListBlocked

> ListBlockedRelationsReply RelationServiceListBlocked(ctx).ListBlockedRelationsRequest(listBlockedRelationsRequest).Execute()





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
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceListBlocked(context.Background()).ListBlockedRelationsRequest(listBlockedRelationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServiceListBlocked``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServiceListBlocked`: ListBlockedRelationsReply
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServiceListBlocked`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServiceListBlockedRequest struct via the builder pattern


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


## RelationServiceListFollowers

> ListFollowersRelationsReply RelationServiceListFollowers(ctx).ListFollowersRelationsRequest(listFollowersRelationsRequest).Execute()





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
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceListFollowers(context.Background()).ListFollowersRelationsRequest(listFollowersRelationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServiceListFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServiceListFollowers`: ListFollowersRelationsReply
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServiceListFollowers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServiceListFollowersRequest struct via the builder pattern


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


## RelationServiceListFollowing

> ListFollowingRelationsReply RelationServiceListFollowing(ctx).ListFollowingRelationsRequest(listFollowingRelationsRequest).Execute()





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
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceListFollowing(context.Background()).ListFollowingRelationsRequest(listFollowingRelationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServiceListFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServiceListFollowing`: ListFollowingRelationsReply
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServiceListFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServiceListFollowingRequest struct via the builder pattern


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


## RelationServiceUnblock

> map[string]interface{} RelationServiceUnblock(ctx).UnblockRelationRequest(unblockRelationRequest).Execute()





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
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceUnblock(context.Background()).UnblockRelationRequest(unblockRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServiceUnblock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServiceUnblock`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServiceUnblock`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServiceUnblockRequest struct via the builder pattern


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


## RelationServiceUnfollow

> map[string]interface{} RelationServiceUnfollow(ctx).UnfollowRelationRequest(unfollowRelationRequest).Execute()





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
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceUnfollow(context.Background()).UnfollowRelationRequest(unfollowRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServiceUnfollow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServiceUnfollow`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServiceUnfollow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServiceUnfollowRequest struct via the builder pattern


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

