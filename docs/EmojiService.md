# EmojiService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addItem**](EmojiService.md#additem) | **POST** /v1/im/emoji/items |  |
| [**createGroup**](EmojiService.md#creategroup) | **POST** /v1/im/emoji/groups |  |
| [**createShare**](EmojiService.md#createshare) | **POST** /v1/im/emoji/shares |  |
| [**deleteGroup**](EmojiService.md#deletegroup) | **POST** /v1/im/emoji/groups/delete |  |
| [**deleteItem**](EmojiService.md#deleteitem) | **POST** /v1/im/emoji/items/delete |  |
| [**importShare**](EmojiService.md#importshare) | **POST** /v1/im/emoji/shares/import |  |
| [**list**](EmojiService.md#list) | **POST** /v1/im/emoji/groups/list |  |
| [**moveItem**](EmojiService.md#moveitem) | **POST** /v1/im/emoji/items/move |  |
| [**renameGroup**](EmojiService.md#renamegroup) | **POST** /v1/im/emoji/groups/rename |  |
| [**renameItem**](EmojiService.md#renameitem) | **POST** /v1/im/emoji/items/rename |  |
| [**reorderGroups**](EmojiService.md#reordergroups) | **POST** /v1/im/emoji/groups/reorder |  |
| [**reorderItems**](EmojiService.md#reorderitems) | **POST** /v1/im/emoji/items/reorder |  |



## addItem

> AddEmojiItemResp addItem(addEmojiItemReq)



向分组加入已上传图片。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { AddItemRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // AddEmojiItemReq
    addEmojiItemReq: ...,
  } satisfies AddItemRequest;

  try {
    const data = await api.addItem(body);
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
| **addEmojiItemReq** | [AddEmojiItemReq](AddEmojiItemReq.md) |  | |

### Return type

[**AddEmojiItemResp**](AddEmojiItemResp.md)

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


## createGroup

> CreateEmojiGroupResp createGroup(createEmojiGroupReq)



创建一个表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { CreateGroupRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // CreateEmojiGroupReq
    createEmojiGroupReq: ...,
  } satisfies CreateGroupRequest;

  try {
    const data = await api.createGroup(body);
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
| **createEmojiGroupReq** | [CreateEmojiGroupReq](CreateEmojiGroupReq.md) |  | |

### Return type

[**CreateEmojiGroupResp**](CreateEmojiGroupResp.md)

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


## createShare

> CreateEmojiShareResp createShare(createEmojiShareReq)



生成七天有效的分组分享码。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { CreateShareRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // CreateEmojiShareReq
    createEmojiShareReq: ...,
  } satisfies CreateShareRequest;

  try {
    const data = await api.createShare(body);
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
| **createEmojiShareReq** | [CreateEmojiShareReq](CreateEmojiShareReq.md) |  | |

### Return type

[**CreateEmojiShareResp**](CreateEmojiShareResp.md)

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


## deleteGroup

> object deleteGroup(deleteEmojiGroupReq)



删除一个表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { DeleteGroupRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // DeleteEmojiGroupReq
    deleteEmojiGroupReq: ...,
  } satisfies DeleteGroupRequest;

  try {
    const data = await api.deleteGroup(body);
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
| **deleteEmojiGroupReq** | [DeleteEmojiGroupReq](DeleteEmojiGroupReq.md) |  | |

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


## deleteItem

> object deleteItem(deleteEmojiItemReq)



从分组移除表情。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { DeleteItemRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // DeleteEmojiItemReq
    deleteEmojiItemReq: ...,
  } satisfies DeleteItemRequest;

  try {
    const data = await api.deleteItem(body);
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
| **deleteEmojiItemReq** | [DeleteEmojiItemReq](DeleteEmojiItemReq.md) |  | |

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


## importShare

> ImportEmojiShareResp importShare(importEmojiShareReq)



把分享码导入为一个新分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { ImportShareRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // ImportEmojiShareReq
    importEmojiShareReq: ...,
  } satisfies ImportShareRequest;

  try {
    const data = await api.importShare(body);
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
| **importEmojiShareReq** | [ImportEmojiShareReq](ImportEmojiShareReq.md) |  | |

### Return type

[**ImportEmojiShareResp**](ImportEmojiShareResp.md)

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


## list

> ListEmojiGroupsResp list(listEmojiGroupsReq)



查询当前账号的表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { ListRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // ListEmojiGroupsReq
    listEmojiGroupsReq: ...,
  } satisfies ListRequest;

  try {
    const data = await api.list(body);
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
| **listEmojiGroupsReq** | [ListEmojiGroupsReq](ListEmojiGroupsReq.md) |  | |

### Return type

[**ListEmojiGroupsResp**](ListEmojiGroupsResp.md)

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


## moveItem

> MoveEmojiItemResp moveItem(moveEmojiItemReq)



将表情移动到另一个当前账号分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { MoveItemRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // MoveEmojiItemReq
    moveEmojiItemReq: ...,
  } satisfies MoveItemRequest;

  try {
    const data = await api.moveItem(body);
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
| **moveEmojiItemReq** | [MoveEmojiItemReq](MoveEmojiItemReq.md) |  | |

### Return type

[**MoveEmojiItemResp**](MoveEmojiItemResp.md)

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


## renameGroup

> RenameEmojiGroupResp renameGroup(renameEmojiGroupReq)



重命名一个表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { RenameGroupRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // RenameEmojiGroupReq
    renameEmojiGroupReq: ...,
  } satisfies RenameGroupRequest;

  try {
    const data = await api.renameGroup(body);
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
| **renameEmojiGroupReq** | [RenameEmojiGroupReq](RenameEmojiGroupReq.md) |  | |

### Return type

[**RenameEmojiGroupResp**](RenameEmojiGroupResp.md)

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


## renameItem

> RenameEmojiItemResp renameItem(renameEmojiItemReq)



重命名一个表情别名。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { RenameItemRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // RenameEmojiItemReq
    renameEmojiItemReq: ...,
  } satisfies RenameItemRequest;

  try {
    const data = await api.renameItem(body);
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
| **renameEmojiItemReq** | [RenameEmojiItemReq](RenameEmojiItemReq.md) |  | |

### Return type

[**RenameEmojiItemResp**](RenameEmojiItemResp.md)

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


## reorderGroups

> object reorderGroups(reorderEmojiGroupsReq)



重排表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { ReorderGroupsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // ReorderEmojiGroupsReq
    reorderEmojiGroupsReq: ...,
  } satisfies ReorderGroupsRequest;

  try {
    const data = await api.reorderGroups(body);
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
| **reorderEmojiGroupsReq** | [ReorderEmojiGroupsReq](ReorderEmojiGroupsReq.md) |  | |

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


## reorderItems

> object reorderItems(reorderEmojiItemsReq)



重排同一分组内的表情。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { ReorderItemsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // ReorderEmojiItemsReq
    reorderEmojiItemsReq: ...,
  } satisfies ReorderItemsRequest;

  try {
    const data = await api.reorderItems(body);
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
| **reorderEmojiItemsReq** | [ReorderEmojiItemsReq](ReorderEmojiItemsReq.md) |  | |

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

