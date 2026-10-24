# GlobalChatService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getGlobalChatTopic**](GlobalChatService.md#getglobalchattopic) | **POST** /v1/im/global-chat/topic/get |  |
| [**listGlobalChatBarrages**](GlobalChatService.md#listglobalchatbarrages) | **POST** /v1/im/global-chat/barrages/list |  |
| [**listGlobalChatMessages**](GlobalChatService.md#listglobalchatmessages) | **POST** /v1/im/global-chat/messages/list |  |
| [**listGlobalChatOnlineMembers**](GlobalChatService.md#listglobalchatonlinemembers) | **POST** /v1/im/global-chat/online-members/list |  |
| [**revokeGlobalChatMessage**](GlobalChatService.md#revokeglobalchatmessage) | **POST** /v1/im/global-chat/messages/revoke |  |
| [**sendGlobalChatBarrage**](GlobalChatService.md#sendglobalchatbarrage) | **POST** /v1/im/global-chat/barrages |  |
| [**sendGlobalChatMessage**](GlobalChatService.md#sendglobalchatmessage) | **POST** /v1/im/global-chat/messages |  |
| [**setGlobalChatTopic**](GlobalChatService.md#setglobalchattopic) | **POST** /v1/im/global-chat/topic/set |  |



## getGlobalChatTopic

> GetGlobalChatTopicResp getGlobalChatTopic(body)



查询当前全局聊天室话题。

### Example

```ts
import {
  Configuration,
  GlobalChatService,
} from '@bass/bbs-sdk-fetch';
import type { GetGlobalChatTopicRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new GlobalChatService();

  const body = {
    // object
    body: Object,
  } satisfies GetGlobalChatTopicRequest;

  try {
    const data = await api.getGlobalChatTopic(body);
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

[**GetGlobalChatTopicResp**](GetGlobalChatTopicResp.md)

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


## listGlobalChatBarrages

> ListGlobalChatBarragesResp listGlobalChatBarrages(listGlobalChatBarragesReq)



分页查询已发布的全局聊天室弹幕。

### Example

```ts
import {
  Configuration,
  GlobalChatService,
} from '@bass/bbs-sdk-fetch';
import type { ListGlobalChatBarragesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new GlobalChatService();

  const body = {
    // ListGlobalChatBarragesReq
    listGlobalChatBarragesReq: ...,
  } satisfies ListGlobalChatBarragesRequest;

  try {
    const data = await api.listGlobalChatBarrages(body);
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
| **listGlobalChatBarragesReq** | [ListGlobalChatBarragesReq](ListGlobalChatBarragesReq.md) |  | |

### Return type

[**ListGlobalChatBarragesResp**](ListGlobalChatBarragesResp.md)

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
  GlobalChatService,
} from '@bass/bbs-sdk-fetch';
import type { ListGlobalChatMessagesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new GlobalChatService();

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
  GlobalChatService,
} from '@bass/bbs-sdk-fetch';
import type { ListGlobalChatOnlineMembersRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new GlobalChatService();

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


## revokeGlobalChatMessage

> RevokeGlobalChatMessageResp revokeGlobalChatMessage(revokeGlobalChatMessageReq)



撤回自己的全局聊天室消息。

### Example

```ts
import {
  Configuration,
  GlobalChatService,
} from '@bass/bbs-sdk-fetch';
import type { RevokeGlobalChatMessageRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new GlobalChatService();

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


## sendGlobalChatBarrage

> SendGlobalChatBarrageResp sendGlobalChatBarrage(sendGlobalChatBarrageReq)



发送全局聊天室弹幕。

### Example

```ts
import {
  Configuration,
  GlobalChatService,
} from '@bass/bbs-sdk-fetch';
import type { SendGlobalChatBarrageRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new GlobalChatService();

  const body = {
    // SendGlobalChatBarrageReq
    sendGlobalChatBarrageReq: ...,
  } satisfies SendGlobalChatBarrageRequest;

  try {
    const data = await api.sendGlobalChatBarrage(body);
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
| **sendGlobalChatBarrageReq** | [SendGlobalChatBarrageReq](SendGlobalChatBarrageReq.md) |  | |

### Return type

[**SendGlobalChatBarrageResp**](SendGlobalChatBarrageResp.md)

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
  GlobalChatService,
} from '@bass/bbs-sdk-fetch';
import type { SendGlobalChatMessageRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new GlobalChatService();

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


## setGlobalChatTopic

> SetGlobalChatTopicResp setGlobalChatTopic(setGlobalChatTopicReq)



设置全局聊天室话题。

### Example

```ts
import {
  Configuration,
  GlobalChatService,
} from '@bass/bbs-sdk-fetch';
import type { SetGlobalChatTopicRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new GlobalChatService();

  const body = {
    // SetGlobalChatTopicReq
    setGlobalChatTopicReq: ...,
  } satisfies SetGlobalChatTopicRequest;

  try {
    const data = await api.setGlobalChatTopic(body);
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
| **setGlobalChatTopicReq** | [SetGlobalChatTopicReq](SetGlobalChatTopicReq.md) |  | |

### Return type

[**SetGlobalChatTopicResp**](SetGlobalChatTopicResp.md)

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

