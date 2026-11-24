# EmojiService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**addEmojiItem**](#addemojiitem) | **POST** /v1/im/emoji/items | |
|[**createEmojiGroup**](#createemojigroup) | **POST** /v1/im/emoji/groups | |
|[**createEmojiShare**](#createemojishare) | **POST** /v1/im/emoji/shares | |
|[**deleteEmojiGroup**](#deleteemojigroup) | **POST** /v1/im/emoji/groups/delete | |
|[**deleteEmojiItem**](#deleteemojiitem) | **POST** /v1/im/emoji/items/delete | |
|[**importEmojiShare**](#importemojishare) | **POST** /v1/im/emoji/shares/import | |
|[**listEmojiGroups**](#listemojigroups) | **POST** /v1/im/emoji/groups/list | |
|[**moveEmojiItem**](#moveemojiitem) | **POST** /v1/im/emoji/items/move | |
|[**renameEmojiGroup**](#renameemojigroup) | **POST** /v1/im/emoji/groups/rename | |
|[**renameEmojiItem**](#renameemojiitem) | **POST** /v1/im/emoji/items/rename | |
|[**reorderEmojiGroups**](#reorderemojigroups) | **POST** /v1/im/emoji/groups/reorder | |
|[**reorderEmojiItems**](#reorderemojiitems) | **POST** /v1/im/emoji/items/reorder | |

# **addEmojiItem**
> AddEmojiItemResp addEmojiItem(addEmojiItemReq)

向分组加入已上传图片。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    AddEmojiItemReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let addEmojiItemReq: AddEmojiItemReq; //

const { status, data } = await apiInstance.addEmojiItem(
    addEmojiItemReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **addEmojiItemReq** | **AddEmojiItemReq**|  | |


### Return type

**AddEmojiItemResp**

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

# **createEmojiGroup**
> CreateEmojiGroupResp createEmojiGroup(createEmojiGroupReq)

创建一个表情分组。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    CreateEmojiGroupReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let createEmojiGroupReq: CreateEmojiGroupReq; //

const { status, data } = await apiInstance.createEmojiGroup(
    createEmojiGroupReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createEmojiGroupReq** | **CreateEmojiGroupReq**|  | |


### Return type

**CreateEmojiGroupResp**

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

# **createEmojiShare**
> CreateEmojiShareResp createEmojiShare(createEmojiShareReq)

生成七天有效的分组分享码。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    CreateEmojiShareReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let createEmojiShareReq: CreateEmojiShareReq; //

const { status, data } = await apiInstance.createEmojiShare(
    createEmojiShareReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createEmojiShareReq** | **CreateEmojiShareReq**|  | |


### Return type

**CreateEmojiShareResp**

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

# **deleteEmojiGroup**
> object deleteEmojiGroup(deleteEmojiGroupReq)

删除一个表情分组。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    DeleteEmojiGroupReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let deleteEmojiGroupReq: DeleteEmojiGroupReq; //

const { status, data } = await apiInstance.deleteEmojiGroup(
    deleteEmojiGroupReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **deleteEmojiGroupReq** | **DeleteEmojiGroupReq**|  | |


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

# **deleteEmojiItem**
> object deleteEmojiItem(deleteEmojiItemReq)

从分组移除表情。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    DeleteEmojiItemReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let deleteEmojiItemReq: DeleteEmojiItemReq; //

const { status, data } = await apiInstance.deleteEmojiItem(
    deleteEmojiItemReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **deleteEmojiItemReq** | **DeleteEmojiItemReq**|  | |


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

# **importEmojiShare**
> ImportEmojiShareResp importEmojiShare(importEmojiShareReq)

把分享码导入为一个新分组。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    ImportEmojiShareReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let importEmojiShareReq: ImportEmojiShareReq; //

const { status, data } = await apiInstance.importEmojiShare(
    importEmojiShareReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **importEmojiShareReq** | **ImportEmojiShareReq**|  | |


### Return type

**ImportEmojiShareResp**

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

# **listEmojiGroups**
> ListEmojiGroupsResp listEmojiGroups(listEmojiGroupsReq)

查询当前账号的表情分组。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    ListEmojiGroupsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let listEmojiGroupsReq: ListEmojiGroupsReq; //

const { status, data } = await apiInstance.listEmojiGroups(
    listEmojiGroupsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listEmojiGroupsReq** | **ListEmojiGroupsReq**|  | |


### Return type

**ListEmojiGroupsResp**

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

# **moveEmojiItem**
> MoveEmojiItemResp moveEmojiItem(moveEmojiItemReq)

将表情移动到另一个当前账号分组。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    MoveEmojiItemReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let moveEmojiItemReq: MoveEmojiItemReq; //

const { status, data } = await apiInstance.moveEmojiItem(
    moveEmojiItemReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **moveEmojiItemReq** | **MoveEmojiItemReq**|  | |


### Return type

**MoveEmojiItemResp**

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

# **renameEmojiGroup**
> RenameEmojiGroupResp renameEmojiGroup(renameEmojiGroupReq)

重命名一个表情分组。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    RenameEmojiGroupReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let renameEmojiGroupReq: RenameEmojiGroupReq; //

const { status, data } = await apiInstance.renameEmojiGroup(
    renameEmojiGroupReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **renameEmojiGroupReq** | **RenameEmojiGroupReq**|  | |


### Return type

**RenameEmojiGroupResp**

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

# **renameEmojiItem**
> RenameEmojiItemResp renameEmojiItem(renameEmojiItemReq)

重命名一个表情别名。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    RenameEmojiItemReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let renameEmojiItemReq: RenameEmojiItemReq; //

const { status, data } = await apiInstance.renameEmojiItem(
    renameEmojiItemReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **renameEmojiItemReq** | **RenameEmojiItemReq**|  | |


### Return type

**RenameEmojiItemResp**

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

# **reorderEmojiGroups**
> object reorderEmojiGroups(reorderEmojiGroupsReq)

重排表情分组。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    ReorderEmojiGroupsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let reorderEmojiGroupsReq: ReorderEmojiGroupsReq; //

const { status, data } = await apiInstance.reorderEmojiGroups(
    reorderEmojiGroupsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **reorderEmojiGroupsReq** | **ReorderEmojiGroupsReq**|  | |


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

# **reorderEmojiItems**
> object reorderEmojiItems(reorderEmojiItemsReq)

重排同一分组内的表情。

### Example

```typescript
import {
    EmojiService,
    Configuration,
    ReorderEmojiItemsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new EmojiService(configuration);

let reorderEmojiItemsReq: ReorderEmojiItemsReq; //

const { status, data } = await apiInstance.reorderEmojiItems(
    reorderEmojiItemsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **reorderEmojiItemsReq** | **ReorderEmojiItemsReq**|  | |


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

