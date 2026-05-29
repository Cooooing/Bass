# RelationServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**relationServiceBlock**](RelationServiceApi.md#relationserviceblock) | **POST** /v1/user/relation/block |  |
| [**relationServiceFollow**](RelationServiceApi.md#relationservicefollow) | **POST** /v1/user/relation/follow |  |
| [**relationServiceGetStatus**](RelationServiceApi.md#relationservicegetstatus) | **POST** /v1/user/relation/get-status |  |
| [**relationServiceListBlocked**](RelationServiceApi.md#relationservicelistblocked) | **POST** /v1/user/relation/list-blocked |  |
| [**relationServiceListFollowers**](RelationServiceApi.md#relationservicelistfollowers) | **POST** /v1/user/relation/list-followers |  |
| [**relationServiceListFollowing**](RelationServiceApi.md#relationservicelistfollowing) | **POST** /v1/user/relation/list-following |  |
| [**relationServiceUnblock**](RelationServiceApi.md#relationserviceunblock) | **POST** /v1/user/relation/unblock |  |
| [**relationServiceUnfollow**](RelationServiceApi.md#relationserviceunfollow) | **POST** /v1/user/relation/unfollow |  |



## relationServiceBlock

> object relationServiceBlock(blockRelationRequest)



当前账号拉黑目标账号

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceBlockRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // BlockRelationRequest
    blockRelationRequest: ...,
  } satisfies RelationServiceBlockRequest;

  try {
    const data = await api.relationServiceBlock(body);
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
| **blockRelationRequest** | [BlockRelationRequest](BlockRelationRequest.md) |  | |

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


## relationServiceFollow

> object relationServiceFollow(followRelationRequest)



当前账号关注目标账号

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceFollowRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // FollowRelationRequest
    followRelationRequest: ...,
  } satisfies RelationServiceFollowRequest;

  try {
    const data = await api.relationServiceFollow(body);
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
| **followRelationRequest** | [FollowRelationRequest](FollowRelationRequest.md) |  | |

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


## relationServiceGetStatus

> GetStatusRelationReply relationServiceGetStatus(getStatusRelationRequest)



查询当前账号与目标账号之间的关系

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceGetStatusRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // GetStatusRelationRequest
    getStatusRelationRequest: ...,
  } satisfies RelationServiceGetStatusRequest;

  try {
    const data = await api.relationServiceGetStatus(body);
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
| **getStatusRelationRequest** | [GetStatusRelationRequest](GetStatusRelationRequest.md) |  | |

### Return type

[**GetStatusRelationReply**](GetStatusRelationReply.md)

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


## relationServiceListBlocked

> ListBlockedRelationsReply relationServiceListBlocked(listBlockedRelationsRequest)



分页查询当前账号拉黑的账号列表

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceListBlockedRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // ListBlockedRelationsRequest
    listBlockedRelationsRequest: ...,
  } satisfies RelationServiceListBlockedRequest;

  try {
    const data = await api.relationServiceListBlocked(body);
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
| **listBlockedRelationsRequest** | [ListBlockedRelationsRequest](ListBlockedRelationsRequest.md) |  | |

### Return type

[**ListBlockedRelationsReply**](ListBlockedRelationsReply.md)

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


## relationServiceListFollowers

> ListFollowersRelationsReply relationServiceListFollowers(listFollowersRelationsRequest)



分页查询当前账号的粉丝账号列表

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceListFollowersRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // ListFollowersRelationsRequest
    listFollowersRelationsRequest: ...,
  } satisfies RelationServiceListFollowersRequest;

  try {
    const data = await api.relationServiceListFollowers(body);
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
| **listFollowersRelationsRequest** | [ListFollowersRelationsRequest](ListFollowersRelationsRequest.md) |  | |

### Return type

[**ListFollowersRelationsReply**](ListFollowersRelationsReply.md)

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


## relationServiceListFollowing

> ListFollowingRelationsReply relationServiceListFollowing(listFollowingRelationsRequest)



分页查询当前账号关注的账号列表

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceListFollowingRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // ListFollowingRelationsRequest
    listFollowingRelationsRequest: ...,
  } satisfies RelationServiceListFollowingRequest;

  try {
    const data = await api.relationServiceListFollowing(body);
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
| **listFollowingRelationsRequest** | [ListFollowingRelationsRequest](ListFollowingRelationsRequest.md) |  | |

### Return type

[**ListFollowingRelationsReply**](ListFollowingRelationsReply.md)

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


## relationServiceUnblock

> object relationServiceUnblock(unblockRelationRequest)



当前账号取消拉黑目标账号

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceUnblockRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // UnblockRelationRequest
    unblockRelationRequest: ...,
  } satisfies RelationServiceUnblockRequest;

  try {
    const data = await api.relationServiceUnblock(body);
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
| **unblockRelationRequest** | [UnblockRelationRequest](UnblockRelationRequest.md) |  | |

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


## relationServiceUnfollow

> object relationServiceUnfollow(unfollowRelationRequest)



当前账号取消关注目标账号

### Example

```ts
import {
  Configuration,
  RelationServiceApi,
} from '@bass/bbs-sdk';
import type { RelationServiceUnfollowRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new RelationServiceApi();

  const body = {
    // UnfollowRelationRequest
    unfollowRelationRequest: ...,
  } satisfies RelationServiceUnfollowRequest;

  try {
    const data = await api.relationServiceUnfollow(body);
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
| **unfollowRelationRequest** | [UnfollowRelationRequest](UnfollowRelationRequest.md) |  | |

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

