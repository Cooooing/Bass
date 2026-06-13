# NotificationService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**countUnread**](#countunread) | **POST** /v1/notify/notification/count-unread | |
|[**list**](#list) | **POST** /v1/notify/notification/list | |
|[**markRead**](#markread) | **POST** /v1/notify/notification/mark-read | |

# **countUnread**
> CountUnreadNotificationsReply countUnread(body)

统计未读通知数量。

### Example

```typescript
import {
    NotificationService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new NotificationService(configuration);

let body: object; //

const { status, data } = await apiInstance.countUnread(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**CountUnreadNotificationsReply**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list**
> ListNotificationsReply list(listNotificationsRequest)

分页查询通知列表。

### Example

```typescript
import {
    NotificationService,
    Configuration,
    ListNotificationsRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new NotificationService(configuration);

let listNotificationsRequest: ListNotificationsRequest; //

const { status, data } = await apiInstance.list(
    listNotificationsRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listNotificationsRequest** | **ListNotificationsRequest**|  | |


### Return type

**ListNotificationsReply**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **markRead**
> MarkReadNotificationReply markRead(markReadNotificationRequest)

标记通知为已读。

### Example

```typescript
import {
    NotificationService,
    Configuration,
    MarkReadNotificationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new NotificationService(configuration);

let markReadNotificationRequest: MarkReadNotificationRequest; //

const { status, data } = await apiInstance.markRead(
    markReadNotificationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **markReadNotificationRequest** | **MarkReadNotificationRequest**|  | |


### Return type

**MarkReadNotificationReply**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

