# \PrivateChatService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountPrivateChatUnread**](PrivateChatService.md#CountPrivateChatUnread) | **Post** /v1/im/private-chat/unread-count | 
[**ListPrivateChatMessages**](PrivateChatService.md#ListPrivateChatMessages) | **Post** /v1/im/private-chat/messages/list | 
[**ListPrivateChatSessions**](PrivateChatService.md#ListPrivateChatSessions) | **Get** /v1/im/private-chat/sessions | 
[**MarkPrivateChatSessionRead**](PrivateChatService.md#MarkPrivateChatSessionRead) | **Post** /v1/im/private-chat/read | 
[**RevokePrivateChatMessage**](PrivateChatService.md#RevokePrivateChatMessage) | **Post** /v1/im/private-chat/messages/revoke | 
[**SendPrivateChatMessage**](PrivateChatService.md#SendPrivateChatMessage) | **Post** /v1/im/private-chat/messages | 
[**SetPrivateChatSessionMuted**](PrivateChatService.md#SetPrivateChatSessionMuted) | **Post** /v1/im/private-chat/muted | 
[**SetPrivateChatSessionPinned**](PrivateChatService.md#SetPrivateChatSessionPinned) | **Post** /v1/im/private-chat/pinned | 



## CountPrivateChatUnread

> CountPrivateChatUnreadResp CountPrivateChatUnread(ctx).Body(body).Execute()





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
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateChatService.CountPrivateChatUnread(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateChatService.CountPrivateChatUnread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountPrivateChatUnread`: CountPrivateChatUnreadResp
	fmt.Fprintf(os.Stdout, "Response from `PrivateChatService.CountPrivateChatUnread`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountPrivateChatUnreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**CountPrivateChatUnreadResp**](CountPrivateChatUnreadResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateChatMessages

> ListPrivateChatMessagesResp ListPrivateChatMessages(ctx).ListPrivateChatMessagesReq(listPrivateChatMessagesReq).Execute()





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
	listPrivateChatMessagesReq := *openapiclient.NewListPrivateChatMessagesReq() // ListPrivateChatMessagesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateChatService.ListPrivateChatMessages(context.Background()).ListPrivateChatMessagesReq(listPrivateChatMessagesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateChatService.ListPrivateChatMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrivateChatMessages`: ListPrivateChatMessagesResp
	fmt.Fprintf(os.Stdout, "Response from `PrivateChatService.ListPrivateChatMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateChatMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listPrivateChatMessagesReq** | [**ListPrivateChatMessagesReq**](ListPrivateChatMessagesReq.md) |  | 

### Return type

[**ListPrivateChatMessagesResp**](ListPrivateChatMessagesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateChatSessions

> ListPrivateChatSessionsResp ListPrivateChatSessions(ctx).PagePage(pagePage).PageSize(pageSize).Execute()





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
	pagePage := "pagePage_example" // string | 页码 (optional)
	pageSize := "pageSize_example" // string | 页大小 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateChatService.ListPrivateChatSessions(context.Background()).PagePage(pagePage).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateChatService.ListPrivateChatSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPrivateChatSessions`: ListPrivateChatSessionsResp
	fmt.Fprintf(os.Stdout, "Response from `PrivateChatService.ListPrivateChatSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateChatSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pagePage** | **string** | 页码 | 
 **pageSize** | **string** | 页大小 | 

### Return type

[**ListPrivateChatSessionsResp**](ListPrivateChatSessionsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkPrivateChatSessionRead

> map[string]interface{} MarkPrivateChatSessionRead(ctx).MarkPrivateChatSessionReadReq(markPrivateChatSessionReadReq).Execute()





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
	markPrivateChatSessionReadReq := *openapiclient.NewMarkPrivateChatSessionReadReq() // MarkPrivateChatSessionReadReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateChatService.MarkPrivateChatSessionRead(context.Background()).MarkPrivateChatSessionReadReq(markPrivateChatSessionReadReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateChatService.MarkPrivateChatSessionRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkPrivateChatSessionRead`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrivateChatService.MarkPrivateChatSessionRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkPrivateChatSessionReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markPrivateChatSessionReadReq** | [**MarkPrivateChatSessionReadReq**](MarkPrivateChatSessionReadReq.md) |  | 

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


## RevokePrivateChatMessage

> RevokePrivateChatMessageResp RevokePrivateChatMessage(ctx).RevokePrivateChatMessageReq(revokePrivateChatMessageReq).Execute()





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
	revokePrivateChatMessageReq := *openapiclient.NewRevokePrivateChatMessageReq() // RevokePrivateChatMessageReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateChatService.RevokePrivateChatMessage(context.Background()).RevokePrivateChatMessageReq(revokePrivateChatMessageReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateChatService.RevokePrivateChatMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokePrivateChatMessage`: RevokePrivateChatMessageResp
	fmt.Fprintf(os.Stdout, "Response from `PrivateChatService.RevokePrivateChatMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokePrivateChatMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revokePrivateChatMessageReq** | [**RevokePrivateChatMessageReq**](RevokePrivateChatMessageReq.md) |  | 

### Return type

[**RevokePrivateChatMessageResp**](RevokePrivateChatMessageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPrivateChatMessage

> SendPrivateChatMessageResp SendPrivateChatMessage(ctx).SendPrivateChatMessageReq(sendPrivateChatMessageReq).Execute()





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
	sendPrivateChatMessageReq := *openapiclient.NewSendPrivateChatMessageReq() // SendPrivateChatMessageReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateChatService.SendPrivateChatMessage(context.Background()).SendPrivateChatMessageReq(sendPrivateChatMessageReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateChatService.SendPrivateChatMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendPrivateChatMessage`: SendPrivateChatMessageResp
	fmt.Fprintf(os.Stdout, "Response from `PrivateChatService.SendPrivateChatMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendPrivateChatMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendPrivateChatMessageReq** | [**SendPrivateChatMessageReq**](SendPrivateChatMessageReq.md) |  | 

### Return type

[**SendPrivateChatMessageResp**](SendPrivateChatMessageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPrivateChatSessionMuted

> map[string]interface{} SetPrivateChatSessionMuted(ctx).SetPrivateChatSessionMutedReq(setPrivateChatSessionMutedReq).Execute()





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
	setPrivateChatSessionMutedReq := *openapiclient.NewSetPrivateChatSessionMutedReq() // SetPrivateChatSessionMutedReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateChatService.SetPrivateChatSessionMuted(context.Background()).SetPrivateChatSessionMutedReq(setPrivateChatSessionMutedReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateChatService.SetPrivateChatSessionMuted``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPrivateChatSessionMuted`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrivateChatService.SetPrivateChatSessionMuted`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPrivateChatSessionMutedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPrivateChatSessionMutedReq** | [**SetPrivateChatSessionMutedReq**](SetPrivateChatSessionMutedReq.md) |  | 

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


## SetPrivateChatSessionPinned

> map[string]interface{} SetPrivateChatSessionPinned(ctx).SetPrivateChatSessionPinnedReq(setPrivateChatSessionPinnedReq).Execute()





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
	setPrivateChatSessionPinnedReq := *openapiclient.NewSetPrivateChatSessionPinnedReq() // SetPrivateChatSessionPinnedReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrivateChatService.SetPrivateChatSessionPinned(context.Background()).SetPrivateChatSessionPinnedReq(setPrivateChatSessionPinnedReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrivateChatService.SetPrivateChatSessionPinned``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetPrivateChatSessionPinned`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrivateChatService.SetPrivateChatSessionPinned`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPrivateChatSessionPinnedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPrivateChatSessionPinnedReq** | [**SetPrivateChatSessionPinnedReq**](SetPrivateChatSessionPinnedReq.md) |  | 

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

