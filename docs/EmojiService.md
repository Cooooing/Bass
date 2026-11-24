# \EmojiService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEmojiItem**](EmojiService.md#AddEmojiItem) | **Post** /v1/im/emoji/items | 
[**CreateEmojiGroup**](EmojiService.md#CreateEmojiGroup) | **Post** /v1/im/emoji/groups | 
[**CreateEmojiShare**](EmojiService.md#CreateEmojiShare) | **Post** /v1/im/emoji/shares | 
[**DeleteEmojiGroup**](EmojiService.md#DeleteEmojiGroup) | **Post** /v1/im/emoji/groups/delete | 
[**DeleteEmojiItem**](EmojiService.md#DeleteEmojiItem) | **Post** /v1/im/emoji/items/delete | 
[**ImportEmojiShare**](EmojiService.md#ImportEmojiShare) | **Post** /v1/im/emoji/shares/import | 
[**ListEmojiGroups**](EmojiService.md#ListEmojiGroups) | **Post** /v1/im/emoji/groups/list | 
[**MoveEmojiItem**](EmojiService.md#MoveEmojiItem) | **Post** /v1/im/emoji/items/move | 
[**RenameEmojiGroup**](EmojiService.md#RenameEmojiGroup) | **Post** /v1/im/emoji/groups/rename | 
[**RenameEmojiItem**](EmojiService.md#RenameEmojiItem) | **Post** /v1/im/emoji/items/rename | 
[**ReorderEmojiGroups**](EmojiService.md#ReorderEmojiGroups) | **Post** /v1/im/emoji/groups/reorder | 
[**ReorderEmojiItems**](EmojiService.md#ReorderEmojiItems) | **Post** /v1/im/emoji/items/reorder | 



## AddEmojiItem

> AddEmojiItemResp AddEmojiItem(ctx).AddEmojiItemReq(addEmojiItemReq).Execute()





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
	resp, r, err := apiClient.EmojiService.AddEmojiItem(context.Background()).AddEmojiItemReq(addEmojiItemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.AddEmojiItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddEmojiItem`: AddEmojiItemResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.AddEmojiItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddEmojiItemRequest struct via the builder pattern


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


## CreateEmojiGroup

> CreateEmojiGroupResp CreateEmojiGroup(ctx).CreateEmojiGroupReq(createEmojiGroupReq).Execute()





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
	resp, r, err := apiClient.EmojiService.CreateEmojiGroup(context.Background()).CreateEmojiGroupReq(createEmojiGroupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.CreateEmojiGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmojiGroup`: CreateEmojiGroupResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.CreateEmojiGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmojiGroupRequest struct via the builder pattern


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


## CreateEmojiShare

> CreateEmojiShareResp CreateEmojiShare(ctx).CreateEmojiShareReq(createEmojiShareReq).Execute()





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
	resp, r, err := apiClient.EmojiService.CreateEmojiShare(context.Background()).CreateEmojiShareReq(createEmojiShareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.CreateEmojiShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmojiShare`: CreateEmojiShareResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.CreateEmojiShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmojiShareRequest struct via the builder pattern


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


## DeleteEmojiGroup

> map[string]interface{} DeleteEmojiGroup(ctx).DeleteEmojiGroupReq(deleteEmojiGroupReq).Execute()





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
	resp, r, err := apiClient.EmojiService.DeleteEmojiGroup(context.Background()).DeleteEmojiGroupReq(deleteEmojiGroupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.DeleteEmojiGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEmojiGroup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.DeleteEmojiGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmojiGroupRequest struct via the builder pattern


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


## DeleteEmojiItem

> map[string]interface{} DeleteEmojiItem(ctx).DeleteEmojiItemReq(deleteEmojiItemReq).Execute()





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
	resp, r, err := apiClient.EmojiService.DeleteEmojiItem(context.Background()).DeleteEmojiItemReq(deleteEmojiItemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.DeleteEmojiItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEmojiItem`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.DeleteEmojiItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmojiItemRequest struct via the builder pattern


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


## ImportEmojiShare

> ImportEmojiShareResp ImportEmojiShare(ctx).ImportEmojiShareReq(importEmojiShareReq).Execute()





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
	resp, r, err := apiClient.EmojiService.ImportEmojiShare(context.Background()).ImportEmojiShareReq(importEmojiShareReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.ImportEmojiShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportEmojiShare`: ImportEmojiShareResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.ImportEmojiShare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportEmojiShareRequest struct via the builder pattern


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


## ListEmojiGroups

> ListEmojiGroupsResp ListEmojiGroups(ctx).ListEmojiGroupsReq(listEmojiGroupsReq).Execute()





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
	resp, r, err := apiClient.EmojiService.ListEmojiGroups(context.Background()).ListEmojiGroupsReq(listEmojiGroupsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.ListEmojiGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEmojiGroups`: ListEmojiGroupsResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.ListEmojiGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmojiGroupsRequest struct via the builder pattern


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


## MoveEmojiItem

> MoveEmojiItemResp MoveEmojiItem(ctx).MoveEmojiItemReq(moveEmojiItemReq).Execute()





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
	resp, r, err := apiClient.EmojiService.MoveEmojiItem(context.Background()).MoveEmojiItemReq(moveEmojiItemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.MoveEmojiItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveEmojiItem`: MoveEmojiItemResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.MoveEmojiItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveEmojiItemRequest struct via the builder pattern


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


## RenameEmojiGroup

> RenameEmojiGroupResp RenameEmojiGroup(ctx).RenameEmojiGroupReq(renameEmojiGroupReq).Execute()





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
	resp, r, err := apiClient.EmojiService.RenameEmojiGroup(context.Background()).RenameEmojiGroupReq(renameEmojiGroupReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.RenameEmojiGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameEmojiGroup`: RenameEmojiGroupResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.RenameEmojiGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameEmojiGroupRequest struct via the builder pattern


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


## RenameEmojiItem

> RenameEmojiItemResp RenameEmojiItem(ctx).RenameEmojiItemReq(renameEmojiItemReq).Execute()





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
	resp, r, err := apiClient.EmojiService.RenameEmojiItem(context.Background()).RenameEmojiItemReq(renameEmojiItemReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.RenameEmojiItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameEmojiItem`: RenameEmojiItemResp
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.RenameEmojiItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameEmojiItemRequest struct via the builder pattern


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


## ReorderEmojiGroups

> map[string]interface{} ReorderEmojiGroups(ctx).ReorderEmojiGroupsReq(reorderEmojiGroupsReq).Execute()





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
	resp, r, err := apiClient.EmojiService.ReorderEmojiGroups(context.Background()).ReorderEmojiGroupsReq(reorderEmojiGroupsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.ReorderEmojiGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderEmojiGroups`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.ReorderEmojiGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderEmojiGroupsRequest struct via the builder pattern


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


## ReorderEmojiItems

> map[string]interface{} ReorderEmojiItems(ctx).ReorderEmojiItemsReq(reorderEmojiItemsReq).Execute()





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
	resp, r, err := apiClient.EmojiService.ReorderEmojiItems(context.Background()).ReorderEmojiItemsReq(reorderEmojiItemsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmojiService.ReorderEmojiItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderEmojiItems`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `EmojiService.ReorderEmojiItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderEmojiItemsRequest struct via the builder pattern


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

