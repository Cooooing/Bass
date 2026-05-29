# \NotificationServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationServiceCountUnread**](NotificationServiceAPI.md#NotificationServiceCountUnread) | **Post** /v1/notify/notification/count-unread | 
[**NotificationServiceList**](NotificationServiceAPI.md#NotificationServiceList) | **Post** /v1/notify/notification/list | 
[**NotificationServiceMarkRead**](NotificationServiceAPI.md#NotificationServiceMarkRead) | **Post** /v1/notify/notification/mark-read | 



## NotificationServiceCountUnread

> CountUnreadNotificationsReply NotificationServiceCountUnread(ctx).Body(body).Execute()



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
	resp, r, err := apiClient.NotificationServiceAPI.NotificationServiceCountUnread(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceAPI.NotificationServiceCountUnread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationServiceCountUnread`: CountUnreadNotificationsReply
	fmt.Fprintf(os.Stdout, "Response from `NotificationServiceAPI.NotificationServiceCountUnread`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceCountUnreadRequest struct via the builder pattern


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


## NotificationServiceList

> ListNotificationsReply NotificationServiceList(ctx).ListNotificationsRequest(listNotificationsRequest).Execute()



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
	resp, r, err := apiClient.NotificationServiceAPI.NotificationServiceList(context.Background()).ListNotificationsRequest(listNotificationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceAPI.NotificationServiceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationServiceList`: ListNotificationsReply
	fmt.Fprintf(os.Stdout, "Response from `NotificationServiceAPI.NotificationServiceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceListRequest struct via the builder pattern


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


## NotificationServiceMarkRead

> MarkReadNotificationReply NotificationServiceMarkRead(ctx).MarkReadNotificationRequest(markReadNotificationRequest).Execute()



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
	resp, r, err := apiClient.NotificationServiceAPI.NotificationServiceMarkRead(context.Background()).MarkReadNotificationRequest(markReadNotificationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationServiceAPI.NotificationServiceMarkRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationServiceMarkRead`: MarkReadNotificationReply
	fmt.Fprintf(os.Stdout, "Response from `NotificationServiceAPI.NotificationServiceMarkRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNotificationServiceMarkReadRequest struct via the builder pattern


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

