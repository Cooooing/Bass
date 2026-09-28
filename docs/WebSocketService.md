# \WebSocketService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTicket**](WebSocketService.md#CreateTicket) | **Post** /v1/game-idle/ws/create-ticket | 



## CreateTicket

> CreateWebSocketTicketResp CreateTicket(ctx).CreateWebSocketTicketReq(createWebSocketTicketReq).Execute()





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
	createWebSocketTicketReq := *openapiclient.NewCreateWebSocketTicketReq() // CreateWebSocketTicketReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebSocketService.CreateTicket(context.Background()).CreateWebSocketTicketReq(createWebSocketTicketReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebSocketService.CreateTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTicket`: CreateWebSocketTicketResp
	fmt.Fprintf(os.Stdout, "Response from `WebSocketService.CreateTicket`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebSocketTicketReq** | [**CreateWebSocketTicketReq**](CreateWebSocketTicketReq.md) |  | 

### Return type

[**CreateWebSocketTicketResp**](CreateWebSocketTicketResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

