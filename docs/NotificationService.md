# NotificationService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**countUnreadNotifications**](#countunreadnotifications) | **POST** /v1/notify/notification/count-unread | |
|[**listNotifications**](#listnotifications) | **POST** /v1/notify/notification/list | |
|[**markReadNotification**](#markreadnotification) | **POST** /v1/notify/notification/mark-read | |

# **countUnreadNotifications**
> CountUnreadNotificationsResp countUnreadNotifications(body)

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

const { status, data } = await apiInstance.countUnreadNotifications(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**CountUnreadNotificationsResp**

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

# **listNotifications**
> ListNotificationsResp listNotifications(listNotificationsReq)

分页查询通知列表。

### Example

```typescript
import {
    NotificationService,
    Configuration,
    ListNotificationsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new NotificationService(configuration);

let listNotificationsReq: ListNotificationsReq; //

const { status, data } = await apiInstance.listNotifications(
    listNotificationsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listNotificationsReq** | **ListNotificationsReq**|  | |


### Return type

**ListNotificationsResp**

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

# **markReadNotification**
> MarkReadNotificationResp markReadNotification(markReadNotificationReq)

标记通知为已读。

### Example

```typescript
import {
    NotificationService,
    Configuration,
    MarkReadNotificationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new NotificationService(configuration);

let markReadNotificationReq: MarkReadNotificationReq; //

const { status, data } = await apiInstance.markReadNotification(
    markReadNotificationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **markReadNotificationReq** | **MarkReadNotificationReq**|  | |


### Return type

**MarkReadNotificationResp**

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

