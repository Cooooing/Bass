# \NotificationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountUnread**](NotificationService.md#CountUnread) | **Post** /v1/notify/notification/count-unread | 
[**List**](NotificationService.md#List) | **Post** /v1/notify/notification/list | 
[**MarkRead**](NotificationService.md#MarkRead) | **Post** /v1/notify/notification/mark-read | 



## CountUnread

> CountUnreadNotificationsReply CountUnread(ctx).Body(body).Execute()



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
	resp, r, err := apiClient.NotificationService.CountUnread(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationService.CountUnread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountUnread`: CountUnreadNotificationsReply
	fmt.Fprintf(os.Stdout, "Response from `NotificationService.CountUnread`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountUnreadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**CountUnreadNotificationsReply**](CountUnreadNotificationsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListNotificationsReply List(ctx).ListNotificationsRequest(listNotificationsRequest).Execute()



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
	listNotificationsRequest := *openapiclient.NewListNotificationsRequest() // ListNotificationsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationService.List(context.Background()).ListNotificationsRequest(listNotificationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListNotificationsReply
	fmt.Fprintf(os.Stdout, "Response from `NotificationService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listNotificationsRequest** | [**ListNotificationsRequest**](ListNotificationsRequest.md) |  | 

### Return type

[**ListNotificationsReply**](ListNotificationsReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkRead

> MarkReadNotificationReply MarkRead(ctx).MarkReadNotificationRequest(markReadNotificationRequest).Execute()



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
	markReadNotificationRequest := *openapiclient.NewMarkReadNotificationRequest() // MarkReadNotificationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationService.MarkRead(context.Background()).MarkReadNotificationRequest(markReadNotificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationService.MarkRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkRead`: MarkReadNotificationReply
	fmt.Fprintf(os.Stdout, "Response from `NotificationService.MarkRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markReadNotificationRequest** | [**MarkReadNotificationRequest**](MarkReadNotificationRequest.md) |  | 

### Return type

[**MarkReadNotificationReply**](MarkReadNotificationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

