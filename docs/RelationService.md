# RelationService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**block**](#block) | **POST** /v1/user/relation/block | |
|[**follow**](#follow) | **POST** /v1/user/relation/follow | |
|[**getStatus**](#getstatus) | **POST** /v1/user/relation/get-status | |
|[**listBlocked**](#listblocked) | **POST** /v1/user/relation/list-blocked | |
|[**listFollowers**](#listfollowers) | **POST** /v1/user/relation/list-followers | |
|[**listFollowing**](#listfollowing) | **POST** /v1/user/relation/list-following | |
|[**unblock**](#unblock) | **POST** /v1/user/relation/unblock | |
|[**unfollow**](#unfollow) | **POST** /v1/user/relation/unfollow | |

# **block**
> object block(blockRelationReq)

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

const { status, data } = await apiInstance.block(
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

# **follow**
> object follow(followRelationReq)

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

const { status, data } = await apiInstance.follow(
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

# **getStatus**
> GetStatusRelationResp getStatus(getStatusRelationReq)

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

const { status, data } = await apiInstance.getStatus(
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

# **listBlocked**
> ListBlockedRelationsResp listBlocked(listBlockedRelationsReq)

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

const { status, data } = await apiInstance.listBlocked(
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

# **listFollowers**
> ListFollowersRelationsResp listFollowers(listFollowersRelationsReq)

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

const { status, data } = await apiInstance.listFollowers(
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

# **listFollowing**
> ListFollowingRelationsResp listFollowing(listFollowingRelationsReq)

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

const { status, data } = await apiInstance.listFollowing(
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

# **unblock**
> object unblock(unblockRelationReq)

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

const { status, data } = await apiInstance.unblock(
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

# **unfollow**
> object unfollow(unfollowRelationReq)

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

const { status, data } = await apiInstance.unfollow(
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

