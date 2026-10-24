# \GlobalChatService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGlobalChatTopic**](GlobalChatService.md#GetGlobalChatTopic) | **Post** /v1/im/global-chat/topic/get | 
[**ListGlobalChatBarrages**](GlobalChatService.md#ListGlobalChatBarrages) | **Post** /v1/im/global-chat/barrages/list | 
[**ListGlobalChatMessages**](GlobalChatService.md#ListGlobalChatMessages) | **Post** /v1/im/global-chat/messages/list | 
[**ListGlobalChatOnlineMembers**](GlobalChatService.md#ListGlobalChatOnlineMembers) | **Post** /v1/im/global-chat/online-members/list | 
[**RevokeGlobalChatMessage**](GlobalChatService.md#RevokeGlobalChatMessage) | **Post** /v1/im/global-chat/messages/revoke | 
[**SendGlobalChatBarrage**](GlobalChatService.md#SendGlobalChatBarrage) | **Post** /v1/im/global-chat/barrages | 
[**SendGlobalChatMessage**](GlobalChatService.md#SendGlobalChatMessage) | **Post** /v1/im/global-chat/messages | 
[**SetGlobalChatTopic**](GlobalChatService.md#SetGlobalChatTopic) | **Post** /v1/im/global-chat/topic/set | 



## GetGlobalChatTopic

> GetGlobalChatTopicResp GetGlobalChatTopic(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.GlobalChatService.GetGlobalChatTopic(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalChatService.GetGlobalChatTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalChatTopic`: GetGlobalChatTopicResp
	fmt.Fprintf(os.Stdout, "Response from `GlobalChatService.GetGlobalChatTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalChatTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**GetGlobalChatTopicResp**](GetGlobalChatTopicResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalChatBarrages

> ListGlobalChatBarragesResp ListGlobalChatBarrages(ctx).ListGlobalChatBarragesReq(listGlobalChatBarragesReq).Execute()





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
	listGlobalChatBarragesReq := *openapiclient.NewListGlobalChatBarragesReq() // ListGlobalChatBarragesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalChatService.ListGlobalChatBarrages(context.Background()).ListGlobalChatBarragesReq(listGlobalChatBarragesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalChatService.ListGlobalChatBarrages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalChatBarrages`: ListGlobalChatBarragesResp
	fmt.Fprintf(os.Stdout, "Response from `GlobalChatService.ListGlobalChatBarrages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalChatBarragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listGlobalChatBarragesReq** | [**ListGlobalChatBarragesReq**](ListGlobalChatBarragesReq.md) |  | 

### Return type

[**ListGlobalChatBarragesResp**](ListGlobalChatBarragesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalChatMessages

> ListGlobalChatMessagesResp ListGlobalChatMessages(ctx).ListGlobalChatMessagesReq(listGlobalChatMessagesReq).Execute()





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
	listGlobalChatMessagesReq := *openapiclient.NewListGlobalChatMessagesReq() // ListGlobalChatMessagesReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalChatService.ListGlobalChatMessages(context.Background()).ListGlobalChatMessagesReq(listGlobalChatMessagesReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalChatService.ListGlobalChatMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalChatMessages`: ListGlobalChatMessagesResp
	fmt.Fprintf(os.Stdout, "Response from `GlobalChatService.ListGlobalChatMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalChatMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listGlobalChatMessagesReq** | [**ListGlobalChatMessagesReq**](ListGlobalChatMessagesReq.md) |  | 

### Return type

[**ListGlobalChatMessagesResp**](ListGlobalChatMessagesResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalChatOnlineMembers

> ListGlobalChatOnlineMembersResp ListGlobalChatOnlineMembers(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.GlobalChatService.ListGlobalChatOnlineMembers(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalChatService.ListGlobalChatOnlineMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGlobalChatOnlineMembers`: ListGlobalChatOnlineMembersResp
	fmt.Fprintf(os.Stdout, "Response from `GlobalChatService.ListGlobalChatOnlineMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalChatOnlineMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**ListGlobalChatOnlineMembersResp**](ListGlobalChatOnlineMembersResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeGlobalChatMessage

> RevokeGlobalChatMessageResp RevokeGlobalChatMessage(ctx).RevokeGlobalChatMessageReq(revokeGlobalChatMessageReq).Execute()





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
	revokeGlobalChatMessageReq := *openapiclient.NewRevokeGlobalChatMessageReq() // RevokeGlobalChatMessageReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalChatService.RevokeGlobalChatMessage(context.Background()).RevokeGlobalChatMessageReq(revokeGlobalChatMessageReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalChatService.RevokeGlobalChatMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokeGlobalChatMessage`: RevokeGlobalChatMessageResp
	fmt.Fprintf(os.Stdout, "Response from `GlobalChatService.RevokeGlobalChatMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeGlobalChatMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revokeGlobalChatMessageReq** | [**RevokeGlobalChatMessageReq**](RevokeGlobalChatMessageReq.md) |  | 

### Return type

[**RevokeGlobalChatMessageResp**](RevokeGlobalChatMessageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendGlobalChatBarrage

> SendGlobalChatBarrageResp SendGlobalChatBarrage(ctx).SendGlobalChatBarrageReq(sendGlobalChatBarrageReq).Execute()





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
	sendGlobalChatBarrageReq := *openapiclient.NewSendGlobalChatBarrageReq() // SendGlobalChatBarrageReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalChatService.SendGlobalChatBarrage(context.Background()).SendGlobalChatBarrageReq(sendGlobalChatBarrageReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalChatService.SendGlobalChatBarrage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendGlobalChatBarrage`: SendGlobalChatBarrageResp
	fmt.Fprintf(os.Stdout, "Response from `GlobalChatService.SendGlobalChatBarrage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendGlobalChatBarrageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendGlobalChatBarrageReq** | [**SendGlobalChatBarrageReq**](SendGlobalChatBarrageReq.md) |  | 

### Return type

[**SendGlobalChatBarrageResp**](SendGlobalChatBarrageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendGlobalChatMessage

> SendGlobalChatMessageResp SendGlobalChatMessage(ctx).SendGlobalChatMessageReq(sendGlobalChatMessageReq).Execute()





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
	sendGlobalChatMessageReq := *openapiclient.NewSendGlobalChatMessageReq() // SendGlobalChatMessageReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalChatService.SendGlobalChatMessage(context.Background()).SendGlobalChatMessageReq(sendGlobalChatMessageReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalChatService.SendGlobalChatMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendGlobalChatMessage`: SendGlobalChatMessageResp
	fmt.Fprintf(os.Stdout, "Response from `GlobalChatService.SendGlobalChatMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendGlobalChatMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendGlobalChatMessageReq** | [**SendGlobalChatMessageReq**](SendGlobalChatMessageReq.md) |  | 

### Return type

[**SendGlobalChatMessageResp**](SendGlobalChatMessageResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGlobalChatTopic

> SetGlobalChatTopicResp SetGlobalChatTopic(ctx).SetGlobalChatTopicReq(setGlobalChatTopicReq).Execute()





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
	setGlobalChatTopicReq := *openapiclient.NewSetGlobalChatTopicReq() // SetGlobalChatTopicReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GlobalChatService.SetGlobalChatTopic(context.Background()).SetGlobalChatTopicReq(setGlobalChatTopicReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GlobalChatService.SetGlobalChatTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetGlobalChatTopic`: SetGlobalChatTopicResp
	fmt.Fprintf(os.Stdout, "Response from `GlobalChatService.SetGlobalChatTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGlobalChatTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setGlobalChatTopicReq** | [**SetGlobalChatTopicReq**](SetGlobalChatTopicReq.md) |  | 

### Return type

[**SetGlobalChatTopicResp**](SetGlobalChatTopicResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

