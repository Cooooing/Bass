# \NotificationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountUnreadNotifications**](NotificationService.md#CountUnreadNotifications) | **Post** /v1/notify/notification/count-unread | 
[**ListNotifications**](NotificationService.md#ListNotifications) | **Post** /v1/notify/notification/list | 
[**MarkReadNotification**](NotificationService.md#MarkReadNotification) | **Post** /v1/notify/notification/mark-read | 



## CountUnreadNotifications

> CountUnreadNotificationsResp CountUnreadNotifications(ctx).Body(body).Execute()





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
	resp, r, err := apiClient.NotificationService.CountUnreadNotifications(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationService.CountUnreadNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountUnreadNotifications`: CountUnreadNotificationsResp
	fmt.Fprintf(os.Stdout, "Response from `NotificationService.CountUnreadNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountUnreadNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**CountUnreadNotificationsResp**](CountUnreadNotificationsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotifications

> ListNotificationsResp ListNotifications(ctx).ListNotificationsReq(listNotificationsReq).Execute()





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
	listNotificationsReq := *openapiclient.NewListNotificationsReq() // ListNotificationsReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationService.ListNotifications(context.Background()).ListNotificationsReq(listNotificationsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationService.ListNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotifications`: ListNotificationsResp
	fmt.Fprintf(os.Stdout, "Response from `NotificationService.ListNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listNotificationsReq** | [**ListNotificationsReq**](ListNotificationsReq.md) |  | 

### Return type

[**ListNotificationsResp**](ListNotificationsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkReadNotification

> MarkReadNotificationResp MarkReadNotification(ctx).MarkReadNotificationReq(markReadNotificationReq).Execute()





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
	markReadNotificationReq := *openapiclient.NewMarkReadNotificationReq([]string{"Ids_example"}) // MarkReadNotificationReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationService.MarkReadNotification(context.Background()).MarkReadNotificationReq(markReadNotificationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationService.MarkReadNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkReadNotification`: MarkReadNotificationResp
	fmt.Fprintf(os.Stdout, "Response from `NotificationService.MarkReadNotification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkReadNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markReadNotificationReq** | [**MarkReadNotificationReq**](MarkReadNotificationReq.md) |  | 

### Return type

[**MarkReadNotificationResp**](MarkReadNotificationResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

