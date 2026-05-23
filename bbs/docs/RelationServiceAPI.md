# \RelationServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RelationServiceBatchGetStatus**](RelationServiceAPI.md#RelationServiceBatchGetStatus) | **Post** /v1/user/relation/batch-get-status | 
[**RelationServiceBlock**](RelationServiceAPI.md#RelationServiceBlock) | **Post** /v1/user/relation/block | 
[**RelationServiceFollow**](RelationServiceAPI.md#RelationServiceFollow) | **Post** /v1/user/relation/follow | 
[**RelationServicePageBlocked**](RelationServiceAPI.md#RelationServicePageBlocked) | **Post** /v1/user/relation/page-blocked | 
[**RelationServicePageFollowers**](RelationServiceAPI.md#RelationServicePageFollowers) | **Post** /v1/user/relation/page-followers | 
[**RelationServicePageFollowing**](RelationServiceAPI.md#RelationServicePageFollowing) | **Post** /v1/user/relation/page-following | 
[**RelationServiceUnblock**](RelationServiceAPI.md#RelationServiceUnblock) | **Post** /v1/user/relation/unblock | 
[**RelationServiceUnfollow**](RelationServiceAPI.md#RelationServiceUnfollow) | **Post** /v1/user/relation/unfollow | 



## RelationServiceBatchGetStatus

> CommonApiAppBbsV1UserBatchGetStatusRelationReply RelationServiceBatchGetStatus(ctx).CommonApiAppBbsV1UserBatchGetStatusRelationRequest(commonApiAppBbsV1UserBatchGetStatusRelationRequest).Execute()





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
	commonApiAppBbsV1UserBatchGetStatusRelationRequest := *openapiclient.NewCommonApiAppBbsV1UserBatchGetStatusRelationRequest() // CommonApiAppBbsV1UserBatchGetStatusRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceBatchGetStatus(context.Background()).CommonApiAppBbsV1UserBatchGetStatusRelationRequest(commonApiAppBbsV1UserBatchGetStatusRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServiceBatchGetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServiceBatchGetStatus`: CommonApiAppBbsV1UserBatchGetStatusRelationReply
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServiceBatchGetStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServiceBatchGetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserBatchGetStatusRelationRequest** | [**CommonApiAppBbsV1UserBatchGetStatusRelationRequest**](CommonApiAppBbsV1UserBatchGetStatusRelationRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserBatchGetStatusRelationReply**](CommonApiAppBbsV1UserBatchGetStatusRelationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelationServiceBlock

> map[string]interface{} RelationServiceBlock(ctx).CommonApiAppBbsV1UserBlockRelationRequest(commonApiAppBbsV1UserBlockRelationRequest).Execute()





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
	commonApiAppBbsV1UserBlockRelationRequest := *openapiclient.NewCommonApiAppBbsV1UserBlockRelationRequest() // CommonApiAppBbsV1UserBlockRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceBlock(context.Background()).CommonApiAppBbsV1UserBlockRelationRequest(commonApiAppBbsV1UserBlockRelationRequest).Execute()
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
 **commonApiAppBbsV1UserBlockRelationRequest** | [**CommonApiAppBbsV1UserBlockRelationRequest**](CommonApiAppBbsV1UserBlockRelationRequest.md) |  | 

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

> map[string]interface{} RelationServiceFollow(ctx).CommonApiAppBbsV1UserFollowRelationRequest(commonApiAppBbsV1UserFollowRelationRequest).Execute()





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
	commonApiAppBbsV1UserFollowRelationRequest := *openapiclient.NewCommonApiAppBbsV1UserFollowRelationRequest() // CommonApiAppBbsV1UserFollowRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceFollow(context.Background()).CommonApiAppBbsV1UserFollowRelationRequest(commonApiAppBbsV1UserFollowRelationRequest).Execute()
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
 **commonApiAppBbsV1UserFollowRelationRequest** | [**CommonApiAppBbsV1UserFollowRelationRequest**](CommonApiAppBbsV1UserFollowRelationRequest.md) |  | 

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


## RelationServicePageBlocked

> CommonApiAppBbsV1UserPageBlockedRelationReply RelationServicePageBlocked(ctx).CommonApiAppBbsV1UserPageBlockedRelationRequest(commonApiAppBbsV1UserPageBlockedRelationRequest).Execute()





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
	commonApiAppBbsV1UserPageBlockedRelationRequest := *openapiclient.NewCommonApiAppBbsV1UserPageBlockedRelationRequest() // CommonApiAppBbsV1UserPageBlockedRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationServiceAPI.RelationServicePageBlocked(context.Background()).CommonApiAppBbsV1UserPageBlockedRelationRequest(commonApiAppBbsV1UserPageBlockedRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServicePageBlocked``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServicePageBlocked`: CommonApiAppBbsV1UserPageBlockedRelationReply
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServicePageBlocked`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServicePageBlockedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserPageBlockedRelationRequest** | [**CommonApiAppBbsV1UserPageBlockedRelationRequest**](CommonApiAppBbsV1UserPageBlockedRelationRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserPageBlockedRelationReply**](CommonApiAppBbsV1UserPageBlockedRelationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelationServicePageFollowers

> CommonApiAppBbsV1UserPageFollowersRelationReply RelationServicePageFollowers(ctx).CommonApiAppBbsV1UserPageFollowersRelationRequest(commonApiAppBbsV1UserPageFollowersRelationRequest).Execute()





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
	commonApiAppBbsV1UserPageFollowersRelationRequest := *openapiclient.NewCommonApiAppBbsV1UserPageFollowersRelationRequest() // CommonApiAppBbsV1UserPageFollowersRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationServiceAPI.RelationServicePageFollowers(context.Background()).CommonApiAppBbsV1UserPageFollowersRelationRequest(commonApiAppBbsV1UserPageFollowersRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServicePageFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServicePageFollowers`: CommonApiAppBbsV1UserPageFollowersRelationReply
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServicePageFollowers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServicePageFollowersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserPageFollowersRelationRequest** | [**CommonApiAppBbsV1UserPageFollowersRelationRequest**](CommonApiAppBbsV1UserPageFollowersRelationRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserPageFollowersRelationReply**](CommonApiAppBbsV1UserPageFollowersRelationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelationServicePageFollowing

> CommonApiAppBbsV1UserPageFollowingRelationReply RelationServicePageFollowing(ctx).CommonApiAppBbsV1UserPageFollowingRelationRequest(commonApiAppBbsV1UserPageFollowingRelationRequest).Execute()





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
	commonApiAppBbsV1UserPageFollowingRelationRequest := *openapiclient.NewCommonApiAppBbsV1UserPageFollowingRelationRequest() // CommonApiAppBbsV1UserPageFollowingRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationServiceAPI.RelationServicePageFollowing(context.Background()).CommonApiAppBbsV1UserPageFollowingRelationRequest(commonApiAppBbsV1UserPageFollowingRelationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationServiceAPI.RelationServicePageFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelationServicePageFollowing`: CommonApiAppBbsV1UserPageFollowingRelationReply
	fmt.Fprintf(os.Stdout, "Response from `RelationServiceAPI.RelationServicePageFollowing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRelationServicePageFollowingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonApiAppBbsV1UserPageFollowingRelationRequest** | [**CommonApiAppBbsV1UserPageFollowingRelationRequest**](CommonApiAppBbsV1UserPageFollowingRelationRequest.md) |  | 

### Return type

[**CommonApiAppBbsV1UserPageFollowingRelationReply**](CommonApiAppBbsV1UserPageFollowingRelationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RelationServiceUnblock

> map[string]interface{} RelationServiceUnblock(ctx).CommonApiAppBbsV1UserUnblockRelationRequest(commonApiAppBbsV1UserUnblockRelationRequest).Execute()





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
	commonApiAppBbsV1UserUnblockRelationRequest := *openapiclient.NewCommonApiAppBbsV1UserUnblockRelationRequest() // CommonApiAppBbsV1UserUnblockRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceUnblock(context.Background()).CommonApiAppBbsV1UserUnblockRelationRequest(commonApiAppBbsV1UserUnblockRelationRequest).Execute()
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
 **commonApiAppBbsV1UserUnblockRelationRequest** | [**CommonApiAppBbsV1UserUnblockRelationRequest**](CommonApiAppBbsV1UserUnblockRelationRequest.md) |  | 

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

> map[string]interface{} RelationServiceUnfollow(ctx).CommonApiAppBbsV1UserUnfollowRelationRequest(commonApiAppBbsV1UserUnfollowRelationRequest).Execute()





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
	commonApiAppBbsV1UserUnfollowRelationRequest := *openapiclient.NewCommonApiAppBbsV1UserUnfollowRelationRequest() // CommonApiAppBbsV1UserUnfollowRelationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelationServiceAPI.RelationServiceUnfollow(context.Background()).CommonApiAppBbsV1UserUnfollowRelationRequest(commonApiAppBbsV1UserUnfollowRelationRequest).Execute()
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
 **commonApiAppBbsV1UserUnfollowRelationRequest** | [**CommonApiAppBbsV1UserUnfollowRelationRequest**](CommonApiAppBbsV1UserUnfollowRelationRequest.md) |  | 

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

