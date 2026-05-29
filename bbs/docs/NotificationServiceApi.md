# NotificationServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**notificationServiceCountUnread**](NotificationServiceApi.md#notificationservicecountunread) | **POST** /v1/notify/notification/count-unread |  |
| [**notificationServiceList**](NotificationServiceApi.md#notificationservicelist) | **POST** /v1/notify/notification/list |  |
| [**notificationServiceMarkRead**](NotificationServiceApi.md#notificationservicemarkread) | **POST** /v1/notify/notification/mark-read |  |



## notificationServiceCountUnread

> CountUnreadNotificationsReply notificationServiceCountUnread(body)



### Example

```ts
import {
  Configuration,
  NotificationServiceApi,
} from '@bass/bbs-sdk';
import type { NotificationServiceCountUnreadRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new NotificationServiceApi();

  const body = {
    // object
    body: Object,
  } satisfies NotificationServiceCountUnreadRequest;

  try {
    const data = await api.notificationServiceCountUnread(body);
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


## notificationServiceList

> ListNotificationsReply notificationServiceList(listNotificationsRequest)



### Example

```ts
import {
  Configuration,
  NotificationServiceApi,
} from '@bass/bbs-sdk';
import type { NotificationServiceListRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new NotificationServiceApi();

  const body = {
    // ListNotificationsRequest
    listNotificationsRequest: ...,
  } satisfies NotificationServiceListRequest;

  try {
    const data = await api.notificationServiceList(body);
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


## notificationServiceMarkRead

> MarkReadNotificationReply notificationServiceMarkRead(markReadNotificationRequest)



### Example

```ts
import {
  Configuration,
  NotificationServiceApi,
} from '@bass/bbs-sdk';
import type { NotificationServiceMarkReadRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new NotificationServiceApi();

  const body = {
    // MarkReadNotificationRequest
    markReadNotificationRequest: ...,
  } satisfies NotificationServiceMarkReadRequest;

  try {
    const data = await api.notificationServiceMarkRead(body);
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

