# NotificationServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**notificationServiceCountUnread**](NotificationServiceApi.md#notificationServiceCountUnread) | **POST** /v1/notify/notification/count-unread |  |
| [**notificationServiceCountUnreadWithHttpInfo**](NotificationServiceApi.md#notificationServiceCountUnreadWithHttpInfo) | **POST** /v1/notify/notification/count-unread |  |
| [**notificationServiceList**](NotificationServiceApi.md#notificationServiceList) | **POST** /v1/notify/notification/list |  |
| [**notificationServiceListWithHttpInfo**](NotificationServiceApi.md#notificationServiceListWithHttpInfo) | **POST** /v1/notify/notification/list |  |
| [**notificationServiceMarkRead**](NotificationServiceApi.md#notificationServiceMarkRead) | **POST** /v1/notify/notification/mark-read |  |
| [**notificationServiceMarkReadWithHttpInfo**](NotificationServiceApi.md#notificationServiceMarkReadWithHttpInfo) | **POST** /v1/notify/notification/mark-read |  |



## notificationServiceCountUnread

> CountUnreadNotificationsReply notificationServiceCountUnread(notificationServiceCountUnreadRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationServiceApi;
import com.bass.bbs.api.NotificationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationServiceApi apiInstance = new NotificationServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APInotificationServiceCountUnreadRequest request = APInotificationServiceCountUnreadRequest.newBuilder()
                .body(body)
                .build();
            CountUnreadNotificationsReply result = apiInstance.notificationServiceCountUnread(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationServiceApi#notificationServiceCountUnread");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| notificationServiceCountUnreadRequest | [**APInotificationServiceCountUnreadRequest**](NotificationServiceApi.md#APInotificationServiceCountUnreadRequest)|-|-|

### Return type

[**CountUnreadNotificationsReply**](CountUnreadNotificationsReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## notificationServiceCountUnreadWithHttpInfo

> ApiResponse<CountUnreadNotificationsReply> notificationServiceCountUnreadWithHttpInfo(notificationServiceCountUnreadRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationServiceApi;
import com.bass.bbs.api.NotificationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationServiceApi apiInstance = new NotificationServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APInotificationServiceCountUnreadRequest request = APInotificationServiceCountUnreadRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<CountUnreadNotificationsReply> response = apiInstance.notificationServiceCountUnreadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationServiceApi#notificationServiceCountUnread");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| notificationServiceCountUnreadRequest | [**APInotificationServiceCountUnreadRequest**](NotificationServiceApi.md#APInotificationServiceCountUnreadRequest)|-|-|

### Return type

ApiResponse<[**CountUnreadNotificationsReply**](CountUnreadNotificationsReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APInotificationServiceCountUnreadRequest"></a>
## APInotificationServiceCountUnreadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## notificationServiceList

> ListNotificationsReply notificationServiceList(notificationServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationServiceApi;
import com.bass.bbs.api.NotificationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationServiceApi apiInstance = new NotificationServiceApi(defaultClient);
        ListNotificationsRequest listNotificationsRequest = new ListNotificationsRequest(); // ListNotificationsRequest | 
        try {
            APInotificationServiceListRequest request = APInotificationServiceListRequest.newBuilder()
                .listNotificationsRequest(listNotificationsRequest)
                .build();
            ListNotificationsReply result = apiInstance.notificationServiceList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationServiceApi#notificationServiceList");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| notificationServiceListRequest | [**APInotificationServiceListRequest**](NotificationServiceApi.md#APInotificationServiceListRequest)|-|-|

### Return type

[**ListNotificationsReply**](ListNotificationsReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## notificationServiceListWithHttpInfo

> ApiResponse<ListNotificationsReply> notificationServiceListWithHttpInfo(notificationServiceListRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationServiceApi;
import com.bass.bbs.api.NotificationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationServiceApi apiInstance = new NotificationServiceApi(defaultClient);
        ListNotificationsRequest listNotificationsRequest = new ListNotificationsRequest(); // ListNotificationsRequest | 
        try {
            APInotificationServiceListRequest request = APInotificationServiceListRequest.newBuilder()
                .listNotificationsRequest(listNotificationsRequest)
                .build();
            ApiResponse<ListNotificationsReply> response = apiInstance.notificationServiceListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationServiceApi#notificationServiceList");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| notificationServiceListRequest | [**APInotificationServiceListRequest**](NotificationServiceApi.md#APInotificationServiceListRequest)|-|-|

### Return type

ApiResponse<[**ListNotificationsReply**](ListNotificationsReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APInotificationServiceListRequest"></a>
## APInotificationServiceListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listNotificationsRequest** | [**ListNotificationsRequest**](ListNotificationsRequest.md) |  | |



## notificationServiceMarkRead

> MarkReadNotificationReply notificationServiceMarkRead(notificationServiceMarkReadRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationServiceApi;
import com.bass.bbs.api.NotificationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationServiceApi apiInstance = new NotificationServiceApi(defaultClient);
        MarkReadNotificationRequest markReadNotificationRequest = new MarkReadNotificationRequest(); // MarkReadNotificationRequest | 
        try {
            APInotificationServiceMarkReadRequest request = APInotificationServiceMarkReadRequest.newBuilder()
                .markReadNotificationRequest(markReadNotificationRequest)
                .build();
            MarkReadNotificationReply result = apiInstance.notificationServiceMarkRead(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationServiceApi#notificationServiceMarkRead");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| notificationServiceMarkReadRequest | [**APInotificationServiceMarkReadRequest**](NotificationServiceApi.md#APInotificationServiceMarkReadRequest)|-|-|

### Return type

[**MarkReadNotificationReply**](MarkReadNotificationReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## notificationServiceMarkReadWithHttpInfo

> ApiResponse<MarkReadNotificationReply> notificationServiceMarkReadWithHttpInfo(notificationServiceMarkReadRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationServiceApi;
import com.bass.bbs.api.NotificationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationServiceApi apiInstance = new NotificationServiceApi(defaultClient);
        MarkReadNotificationRequest markReadNotificationRequest = new MarkReadNotificationRequest(); // MarkReadNotificationRequest | 
        try {
            APInotificationServiceMarkReadRequest request = APInotificationServiceMarkReadRequest.newBuilder()
                .markReadNotificationRequest(markReadNotificationRequest)
                .build();
            ApiResponse<MarkReadNotificationReply> response = apiInstance.notificationServiceMarkReadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationServiceApi#notificationServiceMarkRead");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| notificationServiceMarkReadRequest | [**APInotificationServiceMarkReadRequest**](NotificationServiceApi.md#APInotificationServiceMarkReadRequest)|-|-|

### Return type

ApiResponse<[**MarkReadNotificationReply**](MarkReadNotificationReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APInotificationServiceMarkReadRequest"></a>
## APInotificationServiceMarkReadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **markReadNotificationRequest** | [**MarkReadNotificationRequest**](MarkReadNotificationRequest.md) |  | |


