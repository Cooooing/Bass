# RelationService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**blockRelation**](#blockrelation) | **POST** /v1/user/relation/block | |
|[**followRelation**](#followrelation) | **POST** /v1/user/relation/follow | |
|[**getStatusRelation**](#getstatusrelation) | **POST** /v1/user/relation/get-status | |
|[**listBlockedRelations**](#listblockedrelations) | **POST** /v1/user/relation/list-blocked | |
|[**listFollowersRelations**](#listfollowersrelations) | **POST** /v1/user/relation/list-followers | |
|[**listFollowingRelations**](#listfollowingrelations) | **POST** /v1/user/relation/list-following | |
|[**unblockRelation**](#unblockrelation) | **POST** /v1/user/relation/unblock | |
|[**unfollowRelation**](#unfollowrelation) | **POST** /v1/user/relation/unfollow | |

# **blockRelation**
> object blockRelation(blockRelationReq)

当前账号拉黑目标账号。

### Example

```typescript
import {
    RelationService,
    Configuration,
    BlockRelationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let blockRelationReq: BlockRelationReq; //

const { status, data } = await apiInstance.blockRelation(
    blockRelationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **blockRelationReq** | **BlockRelationReq**|  | |


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

# **followRelation**
> object followRelation(followRelationReq)

当前账号关注目标账号。

### Example

```typescript
import {
    RelationService,
    Configuration,
    FollowRelationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let followRelationReq: FollowRelationReq; //

const { status, data } = await apiInstance.followRelation(
    followRelationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **followRelationReq** | **FollowRelationReq**|  | |


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

# **getStatusRelation**
> GetStatusRelationResp getStatusRelation(getStatusRelationReq)

查询当前账号与目标账号之间的关系。

### Example

```typescript
import {
    RelationService,
    Configuration,
    GetStatusRelationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let getStatusRelationReq: GetStatusRelationReq; //

const { status, data } = await apiInstance.getStatusRelation(
    getStatusRelationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **getStatusRelationReq** | **GetStatusRelationReq**|  | |


### Return type

**GetStatusRelationResp**

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

# **listBlockedRelations**
> ListBlockedRelationsResp listBlockedRelations(listBlockedRelationsReq)

分页查询当前账号拉黑的账号列表。

### Example

```typescript
import {
    RelationService,
    Configuration,
    ListBlockedRelationsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let listBlockedRelationsReq: ListBlockedRelationsReq; //

const { status, data } = await apiInstance.listBlockedRelations(
    listBlockedRelationsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listBlockedRelationsReq** | **ListBlockedRelationsReq**|  | |


### Return type

**ListBlockedRelationsResp**

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

# **listFollowersRelations**
> ListFollowersRelationsResp listFollowersRelations(listFollowersRelationsReq)

分页查询当前账号的粉丝账号列表。

### Example

```typescript
import {
    RelationService,
    Configuration,
    ListFollowersRelationsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let listFollowersRelationsReq: ListFollowersRelationsReq; //

const { status, data } = await apiInstance.listFollowersRelations(
    listFollowersRelationsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listFollowersRelationsReq** | **ListFollowersRelationsReq**|  | |


### Return type

**ListFollowersRelationsResp**

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

# **listFollowingRelations**
> ListFollowingRelationsResp listFollowingRelations(listFollowingRelationsReq)

分页查询当前账号关注的账号列表。

### Example

```typescript
import {
    RelationService,
    Configuration,
    ListFollowingRelationsReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let listFollowingRelationsReq: ListFollowingRelationsReq; //

const { status, data } = await apiInstance.listFollowingRelations(
    listFollowingRelationsReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listFollowingRelationsReq** | **ListFollowingRelationsReq**|  | |


### Return type

**ListFollowingRelationsResp**

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

# **unblockRelation**
> object unblockRelation(unblockRelationReq)

当前账号取消拉黑目标账号。

### Example

```typescript
import {
    RelationService,
    Configuration,
    UnblockRelationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let unblockRelationReq: UnblockRelationReq; //

const { status, data } = await apiInstance.unblockRelation(
    unblockRelationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **unblockRelationReq** | **UnblockRelationReq**|  | |


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

# **unfollowRelation**
> object unfollowRelation(unfollowRelationReq)

当前账号取消关注目标账号。

### Example

```typescript
import {
    RelationService,
    Configuration,
    UnfollowRelationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let unfollowRelationReq: UnfollowRelationReq; //

const { status, data } = await apiInstance.unfollowRelation(
    unfollowRelationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **unfollowRelationReq** | **UnfollowRelationReq**|  | |


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

