# IMService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**countPrivateChatUnread**](IMService.md#countprivatechatunread) | **POST** /v1/im/private-chat/unread-count |  |
| [**listGlobalChatMessages**](IMService.md#listglobalchatmessages) | **POST** /v1/im/global-chat/messages/list |  |
| [**listGlobalChatOnlineMembers**](IMService.md#listglobalchatonlinemembers) | **POST** /v1/im/global-chat/online-members/list |  |
| [**listPrivateChatMessages**](IMService.md#listprivatechatmessages) | **POST** /v1/im/private-chat/messages/list |  |
| [**listPrivateChatSessions**](IMService.md#listprivatechatsessions) | **GET** /v1/im/private-chat/sessions |  |
| [**markPrivateChatSessionRead**](IMService.md#markprivatechatsessionread) | **POST** /v1/im/private-chat/read |  |
| [**revokeGlobalChatMessage**](IMService.md#revokeglobalchatmessage) | **POST** /v1/im/global-chat/messages/revoke |  |
| [**revokePrivateChatMessage**](IMService.md#revokeprivatechatmessage) | **POST** /v1/im/private-chat/messages/revoke |  |
| [**sendGlobalChatMessage**](IMService.md#sendglobalchatmessage) | **POST** /v1/im/global-chat/messages |  |
| [**sendPrivateChatMessage**](IMService.md#sendprivatechatmessage) | **POST** /v1/im/private-chat/messages |  |
| [**setPrivateChatSessionMuted**](IMService.md#setprivatechatsessionmuted) | **POST** /v1/im/private-chat/muted |  |
| [**setPrivateChatSessionPinned**](IMService.md#setprivatechatsessionpinned) | **POST** /v1/im/private-chat/pinned |  |



## countPrivateChatUnread

> CountPrivateChatUnreadResp countPrivateChatUnread(body)



统计当前账号的私聊未读数。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { CountPrivateChatUnreadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // object
    body: Object,
  } satisfies CountPrivateChatUnreadRequest;

  try {
    const data = await api.countPrivateChatUnread(body);
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

[**CountPrivateChatUnreadResp**](CountPrivateChatUnreadResp.md)

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


## listGlobalChatMessages

> ListGlobalChatMessagesResp listGlobalChatMessages(listGlobalChatMessagesReq)



查询全局聊天室消息。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { ListGlobalChatMessagesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // ListGlobalChatMessagesReq
    listGlobalChatMessagesReq: ...,
  } satisfies ListGlobalChatMessagesRequest;

  try {
    const data = await api.listGlobalChatMessages(body);
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
| **listGlobalChatMessagesReq** | [ListGlobalChatMessagesReq](ListGlobalChatMessagesReq.md) |  | |

### Return type

[**ListGlobalChatMessagesResp**](ListGlobalChatMessagesResp.md)

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


## listGlobalChatOnlineMembers

> ListGlobalChatOnlineMembersResp listGlobalChatOnlineMembers(body)



查询全部在线的聊天室成员。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { ListGlobalChatOnlineMembersRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // object
    body: Object,
  } satisfies ListGlobalChatOnlineMembersRequest;

  try {
    const data = await api.listGlobalChatOnlineMembers(body);
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

[**ListGlobalChatOnlineMembersResp**](ListGlobalChatOnlineMembersResp.md)

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


## listPrivateChatMessages

> ListPrivateChatMessagesResp listPrivateChatMessages(listPrivateChatMessagesReq)



查询私聊消息。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { ListPrivateChatMessagesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // ListPrivateChatMessagesReq
    listPrivateChatMessagesReq: ...,
  } satisfies ListPrivateChatMessagesRequest;

  try {
    const data = await api.listPrivateChatMessages(body);
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
| **listPrivateChatMessagesReq** | [ListPrivateChatMessagesReq](ListPrivateChatMessagesReq.md) |  | |

### Return type

[**ListPrivateChatMessagesResp**](ListPrivateChatMessagesResp.md)

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


## listPrivateChatSessions

> ListPrivateChatSessionsResp listPrivateChatSessions(pagePage, pageSize)



查询当前账号的私聊会话。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { ListPrivateChatSessionsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // string | 页码 (optional)
    pagePage: pagePage_example,
    // string | 页大小 (optional)
    pageSize: pageSize_example,
  } satisfies ListPrivateChatSessionsRequest;

  try {
    const data = await api.listPrivateChatSessions(body);
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
| **pagePage** | `string` | 页码 | [Optional] [Defaults to `undefined`] |
| **pageSize** | `string` | 页大小 | [Optional] [Defaults to `undefined`] |

### Return type

[**ListPrivateChatSessionsResp**](ListPrivateChatSessionsResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## markPrivateChatSessionRead

> object markPrivateChatSessionRead(markPrivateChatSessionReadReq)



标记私聊会话已读。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { MarkPrivateChatSessionReadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // MarkPrivateChatSessionReadReq
    markPrivateChatSessionReadReq: ...,
  } satisfies MarkPrivateChatSessionReadRequest;

  try {
    const data = await api.markPrivateChatSessionRead(body);
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
| **markPrivateChatSessionReadReq** | [MarkPrivateChatSessionReadReq](MarkPrivateChatSessionReadReq.md) |  | |

### Return type

**object**

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


## revokeGlobalChatMessage

> RevokeGlobalChatMessageResp revokeGlobalChatMessage(revokeGlobalChatMessageReq)



撤回自己的全局聊天室消息。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { RevokeGlobalChatMessageRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // RevokeGlobalChatMessageReq
    revokeGlobalChatMessageReq: ...,
  } satisfies RevokeGlobalChatMessageRequest;

  try {
    const data = await api.revokeGlobalChatMessage(body);
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
| **revokeGlobalChatMessageReq** | [RevokeGlobalChatMessageReq](RevokeGlobalChatMessageReq.md) |  | |

### Return type

[**RevokeGlobalChatMessageResp**](RevokeGlobalChatMessageResp.md)

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


## revokePrivateChatMessage

> RevokePrivateChatMessageResp revokePrivateChatMessage(revokePrivateChatMessageReq)



撤回私聊消息。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { RevokePrivateChatMessageRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // RevokePrivateChatMessageReq
    revokePrivateChatMessageReq: ...,
  } satisfies RevokePrivateChatMessageRequest;

  try {
    const data = await api.revokePrivateChatMessage(body);
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
| **revokePrivateChatMessageReq** | [RevokePrivateChatMessageReq](RevokePrivateChatMessageReq.md) |  | |

### Return type

[**RevokePrivateChatMessageResp**](RevokePrivateChatMessageResp.md)

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


## sendGlobalChatMessage

> SendGlobalChatMessageResp sendGlobalChatMessage(sendGlobalChatMessageReq)



发送全局聊天室消息。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { SendGlobalChatMessageRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // SendGlobalChatMessageReq
    sendGlobalChatMessageReq: ...,
  } satisfies SendGlobalChatMessageRequest;

  try {
    const data = await api.sendGlobalChatMessage(body);
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
| **sendGlobalChatMessageReq** | [SendGlobalChatMessageReq](SendGlobalChatMessageReq.md) |  | |

### Return type

[**SendGlobalChatMessageResp**](SendGlobalChatMessageResp.md)

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


## sendPrivateChatMessage

> SendPrivateChatMessageResp sendPrivateChatMessage(sendPrivateChatMessageReq)



发送私聊消息。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { SendPrivateChatMessageRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // SendPrivateChatMessageReq
    sendPrivateChatMessageReq: ...,
  } satisfies SendPrivateChatMessageRequest;

  try {
    const data = await api.sendPrivateChatMessage(body);
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
| **sendPrivateChatMessageReq** | [SendPrivateChatMessageReq](SendPrivateChatMessageReq.md) |  | |

### Return type

[**SendPrivateChatMessageResp**](SendPrivateChatMessageResp.md)

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


## setPrivateChatSessionMuted

> object setPrivateChatSessionMuted(setPrivateChatSessionMutedReq)



设置私聊会话免打扰。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { SetPrivateChatSessionMutedRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // SetPrivateChatSessionMutedReq
    setPrivateChatSessionMutedReq: ...,
  } satisfies SetPrivateChatSessionMutedRequest;

  try {
    const data = await api.setPrivateChatSessionMuted(body);
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
| **setPrivateChatSessionMutedReq** | [SetPrivateChatSessionMutedReq](SetPrivateChatSessionMutedReq.md) |  | |

### Return type

**object**

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


## setPrivateChatSessionPinned

> object setPrivateChatSessionPinned(setPrivateChatSessionPinnedReq)



设置私聊会话置顶。

### Example

```ts
import {
  Configuration,
  IMService,
} from '@bass/bbs-sdk-fetch';
import type { SetPrivateChatSessionPinnedRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new IMService();

  const body = {
    // SetPrivateChatSessionPinnedReq
    setPrivateChatSessionPinnedReq: ...,
  } satisfies SetPrivateChatSessionPinnedRequest;

  try {
    const data = await api.setPrivateChatSessionPinned(body);
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
| **setPrivateChatSessionPinnedReq** | [SetPrivateChatSessionPinnedReq](SetPrivateChatSessionPinnedReq.md) |  | |

### Return type

**object**

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

