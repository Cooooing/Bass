# PrivateChatService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**countPrivateChatUnread**](PrivateChatService.md#countprivatechatunread) | **POST** /v1/im/private-chat/unread-count |  |
| [**listPrivateChatMessages**](PrivateChatService.md#listprivatechatmessages) | **POST** /v1/im/private-chat/messages/list |  |
| [**listPrivateChatSessions**](PrivateChatService.md#listprivatechatsessions) | **GET** /v1/im/private-chat/sessions |  |
| [**markPrivateChatSessionRead**](PrivateChatService.md#markprivatechatsessionread) | **POST** /v1/im/private-chat/read |  |
| [**revokePrivateChatMessage**](PrivateChatService.md#revokeprivatechatmessage) | **POST** /v1/im/private-chat/messages/revoke |  |
| [**sendPrivateChatMessage**](PrivateChatService.md#sendprivatechatmessage) | **POST** /v1/im/private-chat/messages |  |
| [**setPrivateChatSessionMuted**](PrivateChatService.md#setprivatechatsessionmuted) | **POST** /v1/im/private-chat/muted |  |
| [**setPrivateChatSessionPinned**](PrivateChatService.md#setprivatechatsessionpinned) | **POST** /v1/im/private-chat/pinned |  |



## countPrivateChatUnread

> CountPrivateChatUnreadResp countPrivateChatUnread(body)



统计当前账号的私聊未读数。

### Example

```ts
import {
  Configuration,
  PrivateChatService,
} from '@bass/bbs-sdk-fetch';
import type { CountPrivateChatUnreadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivateChatService();

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


## listPrivateChatMessages

> ListPrivateChatMessagesResp listPrivateChatMessages(listPrivateChatMessagesReq)



查询私聊消息。

### Example

```ts
import {
  Configuration,
  PrivateChatService,
} from '@bass/bbs-sdk-fetch';
import type { ListPrivateChatMessagesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivateChatService();

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
  PrivateChatService,
} from '@bass/bbs-sdk-fetch';
import type { ListPrivateChatSessionsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivateChatService();

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
  PrivateChatService,
} from '@bass/bbs-sdk-fetch';
import type { MarkPrivateChatSessionReadRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivateChatService();

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


## revokePrivateChatMessage

> RevokePrivateChatMessageResp revokePrivateChatMessage(revokePrivateChatMessageReq)



撤回私聊消息。

### Example

```ts
import {
  Configuration,
  PrivateChatService,
} from '@bass/bbs-sdk-fetch';
import type { RevokePrivateChatMessageRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivateChatService();

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


## sendPrivateChatMessage

> SendPrivateChatMessageResp sendPrivateChatMessage(sendPrivateChatMessageReq)



发送私聊消息。

### Example

```ts
import {
  Configuration,
  PrivateChatService,
} from '@bass/bbs-sdk-fetch';
import type { SendPrivateChatMessageRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivateChatService();

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
  PrivateChatService,
} from '@bass/bbs-sdk-fetch';
import type { SetPrivateChatSessionMutedRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivateChatService();

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
  PrivateChatService,
} from '@bass/bbs-sdk-fetch';
import type { SetPrivateChatSessionPinnedRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivateChatService();

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

