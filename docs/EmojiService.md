# \EmojiService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddItem**](EmojiService.md#AddItem) | **Post** /v1/im/emoji/items | 
[**CreateGroup**](EmojiService.md#CreateGroup) | **Post** /v1/im/emoji/groups | 
[**CreateShare**](EmojiService.md#CreateShare) | **Post** /v1/im/emoji/shares | 
[**DeleteGroup**](EmojiService.md#DeleteGroup) | **Post** /v1/im/emoji/groups/delete | 
[**DeleteItem**](EmojiService.md#DeleteItem) | **Post** /v1/im/emoji/items/delete | 
[**ImportShare**](EmojiService.md#ImportShare) | **Post** /v1/im/emoji/shares/import | 
[**List**](EmojiService.md#List) | **Post** /v1/im/emoji/groups/list | 
[**MoveItem**](EmojiService.md#MoveItem) | **Post** /v1/im/emoji/items/move | 
[**RenameGroup**](EmojiService.md#RenameGroup) | **Post** /v1/im/emoji/groups/rename | 
[**RenameItem**](EmojiService.md#RenameItem) | **Post** /v1/im/emoji/items/rename | 
[**ReorderGroups**](EmojiService.md#ReorderGroups) | **Post** /v1/im/emoji/groups/reorder | 
[**ReorderItems**](EmojiService.md#ReorderItems) | **Post** /v1/im/emoji/items/reorder | 



## AddItem

> AddEmojiItemResp AddItem(ctx).AddEmojiItemReq(addEmojiItemReq).Execute()





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
	addEmojiItemReq := *openapiclient.NewAddEmojiItemReq() // AddEmojiItemReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.AddItem(context.Background()).AddEmojiItemReq(addEmojiItemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.AddItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddItem`: AddEmojiItemResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.AddItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addEmojiItemReq** | [**AddEmojiItemReq**](AddEmojiItemReq.md) |  | 

### Return type

[**AddEmojiItemResp**](AddEmojiItemResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> CreateEmojiGroupResp CreateGroup(ctx).CreateEmojiGroupReq(createEmojiGroupReq).Execute()





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
	createEmojiGroupReq := *openapiclient.NewCreateEmojiGroupReq() // CreateEmojiGroupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.CreateGroup(context.Background()).CreateEmojiGroupReq(createEmojiGroupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: CreateEmojiGroupResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmojiGroupReq** | [**CreateEmojiGroupReq**](CreateEmojiGroupReq.md) |  | 

### Return type

[**CreateEmojiGroupResp**](CreateEmojiGroupResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShare

> CreateEmojiShareResp CreateShare(ctx).CreateEmojiShareReq(createEmojiShareReq).Execute()





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
	createEmojiShareReq := *openapiclient.NewCreateEmojiShareReq() // CreateEmojiShareReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.CreateShare(context.Background()).CreateEmojiShareReq(createEmojiShareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.CreateShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShare`: CreateEmojiShareResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.CreateShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmojiShareReq** | [**CreateEmojiShareReq**](CreateEmojiShareReq.md) |  | 

### Return type

[**CreateEmojiShareResp**](CreateEmojiShareResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> map[string]interface{} DeleteGroup(ctx).DeleteEmojiGroupReq(deleteEmojiGroupReq).Execute()





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
	deleteEmojiGroupReq := *openapiclient.NewDeleteEmojiGroupReq() // DeleteEmojiGroupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.DeleteGroup(context.Background()).DeleteEmojiGroupReq(deleteEmojiGroupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.DeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteGroup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.DeleteGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteEmojiGroupReq** | [**DeleteEmojiGroupReq**](DeleteEmojiGroupReq.md) |  | 

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


## DeleteItem

> map[string]interface{} DeleteItem(ctx).DeleteEmojiItemReq(deleteEmojiItemReq).Execute()





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
	deleteEmojiItemReq := *openapiclient.NewDeleteEmojiItemReq() // DeleteEmojiItemReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.DeleteItem(context.Background()).DeleteEmojiItemReq(deleteEmojiItemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.DeleteItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteItem`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.DeleteItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteEmojiItemReq** | [**DeleteEmojiItemReq**](DeleteEmojiItemReq.md) |  | 

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


## ImportShare

> ImportEmojiShareResp ImportShare(ctx).ImportEmojiShareReq(importEmojiShareReq).Execute()





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
	importEmojiShareReq := *openapiclient.NewImportEmojiShareReq() // ImportEmojiShareReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.ImportShare(context.Background()).ImportEmojiShareReq(importEmojiShareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.ImportShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportShare`: ImportEmojiShareResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.ImportShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importEmojiShareReq** | [**ImportEmojiShareReq**](ImportEmojiShareReq.md) |  | 

### Return type

[**ImportEmojiShareResp**](ImportEmojiShareResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListEmojiGroupsResp List(ctx).ListEmojiGroupsReq(listEmojiGroupsReq).Execute()





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
	listEmojiGroupsReq := *openapiclient.NewListEmojiGroupsReq() // ListEmojiGroupsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.List(context.Background()).ListEmojiGroupsReq(listEmojiGroupsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListEmojiGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listEmojiGroupsReq** | [**ListEmojiGroupsReq**](ListEmojiGroupsReq.md) |  | 

### Return type

[**ListEmojiGroupsResp**](ListEmojiGroupsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveItem

> MoveEmojiItemResp MoveItem(ctx).MoveEmojiItemReq(moveEmojiItemReq).Execute()





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
	moveEmojiItemReq := *openapiclient.NewMoveEmojiItemReq() // MoveEmojiItemReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.MoveItem(context.Background()).MoveEmojiItemReq(moveEmojiItemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.MoveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveItem`: MoveEmojiItemResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.MoveItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moveEmojiItemReq** | [**MoveEmojiItemReq**](MoveEmojiItemReq.md) |  | 

### Return type

[**MoveEmojiItemResp**](MoveEmojiItemResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameGroup

> RenameEmojiGroupResp RenameGroup(ctx).RenameEmojiGroupReq(renameEmojiGroupReq).Execute()





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
	renameEmojiGroupReq := *openapiclient.NewRenameEmojiGroupReq() // RenameEmojiGroupReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.RenameGroup(context.Background()).RenameEmojiGroupReq(renameEmojiGroupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.RenameGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameGroup`: RenameEmojiGroupResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.RenameGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameEmojiGroupReq** | [**RenameEmojiGroupReq**](RenameEmojiGroupReq.md) |  | 

### Return type

[**RenameEmojiGroupResp**](RenameEmojiGroupResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameItem

> RenameEmojiItemResp RenameItem(ctx).RenameEmojiItemReq(renameEmojiItemReq).Execute()





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
	renameEmojiItemReq := *openapiclient.NewRenameEmojiItemReq() // RenameEmojiItemReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.RenameItem(context.Background()).RenameEmojiItemReq(renameEmojiItemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.RenameItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameItem`: RenameEmojiItemResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.RenameItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameEmojiItemReq** | [**RenameEmojiItemReq**](RenameEmojiItemReq.md) |  | 

### Return type

[**RenameEmojiItemResp**](RenameEmojiItemResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderGroups

> map[string]interface{} ReorderGroups(ctx).ReorderEmojiGroupsReq(reorderEmojiGroupsReq).Execute()





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
	reorderEmojiGroupsReq := *openapiclient.NewReorderEmojiGroupsReq() // ReorderEmojiGroupsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.ReorderGroups(context.Background()).ReorderEmojiGroupsReq(reorderEmojiGroupsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.ReorderGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderGroups`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.ReorderGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reorderEmojiGroupsReq** | [**ReorderEmojiGroupsReq**](ReorderEmojiGroupsReq.md) |  | 

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


## ReorderItems

> map[string]interface{} ReorderItems(ctx).ReorderEmojiItemsReq(reorderEmojiItemsReq).Execute()





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
	reorderEmojiItemsReq := *openapiclient.NewReorderEmojiItemsReq() // ReorderEmojiItemsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmojiService.ReorderItems(context.Background()).ReorderEmojiItemsReq(reorderEmojiItemsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.ReorderItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.ReorderItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reorderEmojiItemsReq** | [**ReorderEmojiItemsReq**](ReorderEmojiItemsReq.md) |  | 

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

