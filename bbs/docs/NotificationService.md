# \NotificationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**count_unread**](NotificationService.md#count_unread) | **POST** /v1/notify/notification/count-unread | 
[**list**](NotificationService.md#list) | **POST** /v1/notify/notification/list | 
[**mark_read**](NotificationService.md#mark_read) | **POST** /v1/notify/notification/mark-read | 



## count_unread

> models::CountUnreadNotificationsResp count_unread(body)


统计未读通知数量。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::CountUnreadNotificationsResp**](CountUnreadNotifications_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list

> models::ListNotificationsResp list(list_notifications_req)


分页查询通知列表。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_notifications_req** | [**ListNotificationsReq**](ListNotificationsReq.md) |  | [required] |

### Return type

[**models::ListNotificationsResp**](ListNotifications_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## mark_read

> models::MarkReadNotificationResp mark_read(mark_read_notification_req)


标记通知为已读。

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**mark_read_notification_req** | [**MarkReadNotificationReq**](MarkReadNotificationReq.md) |  | [required] |

### Return type

[**models::MarkReadNotificationResp**](MarkReadNotification_Resp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

