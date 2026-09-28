# \NotificationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountUnread**](NotificationService.md#CountUnread) | **Post** /v1/notify/notification/count-unread | 
[**List**](NotificationService.md#List) | **Post** /v1/notify/notification/list | 
[**MarkRead**](NotificationService.md#MarkRead) | **Post** /v1/notify/notification/mark-read | 



## CountUnread

> CountUnreadNotificationsResp CountUnread(ctx).Body(body).Execute()





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
	// response from `CountUnread`: CountUnreadNotificationsResp
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

[**CountUnreadNotificationsResp**](CountUnreadNotificationsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListNotificationsResp List(ctx).ListNotificationsReq(listNotificationsReq).Execute()





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
	resp, r, err := apiClient.NotificationService.List(context.Background()).ListNotificationsReq(listNotificationsReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationService.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListNotificationsResp
	fmt.Fprintf(os.Stdout, "Response from `NotificationService.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


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


## MarkRead

> MarkReadNotificationResp MarkRead(ctx).MarkReadNotificationReq(markReadNotificationReq).Execute()





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
	resp, r, err := apiClient.NotificationService.MarkRead(context.Background()).MarkReadNotificationReq(markReadNotificationReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationService.MarkRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkRead`: MarkReadNotificationResp
	fmt.Fprintf(os.Stdout, "Response from `NotificationService.MarkRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkReadRequest struct via the builder pattern


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

