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
> object block(blockRelationRequest)

当前账号拉黑目标账号。

### Example

```typescript
import {
    RelationService,
    Configuration,
    BlockRelationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let blockRelationRequest: BlockRelationRequest; //

const { status, data } = await apiInstance.block(
    blockRelationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **blockRelationRequest** | **BlockRelationRequest**|  | |


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
> object follow(followRelationRequest)

当前账号关注目标账号。

### Example

```typescript
import {
    RelationService,
    Configuration,
    FollowRelationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let followRelationRequest: FollowRelationRequest; //

const { status, data } = await apiInstance.follow(
    followRelationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **followRelationRequest** | **FollowRelationRequest**|  | |


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
> GetStatusRelationReply getStatus(getStatusRelationRequest)

查询当前账号与目标账号之间的关系。

### Example

```typescript
import {
    RelationService,
    Configuration,
    GetStatusRelationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let getStatusRelationRequest: GetStatusRelationRequest; //

const { status, data } = await apiInstance.getStatus(
    getStatusRelationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **getStatusRelationRequest** | **GetStatusRelationRequest**|  | |


### Return type

**GetStatusRelationReply**

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
> ListBlockedRelationsReply listBlocked(listBlockedRelationsRequest)

分页查询当前账号拉黑的账号列表。

### Example

```typescript
import {
    RelationService,
    Configuration,
    ListBlockedRelationsRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let listBlockedRelationsRequest: ListBlockedRelationsRequest; //

const { status, data } = await apiInstance.listBlocked(
    listBlockedRelationsRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listBlockedRelationsRequest** | **ListBlockedRelationsRequest**|  | |


### Return type

**ListBlockedRelationsReply**

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
> ListFollowersRelationsReply listFollowers(listFollowersRelationsRequest)

分页查询当前账号的粉丝账号列表。

### Example

```typescript
import {
    RelationService,
    Configuration,
    ListFollowersRelationsRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let listFollowersRelationsRequest: ListFollowersRelationsRequest; //

const { status, data } = await apiInstance.listFollowers(
    listFollowersRelationsRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listFollowersRelationsRequest** | **ListFollowersRelationsRequest**|  | |


### Return type

**ListFollowersRelationsReply**

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
> ListFollowingRelationsReply listFollowing(listFollowingRelationsRequest)

分页查询当前账号关注的账号列表。

### Example

```typescript
import {
    RelationService,
    Configuration,
    ListFollowingRelationsRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let listFollowingRelationsRequest: ListFollowingRelationsRequest; //

const { status, data } = await apiInstance.listFollowing(
    listFollowingRelationsRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **listFollowingRelationsRequest** | **ListFollowingRelationsRequest**|  | |


### Return type

**ListFollowingRelationsReply**

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
> object unblock(unblockRelationRequest)

当前账号取消拉黑目标账号。

### Example

```typescript
import {
    RelationService,
    Configuration,
    UnblockRelationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let unblockRelationRequest: UnblockRelationRequest; //

const { status, data } = await apiInstance.unblock(
    unblockRelationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **unblockRelationRequest** | **UnblockRelationRequest**|  | |


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
> object unfollow(unfollowRelationRequest)

当前账号取消关注目标账号。

### Example

```typescript
import {
    RelationService,
    Configuration,
    UnfollowRelationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new RelationService(configuration);

let unfollowRelationRequest: UnfollowRelationRequest; //

const { status, data } = await apiInstance.unfollow(
    unfollowRelationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **unfollowRelationRequest** | **UnfollowRelationRequest**|  | |


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

