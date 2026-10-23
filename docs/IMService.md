# IMService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**countPrivateChatUnread**](#countprivatechatunread) | **POST** /v1/im/private-chat/unread-count | |
|[**listGlobalChatMessages**](#listglobalchatmessages) | **POST** /v1/im/global-chat/messages/list | |
|[**listGlobalChatOnlineMembers**](#listglobalchatonlinemembers) | **POST** /v1/im/global-chat/online-members/list | |
|[**listPrivateChatMessages**](#listprivatechatmessages) | **POST** /v1/im/private-chat/messages/list | |
|[**listPrivateChatSessions**](#listprivatechatsessions) | **GET** /v1/im/private-chat/sessions | |
|[**markPrivateChatSessionRead**](#markprivatechatsessionread) | **POST** /v1/im/private-chat/read | |
|[**revokeGlobalChatMessage**](#revokeglobalchatmessage) | **POST** /v1/im/global-chat/messages/revoke | |
|[**revokePrivateChatMessage**](#revokeprivatechatmessage) | **POST** /v1/im/private-chat/messages/revoke | |
|[**sendGlobalChatMessage**](#sendglobalchatmessage) | **POST** /v1/im/global-chat/messages | |
|[**sendPrivateChatMessage**](#sendprivatechatmessage) | **POST** /v1/im/private-chat/messages | |
|[**setPrivateChatSessionMuted**](#setprivatechatsessionmuted) | **POST** /v1/im/private-chat/muted | |
|[**setPrivateChatSessionPinned**](#setprivatechatsessionpinned) | **POST** /v1/im/private-chat/pinned | |

# **countPrivateChatUnread**
> CountPrivateChatUnreadResp countPrivateChatUnread(body)

统计当前账号的私聊未读数。

### Example

```typescript
import {
    IMService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

let body: object; //

const { status, data } = await apiInstance.countPrivateChatUnread(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**CountPrivateChatUnreadResp**

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
    IMService,
    Configuration,
    ListGlobalChatMessagesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

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
    IMService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

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

# **listPrivateChatMessages**
> ListPrivateChatMessagesResp listPrivateChatMessages(listPrivateChatMessagesReq)

查询私聊消息。

### Example

```typescript
import {
    IMService,
    Configuration,
    ListPrivateChatMessagesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

let listPrivateChatMessagesReq: ListPrivateChatMessagesReq; //

const { status, data } = await apiInstance.listPrivateChatMessages(
    listPrivateChatMessagesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listPrivateChatMessagesReq** | **ListPrivateChatMessagesReq**|  | |


### Return type

**ListPrivateChatMessagesResp**

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

# **listPrivateChatSessions**
> ListPrivateChatSessionsResp listPrivateChatSessions()

查询当前账号的私聊会话。

### Example

```typescript
import {
    IMService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

let pagePage: string; //页码 (optional) (default to undefined)
let pageSize: string; //页大小 (optional) (default to undefined)

const { status, data } = await apiInstance.listPrivateChatSessions(
    pagePage,
    pageSize
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pagePage** | [**string**] | 页码 | (optional) defaults to undefined|
| **pageSize** | [**string**] | 页大小 | (optional) defaults to undefined|


### Return type

**ListPrivateChatSessionsResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **markPrivateChatSessionRead**
> object markPrivateChatSessionRead(markPrivateChatSessionReadReq)

标记私聊会话已读。

### Example

```typescript
import {
    IMService,
    Configuration,
    MarkPrivateChatSessionReadReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

let markPrivateChatSessionReadReq: MarkPrivateChatSessionReadReq; //

const { status, data } = await apiInstance.markPrivateChatSessionRead(
    markPrivateChatSessionReadReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **markPrivateChatSessionReadReq** | **MarkPrivateChatSessionReadReq**|  | |


### Return type

**object**

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
    IMService,
    Configuration,
    RevokeGlobalChatMessageReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

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

# **revokePrivateChatMessage**
> RevokePrivateChatMessageResp revokePrivateChatMessage(revokePrivateChatMessageReq)

撤回私聊消息。

### Example

```typescript
import {
    IMService,
    Configuration,
    RevokePrivateChatMessageReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

let revokePrivateChatMessageReq: RevokePrivateChatMessageReq; //

const { status, data } = await apiInstance.revokePrivateChatMessage(
    revokePrivateChatMessageReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **revokePrivateChatMessageReq** | **RevokePrivateChatMessageReq**|  | |


### Return type

**RevokePrivateChatMessageResp**

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
    IMService,
    Configuration,
    SendGlobalChatMessageReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

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

# **sendPrivateChatMessage**
> SendPrivateChatMessageResp sendPrivateChatMessage(sendPrivateChatMessageReq)

发送私聊消息。

### Example

```typescript
import {
    IMService,
    Configuration,
    SendPrivateChatMessageReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

let sendPrivateChatMessageReq: SendPrivateChatMessageReq; //

const { status, data } = await apiInstance.sendPrivateChatMessage(
    sendPrivateChatMessageReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **sendPrivateChatMessageReq** | **SendPrivateChatMessageReq**|  | |


### Return type

**SendPrivateChatMessageResp**

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

# **setPrivateChatSessionMuted**
> object setPrivateChatSessionMuted(setPrivateChatSessionMutedReq)

设置私聊会话免打扰。

### Example

```typescript
import {
    IMService,
    Configuration,
    SetPrivateChatSessionMutedReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

let setPrivateChatSessionMutedReq: SetPrivateChatSessionMutedReq; //

const { status, data } = await apiInstance.setPrivateChatSessionMuted(
    setPrivateChatSessionMutedReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **setPrivateChatSessionMutedReq** | **SetPrivateChatSessionMutedReq**|  | |


### Return type

**object**

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

# **setPrivateChatSessionPinned**
> object setPrivateChatSessionPinned(setPrivateChatSessionPinnedReq)

设置私聊会话置顶。

### Example

```typescript
import {
    IMService,
    Configuration,
    SetPrivateChatSessionPinnedReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new IMService(configuration);

let setPrivateChatSessionPinnedReq: SetPrivateChatSessionPinnedReq; //

const { status, data } = await apiInstance.setPrivateChatSessionPinned(
    setPrivateChatSessionPinnedReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **setPrivateChatSessionPinnedReq** | **SetPrivateChatSessionPinnedReq**|  | |


### Return type

**object**

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

