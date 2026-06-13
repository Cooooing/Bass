# NotificationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**countUnread**](NotificationService.md#countunread) | **POST** /v1/notify/notification/count-unread |  |
| [**list**](NotificationService.md#list) | **POST** /v1/notify/notification/list |  |
| [**markRead**](NotificationService.md#markread) | **POST** /v1/notify/notification/mark-read |  |



## countUnread

> CountUnreadNotificationsReply countUnread(body)



统计未读通知数量。

### Example

```ts
import {
  Configuration,
  NotificationService,
} from '@bass/bbs-sdk-fetch';
import type { CountUnreadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new NotificationService();

  const body = {
    // object
    body: Object,
  } satisfies CountUnreadRequest;

  try {
    const data = await api.countUnread(body);
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

[**CountUnreadNotificationsReply**](CountUnreadNotificationsReply.md)

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


## list

> ListNotificationsReply list(listNotificationsRequest)



分页查询通知列表。

### Example

```ts
import {
  Configuration,
  NotificationService,
} from '@bass/bbs-sdk-fetch';
import type { ListRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new NotificationService();

  const body = {
    // ListNotificationsRequest
    listNotificationsRequest: ...,
  } satisfies ListRequest;

  try {
    const data = await api.list(body);
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
| **listNotificationsRequest** | [ListNotificationsRequest](ListNotificationsRequest.md) |  | |

### Return type

[**ListNotificationsReply**](ListNotificationsReply.md)

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


## markRead

> MarkReadNotificationReply markRead(markReadNotificationRequest)



标记通知为已读。

### Example

```ts
import {
  Configuration,
  NotificationService,
} from '@bass/bbs-sdk-fetch';
import type { MarkReadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new NotificationService();

  const body = {
    // MarkReadNotificationRequest
    markReadNotificationRequest: ...,
  } satisfies MarkReadRequest;

  try {
    const data = await api.markRead(body);
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
| **markReadNotificationRequest** | [MarkReadNotificationRequest](MarkReadNotificationRequest.md) |  | |

### Return type

[**MarkReadNotificationReply**](MarkReadNotificationReply.md)

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

