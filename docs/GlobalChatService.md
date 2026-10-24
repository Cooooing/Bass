# GlobalChatService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getGlobalChatTopic**](#getglobalchattopic) | **POST** /v1/im/global-chat/topic/get | |
|[**listGlobalChatBarrages**](#listglobalchatbarrages) | **POST** /v1/im/global-chat/barrages/list | |
|[**listGlobalChatMessages**](#listglobalchatmessages) | **POST** /v1/im/global-chat/messages/list | |
|[**listGlobalChatOnlineMembers**](#listglobalchatonlinemembers) | **POST** /v1/im/global-chat/online-members/list | |
|[**revokeGlobalChatMessage**](#revokeglobalchatmessage) | **POST** /v1/im/global-chat/messages/revoke | |
|[**sendGlobalChatBarrage**](#sendglobalchatbarrage) | **POST** /v1/im/global-chat/barrages | |
|[**sendGlobalChatMessage**](#sendglobalchatmessage) | **POST** /v1/im/global-chat/messages | |
|[**setGlobalChatTopic**](#setglobalchattopic) | **POST** /v1/im/global-chat/topic/set | |

# **getGlobalChatTopic**
> GetGlobalChatTopicResp getGlobalChatTopic(body)

查询当前全局聊天室话题。

### Example

```typescript
import {
    GlobalChatService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new GlobalChatService(configuration);

let body: object; //

const { status, data } = await apiInstance.getGlobalChatTopic(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**GetGlobalChatTopicResp**

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

# **listGlobalChatBarrages**
> ListGlobalChatBarragesResp listGlobalChatBarrages(listGlobalChatBarragesReq)

分页查询已发布的全局聊天室弹幕。

### Example

```typescript
import {
    GlobalChatService,
    Configuration,
    ListGlobalChatBarragesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new GlobalChatService(configuration);

let listGlobalChatBarragesReq: ListGlobalChatBarragesReq; //

const { status, data } = await apiInstance.listGlobalChatBarrages(
    listGlobalChatBarragesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listGlobalChatBarragesReq** | **ListGlobalChatBarragesReq**|  | |


### Return type

**ListGlobalChatBarragesResp**

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

# **listGlobalChatMessages**
> ListGlobalChatMessagesResp listGlobalChatMessages(listGlobalChatMessagesReq)

查询全局聊天室消息。

### Example

```typescript
import {
    GlobalChatService,
    Configuration,
    ListGlobalChatMessagesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new GlobalChatService(configuration);

let listGlobalChatMessagesReq: ListGlobalChatMessagesReq; //

const { status, data } = await apiInstance.listGlobalChatMessages(
    listGlobalChatMessagesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listGlobalChatMessagesReq** | **ListGlobalChatMessagesReq**|  | |


### Return type

**ListGlobalChatMessagesResp**

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

# **listGlobalChatOnlineMembers**
> ListGlobalChatOnlineMembersResp listGlobalChatOnlineMembers(body)

查询全部在线的聊天室成员。

### Example

```typescript
import {
    GlobalChatService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new GlobalChatService(configuration);

let body: object; //

const { status, data } = await apiInstance.listGlobalChatOnlineMembers(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**ListGlobalChatOnlineMembersResp**

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

# **revokeGlobalChatMessage**
> RevokeGlobalChatMessageResp revokeGlobalChatMessage(revokeGlobalChatMessageReq)

撤回自己的全局聊天室消息。

### Example

```typescript
import {
    GlobalChatService,
    Configuration,
    RevokeGlobalChatMessageReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new GlobalChatService(configuration);

let revokeGlobalChatMessageReq: RevokeGlobalChatMessageReq; //

const { status, data } = await apiInstance.revokeGlobalChatMessage(
    revokeGlobalChatMessageReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **revokeGlobalChatMessageReq** | **RevokeGlobalChatMessageReq**|  | |


### Return type

**RevokeGlobalChatMessageResp**

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

# **sendGlobalChatBarrage**
> SendGlobalChatBarrageResp sendGlobalChatBarrage(sendGlobalChatBarrageReq)

发送全局聊天室弹幕。

### Example

```typescript
import {
    GlobalChatService,
    Configuration,
    SendGlobalChatBarrageReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new GlobalChatService(configuration);

let sendGlobalChatBarrageReq: SendGlobalChatBarrageReq; //

const { status, data } = await apiInstance.sendGlobalChatBarrage(
    sendGlobalChatBarrageReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **sendGlobalChatBarrageReq** | **SendGlobalChatBarrageReq**|  | |


### Return type

**SendGlobalChatBarrageResp**

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

# **sendGlobalChatMessage**
> SendGlobalChatMessageResp sendGlobalChatMessage(sendGlobalChatMessageReq)

发送全局聊天室消息。

### Example

```typescript
import {
    GlobalChatService,
    Configuration,
    SendGlobalChatMessageReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new GlobalChatService(configuration);

let sendGlobalChatMessageReq: SendGlobalChatMessageReq; //

const { status, data } = await apiInstance.sendGlobalChatMessage(
    sendGlobalChatMessageReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **sendGlobalChatMessageReq** | **SendGlobalChatMessageReq**|  | |


### Return type

**SendGlobalChatMessageResp**

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

# **setGlobalChatTopic**
> SetGlobalChatTopicResp setGlobalChatTopic(setGlobalChatTopicReq)

设置全局聊天室话题。

### Example

```typescript
import {
    GlobalChatService,
    Configuration,
    SetGlobalChatTopicReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new GlobalChatService(configuration);

let setGlobalChatTopicReq: SetGlobalChatTopicReq; //

const { status, data } = await apiInstance.setGlobalChatTopic(
    setGlobalChatTopicReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **setGlobalChatTopicReq** | **SetGlobalChatTopicReq**|  | |


### Return type

**SetGlobalChatTopicResp**

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

