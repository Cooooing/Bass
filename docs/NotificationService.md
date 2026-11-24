# NotificationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**countUnreadNotifications**](NotificationService.md#countUnreadNotifications) | **POST** /v1/notify/notification/count-unread |  |
| [**countUnreadNotificationsWithHttpInfo**](NotificationService.md#countUnreadNotificationsWithHttpInfo) | **POST** /v1/notify/notification/count-unread |  |
| [**listNotifications**](NotificationService.md#listNotifications) | **POST** /v1/notify/notification/list |  |
| [**listNotificationsWithHttpInfo**](NotificationService.md#listNotificationsWithHttpInfo) | **POST** /v1/notify/notification/list |  |
| [**markReadNotification**](NotificationService.md#markReadNotification) | **POST** /v1/notify/notification/mark-read |  |
| [**markReadNotificationWithHttpInfo**](NotificationService.md#markReadNotificationWithHttpInfo) | **POST** /v1/notify/notification/mark-read |  |



## countUnreadNotifications

> CountUnreadNotificationsResp countUnreadNotifications(countUnreadNotificationsRequest)



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
            APIcountUnreadNotificationsRequest request = APIcountUnreadNotificationsRequest.newBuilder()
                .body(body)
                .build();
            CountUnreadNotificationsResp result = apiInstance.countUnreadNotifications(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#countUnreadNotifications");
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
| countUnreadNotificationsRequest | [**APIcountUnreadNotificationsRequest**](NotificationService.md#APIcountUnreadNotificationsRequest)|-|-|

### Return type

[**CountUnreadNotificationsResp**](CountUnreadNotificationsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## countUnreadNotificationsWithHttpInfo

> ApiResponse<CountUnreadNotificationsResp> countUnreadNotificationsWithHttpInfo(countUnreadNotificationsRequest)



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
            APIcountUnreadNotificationsRequest request = APIcountUnreadNotificationsRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<CountUnreadNotificationsResp> response = apiInstance.countUnreadNotificationsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#countUnreadNotifications");
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
| countUnreadNotificationsRequest | [**APIcountUnreadNotificationsRequest**](NotificationService.md#APIcountUnreadNotificationsRequest)|-|-|

### Return type

ApiResponse<[**CountUnreadNotificationsResp**](CountUnreadNotificationsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcountUnreadNotificationsRequest"></a>
## APIcountUnreadNotificationsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## listNotifications

> ListNotificationsResp listNotifications(listNotificationsRequest)



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
        ListNotificationsReq listNotificationsReq = new ListNotificationsReq(); // ListNotificationsReq | 
        try {
            APIlistNotificationsRequest request = APIlistNotificationsRequest.newBuilder()
                .listNotificationsReq(listNotificationsReq)
                .build();
            ListNotificationsResp result = apiInstance.listNotifications(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#listNotifications");
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
| listNotificationsRequest | [**APIlistNotificationsRequest**](NotificationService.md#APIlistNotificationsRequest)|-|-|

### Return type

[**ListNotificationsResp**](ListNotificationsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listNotificationsWithHttpInfo

> ApiResponse<ListNotificationsResp> listNotificationsWithHttpInfo(listNotificationsRequest)



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
        ListNotificationsReq listNotificationsReq = new ListNotificationsReq(); // ListNotificationsReq | 
        try {
            APIlistNotificationsRequest request = APIlistNotificationsRequest.newBuilder()
                .listNotificationsReq(listNotificationsReq)
                .build();
            ApiResponse<ListNotificationsResp> response = apiInstance.listNotificationsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#listNotifications");
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
| listNotificationsRequest | [**APIlistNotificationsRequest**](NotificationService.md#APIlistNotificationsRequest)|-|-|

### Return type

ApiResponse<[**ListNotificationsResp**](ListNotificationsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistNotificationsRequest"></a>
## APIlistNotificationsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listNotificationsReq** | [**ListNotificationsReq**](ListNotificationsReq.md) |  | |



## markReadNotification

> MarkReadNotificationResp markReadNotification(markReadNotificationRequest)



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
        MarkReadNotificationReq markReadNotificationReq = new MarkReadNotificationReq(); // MarkReadNotificationReq | 
        try {
            APImarkReadNotificationRequest request = APImarkReadNotificationRequest.newBuilder()
                .markReadNotificationReq(markReadNotificationReq)
                .build();
            MarkReadNotificationResp result = apiInstance.markReadNotification(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#markReadNotification");
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
| markReadNotificationRequest | [**APImarkReadNotificationRequest**](NotificationService.md#APImarkReadNotificationRequest)|-|-|

### Return type

[**MarkReadNotificationResp**](MarkReadNotificationResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## markReadNotificationWithHttpInfo

> ApiResponse<MarkReadNotificationResp> markReadNotificationWithHttpInfo(markReadNotificationRequest)



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
        MarkReadNotificationReq markReadNotificationReq = new MarkReadNotificationReq(); // MarkReadNotificationReq | 
        try {
            APImarkReadNotificationRequest request = APImarkReadNotificationRequest.newBuilder()
                .markReadNotificationReq(markReadNotificationReq)
                .build();
            ApiResponse<MarkReadNotificationResp> response = apiInstance.markReadNotificationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling NotificationService#markReadNotification");
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
| markReadNotificationRequest | [**APImarkReadNotificationRequest**](NotificationService.md#APImarkReadNotificationRequest)|-|-|

### Return type

ApiResponse<[**MarkReadNotificationResp**](MarkReadNotificationResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APImarkReadNotificationRequest"></a>
## APImarkReadNotificationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **markReadNotificationReq** | [**MarkReadNotificationReq**](MarkReadNotificationReq.md) |  | |


