# \NotificationService

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**count_unread**](NotificationService.md#count_unread) | **POST** /v1/notify/notification/count-unread | 
[**list**](NotificationService.md#list) | **POST** /v1/notify/notification/list | 
[**mark_read**](NotificationService.md#mark_read) | **POST** /v1/notify/notification/mark-read | 



## count_unread

> models::CountUnreadNotificationsReply count_unread(body)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**body** | **serde_json::Value** |  | [required] |

### Return type

[**models::CountUnreadNotificationsReply**](CountUnreadNotifications_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list

> models::ListNotificationsReply list(list_notifications_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**list_notifications_request** | [**ListNotificationsRequest**](ListNotificationsRequest.md) |  | [required] |

### Return type

[**models::ListNotificationsReply**](ListNotifications_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## mark_read

> models::MarkReadNotificationReply mark_read(mark_read_notification_request)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**mark_read_notification_request** | [**MarkReadNotificationRequest**](MarkReadNotificationRequest.md) |  | [required] |

### Return type

[**models::MarkReadNotificationReply**](MarkReadNotification_Reply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

