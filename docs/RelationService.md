# \RelationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockRelation**](RelationService.md#BlockRelation) | **Post** /v1/user/relation/block | 
[**FollowRelation**](RelationService.md#FollowRelation) | **Post** /v1/user/relation/follow | 
[**GetStatusRelation**](RelationService.md#GetStatusRelation) | **Post** /v1/user/relation/get-status | 
[**ListBlockedRelations**](RelationService.md#ListBlockedRelations) | **Post** /v1/user/relation/list-blocked | 
[**ListFollowersRelations**](RelationService.md#ListFollowersRelations) | **Post** /v1/user/relation/list-followers | 
[**ListFollowingRelations**](RelationService.md#ListFollowingRelations) | **Post** /v1/user/relation/list-following | 
[**UnblockRelation**](RelationService.md#UnblockRelation) | **Post** /v1/user/relation/unblock | 
[**UnfollowRelation**](RelationService.md#UnfollowRelation) | **Post** /v1/user/relation/unfollow | 



## BlockRelation

> map[string]interface{} BlockRelation(ctx).BlockRelationReq(blockRelationReq).Execute()





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
	resp, r, err := apiClient.RelationService.BlockRelation(context.Background()).BlockRelationReq(blockRelationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.BlockRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockRelation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationService.BlockRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockRelationRequest struct via the builder pattern


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


## FollowRelation

> map[string]interface{} FollowRelation(ctx).FollowRelationReq(followRelationReq).Execute()





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
	resp, r, err := apiClient.RelationService.FollowRelation(context.Background()).FollowRelationReq(followRelationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.FollowRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowRelation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationService.FollowRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFollowRelationRequest struct via the builder pattern


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


## GetStatusRelation

> GetStatusRelationResp GetStatusRelation(ctx).GetStatusRelationReq(getStatusRelationReq).Execute()





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
	resp, r, err := apiClient.RelationService.GetStatusRelation(context.Background()).GetStatusRelationReq(getStatusRelationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.GetStatusRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusRelation`: GetStatusRelationResp
	fmt.Fprintf(os.Stdout, "Response from `RelationService.GetStatusRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRelationRequest struct via the builder pattern


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


## ListBlockedRelations

> ListBlockedRelationsResp ListBlockedRelations(ctx).ListBlockedRelationsReq(listBlockedRelationsReq).Execute()





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
	resp, r, err := apiClient.RelationService.ListBlockedRelations(context.Background()).ListBlockedRelationsReq(listBlockedRelationsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.ListBlockedRelations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlockedRelations`: ListBlockedRelationsResp
	fmt.Fprintf(os.Stdout, "Response from `RelationService.ListBlockedRelations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlockedRelationsRequest struct via the builder pattern


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


## ListFollowersRelations

> ListFollowersRelationsResp ListFollowersRelations(ctx).ListFollowersRelationsReq(listFollowersRelationsReq).Execute()





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
	resp, r, err := apiClient.RelationService.ListFollowersRelations(context.Background()).ListFollowersRelationsReq(listFollowersRelationsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.ListFollowersRelations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowersRelations`: ListFollowersRelationsResp
	fmt.Fprintf(os.Stdout, "Response from `RelationService.ListFollowersRelations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowersRelationsRequest struct via the builder pattern


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


## ListFollowingRelations

> ListFollowingRelationsResp ListFollowingRelations(ctx).ListFollowingRelationsReq(listFollowingRelationsReq).Execute()





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
	resp, r, err := apiClient.RelationService.ListFollowingRelations(context.Background()).ListFollowingRelationsReq(listFollowingRelationsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.ListFollowingRelations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFollowingRelations`: ListFollowingRelationsResp
	fmt.Fprintf(os.Stdout, "Response from `RelationService.ListFollowingRelations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFollowingRelationsRequest struct via the builder pattern


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


## UnblockRelation

> map[string]interface{} UnblockRelation(ctx).UnblockRelationReq(unblockRelationReq).Execute()





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
	resp, r, err := apiClient.RelationService.UnblockRelation(context.Background()).UnblockRelationReq(unblockRelationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.UnblockRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnblockRelation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationService.UnblockRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnblockRelationRequest struct via the builder pattern


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


## UnfollowRelation

> map[string]interface{} UnfollowRelation(ctx).UnfollowRelationReq(unfollowRelationReq).Execute()





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
	resp, r, err := apiClient.RelationService.UnfollowRelation(context.Background()).UnfollowRelationReq(unfollowRelationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelationService.UnfollowRelation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnfollowRelation`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelationService.UnfollowRelation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowRelationRequest struct via the builder pattern


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

