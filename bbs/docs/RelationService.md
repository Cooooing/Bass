# RelationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**block**](RelationService.md#block) | **POST** /v1/user/relation/block |  |
| [**follow**](RelationService.md#follow) | **POST** /v1/user/relation/follow |  |
| [**getStatus**](RelationService.md#getstatus) | **POST** /v1/user/relation/get-status |  |
| [**listBlocked**](RelationService.md#listblocked) | **POST** /v1/user/relation/list-blocked |  |
| [**listFollowers**](RelationService.md#listfollowers) | **POST** /v1/user/relation/list-followers |  |
| [**listFollowing**](RelationService.md#listfollowing) | **POST** /v1/user/relation/list-following |  |
| [**unblock**](RelationService.md#unblock) | **POST** /v1/user/relation/unblock |  |
| [**unfollow**](RelationService.md#unfollow) | **POST** /v1/user/relation/unfollow |  |



## block

> object block(blockRelationRequest)



当前账号拉黑目标账号。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { BlockRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // BlockRelationRequest
    blockRelationRequest: ...,
  } satisfies BlockRequest;

  try {
    const data = await api.block(body);
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


## follow

> object follow(followRelationRequest)



当前账号关注目标账号。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { FollowRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // FollowRelationRequest
    followRelationRequest: ...,
  } satisfies FollowRequest;

  try {
    const data = await api.follow(body);
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


## getStatus

> GetStatusRelationReply getStatus(getStatusRelationRequest)



查询当前账号与目标账号之间的关系。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { GetStatusRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // GetStatusRelationRequest
    getStatusRelationRequest: ...,
  } satisfies GetStatusRequest;

  try {
    const data = await api.getStatus(body);
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


## listBlocked

> ListBlockedRelationsReply listBlocked(listBlockedRelationsRequest)



分页查询当前账号拉黑的账号列表。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { ListBlockedRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // ListBlockedRelationsRequest
    listBlockedRelationsRequest: ...,
  } satisfies ListBlockedRequest;

  try {
    const data = await api.listBlocked(body);
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


## listFollowers

> ListFollowersRelationsReply listFollowers(listFollowersRelationsRequest)



分页查询当前账号的粉丝账号列表。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { ListFollowersRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // ListFollowersRelationsRequest
    listFollowersRelationsRequest: ...,
  } satisfies ListFollowersRequest;

  try {
    const data = await api.listFollowers(body);
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


## listFollowing

> ListFollowingRelationsReply listFollowing(listFollowingRelationsRequest)



分页查询当前账号关注的账号列表。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { ListFollowingRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // ListFollowingRelationsRequest
    listFollowingRelationsRequest: ...,
  } satisfies ListFollowingRequest;

  try {
    const data = await api.listFollowing(body);
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


## unblock

> object unblock(unblockRelationRequest)



当前账号取消拉黑目标账号。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { UnblockRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // UnblockRelationRequest
    unblockRelationRequest: ...,
  } satisfies UnblockRequest;

  try {
    const data = await api.unblock(body);
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


## unfollow

> object unfollow(unfollowRelationRequest)



当前账号取消关注目标账号。

### Example

```ts
import {
  Configuration,
  RelationService,
} from '@bass/bbs-sdk-fetch';
import type { UnfollowRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new RelationService();

  const body = {
    // UnfollowRelationRequest
    unfollowRelationRequest: ...,
  } satisfies UnfollowRequest;

  try {
    const data = await api.unfollow(body);
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

