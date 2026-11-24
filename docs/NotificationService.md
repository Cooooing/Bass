# NotificationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**countUnreadNotifications**](NotificationService.md#countunreadnotifications) | **POST** /v1/notify/notification/count-unread |  |
| [**listNotifications**](NotificationService.md#listnotifications) | **POST** /v1/notify/notification/list |  |
| [**markReadNotification**](NotificationService.md#markreadnotification) | **POST** /v1/notify/notification/mark-read |  |



## countUnreadNotifications

> CountUnreadNotificationsResp countUnreadNotifications(body)



统计未读通知数量。

### Example

```ts
import {
  Configuration,
  NotificationService,
} from '@bass/bbs-sdk-fetch';
import type { CountUnreadNotificationsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new NotificationService();

  const body = {
    // object
    body: Object,
  } satisfies CountUnreadNotificationsRequest;

  try {
    const data = await api.countUnreadNotifications(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | `object` |  | |

### Return type

[**CountUnreadNotificationsResp**](CountUnreadNotificationsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## listNotifications

> ListNotificationsResp listNotifications(listNotificationsReq)



分页查询通知列表。

### Example

```ts
import {
  Configuration,
  NotificationService,
} from '@bass/bbs-sdk-fetch';
import type { ListNotificationsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new NotificationService();

  const body = {
    // ListNotificationsReq
    listNotificationsReq: ...,
  } satisfies ListNotificationsRequest;

  try {
    const data = await api.listNotifications(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **listNotificationsReq** | [ListNotificationsReq](ListNotificationsReq.md) |  | |

### Return type

[**ListNotificationsResp**](ListNotificationsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## markReadNotification

> MarkReadNotificationResp markReadNotification(markReadNotificationReq)



标记通知为已读。

### Example

```ts
import {
  Configuration,
  NotificationService,
} from '@bass/bbs-sdk-fetch';
import type { MarkReadNotificationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new NotificationService();

  const body = {
    // MarkReadNotificationReq
    markReadNotificationReq: ...,
  } satisfies MarkReadNotificationRequest;

  try {
    const data = await api.markReadNotification(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **markReadNotificationReq** | [MarkReadNotificationReq](MarkReadNotificationReq.md) |  | |

### Return type

[**MarkReadNotificationResp**](MarkReadNotificationResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

