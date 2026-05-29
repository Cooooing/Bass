# \NotificationServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**notification_service_count_unread**](NotificationServiceApi.md#notification_service_count_unread) | **POST** /v1/notify/notification/count-unread | 
[**notification_service_list**](NotificationServiceApi.md#notification_service_list) | **POST** /v1/notify/notification/list | 
[**notification_service_mark_read**](NotificationServiceApi.md#notification_service_mark_read) | **POST** /v1/notify/notification/mark-read | 



## notification_service_count_unread

> models::CountUnreadNotificationsReply notification_service_count_unread(body)


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


## notification_service_list

> models::ListNotificationsReply notification_service_list(list_notifications_request)


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


## notification_service_mark_read

> models::MarkReadNotificationReply notification_service_mark_read(mark_read_notification_request)


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

