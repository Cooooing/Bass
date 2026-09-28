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

> map[string]interface{} Block(ctx).BlockRelationReq(blockRelationReq).Execute()





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
	blockRelationReq := *openapiclient.NewBlockRelationReq("TargetId_example") // BlockRelationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.Block(context.Background()).BlockRelationReq(blockRelationReq).Execute()
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
 **blockRelationReq** | [**BlockRelationReq**](BlockRelationReq.md) |  | 

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

> map[string]interface{} Follow(ctx).FollowRelationReq(followRelationReq).Execute()





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
	followRelationReq := *openapiclient.NewFollowRelationReq("TargetId_example") // FollowRelationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.Follow(context.Background()).FollowRelationReq(followRelationReq).Execute()
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
 **followRelationReq** | [**FollowRelationReq**](FollowRelationReq.md) |  | 

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

> GetStatusRelationResp GetStatus(ctx).GetStatusRelationReq(getStatusRelationReq).Execute()





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
	getStatusRelationReq := *openapiclient.NewGetStatusRelationReq("TargetId_example") // GetStatusRelationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.GetStatus(context.Background()).GetStatusRelationReq(getStatusRelationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: GetStatusRelationResp
	fmt.Fprintf(os.Stdout, "Response from `RelationService.GetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getStatusRelationReq** | [**GetStatusRelationReq**](GetStatusRelationReq.md) |  | 

### Return type

[**GetStatusRelationResp**](GetStatusRelationResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlocked

> ListBlockedRelationsResp ListBlocked(ctx).ListBlockedRelationsReq(listBlockedRelationsReq).Execute()





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
	listBlockedRelationsReq := *openapiclient.NewListBlockedRelationsReq() // ListBlockedRelationsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.ListBlocked(context.Background()).ListBlockedRelationsReq(listBlockedRelationsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.ListBlocked``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlocked`: ListBlockedRelationsResp
	fmt.Fprintf(os.Stdout, "Response from `RelationService.ListBlocked`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listBlockedRelationsReq** | [**ListBlockedRelationsReq**](ListBlockedRelationsReq.md) |  | 

### Return type

[**ListBlockedRelationsResp**](ListBlockedRelationsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFollowers

> ListFollowersRelationsResp ListFollowers(ctx).ListFollowersRelationsReq(listFollowersRelationsReq).Execute()





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
	listFollowersRelationsReq := *openapiclient.NewListFollowersRelationsReq() // ListFollowersRelationsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.ListFollowers(context.Background()).ListFollowersRelationsReq(listFollowersRelationsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.ListFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowers`: ListFollowersRelationsResp
	fmt.Fprintf(os.Stdout, "Response from `RelationService.ListFollowers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listFollowersRelationsReq** | [**ListFollowersRelationsReq**](ListFollowersRelationsReq.md) |  | 

### Return type

[**ListFollowersRelationsResp**](ListFollowersRelationsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFollowing

> ListFollowingRelationsResp ListFollowing(ctx).ListFollowingRelationsReq(listFollowingRelationsReq).Execute()





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
	listFollowingRelationsReq := *openapiclient.NewListFollowingRelationsReq() // ListFollowingRelationsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.ListFollowing(context.Background()).ListFollowingRelationsReq(listFollowingRelationsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.ListFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowing`: ListFollowingRelationsResp
	fmt.Fprintf(os.Stdout, "Response from `RelationService.ListFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listFollowingRelationsReq** | [**ListFollowingRelationsReq**](ListFollowingRelationsReq.md) |  | 

### Return type

[**ListFollowingRelationsResp**](ListFollowingRelationsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unblock

> map[string]interface{} Unblock(ctx).UnblockRelationReq(unblockRelationReq).Execute()





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
	unblockRelationReq := *openapiclient.NewUnblockRelationReq("TargetId_example") // UnblockRelationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.Unblock(context.Background()).UnblockRelationReq(unblockRelationReq).Execute()
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
 **unblockRelationReq** | [**UnblockRelationReq**](UnblockRelationReq.md) |  | 

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

> map[string]interface{} Unfollow(ctx).UnfollowRelationReq(unfollowRelationReq).Execute()





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
	unfollowRelationReq := *openapiclient.NewUnfollowRelationReq("TargetId_example") // UnfollowRelationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationService.Unfollow(context.Background()).UnfollowRelationReq(unfollowRelationReq).Execute()
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
 **unfollowRelationReq** | [**UnfollowRelationReq**](UnfollowRelationReq.md) |  | 

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

