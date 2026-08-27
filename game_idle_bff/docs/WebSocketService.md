# \WebSocketService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSession**](WebSocketService.md#CreateSession) | **Post** /v1/game-idle/ws/create-session | 



## CreateSession

> CreateWebSocketSessionResp CreateSession(ctx).CreateWebSocketSessionReq(createWebSocketSessionReq).Execute()





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
	createWebSocketSessionReq := *openapiclient.NewCreateWebSocketSessionReq() // CreateWebSocketSessionReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebSocketService.CreateSession(context.Background()).CreateWebSocketSessionReq(createWebSocketSessionReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSocketService.CreateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSession`: CreateWebSocketSessionResp
	fmt.Fprintf(os.Stdout, "Response from `WebSocketService.CreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebSocketSessionReq** | [**CreateWebSocketSessionReq**](CreateWebSocketSessionReq.md) |  | 

### Return type

[**CreateWebSocketSessionResp**](CreateWebSocketSessionResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

