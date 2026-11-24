# EmojiService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addEmojiItem**](EmojiService.md#addemojiitem) | **POST** /v1/im/emoji/items |  |
| [**createEmojiGroup**](EmojiService.md#createemojigroup) | **POST** /v1/im/emoji/groups |  |
| [**createEmojiShare**](EmojiService.md#createemojishare) | **POST** /v1/im/emoji/shares |  |
| [**deleteEmojiGroup**](EmojiService.md#deleteemojigroup) | **POST** /v1/im/emoji/groups/delete |  |
| [**deleteEmojiItem**](EmojiService.md#deleteemojiitem) | **POST** /v1/im/emoji/items/delete |  |
| [**importEmojiShare**](EmojiService.md#importemojishare) | **POST** /v1/im/emoji/shares/import |  |
| [**listEmojiGroups**](EmojiService.md#listemojigroups) | **POST** /v1/im/emoji/groups/list |  |
| [**moveEmojiItem**](EmojiService.md#moveemojiitem) | **POST** /v1/im/emoji/items/move |  |
| [**renameEmojiGroup**](EmojiService.md#renameemojigroup) | **POST** /v1/im/emoji/groups/rename |  |
| [**renameEmojiItem**](EmojiService.md#renameemojiitem) | **POST** /v1/im/emoji/items/rename |  |
| [**reorderEmojiGroups**](EmojiService.md#reorderemojigroups) | **POST** /v1/im/emoji/groups/reorder |  |
| [**reorderEmojiItems**](EmojiService.md#reorderemojiitems) | **POST** /v1/im/emoji/items/reorder |  |



## addEmojiItem

> AddEmojiItemResp addEmojiItem(addEmojiItemReq)



向分组加入已上传图片。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { AddEmojiItemRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // AddEmojiItemReq
    addEmojiItemReq: ...,
  } satisfies AddEmojiItemRequest;

  try {
    const data = await api.addEmojiItem(body);
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


## createEmojiGroup

> CreateEmojiGroupResp createEmojiGroup(createEmojiGroupReq)



创建一个表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { CreateEmojiGroupRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // CreateEmojiGroupReq
    createEmojiGroupReq: ...,
  } satisfies CreateEmojiGroupRequest;

  try {
    const data = await api.createEmojiGroup(body);
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


## createEmojiShare

> CreateEmojiShareResp createEmojiShare(createEmojiShareReq)



生成七天有效的分组分享码。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { CreateEmojiShareRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // CreateEmojiShareReq
    createEmojiShareReq: ...,
  } satisfies CreateEmojiShareRequest;

  try {
    const data = await api.createEmojiShare(body);
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


## deleteEmojiGroup

> object deleteEmojiGroup(deleteEmojiGroupReq)



删除一个表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { DeleteEmojiGroupRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // DeleteEmojiGroupReq
    deleteEmojiGroupReq: ...,
  } satisfies DeleteEmojiGroupRequest;

  try {
    const data = await api.deleteEmojiGroup(body);
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


## deleteEmojiItem

> object deleteEmojiItem(deleteEmojiItemReq)



从分组移除表情。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { DeleteEmojiItemRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // DeleteEmojiItemReq
    deleteEmojiItemReq: ...,
  } satisfies DeleteEmojiItemRequest;

  try {
    const data = await api.deleteEmojiItem(body);
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


## importEmojiShare

> ImportEmojiShareResp importEmojiShare(importEmojiShareReq)



把分享码导入为一个新分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { ImportEmojiShareRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // ImportEmojiShareReq
    importEmojiShareReq: ...,
  } satisfies ImportEmojiShareRequest;

  try {
    const data = await api.importEmojiShare(body);
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


## listEmojiGroups

> ListEmojiGroupsResp listEmojiGroups(listEmojiGroupsReq)



查询当前账号的表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { ListEmojiGroupsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // ListEmojiGroupsReq
    listEmojiGroupsReq: ...,
  } satisfies ListEmojiGroupsRequest;

  try {
    const data = await api.listEmojiGroups(body);
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


## moveEmojiItem

> MoveEmojiItemResp moveEmojiItem(moveEmojiItemReq)



将表情移动到另一个当前账号分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { MoveEmojiItemRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // MoveEmojiItemReq
    moveEmojiItemReq: ...,
  } satisfies MoveEmojiItemRequest;

  try {
    const data = await api.moveEmojiItem(body);
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


## renameEmojiGroup

> RenameEmojiGroupResp renameEmojiGroup(renameEmojiGroupReq)



重命名一个表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { RenameEmojiGroupRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // RenameEmojiGroupReq
    renameEmojiGroupReq: ...,
  } satisfies RenameEmojiGroupRequest;

  try {
    const data = await api.renameEmojiGroup(body);
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


## renameEmojiItem

> RenameEmojiItemResp renameEmojiItem(renameEmojiItemReq)



重命名一个表情别名。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { RenameEmojiItemRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // RenameEmojiItemReq
    renameEmojiItemReq: ...,
  } satisfies RenameEmojiItemRequest;

  try {
    const data = await api.renameEmojiItem(body);
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


## reorderEmojiGroups

> object reorderEmojiGroups(reorderEmojiGroupsReq)



重排表情分组。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { ReorderEmojiGroupsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // ReorderEmojiGroupsReq
    reorderEmojiGroupsReq: ...,
  } satisfies ReorderEmojiGroupsRequest;

  try {
    const data = await api.reorderEmojiGroups(body);
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


## reorderEmojiItems

> object reorderEmojiItems(reorderEmojiItemsReq)



重排同一分组内的表情。

### Example

```ts
import {
  Configuration,
  EmojiService,
} from '@bass/bbs-sdk-fetch';
import type { ReorderEmojiItemsRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new EmojiService();

  const body = {
    // ReorderEmojiItemsReq
    reorderEmojiItemsReq: ...,
  } satisfies ReorderEmojiItemsRequest;

  try {
    const data = await api.reorderEmojiItems(body);
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

