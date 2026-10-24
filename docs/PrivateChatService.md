# PrivateChatService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**countPrivateChatUnread**](PrivateChatService.md#countPrivateChatUnread) | **POST** /v1/im/private-chat/unread-count |  |
| [**countPrivateChatUnreadWithHttpInfo**](PrivateChatService.md#countPrivateChatUnreadWithHttpInfo) | **POST** /v1/im/private-chat/unread-count |  |
| [**listPrivateChatMessages**](PrivateChatService.md#listPrivateChatMessages) | **POST** /v1/im/private-chat/messages/list |  |
| [**listPrivateChatMessagesWithHttpInfo**](PrivateChatService.md#listPrivateChatMessagesWithHttpInfo) | **POST** /v1/im/private-chat/messages/list |  |
| [**listPrivateChatSessions**](PrivateChatService.md#listPrivateChatSessions) | **GET** /v1/im/private-chat/sessions |  |
| [**listPrivateChatSessionsWithHttpInfo**](PrivateChatService.md#listPrivateChatSessionsWithHttpInfo) | **GET** /v1/im/private-chat/sessions |  |
| [**markPrivateChatSessionRead**](PrivateChatService.md#markPrivateChatSessionRead) | **POST** /v1/im/private-chat/read |  |
| [**markPrivateChatSessionReadWithHttpInfo**](PrivateChatService.md#markPrivateChatSessionReadWithHttpInfo) | **POST** /v1/im/private-chat/read |  |
| [**revokePrivateChatMessage**](PrivateChatService.md#revokePrivateChatMessage) | **POST** /v1/im/private-chat/messages/revoke |  |
| [**revokePrivateChatMessageWithHttpInfo**](PrivateChatService.md#revokePrivateChatMessageWithHttpInfo) | **POST** /v1/im/private-chat/messages/revoke |  |
| [**sendPrivateChatMessage**](PrivateChatService.md#sendPrivateChatMessage) | **POST** /v1/im/private-chat/messages |  |
| [**sendPrivateChatMessageWithHttpInfo**](PrivateChatService.md#sendPrivateChatMessageWithHttpInfo) | **POST** /v1/im/private-chat/messages |  |
| [**setPrivateChatSessionMuted**](PrivateChatService.md#setPrivateChatSessionMuted) | **POST** /v1/im/private-chat/muted |  |
| [**setPrivateChatSessionMutedWithHttpInfo**](PrivateChatService.md#setPrivateChatSessionMutedWithHttpInfo) | **POST** /v1/im/private-chat/muted |  |
| [**setPrivateChatSessionPinned**](PrivateChatService.md#setPrivateChatSessionPinned) | **POST** /v1/im/private-chat/pinned |  |
| [**setPrivateChatSessionPinnedWithHttpInfo**](PrivateChatService.md#setPrivateChatSessionPinnedWithHttpInfo) | **POST** /v1/im/private-chat/pinned |  |



## countPrivateChatUnread

> CountPrivateChatUnreadResp countPrivateChatUnread(countPrivateChatUnreadRequest)



统计当前账号的私聊未读数。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcountPrivateChatUnreadRequest request = APIcountPrivateChatUnreadRequest.newBuilder()
                .body(body)
                .build();
            CountPrivateChatUnreadResp result = apiInstance.countPrivateChatUnread(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#countPrivateChatUnread");
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
| countPrivateChatUnreadRequest | [**APIcountPrivateChatUnreadRequest**](PrivateChatService.md#APIcountPrivateChatUnreadRequest)|-|-|

### Return type

[**CountPrivateChatUnreadResp**](CountPrivateChatUnreadResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## countPrivateChatUnreadWithHttpInfo

> ApiResponse<CountPrivateChatUnreadResp> countPrivateChatUnreadWithHttpInfo(countPrivateChatUnreadRequest)



统计当前账号的私聊未读数。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcountPrivateChatUnreadRequest request = APIcountPrivateChatUnreadRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<CountPrivateChatUnreadResp> response = apiInstance.countPrivateChatUnreadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#countPrivateChatUnread");
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
| countPrivateChatUnreadRequest | [**APIcountPrivateChatUnreadRequest**](PrivateChatService.md#APIcountPrivateChatUnreadRequest)|-|-|

### Return type

ApiResponse<[**CountPrivateChatUnreadResp**](CountPrivateChatUnreadResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcountPrivateChatUnreadRequest"></a>
## APIcountPrivateChatUnreadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## listPrivateChatMessages

> ListPrivateChatMessagesResp listPrivateChatMessages(listPrivateChatMessagesRequest)



查询私聊消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        ListPrivateChatMessagesReq listPrivateChatMessagesReq = new ListPrivateChatMessagesReq(); // ListPrivateChatMessagesReq | 
        try {
            APIlistPrivateChatMessagesRequest request = APIlistPrivateChatMessagesRequest.newBuilder()
                .listPrivateChatMessagesReq(listPrivateChatMessagesReq)
                .build();
            ListPrivateChatMessagesResp result = apiInstance.listPrivateChatMessages(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#listPrivateChatMessages");
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
| listPrivateChatMessagesRequest | [**APIlistPrivateChatMessagesRequest**](PrivateChatService.md#APIlistPrivateChatMessagesRequest)|-|-|

### Return type

[**ListPrivateChatMessagesResp**](ListPrivateChatMessagesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listPrivateChatMessagesWithHttpInfo

> ApiResponse<ListPrivateChatMessagesResp> listPrivateChatMessagesWithHttpInfo(listPrivateChatMessagesRequest)



查询私聊消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        ListPrivateChatMessagesReq listPrivateChatMessagesReq = new ListPrivateChatMessagesReq(); // ListPrivateChatMessagesReq | 
        try {
            APIlistPrivateChatMessagesRequest request = APIlistPrivateChatMessagesRequest.newBuilder()
                .listPrivateChatMessagesReq(listPrivateChatMessagesReq)
                .build();
            ApiResponse<ListPrivateChatMessagesResp> response = apiInstance.listPrivateChatMessagesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#listPrivateChatMessages");
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
| listPrivateChatMessagesRequest | [**APIlistPrivateChatMessagesRequest**](PrivateChatService.md#APIlistPrivateChatMessagesRequest)|-|-|

### Return type

ApiResponse<[**ListPrivateChatMessagesResp**](ListPrivateChatMessagesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistPrivateChatMessagesRequest"></a>
## APIlistPrivateChatMessagesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listPrivateChatMessagesReq** | [**ListPrivateChatMessagesReq**](ListPrivateChatMessagesReq.md) |  | |



## listPrivateChatSessions

> ListPrivateChatSessionsResp listPrivateChatSessions(listPrivateChatSessionsRequest)



查询当前账号的私聊会话。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        String pagePage = "pagePage_example"; // String | 页码
        String pageSize = "pageSize_example"; // String | 页大小
        try {
            APIlistPrivateChatSessionsRequest request = APIlistPrivateChatSessionsRequest.newBuilder()
                .pagePage(pagePage)
                .pageSize(pageSize)
                .build();
            ListPrivateChatSessionsResp result = apiInstance.listPrivateChatSessions(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#listPrivateChatSessions");
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
| listPrivateChatSessionsRequest | [**APIlistPrivateChatSessionsRequest**](PrivateChatService.md#APIlistPrivateChatSessionsRequest)|-|-|

### Return type

[**ListPrivateChatSessionsResp**](ListPrivateChatSessionsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listPrivateChatSessionsWithHttpInfo

> ApiResponse<ListPrivateChatSessionsResp> listPrivateChatSessionsWithHttpInfo(listPrivateChatSessionsRequest)



查询当前账号的私聊会话。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        String pagePage = "pagePage_example"; // String | 页码
        String pageSize = "pageSize_example"; // String | 页大小
        try {
            APIlistPrivateChatSessionsRequest request = APIlistPrivateChatSessionsRequest.newBuilder()
                .pagePage(pagePage)
                .pageSize(pageSize)
                .build();
            ApiResponse<ListPrivateChatSessionsResp> response = apiInstance.listPrivateChatSessionsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#listPrivateChatSessions");
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
| listPrivateChatSessionsRequest | [**APIlistPrivateChatSessionsRequest**](PrivateChatService.md#APIlistPrivateChatSessionsRequest)|-|-|

### Return type

ApiResponse<[**ListPrivateChatSessionsResp**](ListPrivateChatSessionsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistPrivateChatSessionsRequest"></a>
## APIlistPrivateChatSessionsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **pagePage** | **String** | 页码 | [optional] |
| **pageSize** | **String** | 页大小 | [optional] |



## markPrivateChatSessionRead

> Object markPrivateChatSessionRead(markPrivateChatSessionReadRequest)



标记私聊会话已读。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        MarkPrivateChatSessionReadReq markPrivateChatSessionReadReq = new MarkPrivateChatSessionReadReq(); // MarkPrivateChatSessionReadReq | 
        try {
            APImarkPrivateChatSessionReadRequest request = APImarkPrivateChatSessionReadRequest.newBuilder()
                .markPrivateChatSessionReadReq(markPrivateChatSessionReadReq)
                .build();
            Object result = apiInstance.markPrivateChatSessionRead(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#markPrivateChatSessionRead");
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
| markPrivateChatSessionReadRequest | [**APImarkPrivateChatSessionReadRequest**](PrivateChatService.md#APImarkPrivateChatSessionReadRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## markPrivateChatSessionReadWithHttpInfo

> ApiResponse<Object> markPrivateChatSessionReadWithHttpInfo(markPrivateChatSessionReadRequest)



标记私聊会话已读。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        MarkPrivateChatSessionReadReq markPrivateChatSessionReadReq = new MarkPrivateChatSessionReadReq(); // MarkPrivateChatSessionReadReq | 
        try {
            APImarkPrivateChatSessionReadRequest request = APImarkPrivateChatSessionReadRequest.newBuilder()
                .markPrivateChatSessionReadReq(markPrivateChatSessionReadReq)
                .build();
            ApiResponse<Object> response = apiInstance.markPrivateChatSessionReadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#markPrivateChatSessionRead");
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
| markPrivateChatSessionReadRequest | [**APImarkPrivateChatSessionReadRequest**](PrivateChatService.md#APImarkPrivateChatSessionReadRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APImarkPrivateChatSessionReadRequest"></a>
## APImarkPrivateChatSessionReadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **markPrivateChatSessionReadReq** | [**MarkPrivateChatSessionReadReq**](MarkPrivateChatSessionReadReq.md) |  | |



## revokePrivateChatMessage

> RevokePrivateChatMessageResp revokePrivateChatMessage(revokePrivateChatMessageRequest)



撤回私聊消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        RevokePrivateChatMessageReq revokePrivateChatMessageReq = new RevokePrivateChatMessageReq(); // RevokePrivateChatMessageReq | 
        try {
            APIrevokePrivateChatMessageRequest request = APIrevokePrivateChatMessageRequest.newBuilder()
                .revokePrivateChatMessageReq(revokePrivateChatMessageReq)
                .build();
            RevokePrivateChatMessageResp result = apiInstance.revokePrivateChatMessage(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#revokePrivateChatMessage");
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
| revokePrivateChatMessageRequest | [**APIrevokePrivateChatMessageRequest**](PrivateChatService.md#APIrevokePrivateChatMessageRequest)|-|-|

### Return type

[**RevokePrivateChatMessageResp**](RevokePrivateChatMessageResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## revokePrivateChatMessageWithHttpInfo

> ApiResponse<RevokePrivateChatMessageResp> revokePrivateChatMessageWithHttpInfo(revokePrivateChatMessageRequest)



撤回私聊消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        RevokePrivateChatMessageReq revokePrivateChatMessageReq = new RevokePrivateChatMessageReq(); // RevokePrivateChatMessageReq | 
        try {
            APIrevokePrivateChatMessageRequest request = APIrevokePrivateChatMessageRequest.newBuilder()
                .revokePrivateChatMessageReq(revokePrivateChatMessageReq)
                .build();
            ApiResponse<RevokePrivateChatMessageResp> response = apiInstance.revokePrivateChatMessageWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#revokePrivateChatMessage");
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
| revokePrivateChatMessageRequest | [**APIrevokePrivateChatMessageRequest**](PrivateChatService.md#APIrevokePrivateChatMessageRequest)|-|-|

### Return type

ApiResponse<[**RevokePrivateChatMessageResp**](RevokePrivateChatMessageResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrevokePrivateChatMessageRequest"></a>
## APIrevokePrivateChatMessageRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **revokePrivateChatMessageReq** | [**RevokePrivateChatMessageReq**](RevokePrivateChatMessageReq.md) |  | |



## sendPrivateChatMessage

> SendPrivateChatMessageResp sendPrivateChatMessage(sendPrivateChatMessageRequest)



发送私聊消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        SendPrivateChatMessageReq sendPrivateChatMessageReq = new SendPrivateChatMessageReq(); // SendPrivateChatMessageReq | 
        try {
            APIsendPrivateChatMessageRequest request = APIsendPrivateChatMessageRequest.newBuilder()
                .sendPrivateChatMessageReq(sendPrivateChatMessageReq)
                .build();
            SendPrivateChatMessageResp result = apiInstance.sendPrivateChatMessage(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#sendPrivateChatMessage");
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
| sendPrivateChatMessageRequest | [**APIsendPrivateChatMessageRequest**](PrivateChatService.md#APIsendPrivateChatMessageRequest)|-|-|

### Return type

[**SendPrivateChatMessageResp**](SendPrivateChatMessageResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## sendPrivateChatMessageWithHttpInfo

> ApiResponse<SendPrivateChatMessageResp> sendPrivateChatMessageWithHttpInfo(sendPrivateChatMessageRequest)



发送私聊消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        SendPrivateChatMessageReq sendPrivateChatMessageReq = new SendPrivateChatMessageReq(); // SendPrivateChatMessageReq | 
        try {
            APIsendPrivateChatMessageRequest request = APIsendPrivateChatMessageRequest.newBuilder()
                .sendPrivateChatMessageReq(sendPrivateChatMessageReq)
                .build();
            ApiResponse<SendPrivateChatMessageResp> response = apiInstance.sendPrivateChatMessageWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#sendPrivateChatMessage");
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
| sendPrivateChatMessageRequest | [**APIsendPrivateChatMessageRequest**](PrivateChatService.md#APIsendPrivateChatMessageRequest)|-|-|

### Return type

ApiResponse<[**SendPrivateChatMessageResp**](SendPrivateChatMessageResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsendPrivateChatMessageRequest"></a>
## APIsendPrivateChatMessageRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **sendPrivateChatMessageReq** | [**SendPrivateChatMessageReq**](SendPrivateChatMessageReq.md) |  | |



## setPrivateChatSessionMuted

> Object setPrivateChatSessionMuted(setPrivateChatSessionMutedRequest)



设置私聊会话免打扰。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        SetPrivateChatSessionMutedReq setPrivateChatSessionMutedReq = new SetPrivateChatSessionMutedReq(); // SetPrivateChatSessionMutedReq | 
        try {
            APIsetPrivateChatSessionMutedRequest request = APIsetPrivateChatSessionMutedRequest.newBuilder()
                .setPrivateChatSessionMutedReq(setPrivateChatSessionMutedReq)
                .build();
            Object result = apiInstance.setPrivateChatSessionMuted(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#setPrivateChatSessionMuted");
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
| setPrivateChatSessionMutedRequest | [**APIsetPrivateChatSessionMutedRequest**](PrivateChatService.md#APIsetPrivateChatSessionMutedRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## setPrivateChatSessionMutedWithHttpInfo

> ApiResponse<Object> setPrivateChatSessionMutedWithHttpInfo(setPrivateChatSessionMutedRequest)



设置私聊会话免打扰。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        SetPrivateChatSessionMutedReq setPrivateChatSessionMutedReq = new SetPrivateChatSessionMutedReq(); // SetPrivateChatSessionMutedReq | 
        try {
            APIsetPrivateChatSessionMutedRequest request = APIsetPrivateChatSessionMutedRequest.newBuilder()
                .setPrivateChatSessionMutedReq(setPrivateChatSessionMutedReq)
                .build();
            ApiResponse<Object> response = apiInstance.setPrivateChatSessionMutedWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#setPrivateChatSessionMuted");
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
| setPrivateChatSessionMutedRequest | [**APIsetPrivateChatSessionMutedRequest**](PrivateChatService.md#APIsetPrivateChatSessionMutedRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsetPrivateChatSessionMutedRequest"></a>
## APIsetPrivateChatSessionMutedRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **setPrivateChatSessionMutedReq** | [**SetPrivateChatSessionMutedReq**](SetPrivateChatSessionMutedReq.md) |  | |



## setPrivateChatSessionPinned

> Object setPrivateChatSessionPinned(setPrivateChatSessionPinnedRequest)



设置私聊会话置顶。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        SetPrivateChatSessionPinnedReq setPrivateChatSessionPinnedReq = new SetPrivateChatSessionPinnedReq(); // SetPrivateChatSessionPinnedReq | 
        try {
            APIsetPrivateChatSessionPinnedRequest request = APIsetPrivateChatSessionPinnedRequest.newBuilder()
                .setPrivateChatSessionPinnedReq(setPrivateChatSessionPinnedReq)
                .build();
            Object result = apiInstance.setPrivateChatSessionPinned(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#setPrivateChatSessionPinned");
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
| setPrivateChatSessionPinnedRequest | [**APIsetPrivateChatSessionPinnedRequest**](PrivateChatService.md#APIsetPrivateChatSessionPinnedRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## setPrivateChatSessionPinnedWithHttpInfo

> ApiResponse<Object> setPrivateChatSessionPinnedWithHttpInfo(setPrivateChatSessionPinnedRequest)



设置私聊会话置顶。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivateChatService;
import com.bass.bbs.api.PrivateChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivateChatService apiInstance = new PrivateChatService(defaultClient);
        SetPrivateChatSessionPinnedReq setPrivateChatSessionPinnedReq = new SetPrivateChatSessionPinnedReq(); // SetPrivateChatSessionPinnedReq | 
        try {
            APIsetPrivateChatSessionPinnedRequest request = APIsetPrivateChatSessionPinnedRequest.newBuilder()
                .setPrivateChatSessionPinnedReq(setPrivateChatSessionPinnedReq)
                .build();
            ApiResponse<Object> response = apiInstance.setPrivateChatSessionPinnedWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivateChatService#setPrivateChatSessionPinned");
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
| setPrivateChatSessionPinnedRequest | [**APIsetPrivateChatSessionPinnedRequest**](PrivateChatService.md#APIsetPrivateChatSessionPinnedRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsetPrivateChatSessionPinnedRequest"></a>
## APIsetPrivateChatSessionPinnedRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **setPrivateChatSessionPinnedReq** | [**SetPrivateChatSessionPinnedReq**](SetPrivateChatSessionPinnedReq.md) |  | |


