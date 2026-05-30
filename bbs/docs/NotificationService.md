# NotificationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**callList**](NotificationService.md#callList) | **POST** /v1/notify/notification/list |  |
| [**callListWithHttpInfo**](NotificationService.md#callListWithHttpInfo) | **POST** /v1/notify/notification/list |  |
| [**countUnread**](NotificationService.md#countUnread) | **POST** /v1/notify/notification/count-unread |  |
| [**countUnreadWithHttpInfo**](NotificationService.md#countUnreadWithHttpInfo) | **POST** /v1/notify/notification/count-unread |  |
| [**markRead**](NotificationService.md#markRead) | **POST** /v1/notify/notification/mark-read |  |
| [**markReadWithHttpInfo**](NotificationService.md#markReadWithHttpInfo) | **POST** /v1/notify/notification/mark-read |  |



## callList

> ListNotificationsReply callList(callListRequest)



分页查询通知列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationService;
import com.bass.bbs.api.NotificationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationService apiInstance = new NotificationService(defaultClient);
        ListNotificationsRequest listNotificationsRequest = new ListNotificationsRequest(); // ListNotificationsRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listNotificationsRequest(listNotificationsRequest)
                .build();
            ListNotificationsReply result = apiInstance.callList(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#callList");
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
| callListRequest | [**APIcallListRequest**](NotificationService.md#APIcallListRequest)|-|-|

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

## callListWithHttpInfo

> ApiResponse<ListNotificationsReply> callListWithHttpInfo(callListRequest)



分页查询通知列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationService;
import com.bass.bbs.api.NotificationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationService apiInstance = new NotificationService(defaultClient);
        ListNotificationsRequest listNotificationsRequest = new ListNotificationsRequest(); // ListNotificationsRequest | 
        try {
            APIcallListRequest request = APIcallListRequest.newBuilder()
                .listNotificationsRequest(listNotificationsRequest)
                .build();
            ApiResponse<ListNotificationsReply> response = apiInstance.callListWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#callList");
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
| callListRequest | [**APIcallListRequest**](NotificationService.md#APIcallListRequest)|-|-|

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


<a id="APIcallListRequest"></a>
## APIcallListRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listNotificationsRequest** | [**ListNotificationsRequest**](ListNotificationsRequest.md) |  | |



## countUnread

> CountUnreadNotificationsReply countUnread(countUnreadRequest)



统计未读通知数量。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationService;
import com.bass.bbs.api.NotificationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationService apiInstance = new NotificationService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcountUnreadRequest request = APIcountUnreadRequest.newBuilder()
                .body(body)
                .build();
            CountUnreadNotificationsReply result = apiInstance.countUnread(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#countUnread");
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
| countUnreadRequest | [**APIcountUnreadRequest**](NotificationService.md#APIcountUnreadRequest)|-|-|

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

## countUnreadWithHttpInfo

> ApiResponse<CountUnreadNotificationsReply> countUnreadWithHttpInfo(countUnreadRequest)



统计未读通知数量。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationService;
import com.bass.bbs.api.NotificationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationService apiInstance = new NotificationService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcountUnreadRequest request = APIcountUnreadRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<CountUnreadNotificationsReply> response = apiInstance.countUnreadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#countUnread");
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
| countUnreadRequest | [**APIcountUnreadRequest**](NotificationService.md#APIcountUnreadRequest)|-|-|

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


<a id="APIcountUnreadRequest"></a>
## APIcountUnreadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## markRead

> MarkReadNotificationReply markRead(markReadRequest)



标记通知为已读。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationService;
import com.bass.bbs.api.NotificationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationService apiInstance = new NotificationService(defaultClient);
        MarkReadNotificationRequest markReadNotificationRequest = new MarkReadNotificationRequest(); // MarkReadNotificationRequest | 
        try {
            APImarkReadRequest request = APImarkReadRequest.newBuilder()
                .markReadNotificationRequest(markReadNotificationRequest)
                .build();
            MarkReadNotificationReply result = apiInstance.markRead(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#markRead");
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
| markReadRequest | [**APImarkReadRequest**](NotificationService.md#APImarkReadRequest)|-|-|

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

## markReadWithHttpInfo

> ApiResponse<MarkReadNotificationReply> markReadWithHttpInfo(markReadRequest)



标记通知为已读。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.NotificationService;
import com.bass.bbs.api.NotificationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        NotificationService apiInstance = new NotificationService(defaultClient);
        MarkReadNotificationRequest markReadNotificationRequest = new MarkReadNotificationRequest(); // MarkReadNotificationRequest | 
        try {
            APImarkReadRequest request = APImarkReadRequest.newBuilder()
                .markReadNotificationRequest(markReadNotificationRequest)
                .build();
            ApiResponse<MarkReadNotificationReply> response = apiInstance.markReadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#markRead");
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
| markReadRequest | [**APImarkReadRequest**](NotificationService.md#APImarkReadRequest)|-|-|

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


<a id="APImarkReadRequest"></a>
## APImarkReadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **markReadNotificationRequest** | [**MarkReadNotificationRequest**](MarkReadNotificationRequest.md) |  | |


