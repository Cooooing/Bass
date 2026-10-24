# GlobalChatService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getGlobalChatTopic**](GlobalChatService.md#getGlobalChatTopic) | **POST** /v1/im/global-chat/topic/get |  |
| [**getGlobalChatTopicWithHttpInfo**](GlobalChatService.md#getGlobalChatTopicWithHttpInfo) | **POST** /v1/im/global-chat/topic/get |  |
| [**listGlobalChatBarrages**](GlobalChatService.md#listGlobalChatBarrages) | **POST** /v1/im/global-chat/barrages/list |  |
| [**listGlobalChatBarragesWithHttpInfo**](GlobalChatService.md#listGlobalChatBarragesWithHttpInfo) | **POST** /v1/im/global-chat/barrages/list |  |
| [**listGlobalChatMessages**](GlobalChatService.md#listGlobalChatMessages) | **POST** /v1/im/global-chat/messages/list |  |
| [**listGlobalChatMessagesWithHttpInfo**](GlobalChatService.md#listGlobalChatMessagesWithHttpInfo) | **POST** /v1/im/global-chat/messages/list |  |
| [**listGlobalChatOnlineMembers**](GlobalChatService.md#listGlobalChatOnlineMembers) | **POST** /v1/im/global-chat/online-members/list |  |
| [**listGlobalChatOnlineMembersWithHttpInfo**](GlobalChatService.md#listGlobalChatOnlineMembersWithHttpInfo) | **POST** /v1/im/global-chat/online-members/list |  |
| [**revokeGlobalChatMessage**](GlobalChatService.md#revokeGlobalChatMessage) | **POST** /v1/im/global-chat/messages/revoke |  |
| [**revokeGlobalChatMessageWithHttpInfo**](GlobalChatService.md#revokeGlobalChatMessageWithHttpInfo) | **POST** /v1/im/global-chat/messages/revoke |  |
| [**sendGlobalChatBarrage**](GlobalChatService.md#sendGlobalChatBarrage) | **POST** /v1/im/global-chat/barrages |  |
| [**sendGlobalChatBarrageWithHttpInfo**](GlobalChatService.md#sendGlobalChatBarrageWithHttpInfo) | **POST** /v1/im/global-chat/barrages |  |
| [**sendGlobalChatMessage**](GlobalChatService.md#sendGlobalChatMessage) | **POST** /v1/im/global-chat/messages |  |
| [**sendGlobalChatMessageWithHttpInfo**](GlobalChatService.md#sendGlobalChatMessageWithHttpInfo) | **POST** /v1/im/global-chat/messages |  |
| [**setGlobalChatTopic**](GlobalChatService.md#setGlobalChatTopic) | **POST** /v1/im/global-chat/topic/set |  |
| [**setGlobalChatTopicWithHttpInfo**](GlobalChatService.md#setGlobalChatTopicWithHttpInfo) | **POST** /v1/im/global-chat/topic/set |  |



## getGlobalChatTopic

> GetGlobalChatTopicResp getGlobalChatTopic(getGlobalChatTopicRequest)



查询当前全局聊天室话题。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetGlobalChatTopicRequest request = APIgetGlobalChatTopicRequest.newBuilder()
                .body(body)
                .build();
            GetGlobalChatTopicResp result = apiInstance.getGlobalChatTopic(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#getGlobalChatTopic");
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
| getGlobalChatTopicRequest | [**APIgetGlobalChatTopicRequest**](GlobalChatService.md#APIgetGlobalChatTopicRequest)|-|-|

### Return type

[**GetGlobalChatTopicResp**](GetGlobalChatTopicResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## getGlobalChatTopicWithHttpInfo

> ApiResponse<GetGlobalChatTopicResp> getGlobalChatTopicWithHttpInfo(getGlobalChatTopicRequest)



查询当前全局聊天室话题。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetGlobalChatTopicRequest request = APIgetGlobalChatTopicRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetGlobalChatTopicResp> response = apiInstance.getGlobalChatTopicWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#getGlobalChatTopic");
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
| getGlobalChatTopicRequest | [**APIgetGlobalChatTopicRequest**](GlobalChatService.md#APIgetGlobalChatTopicRequest)|-|-|

### Return type

ApiResponse<[**GetGlobalChatTopicResp**](GetGlobalChatTopicResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIgetGlobalChatTopicRequest"></a>
## APIgetGlobalChatTopicRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## listGlobalChatBarrages

> ListGlobalChatBarragesResp listGlobalChatBarrages(listGlobalChatBarragesRequest)



分页查询已发布的全局聊天室弹幕。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        ListGlobalChatBarragesReq listGlobalChatBarragesReq = new ListGlobalChatBarragesReq(); // ListGlobalChatBarragesReq | 
        try {
            APIlistGlobalChatBarragesRequest request = APIlistGlobalChatBarragesRequest.newBuilder()
                .listGlobalChatBarragesReq(listGlobalChatBarragesReq)
                .build();
            ListGlobalChatBarragesResp result = apiInstance.listGlobalChatBarrages(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#listGlobalChatBarrages");
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
| listGlobalChatBarragesRequest | [**APIlistGlobalChatBarragesRequest**](GlobalChatService.md#APIlistGlobalChatBarragesRequest)|-|-|

### Return type

[**ListGlobalChatBarragesResp**](ListGlobalChatBarragesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listGlobalChatBarragesWithHttpInfo

> ApiResponse<ListGlobalChatBarragesResp> listGlobalChatBarragesWithHttpInfo(listGlobalChatBarragesRequest)



分页查询已发布的全局聊天室弹幕。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        ListGlobalChatBarragesReq listGlobalChatBarragesReq = new ListGlobalChatBarragesReq(); // ListGlobalChatBarragesReq | 
        try {
            APIlistGlobalChatBarragesRequest request = APIlistGlobalChatBarragesRequest.newBuilder()
                .listGlobalChatBarragesReq(listGlobalChatBarragesReq)
                .build();
            ApiResponse<ListGlobalChatBarragesResp> response = apiInstance.listGlobalChatBarragesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#listGlobalChatBarrages");
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
| listGlobalChatBarragesRequest | [**APIlistGlobalChatBarragesRequest**](GlobalChatService.md#APIlistGlobalChatBarragesRequest)|-|-|

### Return type

ApiResponse<[**ListGlobalChatBarragesResp**](ListGlobalChatBarragesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistGlobalChatBarragesRequest"></a>
## APIlistGlobalChatBarragesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listGlobalChatBarragesReq** | [**ListGlobalChatBarragesReq**](ListGlobalChatBarragesReq.md) |  | |



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
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        ListGlobalChatMessagesReq listGlobalChatMessagesReq = new ListGlobalChatMessagesReq(); // ListGlobalChatMessagesReq | 
        try {
            APIlistGlobalChatMessagesRequest request = APIlistGlobalChatMessagesRequest.newBuilder()
                .listGlobalChatMessagesReq(listGlobalChatMessagesReq)
                .build();
            ListGlobalChatMessagesResp result = apiInstance.listGlobalChatMessages(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#listGlobalChatMessages");
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
| listGlobalChatMessagesRequest | [**APIlistGlobalChatMessagesRequest**](GlobalChatService.md#APIlistGlobalChatMessagesRequest)|-|-|

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
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
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
            System.err.println("Exception when calling GlobalChatService#listGlobalChatMessages");
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
| listGlobalChatMessagesRequest | [**APIlistGlobalChatMessagesRequest**](GlobalChatService.md#APIlistGlobalChatMessagesRequest)|-|-|

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
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        Object body = null; // Object | 
        try {
            APIlistGlobalChatOnlineMembersRequest request = APIlistGlobalChatOnlineMembersRequest.newBuilder()
                .body(body)
                .build();
            ListGlobalChatOnlineMembersResp result = apiInstance.listGlobalChatOnlineMembers(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#listGlobalChatOnlineMembers");
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
| listGlobalChatOnlineMembersRequest | [**APIlistGlobalChatOnlineMembersRequest**](GlobalChatService.md#APIlistGlobalChatOnlineMembersRequest)|-|-|

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
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
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
            System.err.println("Exception when calling GlobalChatService#listGlobalChatOnlineMembers");
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
| listGlobalChatOnlineMembersRequest | [**APIlistGlobalChatOnlineMembersRequest**](GlobalChatService.md#APIlistGlobalChatOnlineMembersRequest)|-|-|

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
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        RevokeGlobalChatMessageReq revokeGlobalChatMessageReq = new RevokeGlobalChatMessageReq(); // RevokeGlobalChatMessageReq | 
        try {
            APIrevokeGlobalChatMessageRequest request = APIrevokeGlobalChatMessageRequest.newBuilder()
                .revokeGlobalChatMessageReq(revokeGlobalChatMessageReq)
                .build();
            RevokeGlobalChatMessageResp result = apiInstance.revokeGlobalChatMessage(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#revokeGlobalChatMessage");
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
| revokeGlobalChatMessageRequest | [**APIrevokeGlobalChatMessageRequest**](GlobalChatService.md#APIrevokeGlobalChatMessageRequest)|-|-|

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
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
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
            System.err.println("Exception when calling GlobalChatService#revokeGlobalChatMessage");
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
| revokeGlobalChatMessageRequest | [**APIrevokeGlobalChatMessageRequest**](GlobalChatService.md#APIrevokeGlobalChatMessageRequest)|-|-|

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



## sendGlobalChatBarrage

> SendGlobalChatBarrageResp sendGlobalChatBarrage(sendGlobalChatBarrageRequest)



发送全局聊天室弹幕。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        SendGlobalChatBarrageReq sendGlobalChatBarrageReq = new SendGlobalChatBarrageReq(); // SendGlobalChatBarrageReq | 
        try {
            APIsendGlobalChatBarrageRequest request = APIsendGlobalChatBarrageRequest.newBuilder()
                .sendGlobalChatBarrageReq(sendGlobalChatBarrageReq)
                .build();
            SendGlobalChatBarrageResp result = apiInstance.sendGlobalChatBarrage(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#sendGlobalChatBarrage");
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
| sendGlobalChatBarrageRequest | [**APIsendGlobalChatBarrageRequest**](GlobalChatService.md#APIsendGlobalChatBarrageRequest)|-|-|

### Return type

[**SendGlobalChatBarrageResp**](SendGlobalChatBarrageResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## sendGlobalChatBarrageWithHttpInfo

> ApiResponse<SendGlobalChatBarrageResp> sendGlobalChatBarrageWithHttpInfo(sendGlobalChatBarrageRequest)



发送全局聊天室弹幕。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        SendGlobalChatBarrageReq sendGlobalChatBarrageReq = new SendGlobalChatBarrageReq(); // SendGlobalChatBarrageReq | 
        try {
            APIsendGlobalChatBarrageRequest request = APIsendGlobalChatBarrageRequest.newBuilder()
                .sendGlobalChatBarrageReq(sendGlobalChatBarrageReq)
                .build();
            ApiResponse<SendGlobalChatBarrageResp> response = apiInstance.sendGlobalChatBarrageWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#sendGlobalChatBarrage");
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
| sendGlobalChatBarrageRequest | [**APIsendGlobalChatBarrageRequest**](GlobalChatService.md#APIsendGlobalChatBarrageRequest)|-|-|

### Return type

ApiResponse<[**SendGlobalChatBarrageResp**](SendGlobalChatBarrageResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsendGlobalChatBarrageRequest"></a>
## APIsendGlobalChatBarrageRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **sendGlobalChatBarrageReq** | [**SendGlobalChatBarrageReq**](SendGlobalChatBarrageReq.md) |  | |



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
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        SendGlobalChatMessageReq sendGlobalChatMessageReq = new SendGlobalChatMessageReq(); // SendGlobalChatMessageReq | 
        try {
            APIsendGlobalChatMessageRequest request = APIsendGlobalChatMessageRequest.newBuilder()
                .sendGlobalChatMessageReq(sendGlobalChatMessageReq)
                .build();
            SendGlobalChatMessageResp result = apiInstance.sendGlobalChatMessage(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#sendGlobalChatMessage");
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
| sendGlobalChatMessageRequest | [**APIsendGlobalChatMessageRequest**](GlobalChatService.md#APIsendGlobalChatMessageRequest)|-|-|

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
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
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
            System.err.println("Exception when calling GlobalChatService#sendGlobalChatMessage");
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
| sendGlobalChatMessageRequest | [**APIsendGlobalChatMessageRequest**](GlobalChatService.md#APIsendGlobalChatMessageRequest)|-|-|

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



## setGlobalChatTopic

> SetGlobalChatTopicResp setGlobalChatTopic(setGlobalChatTopicRequest)



设置全局聊天室话题。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        SetGlobalChatTopicReq setGlobalChatTopicReq = new SetGlobalChatTopicReq(); // SetGlobalChatTopicReq | 
        try {
            APIsetGlobalChatTopicRequest request = APIsetGlobalChatTopicRequest.newBuilder()
                .setGlobalChatTopicReq(setGlobalChatTopicReq)
                .build();
            SetGlobalChatTopicResp result = apiInstance.setGlobalChatTopic(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#setGlobalChatTopic");
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
| setGlobalChatTopicRequest | [**APIsetGlobalChatTopicRequest**](GlobalChatService.md#APIsetGlobalChatTopicRequest)|-|-|

### Return type

[**SetGlobalChatTopicResp**](SetGlobalChatTopicResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## setGlobalChatTopicWithHttpInfo

> ApiResponse<SetGlobalChatTopicResp> setGlobalChatTopicWithHttpInfo(setGlobalChatTopicRequest)



设置全局聊天室话题。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.GlobalChatService;
import com.bass.bbs.api.GlobalChatService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        GlobalChatService apiInstance = new GlobalChatService(defaultClient);
        SetGlobalChatTopicReq setGlobalChatTopicReq = new SetGlobalChatTopicReq(); // SetGlobalChatTopicReq | 
        try {
            APIsetGlobalChatTopicRequest request = APIsetGlobalChatTopicRequest.newBuilder()
                .setGlobalChatTopicReq(setGlobalChatTopicReq)
                .build();
            ApiResponse<SetGlobalChatTopicResp> response = apiInstance.setGlobalChatTopicWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling GlobalChatService#setGlobalChatTopic");
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
| setGlobalChatTopicRequest | [**APIsetGlobalChatTopicRequest**](GlobalChatService.md#APIsetGlobalChatTopicRequest)|-|-|

### Return type

ApiResponse<[**SetGlobalChatTopicResp**](SetGlobalChatTopicResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsetGlobalChatTopicRequest"></a>
## APIsetGlobalChatTopicRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **setGlobalChatTopicReq** | [**SetGlobalChatTopicReq**](SetGlobalChatTopicReq.md) |  | |


