# ImService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**countPrivateChatUnread**](ImService.md#countPrivateChatUnread) | **POST** /v1/im/private-chat/unread-count |  |
| [**countPrivateChatUnreadWithHttpInfo**](ImService.md#countPrivateChatUnreadWithHttpInfo) | **POST** /v1/im/private-chat/unread-count |  |
| [**listGlobalChatMessages**](ImService.md#listGlobalChatMessages) | **POST** /v1/im/global-chat/messages/list |  |
| [**listGlobalChatMessagesWithHttpInfo**](ImService.md#listGlobalChatMessagesWithHttpInfo) | **POST** /v1/im/global-chat/messages/list |  |
| [**listGlobalChatOnlineMembers**](ImService.md#listGlobalChatOnlineMembers) | **POST** /v1/im/global-chat/online-members/list |  |
| [**listGlobalChatOnlineMembersWithHttpInfo**](ImService.md#listGlobalChatOnlineMembersWithHttpInfo) | **POST** /v1/im/global-chat/online-members/list |  |
| [**listPrivateChatMessages**](ImService.md#listPrivateChatMessages) | **POST** /v1/im/private-chat/messages/list |  |
| [**listPrivateChatMessagesWithHttpInfo**](ImService.md#listPrivateChatMessagesWithHttpInfo) | **POST** /v1/im/private-chat/messages/list |  |
| [**listPrivateChatSessions**](ImService.md#listPrivateChatSessions) | **GET** /v1/im/private-chat/sessions |  |
| [**listPrivateChatSessionsWithHttpInfo**](ImService.md#listPrivateChatSessionsWithHttpInfo) | **GET** /v1/im/private-chat/sessions |  |
| [**markPrivateChatSessionRead**](ImService.md#markPrivateChatSessionRead) | **POST** /v1/im/private-chat/read |  |
| [**markPrivateChatSessionReadWithHttpInfo**](ImService.md#markPrivateChatSessionReadWithHttpInfo) | **POST** /v1/im/private-chat/read |  |
| [**revokeGlobalChatMessage**](ImService.md#revokeGlobalChatMessage) | **POST** /v1/im/global-chat/messages/revoke |  |
| [**revokeGlobalChatMessageWithHttpInfo**](ImService.md#revokeGlobalChatMessageWithHttpInfo) | **POST** /v1/im/global-chat/messages/revoke |  |
| [**revokePrivateChatMessage**](ImService.md#revokePrivateChatMessage) | **POST** /v1/im/private-chat/messages/revoke |  |
| [**revokePrivateChatMessageWithHttpInfo**](ImService.md#revokePrivateChatMessageWithHttpInfo) | **POST** /v1/im/private-chat/messages/revoke |  |
| [**sendGlobalChatMessage**](ImService.md#sendGlobalChatMessage) | **POST** /v1/im/global-chat/messages |  |
| [**sendGlobalChatMessageWithHttpInfo**](ImService.md#sendGlobalChatMessageWithHttpInfo) | **POST** /v1/im/global-chat/messages |  |
| [**sendPrivateChatMessage**](ImService.md#sendPrivateChatMessage) | **POST** /v1/im/private-chat/messages |  |
| [**sendPrivateChatMessageWithHttpInfo**](ImService.md#sendPrivateChatMessageWithHttpInfo) | **POST** /v1/im/private-chat/messages |  |
| [**setPrivateChatSessionMuted**](ImService.md#setPrivateChatSessionMuted) | **POST** /v1/im/private-chat/muted |  |
| [**setPrivateChatSessionMutedWithHttpInfo**](ImService.md#setPrivateChatSessionMutedWithHttpInfo) | **POST** /v1/im/private-chat/muted |  |
| [**setPrivateChatSessionPinned**](ImService.md#setPrivateChatSessionPinned) | **POST** /v1/im/private-chat/pinned |  |
| [**setPrivateChatSessionPinnedWithHttpInfo**](ImService.md#setPrivateChatSessionPinnedWithHttpInfo) | **POST** /v1/im/private-chat/pinned |  |



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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcountPrivateChatUnreadRequest request = APIcountPrivateChatUnreadRequest.newBuilder()
                .body(body)
                .build();
            CountPrivateChatUnreadResp result = apiInstance.countPrivateChatUnread(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#countPrivateChatUnread");
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
| countPrivateChatUnreadRequest | [**APIcountPrivateChatUnreadRequest**](ImService.md#APIcountPrivateChatUnreadRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
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
            System.err.println("Exception when calling ImService#countPrivateChatUnread");
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
| countPrivateChatUnreadRequest | [**APIcountPrivateChatUnreadRequest**](ImService.md#APIcountPrivateChatUnreadRequest)|-|-|

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



## listGlobalChatMessages

> ListGlobalChatMessagesResp listGlobalChatMessages(listGlobalChatMessagesRequest)



查询全局聊天室消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        ListGlobalChatMessagesReq listGlobalChatMessagesReq = new ListGlobalChatMessagesReq(); // ListGlobalChatMessagesReq | 
        try {
            APIlistGlobalChatMessagesRequest request = APIlistGlobalChatMessagesRequest.newBuilder()
                .listGlobalChatMessagesReq(listGlobalChatMessagesReq)
                .build();
            ListGlobalChatMessagesResp result = apiInstance.listGlobalChatMessages(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#listGlobalChatMessages");
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
| listGlobalChatMessagesRequest | [**APIlistGlobalChatMessagesRequest**](ImService.md#APIlistGlobalChatMessagesRequest)|-|-|

### Return type

[**ListGlobalChatMessagesResp**](ListGlobalChatMessagesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listGlobalChatMessagesWithHttpInfo

> ApiResponse<ListGlobalChatMessagesResp> listGlobalChatMessagesWithHttpInfo(listGlobalChatMessagesRequest)



查询全局聊天室消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        ListGlobalChatMessagesReq listGlobalChatMessagesReq = new ListGlobalChatMessagesReq(); // ListGlobalChatMessagesReq | 
        try {
            APIlistGlobalChatMessagesRequest request = APIlistGlobalChatMessagesRequest.newBuilder()
                .listGlobalChatMessagesReq(listGlobalChatMessagesReq)
                .build();
            ApiResponse<ListGlobalChatMessagesResp> response = apiInstance.listGlobalChatMessagesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#listGlobalChatMessages");
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
| listGlobalChatMessagesRequest | [**APIlistGlobalChatMessagesRequest**](ImService.md#APIlistGlobalChatMessagesRequest)|-|-|

### Return type

ApiResponse<[**ListGlobalChatMessagesResp**](ListGlobalChatMessagesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistGlobalChatMessagesRequest"></a>
## APIlistGlobalChatMessagesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listGlobalChatMessagesReq** | [**ListGlobalChatMessagesReq**](ListGlobalChatMessagesReq.md) |  | |



## listGlobalChatOnlineMembers

> ListGlobalChatOnlineMembersResp listGlobalChatOnlineMembers(listGlobalChatOnlineMembersRequest)



查询全部在线的聊天室成员。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        Object body = null; // Object | 
        try {
            APIlistGlobalChatOnlineMembersRequest request = APIlistGlobalChatOnlineMembersRequest.newBuilder()
                .body(body)
                .build();
            ListGlobalChatOnlineMembersResp result = apiInstance.listGlobalChatOnlineMembers(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#listGlobalChatOnlineMembers");
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
| listGlobalChatOnlineMembersRequest | [**APIlistGlobalChatOnlineMembersRequest**](ImService.md#APIlistGlobalChatOnlineMembersRequest)|-|-|

### Return type

[**ListGlobalChatOnlineMembersResp**](ListGlobalChatOnlineMembersResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listGlobalChatOnlineMembersWithHttpInfo

> ApiResponse<ListGlobalChatOnlineMembersResp> listGlobalChatOnlineMembersWithHttpInfo(listGlobalChatOnlineMembersRequest)



查询全部在线的聊天室成员。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        Object body = null; // Object | 
        try {
            APIlistGlobalChatOnlineMembersRequest request = APIlistGlobalChatOnlineMembersRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<ListGlobalChatOnlineMembersResp> response = apiInstance.listGlobalChatOnlineMembersWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#listGlobalChatOnlineMembers");
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
| listGlobalChatOnlineMembersRequest | [**APIlistGlobalChatOnlineMembersRequest**](ImService.md#APIlistGlobalChatOnlineMembersRequest)|-|-|

### Return type

ApiResponse<[**ListGlobalChatOnlineMembersResp**](ListGlobalChatOnlineMembersResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistGlobalChatOnlineMembersRequest"></a>
## APIlistGlobalChatOnlineMembersRequest
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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        ListPrivateChatMessagesReq listPrivateChatMessagesReq = new ListPrivateChatMessagesReq(); // ListPrivateChatMessagesReq | 
        try {
            APIlistPrivateChatMessagesRequest request = APIlistPrivateChatMessagesRequest.newBuilder()
                .listPrivateChatMessagesReq(listPrivateChatMessagesReq)
                .build();
            ListPrivateChatMessagesResp result = apiInstance.listPrivateChatMessages(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#listPrivateChatMessages");
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
| listPrivateChatMessagesRequest | [**APIlistPrivateChatMessagesRequest**](ImService.md#APIlistPrivateChatMessagesRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
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
            System.err.println("Exception when calling ImService#listPrivateChatMessages");
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
| listPrivateChatMessagesRequest | [**APIlistPrivateChatMessagesRequest**](ImService.md#APIlistPrivateChatMessagesRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
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
            System.err.println("Exception when calling ImService#listPrivateChatSessions");
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
| listPrivateChatSessionsRequest | [**APIlistPrivateChatSessionsRequest**](ImService.md#APIlistPrivateChatSessionsRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
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
            System.err.println("Exception when calling ImService#listPrivateChatSessions");
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
| listPrivateChatSessionsRequest | [**APIlistPrivateChatSessionsRequest**](ImService.md#APIlistPrivateChatSessionsRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        MarkPrivateChatSessionReadReq markPrivateChatSessionReadReq = new MarkPrivateChatSessionReadReq(); // MarkPrivateChatSessionReadReq | 
        try {
            APImarkPrivateChatSessionReadRequest request = APImarkPrivateChatSessionReadRequest.newBuilder()
                .markPrivateChatSessionReadReq(markPrivateChatSessionReadReq)
                .build();
            Object result = apiInstance.markPrivateChatSessionRead(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#markPrivateChatSessionRead");
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
| markPrivateChatSessionReadRequest | [**APImarkPrivateChatSessionReadRequest**](ImService.md#APImarkPrivateChatSessionReadRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
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
            System.err.println("Exception when calling ImService#markPrivateChatSessionRead");
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
| markPrivateChatSessionReadRequest | [**APImarkPrivateChatSessionReadRequest**](ImService.md#APImarkPrivateChatSessionReadRequest)|-|-|

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



## revokeGlobalChatMessage

> RevokeGlobalChatMessageResp revokeGlobalChatMessage(revokeGlobalChatMessageRequest)



撤回自己的全局聊天室消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        RevokeGlobalChatMessageReq revokeGlobalChatMessageReq = new RevokeGlobalChatMessageReq(); // RevokeGlobalChatMessageReq | 
        try {
            APIrevokeGlobalChatMessageRequest request = APIrevokeGlobalChatMessageRequest.newBuilder()
                .revokeGlobalChatMessageReq(revokeGlobalChatMessageReq)
                .build();
            RevokeGlobalChatMessageResp result = apiInstance.revokeGlobalChatMessage(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#revokeGlobalChatMessage");
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
| revokeGlobalChatMessageRequest | [**APIrevokeGlobalChatMessageRequest**](ImService.md#APIrevokeGlobalChatMessageRequest)|-|-|

### Return type

[**RevokeGlobalChatMessageResp**](RevokeGlobalChatMessageResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## revokeGlobalChatMessageWithHttpInfo

> ApiResponse<RevokeGlobalChatMessageResp> revokeGlobalChatMessageWithHttpInfo(revokeGlobalChatMessageRequest)



撤回自己的全局聊天室消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        RevokeGlobalChatMessageReq revokeGlobalChatMessageReq = new RevokeGlobalChatMessageReq(); // RevokeGlobalChatMessageReq | 
        try {
            APIrevokeGlobalChatMessageRequest request = APIrevokeGlobalChatMessageRequest.newBuilder()
                .revokeGlobalChatMessageReq(revokeGlobalChatMessageReq)
                .build();
            ApiResponse<RevokeGlobalChatMessageResp> response = apiInstance.revokeGlobalChatMessageWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#revokeGlobalChatMessage");
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
| revokeGlobalChatMessageRequest | [**APIrevokeGlobalChatMessageRequest**](ImService.md#APIrevokeGlobalChatMessageRequest)|-|-|

### Return type

ApiResponse<[**RevokeGlobalChatMessageResp**](RevokeGlobalChatMessageResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrevokeGlobalChatMessageRequest"></a>
## APIrevokeGlobalChatMessageRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **revokeGlobalChatMessageReq** | [**RevokeGlobalChatMessageReq**](RevokeGlobalChatMessageReq.md) |  | |



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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        RevokePrivateChatMessageReq revokePrivateChatMessageReq = new RevokePrivateChatMessageReq(); // RevokePrivateChatMessageReq | 
        try {
            APIrevokePrivateChatMessageRequest request = APIrevokePrivateChatMessageRequest.newBuilder()
                .revokePrivateChatMessageReq(revokePrivateChatMessageReq)
                .build();
            RevokePrivateChatMessageResp result = apiInstance.revokePrivateChatMessage(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#revokePrivateChatMessage");
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
| revokePrivateChatMessageRequest | [**APIrevokePrivateChatMessageRequest**](ImService.md#APIrevokePrivateChatMessageRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
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
            System.err.println("Exception when calling ImService#revokePrivateChatMessage");
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
| revokePrivateChatMessageRequest | [**APIrevokePrivateChatMessageRequest**](ImService.md#APIrevokePrivateChatMessageRequest)|-|-|

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



## sendGlobalChatMessage

> SendGlobalChatMessageResp sendGlobalChatMessage(sendGlobalChatMessageRequest)



发送全局聊天室消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        SendGlobalChatMessageReq sendGlobalChatMessageReq = new SendGlobalChatMessageReq(); // SendGlobalChatMessageReq | 
        try {
            APIsendGlobalChatMessageRequest request = APIsendGlobalChatMessageRequest.newBuilder()
                .sendGlobalChatMessageReq(sendGlobalChatMessageReq)
                .build();
            SendGlobalChatMessageResp result = apiInstance.sendGlobalChatMessage(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#sendGlobalChatMessage");
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
| sendGlobalChatMessageRequest | [**APIsendGlobalChatMessageRequest**](ImService.md#APIsendGlobalChatMessageRequest)|-|-|

### Return type

[**SendGlobalChatMessageResp**](SendGlobalChatMessageResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## sendGlobalChatMessageWithHttpInfo

> ApiResponse<SendGlobalChatMessageResp> sendGlobalChatMessageWithHttpInfo(sendGlobalChatMessageRequest)



发送全局聊天室消息。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        SendGlobalChatMessageReq sendGlobalChatMessageReq = new SendGlobalChatMessageReq(); // SendGlobalChatMessageReq | 
        try {
            APIsendGlobalChatMessageRequest request = APIsendGlobalChatMessageRequest.newBuilder()
                .sendGlobalChatMessageReq(sendGlobalChatMessageReq)
                .build();
            ApiResponse<SendGlobalChatMessageResp> response = apiInstance.sendGlobalChatMessageWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#sendGlobalChatMessage");
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
| sendGlobalChatMessageRequest | [**APIsendGlobalChatMessageRequest**](ImService.md#APIsendGlobalChatMessageRequest)|-|-|

### Return type

ApiResponse<[**SendGlobalChatMessageResp**](SendGlobalChatMessageResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsendGlobalChatMessageRequest"></a>
## APIsendGlobalChatMessageRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **sendGlobalChatMessageReq** | [**SendGlobalChatMessageReq**](SendGlobalChatMessageReq.md) |  | |



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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        SendPrivateChatMessageReq sendPrivateChatMessageReq = new SendPrivateChatMessageReq(); // SendPrivateChatMessageReq | 
        try {
            APIsendPrivateChatMessageRequest request = APIsendPrivateChatMessageRequest.newBuilder()
                .sendPrivateChatMessageReq(sendPrivateChatMessageReq)
                .build();
            SendPrivateChatMessageResp result = apiInstance.sendPrivateChatMessage(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#sendPrivateChatMessage");
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
| sendPrivateChatMessageRequest | [**APIsendPrivateChatMessageRequest**](ImService.md#APIsendPrivateChatMessageRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
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
            System.err.println("Exception when calling ImService#sendPrivateChatMessage");
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
| sendPrivateChatMessageRequest | [**APIsendPrivateChatMessageRequest**](ImService.md#APIsendPrivateChatMessageRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        SetPrivateChatSessionMutedReq setPrivateChatSessionMutedReq = new SetPrivateChatSessionMutedReq(); // SetPrivateChatSessionMutedReq | 
        try {
            APIsetPrivateChatSessionMutedRequest request = APIsetPrivateChatSessionMutedRequest.newBuilder()
                .setPrivateChatSessionMutedReq(setPrivateChatSessionMutedReq)
                .build();
            Object result = apiInstance.setPrivateChatSessionMuted(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#setPrivateChatSessionMuted");
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
| setPrivateChatSessionMutedRequest | [**APIsetPrivateChatSessionMutedRequest**](ImService.md#APIsetPrivateChatSessionMutedRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
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
            System.err.println("Exception when calling ImService#setPrivateChatSessionMuted");
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
| setPrivateChatSessionMutedRequest | [**APIsetPrivateChatSessionMutedRequest**](ImService.md#APIsetPrivateChatSessionMutedRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
        SetPrivateChatSessionPinnedReq setPrivateChatSessionPinnedReq = new SetPrivateChatSessionPinnedReq(); // SetPrivateChatSessionPinnedReq | 
        try {
            APIsetPrivateChatSessionPinnedRequest request = APIsetPrivateChatSessionPinnedRequest.newBuilder()
                .setPrivateChatSessionPinnedReq(setPrivateChatSessionPinnedReq)
                .build();
            Object result = apiInstance.setPrivateChatSessionPinned(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling ImService#setPrivateChatSessionPinned");
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
| setPrivateChatSessionPinnedRequest | [**APIsetPrivateChatSessionPinnedRequest**](ImService.md#APIsetPrivateChatSessionPinnedRequest)|-|-|

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
import com.bass.bbs.api.ImService;
import com.bass.bbs.api.ImService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        ImService apiInstance = new ImService(defaultClient);
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
            System.err.println("Exception when calling ImService#setPrivateChatSessionPinned");
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
| setPrivateChatSessionPinnedRequest | [**APIsetPrivateChatSessionPinnedRequest**](ImService.md#APIsetPrivateChatSessionPinnedRequest)|-|-|

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


